# Unora UI Specification — Streak Detail View

**Document ID:** Spec-12  
**Screen Name:** Streak Detail View  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Streak Detail View** — Single connection detail and daily check-in screen

### 1.2 User Flow Reference
**Phase 4 (The Streak Loop)** — This is where the daily work happens — checking in, viewing progress, and maintaining the connection.

**Sequence:**
```
Streaks Dashboard → [Streak Detail View] → Check-In → Hobby Echo
                            ↓
                    Reveal Roadmap → Reveal Modal
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.D](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Build trust through daily consistency. Users check in, see partner activity (Hobby Echo), track progress toward reveals, and maintain streaks through nudges or recovery.

### 1.4 The Core Loop

| Action | Trigger | Outcome |
|--------|---------|---------|
| **Check In** | Tap daily prompt | Extends streak |
| **See Progress** | View Reveal Roadmap | Track path to Day 15 |
| **Maintain** | Nudge / Recover | Save at-risk streaks |
| **Hobby Echo** | Both users check in | See partner's focus |

---

## 2. Low-Fidelity Wireframes (ASCII)

### 2.1 State A: Action Needed (Check-In Required)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   ←  A****  🔥                                    ⚙️  ⋮    │  ← Header
│                                                             │
│                                                             │
│                        Day 7                                │  ← Hero counter
│                      ───────                                │
│                   of 15 days                                │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │       What's your focus today?                      │   │  ← Prompt
│   │                                                     │   │
│   │       Last time: Ran 5k                             │   │  ← Context
│   │                                                     │   │
│   │   ┌──────────┐  ┌──────────┐  ┌──────────┐          │   │
│   │   │   Rest   │  │  Cardio  │  │ Strength │          │   │  ← Choice chips
│   │   └──────────┘  └──────────┘  └──────────┘          │   │
│   │                                                     │   │
│   │   ┌──────────┐  ┌────────────┐                      │   │
│   │   │  Yoga    │  │  Stretching │                     │   │
│   │   └──────────┘  └────────────┘                      │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ─────────────────────────────────────────────────────     │
│                                                             │
│   CONNECTION CORE                                           │
│                                                             │
│                        ┌───────────┐                          │
│                        │  ✦✦✦✦✦✦✦  │  ← Trust Orb           │
│                        │  ✦✦✦✦✦✦✦  │    (glowing sphere,    │
│                        │  ✦✦✦✦✦✦✦  │     intensity grows)   │
│                        └───────────┘                          │
│                                                             │
│                "Trust building..."                          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 State B: Waiting (User Checked In)

```
┌─────────────────────────────────────────────────────────────┐
│   ←  A****  🔥                                    ⚙️  ⋮    │
│                                                             │
│                        Day 7                                │
│                      ───────                                │
│                   of 15 days                                │
│                                                             │
│   ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐   │
│   ╎                                                     ╎   │  ← Dashed border
│   ╎       ✓ You checked in                              ╎   │     (waiting)
│   ╎                                                     ╎   │
│   ╎       Waiting for partner...                        ╎   │
│   ╎                                                     ╎   │
│   ╎       ○ ○ ○ (pulsing animation)                     ╎   │
│   ╎                                                     ╎   │
│   ╎       They have until midnight to check in.         ╎   │
│   ╎                                                     ╎   │
│   └ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┘   │
│                                                             │
│   CONNECTION CORE                                           │
│   ✦ Trust building... (subtle pulse)                        │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.3 State C: Success (Both Checked In — Hobby Echo)

