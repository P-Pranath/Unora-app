'use client'

import Section from './Section'
import { motion } from 'framer-motion'

const audiences = [
  'People who prefer intention over impulse',
  'People who value showing up, not showing off',
  'People who want connection to support personal growth',
]

export default function WhoItsFor() {
  return (
    <Section className="py-24 sm:py-32 px-4 sm:px-6 lg:px-8">
      <div className="max-w-3xl mx-auto text-center">
        <motion.h2
          className="text-3xl sm:text-4xl md:text-5xl font-light mb-12"
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6 }}
        >
          Who it's for
        </motion.h2>

        <div className="space-y-6">
          {audiences.map((audience, index) => (
            <motion.p
              key={index}
              className="text-xl sm:text-2xl text-neutral-300 font-light"
              initial={{ opacity: 0, x: -20 }}
              whileInView={{ opacity: 1, x: 0 }}
              viewport={{ once: false }}
              transition={{ duration: 0.6, delay: index * 0.1 }}
            >
              {audience}
            </motion.p>
          ))}
        </div>
      </div>
    </Section>
  )
}

