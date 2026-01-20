# Unora UI Specification — Card Detail Modal

**Document ID:** Spec-08  
**Screen Name:** Card Detail Modal  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Card Detail Modal** — Expanded profile view overlay

### 1.2 User Flow Reference
**Phase 2 (Discovery) → Phase 4 (Connection)** — This modal provides the detailed view needed to decide whether to express interest.

**Sequence:**
```
Discovery Feed → [Card Detail Modal] → "Connect" → Interest Sent → Active Streak
                        ↓
                    "Close" → Back to Discovery Feed
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.C.2](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Provide sufficient depth about a potential connection — hobbies with context, consistency profile, and intent statement — to help users make a meaningful decision before expressing interest.

### 1.4 Primary Action
- **Connect** — Express interest in this person
- **Close** — Return to Discovery Feed

---

## 2. Low-Fidelity Wireframe (ASCII)

### 2.1 Modal Overlay Structure

```
┌─────────────────────────────────────────────────────────────┐
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│  ← Dimmed background
│▒▒▒                                                     ▒▒▒│     (Discovery Feed)
│▒▒▒   ┌───────────────────────────────────────────┐     ▒▒▒│
│▒▒▒   │                                     ✕     │     ▒▒▒│  ← Close button
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ┌─────────────────────────────────┐     │     ▒▒▒│
│▒▒▒   │   │  🏋️  Gym                        │     │     ▒▒▒│  ← Hobby Anchor 1
│▒▒▒   │   │  "Consistency over intensity"   │     │     ▒▒▒│
│▒▒▒   │   └─────────────────────────────────┘     │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ┌─────────────────────────────────┐     │     ▒▒▒│
│▒▒▒   │   │  🎨  Painting                   │     │     ▒▒▒│  ← Hobby Anchor 2
│▒▒▒   │   │  "Late nights, slow strokes"    │     │     ▒▒▒│
│▒▒▒   │   └─────────────────────────────────┘     │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ─────────────────────────────────────   │     ▒▒▒│  ← Divider
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   Consistency                             │     ▒▒▒│
│▒▒▒   │   ┌─────────────────────────────────┐     │     ▒▒▒│
│▒▒▒   │   │  ████████████████░░░░░░░░       │     │     ▒▒▒│  ← Consistency Band
│▒▒▒   │   │         Regular                 │     │     ▒▒▒│
│▒▒▒   │   └─────────────────────────────────┘     │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ─────────────────────────────────────   │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   What I'm looking for                    │     ▒▒▒│
│▒▒▒   │   "Someone who wants to grow together,    │     ▒▒▒│  ← Intent Statement
│▒▒▒   │    not just pass the time."               │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ┌─────────────────────────────────┐     │     ▒▒▒│
│▒▒▒   │   │            Connect              │     │     ▒▒▒│  ← Primary CTA
│▒▒▒   │   └─────────────────────────────────┘     │     ▒▒▒│     (Server themed)
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   └───────────────────────────────────────────┘     ▒▒▒│
│▒▒▒                                                     ▒▒▒│
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Modal Detail (Internal Layout)

