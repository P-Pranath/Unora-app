# Unora UI Specification — Daily Check-In Interaction

**Document ID:** Spec-13  
**Screen Name:** Daily Check-In Interaction  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Daily Check-In Interaction** — Focused overlay for capturing daily effort signal

### 1.2 User Flow Reference
**Phase 4 (Streak Loop) → Daily Action** — This is the core interaction where users log their daily activity.

**Sequence:**
```
Streak Detail → [Check-In Interaction] → Success → Hobby Echo
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.D.2](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Capture the daily "Effort Signal" with **minimal friction**. This signal generates the Hobby Echo that the partner sees when they check in.

### 1.4 Core Philosophy

> **"Presence over Performance."**

The interaction must be:
- **Tap-Based Only** — No free text typing allowed
- **Low Friction** — Auto-submit on tap (no confirmation button)
- **Context-Aware** — Questions adapt to user's primary hobby
- **Encouraging** — Celebrate consistency, not intensity

---

## 2. Low-Fidelity Wireframes (ASCII)

### 2.1 Check-In Overlay (Pending State)

```
┌─────────────────────────────────────────────────────────────┐
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│  ← Streak Detail
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│     (dimmed)
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│
│▒▒▒                                                     ▒▒▒│
│▒▒▒   ┌───────────────────────────────────────────┐     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   Day 7 Check-In                    ✕     │     ▒▒▒│  ← Header
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ─────────────────────────────────────   │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   🏋️  What's your workout today?          │     ▒▒▒│  ← Prompt
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │       Last time: Upper Body              │     ▒▒▒│  ← Context
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ┌──────────────┐  ┌──────────────┐      │     ▒▒▒│
│▒▒▒   │   │   Leg Day    │  │  Upper Body  │      │     ▒▒▒│  ← Choice chips
│▒▒▒   │   └──────────────┘  └──────────────┘      │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │   ┌──────────────┐  ┌──────────────┐      │     ▒▒▒│
│▒▒▒   │   │    Cardio    │  │ Rest/Recovery│      │     ▒▒▒│
│▒▒▒   │   └──────────────┘  └──────────────┘      │     ▒▒▒│
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   │       Tap to check in                     │     ▒▒▒│  ← Helper
│▒▒▒   │                                           │     ▒▒▒│
│▒▒▒   └───────────────────────────────────────────┘     ▒▒▒│
│▒▒▒                                                     ▒▒▒│
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Selected State (Before Submit)

```
┌───────────────────────────────────────────┐
│                                           │
│   Day 7 Check-In                    ✕     │
│                                           │
│   🏋️  What's your workout today?          │
│                                           │
│   ┌──────────────┐  ┏━━━━━━━━━━━━━━┓      │
│   │   Leg Day    │  ┃ ✓ Upper Body ┃      │  ← SELECTED
│   └──────────────┘  ┗━━━━━━━━━━━━━━┛      │     Server color bg
│                                           │
│   ┌──────────────┐  ┌──────────────┐      │
│   │    Cardio    │  │ Rest/Recovery│      │
│   └──────────────┘  └──────────────┘      │
│                                           │
└───────────────────────────────────────────┘
```

### 2.3 Success State (After Submit)

```
┌───────────────────────────────────────────┐
│                                           │
│                                           │
│                    ✓                      │  ← Checkmark (animated)
│                                           │
│             Checked in!                   │
│                                           │
│          🏋️  Upper Body                   │  ← Selected option
│                                           │
│       Waiting for partner...              │  ← Next state preview
│                                           │
│                                           │
└───────────────────────────────────────────┘
      ↓ (Auto-closes in 1.5s)
```

### 2.4 Optional Personality Micro-Question (Post Check-In)

> [!NOTE]
> Occasionally (~1 in 3-5 check-ins), after the success state, a single personality-related micro-question may appear. This is **always optional and skippable**.

```
┌───────────────────────────────────────────┐
│                                           │
│                    ✓                      │  ← Success checkmark
│             Checked in!                   │
│                                           │
│   ─────────────────────────────────────   │
│                                           │
│   Quick thought                           │  ← Label (subtle)
│                                           │
│   When you have free time, do you         │  ← Question text
│   usually...                              │
│                                           │
│   ┌──────────────┐  ┌──────────────┐      │
│   │ Seek others  │  │ Solo time    │      │  ← 2 options (chips)
│   └──────────────┘  └──────────────┘      │
│                                           │
│                   Skip                    │  ← Always visible
│                                           │
└───────────────────────────────────────────┘
```

