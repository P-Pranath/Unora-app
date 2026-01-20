'use client'

import { APP_NAME } from '@/lib/constants'
import Image from 'next/image' // Uncomment when adding your logo

export default function Header() {
  return (
    <header className="fixed top-0 left-0 right-0 z-50 pointer-events-none">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-6">
        <div className="flex items-center gap-2 pointer-events-auto">
          {/* Logo placeholder - replace with your actual logo */}
          <div className="w-10 h-10 rounded bg-neutral-700/50 flex items-center justify-center">
            <Image
              src="/logo.png"
              alt="Unora Logo"
              width={64}
              height={64}
              quality={100}
              priority
              className="w-[26px] h-[26px]"
            />
          </div>

          {/* App name */}
          <span className="text-sm text-neutral-500 font-light tracking-wide">
            {APP_NAME}
          </span>
        </div>
      </div>
    </header>
  )
}

