'use client'

import Section from './Section'
import { motion } from 'framer-motion'
import { useInView } from 'framer-motion'
import { useRef } from 'react'

const steps = [
  {
    number: '01',
    title: 'Choose who you want to build with',
    description: 'Connect through interests, habits, and how someone shows up. Thoughtful AI helps surface connections that grow more meaningful over time.',
  },
  {
    number: '02',
    title: 'Build something real together',
    description: 'A shared pulse builds accountability — with each other and with yourself — turning consistency into focus, discipline, and real connection.',
  },
  {
    number: '03',
    title: 'Earn discovery, one layer at a time',
    description: 'Each milestone unlocks something real - facts, traits, details. Curiosity turns into clarity.',
  },
]

function StepItem({ step, index }: { step: typeof steps[0]; index: number }) {
  const ref = useRef(null)
  const isInView = useInView(ref, { once: false, margin: '-50px' })

  return (
    <motion.div
      ref={ref}
      className="flex gap-6 md:gap-8"
      initial={{ opacity: 0, x: -30 }}
      animate={isInView ? { opacity: 1, x: 0 } : { opacity: 0, x: -30 }}
      transition={{ duration: 0.8, delay: index * 0.2, ease: [0.16, 1, 0.3, 1] }}
    >
      <div className="flex-shrink-0">
        <div className="text-2xl sm:text-3xl font-light text-neutral-500 mb-2">
          {step.number}
        </div>
        {index < steps.length - 1 && (
          <div className="w-px h-full bg-gradient-to-b from-neutral-700 to-transparent ml-3 mt-2" />
        )}
      </div>
      <div className="pb-12">
        <h3 className="text-xl sm:text-2xl font-light mb-3">{step.title}</h3>
        <p className="text-neutral-400 font-light leading-relaxed">
          {step.description}
        </p>
      </div>
    </motion.div>
  )
}

export default function Solution() {
  return (
    <Section className="py-24 sm:py-32 px-4 sm:px-6 lg:px-8" id="how-it-works">
      <div className="max-w-3xl mx-auto">
        <motion.h2
          className="text-3xl sm:text-4xl md:text-5xl font-light mb-16 text-center"
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6 }}
        >
          We redesigned how connection begins.
        </motion.h2>

        <div className="space-y-0">
          {steps.map((step, index) => (
            <StepItem key={index} step={step} index={index} />
          ))}
        </div>
      </div>
    </Section>
  )
}