**Trigger Rules:**

| Rule | Value |
|------|-------|
| Frequency | ~1 in 3-5 check-ins (system-controlled) |
| Never on | Day 1-3 of streak |
| Max per session | 1 question only |

**Behavior:**

| Action | Result |
|--------|--------|
| Tap option | Update numeric personality state → Auto-close |
| Tap "Skip" | No penalty → Auto-close |
| Wait 10s | Auto-close (no action recorded) |

**Critical Rules:**
- ❌ AI NEVER generates questions — all questions are system-defined
- ❌ NEVER blocks streak completion (streak already recorded)
- ❌ No progress/score shown to user
- ✓ Skip always available, always consequence-free

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
CHECK-IN OVERLAY
├── Type: Central Card (Modal)
├── Position: Centered, vertically and horizontally
├── Width: min(340px, 90vw)
├── Max-height: 400px
├── Background: var(--surface-card)
├── Border radius: var(--radius-xl) → 20px
├── Shadow: 0 24px 48px rgba(0,0,0,0.20)
├── Padding: 24px
├── Z-Index: 100
│
└── [SCRIM/BACKDROP]
    ├── Background: rgba(0,0,0,0.5)
    ├── Z-Index: 99
    └── Tap: Close overlay (cancel check-in)

Premium Dark Mode (Default):
├── Modal surface: var(--pdm-surface-3) → #222226
├── Backdrop: Glassmorphism blur (16px) with 85% opacity
├── Selected chip: Server color with outer glow
└── Success checkmark: Teal with expanding glow burst
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Backdrop** | Glass blur: `blur(16px)`, bg `rgba(13,13,15,0.85)` |
| **Modal surface** | Elevated charcoal `#222226` with inner glow highlight |
| **Chips (default)** | Border `#2A2A2E`, transparent bg |
| **Chips (hover)** | Border server color, server @ 8% bg |
| **Chips (selected)** | Server bg + glow: `0 0 12px rgba(server,0.35)` |
| **Success checkmark** | Teal with glow burst animation |
| **Context text** | Gold-tinted subtle text |

**Chip Selection Glow:**
```css
/* Selected chip with server glow */
.chip.selected {
  background: var(--server-color);
  box-shadow: 0 0 12px rgba(var(--server-color-rgb), 0.4);
  transition: all 150ms ease-out;
}

/* Success state glow burst */
.success-checkmark {
  filter: drop-shadow(0 0 24px rgba(74, 155, 140, 0.5));
}
```



### 3.2 Chip Grid

```
CHOICE CHIP GRID
├── Display: Grid
├── Columns: 2
├── Gap: var(--space-3) → 12px
├── Chip min-height: 52px
├── Chip min-width: 140px
└── Touch target: Full chip area
```

### 3.3 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Card padding | `var(--space-6)` | 24px |
| Header margin-bottom | `var(--space-4)` | 16px |
| Prompt margin-bottom | `var(--space-2)` | 8px |
| Context margin-bottom | `var(--space-5)` | 20px |
| Chip gap | `var(--space-3)` | 12px |
| Helper margin-top | `var(--space-4)` | 16px |

### 3.4 Z-Index Stack

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Streak Detail | 1 | Background |
| Scrim | 99 | Dimmed overlay |
| Check-In Card | 100 | Modal content |
| Success Animation | 101 | Checkmark |

---

## 4. Component Inventory

### 4.1 Dynamic Server Theming

**Selected chip color matches active server:**

| Server | Token | Hex | Usage |
|--------|-------|-----|-------|
| **Partner** | `var(--server-partner-primary)` | #C9785D (Terracotta) | Selected chip bg |
| **Friend** | `var(--server-friend-primary)` | #4A9B8C (Teal) | Selected chip bg |
| **Growth** | `var(--server-growth-primary)` | #6B5B95 (Indigo) | Selected chip bg |

### 4.2 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Header "Day X Check-In" | Outfit | 600 | 18px | `--unora-ink-primary` |
| Prompt question | Outfit | 600 | 20px | `--unora-ink-primary` |
| Context "Last time" | Inter | 400 | 14px | `--unora-ink-tertiary` |
| Chip text | Inter | 500 | 16px | `--unora-ink-primary` / White |
| Helper text | Inter | 400 | 12px | `--unora-ink-tertiary` |
| Success text | Outfit | 600 | 24px | `--feedback-success` |

