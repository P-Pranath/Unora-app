import { NextRequest, NextResponse } from 'next/server'
import { validateEmail } from '@/lib/emailValidation'

const SUPABASE_URL = process.env.NEXT_PUBLIC_SUPABASE_URL || ''
const SUPABASE_ANON_KEY = process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY || ''

export async function POST(request: NextRequest) {
  try {
    const { email } = await request.json()

    // Validate email
    const validation = validateEmail(email)
    if (!validation.valid) {
      return NextResponse.json(
        { error: validation.error },
        { status: 400 }
      )
    }

    // If Supabase is not configured, just return success (for development)
    if (!SUPABASE_URL || !SUPABASE_ANON_KEY) {
      console.log('Supabase not configured. Email would be saved:', email)
      return NextResponse.json(
        { message: 'Email received (Supabase not configured)' },
        { status: 200 }
      )
    }

    // Save to Supabase
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      'Prefer': 'return=minimal',
    }
    
    // Add authentication headers
    if (SUPABASE_ANON_KEY) {
      headers['apikey'] = SUPABASE_ANON_KEY
      headers['Authorization'] = `Bearer ${SUPABASE_ANON_KEY}`
    }
    
    const response = await fetch(`${SUPABASE_URL}/rest/v1/waitlist`, {
      method: 'POST',
      headers,
      body: JSON.stringify({
        email: email.toLowerCase().trim(),
        created_at: new Date().toISOString(),
      }),
    })

    if (!response.ok) {
      // Check if it's a duplicate
      if (response.status === 409) {
        return NextResponse.json(
          { error: 'This email is already on the waitlist' },
          { status: 409 }
        )
      }
      throw new Error('Failed to save email')
    }

    // Send confirmation email (via separate API route)
    try {
      await fetch(`${process.env.NEXT_PUBLIC_APP_URL || 'http://localhost:3000'}/api/send-confirmation`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email }),
      })
    } catch (emailError) {
      // Don't fail if email fails, just log it
      console.error('Failed to send confirmation email:', emailError)
    }

    return NextResponse.json(
      { message: 'Successfully joined the waitlist!' },
      { status: 200 }
    )
  } catch (error) {
    console.error('Waitlist API error:', error)
    return NextResponse.json(
      { error: 'Something went wrong. Please try again.' },
      { status: 500 }
    )
  }
}

