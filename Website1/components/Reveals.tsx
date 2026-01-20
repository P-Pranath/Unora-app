'use client'

import Section from './Section'
import { motion } from 'framer-motion'

const reveals = [
  {
    day: 'Day 3',
    title: 'Facts revealed gradually',
    description: 'Small details emerge as trust builds. No pressure, just natural progression.',
  },
  {
    day: 'Day 10',
    title: 'Identity revealed',
    description: 'Who you are comes after you\'ve shown who you are through consistency.',
  },
  {
    day: 'Day 15',
    title: 'Conversation opens',
    description: 'What you’ve built now has a voice.',
  },
]

export default function Reveals() {
  return (
    <Section className="py-24 sm:py-32 px-4 sm:px-6 lg:px-8">
      <div className="max-w-4xl mx-auto">
        <motion.h2
          className="text-3xl sm:text-4xl md:text-5xl font-light mb-8 text-center"
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6 }}
        >
          You don't unlock attention. You unlock reality.
        </motion.h2>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mt-16">
          {reveals.map((reveal, index) => (
            <motion.div
              key={index}
              className="p-6 sm:p-8 border border-neutral-800 rounded-2xl bg-neutral-900/30 hover:border-neutral-700 transition-colors"
              initial={{ opacity: 0, y: 30 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: false }}
              transition={{ duration: 0.6, delay: index * 0.15 }}
              whileHover={{ y: -4 }}
            >
              <div className="text-sm text-neutral-500 font-light mb-3">
                {reveal.day}
              </div>
              <h3 className="text-xl sm:text-2xl font-light mb-3">
                {reveal.title}
              </h3>
              <p className="text-neutral-400 font-light leading-relaxed text-sm sm:text-base">
                {reveal.description}
              </p>
            </motion.div>
          ))}
        </div>
      </div>
    </Section>
  )
}

