'use client'

import { motion } from 'framer-motion'
import ScrollIndicator from './ScrollIndicator'
import { APP_NAME, APP_DESCRIPTION } from '@/lib/constants'

export default function Hero() {
  return (
    <section className="min-h-screen flex flex-col items-center justify-center px-4 sm:px-6 lg:px-8 relative">
      <div className="max-w-4xl mx-auto text-center">
        <motion.h1
          className="text-4xl sm:text-5xl md:text-6xl lg:text-7xl font-light tracking-tight mb-6"
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, ease: [0.16, 1, 0.3, 1] }}
        >
          <span className="text-white">{APP_NAME}</span>
          <span className="block text-2xl sm:text-3xl md:text-4xl lg:text-5xl text-neutral-400 mt-4">
            Some things are worth waiting for.
          </span>
        </motion.h1>

        <motion.p
          className="text-lg sm:text-xl md:text-2xl text-neutral-400 font-light max-w-2xl mx-auto mb-12 leading-relaxed"
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, delay: 0.2, ease: [0.16, 1, 0.3, 1] }}
        >
          A platform for building meaningful connections through shared commitment.
        </motion.p>

        <motion.div
          className="flex flex-col sm:flex-row gap-4 justify-center items-center"
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, delay: 0.4, ease: [0.16, 1, 0.3, 1] }}
        >
          <motion.a
            href="#waitlist"
            className="px-8 py-3 bg-white text-black rounded-full font-medium hover:bg-neutral-200 transition-colors"
            whileHover={{ scale: 1.02 }}
            whileTap={{ scale: 0.98 }}
          >
            Join the waitlist
          </motion.a>

          <motion.a
            href="#how-it-works"
            className="px-8 py-3 border border-neutral-700 rounded-full font-medium hover:border-neutral-500 transition-colors"
            whileHover={{ scale: 1.02 }}
            whileTap={{ scale: 0.98 }}
          >
            See how it works
          </motion.a>
        </motion.div>
      </div>

      <ScrollIndicator />
    </section>
  )
}

