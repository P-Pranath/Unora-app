'use client'

import Section from './Section'
import { motion } from 'framer-motion'

const philosophy = [
  'We believe what you wait for matters.',
  'What you build lasts.',
  'And real connection isn’t rushed.',
]

export default function Why() {
  return (
    <Section className="py-24 sm:py-32 px-4 sm:px-6 lg:px-8">
      <div className="max-w-3xl mx-auto">
        <div className="space-y-8">
          {philosophy.map((line, index) => (
            <motion.p
              key={index}
              className="text-2xl sm:text-3xl md:text-4xl font-light text-center leading-relaxed"
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: false }}
              transition={{ duration: 0.8, delay: index * 0.2, ease: [0.16, 1, 0.3, 1] }}
            >
              {line}
            </motion.p>
          ))}
        </div>
      </div>
    </Section>
  )
}

