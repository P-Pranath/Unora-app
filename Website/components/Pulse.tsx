'use client'

import Section from './Section'
import { motion, useScroll, useTransform, AnimatePresence } from 'framer-motion'
import { useRef, useState } from 'react'

// Reveal days - 4 reveals across 15 days
const REWARD_DAYS = [3, 6, 8, 14]

// Simulated current day for demo (can be adjusted)
const CURRENT_DAY = 7

type DayState = 'reveal' | 'no-reveal' | 'future'

function getDayState(day: number): DayState {
  if (day > CURRENT_DAY) return 'future'
  if (REWARD_DAYS.includes(day)) return 'reveal'
  return 'no-reveal'
}

function getDayMessage(day: number, state: DayState): { title: string; description: string } {
  switch (state) {
    case 'reveal':
      return {
        title: 'Something revealed',
        description: `Day ${day} — your shared pulse unlocked a new detail about them.`
      }
    case 'no-reveal':
      return {
        title: 'Pulse felt',
        description: `Day ${day} — no reveal today. The heartbeat continues.`
      }
    case 'future':
      return {
        title: 'Not yet reached',
        description: `Day ${day} — keep the pulse alive to reach this moment.`
      }
  }
}

export default function Pulse() {
  const ref = useRef<HTMLDivElement>(null)
  const [selectedDay, setSelectedDay] = useState<number | null>(null)

  const { scrollYProgress } = useScroll({
    target: ref,
    offset: ['start end', 'end start'],
  })

  const progress = useTransform(scrollYProgress, [0, 0.5, 0.59], [0.65, 0.65, 1])

  const days = Array.from({ length: 15 }, (_, i) => i + 1)

  const handleDayClick = (day: number) => {
    setSelectedDay(selectedDay === day ? null : day)
  }

  const handleClose = () => {
    setSelectedDay(null)
  }

  return (
    <Section className="py-24 sm:py-32 px-4 sm:px-6 lg:px-8" ref={ref}>
      <div className="max-w-4xl mx-auto">
        <motion.h2
          className="text-3xl sm:text-4xl md:text-5xl font-light mb-8 text-center"
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6 }}
        >
          One heart, two people
        </motion.h2>

        <motion.p
          className="text-lg sm:text-xl text-neutral-400 font-light text-center mb-16 max-w-2xl mx-auto leading-relaxed"
          initial={{ opacity: 0 }}
          whileInView={{ opacity: 1 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6, delay: 0.2 }}
        >
          Your pulse is the shared heartbeat of your connection. Each day you both show up, it grows stronger. Miss a day, and it fades.
        </motion.p>

        <div className="mb-12">
          <div className="flex justify-between items-center mb-4">
            <span className="text-sm text-neutral-500 font-light">Day 1</span>
            <span className="text-sm text-neutral-500 font-light">Day 15</span>
          </div>

          <div className="relative h-2 bg-neutral-800 rounded-full overflow-hidden">
            <motion.div
              className="absolute inset-y-0 left-0 bg-gradient-to-r from-neutral-600 via-neutral-400 to-neutral-200 rounded-full"
              style={{ width: useTransform(progress, (p) => `${p * 100}%`) }}
            />
          </div>
        </div>

        {/* Interactive Day Grid */}
        <div className="relative grid gap-2 mb-12 grid-cols-5 sm:grid-cols-10 lg:grid-cols-15">
          {days.map((day) => {
            const state = getDayState(day)
            const isRevealDay = REWARD_DAYS.includes(day)
            const isPast = day <= CURRENT_DAY
            const isSelected = selectedDay === day

            return (
              <motion.button
                key={day}
                onClick={() => handleDayClick(day)}
                className={`
                  aspect-square rounded-lg border flex items-center justify-center text-xs font-medium
                  cursor-pointer transition-all duration-200 relative
                  ${isSelected ? 'ring-2 ring-offset-2 ring-offset-neutral-900 ring-neutral-400' : ''}
                  ${isRevealDay && isPast
                    ? 'bg-neutral-700/80 border-neutral-500 text-neutral-200'
                    : isPast
                      ? 'bg-neutral-700 border-neutral-600 text-neutral-300'
                      : 'bg-neutral-800/50 border-neutral-700/50 text-neutral-500'
                  }
                  hover:scale-105 hover:border-neutral-500
                  active:scale-95
                `}
                initial={{ opacity: 0.3 }}
                whileInView={{ opacity: 1 }}
                viewport={{ once: false }}
                transition={{ duration: 0.3, delay: day * 0.02 }}
                whileHover={{ y: -2 }}
                whileTap={{ scale: 0.95 }}
              >
                {day}
                {isRevealDay && (
                  <span className="absolute -top-0.5 -right-0.5 w-1.5 h-1.5 rounded-full bg-neutral-400" />
                )}
              </motion.button>
            )
          })}
        </div>

        {/* Popup/Modal for selected day */}
        <AnimatePresence>
          {selectedDay !== null && (
            <motion.div
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: 10 }}
              transition={{ duration: 0.2 }}
              className="mb-12"
            >
              <div
                className={`
                  relative p-6 rounded-2xl border backdrop-blur-sm
                  ${getDayState(selectedDay) === 'reveal'
                    ? 'bg-neutral-800/60 border-neutral-600/40'
                    : getDayState(selectedDay) === 'future'
                      ? 'bg-neutral-800/50 border-neutral-700/50'
                      : 'bg-neutral-800/80 border-neutral-700'
                  }
                `}
              >
                <button
                  onClick={handleClose}
                  className="absolute top-4 right-4 text-neutral-400 hover:text-neutral-200 transition-colors"
                  aria-label="Close"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                    <line x1="18" y1="6" x2="6" y2="18"></line>
                    <line x1="6" y1="6" x2="18" y2="18"></line>
                  </svg>
                </button>

                <div className="flex items-start gap-4">
                  <div className={`
                    flex-shrink-0 w-12 h-12 rounded-xl flex items-center justify-center text-lg font-semibold
                    ${getDayState(selectedDay) === 'reveal'
                      ? 'bg-neutral-600/30 text-neutral-200'
                      : getDayState(selectedDay) === 'future'
                        ? 'bg-neutral-700/50 text-neutral-500'
                        : 'bg-neutral-700 text-neutral-300'
                    }
                  `}>
                    {selectedDay}
                  </div>
                  <div className="flex-1">
                    <h4 className={`
                      text-lg font-medium mb-1
                      ${getDayState(selectedDay) === 'reveal'
                        ? 'text-neutral-100'
                        : 'text-neutral-200'
                      }
                    `}>
                      {getDayMessage(selectedDay, getDayState(selectedDay)).title}
                    </h4>
                    <p className="text-neutral-400 text-sm font-light leading-relaxed">
                      {getDayMessage(selectedDay, getDayState(selectedDay)).description}
                    </p>
                  </div>
                </div>
              </div>
            </motion.div>
          )}
        </AnimatePresence>

        <motion.div
          className="text-center space-y-4"
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: false }}
          transition={{ duration: 0.6, delay: 0.4 }}
        >
          <p className="text-neutral-300 font-light leading-relaxed">
            Each day you both show up, your pulse grows stronger. The longer it beats, the more you reveal to each other.
          </p>
          <p className="text-neutral-400 text-sm font-light">
            AI nudges you gently when it's time to check in. Miss a day and your pulse weakens — but one check-in brings it back to life.
          </p>
        </motion.div>
      </div>
    </Section>
  )
}

