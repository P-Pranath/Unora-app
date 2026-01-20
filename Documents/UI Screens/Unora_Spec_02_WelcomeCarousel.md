# Unora UI Specification — Welcome / Philosophy Carousel

**Document ID:** Spec-02  
**Screen Name:** Welcome / Philosophy Carousel  
**Version:** 1.0  
**Date:** January 2026  
**Classification:** Internal / Engineering / Design  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name

**Welcome / Philosophy Carousel** (also: Onboarding Carousel, Philosophy Introduction)

### 1.2 User Flow Reference

**Phase 1: Verified Onboarding** — The Welcome Carousel is the first interactive experience after the Splash Screen for new or logged-out users. It introduces Unora's core philosophy before the user begins profile creation.

**Sequence Position:**
```
Splash Screen → [Welcome Carousel] → Profile Creation → Identity Verification → Server Selection → Discovery
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.A (Onboarding & Hard-Gate Verification)](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose

The Welcome Carousel educates new users on Unora's "Anti-Tinder" philosophy — **anonymous discovery, earned communication, and gradual reveals** — setting expectations before they commit to the verification-gated onboarding process.

### 1.4 Primary Action

**Navigation through carousel slides** via swipe gesture or "Next" button, culminating in a "Get Started" CTA that initiates profile creation.

---

## 2. Low-Fidelity Wireframe (ASCII)

### 2.1 Welcome Landing (Pre-Carousel)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │   ← System status bar
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                   ┌───────────────┐                         │
│                   │               │                         │
│                   │   [ UNORA ]   │                         │   ← App Logo (64px)
│                   │     LOGO      │                         │
│                   │               │                         │
│                   └───────────────┘                         │
│                                                             │
│                                                             │
│         Connection worth waiting for.                       │   ← Tagline (H2)
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                   Get Started                       │   │   ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│               Already have an account?                      │   ← Caption + link
│                    Sign in →                                │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │   ← Safe area (bottom)
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Philosophy Carousel Slide (Generic Structure)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │   ← System status bar
│                                                             │
│                                                             │
│                                                             │
│              ┌─────────────────────────┐                    │
│              │                         │                    │
│              │    ┌─────────────────┐  │                    │
│              │    │                 │  │                    │
│              │    │  [Illustration] │  │                    │   ← Illustration area
│              │    │                 │  │                    │     200px height max
│              │    └─────────────────┘  │                    │
│              │                         │                    │
│              └─────────────────────────┘                    │
│                                                             │
│                                                             │
│                                                             │
│         No photos. Just hobbies.                            │   ← Headline (H2)
│                                                             │
│         Discover people through                             │   ← Body text
│         what they do, not how they look.                    │
│                                                             │
│                                                             │
│                       ● ○ ○                                 │   ← Pagination dots
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                     Next                            │   │   ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                      Skip →                                 │   ← Tertiary link
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │   ← Safe area (bottom)
└─────────────────────────────────────────────────────────────┘
```

