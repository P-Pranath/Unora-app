# Source Code Structure

This document explains the organization of the `src/` directory and the conventions used throughout the codebase.

## 📁 Folder Overview

```
src/
├── app/                    # Application initialization and providers
├── features/               # Feature modules (main business logic)
├── components/             # Shared UI components (atomic design)
├── navigation/             # Navigation configuration
├── services/               # External services and utilities
├── hooks/                  # Shared custom React hooks
├── stores/                 # Global state management (Zustand)
├── theme/                  # Design system and styling
├── utils/                  # Utility functions
├── types/                  # TypeScript type definitions
├── assets/                 # Static assets (images, fonts, icons)
└── constants/              # Application constants
```

## 🎯 Core Principles

### 1. Feature-First Architecture

Each feature is self-contained with its own:

- Screens
- Components
- API hooks (React Query)
- Business logic hooks
- Local state (Zustand store if needed)
- TypeScript types

### 2. State Management Strategy

**Server State (React Query)**:

- All API calls
- Data fetching and caching
- Background synchronization

**Client State (Zustand)**:

- Authentication state
- App preferences (theme, language)
- UI state (modals, drawers)

**Local Component State**:

- Form inputs
- UI interactions
- Transient state

### 3. Component Organization (Atomic Design)

**Atoms** (`components/atoms/`):

- Basic building blocks (Button, Input, Text, Icon)
- Cannot be broken down further
- Highly reusable

**Molecules** (`components/molecules/`):

- Combinations of atoms (FormField, Card, SearchBar)
- Serve a single purpose
- Reusable across features

**Organisms** (`components/organisms/`):

- Complex UI sections (Header, Modal, Form)
- Can contain molecules and atoms
- Feature-agnostic

**Layout** (`components/layout/`):

- Screen-level wrappers (Screen, Content, ScrollView)
- Handle safe areas and responsive behavior

**Feedback** (`components/feedback/`):

- Loading states, errors, toasts
- User feedback components

### 4. Responsive Design

The app uses a **standard screen dimension-based scaling system** optimized for mobile portrait-only applications.

#### Core Philosophy

1. **Standard Reference Screen**: iPhone X (375 x 812 pixels) is used as the standard reference screen
2. **Proportional Scaling**: All dimensions scale proportionally based on the ratio between actual device dimensions and standard dimensions
3. **Function-Based Approach**: All responsive sizing is done through utility functions rather than hardcoded values
4. **Font Scaling**: Separate responsive font sizing using `react-native-responsive-fontsize`
5. **Mobile Portrait Only**: This approach is optimized for mobile devices in portrait orientation only

#### Standard Screen Dimensions

**Location**: `src/constant/index.ts` and `src/utils/constants.ts`

```typescript
export const STANDARD_SCREEN_SIZE = 812;
export const STANDARD_SCREEN_DIMENSIONS = {height: 812, width: 375};
```

All responsive calculations are based on these standard dimensions (iPhone X).

#### Responsive Functions

**Primary Functions** (`src/utils/helper.tsx` and `src/utils/helpers.ts`):

1. **`RfW(value: number)`** - Responsive Width

   - Scales width values proportionally based on device width
   - Formula: `SCREEN_WIDTH * (value / STANDARD_SCREEN_DIMENSIONS.width)`
   - Use for: padding (horizontal), margin (horizontal), width, left/right positioning

2. **`RfH(value: number)`** - Responsive Height
   - Scales height values proportionally based on device height
   - Formula: `SCREEN_HEIGHT * (value / STANDARD_SCREEN_DIMENSIONS.height)`
   - Use for: padding (vertical), margin (vertical), height, top/bottom positioning

**Usage**:

```typescript
import {RfW, RfH} from '@/utils/helper';

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: RfW(24),
    paddingVertical: RfH(20),
    width: RfW(327),
    height: RfH(240),
  },
});
```

#### Font Sizing

**Package**: `react-native-responsive-fontsize`

**Location**: `src/theme/sizes.ts`

All font sizes use `RFValue` function:

```typescript
import {RFValue} from 'react-native-responsive-fontsize';
import {STANDARD_SCREEN_SIZE} from '../constant';

export const FontSize = {
  10: RFValue(10, STANDARD_SCREEN_SIZE),
  14: RFValue(14, STANDARD_SCREEN_SIZE),
  16: RFValue(16, STANDARD_SCREEN_SIZE),
  18: RFValue(18, STANDARD_SCREEN_SIZE),
  // ... more sizes
};

// Usage
const styles = StyleSheet.create({
  text: {
    fontSize: FontSize[14],
  },
});
```

#### Theme Constants

**Location**: `src/theme/sizes.ts`

Pre-calculated responsive values for consistency:

