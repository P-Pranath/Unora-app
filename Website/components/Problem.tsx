'use client'

import Section from './Section'
import { motion } from 'framer-motion'

const problems = [
  {
    icon: '🕰️',
    title: 'Consistency matters more than speed',
    description: 'Real connection forms when people show up repeatedly - not instantly.',
  },
  {
    icon: '🧠',
    title: 'Attention is fragile',
    description: 'Endless choice and constant distraction make it hard to stay present with anyone.',
  },
  {
    icon: '🌱',
    title: 'Connection should support growth',
    description: 'The right systems help people stay focused on themselves while building something with someone else.',
  },
]

export default function Problem() {
  return (
    <Section className="py-24 sm:py-32 px-4 sm:px-6 lg:px-8">
      <div className="max-w-4xl mx-auto">
        <motion.h2
          className="text-3xl sm:text-4xl md:text-5xl font-light mb-16 text-center"
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6 }}
        >
          What we believe about connection
        </motion.h2>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-8 md:gap-12">
          {problems.map((problem, index) => (
            <motion.div
              key={index}
              className="text-center"
              initial={{ opacity: 0, y: 30 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: false }}
              transition={{ duration: 0.6, delay: index * 0.15 }}
            >
              <div className="text-4xl mb-4 opacity-60">{problem.icon}</div>
              <h3 className="text-lg sm:text-xl text-neutral-300 font-medium mb-2">
                {problem.title}
              </h3>
              <p className="text-base sm:text-lg text-neutral-400 font-light">
                {problem.description}
              </p>
            </motion.div>
          ))}
        </div>
      </div>
    </Section>
  )
}

