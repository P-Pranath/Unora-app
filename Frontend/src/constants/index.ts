import {Dimensions} from 'react-native';

// Standard Reference Screen: iPhone X (375 x 812 pixels)
export const STANDARD_SCREEN_SIZE = 812;
export const STANDARD_SCREEN_DIMENSIONS = {height: 812, width: 375};

// Current device dimensions
export const SCREEN_WIDTH = Dimensions.get('window').width;
export const SCREEN_HEIGHT = Dimensions.get('window').height;

// API Configuration
export const API_CONFIG = {
  TIMEOUT: 15000, // 15s timeout for mobile
  RETRY_COUNT: 1,
};

// Storage Keys
export const STORAGE_KEYS = {
  AUTH: 'auth',
  UI: 'ui',
  SETTINGS: 'settings',
} as const;

// Route Names
export const ROUTES = {
  // Auth Stack
  AUTH: {
    LOGIN: 'Login',
    REGISTER: 'Register',
    FORGOT_PASSWORD: 'ForgotPassword',
  },
  // Main Stack
  MAIN: {
    HOME: 'Home',
    PROFILE: 'Profile',
    SETTINGS: 'Settings',
  },
  // Tab Navigator
  TABS: {
    HOME_TAB: 'HomeTab',
    PROFILE_TAB: 'ProfileTab',
  },
} as const;
