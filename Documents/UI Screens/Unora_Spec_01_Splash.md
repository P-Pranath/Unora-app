# Unora UI Specification — Splash Screen

**Document ID:** Spec-01  
**Screen Name:** Splash Screen  
**Version:** 1.0  
**Date:** January 2026  
**Classification:** Internal / Engineering / Design  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name

**Splash Screen** (also: Launch Screen, App Start Screen)

### 1.2 User Flow Reference

**Phase 1: Verified Onboarding** — The Splash Screen is the very first visual the user encounters when launching the Unora app. It precedes all user flows, serving as the app initialization screen before transitioning to either:
- **Welcome Screen** (new/logged-out users)
- **Discovery/Home Screen** (authenticated, verified users)

**Reference:** [Unora_UserFlow_Logic.md — Section 1 (High-Level Macro Flow)](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose

The Splash Screen communicates brand identity and provides visual continuity during app initialization, establishing the calm and intentional tone that defines the Unora experience.

### 1.4 Primary Action

**None** — This is a passive screen. The user does not interact with the Splash Screen. The system automatically transitions out once initialization completes.

---

## 2. Low-Fidelity Wireframe (ASCII)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                     [Status Bar]                            │   ← System status bar
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                   ┌───────────────┐                         │
│                   │               │                         │
│                   │   [ UNORA ]   │                         │   ← App Logo (centered)
│                   │     LOGO      │                         │     64px, subtle pulse animation
│                   │               │                         │
│                   └───────────────┘                         │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │   ← Safe area (bottom)
└─────────────────────────────────────────────────────────────┘
```

### Layout Constraint Check

| Validation | Result |
|------------|--------|
| **Layout Type** | Static (no scrolling, no interaction) |
| **DSD Alignment** | Matches onboarding screen philosophy (DSD Section 9) |
| **Single Focus** | ✓ One centered element (logo) |
| **No CTAs** | ✓ No interactive elements |
| **Minimal UI** | ✓ Maximum visual calm |

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  SPLASH SCREEN CONTAINER                                                      │
│                                                                               │
│  ├── Position: fixed                                                          │
│  ├── Width: 100vw                                                             │
│  ├── Height: 100vh                                                            │
│  ├── Padding: var(--padding-screen) → 20px (all sides)                        │
│  ├── Display: flex                                                            │
│  ├── Flex-direction: column                                                   │
│  ├── Justify-content: center                                                  │
│  ├── Align-items: center                                                      │
│  └── Background: var(--surface-background) → #FAFAF8                          │
│                                                                               │
│  Premium Dark Mode (Default):                                                 │
│  ├── Background: var(--pdm-background) → #0D0D0F                              │
│  ├── Subtle radial gradient from center:                                      │
│  │   radial-gradient(ellipse at center,                                       │
│  │     rgba(196, 167, 125, 0.03) 0%,    /* Logo gold ambient */               │
│  │     transparent 50%)                                                       │
│  └── Creates warm, inviting depth around logo area                            │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` with subtle warm radial gradient from center |
| **Logo** | Soft ambient glow: `0 0 40px rgba(196, 167, 125, 0.2)` |
| **Atmosphere** | Warm, premium, anticipatory — not cold or sterile |

**Logo Glow Effect:**
```css
/* Logo ambient glow in premium dark mode */
.splash-logo {
  filter: drop-shadow(0 0 24px rgba(196, 167, 125, 0.25));
}
```


### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen padding | `var(--padding-screen)` | 20px |
| Safe area (iOS) | System | env(safe-area-inset-top), env(safe-area-inset-bottom) |
| Logo container margin | — | Auto-centered (flexbox) |

### 3.3 Z-Index Layers

| Layer Name | Z-Index | Contents |
|------------|---------|----------|
| **Background Layer** | 0 | Solid background color |
| **Content Layer** | 1 | Logo element |
| **System Layer** | 100+ | iOS/Android status bar (system-controlled) |

> [!NOTE]
> The Splash Screen has no overlay or modal layers. It exists as a single flat layer.

---

## 4. Component Inventory

### 4.1 Logo Component

**Component Name:** Brand Logo (Centered)

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  LOGO COMPONENT                                                               │
│                                                                               │
│  ┌─────────────────────────────────────┐                                      │
│  │                                     │                                      │
│  │            [ UNORA ]                │   ← Logo Asset                       │
│  │                                     │                                      │
│  └─────────────────────────────────────┘                                      │
│                                                                               │
│  Specifications:                                                              │
│  ├── Asset Type: SVG (preferred) or PNG @3x                                   │
│  ├── Size: 64px height (width proportional)                                   │
│  ├── Max Width: 180px                                                         │
│  ├── Position: Centered (horizontal + vertical)                               │
│  ├── Color: Brand logotype in var(--unora-primary-accent) → #C4A77D           │
│  │           OR full-color brand asset                                        │
│  └── Alt Text: "Unora" (for accessibility)                                    │
│                                                                               │
│  Dark Mode:                                                                   │
│  └── Logo variant: Light version for dark backgrounds                         │
│      Color: #F5F5F3 or dedicated dark mode asset                              │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 4.2 Typography Specifications

The Splash Screen contains **no text elements** beyond the logo. If implementing a text-based logotype:

| Element | Font Family | Weight | Size | Color |
|---------|-------------|--------|------|-------|
| **Logotype (text)** | `var(--font-display)` → Outfit | 700 | 28px | `var(--unora-primary-accent)` → #C4A77D |

### 4.3 Color Tokens

| Token | Light Mode | Dark Mode | Usage |
|-------|------------|-----------|-------|
| `--surface-background` | #FAFAF8 | #121212 | Screen background |
| `--unora-primary-accent` | #C4A77D | #C4A77D | Logo tint (if SVG) |
| `--dm-ink-primary` | — | #F5F5F3 | Logo tint (dark mode variant) |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Description |
|---------|---------|-------------|
| **None** | — | The Splash Screen has no interactive elements |

### 5.2 Behaviors

The Splash Screen is entirely system-driven. No user taps or gestures are processed.

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  SYSTEM BEHAVIOR FLOW                                                         │
│                                                                               │
│  APP LAUNCH EVENT                                                             │
│        │                                                                      │
│        ▼                                                                      │
│  ┌─────────────────────────────────────────────────────────────────────────┐  │
│  │  SYSTEM: Display Splash Screen immediately                              │  │
│  │          (no delay, no pre-render)                                      │  │
│  └─────────────────────────────────────────────────────────────────────────┘  │
│        │                                                                      │
│        ▼                                                                      │
│  SYSTEM: Initialize app services in background                                │
│        ├── Load user session (if exists)                                      │
│        ├── Check authentication state                                         │
│        ├── Prefetch initial data                                              │
│        └── Initialize analytics                                               │
│        │                                                                      │
│        ▼                                                                      │
│  ┌─────────────────────────────────────────────────────────────────────────┐  │
│  │  MINIMUM DISPLAY TIME: 1200ms (1.2 seconds)                             │  │
│  │                                                                          │  │
│  │  Purpose: Ensures brand impression regardless of fast init               │  │
│  │  ⚠️ Do NOT display indefinitely — max 5 seconds before timeout           │  │
│  └─────────────────────────────────────────────────────────────────────────┘  │
│        │                                                                      │
│        ▼                                                                      │
│  ┌─────────────────────────────────────────────────────────────────────────┐  │
│  │  [DECISION] Is user authenticated AND verified?                         │  │
│  └─────────────────────────────────────────────────────────────────────────┘  │
│        │                           │                                          │
│        ▼                           ▼                                          │
│  ┌─────────────────┐      ┌─────────────────────────────────────────────────┐ │
│  │     ✓ YES       │      │                    ✗ NO                         │ │
│  │                 │      │                                                 │ │
│  │ Transition to   │      │ Transition to Welcome Screen                    │ │
│  │ Discovery/Home  │      │ (Onboarding flow)                               │ │
│  │ Screen          │      │                                                 │ │
│  └─────────────────┘      └─────────────────────────────────────────────────┘ │
│                                                                               │
└───────────────────────────────────────────────────────────────────────────────┘
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.A.2 (Identity Verification Hard Gate)](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 5.3 Transitions

#### Splash → Next Screen Transition

| Property | Value | Reference |
|----------|-------|-----------|
| **Animation Type** | Cross-fade | DSD Section 5.4 |
| **Duration** | 200ms | DSD Section 5.4 (Server switch duration) |
| **Easing** | ease-out | DSD Page Transitions |
| **Direction** | Fade out splash, fade in next screen |

```
Transition Sequence:
[0ms]      Initialization complete, ready to transition
[0ms]      Splash screen begins fade-out (opacity 1 → 0)
[0ms]      Next screen begins fade-in (opacity 0 → 1)
[200ms]    Transition complete, splash screen removed from DOM
```

#### Logo Pulse Animation (During Display)

| Property | Value |
|----------|-------|
| **Animation Name** | `splash-logo-pulse` |
| **Duration** | 1.5s |
| **Iteration** | infinite (while splash is visible) |
| **Easing** | ease-in-out |

```css
/* Premium Dark Mode: Logo pulse with glow intensity variation */
@keyframes splash-logo-pulse {
  0%   { 
    opacity: 0.85; 
    transform: scale(1.0);
    filter: drop-shadow(0 0 20px rgba(196, 167, 125, 0.15));
  }
  50%  { 
    opacity: 1.0; 
    transform: scale(1.02);
    filter: drop-shadow(0 0 32px rgba(196, 167, 125, 0.35));
  }
  100% { 
    opacity: 0.85; 
    transform: scale(1.0);
    filter: drop-shadow(0 0 20px rgba(196, 167, 125, 0.15));
  }
}

/* Animation feels like the logo is "breathing" with warmth */
```


> [!TIP]
> Respect `prefers-reduced-motion`: When enabled, display static logo at full opacity with no animation.

---

## 6. State Definitions

### 6.1 State Matrix

| State | Visual Appearance | Conditions | Duration |
|-------|-------------------|------------|----------|
| **Default** | Logo centered, pulsing | App launching, services initializing | 1.2s – 5s |
| **Loading (Skeleton)** | Not applicable | — | — |
| **Empty** | Not applicable | — | — |
| **Error** | Timeout message | Initialization fails after 5s | Until retry/dismiss |
| **Locked/Disabled** | Not applicable | — | — |
| **Cooldown** | Not applicable | — | — |

### 6.2 Default State (Primary)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                   ┌───────────────┐                         │
│                   │               │                         │
│                   │   [ UNORA ]   │                         │   ← Logo (pulsing)
│                   │               │                         │
│                   └───────────────┘                         │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Background: var(--surface-background) → #FAFAF8
Logo: Centered, 64px, subtle pulse animation
Tone: Calm, unhurried, brand-forward
```

### 6.3 Error State (Timeout)

If app initialization fails after **5 seconds**, display a non-alarming error state:

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                        📡                                   │   ← Icon: 48px
│                                                             │       var(--unora-ink-tertiary)
│                                                             │
│              Couldn't start Unora                           │   ← H3 style
│                                                             │
│   Check your connection and try again.                      │   ← Body style, muted
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                    Try Again                        │   │   ← Primary Button
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Error State Styling:
├── Background: var(--surface-background)
├── Icon: WifiSlash or SignalSlash, 48px, var(--unora-ink-tertiary)
├── Headline: H3, var(--unora-ink-primary)
├── Body: Body size, var(--unora-ink-secondary)
├── Button: Primary style (see DSD Section 3.1)
└── Tone: Calm, simple, no technical jargon

⚠️ Do NOT use red colors or alarming language.
   This is a recoverable state, not a catastrophe.
```

**Reference:** [Unora_DesignSystem.md — Section 10.2 (Network Error)](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

---

## 7. Content & Copy Guidelines

### 7.1 Text Strings

| Element | Copy | Fallback |
|---------|------|----------|
| **Logo alt text** | "Unora" | — |
| **Screen reader announcement** | "Unora is loading" | — |
| **Error headline** | "Couldn't start Unora" | "Unable to load" |
| **Error body** | "Check your connection and try again." | "Please try again." |
| **Error button** | "Try Again" | "Retry" |

### 7.2 Tone Guidelines

The Splash Screen embodies the core Unora philosophy:

| Principle | Application |
|-----------|-------------|
| **Presence over Performance** | No loading percentages, no spinning wheels. Just the logo, breathing. |
| **Clarity over Clutter** | Zero text, zero distractions. Pure brand identity. |
| **Anticipation over Gratification** | The pause creates a moment of arrival, not impatience. |
| **Firmness without Friction** | Even in error states, language is calm and empowering, not punitive. |

**Reference:** [Unora_PRD.md — Section 7 (Core Philosophy)](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)

> [!IMPORTANT]
> **Anti-Tinder Positioning:** Unlike apps optimized for speed and urgency, Unora's splash screen deliberately avoids countdown timers, progress bars, or loading indicators. The brief pause signals intentionality.

### 7.3 Avoid

| ❌ Don't Use | ✓ Instead |
|-------------|-----------|
| "Loading..." | (No text needed) |
| "Please wait" | (No text needed) |
| "Preparing your experience" | (No text needed) |
| Spinning loader | Subtle logo pulse |
| Progress percentage | (No indicator needed) |
| Technical error codes | "Check your connection and try again" |

---

## 8. Accessibility Specifications

### 8.1 Screen Reader Support

```
On Splash Screen Appear:
├── Announce: "Unora is loading"
├── Role: progressbar (implicit)
└── aria-busy: true

On Transition Complete:
├── Announce: "Loading complete"
├── Focus: Move to first focusable element of next screen
└── aria-busy: false
```

### 8.2 Motion Sensitivity

```css
@media (prefers-reduced-motion: reduce) {
  .splash-logo {
    animation: none;
    opacity: 1;
  }
}
```

### 8.3 Color Contrast

| Element | Foreground | Background | Ratio | WCAG |
|---------|------------|------------|-------|------|
| Logo (golden) | #C4A77D | #FAFAF8 | 3.1:1 | ✓ AA (large) |
| Error headline | #1A1A1A | #FAFAF8 | 16:1 | ✓ AAA |
| Error body | #4A4A4A | #FAFAF8 | 8.5:1 | ✓ AAA |

---

## 9. Platform-Specific Notes

### 9.1 iOS

- **LaunchScreen.storyboard**: Implement as a static launch screen matching splash visuals
- **Safe areas**: Respect `env(safe-area-inset-top)` and `env(safe-area-inset-bottom)`
- **Status bar**: Use default style (dark content on light background)

### 9.2 Android

- **splash_screen.xml**: Configure via Android 12+ Splash Screen API
- **windowBackground**: Set to `--surface-background` color
- **Animated icon support**: Use animated vector drawable for logo pulse (Android 12+)
- **Legacy (< Android 12)**: Use themed activity with centered logo

### 9.3 Web (PWA)

- **manifest.json**: Set `background_color: "#FAFAF8"`
- **theme_color**: Set to `#FAFAF8`
- **splash_pages**: Generate appropriate splash images for iOS Safari

---

## 10. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Logo asset at correct size (64px height) | Critical | ☐ |
| Background color matches design token | Critical | ☐ |
| Logo centered vertically and horizontally | Critical | ☐ |
| Pulse animation (with reduced-motion support) | High | ☐ |
| Minimum display time (1200ms) | High | ☐ |
| Maximum display time / timeout (5000ms) | High | ☐ |
| Cross-fade transition to next screen | High | ☐ |
| Error state implementation | Medium | ☐ |
| Screen reader announcements | Medium | ☐ |
| Dark mode support | Medium | ☐ |
| Platform-specific launch screens | Medium | ☐ |
| Safe area compliance | Medium | ☐ |

---

## 11. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Core philosophy, brand positioning |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Color tokens, typography, animation specs |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Navigation flow, authentication gates |
| Unora_Spec_02_Welcome.md (planned) | Next screen in onboarding sequence |

---

**Document maintained by:** Product Design Team  
**Last updated:** January 2026  
**Next review:** Upon design system updates

---

*This specification is developer-ready and should be implemented as defined. Any deviations require design review.*