```
┌─────────────────────────────────────────────────────────────┐
│   ←  A****  🔥                                    ⚙️  ⋮    │
│                                                             │
│                        Day 7                                │
│                      ───────                                │
│                    Streak Extended! ✓                       │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓  │   │
│   │  ┃                                               ┃  │   │  ← Hobby Echo
│   │  ┃   Your partner is focusing on:                ┃  │   │     (highlighted)
│   │  ┃                                               ┃  │   │
│   │  ┃          🏋️  Deep Work                       ┃  │   │
│   │  ┃                                               ┃  │   │
│   │  ┃   "Crushing the morning routine"              ┃  │   │
│   │  ┃                                               ┃  │   │
│   │  ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛  │   │
│   │                                                     │   │
│   │   You checked in: Cardio ✓                          │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   CONNECTION CORE                                           │
│   ✦✦✦ Connection deepening... (bright glow)                  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.4 State D: At Risk (Partner Missed)

```
┌─────────────────────────────────────────────────────────────┐
│   ←  A****  🔥                                    ⚙️  ⋮    │
│                                                             │
│                        Day 7                                │
│                      ───────                                │
│                    ⚠️ At Risk                               │  ← Amber theme
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │       Partner missed yesterday                      │   │
│   │                                                     │   │
│   │       Send them a nudge to remind them.             │   │
│   │                                                     │   │
│   │       ┌─────────────────────────────────────┐       │   │
│   │       │           Send Nudge                │       │   │  ← Amber button
│   │       └─────────────────────────────────────┘       │   │
│   │                                                     │   │
│   │       Nudges left today: 2                          │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.5 State E: Payment Required (User Missed)

```
┌─────────────────────────────────────────────────────────────┐
│   ←  A****  🔥                                    ⚙️  ⋮    │
│                                                             │
│                        Day 7                                │
│                      ───────                                │
│                    🚨 Streak at Risk                        │  ← Terracotta theme
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │       You missed your check-in yesterday.           │   │
│   │                                                     │   │
│   │       Recover your streak before it breaks.         │   │
│   │                                                     │   │
│   │       Time remaining: 4 hours 32 minutes            │   │
│   │                                                     │   │
│   │       ┌─────────────────────────────────────┐       │   │
│   │       │        Recover Streak — ₹49         │       │   │  ← Terracotta
│   │       └─────────────────────────────────────┘       │   │
│   │                                                     │   │
│   │       or                                            │   │
│   │                                                     │   │
│   │       Let streak end →                              │   │  ← Destructive
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
STREAK DETAIL CONTAINER
├── Position: fixed, 100vw × 100vh
├── Background: var(--surface-background)
├── Display: flex, column
├── Overflow-y: scroll
│
├── [HEADER] — 60px
│   ├── Padding: 20px horizontal
│   ├── Left: Back arrow + Partner alias + Server icon
│   └── Right: Settings + Overflow menu
│
├── [HERO SECTION] — 120px
│   ├── Align: center
│   ├── Day counter: Large number
│   └── Status subtitle
│
├── [INTERACTION CARD] — dynamic
│   ├── Margin: 0 20px
│   ├── Min-height: 200px
│   └── Contains: Prompt/Status/Hobby Echo
│
├── [REVEAL ROADMAP] — 100px
│   ├── Margin: 24px 20px
│   └── Horizontal timeline
│
└── [SAFE AREA PADDING]

Premium Dark Mode (Default):
├── Background: var(--pdm-background) → #0D0D0F
├── Day counter: Server color with glow on milestones
├── Interaction card: Elevated surface with state-specific glows
└── Reveal roadmap: Glowing milestone nodes
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Day counter** | Server color, milestone days have glow: `0 0 24px rgba(server,0.35)` |
| **Interaction card** | Surface `#1A1A1E`, border `#2A2A2E`, inner glow highlight |
| **Choice chips (default)** | Border `#2A2A2E`, surface transparent |
| **Choice chips (selected)** | Server bg + glow: `0 0 12px rgba(server,0.35)` |
| **Hobby Echo card** | Server accent @ 15% bg, server border with subtle glow |
| **Roadmap completed nodes** | Server color with subtle glow |
| **Roadmap milestone nodes** | Larger, server color, pulsing glow on current/next |

**Milestone Day Glow (Premium):**
```css
/* Day counter at milestone (5, 8, 12, 15) */
.day-counter.milestone {
  color: var(--server-color);
  text-shadow: 0 0 24px rgba(var(--server-color-rgb), 0.4);
}

/* Current milestone node on roadmap */
.roadmap-node.milestone.current {
  background: var(--server-color);
  box-shadow: 0 0 16px rgba(var(--server-color-rgb), 0.5);
  animation: milestone-pulse 2s infinite;
}
```



### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen padding | `var(--padding-screen)` | 20px |
| Hero counter size | — | 64px text |
| Card padding | `var(--space-5)` | 20px |
| Section gap | `var(--space-6)` | 24px |
| Chip gap | `var(--space-2)` | 8px |
| Roadmap node size | — | 12px |

---

## 4. Component Inventory

### 4.1 Dynamic Server Theming

