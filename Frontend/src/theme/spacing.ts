import {RfW, RfH} from '../utils/helpers';

/**
 * Spacing scale for consistent margins and paddings
 * Based on 4px base unit
 */
export const spacing = {
  // 0px
  none: 0,
  // 4px
  xs: 4,
  // 8px
  sm: 8,
  // 12px
  md: 12,
  // 16px
  lg: 16,
  // 20px
  xl: 20,
  // 24px
  '2xl': 24,
  // 32px
  '3xl': 32,
  // 40px
  '4xl': 40,
  // 48px
  '5xl': 48,
  // 64px
  '6xl': 64,
} as const;

/**
 * Responsive horizontal spacing (based on width)
 */
export const horizontalSpacing = {
  xs: RfW(4),
  sm: RfW(8),
  md: RfW(12),
  lg: RfW(16),
  xl: RfW(20),
  '2xl': RfW(24),
  '3xl': RfW(32),
  '4xl': RfW(40),
} as const;

/**
 * Responsive vertical spacing (based on height)
 */
export const verticalSpacing = {
  xs: RfH(4),
  sm: RfH(8),
  md: RfH(12),
  lg: RfH(16),
  xl: RfH(20),
  '2xl': RfH(24),
  '3xl': RfH(32),
  '4xl': RfH(40),
} as const;

/**
 * Screen padding (standard content padding)
 */
export const screenPadding = {
  horizontal: RfW(24),
  vertical: RfH(16),
} as const;

export type Spacing = typeof spacing;
