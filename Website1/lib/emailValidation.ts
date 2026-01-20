// Email validation utility
const commonDomains = [
  'gmail.com',
  'yahoo.com',
  'hotmail.com',
  'outlook.com',
  'icloud.com',
  'mail.com',
  'protonmail.com',
  'aol.com',
  'live.com',
  'msn.com',
  'yandex.com',
  'zoho.com',
  'gmx.com',
  'mail.ru',
  'qq.com',
]

export function validateEmail(email: string): { valid: boolean; error?: string } {
  // Basic checks
  if (!email || email.trim() === '') {
    return { valid: false, error: 'Email is required' }
  }

  // Check for @ symbol
  if (!email.includes('@')) {
    return { valid: false, error: 'Email must contain @ symbol' }
  }

  // Split email into local and domain parts
  const parts = email.split('@')
  if (parts.length !== 2) {
    return { valid: false, error: 'Invalid email format' }
  }

  const [localPart, domain] = parts

  // Check local part (before @)
  if (!localPart || localPart.length === 0) {
    return { valid: false, error: 'Email must have a username before @' }
  }

  if (localPart.length > 64) {
    return { valid: false, error: 'Email username is too long' }
  }

  // Check domain part (after @)
  if (!domain || domain.length === 0) {
    return { valid: false, error: 'Email must have a domain after @' }
  }

  // Check for valid domain format
  const domainRegex = /^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*\.[a-zA-Z]{2,}$/
  if (!domainRegex.test(domain)) {
    return { valid: false, error: 'Invalid domain format' }
  }

  // Check if domain is in common domains list (optional but helpful)
  const domainLower = domain.toLowerCase()
  const isCommonDomain = commonDomains.some(d => domainLower === d || domainLower.endsWith('.' + d))
  
  // Basic email regex check
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(email)) {
    return { valid: false, error: 'Invalid email format' }
  }

  return { valid: true }
}

