# Unora UI Specification — Personality Setup

**Document ID:** Spec-24  
**Screen Name:** Personality Setup  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Personality Setup** — Optional situational question flow for personality signal collection

### 1.2 User Flow Reference
**Accessible from:** Onboarding (optional step) OR Profile tab entry point

**Sequence:**
```
[Onboarding Step 4.5] OR [Profile → Personality Context]
        ↓
[Personality Setup] → Answer questions → Return to origin
        ↓
       OR
[Exit early] → Partial progress saved → Return to origin
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.A.4.1](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Collect personality signals through situational questions to generate a **Personality Intelligence Summary** visible to others. This is **100% optional** — skipping does not affect any core app functionality.

### 1.4 Core Philosophy

> **"Context, not Categorization."**

The flow must feel:
- **Optional** — Exit always available
- **Calm** — No urgency, no pressure
- **Tap-based** — No free text entry
- **Human** — Situational questions, not clinical assessments

---

## 2. Low-Fidelity Wireframes (ASCII)

### 2.1 Question Screen (Standard)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   ←                                                         │  ← Exit (with confirm)
│                                                             │
│         Question 3 of 10                                    │  ← Progress text
│                                                             │
│   ━━━━━━━━━━━━━━━━━░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░━━   │  ← Progress bar
│                                                             │
│                                                             │
│                        🎨                                   │  ← Visual icon
│                                                             │
│         When starting a creative project,                   │  ← Question text
│         how do you usually begin?                           │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │          Plan everything first                      │   │  ← Option 1
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │          Dive in and figure it out                  │   │  ← Option 2
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │          Depends on the situation                   │   │  ← Option 3
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
│                        Skip question                        │  ← Always visible
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Completion Screen

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                         ✓                                   │  ← Success checkmark
│                                                             │
│               All done!                                     │  ← Headline
│                                                             │
│         Others can now see a bit more                       │  ← Subtext
│         about how you connect.                              │
│                                                             │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                    Continue                         │   │  ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.3 Exit Confirmation Modal

```
┌─────────────────────────────────────────────────────────────┐
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│
│▒▒▒                                                     ▒▒▒│
│▒▒▒   ╭─────────────────────────────────────────────╮   ▒▒▒│
│▒▒▒   │                                             │   ▒▒▒│
│▒▒▒   │         Exit personality setup?             │   ▒▒▒│  ← Title
│▒▒▒   │                                             │   ▒▒▒│
│▒▒▒   │  Your progress is saved. You can            │   ▒▒▒│  ← Body
│▒▒▒   │  continue anytime from your profile.        │   ▒▒▒│
│▒▒▒   │                                             │   ▒▒▒│
│▒▒▒   │  ┌───────────────────────────────────┐      │   ▒▒▒│
│▒▒▒   │  │           Exit                    │      │   ▒▒▒│  ← Secondary
│▒▒▒   │  └───────────────────────────────────┘      │   ▒▒▒│
│▒▒▒   │                                             │   ▒▒▒│
│▒▒▒   │  ┌───────────────────────────────────┐      │   ▒▒▒│
│▒▒▒   │  │           Keep going              │      │   ▒▒▒│  ← Primary
│▒▒▒   │  └───────────────────────────────────┘      │   ▒▒▒│
│▒▒▒   │                                             │   ▒▒▒│
│▒▒▒   ╰─────────────────────────────────────────────╯   ▒▒▒│
│▒▒▒                                                     ▒▒▒│
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
PERSONALITY SETUP SCREEN
├── Position: fixed, 100vw × 100vh
├── Background: var(--pdm-background) → #0D0D0F
├── Display: flex, column
│
├── [HEADER] — 56px height
│   ├── Back/Exit button: left-aligned
│   └── Progress text: right-aligned (optional)
│
├── [PROGRESS BAR] — 4px height
│   └── Fill: var(--accent-gold)
│
├── [CONTENT AREA] — flex: 1
│   ├── Padding: 24px horizontal
│   ├── Align: center
│   └── Gap: 32px
│
└── [ACTION AREA]
    ├── Padding: 20px
    └── Skip link: centered
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Progress bar** | Gold fill with glow: `0 0 8px rgba(196,167,125,0.4)` |
| **Question icon** | Tinted with gold @ 20% bg |
| **Option chips** | Surface `#1A1A1E`, border `#2A2A2E` |
| **Selected chip** | Gold border + gold @ 10% fill |
| **Skip link** | Tertiary text, no decoration |