### 2.3 Final Carousel Slide

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│                                                             │
│              ┌─────────────────────────┐                    │
│              │                         │                    │
│              │    ┌─────────────────┐  │                    │
│              │    │                 │  │                    │
│              │    │  [Illustration] │  │                    │   ← Reveals unlocking
│              │    │                 │  │                    │
│              │    └─────────────────┘  │                    │
│              │                         │                    │
│              └─────────────────────────┘                    │
│                                                             │
│                                                             │
│         Earned, not given.                                  │   ← Headline (H2)
│                                                             │
│         Identity is revealed gradually                      │   ← Body text
│         as trust builds.                                    │
│                                                             │
│                                                             │
│                       ○ ○ ●                                 │   ← Pagination dots
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                  Get Started                        │   │   ← Primary CTA (changed)
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
├─────────────────────────────────────────────────────────────┤
└─────────────────────────────────────────────────────────────┘
```

### Layout Constraint Check

| Validation | Result |
|------------|--------|
| **Layout Type** | Vertical Stack (Discovery pattern per DSD v1.2 Section 1.2) |
| **DSD Alignment** | Matches onboarding philosophy (DSD Section 9.3) |
| **Single Focus** | ✓ One primary CTA per slide |
| **No Modal** | ✓ Full-screen carousel, not modal overlay |
| **Minimal UI** | ✓ Clean visual hierarchy, maximum whitespace |

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  CAROUSEL SCREEN CONTAINER                                                    │
│                                                                               │
│  ├── Position: fixed                                                          │
│  ├── Width: 100vw                                                             │
│  ├── Height: 100vh                                                            │
│  ├── Display: flex                                                            │
│  ├── Flex-direction: column                                                   │
│  ├── Background: var(--surface-background) → #FAFAF8                          │
│  │                                                                            │
│  ├── [SAFE AREA TOP]                                                          │
│  │   └── Height: env(safe-area-inset-top) + var(--space-4) → 16px             │
│  │                                                                            │
│  ├── [CONTENT AREA]                                                           │
│  │   ├── Flex: 1 (fills available space)                                      │
│  │   ├── Padding: var(--padding-screen) → 20px (horizontal)                   │
│  │   ├── Display: flex                                                        │
│  │   ├── Flex-direction: column                                               │
│  │   └── Justify-content: center                                              │
│  │                                                                            │
│  ├── [ACTION AREA]                                                            │
│  │   ├── Padding: var(--padding-screen) → 20px (all sides)                    │
│  │   ├── Padding-bottom: var(--space-8) → 40px                                │
│  │   └── Gap: var(--gap-stack) → 12px                                         │
│  │                                                                            │
│  └── [SAFE AREA BOTTOM]                                                       │
│      └── Height: env(safe-area-inset-bottom)                                  │
│                                                                               │
│  Premium Dark Mode (Default):                                                 │
│  ├── Background: var(--pdm-background) → #0D0D0F                              │
│  ├── Subtle top-gradient overlay for depth:                                   │
│  │   linear-gradient(180deg, rgba(20,20,22,0.8) 0%, transparent 30%)          │
│  └── Illustration cards receive subtle inner glow                             │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` with subtle top gradient for depth |
| **Logo** | Soft ambient glow: `0 0 32px rgba(196, 167, 125, 0.2)` |
| **Illustration cards** | Elevated surface `#1A1A1E` with inner glow `inset 0 1px 0 rgba(255,255,255,0.03)` |
| **CTA Button** | Gold gradient with outer glow (see DSD Section 12.5) |
| **Pagination dots** | Active dot: Gold with subtle glow `0 0 8px rgba(196,167,125,0.4)` |

### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen horizontal padding | `var(--padding-screen)` | 20px |
| Content section gap | `var(--gap-section)` | 32px |
| Stack gap (between text elements) | `var(--gap-stack)` | 12px |
| Illustration container margin-bottom | `var(--space-8)` | 40px |
| Pagination dots margin | `var(--space-6)` | 24px (top & bottom) |
| CTA button margin-top | `var(--space-6)` | 24px |
| Tertiary link margin-top | `var(--space-4)` | 16px |

### 3.3 Z-Index Layers

| Layer Name | Z-Index | Contents |
|------------|---------|----------|
| **Background Layer** | 0 | Solid/gradient background |
| **Content Layer** | 1 | Illustrations, text, pagination |
| **Action Layer** | 2 | CTA buttons, tertiary links |
| **System Layer** | 100+ | iOS/Android status bar (system-controlled) |

> [!NOTE]
> The Welcome Carousel uses only flat layers — no overlays or modals are required.

---

## 4. Component Inventory

### 4.1 Logo Component (Welcome Landing Only)

**Component Name:** Brand Logo (Centered)

