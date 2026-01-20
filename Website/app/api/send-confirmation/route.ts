import { NextRequest, NextResponse } from 'next/server'
import { APP_NAME } from '@/lib/constants'

const RESEND_API_KEY = process.env.RESEND_API_KEY || ''

export async function POST(request: NextRequest) {
  try {
    // If Resend is not configured, skip sending email
    if (!RESEND_API_KEY) {
      console.log('Resend not configured. Confirmation email would be sent')
      return NextResponse.json(
        { message: 'Email service not configured' },
        { status: 200 }
      )
    }

    const { email } = await request.json()

    // Send email via Resend
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
    }
    
    if (RESEND_API_KEY) {
      headers['Authorization'] = `Bearer ${RESEND_API_KEY}`
    }
    
    const response = await fetch('https://api.resend.com/emails', {
      method: 'POST',
      headers,
      body: JSON.stringify({
        from: `${APP_NAME} <hello@unora.website>`, // Change this to your verified domain
        to: [email],
        subject: `Welcome to the ${APP_NAME} waitlist!`,
        html: `
          <!DOCTYPE html>
          <html>
            <head>
              <meta charset="utf-8">
              <meta name="viewport" content="width=device-width, initial-scale=1.0">
            </head>
            <body style="font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif; line-height: 1.6; color: #333; max-width: 600px; margin: 0 auto; padding: 20px;">
              <div style="text-align: center; margin-bottom: 30px;">
                <h1 style="color: #1a1a1a; font-size: 28px; font-weight: 300; margin: 0;">You're on the list!</h1>
              </div>
              
              <div style="background: #fafafa; padding: 30px; border-radius: 8px; margin-bottom: 20px;">
                <p style="font-size: 16px; color: #666; margin: 0 0 15px 0;">
                  Thanks for joining the ${APP_NAME} waitlist!
                </p>
                <p style="font-size: 16px; color: #666; margin: 0 0 15px 0;">
                  We're building something special, and we'll let you know as soon as we're ready.
                </p>
                <p style="font-size: 16px; color: #666; margin: 0;">
                  We’ll reach out only when there’s something worth sharing.
                </p>
              </div>
              
              <div style="text-align: center; margin-top: 30px; padding-top: 20px; border-top: 1px solid #eee;">
                <p style="font-size: 14px; color: #999; margin: 0;">
                  ${APP_NAME} - Some things are worth waiting for.
                </p>
              </div>
            </body>
          </html>
        `,
        text: `Thanks for joining the ${APP_NAME} waitlist! We're building something special and will let you know when we're ready.`,
      }),
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.message || 'Failed to send email')
    }

    return NextResponse.json(
      { message: 'Confirmation email sent' },
      { status: 200 }
    )
  } catch (error) {
    console.error('Email sending error:', error)
    // Don't fail the request if email fails
    return NextResponse.json(
      { message: 'Email service error (non-critical)' },
      { status: 200 }
    )
  }
}

