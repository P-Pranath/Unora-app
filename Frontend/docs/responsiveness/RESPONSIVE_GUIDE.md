# Responsiveness Documentation

## Overview

This document provides a comprehensive guide to how responsiveness is handled in this React Native codebase. The approach uses a **standard screen dimension-based scaling system** combined with **device detection utilities** to create a consistent responsive experience across different mobile screen sizes in portrait orientation.

---

## Table of Contents

1. [Core Philosophy](#core-philosophy)
2. [Standard Screen Dimensions](#standard-screen-dimensions)
3. [Responsive Functions](#responsive-functions)
4. [Font Sizing](#font-sizing)
5. [Device Detection](#device-detection)
6. [Theme Constants](#theme-constants)
7. [Usage Patterns](#usage-patterns)
8. [Best Practices](#best-practices)
9. [Examples](#examples)
10. [Implementation Guide](#implementation-guide)

**Note**: This documentation is specifically for **mobile portrait-only** applications. Tablet and landscape orientation handling are not covered.

---

## Core Philosophy

The responsiveness approach in this codebase is based on:

1. **Standard Reference Screen**: iPhone X (375 x 812 pixels) is used as the standard reference screen
2. **Proportional Scaling**: All dimensions scale proportionally based on the ratio between actual device dimensions and standard dimensions
3. **Function-Based Approach**: All responsive sizing is done through utility functions rather than hardcoded values
4. **Device-Aware**: Special handling for devices with notches (mobile only)
5. **Font Scaling**: Separate responsive font sizing using `react-native-responsive-fontsize`
6. **Mobile Portrait Only**: This approach is optimized for mobile devices in portrait orientation only

---

## Standard Screen Dimensions

### Constants

**Location**: `src/constant/index.ts` and `src/utils/constants.ts`

```typescript
export const STANDARD_SCREEN_SIZE = 812;
export const STANDARD_SCREEN_DIMENSIONS = {height: 812, width: 375};
```

**Standard Reference**:

- **Width**: 375 pixels (iPhone X width)
- **Height**: 812 pixels (iPhone X height)
- **Base Screen Size**: 812 (used for font scaling)

**Note**: All responsive calculations are based on these standard dimensions. If your design is based on a different screen size, update these constants accordingly.

---

## Responsive Functions

### Primary Functions

#### 1. `RfW(value: number)` - Responsive Width

**Location**: `src/utils/helper.tsx` and `src/utils/helpers.ts`

**Purpose**: Scales width values proportionally based on device width.

**Formula**:

```typescript
const RfW = (value: number) =>
  SCREEN_WIDTH * (value / STANDARD_SCREEN_DIMENSIONS.width);
```

**Example**:

```typescript
// On a device with width 390 (iPhone 12):
RfW(24) = 390 * (24 / 375) = 24.96 ≈ 25 pixels

// On a device with width 428 (iPhone 13 Pro Max):
RfW(24) = 428 * (24 / 375) = 27.39 ≈ 27 pixels
```

**Usage**:

```typescript
import {RfW} from '../../utils/helper';

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: RfW(24),
    width: RfW(327),
  },
});
```

#### 2. `RfH(value: number)` - Responsive Height

**Location**: `src/utils/helper.tsx` and `src/utils/helpers.ts`

**Purpose**: Scales height values proportionally based on device height.

**Formula**:

```typescript
const RfH = (value: number) =>
  SCREEN_HEIGHT * (value / STANDARD_SCREEN_DIMENSIONS.height);
```

**Example**:

```typescript
// On a device with height 844 (iPhone 12):
RfH(50) = 844 * (50 / 812) = 52 pixels

// On a device with height 926 (iPhone 13 Pro Max):
RfH(50) = 926 * (50 / 812) = 57 pixels
```

**Usage**:

```typescript
import {RfH} from '../../utils/helper';

const styles = StyleSheet.create({
  container: {
    paddingVertical: RfH(20),
    height: RfH(240),
  },
});
```

### Screen Dimension Functions

**Location**: `src/constant/index.ts`

```typescript
import {Dimensions} from 'react-native';

export const SCREEN_WIDTH = Dimensions.get('window').width;
export const SCREEN_HEIGHT = Dimensions.get('window').height;
```

**Location**: `src/utils/helpers.ts`

```typescript
export const deviceWidth = () => {
  const dim = Dimensions.get('window');
  return dim.width;
};

export const deviceHeight = () => {
  const dim = Dimensions.get('window');
  return dim.height;
};

export const deviceHeightScreen = () => {
  const dim = Dimensions.get('screen');
  return dim.height;
};
```

**Note**:

- `Dimensions.get('window')` - Gets the visible window dimensions (excludes status bar, navigation bar, etc.)
- `Dimensions.get('screen')` - Gets the full screen dimensions (includes all system UI)

---

## Font Sizing

### RFValue Function

**Package**: `react-native-responsive-fontsize`

**Purpose**: Scales font sizes responsively based on screen height.

**Formula**: Automatically calculates optimal font size based on standard screen size.

**Location**: `src/theme/sizes.ts` and `src/theme/commonStyles.ts`

### FontSize Constants

**Location**: `src/theme/sizes.ts`

```typescript
import {RFValue} from 'react-native-responsive-fontsize';
import {STANDARD_SCREEN_SIZE} from '../constant';

export const FontSize = {
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
  30: RFValue(30, STANDARD_SCREEN_SIZE),
  40: RFValue(40, STANDARD_SCREEN_SIZE),
};
```

**Usage**:

```typescript
import {FontSize} from '../../theme/sizes';

const styles = StyleSheet.create({
  text: {
    fontSize: FontSize[14],
  },
});
```

**Direct RFValue Usage**:

```typescript
import {RFValue} from 'react-native-responsive-fontsize';
import {STANDARD_SCREEN_SIZE} from '../constant';

const styles = StyleSheet.create({
  text: {
    fontSize: RFValue(16, STANDARD_SCREEN_SIZE),
  },
});
```

---

## Device Detection

### Notch Detection

**Location**: `src/utils/helpers.ts`

```typescript
import DeviceInfo from 'react-native-device-info';

export const isDisplayWithNotch = () => {
  return DeviceInfo.hasNotch();
};
```

**Usage**:

```typescript
import {isDisplayWithNotch} from '../../utils/helpers';

const styles = StyleSheet.create({
  container: {
    paddingTop: isDisplayWithNotch() ? RfH(30) : RfH(10),
  },
});
```

---

## Theme Constants

### HEIGHT Constants

**Location**: `src/theme/sizes.ts`

Pre-calculated responsive height values for common use cases.

```typescript
export const HEIGHT = {
  H0: RfH(0),
  H1: RfH(1),
  H2: RfH(2),
  // ... many more
  H20: RfH(20),
  H24: RfH(24),
  H50: RfH(50),
  H100: RfH(100),
  H240: RfH(240),
  // ... up to H554
};
```

**Usage**:

```typescript
import {HEIGHT} from '../../theme/sizes';

const styles = StyleSheet.create({
  container: {
    height: HEIGHT.H240,
    marginTop: HEIGHT.H20,
  },
});
```

**Benefits**:

- Consistent sizing across the app
- Easy to maintain
- Clear intent

### WIDTH Constants

**Location**: `src/theme/sizes.ts`

Pre-calculated responsive width values for common use cases.

```typescript
export const WIDTH = {
  W3: RfW(3),
  W5: RfW(5),
  // ... many more
  W16: RfW(16),
  W24: RfW(24),
  W327: RfW(327),
  // ... up to W376
};
```

**Usage**:

```typescript
import {WIDTH} from '../../theme/sizes';

const styles = StyleSheet.create({
  container: {
    width: WIDTH.W327,
    paddingHorizontal: WIDTH.W24,
  },
});
```

**Benefits**:

- Same as HEIGHT constants
- Ensures design consistency
- Reduces calculation overhead

### BorderRadius Constants

**Location**: `src/theme/sizes.ts`

Fixed border radius values (not responsive - these are design constants).

```typescript
export const BorderRadius = {
  BR0: 0,
  BR5: 5,
  BR10: 10,
  BR12: 12,
  BR15: 15,
  BR20: 20,
};
```

**Usage**:

```typescript
import {BorderRadius} from '../../theme/sizes';

const styles = StyleSheet.create({
  card: {
    borderRadius: BorderRadius.BR12,
  },
});
```

---

## Usage Patterns

### Pattern 1: Basic Responsive Styling

```typescript
import {StyleSheet} from 'react-native';
import {RfW, RfH} from '../../utils/helper';

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: RfW(24),
    paddingVertical: RfH(20),
    width: RfW(327),
    height: RfH(240),
  },
});
```

### Pattern 2: Using Theme Constants

```typescript
import {StyleSheet} from 'react-native';
import {HEIGHT, WIDTH, FontSize} from '../../theme/sizes';

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: WIDTH.W24,
    paddingVertical: HEIGHT.H20,
    width: WIDTH.W327,
  },
  text: {
    fontSize: FontSize[14],
  },
});
```

### Pattern 3: Dynamic Styling (Inline)

```typescript
import {RfW, RfH} from '../../utils/helper';

const MyComponent = () => {
  return (
    <View
      style={{
        paddingHorizontal: RfW(24),
        height: RfH(43),
      }}>
      {/* content */}
    </View>
  );
};
```

### Pattern 4: HTML Content Width

```typescript
import {deviceWidth} from '../../utils/helpers';
import {WIDTH} from '../../theme/sizes';

const contentWidth = deviceWidth() - WIDTH.W50;

// Used for HTML rendering to ensure proper scaling
<CustomRenderHtml contentWidth={contentWidth} source={htmlSource} />;
```

### Pattern 5: Full-Width Components

```typescript
import {deviceWidth} from '../../utils/helpers';

const styles = StyleSheet.create({
  fullWidthImage: {
    width: deviceWidth(),
    height: RfH(335),
  },
});
```

---

## Best Practices

### 1. Always Use Responsive Functions

**❌ Don't:**

```typescript
const styles = StyleSheet.create({
  container: {
    paddingHorizontal: 24, // Fixed value
    width: 327, // Fixed value
  },
});
```

**✅ Do:**

```typescript
const styles = StyleSheet.create({
  container: {
    paddingHorizontal: RfW(24), // Responsive
    width: RfW(327), // Responsive
  },
});
```

### 2. Use Theme Constants When Available

**❌ Don't:**

```typescript
paddingHorizontal: RfW(24),
```

**✅ Do (when constant exists):**

```typescript
paddingHorizontal: WIDTH.W24,
```

**Benefits**:

- Consistency across the app
- Easier maintenance
- Clear design system

### 3. Use RfW for Horizontal Spacing

Use `RfW` for:

- Padding (horizontal)
- Margin (horizontal)
- Width
- Left/Right positioning

### 4. Use RfH for Vertical Spacing

Use `RfH` for:

- Padding (vertical)
- Margin (vertical)
- Height
- Top/Bottom positioning

### 5. Font Sizing

**Always use responsive font sizing:**

```typescript
// ✅ Good
fontSize: FontSize[14];
fontSize: RFValue(16, STANDARD_SCREEN_SIZE);

// ❌ Avoid
fontSize: 14; // Fixed size
```

### 6. Border Radius

**Border radius values are typically fixed (not responsive):**

```typescript
borderRadius: BorderRadius.BR12,  // ✅ Good
borderRadius: RfW(12),            // ❌ Usually unnecessary
```

### 7. Percentage Values

**For full-width or percentage-based layouts:**

```typescript
width: '100%',           // ✅ Good for flex layouts
width: deviceWidth(),    // ✅ Good for absolute widths
width: RfW(375),        // ✅ Good for fixed proportional widths
```

### 8. Screen Dimensions

**Use appropriate dimension source:**

```typescript
// For visible content area (recommended)
const width = Dimensions.get('window').width;

// For full screen (includes system UI)
const height = Dimensions.get('screen').height;
```

### 9. Complex Calculations

**When you need calculated responsive values:**

```typescript
const contentWidth = deviceWidth() - WIDTH.W50;
const imageSize = RfW(253);
const cardHeight = RfH(240);
```

---

## Examples

### Example 1: Card Component

```typescript
import {StyleSheet} from 'react-native';
import {HEIGHT, WIDTH, FontSize, BorderRadius} from '../../theme/sizes';
import {RfH} from '../../utils/helper';

const styles = StyleSheet.create({
  card: {
    width: WIDTH.W327,
    height: HEIGHT.H240,
    borderRadius: BorderRadius.BR12,
    paddingHorizontal: WIDTH.W24,
    paddingVertical: HEIGHT.H20,
  },
  title: {
    fontSize: FontSize[18],
    marginBottom: HEIGHT.H12,
  },
  subtitle: {
    fontSize: FontSize[14],
    marginTop: HEIGHT.H8,
  },
});
```

### Example 2: Input Field

```typescript
import {StyleSheet} from 'react-native';
import {RfW, RfH} from '../../utils/helper';
import {FontSize} from '../../theme/sizes';

const styles = StyleSheet.create({
  inputContainer: {
    width: '100%',
    paddingHorizontal: RfW(16),
  },
  input: {
    height: RfH(43),
    fontSize: FontSize[14],
    paddingHorizontal: RfW(16),
  },
});
```

### Example 3: Full-Width Image Carousel

```typescript
import {StyleSheet} from 'react-native';
import {deviceWidth} from '../../utils/helpers';
import {RfH} from '../../utils/helper';

const styles = StyleSheet.create({
  carouselContainer: {
    width: deviceWidth(),
    height: RfH(335),
  },
  image: {
    width: deviceWidth(),
    height: RfH(335),
  },
});
```

### Example 4: Responsive Grid Layout

```typescript
import {StyleSheet} from 'react-native';
import {RfW, RfH} from '../../utils/helper';
import {WIDTH} from '../../theme/sizes';

const CARD_WIDTH = RfW(156);
const CARD_HEIGHT = RfH(156);

const styles = StyleSheet.create({
  gridContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    paddingHorizontal: WIDTH.W24,
  },
  card: {
    width: CARD_WIDTH,
    height: CARD_HEIGHT,
    marginRight: WIDTH.W15,
    marginBottom: RfH(15),
  },
});
```

### Example 5: Modal with Responsive Padding

```typescript
import {StyleSheet} from 'react-native';
import {RfH} from '../../utils/helper';
import {HEIGHT, WIDTH} from '../../theme/sizes';

const styles = StyleSheet.create({
  modalContainer: {
    paddingHorizontal: WIDTH.W32,
    paddingBottom: HEIGHT.H30,
  },
  content: {
    paddingHorizontal: WIDTH.W24,
    paddingTop: HEIGHT.H20,
  },
});
```

---

## Implementation Guide

### Setting Up in a New Project

#### Step 1: Install Dependencies

```bash
npm install react-native-responsive-fontsize react-native-device-info
# or
yarn add react-native-responsive-fontsize react-native-device-info
```

**Note**: `react-native-device-info` is only needed for notch detection. If you don't need notch detection, you can skip this dependency.

#### Step 2: Define Standard Dimensions

Create `src/constant/index.ts`:

```typescript
import {Dimensions} from 'react-native';

// iPhone X dimensions (or your design's base screen)
export const STANDARD_SCREEN_SIZE = 812;
export const STANDARD_SCREEN_DIMENSIONS = {height: 812, width: 375};

export const SCREEN_WIDTH = Dimensions.get('window').width;
export const SCREEN_HEIGHT = Dimensions.get('window').height;
```

#### Step 3: Create Responsive Functions

Create `src/utils/helper.tsx`:

```typescript
import {
  SCREEN_WIDTH,
  SCREEN_HEIGHT,
  STANDARD_SCREEN_DIMENSIONS,
} from '../constant';

export const RfW = (value: number) =>
  SCREEN_WIDTH * (value / STANDARD_SCREEN_DIMENSIONS.width);

export const RfH = (value: number) =>
  SCREEN_HEIGHT * (value / STANDARD_SCREEN_DIMENSIONS.height);
```

Create `src/utils/helpers.ts`:

```typescript
import {Dimensions} from 'react-native';
import DeviceInfo from 'react-native-device-info';
import {STANDARD_SCREEN_DIMENSIONS} from './constants';

export const deviceWidth = () => {
  const dim = Dimensions.get('window');
  return dim.width;
};

export const deviceHeight = () => {
  const dim = Dimensions.get('window');
  return dim.height;
};

export const isDisplayWithNotch = () => {
  return DeviceInfo.hasNotch();
};

export const RfW = value => {
  const dim = Dimensions.get('window');
  return dim.width * (value / STANDARD_SCREEN_DIMENSIONS.width);
};

export const RfH = value => {
  const dim = Dimensions.get('window');
  return dim.height * (value / STANDARD_SCREEN_DIMENSIONS.height);
};
```

**Note**: Orientation functions (`isPortrait`, `isLandscape`) are not included as this approach is for portrait-only mobile apps.

#### Step 4: Create Theme Constants

Create `src/theme/sizes.ts`:

```typescript
import {RFValue} from 'react-native-responsive-fontsize';
import {STANDARD_SCREEN_SIZE} from '../constant';
import {RfH, RfW} from '../utils/helper';

export const HEIGHT = {
  H0: RfH(0),
  H1: RfH(1),
  H2: RfH(2),
  // Add common values
  H8: RfH(8),
  H10: RfH(10),
  H12: RfH(12),
  H16: RfH(16),
  H20: RfH(20),
  H24: RfH(24),
  // ... add more as needed
};

export const WIDTH = {
  W3: RfW(3),
  W5: RfW(5),
  W8: RfW(8),
  W10: RfW(10),
  W12: RfW(12),
  W16: RfW(16),
  W24: RfW(24),
  // ... add more as needed
};

export const FontSize = {
  10: RFValue(10, STANDARD_SCREEN_SIZE),
  11: RFValue(11, STANDARD_SCREEN_SIZE),
  12: RFValue(12, STANDARD_SCREEN_SIZE),
  14: RFValue(14, STANDARD_SCREEN_SIZE),
  16: RFValue(16, STANDARD_SCREEN_SIZE),
  18: RFValue(18, STANDARD_SCREEN_SIZE),
  20: RFValue(20, STANDARD_SCREEN_SIZE),
  // ... add more as needed
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

#### Step 5: Start Using

```typescript
import {StyleSheet} from 'react-native';
import {RfW, RfH} from '../../utils/helper';
import {HEIGHT, WIDTH, FontSize} from '../../theme/sizes';

const styles = StyleSheet.create({
  container: {
    paddingHorizontal: WIDTH.W24,
    paddingVertical: HEIGHT.H20,
  },
});
```

---

## Key Takeaways

1. **Standard Screen**: Use a reference screen (e.g., iPhone X: 375x812) for all design specs
2. **Proportional Scaling**: `RfW` and `RfH` scale everything proportionally across all mobile screen sizes
3. **Mobile Portrait Only**: This approach is optimized for mobile devices in portrait orientation
4. **Font Scaling**: Always use `RFValue` or `FontSize` constants for fonts
5. **Theme Constants**: Use pre-calculated constants from `HEIGHT` and `WIDTH` for consistency
6. **Consistency**: Stick to the patterns - don't mix fixed and responsive values
7. **Calculation Once**: Theme constants are calculated once at import, improving performance
8. **Notch Awareness**: Use `isDisplayWithNotch()` when you need to adjust for devices with notches

---

## Common Pitfalls to Avoid

1. **❌ Mixing fixed and responsive values**: Always use responsive functions
2. **❌ Hardcoding dimensions**: Never use fixed pixel values directly
3. **❌ Using wrong dimension source**: Use `window` for visible area, `screen` for full screen
4. **❌ Not using theme constants**: Leverage existing constants for consistency
5. **❌ Adding tablet-specific code**: This approach is mobile-only - don't add tablet conditionals
6. **❌ Orientation handling**: Don't add landscape/portrait conditionals - focus on portrait only

---

## Summary

This responsiveness approach provides:

- ✅ **Consistent scaling** across all devices
- ✅ **Easy maintenance** through centralized constants
- ✅ **Type safety** (if using TypeScript)
- ✅ **Performance** through pre-calculated constants
- ✅ **Flexibility** for device-specific adjustments
- ✅ **Design system alignment** with standard screen dimensions

By following these patterns, your app will look consistent and properly scaled across all device sizes while maintaining code that's easy to understand and maintain.