```typescript
export const HEIGHT = {
  H0: RfH(0),
  H1: RfH(1),
  H20: RfH(20),
  H24: RfH(24),
  H240: RfH(240),
  // ... more values
};

export const WIDTH = {
  W3: RfW(3),
  W16: RfW(16),
  W24: RfW(24),
  W327: RfW(327),
  // ... more values
};

export const BorderRadius = {
  BR0: 0,
  BR5: 5,
  BR10: 10,
  BR12: 12,
  BR15: 15,
  BR20: 20,
};
```

**Usage with Theme Constants**:

```typescript
import {HEIGHT, WIDTH, FontSize, BorderRadius} from '@/theme/sizes';

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: WIDTH.W24,
    paddingVertical: HEIGHT.H20,
    width: WIDTH.W327,
    borderRadius: BorderRadius.BR12,
  },
  text: {
    fontSize: FontSize[14],
  },
});
```

#### Best Practices

**✅ Always use responsive functions**:

```typescript
// ✅ Good
paddingHorizontal: RfW(24),
fontSize: FontSize[14],
paddingHorizontal: WIDTH.W24,  // Prefer constants when available

// ❌ Bad
paddingHorizontal: 24,  // Fixed value
fontSize: 14,           // Fixed value
```

**✅ Use RfW for horizontal spacing** (padding, margin, width, left/right)  
**✅ Use RfH for vertical spacing** (padding, margin, height, top/bottom)  
**✅ Use FontSize constants or RFValue for all font sizes**  
**✅ Use theme constants (HEIGHT, WIDTH) when available for consistency**  
**✅ Border radius values are typically fixed** (use BorderRadius constants)

**For more details**, see the [Responsive Guide](../responsiveness/RESPONSIVE_GUIDE.md).

## 📂 Detailed Folder Descriptions

### `app/`

Application initialization and global providers.

**Files**:

- `AppProviders.tsx` - All context providers (React Query, Navigation, Theme)
- `RootNavigator.tsx` - Root navigation component
- `config/` - Firebase, Sentry, and other configurations

### `features/`

Feature modules following consistent structure.

**Standard Feature Structure**:

```
features/[feature-name]/
├── screens/              # Screen components
├── components/           # Feature-specific components
├── api/                  # React Query hooks
├── hooks/                # Feature-specific hooks
├── store/                # Zustand store (if needed)
├── types/                # TypeScript types
└── index.ts              # Public API
```

**Examples**:

- `features/auth/` - Authentication flows
- `features/home/` - Home dashboard
- `features/approvals/` - Approvals management
- `features/tenant/` - Tenant-specific features

### `components/`

Shared components using atomic design pattern.

**Naming Convention**:

- Components: PascalCase (e.g., `Button.tsx`)
- Folders: PascalCase (e.g., `Button/`)
- Files: camelCase for utilities (e.g., `button.styles.ts`)

**Component Structure**:

```
Button/
├── Button.tsx           # Component implementation
├── Button.styles.ts     # StyleSheet
├── Button.types.ts      # Props interface
├── Button.test.tsx      # Tests
└── index.ts             # Public exports
```

### `navigation/`

Navigation configuration using React Navigation.

**Structure**:

- `RootNavigator.tsx` - Root navigation
- `stacks/` - Stack navigators (Auth, Main, Tenant)
- `tabs/` - Tab navigators
- `linking.ts` - Deep linking configuration
- `navigationRef.ts` - Imperative navigation reference
- `types.ts` - Navigation TypeScript types

### `services/`

External services and integrations.

**Key Services**:

- `api/` - API client, interceptors, React Query configuration
- `monitoring/` - Sentry, logger, error handler (already implemented)
- `storage/` - MMKV and secure storage wrappers
- `notifications/` - Push notifications (FCM)
- `analytics/` - Firebase Analytics
- `location/` - Location services

### `hooks/`

Shared custom React hooks.

**Examples**:

- `useResponsiveLayout.ts` - Responsive breakpoint decisions
- `useScreenLoadTime.ts` - Performance monitoring (already implemented)
- `useNetworkStatus.ts` - Network connectivity
- `useKeyboard.ts` - Keyboard events
- `useDebounce.ts` - Debounce utility
- `useThrottle.ts` - Throttle utility

### `stores/`

Global Zustand stores (client-side state only).

**Guidelines**:

- Keep stores small and focused
- Use React Query for server state
- Only use Zustand for client state

**Example Stores**:

- `authStore.ts` - User, tokens, auth status
- `appStore.ts` - Theme, language, preferences
- `notificationStore.ts` - Notification state

### `theme/`

Centralized design system.

**Files**:

