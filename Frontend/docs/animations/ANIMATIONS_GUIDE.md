# Subtle & Professional Animations Guide

This guide provides best practices and patterns for adding subtle, professional animations to features using `react-native-reanimated`. Use this as a reference when implementing animations to ensure consistency, performance, and a polished user experience.

## Table of Contents

- [Principles](#principles)
- [Setup & Dependencies](#setup--dependencies)
- [Animation Patterns](#animation-patterns)
- [Performance Best Practices](#performance-best-practices)
- [Common Use Cases](#common-use-cases)
- [Code Examples](#code-examples)

## Principles

### 1. Subtlety First

- Animations should enhance, not distract
- Keep scale values small (0.9-1.1 range)
- Use short durations (200-300ms for most animations)
- Avoid excessive motion or bouncy effects

### 2. Performance Priority

- Always use `react-native-reanimated` (not Animated API)
- Animations run on UI thread (60fps guaranteed)
- Avoid animating layout properties when possible
- Use `useAnimatedStyle` for style animations

### 3. Consistency

- Use similar timing and easing across the app
- Maintain consistent scale/opacity values
- Follow established patterns for similar interactions

## Setup & Dependencies

### Required Package

```json
"react-native-reanimated": "^3.18.0"
```

### Babel Configuration

Ensure `react-native-reanimated/plugin` is in your `babel.config.js`:

```javascript
plugins: [
  // ... other plugins
  'react-native-reanimated/plugin', // Must be last
],
```

### Import Pattern

```typescript
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withTiming,
  withSpring,
  FadeIn,
  FadeOut,
} from 'react-native-reanimated';
```

## Animation Patterns

### 1. Fade-In on Mount

**Use case**: Components appearing on screen

```typescript
const opacity = useSharedValue(0);
const translateY = useSharedValue(-10);

useEffect(() => {
  opacity.value = withTiming(1, {duration: 300});
  translateY.value = withSpring(0, {
    damping: 20,
    stiffness: 90,
  });
}, []);

const animatedStyle = useAnimatedStyle(() => ({
  opacity: opacity.value,
  transform: [{translateY: translateY.value}],
}));

return <Animated.View style={animatedStyle}>...</Animated.View>;
```

### 2. Press/Interaction Feedback

**Use case**: Buttons, icons, interactive elements

```typescript
const scale = useSharedValue(1);

const animatedStyle = useAnimatedStyle(() => ({
  transform: [{scale: scale.value}],
}));

const handlePressIn = () => {
  scale.value = withSpring(0.92, {
    damping: 15,
    stiffness: 300,
  });
};

const handlePressOut = () => {
  scale.value = withSpring(1, {
    damping: 15,
    stiffness: 300,
  });
};

return (
  <TouchableOpacity
    onPressIn={handlePressIn}
    onPressOut={handlePressOut}
    activeOpacity={1}>
    <Animated.View style={animatedStyle}>{/* Content */}</Animated.View>
  </TouchableOpacity>
);
```

### 3. State-Based Transitions

**Use case**: Active/inactive states, focus changes

```typescript
const scale = useSharedValue(focused ? 1.1 : 1);
const opacity = useSharedValue(focused ? 1 : 0.7);

useEffect(() => {
  scale.value = withSpring(focused ? 1.1 : 1, {
    damping: 15,
    stiffness: 300,
  });
  opacity.value = withTiming(focused ? 1 : 0.7, {duration: 200});
}, [focused]);

const animatedStyle = useAnimatedStyle(() => ({
  transform: [{scale: scale.value}],
  opacity: opacity.value,
}));
```

### 4. Enter/Exit Animations

**Use case**: Lists, modals, conditional rendering

```typescript
// Using entering/exiting props
<Animated.View entering={FadeIn.duration(200)} exiting={FadeOut.duration(150)}>
  {/* Content */}
</Animated.View>
```

## Performance Best Practices

### ✅ DO

1. **Use `useAnimatedStyle` for style animations**

   ```typescript
   const animatedStyle = useAnimatedStyle(() => ({
     opacity: opacity.value,
     transform: [{scale: scale.value}],
   }));
   ```

2. **Keep durations short (200-300ms)**

   ```typescript
   withTiming(1, {duration: 200}); // ✅ Good
   withTiming(1, {duration: 1000}); // ❌ Too slow
   ```

3. **Use spring animations for natural feel**

   ```typescript
   withSpring(1, {
     damping: 15, // Higher = less bouncy
     stiffness: 300, // Higher = faster
   });
   ```

4. **Animate transform and opacity (most performant)**

   ```typescript
   // ✅ Good - runs on UI thread
   transform: [{scale: scale.value}];
   opacity: opacity.value;

   // ⚠️ Avoid - may cause layout recalculations
   width: width.value;
   height: height.value;
   ```

5. **Initialize shared values with correct initial state**
   ```typescript
   const scale = useSharedValue(focused ? 1.1 : 1); // ✅
   const scale = useSharedValue(1); // ❌ If focused should start at 1.1
   ```

### ❌ DON'T

1. **Don't use React Native's Animated API** (use reanimated instead)
2. **Don't animate layout properties** (width, height, padding) unless necessary
3. **Don't use long durations** (>500ms for most cases)
4. **Don't create excessive animations** (every element doesn't need animation)
5. **Don't forget to handle cleanup** (reanimated handles this automatically)

## Common Use Cases

### Header/Component Fade-In

```typescript
const HeaderComponent = () => {
  const opacity = useSharedValue(0);
  const translateY = useSharedValue(-10);

  useEffect(() => {
    opacity.value = withTiming(1, {duration: 300});
    translateY.value = withSpring(0, {
      damping: 20,
      stiffness: 90,
    });
  }, []);

  const animatedStyle = useAnimatedStyle(() => ({
    opacity: opacity.value,
    transform: [{translateY: translateY.value}],
  }));

  return (
    <Animated.View style={[styles.container, animatedStyle]}>
      {/* Header content */}
    </Animated.View>
  );
};
```

### Animated Icon Button

```typescript
const AnimatedIconButton: React.FC<{
  onPress?: () => void;
  children: React.ReactNode;
}> = ({onPress, children}) => {
  const scale = useSharedValue(1);

  const animatedStyle = useAnimatedStyle(() => ({
    transform: [{scale: scale.value}],
  }));

  const handlePressIn = () => {
    scale.value = withSpring(0.92, {
      damping: 15,
      stiffness: 300,
    });
  };

  const handlePressOut = () => {
    scale.value = withSpring(1, {
      damping: 15,
      stiffness: 300,
    });
  };

  return (
    <TouchableOpacity
      onPress={onPress}
      onPressIn={handlePressIn}
      onPressOut={handlePressOut}
      activeOpacity={1}>
      <Animated.View style={[styles.button, animatedStyle]}>
        {children}
      </Animated.View>
    </TouchableOpacity>
  );
};
```

### Tab Bar Icon/Label Animation

```typescript
const TabIcon = ({focused, IconComponent}) => {
  const scale = useSharedValue(focused ? 1.1 : 1);
  const opacity = useSharedValue(focused ? 1 : 0.7);

  useEffect(() => {
    scale.value = withSpring(focused ? 1.1 : 1, {
      damping: 15,
      stiffness: 300,
    });
    opacity.value = withTiming(focused ? 1 : 0.7, {duration: 200});
  }, [focused]);

  const animatedStyle = useAnimatedStyle(() => ({
    transform: [{scale: scale.value}],
    opacity: opacity.value,
  }));

  return (
    <Animated.View style={animatedStyle}>
      <IconComponent size={24} color={focused ? '#46238D' : '#101828'} />
    </Animated.View>
  );
};
```

### Badge/Notification Count Animation

```typescript
<Animated.View entering={FadeIn.duration(200)} style={styles.badgeContainer}>
  <Badge count={unreadCount} />
</Animated.View>
```

## Code Examples

### Complete Example: Animated Header with Interactive Buttons

```typescript
import React, {useEffect} from 'react';
import {TouchableOpacity, StyleSheet} from 'react-native';
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withTiming,
  withSpring,
  FadeIn,
} from 'react-native-reanimated';

export const AnimatedHeader = () => {
  // Header fade-in animation
  const headerOpacity = useSharedValue(0);
  const headerTranslateY = useSharedValue(-10);

  useEffect(() => {
    headerOpacity.value = withTiming(1, {duration: 300});
    headerTranslateY.value = withSpring(0, {
      damping: 20,
      stiffness: 90,
    });
  }, []);

  const animatedHeaderStyle = useAnimatedStyle(() => ({
    opacity: headerOpacity.value,
    transform: [{translateY: headerTranslateY.value}],
  }));

  // Animated button component
  const AnimatedButton = ({onPress, children}) => {
    const scale = useSharedValue(1);

    const animatedStyle = useAnimatedStyle(() => ({
      transform: [{scale: scale.value}],
    }));

    const handlePressIn = () => {
      scale.value = withSpring(0.92, {
        damping: 15,
        stiffness: 300,
      });
    };

    const handlePressOut = () => {
      scale.value = withSpring(1, {
        damping: 15,
        stiffness: 300,
      });
    };

    return (
      <TouchableOpacity
        onPress={onPress}
        onPressIn={handlePressIn}
        onPressOut={handlePressOut}
        activeOpacity={1}>
        <Animated.View style={[styles.button, animatedStyle]}>
          {children}
        </Animated.View>
      </TouchableOpacity>
    );
  };

  return (
    <Animated.View style={[styles.container, animatedHeaderStyle]}>
      {/* Header content */}
      <AnimatedButton onPress={() => {}}>{/* Button content */}</AnimatedButton>
    </Animated.View>
  );
};
```

## Animation Timing Reference

| Animation Type | Duration  | Easing             | Use Case                    |
| -------------- | --------- | ------------------ | --------------------------- |
| Fade-in        | 200-300ms | `withTiming`       | Component appearance        |
| Press feedback | 150-200ms | `withSpring`       | Button interactions         |
| State change   | 200ms     | `withTiming`       | Active/inactive states      |
| Scale on focus | 200ms     | `withSpring`       | Tab icons, focused elements |
| Enter/Exit     | 150-200ms | `FadeIn`/`FadeOut` | Conditional rendering       |

## Spring Configuration Reference

### Subtle (Recommended)

```typescript
{
  damping: 15,
  stiffness: 300,
}
```

### More Bouncy (Use Sparingly)

```typescript
{
  damping: 10,
  stiffness: 200,
}
```

### Less Bouncy (Very Subtle)

```typescript
{
  damping: 20,
  stiffness: 400,
}
```

## Checklist for Adding Animations

When adding animations to a feature, ensure:

- [ ] Using `react-native-reanimated` (not Animated API)
- [ ] Animations are subtle (scale 0.9-1.1, duration 200-300ms)
- [ ] Using `useAnimatedStyle` for style animations
- [ ] Animating transform/opacity (not layout properties)
- [ ] Initial values match component's initial state
- [ ] Spring animations use appropriate damping/stiffness
- [ ] No performance issues (60fps maintained)
- [ ] Animations enhance UX without being distracting
- [ ] Consistent with existing animation patterns
- [ ] Reduced motion preferences are respected

## Accessibility: Reduced Motion Support

### Overview

Users may have enabled "Reduce Motion" or "Reduce Animations" in their device accessibility settings. This preference indicates that they prefer minimal or no animations due to motion sensitivity, vestibular disorders, or personal preference. **All animations must respect this preference.**

### Detecting Reduced Motion Preferences

React Native Reanimated provides a `useReducedMotion` hook to detect the user's motion preference synchronously:

```typescript
import {useReducedMotion} from 'react-native-reanimated';

const reducedMotion = useReducedMotion();
```

**Important Notes:**

- The hook returns the reduced motion preference as it was at app startup (static value)
- It does not update if the user changes the setting while the app is running
- It provides synchronous access, unlike `AccessibilityInfo.isReduceMotionEnabled()` which is async
- For most use cases, checking at startup is sufficient

### Implementation Guidelines

1. **Disable or Drastically Shorten Animations**

   - When reduced motion is enabled, set animation durations to `0` or a minimal value (≤50ms)
   - Skip transform animations (scale, translate) entirely
   - Keep only essential opacity changes if necessary, with instant transitions

2. **Conditional Animation Logic**

   ```typescript
   import {
     useReducedMotion,
     useSharedValue,
     useAnimatedStyle,
   } from 'react-native-reanimated';

   const reducedMotion = useReducedMotion();
   // useReducedMotion is static (doesn't change during runtime), so initialize once
   const reducedMotionValue = useSharedValue(reducedMotion);
   const duration = reducedMotion ? 0 : 300;
   const scale = useSharedValue(1);

   // Skip transform animations when reduced motion is enabled
   const animatedStyle = useAnimatedStyle(() => ({
     opacity: opacity.value,
     // Only apply transform if motion is not reduced (use shared value in worklet)
     ...(reducedMotionValue.value ? {} : {transform: [{scale: scale.value}]}),
   }));
   ```

3. **Provide Non-Animated Alternatives**

   - Use immediate state changes instead of animated transitions
   - Ensure visual feedback is still clear without motion (e.g., color changes, border styles)
   - Maintain all functionality - animations are enhancements, not requirements

4. **Motion-Safe Fallbacks**

   ```typescript
   import {
     useReducedMotion,
     useSharedValue,
     useAnimatedStyle,
     withSpring,
   } from 'react-native-reanimated';
   import {useCallback} from 'react';

   const usePressAnimation = (options?: UsePressAnimationOptions) => {
     const reducedMotion = useReducedMotion();
     // useReducedMotion is static (doesn't change during runtime), so initialize once
     const reducedMotionValue = useSharedValue(reducedMotion);
     const scale = useSharedValue(1);

     const animatedStyle = useAnimatedStyle(() => {
       // Skip scale animation if reduced motion is enabled (use shared value in worklet)
       if (reducedMotionValue.value) {
         return {}; // No transform applied
       }
       return {
         transform: [{scale: scale.value}],
       };
     });

     const handlePressIn = useCallback(() => {
       if (!reducedMotion) {
         scale.value = withSpring(pressedScale, springConfig);
       }
     }, [pressedScale, springConfig, scale, reducedMotion]);

     // ... rest of implementation
   };
   ```

### Example: Updated Hook with Reduced Motion Support

```typescript
// Example: useFadeInAnimation with reduced motion support
import {
  useReducedMotion,
  useSharedValue,
  useAnimatedStyle,
  withTiming,
  withSpring,
} from 'react-native-reanimated';
import {useEffect, useRef} from 'react';

export const useFadeInAnimation = (options?: UseFadeInAnimationOptions) => {
  const reducedMotion = useReducedMotion();
  // useReducedMotion is static (doesn't change during runtime), so initialize once
  const reducedMotionValue = useSharedValue(reducedMotion);
  const duration = reducedMotion ? 0 : options?.duration ?? 300;
  const initialTranslateY = reducedMotion
    ? 0
    : options?.initialTranslateY ?? -10;

  const opacity = useSharedValue(reducedMotion ? 1 : 0);
  const translateY = useSharedValue(initialTranslateY);

  useEffect(() => {
    if (reducedMotion) {
      // Instant transition, no animation
      opacity.value = 1;
      translateY.value = 0;
    } else {
      opacity.value = withTiming(1, {duration});
      translateY.value = withSpring(0, springConfig);
    }
  }, [duration, springConfig, opacity, translateY, reducedMotion]);

  const animatedStyle = useAnimatedStyle(() => ({
    opacity: opacity.value,
    // Skip translateY transform if reduced motion is enabled (use shared value in worklet)
    ...(reducedMotionValue.value
      ? {}
      : {transform: [{translateY: translateY.value}]}),
  }));

  return animatedStyle;
};
```

### Testing Checklist

- [ ] Enable "Reduce Motion" in device accessibility settings (iOS: Settings → Accessibility → Motion → Reduce Motion; Android: Settings → Accessibility → Remove animations)
- [ ] Verify all animations are disabled or instant (≤50ms)
- [ ] Confirm transform animations (scale, translate) are skipped
- [ ] Ensure visual feedback remains clear without animations
- [ ] Test that all functionality works correctly without motion
- [ ] Verify entering/exiting animations respect the preference

### References to Update

When implementing reduced motion support, update these example files:

- `src/hooks/usePressAnimation.ts` - Conditionally skip scale transforms
- `src/hooks/useFocusAnimation.ts` - Skip scale animations, keep instant opacity changes
- `src/hooks/useFadeInAnimation.ts` - Set duration to 0, skip translateY transform
- `src/components/UI/AnimatedIconButton/index.tsx` - Respect reduced motion in press animations

## References

- [React Native Reanimated Documentation](https://docs.swmansion.com/react-native-reanimated/)
- [Reanimated Performance Guide](https://docs.swmansion.com/react-native-reanimated/docs/fundamentals/glossary/#worklets)
- Example implementations:
  - `src/components/layout/ProfileHeader/index.tsx`
  - `src/navigation/tabs/MainTabsNavigator.tsx`

---

**Remember**: Animations should feel natural and enhance the user experience, not draw attention to themselves. When in doubt, keep it subtle!
