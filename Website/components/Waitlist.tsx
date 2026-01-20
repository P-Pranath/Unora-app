'use client'

import Section from './Section'
import { motion } from 'framer-motion'
import { useState } from 'react'
import { validateEmail } from '@/lib/emailValidation'

export default function Waitlist() {
  const [email, setEmail] = useState('')
  const [submitted, setSubmitted] = useState(false)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError(null)
    setSuccess(false)

    // Client-side validation
    const validation = validateEmail(email)
    if (!validation.valid) {
      setError(validation.error || 'Invalid email')
      return
    }

    setLoading(true)

    try {
      const response = await fetch('/api/waitlist', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email }),
      })

      const data = await response.json()

      if (!response.ok) {
        throw new Error(data.error || 'Something went wrong')
      }

      // Success
      setSuccess(true)
      setSubmitted(true)
      setEmail('')
      
      // Reset after 5 seconds
      setTimeout(() => {
        setSubmitted(false)
        setSuccess(false)
      }, 5000)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to join waitlist. Please try again.')
      setSubmitted(false)
    } finally {
      setLoading(false)
    }
  }

  return (
    <Section className="py-24 sm:py-32 px-4 sm:px-6 lg:px-8" id="waitlist">
      <div className="max-w-2xl mx-auto text-center">
        <motion.h2
          className="text-3xl sm:text-4xl md:text-5xl font-light mb-6"
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6 }}
        >
          Build something real.
        </motion.h2>

        <motion.form
          onSubmit={handleSubmit}
          className="mt-12 flex flex-col sm:flex-row gap-4 max-w-md mx-auto"
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6, delay: 0.2 }}
        >
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Enter your email"
            required
            className="flex-1 px-6 py-4 bg-neutral-900 border border-neutral-800 rounded-full text-white placeholder-neutral-500 focus:outline-none focus:border-neutral-600 transition-colors"
          />
          <motion.button
            type="submit"
            disabled={loading}
            className="px-8 py-4 bg-white text-black rounded-full font-medium hover:bg-neutral-200 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            whileHover={!loading ? { scale: 1.02 } : {}}
            whileTap={!loading ? { scale: 0.98 } : {}}
          >
            {loading ? 'Joining...' : submitted ? 'Joined ✓' : 'Join waitlist'}
          </motion.button>
        </motion.form>

        {/* Error message */}
        {error && (
          <motion.p
            className="mt-4 text-sm text-red-400 font-light"
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
          >
            {error}
          </motion.p>
        )}

        {/* Success message */}
        {success && (
          <motion.p
            className="mt-4 text-sm text-green-400 font-light"
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
          >
            Check your email for confirmation! ✓
          </motion.p>
        )}

        <motion.p
          className="mt-6 text-sm text-neutral-500 font-light"
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6, delay: 0.4 }}
        >
          No spam. No pressure.
        </motion.p>
      </div>
    </Section>
  )
}

