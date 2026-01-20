import {Dimensions} from 'react-native';
import DeviceInfo from 'react-native-device-info';
import {STANDARD_SCREEN_DIMENSIONS} from '../constants';

/**
 * Get current device window width
 */
export const deviceWidth = (): number => {
  const dim = Dimensions.get('window');
  return dim.width;
};

/**
 * Get current device window height
 */
export const deviceHeight = (): number => {
  const dim = Dimensions.get('window');
  return dim.height;
};

/**
 * Get full screen height (includes system UI)
 */
export const deviceHeightScreen = (): number => {
  const dim = Dimensions.get('screen');
  return dim.height;
};

/**
 * Check if device has a notch (iPhone X and later)
 */
export const isDisplayWithNotch = (): boolean => {
  return DeviceInfo.hasNotch();
};

/**
 * Responsive Width - Scales width values proportionally based on device width
 * Use for: padding (horizontal), margin (horizontal), width, left/right positioning
 *
 * @param value - The width value based on standard screen (375px)
 * @returns Scaled width value for current device
 */
export const RfW = (value: number): number => {
  const dim = Dimensions.get('window');
  return dim.width * (value / STANDARD_SCREEN_DIMENSIONS.width);
};

/**
 * Responsive Height - Scales height values proportionally based on device height
 * Use for: padding (vertical), margin (vertical), height, top/bottom positioning
 *
 * @param value - The height value based on standard screen (812px)
 * @returns Scaled height value for current device
 */
export const RfH = (value: number): number => {
  const dim = Dimensions.get('window');
  return dim.height * (value / STANDARD_SCREEN_DIMENSIONS.height);
};

/**
 * Clamp a value between min and max
 */
export const clamp = (value: number, min: number, max: number): number => {
  return Math.min(Math.max(value, min), max);
};

/**
 * Format a number as currency
 */
export const formatCurrency = (
  value: number,
  currency = 'USD',
  locale = 'en-US',
): string => {
  return new Intl.NumberFormat(locale, {
    style: 'currency',
    currency,
  }).format(value);
};

/**
 * Debounce function
 */
export const debounce = <T extends (...args: unknown[]) => unknown>(
  func: T,
  wait: number,
): ((...args: Parameters<T>) => void) => {
  let timeoutId: ReturnType<typeof setTimeout> | null = null;

  return (...args: Parameters<T>) => {
    if (timeoutId) {
      clearTimeout(timeoutId);
    }
    timeoutId = setTimeout(() => {
      func(...args);
    }, wait);
  };
};

/**
 * Throttle function
 */
export const throttle = <T extends (...args: unknown[]) => unknown>(
  func: T,
  limit: number,
): ((...args: Parameters<T>) => void) => {
  let inThrottle = false;

  return (...args: Parameters<T>) => {
    if (!inThrottle) {
      func(...args);
      inThrottle = true;
      setTimeout(() => {
        inThrottle = false;
      }, limit);
    }
  };
};

/**
 * Check if a string is a valid email
 */
export const isValidEmail = (email: string): boolean => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
};

/**
 * Capitalize first letter of a string
 */
export const capitalizeFirst = (str: string): string => {
  if (!str) return '';
  return str.charAt(0).toUpperCase() + str.slice(1);
};

/**
 * Truncate string with ellipsis
 */
export const truncateString = (str: string, maxLength: number): string => {
  if (str.length <= maxLength) return str;
  return str.slice(0, maxLength - 3) + '...';
};
