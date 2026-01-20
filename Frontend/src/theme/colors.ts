/**
 * Color palette for the Unora app
 */
export const colors = {
  // Primary colors
  primary: {
    50: '#FFF5F0',
    100: '#FFE6D9',
    200: '#FFCBB3',
    300: '#FFB08C',
    400: '#FF8E5C',
    500: '#FF6B35', // Main primary (Orange)
    600: '#E55A2B',
    700: '#CC4A21',
    800: '#B33A17',
    900: '#992A0D',
  },

  // Secondary colors
  secondary: {
    50: '#F5F3FF',
    100: '#EBE6FF',
    200: '#D4CCFF',
    300: '#BDB3FF',
    400: '#9C8CFF',
    500: '#7B68EE', // Main secondary (Purple)
    600: '#6A58D9',
    700: '#5948C4',
    800: '#4838AF',
    900: '#37289A',
  },

  // Neutral colors
  neutral: {
    0: '#FFFFFF',
    50: '#F9FAFB',
    100: '#F3F4F6',
    200: '#E5E7EB',
    300: '#D1D5DB',
    400: '#9CA3AF',
    500: '#6B7280',
    600: '#4B5563',
    700: '#374151',
    800: '#1F2937',
    900: '#111827',
    950: '#0A0A0F', // Dark background
  },

  // Semantic colors
  success: {
    light: '#D1FAE5',
    main: '#10B981',
    dark: '#059669',
  },

  warning: {
    light: '#FEF3C7',
    main: '#F59E0B',
    dark: '#D97706',
  },

  error: {
    light: '#FEE2E2',
    main: '#EF4444',
    dark: '#DC2626',
  },

  info: {
    light: '#DBEAFE',
    main: '#3B82F6',
    dark: '#2563EB',
  },

  // Background colors
  background: {
    primary: '#0A0A0F',
    secondary: '#111827',
    tertiary: '#1F2937',
    surface: '#FFFFFF',
    surfaceSecondary: '#F9FAFB',
  },

  // Text colors
  text: {
    primary: '#FFFFFF',
    secondary: '#9CA3AF',
    tertiary: '#6B7280',
    inverse: '#111827',
    muted: '#4B5563',
  },

  // Border colors
  border: {
    light: '#E5E7EB',
    default: '#D1D5DB',
    dark: '#374151',
    focus: '#FF6B35',
  },

  // Onboarding/Design colors
  onboarding: {
    primary: '#C4A77D',
    primaryHover: '#B3966D',
    backgroundLight: '#FAFAF8',
    surfaceCard: '#FFFFFF',
    textHead: '#1A1A1A',
    textBody: '#4A4A4A',
    dotInactive: '#D4D4D0',
    cardInnerBack: '#F5F5F4',
    cardInnerFront: '#FAFAF8',
    skeleton: '#E7E5E4',
  },
} as const;

export type Colors = typeof colors;
