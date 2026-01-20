# Asset Management Guide

This document explains how to organize, add, and use assets (images, icons, fonts, animations) in the app.

## Directory Structure

```
src/assets/
├── fonts/           # Custom font files (.ttf, .otf)
├── icons/           # SVG icons or small PNG icons
│   └── index.ts     # Icon exports
├── images/          # Raster images (PNG, JPG, WebP)
│   └── index.ts     # Image exports
└── lottie/          # Lottie animation JSON files
    └── index.ts     # Animation exports
```

## Asset Categories

### 1. Images (`src/assets/images/`)

Static raster images like logos, backgrounds, and illustrations.
**Naming Convention:**

- Use `snake_case` for file names
- Include size variant suffix: `@2x`, `@3x` for resolution variants
- Prefix with category when helpful: `bg_`, `logo_`, `illustration_`
  **Examples:**

```
logo.png
logo@2x.png
logo@3x.png
bg_onboarding.png
illustration_empty_state.png
```

**How to Add:**

1. Place image files in `src/assets/images/`
2. Add export to `src/assets/images/index.ts`:

```typescript
export const Images = {
  logo: require('./logo.png'),
  bgOnboarding: require('./bg_onboarding.png'),
  // Add new images here
} as const;
```

3. Use in components:

```typescript
import {Images} from '@/assets/images';
<Image source={Images.logo} />;
```

### 2. Icons (`src/assets/icons/`)

Small vector or raster icons. Prefer SVG where possible for scalability.
**Naming Convention:**

- Use `snake_case` for file names
- Prefix with `ic_` for consistency
- Use descriptive names: `ic_arrow_right.svg`, `ic_close.svg`
  **How to Add SVG Icons:**

1. Place SVG file in `src/assets/icons/`
2. Use `react-native-svg` or create a component:

```typescript
// src/assets/icons/ArrowRight.tsx
import * as React from 'react';
import Svg, {Path} from 'react-native-svg';

interface Props {
  size?: number;
  color?: string;
}

export const ArrowRightIcon: React.FC<Props> = ({
  size = 24,
  color = '#000',
}) => (
  <Svg width={size} height={size} viewBox="0 0 24 24" fill="none">
    <Path d="M..." fill={color} />
  </Svg>
);
```

3. Export from index:

```typescript
// src/assets/icons/index.ts
export {ArrowRightIcon} from './ArrowRight';
export {CloseIcon} from './Close';
```

### 3. Fonts (`src/assets/fonts/`)

Custom font files for typography.
**Naming Convention:**

- Use the exact font family name: `Roboto-Regular.ttf`
- Include all weight variants you need
  **How to Add:**

1. Place font files in `src/assets/fonts/`
2. Link fonts (done via `react-native.config.js`):

```javascript
module.exports = {
  assets: ['./src/assets/fonts'],
};
```

3. Run: `npx react-native-asset`
4. Reference in styles using the font family name

### 4. Lottie Animations (`src/assets/lottie/`)

JSON animation files from After Effects via Bodymovin.
**Naming Convention:**

- Use `snake_case`: `loading_spinner.json`, `success_check.json`
  **How to Add:**

1. Place JSON file in `src/assets/lottie/`
2. Export from index:

```typescript
// src/assets/lottie/index.ts
export const Animations = {
  loadingSpinner: require('./loading_spinner.json'),
  successCheck: require('./success_check.json'),
} as const;
```

3. Use with `lottie-react-native`:

```typescript
import LottieView from 'lottie-react-native';
import {Animations} from '@/assets/lottie';

<LottieView source={Animations.loadingSpinner} autoPlay loop />;
```

## Best Practices

### Image Optimization

- **Compress images** before adding to the project. Use tools like:
  - [TinyPNG](https://tinypng.com/) for PNG/JPG
  - [SVGOMG](https://jakearchibald.github.io/svgomg/) for SVGs
- **Use WebP** where possible (smaller file size, good quality)
- **Provide @2x and @3x variants** for crisp display on all devices
- **Avoid images larger than 1MB** unless absolutely necessary

### Resolution Variants

React Native automatically picks the right resolution based on device pixel density:

```
logo.png      → 1x (baseline)
logo@2x.png   → 2x (Retina, most modern phones)
logo@3x.png   → 3x (high-density displays like iPhone Pro)
```

Only the base name is used in require():

```typescript
require('./logo.png'); // RN picks @2x or @3x automatically
```

### Type Safety

Always export assets through typed index files:

```typescript
// ✅ Good - typed and centralized
export const Images = {
  logo: require('./logo.png'),
} as const;

export type ImageKey = keyof typeof Images;

// ❌ Bad - scattered requires throughout codebase
<Image source={require('../../../assets/images/logo.png')} />;
```

### Organizing Large Asset Collections

For apps with many assets, organize by feature or category:

```
src/assets/images/
├── index.ts           # Main export (re-exports all)
├── branding/
│   ├── index.ts
│   ├── logo.png
│   └── logo_white.png
├── onboarding/
│   ├── index.ts
│   ├── slide_1.png
│   └── slide_2.png
└── icons/
    ├── index.ts
    ├── home.png
    └── profile.png
```

Main index re-exports:

```typescript
// src/assets/images/index.ts
export {BrandingImages} from './branding';
export {OnboardingImages} from './onboarding';
export {IconImages} from './icons';

// Or flatten:
import {BrandingImages} from './branding';
import {OnboardingImages} from './onboarding';

export const Images = {
  ...BrandingImages,
  ...OnboardingImages,
} as const;
```

## Platform-Specific Assets

For assets that differ by platform:

```typescript
import {Platform} from 'react-native';

export const Images = {
  statusBarBg: Platform.select({
    ios: require('./status_bar_ios.png'),
    android: require('./status_bar_android.png'),
  }),
} as const;
```

Or use file naming:

```
icon.ios.png
icon.android.png
```

## Checklist for Adding Assets

- [ ] File is optimized/compressed
- [ ] File uses correct naming convention
- [ ] Resolution variants provided (@2x, @3x) for images
- [ ] Added to appropriate index.ts export
- [ ] No duplicate assets (check existing exports first)
- [ ] Used through import, not inline require

## Removing Assets

1. Remove the file from the filesystem
2. Remove the export from the index file
3. Search codebase for any remaining references
4. Run TypeScript check to catch broken imports: `yarn typecheck`