```
┌───────────────────────────────────────────────────────────┐
│                                                     ✕     │  ← Close (24px)
│                                                           │
│   🏋️  Gym                                                 │  ← Icon (32px)
│   "Consistency over intensity"                            │  ← Micro-copy
│                                                           │
│   🎨  Painting                                            │
│   "Late nights, slow strokes"                             │
│                                                           │
│   📚  Reading                                             │
│   "Always have a book going"                              │
│                                                           │
│   ───────────────── ✦ ─────────────────                   │  ← Divider
│                                                           │
│   Consistency Profile                                     │
│   ┌───────────────────────────────────────────────────┐   │
│   │  ██████████████████████░░░░░░░░░░░░               │   │  ← Band visual
│   │                                                   │   │
│   │  Regular — Shows up most days                     │   │  ← Band label
│   └───────────────────────────────────────────────────┘   │
│                                                           │
│   Availability              ┌─────────────────────────┐   │
│   ┌────────────────────┐    │    Personality Cues     │   │
│   │ ☀️ Mornings        │    ├─────────────────────────┤   │
│   └────────────────────┘    │  ● Introvert   ○ ○ ○ ◐  │   │  ← Tri-state
│                              │  ● Planner    ○ ○ ◐ ○  │   │     indicator
│                              │  ● Deep-talk  ○ ○ ○ ●  │   │
│                              └─────────────────────────┘   │
│                                                           │
│   ───────────────── ✦ ─────────────────                   │  ← Divider
│                                                           │
│   What I'm looking for                                    │
│                                                           │
│   "Someone who wants to grow together, not just           │  ← Intent text
│    pass the time. I value depth over breadth."            │
│                                                           │
│   ───────────────── ✦ ─────────────────                   │  ← Divider
│                                                           │
│   About their approach                      (if available)│  ← Personality
│                                                           │     Summary
│   "They tend to approach new situations with curiosity,   │     (AI-generated)
│    preferring to observe before engaging. In social       │
│    settings, they value meaningful exchanges over..."     │
│                                                           │
│                                                           │
│   ┌───────────────────────────────────────────────────┐   │
│   │                    Connect                        │   │  ← CTA Button
│   └───────────────────────────────────────────────────┘   │   (Disabled if
│                                                           │    slots full)
└───────────────────────────────────────────────────────────┘
```

> [!NOTE]
> **Personality Summary Absence:** If the viewed user has insufficient personality signals, the "About their approach" section is **omitted entirely**. No placeholder, no locked state, no teaser text is shown.

---

## 3. Layout & Spacing Specs

### 3.1 Modal Container

```
CARD DETAIL MODAL CONTAINER
├── Position: fixed, centered (top: 50%, left: 50%, transform: translate(-50%, -50%))
├── Width: min(360px, 90vw)
├── Max-height: 80vh
├── Background: var(--surface-card) → #FFFFFF
├── Border radius: var(--radius-xl) → 20px
├── Border: 1px solid var(--border-subtle)
├── Shadow: 0 24px 48px rgba(0,0,0,0.20)
├── Padding: 24px
├── Overflow-y: auto (if content exceeds max-height)
├── Z-Index: 100
│
└── [SCRIM/BACKDROP]
    ├── Position: fixed, inset: 0
    ├── Background: rgba(0,0,0,0.5)
    ├── Z-Index: 99
    └── Tap: Close modal

Premium Dark Mode (Default):
├── Modal surface: var(--pdm-surface-3) → #222226
├── Modal shadow: Multi-layer with server glow accent
├── Backdrop: Glassmorphism blur (16px) with rgba(13,13,15,0.85)
└── Server accent: Subtle glow on intent border and CTA
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Backdrop** | Glass blur: `backdrop-filter: blur(16px)`, bg `rgba(13,13,15,0.85)` |
| **Modal surface** | Elevated charcoal `#222226` with inner glow highlight |
| **Modal shadow** | Multi-layer: `0 16px 48px rgba(0,0,0,0.4)`  |
| **Hobby icon bg** | Server accent @ 15% with subtle icon glow |
| **Consistency band** | Server accent fill with glow: `0 0 8px rgba(server,0.25)` |
| **Intent statement** | Server-colored left border with ambient glow |
| **Connect CTA** | Server gradient + outer glow (DSD Section 12.5) |
| **Close button** | Subtle hover glow |

**Modal Entry Animation (Premium):**
```css
/* Premium modal enter with scale and backdrop blur */
@keyframes modal-premium-enter {
  0%   { opacity: 0; transform: translate(-50%, -50%) scale(0.95); }
  100% { opacity: 1; transform: translate(-50%, -50%) scale(1.0); }
}
Duration: 250ms
Easing: cubic-bezier(0.16, 1, 0.3, 1)

/* Backdrop blur-in */
@keyframes backdrop-blur-in {
  0%   { backdrop-filter: blur(0); opacity: 0; }
  100% { backdrop-filter: blur(16px); opacity: 1; }
}
```



