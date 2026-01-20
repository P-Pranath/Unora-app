'use client'

import { motion } from 'framer-motion'
import { useInView } from 'framer-motion'
import { forwardRef, useRef } from 'react'

interface SectionProps {
  children: React.ReactNode
  className?: string
  delay?: number
  id?: string
}

const Section = forwardRef<HTMLElement, SectionProps>(
  ({ children, className = '', delay = 0, id }, forwardedRef) => {
    const internalRef = useRef<HTMLElement>(null)
    const ref = (forwardedRef || internalRef) as React.RefObject<HTMLElement>
    const isInView = useInView(ref as React.RefObject<HTMLElement>, { once: false, margin: '-100px' })

  return (
    <motion.section
        ref={ref as any}
        id={id}
      className={className}
      initial={{ opacity: 0, y: 40 }}
      animate={isInView ? { opacity: 1, y: 0 } : { opacity: 0, y: 40 }}
      transition={{ duration: 0.8, delay, ease: [0.16, 1, 0.3, 1] }}
    >
      {children}
    </motion.section>
  )
}
)

Section.displayName = 'Section'

export default Section