**The entire screen adapts to the connection's server:**

| Server | Token | Applied To |
|--------|-------|------------|
| **Partner** | `var(--server-partner-primary)` | Headers, buttons, progress, hero accent |
| **Friend** | `var(--server-friend-primary)` | Headers, buttons, progress, hero accent |
| **Growth** | `var(--server-growth-primary)` | Headers, buttons, progress, hero accent |

### 4.2 Status Colors (Override)

| State | Color | Usage |
|-------|-------|-------|
| At Risk | `var(--feedback-warning)` #E6A43A | Amber theme overrides server color |
| Payment | `var(--feedback-error)` #C9785D | Terracotta overrides server color |
| Success | `var(--feedback-success)` #4A9B8C | Green for "Streak Extended" |

### 4.3 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Partner alias | Outfit | 600 | 18px | `--unora-ink-primary` |
| Day counter | Outfit | 700 | 64px | [Server color] |
| "of 15 days" | Inter | 400 | 16px | `--unora-ink-tertiary` |
| Status subtitle | Inter | 500 | 14px | [Status color] |
| Prompt headline | Outfit | 600 | 20px | `--unora-ink-primary` |
| Context text | Inter | 400 | 14px | `--unora-ink-secondary` |
| Hobby Echo title | Outfit | 600 | 18px | `--unora-ink-primary` |
| Roadmap labels | Inter | 500 | 10px | `--unora-ink-tertiary` |

### 4.4 Check-In Choice Chips

