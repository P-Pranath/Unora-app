import {RfW, RfH} from '../utils/helpers';

/**
 * Pre-calculated responsive HEIGHT values
 */
export const HEIGHT = {
  H0: RfH(0),
  H1: RfH(1),
  H2: RfH(2),
  H4: RfH(4),
  H6: RfH(6),
  H8: RfH(8),
  H10: RfH(10),
  H12: RfH(12),
  H14: RfH(14),
  H16: RfH(16),
  H18: RfH(18),
  H20: RfH(20),
  H22: RfH(22),
  H24: RfH(24),
  H26: RfH(26),
  H28: RfH(28),
  H30: RfH(30),
  H32: RfH(32),
  H36: RfH(36),
  H40: RfH(40),
  H44: RfH(44),
  H48: RfH(48),
  H50: RfH(50),
  H52: RfH(52),
  H56: RfH(56),
  H60: RfH(60),
  H64: RfH(64),
  H72: RfH(72),
  H80: RfH(80),
  H96: RfH(96),
  H100: RfH(100),
  H120: RfH(120),
  H140: RfH(140),
  H160: RfH(160),
  H180: RfH(180),
  H200: RfH(200),
  H240: RfH(240),
  H280: RfH(280),
  H320: RfH(320),
  H400: RfH(400),
} as const;

/**
 * Pre-calculated responsive WIDTH values
 */
export const WIDTH = {
  W0: RfW(0),
  W1: RfW(1),
  W2: RfW(2),
  W4: RfW(4),
  W6: RfW(6),
  W8: RfW(8),
  W10: RfW(10),
  W12: RfW(12),
  W14: RfW(14),
  W16: RfW(16),
  W18: RfW(18),
  W20: RfW(20),
  W22: RfW(22),
  W24: RfW(24),
  W26: RfW(26),
  W28: RfW(28),
  W30: RfW(30),
  W32: RfW(32),
  W36: RfW(36),
  W40: RfW(40),
  W44: RfW(44),
  W48: RfW(48),
  W50: RfW(50),
  W56: RfW(56),
  W60: RfW(60),
  W64: RfW(64),
  W72: RfW(72),
  W80: RfW(80),
  W96: RfW(96),
  W100: RfW(100),
  W120: RfW(120),
  W140: RfW(140),
  W160: RfW(160),
  W200: RfW(200),
  W240: RfW(240),
  W280: RfW(280),
  W320: RfW(320),
  W327: RfW(327),
  W343: RfW(343),
  W375: RfW(375),
} as const;

/**
 * Border Radius constants (fixed values)
 */
export const BorderRadius = {
  BR0: 0,
  BR2: 2,
  BR4: 4,
  BR5: 5,
  BR6: 6,
  BR8: 8,
  BR10: 10,
  BR12: 12,
  BR14: 14,
  BR15: 15,
  BR16: 16,
  BR20: 20,
  BR24: 24,
  BR28: 28,
  BR32: 32,
  BRFull: 9999,
} as const;

/**
 * Component sizes (buttons, inputs, icons)
 */
export const componentSizes = {
  buttons: {
    sm: {
      height: RfH(32),
      paddingHorizontal: RfW(12),
    },
    md: {
      height: RfH(44),
      paddingHorizontal: RfW(16),
    },
    lg: {
      height: RfH(52),
      paddingHorizontal: RfW(24),
    },
  },
  inputs: {
    sm: {
      height: RfH(36),
      paddingHorizontal: RfW(12),
    },
    md: {
      height: RfH(44),
      paddingHorizontal: RfW(16),
    },
    lg: {
      height: RfH(52),
      paddingHorizontal: RfW(16),
    },
  },
  icons: {
    xs: 12,
    sm: 16,
    md: 20,
    lg: 24,
    xl: 32,
    '2xl': 40,
    '3xl': 48,
  },
} as const;

/**
 * Hit slop for touch targets (accessibility)
 */
export const hitSlop = {
  sm: {top: 8, bottom: 8, left: 8, right: 8},
  md: {top: 12, bottom: 12, left: 12, right: 12},
  lg: {top: 16, bottom: 16, left: 16, right: 16},
} as const;

export type Sizes = {
  HEIGHT: typeof HEIGHT;
  WIDTH: typeof WIDTH;
  BorderRadius: typeof BorderRadius;
  componentSizes: typeof componentSizes;
};
