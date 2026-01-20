# Unora UI Specification — Hobby Echo (Success)

**Document ID:** Spec-14  
**Screen Name:** Hobby Echo (Success)  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Hobby Echo (Success)** — Celebratory success state after mutual check-in

### 1.2 User Flow Reference
**Phase 4 (Streak Loop) → Mutual Success** — Displayed when both users complete their daily check-in.

**Sequence:**
```
User Check-In → Partner Check-In → [Hobby Echo] → Back to Streak Detail
```

**Reference:** [Unora_PRD.md — Section 14.2.1](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)

### 1.3 Purpose
Validate shared effort and build connection through **activity, not just conversation**. The Echo creates a sense of mutual presence — "We're both showing up."

### 1.4 The Echo Concept

Instead of generic success messages, the system displays a **context-aware summary** of partner activity:

| Generic ❌ | Echo ✓ |
|-----------|--------|
| "Good job!" | "Your partner crushed [Leg Day] today." |
| "Check-in complete" | "Your friend spent time [Reading Fiction]." |

> **Privacy Constraint:** The Echo is strictly limited to the **Activity Domain**. No location, time, or other metadata is shared.

---

## 2. Low-Fidelity Wireframes (ASCII)

### 2.1 Hobby Echo Modal

```
┌─────────────────────────────────────────────────────────────┐
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│  ← Streak Detail
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│     (dimmed)
│▒▒▒                                                     ▒▒▒│
│▒▒▒   ┌───────────────────────────────────────────┐     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │               ✨ 🔥 ✨                    │     ▒▒▒│  ← Animated icon
│▒▒▒   │                                           │     ▒▒▒│     (Server color)
│▒▒▒   │          Streak Extended!                 │     ▒▒▒│  ← Headline
│▒▒▒   │              Day 7                        │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ─────────────────────────────────────   │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ┌───────────────────────────────────┐   │     ▒▒▒│
│▒▒▒   │   │                                   │   │     ▒▒▒│
│▒▒▒   │   │   Your partner crushed           │   │     ▒▒▒│  ← Echo Card
│▒▒▒   │   │                                   │   │     ▒▒▒│
│▒▒▒   │   │        🏋️  Leg Day               │   │     ▒▒▒│  ← Activity
│▒▒▒   │   │                                   │   │     ▒▒▒│
│▒▒▒   │   │       "3 sets in, no excuses"     │   │     ▒▒▒│  ← Optional quote
│▒▒▒   │   │                                   │   │     ▒▒▒│
│▒▒▒   │   └───────────────────────────────────┘   │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   You checked in: Upper Body ✓           │     ▒▒▒│  ← User's activity
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ┌───────────────────────────────────┐   │     ▒▒▒│
│▒▒▒   │   │            Continue               │   │     ▒▒▒│  ← CTA
│▒▒▒   │   └───────────────────────────────────┘   │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   └───────────────────────────────────────────┘     ▒▒▒│
│▒▒▒                                                     ▒▒▒│
│▒▒▒             ✦  ✦  ✦  ✦  ✦  ✦  ✦                    ▒▒▒│  ← Confetti
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Echo Card Variants by Server

```
PARTNER SERVER (Terracotta):
┌───────────────────────────────────────┐
│   Your partner finished a             │
│                                       │
│        🔥  High Intensity Workout     │  ← Terracotta accent
│                                       │
└───────────────────────────────────────┘

FRIEND SERVER (Teal):
┌───────────────────────────────────────┐
│   Your friend spent time              │
│                                       │
│        👋  Reading Fiction            │  ← Teal accent
│                                       │
└───────────────────────────────────────┘

GROWTH SERVER (Indigo):
┌───────────────────────────────────────┐
│   Your accountability buddy tackled   │
│                                       │
│        🎯  Deep Work Session          │  ← Indigo accent
│                                       │
└───────────────────────────────────────┘
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
HOBBY ECHO MODAL
├── Type: Centered Modal Overlay
├── Position: Center of screen
├── Width: min(360px, 90vw)
├── Background: var(--surface-elevated) → #FFFFFF
├── Border radius: var(--radius-xl) → 20px
├── Shadow: var(--shadow-elevated) → 0 24px 64px rgba(0,0,0,0.25)
├── Padding: 32px (generous for celebration)
├── Z-Index: 100
│
├── [SCRIM]
│   ├── Background: rgba(0,0,0,0.6)
│   ├── Z-Index: 99
│   └── Non-interactive (modal required)
│
└── [CONFETTI LAYER]
    ├── Z-Index: 101
    └── Particles: Server color palette

Premium Dark Mode (Default):
├── Modal surface: var(--pdm-surface-3) → #222226 with radiant server glow
├── Backdrop: Deep blur (20px) with 90% opacity
├── Echo card: Server accent with inner glow
└── Confetti: Server colors + gold accent particles
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Backdrop** | Glass blur: `blur(20px)`, bg `rgba(13,13,15,0.9)` |
| **Modal surface** | Elevated `#222226` with radial server glow underneath |
| **Success icon** | Server color with pulsing glow: `0 0 32px rgba(server,0.4)` |
| **Headline** | High-contrast white with subtle text glow |
| **Echo card** | Server @ 15% bg, glowing server border |
| **Confetti** | Server color + gold accent, glowing particles |
| **Continue button** | Server gradient + outer glow |