### 3.3 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen padding | `var(--padding-screen)` | 24px |
| Progress bar height | — | 4px |
| Icon size | — | 48px |
| Question margin-top | `var(--space-6)` | 24px |
| Options gap | `var(--space-3)` | 12px |
| Skip margin-top | `var(--space-5)` | 20px |

---

## 4. Component Inventory

### 4.1 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Progress text | Inter | 400 | 14px | `--pdm-ink-tertiary` |
| Question text | Outfit | 600 | 20px | `--pdm-ink-primary` |
| Option text | Inter | 500 | 16px | `--pdm-ink-primary` |
| Skip link | Inter | 400 | 14px | `--pdm-ink-tertiary` |
| Completion headline | Outfit | 600 | 24px | `--pdm-ink-primary` |
| Completion subtext | Inter | 400 | 16px | `--pdm-ink-secondary` |

### 4.2 Progress Bar

```
PROGRESS BAR
├── Width: 100%
├── Height: 4px
├── Background: var(--border-subtle) → #2A2A2E
├── Fill: var(--accent-gold) → #C4A77D
├── Border radius: var(--radius-full)
├── Glow: 0 0 8px rgba(196,167,125,0.35)
└── Animation: width transition 300ms ease-out
```

### 4.3 Question Icon

```
QUESTION ICON
├── Size: 48px container
├── Icon: Phosphor icon (contextual to question theme)
├── Background: var(--accent-gold) @ 15%
├── Border radius: var(--radius-md) → 12px
└── Color: var(--accent-gold)
```

### 4.4 Option Chip (Answer)

```
OPTION CHIP
├── Width: 100%
├── Min-height: 56px
├── Padding: 16px 20px
├── Background: var(--pdm-surface-2) → #1A1A1E
├── Border: 1.5px solid var(--border-subtle) → #2A2A2E
├── Border radius: var(--radius-lg) → 16px
├── Text-align: center
│
├── [DEFAULT STATE]
│   └── As specified above
│
├── [PRESSED]
│   ├── Scale: 0.98
│   └── Duration: 100ms
│
└── [SELECTED] (transient)
    ├── Border: 2px solid var(--accent-gold)
    ├── Background: var(--accent-gold) @ 10%
    └── Auto-proceed to next question
```

### 4.5 Skip Link

```
SKIP LINK
├── Text: "Skip question"
├── Font: Inter 14px / 400
├── Color: var(--pdm-ink-tertiary)
├── Text-decoration: none
├── Padding: 12px 24px (large touch target)
└── Action: Proceed without recording answer
```

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Option chip | Select + auto-proceed |
| Tap | Skip link | Proceed without answer |
| Tap | ← (Exit) | Show exit confirmation modal |

### 5.2 Question Flow

```
USER enters Personality Setup
    │
    ▼
SYSTEM:
├── Load question 1 of 8-12
├── Show progress bar (fraction)
└── Display options (2-3 per question)
    │
    ▼
USER taps an option
    │
    ▼
┌─────────────────────────────────────────┐
│ IMMEDIATE (0ms):                        │
│ ├── Haptic: Impact (light)              │
│ ├── Chip: Scale 0.98                    │
│ └── Chip: Gold border highlight         │
└─────────────────────────────────────────┘
    │
    ▼ (200ms)
┌─────────────────────────────────────────┐
│ RECORD + PROCEED:                       │
│ ├── Update numeric personality state    │
│ ├── Fade out current question           │
│ ├── Fade in next question               │
│ └── Update progress bar                 │
└─────────────────────────────────────────┘
    │
    ▼ (Repeat until last question)
┌─────────────────────────────────────────┐
│ COMPLETION:                             │
│ ├── Show completion screen              │
│ ├── Success checkmark animation         │
│ └── "Continue" returns to origin        │
└─────────────────────────────────────────┘
```

### 5.3 Skip Behavior

| Action | Result |
|--------|--------|
| Skip question | Proceed to next, no penalty |
| Skip all questions | Allowed (no summary generated) |
| Exit mid-flow | Progress saved (can resume later) |

### 5.4 Exit Flow