### 3.2 Internal Spacing

| Element | Token | Value |
|---------|-------|-------|
| Modal padding | `var(--space-6)` | 24px |
| Hobby anchor gap | `var(--space-4)` | 16px |
| Section divider margin | `var(--space-5)` | 20px (top & bottom) |
| Consistency band margin | `var(--space-4)` | 16px |
| Intent text margin | `var(--space-4)` | 16px |
| CTA margin-top | `var(--space-6)` | 24px |
| Close button offset | 16px from top-right |

### 3.3 Z-Index Stack

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Discovery Feed | 1 | Background cards |
| Scrim/Backdrop | 99 | Dimmed overlay |
| Modal Card | 100 | Detail content |
| Close Button | 101 | Above modal content |
| Toast | 200 | Success/error messages |

---

## 4. Component Inventory

### 4.1 Dynamic Server Theming

**All accent colors inherit from the active server:**

| Server | Token | Hex | Applied To |
|--------|-------|-----|------------|
| **Partner** | `var(--server-partner-primary)` | #C9785D | Connect button, icons, band fill |
| **Friend** | `var(--server-friend-primary)` | #4A9B8C | Connect button, icons, band fill |
| **Growth** | `var(--server-growth-primary)` | #6B5B95 | Connect button, icons, band fill |

### 4.2 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Hobby title | Outfit | 600 | 16px | `--unora-ink-primary` |
| Hobby micro-copy | Inter | 400 | 14px | `--unora-ink-secondary` |
| Section label | Inter | 600 | 12px | `--unora-ink-tertiary` |
| Consistency label | Inter | 500 | 14px | `--unora-ink-secondary` |
| Availability tag | Inter | 500 | 13px | `--unora-ink-secondary` |
| Personality cue label | Inter | 500 | 13px | `--unora-ink-secondary` |
| Intent statement | Inter | 400 | 15px | `--unora-ink-primary` |
| Button text | Inter | 600 | 16px | White |

### 4.3 Hobby Anchor Component

