import Hero from '@/components/Hero'
import Problem from '@/components/Problem'
import Solution from '@/components/Solution'
import Pulse from '@/components/Pulse'
import Reveals from '@/components/Reveals'
import WhoItsFor from '@/components/WhoItsFor'
import Why from '@/components/Why'
import Waitlist from '@/components/Waitlist'
import { APP_NAME, APP_DESCRIPTION, APP_URL } from '@/lib/constants'

export default function Home() {
  // Structured Data (JSON-LD) for better SEO
  const jsonLd = {
    '@context': 'https://schema.org',
    '@type': 'WebApplication',
    name: APP_NAME,
    description: APP_DESCRIPTION,
    url: APP_URL,
    applicationCategory: 'LifestyleApplication',
    operatingSystem: 'Web',
    logo: `${APP_URL}logo.png`,
    offers: {
      '@type': 'Offer',
      price: '0',
      priceCurrency: 'USD',
    },
    // aggregateRating: {
    // '@type': 'AggregateRating',
    // Add when you have reviews
    // ratingValue: '4.5',
    // reviewCount: '100',
    // },
  }

  // Organization schema for brand recognition (helps Google show your logo)
  const organizationJsonLd = {
    '@context': 'https://schema.org',
    '@type': 'Organization',
    name: APP_NAME,
    url: APP_URL,
    logo: `${APP_URL}logo.png`,
  }

  // WebSite schema - critical for Google to recognize and display your site name
  const websiteJsonLd = {
    '@context': 'https://schema.org',
    '@type': 'WebSite',
    name: APP_NAME,
    url: APP_URL,
  }

  return (
    <>
      {/* Structured Data for SEO */}
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(organizationJsonLd) }}
      />
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(websiteJsonLd) }}
      />

      <main className="min-h-screen">
        <Hero />
        <Problem />
        <Solution />
        <Pulse />
        <Reveals />
        <WhoItsFor />
        <Why />
        <Waitlist />
      </main>
    </>
  )
}