| Property | Value |
|----------|-------|
| **Asset Type** | SVG (preferred) or PNG @3x |
| **Size** | 64px height (width proportional) |
| **Max Width** | 180px |
| **Position** | Centered horizontally |
| **Color** | `var(--unora-primary-accent)` → #C4A77D |
| **Alt Text** | "Unora" |
| **Dark Mode** | Use light variant (#F5F5F3) |

### 4.2 Illustration Container

**Component Name:** Philosophy Illustration

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  ILLUSTRATION CONTAINER                                                       │
│                                                                               │
│  ├── Width: 100% (max 280px)                                                  │
│  ├── Height: 200px (max)                                                      │
│  ├── Position: Centered horizontally                                          │
│  ├── Margin-bottom: var(--space-8) → 40px                                     │
│  ├── Object-fit: contain                                                      │
│  └── Background: transparent (illustrations are flat/vector)                  │
│                                                                               │
│  Illustration Style:                                                          │
│  ├── Flat vector design                                                       │
│  ├── Primary palette: var(--unora-primary), var(--unora-primary-accent)       │
│  ├── Accent colors: Subtle server theme hints (optional)                      │
│  ├── No gradients or 3D effects                                               │
│  └── Warm, approachable, minimal                                              │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 4.3 Typography Specifications

| Element | Font Family | Weight | Size | Line Height | Letter Spacing | Color |
|---------|-------------|--------|------|-------------|----------------|-------|
| **Tagline** | `var(--font-display)` → Outfit | 600 | 22px | 1.25 | -0.01em | `var(--unora-ink-primary)` → #1A1A1A |
| **Headline (H2)** | `var(--font-display)` → Outfit | 600 | 22px | 1.25 | -0.01em | `var(--unora-ink-primary)` → #1A1A1A |
| **Body Text** | `var(--font-body)` → Inter | 400 | 16px | 1.5 | 0 | `var(--unora-ink-secondary)` → #4A4A4A |
| **Caption** | `var(--font-body)` → Inter | 400 | 14px | 1.5 | 0 | `var(--unora-ink-tertiary)` → #7A7A7A |
| **Tertiary Link** | `var(--font-body)` → Inter | 500 | 14px | 1.5 | 0 | `var(--unora-primary-accent)` → #C4A77D |

### 4.4 Button Components

#### Primary Button ("Get Started" / "Next")

| Property | Value | Reference |
|----------|-------|-----------|
| **Height** | 52px | DSD Section 3.1 |
| **Width** | Full width - 40px (screen padding) | — |
| **Border Radius** | `var(--radius-md)` → 12px | DSD Section 2.4 |
| **Background** | `var(--unora-primary-accent)` → #C4A77D | Brand neutral (no server context) |
| **Text Color** | #FFFFFF | — |
| **Font** | Inter 16px / 600 | DSD Section 3.1 |
| **Padding** | 16px 24px | DSD Section 3.1 |
| **Shadow** | 0 2px 8px rgba(0,0,0,0.08) | DSD Section 3.1 |

**Button States:**

| State | Appearance |
|-------|------------|
| **Default** | Full opacity, background #C4A77D |
| **Pressed** | Scale 0.98, opacity 0.9 |
| **Disabled** | Opacity 0.4, non-interactive |
| **Loading** | Spinner replaces text, text fades to 60% |

#### Tertiary Button ("Skip" / "Sign in")

| Property | Value |
|----------|-------|
| **Height** | 44px (touch target) |
| **Background** | Transparent |
| **Text Color** | `var(--unora-primary-accent)` → #C4A77D |
| **Font** | Inter 14px / 500 |
| **Underline** | None (optional on hover for web) |

### 4.5 Pagination Dots

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  PAGINATION DOTS                                                              │
│                                                                               │
│  ├── Container: Centered horizontally                                         │
│  ├── Gap between dots: var(--space-2) → 8px                                   │
│  ├── Margin: var(--space-6) → 24px (top & bottom)                             │
│  │                                                                            │
│  ├── ACTIVE DOT                                                               │
│  │   ├── Size: 8px × 8px                                                      │
│  │   ├── Border Radius: var(--radius-full) → 9999px                           │
│  │   ├── Background: var(--unora-primary-accent) → #C4A77D                    │
│  │   └── Transition: background 200ms ease                                    │
│  │                                                                            │
│  └── INACTIVE DOT                                                             │
│      ├── Size: 8px × 8px                                                      │
│      ├── Border Radius: var(--radius-full) → 9999px                           │
│      ├── Background: var(--border-medium) → #D4D4D0                           │
│      └── Transition: background 200ms ease                                    │
│                                                                               │
│  Dark Mode:                                                                   │
│  ├── Active: var(--unora-primary-accent) (unchanged)                          │
│  └── Inactive: var(--dm-border-medium) → #3D3D3D                              │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 4.6 Color Tokens Summary

| Token | Light Mode | Dark Mode | Usage |
|-------|------------|-----------|-------|
| `--surface-background` | #FAFAF8 | #121212 | Screen background |
| `--unora-primary-accent` | #C4A77D | #C4A77D | CTA buttons, active dots, links |
| `--unora-ink-primary` | #1A1A1A | #F5F5F3 | Headlines |
| `--unora-ink-secondary` | #4A4A4A | #C4C4C0 | Body text |
| `--unora-ink-tertiary` | #7A7A7A | #8A8A86 | Captions |
| `--border-medium` | #D4D4D0 | #3D3D3D | Inactive pagination dots |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| **Tap** | "Get Started" button (Welcome) | Navigate to first carousel slide |
| **Tap** | "Next" button (Carousel) | Navigate to next slide |
| **Tap** | "Get Started" button (Final) | Navigate to Profile Creation |
| **Tap** | "Skip" link | Navigate directly to Profile Creation |
| **Tap** | "Sign in" link | Navigate to Sign In screen |
| **Swipe Left** | Carousel content area | Navigate to next slide |
| **Swipe Right** | Carousel content area | Navigate to previous slide |
| **Tap** | Pagination dot | Navigate to corresponding slide |

### 5.2 Behaviors

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  WELCOME CAROUSEL BEHAVIOR FLOW                                               │
│                                                                               │
│  USER ARRIVES FROM SPLASH SCREEN                                              │
│        │                                                                      │
│        ▼                                                                      │
│  ┌─────────────────────────────────────────────────────────────────────────┐  │
│  │  WELCOME LANDING SCREEN                                                 │  │
│  │                                                                         │  │
│  │  User sees:                                                             │  │
│  │  ├── Unora logo                                                         │  │
│  │  ├── Tagline: "Connection worth waiting for."                           │  │
│  │  ├── "Get Started" CTA                                                  │  │
│  │  └── "Already have an account? Sign in" link                            │  │
│  └─────────────────────────────────────────────────────────────────────────┘  │
│        │                                                                      │
│        ├── [User taps "Get Started"]                                          │
│        │         │                                                            │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  CAROUSEL SLIDE 1: "No photos. Just hobbies."                │  │
│        │   │                                                               │  │
│        │   │  Philosophy: Anonymous Discovery                              │  │
│        │   │  Message: Discover people through what they do,               │  │
│        │   │           not how they look.                                  │  │
│        │   │                                                               │  │
│        │   │  Pagination: ● ○ ○                                            │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │         │                                                            │
│        │         ├── [Swipe left OR tap "Next"]                               │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  CAROUSEL SLIDE 2: "15 days of showing up."                  │  │
│        │   │                                                               │  │
│        │   │  Philosophy: Earned Communication                             │  │
│        │   │  Message: Chat unlocks after 15 days of                       │  │
│        │   │           mutual consistency.                                 │  │
│        │   │                                                               │  │
│        │   │  Pagination: ○ ● ○                                            │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │         │                                                            │
│        │         ├── [Swipe left OR tap "Next"]                               │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  CAROUSEL SLIDE 3: "Earned, not given."                      │  │
│        │   │                                                               │  │
│        │   │  Philosophy: Gradual Reveals                                  │  │
│        │   │  Message: Identity is revealed gradually                      │  │
│        │   │           as trust builds.                                    │  │
│        │   │                                                               │  │
│        │   │  Pagination: ○ ○ ●                                            │  │
│        │   │  CTA: "Get Started" (changed from "Next")                     │  │
│        │   │  Skip link: HIDDEN on final slide                             │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │         │                                                            │
│        │         ├── [Tap "Get Started"]                                      │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  NAVIGATE TO: Profile Creation Screen                        │  │
│        │   │  Reference: Phase 1 Onboarding (PRD Section 10.1)            │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │                                                                      │
│        ├── [User taps "Sign in"]                                              │
│        │         │                                                            │
│        │         └──→ Navigate to Sign In Screen                              │
│        │                                                                      │
│        └── [User taps "Skip" on any carousel slide]                           │
│                  │                                                            │
│                  └──→ Navigate directly to Profile Creation Screen             │
│                                                                               │
└───────────────────────────────────────────────────────────────────────────────┘
```

**Reference:** 
- [Unora_PRD.md — Section 7 (Core Philosophy)](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)
- [Unora_UserFlow_Logic.md — Section 2.A.1 (Profile Creation)](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 5.3 Transitions

#### Slide-to-Slide Transition (Swipe)

| Property | Value | Reference |
|----------|-------|-----------|
| **Animation Type** | Slide + Fade | — |
| **Duration** | 300ms | DSD Section 5.4 |
| **Easing** | ease-out | DSD Section 5.4 |
| **Direction** | Horizontal (follow swipe direction) |

```
Swipe Left Transition:
[0ms]      Gesture begins
[gesture]  Content translates with finger position, next slide visible
[release]  If threshold met: animate to next slide position
           Outgoing: translateX(-100%), opacity 1 → 0.8
           Incoming: translateX(100% → 0), opacity 0.8 → 1
[300ms]    Transition complete
           Pagination dot updates
```

#### Button Tap Transition (Next/Get Started)

| Property | Value |
|----------|-------|
| **Animation Type** | Cross-fade + subtle slide |
| **Duration** | 250ms |
| **Easing** | ease-out |

```
Button Tap "Next":
[0ms]      Button pressed, scale to 0.98
[50ms]     Haptic: light tap
[80ms]     Button returns to scale 1.0
[100ms]    Current slide fades out, slides left slightly
[250ms]    Next slide fully visible
           Pagination updates
```

#### Welcome → Carousel Transition

| Property | Value |
|----------|-------|
| **Animation Type** | Cross-fade |
| **Duration** | 200ms |
| **Easing** | ease-out |

```
"Get Started" on Welcome Landing:
[0ms]      Button tap registered
[0ms]      Welcome screen begins fade-out
[0ms]      First carousel slide begins fade-in
[200ms]    Transition complete
```

#### Carousel → Profile Creation Transition

| Property | Value | Reference |
|----------|-------|-----------|
| **Animation Type** | Slide from right | DSD Section 5.4 (Screen push) |
| **Duration** | 300ms | DSD Section 5.4 |
| **Easing** | ease-out | — |

> [!TIP]
> Respect `prefers-reduced-motion`: When enabled, use instant cross-fade with no slide animations.

---

## 6. State Definitions

### 6.1 State Matrix

| State | Visual Appearance | Conditions | Duration |
|-------|-------------------|------------|----------|
| **Default** | Welcome landing with CTA | First-time/logged-out user | Until user action |
| **Carousel Active** | Philosophy slide visible | User navigating slides | Until completion or skip |
| **Loading (Skeleton)** | Not applicable | — | — |
| **Empty** | Not applicable | — | — |
| **Error** | Error message card | Failed to load assets | Until retry/dismiss |
| **Locked/Disabled** | Not applicable | — | — |
| **Cooldown** | Not applicable | — | — |

### 6.2 Default State (Welcome Landing)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│                                                             │
│                   ┌───────────────┐                         │
│                   │               │                         │
│                   │   [ UNORA ]   │                         │   ← Logo (static or
│                   │               │                         │      subtle breathe)
│                   └───────────────┘                         │
│                                                             │
│                                                             │
│         Connection worth waiting for.                       │   ← Tagline
│                                                             │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                   Get Started                       │   │   ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│               Already have an account?                      │
│                    Sign in →                                │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Background: var(--surface-background) → #FAFAF8
Gradient (optional): Subtle brand sand → white
Tone: Welcoming, calm, premium
```

### 6.3 Carousel Active State

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│              ┌─────────────────────────┐                    │
│              │                         │                    │
│              │    [Slide-specific      │                    │
│              │     Illustration]       │                    │   ← 200px max height
│              │                         │                    │
│              └─────────────────────────┘                    │
│                                                             │
│                                                             │
│          [Slide Headline]                                   │   ← H2, centered
│                                                             │
│          [Slide Body Text]                                  │   ← Body, centered
│          [Continuation if needed]                           │
│                                                             │
│                                                             │
│                       ● ○ ○                                 │   ← Pagination
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                     Next                            │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                      Skip →                                 │   ← Tertiary (slides 1-2)
│                                                             │
└─────────────────────────────────────────────────────────────┘

Swipe gestures: Active (left = next, right = previous)
Back navigation: Available on slides 2-3
```

### 6.4 Error State (Asset Load Failure)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│                                                             │
│                        ⚠️                                   │   ← Icon: 32px
│                                                             │       var(--feedback-warning)
│                                                             │
│          Something went wrong                               │   ← H4 style
│                                                             │
│   We couldn't load the welcome experience.                  │   ← Body style, muted
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                    Try Again                        │   │   ← Secondary button
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                   Skip to Sign Up →                         │   ← Tertiary fallback
│                                                             │
└─────────────────────────────────────────────────────────────┘

Error styling:
├── Background: var(--surface-background)
├── Icon: Warning, var(--feedback-warning) → #E6A43A
├── No red colors or alarming language
└── Fallback allows user to bypass carousel
```

> [!IMPORTANT]
> Error state should never block onboarding entirely. The "Skip to Sign Up" link ensures users can always proceed to create an account.

---

## 7. Content & Copy Guidelines

### 7.1 Text Strings

#### Welcome Landing

| Element | Copy | Notes |
|---------|------|-------|
| **Tagline** | "Connection worth waiting for." | Central brand promise |
| **Primary CTA** | "Get Started" | Action-oriented |
| **Secondary Link** | "Already have an account? Sign in →" | Returning users |

#### Carousel Slide 1: Anonymous Discovery

| Element | Copy | Philosophy |
|---------|------|------------|
| **Headline** | "No photos. Just hobbies." | Anti-appearance-first |
| **Body** | "Discover people through what they do, not how they look." | Hobby-anchored matching |
| **CTA** | "Next" | — |
| **Skip** | "Skip →" | — |

#### Carousel Slide 2: Earned Communication

| Element | Copy | Philosophy |
|---------|------|------------|
| **Headline** | "15 days of showing up." | Commitment-first |
| **Body** | "Chat unlocks after 15 days of mutual consistency." | Delayed gratification |
| **CTA** | "Next" | — |
| **Skip** | "Skip →" | — |

#### Carousel Slide 3: Gradual Reveals

| Element | Copy | Philosophy |
|---------|------|------------|
| **Headline** | "Earned, not given." | Trust-building |
| **Body** | "Identity is revealed gradually as trust builds." | Progressive disclosure |
| **CTA** | "Get Started" | Final action |
| **Skip** | Hidden | Force commitment at end |

### 7.2 Tone Guidelines

The Welcome Carousel embodies Unora's core philosophy:

| Principle | Application |
|-----------|-------------|
| **Presence over Performance** | Illustrations are warm and human, not gamified or flashy |
| **Anticipation over Gratification** | Copy speaks to delayed reward and building something real |
| **Clarity over Clutter** | One concept per slide, minimal text, maximum impact |
| **Anti-Tinder Positioning** | Explicitly contrasts with photo-first, instant-chat platforms |

**Reference:** [Unora_PRD.md — Section 7 (Core Philosophy)](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)

### 7.3 Illustration Guidance

| Slide | Illustration Concept | Elements |
|-------|---------------------|----------|
| **Slide 1** | Anonymous hobby cards | Abstract cards floating, hobby icons visible (🏋️ 🎨 📚), no faces |
| **Slide 2** | Calendar/streak visualization | 15-day path with checkmarks, two figures maintaining streak |
| **Slide 3** | Progressive reveal locks | Lock icons opening progressively, silhouette becoming identity |

**Style Requirements:**
- Flat vector, minimal detail
- Warm palette matching brand colors
- No realistic faces until Slide 3 (silhouettes only)
- Inclusive representation (diverse body types, no specific demographics)

### 7.4 Avoid

| ❌ Don't Use | ✓ Instead |
|-------------|-----------|
| "Swipe right on your future" | "Discover people through what they do" |
| "Find your match instantly" | "Connection worth waiting for" |
| "Unlimited profiles" | "5 intentional options daily" |
| "Chat now" | "Earn the conversation" |
| Tech jargon | Simple, human language |
| Urgency language | Patience-oriented messaging |

---

## 8. Accessibility Specifications

### 8.1 Screen Reader Support

```
Welcome Landing:
├── Announce: "Unora. Connection worth waiting for. Button: Get Started."
├── Focus order: Logo (decorative) → Tagline → Get Started → Sign in
└── aria-label on logo: "Unora logo"

Carousel Slides:
├── Announce: "Slide 1 of 3. [Headline]. [Body text]. Button: Next. Link: Skip."
├── Illustration: aria-hidden="true" (decorative)
├── Pagination: role="tablist" with aria-label="Carousel navigation"
├── Each dot: role="tab", aria-selected, aria-label="Go to slide X"
└── Content area: role="tabpanel"

Swipe Gesture:
└── Provide visual Next/Previous buttons for users who cannot swipe
```

### 8.2 Motion Sensitivity

```css
@media (prefers-reduced-motion: reduce) {
  .carousel-slide {
    transition: opacity 0.01ms;
    transform: none;
  }
  
  .pagination-dot {
    transition: none;
  }
  
  .illustration {
    animation: none;
  }
}
```

### 8.3 Color Contrast

| Element | Foreground | Background | Ratio | WCAG |
|---------|------------|------------|-------|------|
| Headline | #1A1A1A | #FAFAF8 | 16:1 | ✓ AAA |
| Body text | #4A4A4A | #FAFAF8 | 8.5:1 | ✓ AAA |
| CTA button text | #FFFFFF | #C4A77D | 3.2:1 | ✓ AA (large) |
| Tertiary link | #C4A77D | #FAFAF8 | 3.1:1 | ✓ AA (large) |
| Active dot | #C4A77D | #FAFAF8 | 3.1:1 | ✓ AA |

### 8.4 Touch Targets

| Component | Touch Size | Visual Size |
|-----------|------------|-------------|
| Primary CTA | Full width × 52px | — |
| Tertiary link | 44px height minimum | Text size |
| Pagination dots | 44px × 44px | 8px visual |
| Swipe area | Full content area | — |

---

## 9. Platform-Specific Notes

### 9.1 iOS

- **Safe areas**: Respect `env(safe-area-inset-top)` and bottom
- **Swipe gestures**: Use `UIPageViewController` or SwiftUI `TabView`
- **Back gesture**: Disable edge-swipe-to-back on carousel (custom handling)
- **Status bar**: Light content for dark mode, dark content for light mode

### 9.2 Android

- **ViewPager2**: Use for horizontal carousel implementation
- **Edge-to-edge**: Apply WindowInsets for safe areas
- **Back button**: Handle carousel navigation (previous slide or exit)
- **Gesture navigation**: Ensure swipe doesn't conflict with system gestures

### 9.3 Web (PWA)

- **Keyboard navigation**: Arrow keys for slide navigation
- **Touch/Mouse**: Support both swipe and click on dots/buttons
- **Responsive**: Stack layout maintained at all breakpoints
- **Focus indicators**: Visible focus rings on all interactive elements

---

## 10. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Welcome landing with logo and tagline | Critical | ☐ |
| 3 carousel slides with correct copy | Critical | ☐ |
| Illustrations for each slide | High | ☐ |
| Swipe gesture navigation | Critical | ☐ |
| Pagination dots with active state | High | ☐ |
| "Next" button functionality | Critical | ☐ |
| "Get Started" CTA on final slide | Critical | ☐ |
| "Skip" link on slides 1-2 | High | ☐ |
| "Sign in" link on welcome landing | High | ☐ |
| Slide transitions (300ms) | High | ☐ |
| Dark mode support | Medium | ☐ |
| Screen reader announcements | Medium | ☐ |
| Reduced motion support | Medium | ☐ |
| Error state implementation | Medium | ☐ |
| Safe area compliance | Medium | ☐ |
| Touch target minimums (44px) | Medium | ☐ |

---

## 11. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Core philosophy, brand positioning, Anti-Tinder thesis |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Color tokens, typography, animation specs, Section 9.3 |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Navigation flow, Phase 1 onboarding |
| [Unora_Spec_01_Splash.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_01_Splash.md) | Previous screen in sequence |
| Unora_Spec_03_ProfileCreation.md (planned) | Next screen after Welcome Carousel |

---

**Document maintained by:** Product Design Team  
**Last updated:** January 2026  
**Next review:** Upon design system updates

---

*This specification is developer-ready and should be implemented as defined. Any deviations require design review.*
