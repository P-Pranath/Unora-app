import {RFValue} from 'react-native-responsive-fontsize';
import {STANDARD_SCREEN_SIZE} from '../constants';

/**
 * Font families - update these with your custom fonts
 * Note: 'display' and 'body' are aliases for the design system
 * In production, these should be set to 'Outfit' and 'Inter' respectively
 */
export const fontFamily = {
  regular: 'System',
  medium: 'System',
  semibold: 'System',
  bold: 'System',
  display: 'System', // Should be 'Outfit' when font is loaded
  body: 'System', // Should be 'Inter' when font is loaded
} as const;

/**
 * Responsive font sizes based on standard screen size (812px height)
 */
export const fontSize = {
  // Extra small
  xs: RFValue(10, STANDARD_SCREEN_SIZE),
  // Small
  sm: RFValue(12, STANDARD_SCREEN_SIZE),
  // Base
  base: RFValue(14, STANDARD_SCREEN_SIZE),
  // Medium
  md: RFValue(16, STANDARD_SCREEN_SIZE),
  // Large
  lg: RFValue(18, STANDARD_SCREEN_SIZE),
  // Extra large
  xl: RFValue(20, STANDARD_SCREEN_SIZE),
  // 2xl
  '2xl': RFValue(24, STANDARD_SCREEN_SIZE),
  // 3xl
  '3xl': RFValue(30, STANDARD_SCREEN_SIZE),
  // 4xl
  '4xl': RFValue(36, STANDARD_SCREEN_SIZE),
  // 5xl
  '5xl': RFValue(48, STANDARD_SCREEN_SIZE),
  // 6xl
  '6xl': RFValue(64, STANDARD_SCREEN_SIZE),
} as const;

/**
 * Font size numeric map for dynamic access
 */
export const FontSize: {[key: number]: number} = {
  10: RFValue(10, STANDARD_SCREEN_SIZE),
  11: RFValue(11, STANDARD_SCREEN_SIZE),
  12: RFValue(12, STANDARD_SCREEN_SIZE),
  13: RFValue(13, STANDARD_SCREEN_SIZE),
  14: RFValue(14, STANDARD_SCREEN_SIZE),
  15: RFValue(15, STANDARD_SCREEN_SIZE),
  16: RFValue(16, STANDARD_SCREEN_SIZE),
  17: RFValue(17, STANDARD_SCREEN_SIZE),
  18: RFValue(18, STANDARD_SCREEN_SIZE),
  19: RFValue(19, STANDARD_SCREEN_SIZE),
  20: RFValue(20, STANDARD_SCREEN_SIZE),
  21: RFValue(21, STANDARD_SCREEN_SIZE),
  22: RFValue(22, STANDARD_SCREEN_SIZE),
  23: RFValue(23, STANDARD_SCREEN_SIZE),
  24: RFValue(24, STANDARD_SCREEN_SIZE),
  25: RFValue(25, STANDARD_SCREEN_SIZE),
  26: RFValue(26, STANDARD_SCREEN_SIZE),
  28: RFValue(28, STANDARD_SCREEN_SIZE),
  30: RFValue(30, STANDARD_SCREEN_SIZE),
  32: RFValue(32, STANDARD_SCREEN_SIZE),
  36: RFValue(36, STANDARD_SCREEN_SIZE),
  40: RFValue(40, STANDARD_SCREEN_SIZE),
  48: RFValue(48, STANDARD_SCREEN_SIZE),
  64: RFValue(64, STANDARD_SCREEN_SIZE),
};

/**
 * Font weights
 */
export const fontWeight = {
  regular: '400' as const,
  medium: '500' as const,
  semibold: '600' as const,
  bold: '700' as const,
  extrabold: '800' as const,
};

/**
 * Line heights (multiplier based on font size)
 */
export const lineHeight = {
  tight: 1.1,
  normal: 1.4,
  relaxed: 1.6,
  loose: 1.8,
} as const;

/**
 * Letter spacing
 */
export const letterSpacing = {
  tighter: -0.5,
  tight: -0.25,
  normal: 0,
  wide: 0.25,
  wider: 0.5,
  widest: 1,
} as const;

export type Typography = {
  fontFamily: typeof fontFamily;
  fontSize: typeof fontSize;
  fontWeight: typeof fontWeight;
  lineHeight: typeof lineHeight;
  letterSpacing: typeof letterSpacing;
};

export const typography: Typography = {
  fontFamily,
  fontSize,
  fontWeight,
  lineHeight,
  letterSpacing,
};