### 4.3 Choice Chip Component

**Reference:** [Unora_DesignSystem.md — Section 3.4](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

```
CHOICE CHIP (Check-In)
├── Min-height: 52px
├── Min-width: 140px
├── Padding: 14px 20px
├── Border radius: var(--radius-lg) → 16px
├── Font: Inter 16px / 500
├── Text-align: center
│
├── [DEFAULT STATE]
│   ├── Background: transparent
│   ├── Border: 1.5px solid var(--border-medium)
│   └── Text: var(--unora-ink-primary)
│
├── [HOVER/FOCUS]
│   ├── Border: 1.5px solid [server color]
│   └── Background: [server color] @ 5%
│
├── [PRESSED]
│   ├── Scale: 0.95
│   └── Duration: 100ms
│
├── [SELECTED]
│   ├── Background: [server color]
│   ├── Border: none
│   ├── Text: white
│   ├── Icon: ✓ checkmark (16px, left of text)
│   └── Scale: 1.0 (bounce back from 0.95)
│
└── [SUBMITTING]
    ├── Background: [server color]
    ├── Opacity: 0.8
    └── Spinner: White, 16px (optional)
```

### 4.4 Success Checkmark

```
SUCCESS CHECKMARK
├── Size: 64px
├── Color: var(--feedback-success) → #4A9B8C
├── Animation: Scale 0 → 1.2 → 1.0 (bounce)
├── Duration: 400ms
├── Timing: ease-out
└── Stroke animation: Draw-in effect (optional)
```

### 4.5 Close Button

| Property | Value |
|----------|-------|
| Position | Top-right, 16px offset |
| Size | 32px touch, 20px icon |
| Icon | ✕ |
| Color | `var(--unora-ink-tertiary)` |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Choice chip | Select + auto-submit |
| Tap | Close (✕) | Cancel, return to Streak Detail |
| Tap | Scrim | Cancel, return to Streak Detail |

### 5.2 Check-In Flow

```
USER opens Check-In Overlay
    │
    ▼
┌─────────────────────────────────────────┐
│ SYSTEM:                                 │
│ ├── Retrieve user's primary hobby       │
│ ├── Generate context-aware question     │
│ ├── Display 3-4 relevant options        │
│ └── Show "Last time: [activity]"        │
└─────────────────────────────────────────┘
    │
    ▼
USER taps a choice chip
    │
    ▼
┌─────────────────────────────────────────┐
│ IMMEDIATE (0ms):                        │
│ ├── Haptic: Success (medium impact)     │
│ ├── Chip: Scale 0.95                    │
│ └── Other chips: Fade to 0.5 opacity    │
└─────────────────────────────────────────┘
    │
    ▼ (100ms)
┌─────────────────────────────────────────┐
│ VISUAL FEEDBACK:                        │
│ ├── Chip: Scale back to 1.0             │
│ ├── Chip: Background → [server color]   │
│ ├── Chip: Text → white + ✓ checkmark    │
│ └── Submit to API                       │
└─────────────────────────────────────────┘
    │
    ▼ (300ms)
┌─────────────────────────────────────────┐
│ TRANSITION TO SUCCESS:                  │
│ ├── Card content fades out              │
│ ├── Checkmark animates in (bounce)      │
│ ├── "Checked in!" text                  │
│ └── Selected option displayed           │
└─────────────────────────────────────────┘
    │
    ▼ (1500ms)
┌─────────────────────────────────────────┐
│ AUTO-CLOSE:                             │
│ ├── Card scales down + fades            │
│ ├── Return to Streak Detail             │
│ └── Show State B (Waiting) or C (Echo)  │
└─────────────────────────────────────────┘
```

### 5.3 Haptic Feedback

**Reference:** [Unora_DesignSystem.md — Section 5.2](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

| Event | Haptic Type | Intensity |
|-------|-------------|-----------|
| Chip tap | Impact | Medium |
| Submit success | Notification | Success |
| Card close | Impact | Light |

### 5.4 Auto-Submit Behavior

> **No confirmation button required.** Single tap = immediate submission.

This design choice prioritizes:
- **Low friction** — One tap to complete
- **Speed** — No extra steps
- **Delight** — Immediate feedback

### 5.5 Transitions

| Transition | Animation |
|------------|-----------|
| Overlay open | Scale 0.9 → 1.0 + fade, 250ms |
| Chip press | Scale 0.95, 100ms |
| Chip select | Color fill, 150ms |
| Success checkmark | Scale 0 → 1.2 → 1.0, 400ms |
| Overlay close | Scale 1.0 → 0.95 + fade, 200ms |

---

## 6. State Definitions

### 6.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Pending | All chips default, question visible | Initial |
| Selected | One chip highlighted | User tapped |
| Submitting | Selected chip, brief delay | API call |
| Success | Checkmark + confirmation | Submission complete |

### 6.2 Pending State

```
Question: Visible
Options: All chips in default style
Helper: "Tap to check in"
Close button: Visible
```

### 6.3 Selected State (Transient)

```
Selected chip: Server color bg, white text, ✓
Other chips: Faded (opacity 0.5)
Duration: ~100ms before transitioning
```

### 6.4 Submitting State

```
Selected chip: Slightly dimmed (0.8 opacity)
Spinner: Optional (only if API slow)
Duration: 100-500ms
```

### 6.5 Success State

```
Content: Replaced with success animation
Checkmark: Large, animated, green
Text: "Checked in!"
Sub-text: Selected option name
Auto-close: After 1.5 seconds
```

---

## 7. Content & Copy Guidelines

### 7.1 Header Format

| Format | Example |
|--------|---------|
| Standard | "Day 7 Check-In" |
| Milestone | "Day 5 Check-In ✨" |

### 7.2 Context-Aware Prompts

| Hobby Category | Prompt |
|----------------|--------|
| Gym | "What's your workout today?" |
| Running | "How far are you going today?" |
| Reading | "What are you reading today?" |
| Meditation | "What's your practice today?" |
| Creative | "What are you creating today?" |
| Learning | "What are you studying today?" |

### 7.3 Context Line

| Format | Example |
|--------|---------|
| With history | "Last time: Upper Body" |
| First check-in | "Your first check-in!" |
| Streak milestone | "Day 5 — Keep going!" |

### 7.4 Choice Options by Hobby

**Gym:**
- Leg Day, Upper Body, Cardio, Rest/Recovery

**Running:**
- Short Run, Long Run, Intervals, Rest Day

**Reading:**
- Fiction, Non-Fiction, Study, Audiobook

**Meditation:**
- Morning Sit, Breathwork, Body Scan, Guided

**Creative (Art):**
- Sketching, Painting, Digital, Practice

**Learning:**
- Course, Project, Practice, Research

### 7.5 Helper Text

| State | Text |
|-------|------|
| Pending | "Tap to check in" |
| Success | "Waiting for partner..." |

### 7.6 Success Messages

| Type | Message |
|------|---------|
| Standard | "Checked in!" |
| Milestone | "Day 5 unlocked! 🎉" |
| Streak extended | "Streak Extended!" |

---

## 8. Accessibility

### 8.1 Screen Reader
- Header: "Day 7 Check-In. Close button."
- Prompt: "What's your workout today? Last time: Upper Body."
- Chips: "Leg Day, button. Upper Body, button. Cardio, button. Rest Recovery, button."
- Selected: "Upper Body, selected."
- Success: "Checked in! Upper Body."

### 8.2 Touch Targets
- Chips: 52px height minimum
- Close button: 44px × 44px

### 8.3 Focus Management
- On open: Focus trap within modal
- Tab order: Close → Chips (L→R, T→B)
- On close: Focus returns to trigger

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Choice chip grid (2 columns) | Critical | ☐ |
| Server-colored selected state | Critical | ☐ |
| Auto-submit on tap | Critical | ☐ |
| Haptic feedback | High | ☐ |
| Success animation (checkmark) | High | ☐ |
| Auto-close after success | High | ☐ |
| Context-aware prompts | Medium | ☐ |
| Hobby-specific options | Critical | ☐ |
| Close button | High | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 14 — Streak System |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 3.4 — Choice Chips, Section 5.2 — Haptics |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.D.2 — Check-In Logic |
| [Unora_Spec_12_StreakDetail.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_12_StreakDetail.md) | Parent screen |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