**Celebratory Modal Glow:**
```css
/* Radiant glow underneath modal for celebration */
.hobby-echo-modal {
  box-shadow: 
    0 24px 64px rgba(0, 0, 0, 0.4),
    0 0 100px rgba(var(--server-color-rgb), 0.15);
}

/* Success icon breathing glow */
@keyframes celebration-glow {
  0%, 100% { filter: drop-shadow(0 0 24px rgba(var(--server-color-rgb), 0.3)); }
  50%      { filter: drop-shadow(0 0 40px rgba(var(--server-color-rgb), 0.5)); }
}
```



### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Modal padding | `var(--space-8)` | 32px |
| Icon margin-bottom | `var(--space-4)` | 16px |
| Headline margin-bottom | `var(--space-2)` | 8px |
| Day number margin-bottom | `var(--space-5)` | 20px |
| Echo card margin | `var(--space-5)` | 20px |
| Echo card padding | `var(--space-5)` | 20px |
| CTA margin-top | `var(--space-6)` | 24px |

### 3.3 Z-Index Stack

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Streak Detail | 1 | Background |
| Scrim | 99 | Dimmed overlay |
| Modal | 100 | Echo content |
| Confetti | 101 | Celebration particles |

---

## 4. Component Inventory

### 4.1 Dynamic Server Theming

| Server | Token | Hex | Applied To |
|--------|-------|-----|------------|
| **Partner** | `var(--server-partner-primary)` | #C9785D | Icon, Echo card border, confetti |
| **Friend** | `var(--server-friend-primary)` | #4A9B8C | Icon, Echo card border, confetti |
| **Growth** | `var(--server-growth-primary)` | #6B5B95 | Icon, Echo card border, confetti |

### 4.2 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Headline "Streak Extended!" | Outfit | 700 | 24px | `--unora-ink-primary` |
| Day number | Outfit | 600 | 18px | [Server color] |
| Echo intro | Inter | 400 | 16px | `--unora-ink-secondary` |
| Activity name | Outfit | 600 | 20px | `--unora-ink-primary` |
| Optional quote | Inter | 400 | 14px | `--unora-ink-tertiary` |
| User check-in | Inter | 500 | 14px | `--unora-ink-secondary` |
| Button text | Inter | 600 | 16px | White |

### 4.3 Success Icon (Animated)

```
SUCCESS ICON
├── Type: Flame 🔥 (Partner) / Wave 👋 (Friend) / Target 🎯 (Growth)
├── Size: 48px
├── Container: 80px with glow effect
├── Glow: [Server color] @ 20%, 120px blur
├── Sparkles: ✨ animated around icon
│
└── ANIMATION
    ├── Scale: 0 → 1.2 → 1.0 (bounce)
    ├── Rotation: Slight wobble
    ├── Duration: 500ms
    ├── Delay: 0ms (immediate)
    └── Sparkles: Staggered fade-in
```

### 4.4 Echo Card