- `colors.ts` - Color palette
- `typography.ts` - Font system (sizes, weights, line heights)
- `spacing.ts` - Spacing scale (xs, sm, md, lg, xl)
- `sizes.ts` - Component sizes (buttons, inputs, icons)
- `breakpoints.ts` - Responsive breakpoints
- `shadows.ts` - Shadow styles
- `animations.ts` - Animation configurations
- `ThemeProvider.tsx` - Theme context

### `utils/`

Utility functions.

**Key Utilities**:

- `helper.tsx` / `helpers.ts` - RfW, RfH responsive functions
- `font.ts` - Responsive font sizing
- `layout.ts` - Layout calculations
- `image.ts` - Image utilities
- `date.ts` - Date formatting
- `string.ts` - String manipulation
- `validation.ts` - Form validation
- `format.ts` - Data formatting

### `types/`

TypeScript type definitions.

**Files**:

- `api.types.ts` - API request/response types
- `navigation.types.ts` - Navigation types
- `user.types.ts` - User-related types
- `common.types.ts` - Common utility types
- `monitoring.ts` - Monitoring types (already created)

### `assets/`

Static assets organized by category.

**Structure**:

- `images/` - PNG/JPG images (organized by feature)
- `icons/` - SVG icons
- `fonts/` - Font files
- `lottie/` - Lottie animation JSON files

### `constants/`

Application-wide constants.

**Files**:

- `api.ts` - API endpoint constants
- `storage.ts` - Storage key constants
- `routes.ts` - Route name constants
- `config.ts` - App configuration

## 🏗️ Creating a New Feature

1. **Create feature folder**:

```bash
mkdir -p src/features/[feature-name]/{screens,components,api,hooks,types}
```

2. **Create index.ts** to export public API:

```typescript
// features/[feature-name]/index.ts
export * from './screens';
export * from './components';
export * from './hooks';
```

3. **Create screens** in `screens/` folder

4. **Create API hooks** in `api/`:

```typescript
// api/[feature]Api.ts
export const useFeatureData = () => {
  return useQuery({
    queryKey: ['feature'],
    queryFn: () => apiClient.get('/feature'),
  });
};
```

5. **Create custom hooks** in `hooks/`:

```typescript
// hooks/useFeature.ts
export const useFeature = () => {
  const {data} = useFeatureData();
  // Business logic
  return {
    /* ... */
  };
};
```

6. **Add navigation** in `navigation/stacks/`

## 📝 Naming Conventions

### Files

- Components: `PascalCase.tsx` (e.g., `LoginScreen.tsx`)
- Hooks: `camelCase.ts` (e.g., `useAuth.ts`)
- Utilities: `camelCase.ts` (e.g., `format.ts`)
- Types: `camelCase.types.ts` (e.g., `auth.types.ts`)
- Styles: `PascalCase.styles.ts` (e.g., `Button.styles.ts`)
- Tests: `PascalCase.test.tsx` (e.g., `Button.test.tsx`)

### Variables

- Constants: `SCREAMING_SNAKE_CASE` (e.g., `API_BASE_URL`)
- Functions: `camelCase` (e.g., `getUserProfile`)
- Components: `PascalCase` (e.g., `LoginButton`)
- Types/Interfaces: `PascalCase` (e.g., `UserProfile`)

### Imports

Use path aliases (configure in `tsconfig.json`):

```typescript
import {Button} from '@/components/atoms';
import {useAuth} from '@/features/auth';
import {spacing} from '@/theme';
import {logger} from '@/services/monitoring';
```

## 🧪 Testing

Place tests next to the code they test:

```
Button/
├── Button.tsx
├── Button.test.tsx        # ✅ Co-located
└── index.ts
```

## 📚 Additional Resources

- [Navigation Types](./navigation/types.ts) - Navigation type definitions
- [Theme Documentation](./theme/README.md) - Design system guide
- [Monitoring Setup](../docs/MONITORING_SETUP.md) - Error handling guide
- [Quick Reference](../docs/QUICK_REFERENCE.md) - Monitoring quick start

## ✅ Best Practices

1. **Never hard-code dimensions** - Use theme tokens or responsive utilities
2. **Keep features isolated** - Features should not import from other features directly
3. **Use shared components** - Reuse components from `components/` folder
4. **Type everything** - No `any` types unless absolutely necessary
5. **Co-locate files** - Keep related files together
6. **Export through index.ts** - Control public API of each module
7. **Test important logic** - Write tests for hooks and utilities
8. **Monitor performance** - Use `useScreenLoadTime` on all screens
9. **Log user actions** - Use `logger.logUserAction()` for important events
10. **Handle errors gracefully** - Use error boundaries and error handlers

---

**Questions?** Check the main documentation or create an issue.