**Reference:** [Unora_DesignSystem.md — Section 10.3](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

```
HOBBY ANCHOR
├── Layout: Horizontal (icon left, text right)
├── Gap: 12px
│
├── [ICON]
│   ├── Size: 40px container, 24px icon
│   ├── Background: [server accent] @ 10%
│   ├── Border radius: var(--radius-sm) → 8px
│   └── Icon color: [server accent]
│
└── [TEXT]
    ├── Title: "Gym" — Outfit 16px/600
    └── Micro-copy: "Consistency over intensity" — Inter 14px/400
```

### 4.4 Consistency Band

```
CONSISTENCY BAND
├── Width: 100%
├── Height: 60px
├── Background: var(--surface-background)
├── Border radius: var(--radius-md) → 12px
├── Padding: 12px
│
├── [BAR]
│   ├── Height: 8px
│   ├── Background: var(--border-subtle)
│   ├── Fill: [server accent color]
│   ├── Fill percentage: Based on consistency level
│   │   ├── Calm: 33%
│   │   ├── Regular: 66%
│   │   └── High: 100%
│   └── Border radius: var(--radius-full)
│
└── [LABEL]
    ├── Text: "Regular — Shows up most days"
    └── Style: Inter 14px/500, centered
```

**Consistency Levels:**

| Level | Fill % | Description |
|-------|--------|-------------|
| Calm | 33% | "Shows up occasionally" |
| Regular | 66% | "Shows up most days" |
| High | 100% | "Shows up every day" |

### 4.5 Intent Statement

```
INTENT STATEMENT
├── Background: var(--surface-background)
├── Border radius: var(--radius-md) → 12px
├── Padding: 16px
├── Border-left: 3px solid [server accent]
│
└── [TEXT]
    ├── Style: Inter 15px/400
    ├── Color: var(--unora-ink-primary)
    ├── Line height: 1.6
    └── Max lines: 4 (truncate with "...")
```

### 4.6 Personality Summary Block (Conditional)

> [!IMPORTANT]
> This component is **only rendered if the viewed user has sufficient personality signals**. If unavailable, the entire block is omitted — no placeholder, no locked state.

```
PERSONALITY SUMMARY BLOCK
├── Visibility: Only if personality signals exist for viewed user
├── Background: var(--surface-background)
├── Border radius: var(--radius-md) → 12px
├── Padding: 16px
├── Border-left: 3px solid [server accent]
├── Margin-top: 20px
│
├── [LABEL]
│   ├── Text: "About their approach"
│   ├── Font: Inter 12px / 600
│   ├── Color: var(--unora-ink-tertiary)
│   ├── Text-transform: none
│   └── Margin-bottom: 8px
│
└── [SUMMARY TEXT]
    ├── Style: Inter 14px / 400
    ├── Color: var(--unora-ink-secondary)
    ├── Line height: 1.6
    ├── Max lines: 5 (truncate with "...")
    ├── Generated: Fresh on modal open (stateless, from numeric scores only)
    └── Perspective: Always third-person
```

**Absence Behavior:**
| Condition | Result |
|-----------|--------|
| Signals exist | Show block with AI-generated summary |
| Signals insufficient | **Omit block entirely** |
| Never | Show placeholder, "locked" state, or teaser |

**Copy Principles:**
- ❌ Never use: "Personality type", "Score", "Rating", "Introvert/Extrovert"
- ✓ Use: Behavioral descriptions, context, natural language

### 4.7 Connect Button (Server Themed)

| Property | Value |
|----------|-------|
| Height | 52px |
| Width | 100% |
| Background | [Server accent color] |
| Text | White, Inter 16px/600 |
| Border radius | `var(--radius-md)` → 12px |
| Shadow | 0 4px 12px [server color] @ 20% |

**Button States:**

| State | Appearance |
|-------|------------|
| Default | Full opacity, server color |
| Pressed | Scale 0.98, opacity 0.9 |
| Loading | Spinner replaces text, "Connecting..." |
| Disabled | Opacity 0.4 (if slots full) |

### 4.7 Close Button

| Property | Value |
|----------|-------|
| Position | Top-right, 16px offset |
| Size | 32px touch target, 20px icon |
| Icon | ✕ (close/cross) |
| Color | `var(--unora-ink-tertiary)` |
| Background | Transparent (hover: light gray) |

### 4.8 Availability Tag

```
AVAILABILITY TAG
├── Layout: Inline pill/tag
├── Background: var(--surface-background)
├── Border: 1px solid var(--border-subtle)
├── Border radius: var(--radius-full) → 999px
├── Padding: 6px 12px
├── Height: 28px
│
└── [CONTENT]
    ├── Icon: ☀️ (Mornings) / 🌙 (Evenings) / 📅 (Weekends) — 14px
    ├── Gap: 6px
    └── Text: "Mornings" / "Evenings" / "Weekends" — Inter 13px/500
```

**Availability Options:**

| Value | Icon | Label |
|-------|------|-------|
| Morning | ☀️ | "Mornings" |
| Evening | 🌙 | "Evenings" |
| Weekend | 📅 | "Weekends" |
| Flexible | ⏰ | "Flexible" |

### 4.9 Personality Tri-State Component

```
PERSONALITY TRI-STATE
├── Layout: Vertical stack of cue rows
├── Background: var(--surface-background)
├── Border radius: var(--radius-md) → 12px
├── Padding: 12px
│
└── [CUE ROW] — Repeat per personality trait
    ├── Height: 24px
    ├── Gap: 8px
    │
    ├── [LABEL]
    │   ├── Text: Trait name (e.g., "Introvert", "Planner")
    │   └── Style: Inter 13px/500, var(--unora-ink-secondary)
    │
    └── [INDICATOR]
        ├── 5 circles, 8px each
        ├── Gap: 4px
        ├── Filled: [server accent] — represents position
        ├── Empty: var(--border-subtle)
        └── Visual: Shows spectrum position (1-5)
```

**Personality Cues (PRD Aligned):**

| Trait | Left Extreme (1) | Right Extreme (5) |
|-------|------------------|-------------------|
| Social | Introvert | Extrovert |
| Planning | Spontaneous | Planner |
| Conversation | Small-talk | Deep-talk |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Close button (✕) | Close modal, return to feed |
| Tap | Scrim/backdrop | Close modal, return to feed |
| Tap | Connect button | Execute connect logic |
| Swipe down | Modal | Close modal (optional gesture) |

### 5.2 Entry Transition

```
MODAL ENTRY ANIMATION

Discovery Card Tap
    │
    ▼
┌─────────────────────────────────────────┐
│ 1. Scrim fades in: 0 → 50% black        │
│    Duration: 200ms                       │
│                                          │
│ 2. Modal scales: 0.9 → 1.0              │
│    Modal fades: 0 → 1                   │
│    Duration: 250ms                       │
│    Easing: ease-out                      │
│                                          │
│ 3. Origin: Center of tapped card        │
│    (shared element feel)                │
└─────────────────────────────────────────┘
```

### 5.3 Exit Transition

```
MODAL EXIT ANIMATION

User taps Close/Scrim
    │
    ▼
┌─────────────────────────────────────────┐
│ 1. Modal scales: 1.0 → 0.95             │
│    Modal fades: 1 → 0                   │
│    Duration: 200ms                       │
│                                          │
│ 2. Scrim fades out: 50% → 0             │
│    Duration: 150ms                       │
└─────────────────────────────────────────┘
```

### 5.4 Connect Action Logic

> [!IMPORTANT]
> **Hard Lock Rule:** The Connect button is evaluated and rendered at modal open time. If `Available Slots == 0`, the button renders in the **Disabled** state immediately, preventing any interaction. There is no runtime capacity check or error toast.

```
MODAL OPENS
    │
    ▼
┌─────────────────────────────────────────┐
│ [CHECK] Available Slots > 0?            │
└─────────────────────────────────────────┘
    │                    │
    ▼                    ▼
   YES                   NO
    │                    │
    ▼                    ▼
┌─────────────────────┐  ┌─────────────────────────────────────┐
│ Button: ENABLED     │  │ Button: DISABLED                    │
│ (Default state)     │  │ ├── Opacity: 0.4                    │
│                     │  │ ├── Non-interactive                 │
│                     │  │ └── Cursor: not-allowed             │
└─────────────────────┘  │                                     │
                         │ Helper text below button:           │
                         │ "Complete a streak to unlock"       │
                         └─────────────────────────────────────┘
```

**If Button is ENABLED and USER taps "Connect":**

```
USER taps "Connect" button
    │
    ▼
┌─────────────────────────────────────────┐
│ Button enters LOADING state             │
│ ├── Text: "Connecting..."               │
│ ├── Spinner: 16px white                 │
│ └── Non-interactive                     │
└─────────────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────────────┐
│ API: Send Interest                      │
│                                         │
│ Success:                                │
│ ├── Close modal                         │
│ ├── Haptic: success                     │
│ └── Toast: "Interest Sent! ✓"           │
└─────────────────────────────────────────┘
```

**Capacity Limits by Tier:**
| Tier | Active Slots |
|------|--------------|
| Free | 1 |
| Plus | 2 |
| Pro | 4 |

### 5.5 Mutual Match (If Applicable)

```
IF both users have expressed interest:
    │
    ▼
┌─────────────────────────────────────────┐
│ **MATCH!**                              │
│                                         │
│ Modal shows celebratory animation       │
│ Toast: "It's a match! Streak started."  │
│ Navigate to Active Streaks              │
└─────────────────────────────────────────┘
```

---

## 6. State Definitions

### 6.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Default | Full content visible, Connect enabled | Available slots > 0 |
| Disabled | Connect button dimmed, non-interactive | Available slots == 0 (Hard Lock) |
| Loading | Connect button shows spinner | After tap Connect |
| Success | Modal closing, toast appears | Interest sent |

### 6.2 Default State

```
Modal: Visible, all content loaded
Connect button: Enabled, server-colored
Close button: Visible
Condition: User has at least 1 available connection slot
```

### 6.3 Loading State

```
Connect button: "Connecting..." with spinner
All other elements: Non-interactive (dimmed 0.5)
Close button: Hidden during submit
```

### 6.4 Disabled State (Hard Lock)

> [!CAUTION]
> This state is evaluated **once at modal open**. The button does not become disabled mid-session.

```
Condition: Active Slots == Tier Limit (e.g., 1/1 for Free)

Connect button:
├── Opacity: 0.4
├── Cursor: not-allowed (web) / non-tappable (mobile)
├── Background: [server accent] @ 40%
├── Text: "Connect" (unchanged)
└── No interaction (tap does nothing)

Helper text (below button):
├── Text: "Complete a streak to unlock"
├── Style: Inter 12px/400, var(--unora-ink-tertiary)
└── Alignment: Center

All other modal content: Fully interactive
Close button: Visible and functional
```

### 6.5 Success State

```
Modal: Animating closed
Toast: "Interest Sent! ✓" — 3 second display
Discovery Feed: Visible again
Card: Marked as "Interest Sent" (visual update)
```

---

## 7. Content & Copy Guidelines

### 7.1 Hobby Micro-Copy Examples

| Hobby | Micro-Copy |
|-------|------------|
| Gym | "Consistency over intensity" |
| Painting | "Late nights, slow strokes" |
| Reading | "Always have a book going" |
| Yoga | "Morning person, mat person" |
| Cooking | "Experiments welcome, no recipes" |
| Music | "Guitar in hand, headphones on" |
| Hiking | "Mountains over beaches, always" |

### 7.2 Consistency Labels

| Level | Label |
|-------|-------|
| Calm | "Calm — Shows up occasionally" |
| Regular | "Regular — Shows up most days" |
| High | "High — Shows up every day" |

### 7.3 Intent Statement Examples

- "Someone who wants to grow together, not just pass the time."
- "Looking for depth over breadth. Quality time matters."
- "I want a partner who celebrates the small wins."
- "Seeking a friend who actually follows through."

### 7.4 Button Labels

| State | Label |
|-------|-------|
| Default | "Connect" |
| Loading | "Connecting..." |

### 7.5 Toast Messages

| Type | Message | Duration |
|------|---------|----------|
| Success | "Interest Sent! ✓" | 3s |
| Error (Capacity) | "Connection limit reached. Complete a streak first." | 4s |
| Match | "It's a match! Streak started." | 4s |

---

## 8. Accessibility

### 8.1 Screen Reader
- Modal open: "Card detail. Close button. [Hobby 1], [micro-copy]. [Hobby 2]..."
- Consistency: "Consistency: Regular. Shows up most days."
- Intent: "Looking for: [intent statement]"
- Button: "Connect button" / "Connecting, please wait"

### 8.2 Focus Management
- On open: Focus moves to Close button
- Tab order: Close → Hobbies → Consistency → Intent → Connect
- On close: Focus returns to triggering card

### 8.3 Touch Targets
- Close button: 44px × 44px
- Connect button: 52px × full width

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Centered modal overlay | Critical | ☐ |
| Dimmed backdrop with tap-to-close | Critical | ☐ |
| Scale + fade entry animation | High | ☐ |
| Hobby anchors with micro-copy | Critical | ☐ |
| Consistency band visual | Critical | ☐ |
| Intent statement display | Critical | ☐ |
| Connect button (server themed) | Critical | ☐ |
| Loading state on Connect | High | ☐ |
| Success toast | High | ☐ |
| Capacity error handling | Critical | ☐ |
| Close button | Critical | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 12 — Discovery, Section 13 — Connection Logic |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 10 — Card Variants, Server Tokens |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.C — Discovery, Section 2.D — Connection |
| [Unora_Spec_07_DiscoveryFeed.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_07_DiscoveryFeed.md) | Parent screen |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