**Reference:** [Unora_DesignSystem.md — Section 3.3](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

```
ECHO CARD
├── Width: 100%
├── Background: [server color] @ 8%
├── Border: 2px solid [server color]
├── Border radius: var(--radius-lg) → 16px
├── Padding: 20px
├── Text-align: center
│
├── [INTRO LINE]
│   └── "Your partner crushed" / "Your friend spent time"
│
├── [ACTIVITY]
│   ├── Icon: Hobby emoji (32px)
│   └── Name: Activity category (20px, bold)
│
└── [QUOTE] (optional)
    └── Micro-description in quotes
```

### 4.5 Continue Button

| Property | Value |
|----------|-------|
| Height | 52px |
| Width | 100% |
| Background | [Server color] |
| Text | White, Inter 16px/600 |
| Border radius | `var(--radius-md)` → 12px |
| Shadow | 0 4px 12px [server color] @ 20% |

### 4.6 Confetti Effect

```
CONFETTI PARTICLES
├── Count: 30-50 particles
├── Colors: [Server color], [Server color @ 60%], Gold accent
├── Size: 6-12px (varied)
├── Shape: Squares, circles, lines
├── Duration: 2-3 seconds
├── Trigger: On modal open
├── Easing: Gravity fall + drift
└── Layer: Above modal (Z: 101)
```

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Condition | Action |
|---------|-----------|--------|
| Auto-show | Partner completes check-in | Display Echo modal |
| Tap | Continue button | Close modal, show Streak Detail |
| Tap | Outside modal | No action (modal required) |

### 5.2 Echo Generation Flow

```
PARTNER completes their check-in
    │
    ▼
┌─────────────────────────────────────────┐
│ SYSTEM:                                 │
│ ├── Detect mutual check-in complete     │
│ ├── Retrieve partner's activity         │
│ ├── Map to Echo category (privacy)      │
│ ├── Generate Echo text                  │
│ └── Trigger celebration modal           │
└─────────────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────────────┐
│ DISPLAY:                                │
│ ├── Haptic: Heavy/Success sequence      │
│ ├── Modal: Scale up + fade in           │
│ ├── Confetti: Staggered particle burst  │
│ ├── Icon: Bounce animation              │
│ └── Echo card: Fade in (200ms delay)    │
└─────────────────────────────────────────┘
```

### 5.3 Privacy Mapping

> **CRITICAL:** Raw input is never displayed. Activity is mapped to safe categories.

| Raw Input | Echo Display |
|-----------|--------------|
| "Did squats and deadlifts" | "Leg Day" |
| "Read 50 pages of novel" | "Reading Fiction" |
| "2 hour focus session" | "Deep Work Session" |

### 5.4 Haptic Sequence

**Reference:** [Unora_DesignSystem.md — Section 5.1](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

| Event | Haptic |
|-------|--------|
| Modal open | Heavy impact |
| Confetti burst | Triple light tap sequence |
| Button tap | Medium impact |

### 5.5 Entry Animation

```
MODAL ENTRY (Celebratory)
├── Scrim: Fade in, 200ms
├── Modal: Scale 0.9 → 1.0 + Fade, 400ms, ease-out-back
├── Icon: Scale 0 → 1.2 → 1.0, 500ms (bouncy)
├── Confetti: Burst after 300ms
├── Echo card: Fade in after 400ms
└── Button: Fade in after 600ms
```

### 5.6 Exit Animation

```
MODAL EXIT
├── Modal: Scale 1.0 → 0.95 + Fade, 200ms
├── Scrim: Fade out, 150ms
└── Return: Streak Detail (State C)
```

---

## 6. State Definitions

### 6.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Entering | Animation in progress | Modal opening |
| Active | Full celebration display | Awaiting user action |
| Exiting | Animation out | Button tapped |

### 6.2 Active State (Default)

```
Modal: Visible, centered
Icon: Animated, glowing
Headline: "Streak Extended!"
Echo card: Partner's activity displayed
Confetti: Falling particles
Button: "Continue" enabled
Scrim: Non-interactive
```

---

## 7. Content & Copy Guidelines

### 7.1 Headline

| Format | Example |
|--------|---------|
| Standard | "Streak Extended!" |
| Milestone (Day 5) | "Streak Extended! 🎉" |
| Milestone (Day 15) | "Final Day Complete! 🏆" |

### 7.2 Echo Intro by Server

| Server | Intro Format |
|--------|--------------|
| Partner | "Your partner crushed" / "Your partner finished" |
| Friend | "Your friend spent time" / "Your friend focused on" |
| Growth | "Your accountability buddy tackled" / "Your buddy worked on" |

### 7.3 Echo Examples

**Partner Server:**
| Activity | Echo |
|----------|------|
| Gym - Leg Day | "Your partner crushed **Leg Day**" |
| Running | "Your partner finished a **Morning Run**" |
| Yoga | "Your partner centered with **Yoga**" |

**Friend Server:**
| Activity | Echo |
|----------|------|
| Reading | "Your friend spent time **Reading Fiction**" |
| Gaming | "Your friend enjoyed **Game Time**" |
| Art | "Your friend created some **Art**" |

**Growth Server:**
| Activity | Echo |
|----------|------|
| Deep Work | "Your buddy tackled **Deep Work**" |
| Learning | "Your buddy invested in **Learning**" |
| Habit | "Your buddy maintained **Morning Routine**" |

### 7.4 Button Labels

| Context | Label |
|---------|-------|
| Standard | "Continue" |
| With tomorrow preview | "See Tomorrow's Goal" |

### 7.5 Privacy Note (Internal)

> **DEV NOTE:** Ensure no raw input data is displayed. Only mapped activity categories should appear in the Echo. User's micro-descriptions should be pre-processed before display.

---

## 8. Accessibility

### 8.1 Screen Reader
- Modal: "Streak Extended! Day 7. Your partner crushed Leg Day. You checked in: Upper Body. Continue button."
- Confetti: Ignored (decorative)

### 8.2 Reduced Motion
- If `prefers-reduced-motion`: Disable confetti, scale animations
- Keep essential feedback (color changes)

### 8.3 Touch Targets
- Continue button: 52px height, full width

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Centered modal overlay | Critical | ☐ |
| Server-themed styling | Critical | ☐ |
| Success icon animation | High | ☐ |
| Echo card with partner activity | Critical | ☐ |
| Privacy mapping (no raw data) | Critical | ☐ |
| Confetti effect | Medium | ☐ |
| Haptic sequence | High | ☐ |
| Entry/exit animations | High | ☐ |
| Continue button | High | ☐ |
| Reduced motion support | Medium | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 14.2.1 — Hobby Echo |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 3.3 — Modal Card, Section 5.1 — Haptics |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.D — Streak Loop |
| [Unora_Spec_12_StreakDetail.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_12_StreakDetail.md) | Parent screen (State C) |
| [Unora_Spec_13_CheckInAction.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_13_CheckInAction.md) | Trigger flow |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