```
USER taps ← (back/exit)
    │
    ▼
IF questions answered > 0:
├── Show Exit Confirmation Modal
│   ├── "Exit" → Save progress, return to origin
│   └── "Keep going" → Dismiss modal, continue
│
ELSE (first question):
└── Direct exit, no modal
```

---

## 6. State Definitions

### 6.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Question | Current question displayed | Normal flow |
| Transition | Fading between questions | After answer/skip |
| Completion | Success screen | All questions answered |
| Exit Confirm | Modal overlay | User taps exit mid-flow |

### 6.2 Question State

```
Progress bar: Filled to current/total
Icon: Visible, themed to question
Question: Visible, centered
Options: 2-3 chips, vertically stacked
Skip: Visible, below options
Exit: Visible in header
```

### 6.3 Completion State

```
Progress bar: 100% filled
Icon: Success checkmark (animated)
Headline: "All done!"
Subtext: "Others can now see a bit more about how you connect."
CTA: "Continue" button
```

---

## 7. Content & Copy Guidelines

### 7.1 Progress Text

| Format | Example |
|--------|---------|
| Standard | "Question 3 of 10" |
| Final | "Last question" |

### 7.2 Question Examples

| Theme | Question | Options |
|-------|----------|---------|
| Social | "When meeting new people, you usually..." | "Wait for them to approach" / "Introduce yourself" / "Depends on the setting" |
| Planning | "When starting a creative project..." | "Plan everything first" / "Dive in and figure it out" / "Mix of both" |
| Energy | "After a long week, you prefer to..." | "Recharge alone" / "Meet up with friends" / "Depends on the week" |
| Conflict | "When there's disagreement, you tend to..." | "Seek compromise" / "Stand your ground" / "Avoid the topic" |

### 7.3 Completion Copy

| Element | Copy |
|---------|------|
| Headline | "All done!" |
| Subtext | "Others can now see a bit more about how you connect." |
| CTA | "Continue" |

### 7.4 Exit Modal Copy

| Element | Copy |
|---------|------|
| Title | "Exit personality setup?" |
| Body | "Your progress is saved. You can continue anytime from your profile." |
| Primary CTA | "Keep going" |
| Secondary CTA | "Exit" |

### 7.5 Tone Guidelines

| Principle | Application |
|-----------|-------------|
| Never use | "Test", "Quiz", "Assessment", "Score", "Result" |
| Always use | "Context", "Signals", "Approach" |
| Avoid | Urgency, obligation, pressure |
| Embrace | Calm, optional, curiosity |

---

## 8. Accessibility

### 8.1 Screen Reader

- Progress: "Question 3 of 10"
- Question: "[Question text]"
- Options: "[Option 1], button. [Option 2], button."
- Skip: "Skip question, button"
- Completion: "All done! Others can now see a bit more about how you connect."

### 8.2 Touch Targets

- Option chips: 56px height minimum
- Skip link: 44px × 44px
- Exit button: 44px × 44px

### 8.3 Focus Management

- On load: Focus on first option
- Tab order: Options → Skip → Exit
- On answer: Focus moves to first option of next question

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Question display with icon | Critical | ☐ |
| 2-3 option chips per question | Critical | ☐ |
| Auto-proceed on tap | Critical | ☐ |
| Progress bar animation | High | ☐ |
| Skip functionality | Critical | ☐ |
| Exit confirmation modal | High | ☐ |
| Completion screen | Critical | ☐ |
| Dark mode styling | Medium | ☐ |
| Haptic feedback | Medium | ☐ |
| Resume from partial progress | High | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 18.7 — Personality Intelligence Summary |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 3.5 — Personality Setup UI |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.A.4.1 — Optional Personality Setup |
| [Unora_Spec_04_ProfileCreation.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_04_ProfileCreation.md) | Entry point (onboarding) |
| [Unora_Spec_21_UserProfile.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_21_UserProfile.md) | Entry point (profile tab) |

---

## 11. Critical Constraints (Locked)

> [!CAUTION]
> These constraints are **non-negotiable** and must be enforced at all times:

| Constraint | Enforcement |
|------------|-------------|
| AI NEVER generates questions | All questions are system-defined |
| User NEVER sees their summary | Summary only visible to others |
| Flow is 100% optional | Discovery never blocked |
| No test/quiz language | Copy reviewed for compliance |
| No scores/types shown | Only behavioral descriptions |
| Skip always available | Every question has skip option |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