**Reference:** [Unora_DesignSystem.md — Section 3.4](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

```
CHOICE CHIP (Single Select)
├── Height: 44px
├── Padding: 12px 20px
├── Border radius: var(--radius-lg) → 16px
├── Font: Inter 14px / 500
│
├── [DEFAULT]
│   ├── Background: transparent
│   ├── Border: 1.5px solid var(--border-medium)
│   └── Text: var(--unora-ink-primary)
│
├── [SELECTED]
│   ├── Background: [server color]
│   ├── Border: none
│   ├── Text: white
│   └── Icon: ✓ checkmark
│
└── [PRESSED]
    └── Scale: 0.97, 100ms
```

### 4.5 Interaction Card States

#### Action Needed Card
```
INTERACTION CARD (Check-In)
├── Background: var(--surface-card)
├── Border: 1px solid var(--border-subtle)
├── Border radius: var(--radius-xl) → 20px
├── Padding: 24px
├── Shadow: 0 4px 12px rgba(0,0,0,0.06)
│
├── [PROMPT]
│   └── "What's your focus today?"
│
├── [CONTEXT]
│   └── "Last time: Ran 5k"
│
└── [CHOICE CHIPS]
    └── Grid of options
```

#### Waiting Card
```
INTERACTION CARD (Waiting)
├── Background: var(--surface-card)
├── Border: 2px dashed var(--border-medium) ← Dashed
├── Border radius: var(--radius-xl)
├── Padding: 24px
│
├── [STATUS]
│   ├── "✓ You checked in"
│   └── "Cardio"
│
└── [WAITING]
    ├── "Waiting for partner..."
    └── Pulsing dots animation
```

### 4.6 Hobby Echo Component

```
HOBBY ECHO CARD
├── Background: [server color] @ 10%
├── Border: 2px solid [server color]
├── Border radius: var(--radius-lg)
├── Padding: 20px
│
├── [HEADER]
│   └── "Your partner is focusing on:"
│
├── [CONTENT]
│   ├── Icon: Hobby emoji (32px)
│   └── Activity: "Deep Work"
│
└── [QUOTE] (optional)
    └── "Crushing the morning routine"
```

### 4.7 Connection Core (Trust Orb) — Mystery Reveal UX

> [!IMPORTANT]
> **Users never see day countdowns, milestone markers, or reveal roadmaps.** The Connection Core is an abstract visual that represents streak strength without revealing timing.

```
CONNECTION CORE (Abstract Progress Visual)
├── Width: 100% - padding
├── Height: 80px (centered)
│
├── [ORB CONTAINER]
│   ├── Size: 64px diameter
│   ├── Center aligned
│   └── Glassmorphic container with subtle blur
│
├── [TRUST ORB]
│   ├── Shape: Circular gradient sphere
│   ├── Base color: var(--glass-fill) with [server color] tint
│   ├── Inner glow: var(--glow-color-server) with increasing intensity
│   ├── Intensity: Grows with streak day count (internal logic)
│   │   ├── Day 1-5:   Dim, soft pulse
│   │   ├── Day 6-10:  Medium brightness, steady glow
│   │   ├── Day 11-14: Bright, intensifying pulse
│   │   └── Day 15:    Maximum brilliance, radiant shimmer
│   └── Animation: Continuous subtle breathing pulse (1.0 → 1.05 scale)
│
├── [STATUS TEXT]
│   ├── Font: Inter 14px / 500
│   ├── Color: var(--text-secondary)
│   └── Copy: Abstract, intriguing (never mentions days)
│       ├── "Trust building..."
│       ├── "Connection deepening..."
│       ├── "Something is forming..."
│       └── "The bond strengthens..."
│
└── [PREMIUM VISUAL NOTES]
    ├── Bioluminescent aesthetic: soft organic glow, NOT harsh neon
    ├── Ethereal quality: dreamy transparency layers
    └── Dark mode: Orb is the brightest element on screen
```

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Back arrow | Return to Streaks Dashboard |
| Tap | Choice chip | Select check-in option |
| Tap | Selected chip | Submit check-in |
| Tap | Nudge button | Send nudge notification |
| Tap | Recover button | Open payment flow |
| Tap | Milestone node | Open Reveal Modal (if unlocked) |
| Tap | Settings | Open connection settings |

### 5.2 Check-In Flow

```
USER opens Streak Detail (State A)
    │
    ▼
┌─────────────────────────────────────────┐
│ Show prompt: "What's your focus today?" │
│ Show context: "Last time: [activity]"   │
│ Show choice chips                       │
└─────────────────────────────────────────┘
    │
    ▼
USER taps a chip
    │
    ▼
┌─────────────────────────────────────────┐
│ Chip: Selected state (server color)     │
│ Other chips: Fade slightly              │
│ Haptic: Light impact                    │
└─────────────────────────────────────────┘
    │
    ▼
USER taps selected chip again (confirm)
    │
    ▼
┌─────────────────────────────────────────┐
│ API: Submit check-in                    │
│ Haptic: Success                         │
│ Animation: Card flips/transitions       │
│ State → B (Waiting) OR C (Success)      │
└─────────────────────────────────────────┘
```

### 5.3 Hobby Echo Logic

```
BOTH users have checked in for the day
    │
    ▼
┌─────────────────────────────────────────┐
│ Transform Interaction Card to:          │
│ ├── Status: "Streak Extended! ✓"        │
│ ├── Hobby Echo: Partner's activity      │
│ ├── User's check-in: Smaller display    │
│ └── Celebration: Confetti (optional)    │
└─────────────────────────────────────────┘
```

### 5.4 Reveal Unlock

```
USER reaches milestone day (5, 8, 12, 15)
    │
    ▼
┌─────────────────────────────────────────┐
│ Roadmap node: Glows/pulses              │
│ Toast: "New reveal unlocked!"           │
│ Node becomes tappable                   │
└─────────────────────────────────────────┘
    │
    ▼
USER taps milestone node
    │
    ▼
┌─────────────────────────────────────────┐
│ Open Reveal Modal (Screen #13)          │
│ Show unlocked content                   │
└─────────────────────────────────────────┘
```

### 5.5 Transitions

| Transition | Animation |
|------------|-----------|
| Card state change | Cross-fade, 300ms |
| Chip selection | Scale + color, 150ms |
| Check-in submit | Card flip, 400ms |
| Hobby Echo appear | Slide up + fade, 300ms |
| Roadmap progress | Fill animation, 500ms |

---

## 6. State Definitions

### 6.1 State Matrix

| State | Condition | UI Theme | Primary Action |
|-------|-----------|----------|----------------|
| **A: Action Needed** | User hasn't checked in | Server color | Check-in chips |
| **B: Waiting** | User checked, partner hasn't | Server color | Wait message |
| **C: Success** | Both checked in | Green/Server | Hobby Echo |
| **D: At Risk** | Partner missed yesterday | Amber | Nudge button |
| **E: Payment** | User missed yesterday | Terracotta | Recover button |

### 6.2 State A: Action Needed

```
Hero subtitle: "of 15 days"
Card: Prompt + choice chips
Button: None (chips are the action)
Roadmap: Normal progress
```

### 6.3 State B: Waiting

```
Hero subtitle: "of 15 days"
Card: "✓ You checked in" + waiting message
Border: Dashed
Animation: Pulsing dots
Roadmap: Current day highlighted
```

### 6.4 State C: Success

```
Hero subtitle: "Streak Extended! ✓"
Hero color: Green success
Card: Hobby Echo + user's check-in
Roadmap: Progress fills to current day
Celebration: Optional confetti
```

### 6.5 State D: At Risk

```
Hero subtitle: "⚠️ At Risk"
Hero color: Amber
Card: "Partner missed yesterday" + Nudge button
Button: "Send Nudge" (amber)
Note: "Nudges left today: X"
```

### 6.6 State E: Payment Required

```
Hero subtitle: "🚨 Streak at Risk"
Hero color: Terracotta
Card: "You missed yesterday" + countdown + Recover button
Button: "Recover Streak — ₹49" (terracotta)
Secondary: "Let streak end →"
Timer: "Time remaining: X hours"
```

---

## 7. Content & Copy Guidelines

### 7.1 Check-In Prompts

| Hobby Category | Prompt |
|----------------|--------|
| Fitness | "What's your workout today?" |
| Creative | "What are you creating today?" |
| Learning | "What are you learning today?" |
| Wellness | "What's your wellness focus?" |
| General | "What's your focus today?" |

### 7.2 Context Text

| Format | Example |
|--------|---------|
| Last activity | "Last time: Ran 5k" |
| Streak context | "7 days of consistency!" |
| Encouragement | "Keep the momentum going" |

### 7.3 Choice Chip Options

| Category | Options |
|----------|---------|
| Fitness | Rest, Cardio, Strength, Yoga, Stretching |
| Creative | Writing, Design, Music, Art, Photography |
| Learning | Reading, Course, Practice, Research |
| Wellness | Meditation, Sleep, Nutrition, Mindfulness |

### 7.4 Hobby Echo Messages

| Format | Example |
|--------|---------|
| Activity | "Your partner is focusing on: Deep Work" |
| With quote | "Crushing the morning routine" |
| Encouragement | "You're both showing up!" |

### 7.5 Connection Core Status Copy

> [!IMPORTANT]
> **Mystery Reveal UX:** Copy is abstract and intriguing, never mentions specific days or milestones.

| Intensity | Status Text |
|-----------|--------------|
| Low (Day 1-5) | "Trust building..." |
| Medium (Day 6-10) | "Connection deepening..." |
| High (Day 11-14) | "Something is forming..." |
| Maximum (Day 15) | "The bond is ready." |

### 7.6 State-Specific Copy

| State | Message |
|-------|---------|
| Waiting | "Waiting for partner..." |
| Success | "Streak Extended!" |
| At Risk | "Partner missed yesterday. Send them a nudge." |
| Payment | "You missed your check-in. Recover before it breaks." |

---

## 8. Accessibility

### 8.1 Screen Reader
- Header: "Streak with A, 4 letters hidden. Partner server. Day 7. Trust building."
- Check-in: "What's your focus today? 5 options available. Rest, button."
- Waiting: "You checked in with Cardio. Waiting for partner."
- Hobby Echo: "Streak extended. Partner is focusing on Deep Work."

### 8.2 Touch Targets
- Choice chips: 44px height
- Action buttons: 52px height
- Connection Core: Decorative element (no touch action)

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| 5 distinct states (A-E) | Critical | ☐ |
| Check-in choice chips | Critical | ☐ |
| Hobby Echo display | Critical | ☐ |
| Connection Core (Trust Orb) | High | ☐ |
| State-based theming (Amber/Terracotta) | High | ☐ |
| Server-colored theming | Critical | ☐ |
| Nudge action | High | ☐ |
| Recovery payment flow | Critical | ☐ |
| Card state transitions | Medium | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 14 — Streak System, Section 15 — Reveals |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 3.4 — Choice Chips, Section 6 — Colors |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.D — Streak Loop |
| [Unora_Spec_11_StreaksDashboard.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_11_StreaksDashboard.md) | Parent screen |
| Unora_Spec_13_RevealModal.md (planned) | Milestone reveal |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
</Parameter>
<parameter name="Complexity">7
