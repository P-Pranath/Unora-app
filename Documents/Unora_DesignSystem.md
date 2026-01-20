# Unora — Design System Document (DSD)

**Version:** 1.2  
**Date:** January 2026  
**Classification:** Internal / Design & Engineering  
**Author:** Design Leadership

---

## Table of Contents

1. [Design Philosophy & Principles](#1-design-philosophy--principles)
2. [Foundations](#2-foundations)
3. [Component Library](#3-component-library)
4. [Visualizing Logic & States](#4-visualizing-logic--states)
   - 4.X [Verification UI](#4x-verification-ui)
5. [Feedback & Haptics](#5-feedback--haptics)
6. [Accessibility & Inclusivity](#6-accessibility--inclusivity)

---

## 1. Design Philosophy & Principles

### 1.1 Core Design Pillars

Unora is the **anti-Tinder**. Every design decision must reinforce intentionality over impulsivity, depth over speed, and earned trust over instant access.

| Pillar | Definition | Anti-Pattern |
|--------|------------|--------------|
| **Clarity Over Clutter** | Every screen has one clear action. Minimal UI, maximum focus. | Busy dashboards, multiple CTAs competing for attention |
| **Anticipation Over Gratification** | Delayed reveals create emotional investment. The wait is the feature. | Instant previews, "unlock everything now" prompts |
| **Presence Over Performance** | Check-ins ask for presence, not creativity. Low friction, high meaning. | Free text inputs, elaborate profile customization |
| **Firmness Without Friction** | Constraints feel architectural, not punitive. Limits empower, not frustrate. | Angry error states, "pay to remove limits" pressure |

---

### 1.2 Interaction Philosophy

**Discouraging Mindless Scrolling:**
- No infinite scroll. Discovery ends at 5 cards.
- No pull-to-refresh gesture. Refresh is explicit, button-driven, and cooldown-gated.
- No algorithmic feed. Cards are static until the user deliberately refreshes.
- **All 5 daily cards visible at once** in a Compact Vertical Stack — no scrolling required to see all options.

**Encouraging Focused Action:**
- One primary CTA per screen (e.g., "Check In" or "Connect").
- **Card tap opens a Centered Modal** with full details — Discovery list remains visible behind overlay.
- Generous whitespace to reduce cognitive load.
- Progress is visible and finite (15 days, not infinite content).

**Design Rhythm:**
- The app is designed for **1-2 minute daily visits**, not 30-minute sessions.
- UI should feel calm, not urgent. No countdown timers that create anxiety.

---

## 2. Foundations

### 2.1 Color Palette

#### Brand Core

```css
/* Primary Actions */
--unora-primary: #E8DED5;        /* Warm Sand — Default surfaces */
--unora-primary-accent: #C4A77D; /* Golden Tan — Primary CTA */
--unora-primary-hover: #B08D5B;  /* Deeper Gold — Hover state */

/* Ink & Text */
--unora-ink-primary: #1A1A1A;    /* Near Black — Headlines */
--unora-ink-secondary: #4A4A4A;  /* Dark Gray — Body text */
--unora-ink-tertiary: #7A7A7A;   /* Medium Gray — Captions */
--unora-ink-muted: #A3A3A3;      /* Light Gray — Placeholders */
```

#### Server Themes

Each server has a distinct visual identity. The theme is applied to:
- Server tab indicator (bottom of tab)
- Card accent border (left edge)
- Primary action button background
- Status bar tint

```css
/* 🔥 Partner Server (Romantic) — Warm & Terracotta */
--server-partner-primary: #D4714A;    /* Terracotta */
--server-partner-secondary: #F2E0D5;  /* Blush sand */
--server-partner-accent: #8B3A2F;     /* Deep brick */
--server-partner-surface: #FAF5F2;    /* Warm white */

/* 👋 Friend Server (Platonic) — Teal & Sage */
--server-friend-primary: #4A9B8C;     /* Sage teal */
--server-friend-secondary: #D9EDE8;   /* Mint foam */
--server-friend-accent: #2D6B5E;      /* Deep forest */
--server-friend-surface: #F5FAF8;     /* Cool white */

/* 🎯 Growth Server (Accountability) — Indigo & Electric */
--server-growth-primary: #5B6ABF;     /* Electric indigo */
--server-growth-secondary: #E0E3F5;   /* Lavender mist */
--server-growth-accent: #3D4785;      /* Deep navy */
--server-growth-surface: #F7F8FC;     /* Blue white */
```

#### Status Colors (Streak States)

```css
/* Streak States */
--status-active: #4A9B8C;         /* Teal — Streak healthy */
--status-active-bg: #E8F5F2;      /* Light teal surface */

--status-at-risk: #E6A43A;        /* Amber — At risk */
--status-at-risk-bg: #FDF4E3;     /* Light amber surface */

--status-payment: #D4714A;        /* Terracotta — Payment window */
--status-payment-bg: #F9EAE3;     /* Light terracotta surface */

--status-reset: #7A7A7A;          /* Gray — Reset/Neutral */
--status-reset-bg: #F2F2F2;       /* Light gray surface */

/* System Feedback */
--feedback-success: #4A9B8C;      /* Teal */
--feedback-warning: #E6A43A;      /* Amber */
--feedback-error: #C94A4A;        /* Soft red */
--feedback-info: #5B6ABF;         /* Indigo */
```

#### Neutral/Surface (Light Mode Primary)

```css
/* Surfaces — Light Mode First */
--surface-background: #FAFAF8;    /* Off-white canvas */
--surface-card: #FFFFFF;          /* Pure white cards */
--surface-elevated: #FFFFFF;      /* Modals, sheets */
--surface-overlay: rgba(26, 26, 26, 0.6); /* Backdrop */

/* Borders & Dividers */
--border-subtle: #E8E8E6;         /* Card borders */
--border-medium: #D4D4D0;         /* Input borders */
--border-strong: #A3A3A3;         /* Focused borders */
```

---

### 2.2 Typography

**Font Stack:** [Inter](https://fonts.google.com/specimen/Inter) (Primary), [Outfit](https://fonts.google.com/specimen/Outfit) (Display/Headlines)

Inter provides excellent legibility for UI, especially on mobile. Outfit adds premium character to headlines.

```css
/* Font Imports */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Outfit:wght@500;600;700&display=swap');

/* Type Scale */
--font-display: 'Outfit', sans-serif;
--font-body: 'Inter', sans-serif;

/* Heading Styles */
--h1-size: 28px;
--h1-weight: 700;
--h1-line-height: 1.2;
--h1-letter-spacing: -0.02em;
--h1-font: var(--font-display);

--h2-size: 22px;
--h2-weight: 600;
--h2-line-height: 1.25;
--h2-letter-spacing: -0.01em;
--h2-font: var(--font-display);

--h3-size: 18px;
--h3-weight: 600;
--h3-line-height: 1.3;
--h3-letter-spacing: 0;
--h3-font: var(--font-display);

--h4-size: 16px;
--h4-weight: 600;
--h4-line-height: 1.35;
--h4-letter-spacing: 0;
--h4-font: var(--font-body);

/* Body Styles */
--body-large-size: 16px;
--body-large-weight: 400;
--body-large-line-height: 1.5;

--body-size: 14px;
--body-weight: 400;
--body-line-height: 1.5;

--body-small-size: 13px;
--body-small-weight: 400;
--body-small-line-height: 1.45;

/* Caption & Label */
--caption-size: 12px;
--caption-weight: 500;
--caption-line-height: 1.4;
--caption-letter-spacing: 0.01em;

--label-size: 11px;
--label-weight: 600;
--label-line-height: 1.3;
--label-letter-spacing: 0.04em;
--label-transform: uppercase;
```

---

### 2.3 Iconography

**Style:** Outlined icons (2px stroke weight) for navigation and actions. Filled icons for selected/active states only.

**Icon Library:** [Phosphor Icons](https://phosphoricons.com/) — Flexible, modern, excellent consistency.

```
Server Icons:
├── 🔥 Partner: HeartStraight (outline) / HeartStraightFill (active)
├── 👋 Friend:  HandWaving (outline) / HandWavingFill (active)
└── 🎯 Growth:  Target (outline) / TargetFill (active)

Action Icons:
├── Refresh:    ArrowsClockwise
├── Connect:    Handshake
├── Nudge:      BellRinging
├── Check-in:   CheckCircle
├── Reveal:     Eye / EyeSlash
└── Chat:       ChatCircle

State Icons:
├── Locked:     Lock / LockSimple
├── Unlocked:   LockOpen
├── Timer:      Timer
└── Warning:    Warning
```

**Icon Sizing:**

| Context | Size | Touch Target |
|---------|------|--------------|
| Navigation (bottom bar) | 24px | 48px |
| Action buttons | 20px | 44px |
| Inline icons | 16px | — |
| Card icons (hobbies) | 20px | — |

---

### 2.4 Spacing & Grid

**Base Unit:** 4px. All spacing values are multiples of 4.

```css
/* Spacing Scale */
--space-0: 0px;
--space-1: 4px;
--space-2: 8px;
--space-3: 12px;
--space-4: 16px;
--space-5: 20px;
--space-6: 24px;
--space-7: 32px;
--space-8: 40px;
--space-9: 48px;
--space-10: 64px;

/* Component Spacing */
--padding-card: 16px;
--padding-screen: 20px;
--gap-inline: 8px;
--gap-stack: 12px;
--gap-section: 32px;

/* Border Radius */
--radius-sm: 8px;
--radius-md: 12px;
--radius-lg: 16px;
--radius-xl: 24px;
--radius-full: 9999px;
```

**Grid System:**

- Screen padding: 20px
- Max content width: 400px (centered on larger screens)

**Discovery Layout — Compact Vertical Stack:**

All 5 daily Discovery cards are displayed in a **single-screen vertical stack** of horizontal rectangle cards — similar to a playlist or email inbox. No scrolling is required to view all available options.

```
┌─────────────────────────────────────────────────────────────┐
│   Discovery                               [Filter] [↻]     │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ ▌ 🏋️ Gym                              ●●●●○  High  │   │  ← Card 1
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ ▌ 📚 Reading                          ●●●○○  Reg   │   │  ← Card 2
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ ▌ 🎨 Painting                         ●●●●●  V.Hi  │   │  ← Card 3
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ ▌ 🧘 Meditation                       ●●○○○  Calm  │   │  ← Card 4
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ ▌ 🎸 Guitar                           ●●●●○  High  │   │  ← Card 5
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

Stack Specifications:
├── **Card type:** Discovery Teaser (horizontal rectangle, compact)
├── **Card height:** 80-90px
├── **Card width:** Full width - 40px (screen padding)
├── **Card gap:** 8px (tight stack to fit all 5)
├── **Interaction:** Tap opens Card Detail Modal
└── **Constraint:** All 5 cards must be visible without scrolling

---

## 3. Component Library

### 3.1 Buttons

#### Primary Button (Connect / Check-In)

The primary action per screen. Uses server theme color when on a server screen.

```
┌─────────────────────────────────────────────────┐
│  ┌─────────────────────────────────────────┐    │
│  │                                         │    │
│  │           💚 Check In                   │    │  Height: 52px
│  │                                         │    │  Radius: 12px
│  └─────────────────────────────────────────┘    │  Padding: 16px 24px
│                                                 │
│  Background: var(--server-*-primary)            │
│  Text: #FFFFFF                                  │
│  Font: Inter 16px / 600                         │
│  Shadow: 0 2px 8px rgba(0,0,0,0.08)             │
└─────────────────────────────────────────────────┘

States:
├── Default: Full opacity
├── Pressed: Scale 0.98, opacity 0.9
├── Disabled: Opacity 0.4, no interactions
└── Loading: Spinner replaces icon, text fades to 60%
```

#### Secondary Button (Nudge / View Details)

```
┌─────────────────────────────────────────────────┐
│  ┌─────────────────────────────────────────┐    │
│  │           🔔 Nudge Partner              │    │  Height: 44px
│  └─────────────────────────────────────────┘    │  Radius: 10px
│                                                 │
│  Background: transparent                        │
│  Border: 1.5px solid var(--border-medium)       │
│  Text: var(--unora-ink-secondary)               │
│  Font: Inter 14px / 500                         │
└─────────────────────────────────────────────────┘
```

#### Disabled / Global Lock Button

When Discovery is locked (time or capacity), the Refresh button enters this state.

```
┌─────────────────────────────────────────────────┐
│  ┌─────────────────────────────────────────┐    │
│  │           🔒 Refresh in 4h 23m          │    │  
│  └─────────────────────────────────────────┘    │  
│                                                 │
│  Background: var(--surface-card)                │
│  Border: 1.5px dashed var(--border-subtle)      │
│  Text: var(--unora-ink-tertiary)                │
│  Icon: Lock (solid) in muted color              │
│  Font: Inter 14px / 500                         │
│                                                 │
│  ⚠️ NO red or error colors — this is expected   │
│     behavior, not an error state.               │
└─────────────────────────────────────────────────┘
```

---

### 3.2 Discovery Cards

Cards are the core discovery unit — no photos, only hobby anchors and consistency signals. The Discovery screen uses **two card variants**: a compact teaser for the list and a full detail card shown in a modal.

---

#### Variant A: Discovery Teaser (List Item)

The **Discovery Teaser** is the compact card displayed in the Discovery screen's vertical stack. It provides just enough information to spark curiosity without revealing full details.

```
┌─────────────────────────────────────────────────────────────┐
│ ▌                                                           │
│ ▌  ┌────┐                                                   │
│ ▌  │ 🏋️ │  Gym                              ●●●●○  High    │
│ ▌  └────┘                                                   │
│ ▌                                                           │
└─────────────────────────────────────────────────────────────┘
```

**Discovery Teaser Specifications:**

| Property | Value |
|----------|-------|
| **Height** | 80-90px |
| **Width** | Full width - 40px (screen padding) |
| **Left border** | 4px solid `var(--server-*-primary)` |
| **Background** | `var(--surface-card)` |
| **Radius** | 12px |
| **Shadow** | 0 2px 8px rgba(0,0,0,0.04) |
| **Padding** | 16px |

**Content (Minimal):**

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌────┐                                         ●●●●○      │
│   │ 🏋️ │  Gym                                   High       │
│   └────┘                                                    │
│                                                             │
│   Left: Hobby Icon (24px) + Hobby Title (Inter 16px / 600)  │
│   Right: Consistency Dots (5 dots, 8px each) + Label        │
│                                                             │
│   ⚠️ NO micro-copy, NO intent statement, NO buttons.         │
│      These appear only in the Detail Card (modal).           │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Consistency Dots (Compact Indicator):**

```
Calm              │  Regular           │  Highly Consistent
○○○○○             │  ●●●○○             │  ●●●●●
(all empty)       │  (partial fill)    │  (all filled)

● = Filled dot (server color)
○ = Empty dot (var(--border-subtle))
Dot size: 8px
Gap: 3px
```

**Tap State:**
- Scale: 0.98 on press
- Background: Subtle highlight (`var(--server-*-secondary)` at 20%)
- Opens: Card Detail Modal

---

#### Variant B: Detail Card (Modal Content)

The **Detail Card** is the full-featured card displayed inside the Card Detail Modal when a user taps a Discovery Teaser. It contains all information needed to make a connection decision.

```
┌─────────────────────────────────────────────────────────────┐
│ ▌                                                           │
│ ▌  ┌───────────────────────────────────────────────────┐    │
│ ▌  │                                                   │    │
│ ▌  │  ┌─────────────────────────────────────────────┐  │    │
│ ▌  │  │  🏋️ Gym                                     │  │    │ ← Hobby Anchor
│ ▌  │  │  "Consistency over intensity"               │  │    │ ← Micro-copy
│ ▌  │  └─────────────────────────────────────────────┘  │    │
│ ▌  │                                                   │    │
│ ▌  │  ┌─────────────────────────────────────────────┐  │    │
│ ▌  │  │  🎨 Painting                                │  │    │
│ ▌  │  │  "Late nights, slow strokes"               │  │    │
│ ▌  │  └─────────────────────────────────────────────┘  │    │
│ ▌  │                                                   │    │
│ ▌  │  ─────────────────────────────────────────────    │    │
│ ▌  │                                                   │    │
│ ▌  │  ┌────────────┐                                   │    │
│ ▌  │  │ ▓▓▓▓▓▓░░░ │  Highly Consistent                │    │ ← Consistency Band
│ ▌  │  └────────────┘                                   │    │
│ ▌  │                                                   │    │
│ ▌  │  "Seeking someone who values slow mornings."      │    │ ← Intent Statement
│ ▌  │                                                   │    │
│ ▌  └───────────────────────────────────────────────────┘    │
│ ▌                                                           │
│ ▌  ┌───────────────────────────────────────────────────┐    │
│ ▌  │                   🤝 Connect                      │    │ ← Action
│ ▌  └───────────────────────────────────────────────────┘    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Detail Card Specifications:**

| Property | Value |
|----------|-------|
| **Left border** | 4px solid `var(--server-*-primary)` |
| **Background** | `var(--surface-card)` |
| **Radius** | 16px |
| **Shadow** | 0 2px 12px rgba(0,0,0,0.06) |
| **Padding** | 20px |

**Full Content Hierarchy:**

1. **Hobby Anchors** — Icon + Title + Micro-copy (up to 2)
2. **Consistency Band** — Full visual bar + label
3. **Intent Statement** — User's connection intent (full text)
4. **Connect Button** — Primary CTA

---

#### Hobby Anchor Component

```
┌───────────────────────────────────────────────────┐
│   ┌────┐                                          │
│   │ 🏋️ │  Gym                                     │  Icon: 20px, server color
│   └────┘  "Consistency over intensity"            │  Title: Inter 15px / 600
│                                                   │  Micro: Inter 13px / 400 / italic
│   Background: var(--server-*-secondary)           │     color: var(--unora-ink-tertiary)
│   Padding: 12px 16px                              │
│   Radius: 10px                                    │
└───────────────────────────────────────────────────┘
```

#### Consistency Band (Full)

Three visual states representing streak history:

```
Calm              │  Regular           │  Highly Consistent
░░░░░░░░░░        │  ▓▓▓▓▓░░░░░        │  ▓▓▓▓▓▓▓▓▓▓
Gray blocks       │  Half-filled       │  Fully filled
                  │  (server color)    │  (server color)

Label: Caption style, positioned right of band
Band Width: 80px
Block Height: 6px
Gap between blocks: 2px
```

---

#### Personality Summary Block (Optional Component)

The Personality Summary Block displays an AI-generated personality insight inside the Card Detail Modal. It appears only when the viewed user has sufficient personality signals.

**Placement:**
- Below: Hobby Anchors, Consistency Band, Intent Statement
- Above: Connect Button (CTA)

**Wireframe:**

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ ▌                                                                           │
│ ▌  [Hobby Anchors]                                                          │
│ ▌  [Consistency Band]                                                       │
│ ▌  [Intent Statement]                                                       │
│ ▌                                                                           │
│ ▌  ──────────────────────────────────────────────────────────────────────   │
│ ▌                                                                           │
│ ▌  They tend to process ideas internally before sharing them. In group      │
│ ▌  settings, they often observe before contributing. Their energy seems     │
│ ▌  to renew through solo activities, though they value depth in             │
│ ▌  one-on-one conversations.                                                │
│ ▌                                                                           │
│ ▌  ──────────────────────────────────────────────────────────────────────   │
│ ▌                                                                           │
│ ▌  ┌─────────────────────────────────────────────────────────────────────┐  │
│ ▌  │                        🤝 Connect                                   │  │
│ ▌  └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Visual Specifications:**

| Property | Value |
|----------|-------|
| **Typography** | Inter 14px / 400 (body), line-height 1.6 |
| **Text Color** | `var(--unora-ink-secondary)` |
| **Background** | None (transparent, inherits card background) |
| **Max Lines** | 5-6 lines (approx. 50-90 words) |
| **Padding** | 0 (uses parent card padding) |
| **Top/Bottom Divider** | 1px solid `var(--border-subtle)`, 16px margin |
| **Interaction** | Read-only (no taps, no expansion) |

**Absence Behavior:**
- If user has insufficient personality signals: **Section is omitted entirely**
- No "locked" state, no placeholder, no teaser text
- Card Detail Modal simply displays without this section

**Copy Principles:**

| ❌ Never Use | ✓ Instead Use |
|--------------|---------------|
| "Personality type" | (describe naturally) |
| "Assessment result" | (describe naturally) |
| "Introvert/Extrovert" | Behavioral descriptions |
| "Score", "Rating", "Level" | (omit entirely) |
| First-person ("I", "My") | Third-person ("They", "Their") |
| "Test", "Quiz", "Assessment" | "Context", "Signals", "Insights" |

---

### 3.3 The Streak Component

#### Daily Check-In Button (Before Check-In)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Day 7                                                     │  ← Progress label
│   ═══════════════════════════════════════════════════════                    │  ← Trust indicator
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │   "Last time you crushed Leg Day.                   │   │  ← Hobby Echo prompt
│   │    What's the focus today?"                         │   │
│   │                                                     │   │
│   │   ┌──────────┐  ┌──────────┐  ┌──────────┐          │   │
│   │   │  Rest    │  │ Cardio   │  │  Legs    │          │   │  ← Tap options
│   │   └──────────┘  └──────────┘  └──────────┘          │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Tap Option Styling:
├── Default: White background, subtle border, server-colored text
├── Selected: Server-primary background, white text
└── Size: Min 72px width, 44px height, 8px radius
```

#### Check-In Status (After Check-In)

```
My Check-In                      Partner's Check-In
┌─────────────────────┐          ┌─────────────────────┐
│                     │          │                     │
│   ✓ Checked In      │          │   ⏳ Waiting...     │
│   Today at 9:42 AM  │          │                     │
│                     │          │                     │
└─────────────────────┘          └─────────────────────┘

Background: var(--status-active-bg)     Background: var(--surface-card)
Border: 1px solid var(--status-active)  Border: 1px dashed var(--border-subtle)
Icon: CheckCircle (filled, green)       Icon: Timer (outlined, muted)
```

#### The Hobby Echo (Partner Progress)

Displayed upon mutual check-in completion:

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │         ✨ Streak Extended! Day 8 ✨                │   │
│   │                                                     │   │
│   │   ┌─────────────────────────────────────────────┐   │   │
│   │   │                                             │   │   │
│   │   │  Your partner had a                         │   │   │
│   │   │  💪 High Intensity session today.           │   │   │
│   │   │                                             │   │   │
│   │   └─────────────────────────────────────────────┘   │   │
│   │                                                     │   │
│   │              [ Continue → ]                         │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Container: Centered modal card
Background: var(--surface-elevated)
Shadow: 0 8px 32px rgba(0,0,0,0.12)
Radius: 20px
Animation: Fade-in from center + subtle scale (0.95 → 1.0)

Emoji styling: 28px, positioned left of text
Text: Inter 16px / 500
```

---

### 3.4 Inputs (Tap-Based Only)

No free text inputs for check-ins. All selections are tap-based.

```
Choice Chip Grid (for check-in options):

┌────────────┐  ┌────────────┐  ┌────────────┐
│    Rest    │  │  Cardio    │  │   Legs     │
└────────────┘  └────────────┘  └────────────┘
┌────────────┐  ┌────────────┐
│   Upper    │  │  Full Body │
└────────────┘  └────────────┘

Styling:
├── Background: var(--surface-card)
├── Border: 1.5px solid var(--border-medium)
├── Selected Border: 2px solid var(--server-*-primary)
├── Selected Background: var(--server-*-secondary)
├── Text: Inter 14px / 500
├── Min Size: 72px × 44px
├── Radius: 8px
└── Gap: 8px
```

---

### 3.5 Personality Setup UI (Profile Tab)

The Personality Setup UI is an optional card displayed in the Profile tab, allowing users to add personality context at their own pace.

#### Setup Card (Incomplete State)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                                                                     │   │
│   │   ✦ Add personality context                                        │   │
│   │                                                                     │   │
│   │   Help others understand how you connect.                          │   │
│   │   Takes about 2 minutes.                                           │   │
│   │                                                                     │   │
│   │   ┌────────────────────────┐   ┌────────────────────────┐          │   │
│   │   │    Get Started →       │   │   I'll do this later   │          │   │
│   │   └────────────────────────┘   └────────────────────────┘          │   │
│   │                                                                     │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Specifications:**

| Property | Value |
|----------|-------|
| **Card Background** | `var(--surface-card)` |
| **Border** | 1px solid `var(--border-subtle)` |
| **Border Radius** | 12px |
| **Icon** | ✦ (sparkle), `var(--unora-primary-accent)` |
| **Title** | Inter 16px / 600, `var(--unora-ink-primary)` |
| **Description** | Inter 14px / 400, `var(--unora-ink-secondary)` |
| **Primary CTA** | "Get Started →" — Secondary button style |
| **Skip CTA** | "I'll do this later" — Text link, `var(--unora-ink-tertiary)` |

**Copy Principles:**

| ❌ Never Say | ✓ Instead Say |
|--------------|---------------|
| "Take the personality test" | "Add personality context" |
| "Complete your assessment" | "Help others understand how you connect" |
| "Your profile is incomplete" | (no urgency language) |
| "X% complete" | (no progress bars implying obligation) |

#### Setup Card (Complete State)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │   ✓ Personality context added                                       │   │
│   │   Others can now see insights about how you connect.               │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Complete State Specs:**
- Background: `var(--status-active-bg)` (light teal)
- Icon: ✓ checkmark, `var(--status-active)`
- Text: Inter 14px / 500, `var(--unora-ink-secondary)`
- No edit action shown (signals are accumulative, not editable)

---

## 4. Visualizing Logic & States

### 4.1 The Global Lock Screen

When Discovery is locked (Time or Capacity), the UI must feel **firm but not punitive**. This is intentional product design, not an error.

#### Lock Type: Time-Based (Cooldown)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │             ⏱️                                      │   │  Icon: Timer, 48px
│   │                                                     │   │
│   │   Your next refresh is brewing.                     │   │  Headline: H3
│   │                                                     │   │
│   │   ━━━━━━━━━━━━━━━━━━░░░░░░░                         │   │  Progress bar (time)
│   │                                                     │   │
│   │   Available in 4h 23m                               │   │  Caption: time remaining
│   │                                                     │   │
│   │   ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─                       │   │
│   │                                                     │   │
│   │   While you wait, focus on your active streak.      │   │  Encouraging sub-copy
│   │                                                     │   │
│   │              [ View Streak → ]                      │   │  Secondary CTA
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Background: var(--surface-background)
Card: var(--surface-card) with subtle shadow
Tone: Calm, patient. No urgency language.
```

#### Lock Type: Capacity-Based (Slots Full)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │             🎯                                      │   │  Icon: Target, 48px
│   │                                                     │   │
│   │   You're fully focused.                             │   │  Headline: H3
│   │                                                     │   │
│   │   Your 1 active connection has your full            │   │  Body text
│   │   attention right now.                              │   │
│   │                                                     │   │
│   │   ┌─────────────────────────────────────────────┐   │   │
│   │   │  ▓ 1/1 Connection Active                   │   │   │  Slot indicator
│   │   └─────────────────────────────────────────────┘   │   │
│   │                                                     │   │
│   │   ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─                       │   │
│   │                                                     │   │
│   │   Maintain it, or upgrade for more connections.     │   │  Options
│   │                                                     │   │
│   │   [ View Streak ]        [ Upgrade → ]              │   │  Two CTAs
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Upgrade CTA: Secondary style, not aggressive upsell
Language: "Upgrade" not "Unlock" — emphasizes capability, not restriction
```

---

### 4.2 Bottom Navigation Bar (3-Tab Model)

> [!IMPORTANT]
> Unora uses a **3-Tab Bottom Navigation** system. Tabs represent app sections, NOT servers. Server context is controlled by a top-left dropdown.

**Reference:** [UserFlow Section 1.1 — Navigation Architecture](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

#### Wireframe

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│                      [ Screen Content ]                         │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │  ← Floating Glass Bar
│  ░                                                           ░  │
│  ░   🔥            🃏             👤                         ░  │  ← Icons: 24px
│  ░  Streaks      Discovery      Profile                      ░  │  ← Labels: 11px
│  ░               ═══════                                     ░  │  ← Active indicator
│  ░                                                           ░  │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

#### Tab Specifications

| Tab | Icon (Phosphor) | Behavior | Server Switcher |
|-----|-----------------|----------|-----------------|
| **Streaks** | Fire / Flame | Global Aggregator — shows ALL server connections | ✗ Hidden |
| **Discovery** | Cards / Unora Logo | Contextual — shows current server only | ✓ Visible |
| **Profile** | User / Avatar | Contextual — theme matches current server | ✓ Visible |

#### Floating Glass Aesthetic (Premium Dark Mode)

```css
/* Bottom Navigation — Floating Glass */
.bottom-nav {
  position: fixed;
  bottom: 20px;
  left: 16px;
  right: 16px;
  height: 64px;
  
  /* Glassmorphism */
  background: rgba(13, 13, 15, 0.85);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  
  /* Subtle border + glow */
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 20px;
  
  /* Elevation */
  box-shadow: 
    0 8px 32px rgba(0, 0, 0, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.05);
}
```

#### Tab States

| State | Icon | Label Color | Indicator |
|-------|------|-------------|-----------|
| **Active** | Filled (solid) | `var(--unora-primary-accent)` | 3px bar below icon |
| **Inactive** | Outlined | `var(--dm-ink-tertiary)` | None |

```
Active Tab:
├── Icon: Solid fill, color: var(--unora-primary-accent)
├── Label: var(--unora-primary-accent), Inter 11px/500
├── Indicator: 3px × 24px bar, var(--unora-primary-accent)
└── Touch area: 64px × 64px minimum

Inactive Tab:
├── Icon: Outline only, color: var(--dm-ink-tertiary)
├── Label: var(--dm-ink-tertiary), Inter 11px/400
└── No indicator bar
```

---

### 4.2.1 Server Switcher Component

The **Server Switcher** is a dropdown in the **Top-Left corner** of contextual screens. It is NOT in the bottom navigation.

#### Visibility Rules

| Screen | Server Switcher |
|--------|-----------------|
| Streaks Tab | ✗ **Hidden** (global view) |
| Discovery Tab | ✓ **Visible** |
| Profile Tab | ✓ **Visible** |

#### Wireframe

```
┌─────────────────────────────────────────────────────────────────┐
│   🔥 Partner ▼                                      ⚙️         │  ← Header
│   ════════════                                                  │
│                                                                 │
│   Tapping opens dropdown:                                       │
│   ┌────────────────────────┐                                    │
│   │ 🔥 Partner          ✓  │  ← Active (check mark)             │
│   ├────────────────────────┤                                    │
│   │ 👋 Friend               │                                    │
│   ├────────────────────────┤                                    │
│   │ 🎯 Growth               │                                    │
│   └────────────────────────┘                                    │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

#### Dropdown Specs

```
SERVER SWITCHER DROPDOWN
├── Width: 180px
├── Background: var(--dm-surface-elevated)
├── Border: 1px solid var(--dm-border-subtle)
├── Border radius: 12px
├── Shadow: 0 8px 24px rgba(0,0,0,0.4)
│
├── [ITEMS]
│   ├── Height: 48px per item
│   ├── Padding: 12px 16px
│   ├── Icon: 20px, server color
│   ├── Label: Inter 15px/500, var(--dm-ink-primary)
│   └── Active indicator: Checkmark (right-aligned)
│
└── [ANIMATION]
    ├── Open: Fade in + scale 0.95 → 1.0 (150ms)
    └── Close: Fade out (100ms)
```

---

### 4.2.2 Dynamic Server Theming (Profile Tab)

When on the **Profile Tab**, UI accent colors dynamically change to match the currently selected server.

| Selected Server | Primary Accent | Secondary | CTA Background |
|-----------------|----------------|-----------|----------------|
| Partner | `#E07D5A` (Terracotta) | `#2D2420` | `#E07D5A` |
| Friend | `#5DB3A2` (Teal) | `#1D2926` | `#5DB3A2` |
| Growth | `#7B8AD9` (Indigo) | `#1E2038` | `#7B8AD9` |

**Affected Elements:**
- Section headers
- Button backgrounds
- Toggle/switch accents
- Active border highlights
- Subscription card accents

---

### 4.3 Connection Core (Trust Orb) — Mystery Reveal Visual

> [!IMPORTANT]
> **Users never see numbered milestones, day counts, or reveal timelines.** The Connection Core is an abstract visual that represents streak strength through glowing intensity, not explicit progress markers.

#### Visual Language

```
    Low Intensity          Medium Intensity         High Intensity
    (Day 1-5)              (Day 6-10)              (Day 11-15)

       ┌───┐                  ┌───┐                  ┌───┐
       │ ✦ │                  │✦✦✦│                  │✦✦✦│
       └───┘                  └───┘                  └───┘
      Dim pulse             Steady glow           Radiant shimmer

✦ = Glow intensity (more stars = brighter visual)
```

#### Component Specifications

```
CONNECTION CORE (Abstract Progress Visual)
├── Container: 80px × 80px centered
├── Orb diameter: 64px
├── Background: Glassmorphic with subtle blur
│
├── [VISUAL APPEARANCE]
│   ├── Shape: Circular gradient sphere
│   ├── Base: var(--glass-fill) with [server color] tint
│   ├── Glow: Inner glow using server glow-color token
│   └── Aesthetic: Bioluminescent, ethereal, NOT neon
│
├── [INTENSITY STATES] (internal day mapping)
│   ├── Dim:    Opacity 0.3, scale 0.95, soft pulse
│   ├── Medium: Opacity 0.6, scale 1.0, steady glow
│   ├── Bright: Opacity 0.9, scale 1.05, intensifying pulse
│   └── Maximum: Opacity 1.0, scale 1.1, radiant shimmer
│
└── [ANIMATION]
    ├── Breathing pulse: scale 1.0 → 1.03 (2s loop)
    ├── Glow intensity variation: opacity ±0.1 (3s loop)
    └── Respects prefers-reduced-motion
```

#### Status Text Options

| Intensity | Copy (Never mentions days) |
|-----------|-----------------------------|
| Low | "Trust building..." |
| Medium | "Connection deepening..." |
| High | "Something is forming..." |
| Maximum | "The bond is ready." |

#### Compact Card Version

```
┌─────────────────────────────────────────────────────────────┐
│   Day 7 • ✦ Trust building...                                 │
│   ──────────────────────────────  (subtle glow line)                │
└─────────────────────────────────────────────────────────────┘

No milestone markers, no countdown text, no "X of 15" labels.
```

---

### 4.4 Blur Logic (Reveal Visual Treatment)

| Data Type | Pre-Reveal Treatment | Post-Reveal |
|-----------|---------------------|-------------|
| **Photos** | Not shown at all (no blurred placeholder) — replaced with abstract avatar | Full clarity |
| **Name** | Shows "Your Partner" label only | Full name revealed |
| **Age/Location** | Hidden behind reveal card | Visible inline |
| **Voice Note** | Grayed-out waveform with lock icon | Playable |

**Abstract Avatar (No Photo State):**

```
┌───────────────────────┐
│                       │
│         ◯◯           │
│        ╱──╲          │
│       │    │         │
│                       │
│  Gradient background  │
│  using server colors  │
│                       │
└───────────────────────┘

Size: 64px × 64px (profile), 48px (inline)
Background: Linear gradient of server secondary → primary at 15%
Shape: Rounded square (12px radius)
Icon: Minimal face outline in server-accent color at 40% opacity
```

**Reveal Animation:**
- Duration: 600ms
- Easing: ease-out
- Effect: Fade in from 0% + slight scale from 0.95
- Confetti/particles: Optional 300ms celebration on Reveal 5 only

---

### 4.X Verification UI

**PRD Reference:** [Section 10: Progressive Verification](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)

This section defines the visual language for progressive verification components. All verification UI must feel **calm and trust-forward** — no red warnings, no urgency, no gamified badges.

---

#### 4.X.1 Photo Privacy States

Photos uploaded during onboarding have the following visual states:

| State | Visual Treatment |
|-------|------------------|
| **Locked (Default)** | Heavy blur (20px radius) + lock icon overlay at 40% opacity |
| **Uploading** | Progress indicator over blur, subtle pulse animation |
| **Submitted** | Checkmark overlay (green, subtle), blur maintained |
| **Revealed (Day 15)** | Full clarity, no overlay, smooth 600ms reveal animation |

```css
/* Photo Privacy States */
--photo-locked-blur: 20px;
--photo-locked-overlay: rgba(26, 26, 26, 0.3);
--photo-submitted-icon: var(--semantic-success);
--photo-reveal-duration: 600ms;
--photo-reveal-easing: ease-out;
```

**Lock Icon Overlay:**
```
┌────────────────────────────────┐
│                                │
│        ┌──────────┐            │
│        │    🔒    │            │   Lock icon: 24px
│        └──────────┘            │   Centered vertically/horizontally
│                                │   Opacity: 40%
│       [Blurred Photo]          │   Color: var(--unora-ink-tertiary)
│                                │
└────────────────────────────────┘
```

---

#### 4.X.2 Liveness Challenge Visuals

The liveness challenge UI must feel **calm and non-intrusive**.

**Camera Frame with Guide:**

```
┌────────────────────────────────────────────────────────────────┐
│                                                                │
│                    ┌──────────────────────┐                    │
│                    │                      │                    │
│                    │    ╭──────────────╮  │                    │
│                    │    │              │  │   Dotted oval guide│
│                    │    │   (face)     │  │   Color: --server-*-primary│
│                    │    │              │  │   at 60% opacity   │
│                    │    ╰──────────────╯  │                    │
│                    │                      │                    │
│                    └──────────────────────┘                    │
│                                                                │
│    "Position your face in the frame"                           │
│                                                                │
│               [ Capture ]        [ Skip ]                      │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```

**States:**

| State | Visual Treatment |
|-------|------------------|
| **Awaiting** | Dotted oval guide, instruction text visible |
| **Capturing** | Subtle pulse on oval guide, "Stay still" text |
| **Processing** | Spinner in center, oval fades to 30% opacity |
| **Success** | Green checkmark, "You're verified" text |
| **Retry** | Yellow oval (brief), "Let's try again" text |

**Color Tokens:**
```css
/* Liveness Challenge */
--liveness-guide-default: var(--server-*-primary);
--liveness-guide-success: var(--semantic-success);
--liveness-guide-retry: var(--semantic-warning);
--liveness-text-primary: var(--unora-ink-primary);
--liveness-text-secondary: var(--unora-ink-secondary);
```

---

#### 4.X.3 Trust Indicators

Trust indicators are **visible only to the user** (never to matches) and appear in the Trust & Verification settings screen.

**Timeline Format:**

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Your Trust Status                                                          │
│                                                                             │
│  ────╼━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╾────        │
│                                                                             │
│  ✓ Phone verified                                                           │
│    └── Verified on Jan 9, 2026                                              │
│                                                                             │
│  ✓ Photos submitted                                                          │
│    └── 4 photos uploaded                                                    │
│                                                                             │
│  ⏳ Liveness check                                                          │
│    └── Available during your streak                                         │
│    └── [Complete now →]                                                     │
│                                                                             │
│  ⏳ Consistency                                                              │
│    └── Day 7 of first streak                                                │
│                                                                             │
│  ○ Government ID (Optional)                                                 │
│    └── Unlocks profile flexibility                                          │
│    └── Available after streak completion                                    │
│                                                                             │
│  ──────────────────────────────────────────────────────────────────────────│
│  "Trust builds through consistency. You're always in control."              │
└─────────────────────────────────────────────────────────────────────────────┘
```

**Icon Tokens:**

| State | Icon | Color |
|-------|------|-------|
| Complete | ✓ (checkmark) | `var(--semantic-success)` |
| In Progress | ⏳ (hourglass) | `var(--unora-ink-secondary)` |
| Pending/Optional | ○ (empty circle) | `var(--unora-ink-muted)` |
| Action Available | → (arrow) | `var(--unora-primary-accent)` |

**Design Rules:**
- No red warning icons
- No countdown timers on verification
- No percentage scores or gamified progress bars
- No badges visible to other users
- No "unverified" labels that feel punitive

---

#### 4.X.4 Verification States

General verification states for inline display:

```css
/* Verification State Colors */
--verification-complete: var(--semantic-success);
--verification-pending: var(--unora-ink-secondary);
--verification-available: var(--unora-primary-accent);
--verification-optional: var(--unora-ink-muted);
```

**Inline Text States:**

| State | Text | Color | Icon |
|-------|------|-------|------|
| Complete | "Verified" | --semantic-success | ✓ |
| In Progress | "Building trust" | --unora-ink-secondary | ⏳ |
| Available | "Action available" | --unora-primary-accent | → |
| Skipped | "Skipped — retry anytime" | --unora-ink-muted | ― |

---

#### 4.X.5 Copy Principles for Verification

> [!IMPORTANT]
> **All verification copy must be calm and trust-forward.**

| ❌ Never Say | ✓ Instead Say |
|--------------|---------------|
| "KYC" | "Trust verification" |
| "Government verification required" | "Optional identity confirmation" |
| "You must verify" | "Build your trust profile" |
| "Verification failed" | "Let's try again" |
| "Unverified account" | "Trust building in progress" |
| "Complete verification now!" | "Complete when you're ready" |
| "Your account is limited" | "More options unlock over time" |

**Tone Guidelines:**

- **Calm**: No urgency, no countdown pressure
- **Empowering**: User is in control, not being screened
- **Clear**: Simple language, no jargon
- **Non-punitive**: Skipping is okay, no shaming

---

## 5. Feedback & Haptics

### 5.1 Tactile Design Philosophy

Unora should feel **grounded and intentional**, not gamified. Haptics confirm meaningful actions, not every tap.

| Action | Haptic Type | iOS | Android |
|--------|-------------|-----|---------|
| Check-in complete | Success | `.success` | `CONFIRM` (50ms) |
| Nudge sent | Light tap | `.light` | `TICK` (20ms) |
| Match created | Heavy celebration | `.heavy` × 2 | `HEAVY_CLICK` × 2 |
| Reveal unlocked | Medium + light | `.medium` → `.light` | `CLICK` → `TICK` |
| Error / Lock | Rigid | `.rigid` | `REJECT` |

---

### 5.2 Check-In Haptics & Animation

**Check-In Success Flow:**

1. User taps check-in option
2. Option scales to 0.95 (30ms)
3. Haptic fires (success)
4. Option returns to 1.0 with checkmark overlay (200ms)
5. Card transitions to "Checked In" state (300ms fade)

```
Animation Sequence:
[0ms]    Tap detected
[30ms]   Scale to 0.95
[50ms]   Haptic: success
[80ms]   Scale to 1.0
[100ms]  Checkmark icon fades in
[300ms]  Full "Checked In" state visible
```

---

### 5.3 Nudge Interaction

**Nudge Button Animation:**

1. Bell icon wobbles (2 oscillations, 200ms)
2. Button ripple effect from center
3. Toast appears from bottom: "Nudge sent"

```
Bell Wobble Keyframes:
0%    { rotate: 0deg }
25%   { rotate: 12deg }
50%   { rotate: -12deg }
75%   { rotate: 8deg }
100%  { rotate: 0deg }

Duration: 200ms
Easing: ease-in-out
```

---

### 5.4 Page Transitions

| Transition | Animation | Duration |
|------------|-----------|----------|
| Server switch | Cross-fade | 200ms |
| **Discovery Teaser → Detail Modal** | **Scale up + fade in (centered)** | **250ms** |
| Modal open | Fade + slide up from bottom 20px | 250ms |
| Modal close | Fade + slide down | 200ms |
| Screen push | Slide from right | 300ms |

---

#### Card Detail Modal Transition

When a user taps a **Discovery Teaser**, the **Card Detail Modal** appears as a centered overlay. The Discovery list remains visible but dimmed behind the modal.

```
Open Animation Sequence:
[0ms]    Teaser tap detected
[0ms]    Haptic: Light tap
[0ms]    Overlay fades in (rgba(26,26,26,0.6))
[0ms]    Modal scales from 0.9 → 1.0, opacity 0 → 1
[250ms]  Modal fully visible, centered on screen

Close Animation Sequence:
[0ms]    Close/X tapped, overlay tapped, or swipe down
[0ms]    Modal scales 1.0 → 0.95, opacity 1 → 0
[0ms]    Overlay fades out
[200ms]  Back to Discovery list view
```

**Transition Properties:**

| Property | Value |
|----------|-------|
| **Easing (open)** | ease-out |
| **Easing (close)** | ease-in |
| **Scale origin** | Center of modal |
| **Background** | Discovery list visible but dimmed |
| **Haptic** | Light tap on modal open |

---

## 6. Accessibility & Inclusivity

### 6.1 Contrast Ratios

All text must meet **WCAG AA** standards:

| Element | Foreground | Background | Ratio | Pass |
|---------|------------|------------|-------|------|
| Body text | `#4A4A4A` | `#FAFAF8` | 8.5:1 | ✓ AAA |
| Headlines | `#1A1A1A` | `#FAFAF8` | 16:1 | ✓ AAA |
| Caption | `#7A7A7A` | `#FAFAF8` | 4.8:1 | ✓ AA |
| Primary CTA text | `#FFFFFF` | `#D4714A` | 4.6:1 | ✓ AA |
| Error text | `#C94A4A` | `#FAFAF8` | 5.2:1 | ✓ AA |

### 6.2 Touch Targets

**Minimum touch target:** 44px × 44px (Apple HIG recommendation)

| Component | Touch Size | Visual Size |
|-----------|------------|-------------|
| Navigation icons | 48px × 48px | 24px icon |
| Buttons | Full width × 52px | — |
| Choice chips | Min 72px × 44px | — |
| Card actions | 44px × 44px | 20px icon |
| Close/dismiss | 44px × 44px | 24px icon |

### 6.3 Motion Sensitivity

- All animations respect `prefers-reduced-motion`
- When reduced motion is enabled:
  - Instant transitions (no duration)
  - No wobble/bounce effects
  - Static progress indicators

### 6.4 Screen Reader Considerations

- All icons have `aria-label` text
- Check-in status announced live
- Streak progress read as "Day 7, trust building. Connection deepening."
- Server tabs announce full name + state on focus

### 6.5 Color Independence

- Never use color alone to convey status
- All status indicators include:
  - Icon (✓, ⏳, ⚠️, 🔒)
  - Text label
  - Color as enhancement only

---

## Appendix: Quick Reference Tokens

```css
/* Copy-Paste Token Reference */

:root {
  /* Brand */
  --unora-primary: #E8DED5;
  --unora-primary-accent: #C4A77D;
  --unora-ink-primary: #1A1A1A;
  --unora-ink-secondary: #4A4A4A;
  
  /* Servers */
  --server-partner: #D4714A;
  --server-friend: #4A9B8C;
  --server-growth: #5B6ABF;
  
  /* Status */
  --status-active: #4A9B8C;
  --status-at-risk: #E6A43A;
  --status-payment: #D4714A;
  
  /* Surfaces */
  --surface-bg: #FAFAF8;
  --surface-card: #FFFFFF;
  
  /* Typography */
  --font-display: 'Outfit', sans-serif;
  --font-body: 'Inter', sans-serif;
  
  /* Spacing */
  --space-unit: 4px;
  --padding-card: 16px;
  --padding-screen: 20px;
  --radius-md: 12px;
}
```

---

**Document maintained by:** Design Systems Team  
**Last updated:** January 2026  
**Review cycle:** Monthly

---

## 7. Extended Foundations — Dark Mode

### 7.1 Dark Mode Philosophy

Dark mode is not just inverted colors — it's a **separate experience** optimized for low-light contexts. Unora's dark mode should feel:

- **Restful**: Lower contrast than light mode to reduce eye strain
- **Premium**: Deep, rich surfaces (not pure black)
- **Warm**: Retains brand warmth even in dark context

### 7.2 Dark Mode Color Tokens

```css
/* Dark Mode — Surfaces */
--dm-surface-background: #121212;      /* Deep charcoal */
--dm-surface-card: #1E1E1E;            /* Elevated charcoal */
--dm-surface-elevated: #252525;        /* Modals, sheets */
--dm-surface-overlay: rgba(0, 0, 0, 0.7); /* Backdrop */

/* Dark Mode — Text */
--dm-ink-primary: #F5F5F3;             /* Off-white headlines */
--dm-ink-secondary: #C4C4C0;           /* Body text */
--dm-ink-tertiary: #8A8A86;            /* Captions */
--dm-ink-muted: #5A5A58;               /* Placeholders */

/* Dark Mode — Borders */
--dm-border-subtle: #2A2A2A;           /* Card borders */
--dm-border-medium: #3D3D3D;           /* Input borders */
--dm-border-strong: #5A5A5A;           /* Focused borders */

/* Dark Mode — Server Themes (adjusted for contrast) */
--dm-server-partner-primary: #E07D5A;  /* Lighter terracotta */
--dm-server-partner-secondary: #2D2420;/* Deep warm */
--dm-server-partner-surface: #1A1614;  /* Warm black */

--dm-server-friend-primary: #5DB3A2;   /* Lighter teal */
--dm-server-friend-secondary: #1D2926; /* Deep forest */
--dm-server-friend-surface: #141A18;   /* Cool black */

--dm-server-growth-primary: #7B8AD9;   /* Lighter indigo */
--dm-server-growth-secondary: #1E2038; /* Deep navy */
--dm-server-growth-surface: #14151F;   /* Blue black */

/* Dark Mode — Status Colors (slightly muted) */
--dm-status-active: #5DB3A2;           /* Lighter teal */
--dm-status-active-bg: #1D2926;
--dm-status-at-risk: #E6B35A;          /* Softer amber */
--dm-status-at-risk-bg: #2A2419;
--dm-status-payment: #E07D5A;          /* Softer terracotta */
--dm-status-payment-bg: #2A1F1A;
```

### 7.3 Dark Mode Card Example

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Background: #121212                                       │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │ Card: #1E1E1E
│   │  ▌  🏋️ Gym                                          │   │
│   │  ▌  "Consistency over intensity"                    │   │ Text: #F5F5F3
│   │                                                     │   │ Micro: #8A8A86
│   │     ▓▓▓▓▓▓░░░  Highly Consistent                    │   │
│   │                                                     │   │
│   │  ┌─────────────────────────────────────────────┐    │   │
│   │  │              🤝 Connect                     │    │   │ Button stays
│   │  └─────────────────────────────────────────────┘    │   │ server-colored
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Dark mode adjustments:
├── Card shadow: 0 2px 12px rgba(0,0,0,0.3) (stronger)
├── Border: 1px solid var(--dm-border-subtle) for card edges
├── Server accent bar: Same brightness as light mode
└── CTAs: Maintain full saturation for visibility
```

### 7.4 System Preference Detection

```css
@media (prefers-color-scheme: dark) {
  :root {
    --surface-background: var(--dm-surface-background);
    --surface-card: var(--dm-surface-card);
    --unora-ink-primary: var(--dm-ink-primary);
    --unora-ink-secondary: var(--dm-ink-secondary);
    /* ... map all tokens */
  }
}
```

---

## 8. Extended Component Library

### 8.1 Additional Button Variants

#### Tertiary Button (Text-Only)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│               Skip this step →                              │
│                                                             │
│  Background: transparent                                    │
│  Text: var(--server-*-primary)                              │
│  Font: Inter 14px / 500                                     │
│  Underline on hover (optional)                              │
│  No border                                                  │
│                                                             │
│  Use for: Skip actions, secondary navigation, "Learn more"  │
└─────────────────────────────────────────────────────────────┘
```

#### Icon-Only Button

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌────────┐     ┌────────┐     ┌────────┐                  │
│   │   ✕    │     │   ⚙    │     │   ?    │                  │  Size: 44px × 44px
│   └────────┘     └────────┘     └────────┘                  │  Icon: 20px
│    Close         Settings        Help                       │  Radius: 12px
│                                                             │
│  Background: var(--surface-card)                            │
│  Border: 1px solid var(--border-subtle)                     │
│  Icon: var(--unora-ink-secondary)                           │
│                                                             │
│  Hover: Border darkens to var(--border-medium)              │
│  Active: Background slightly darker                         │
└─────────────────────────────────────────────────────────────┘
```

#### Destructive Button (End Connection)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌─────────────────────────────────────────────┐           │
│   │         End This Connection                 │           │  Height: 44px
│   └─────────────────────────────────────────────┘           │
│                                                             │
│  Background: transparent                                    │
│  Border: 1.5px solid #C94A4A (soft red)                     │
│  Text: #C94A4A                                              │
│  Font: Inter 14px / 500                                     │
│                                                             │
│  ⚠️ Never use solid red background — too aggressive.        │
│     Outline style conveys seriousness without alarm.        │
└─────────────────────────────────────────────────────────────┘
```

#### Floating Action Button (FAB)

Reserved for single most important action per flow (rare use).

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                        ┌────────────┐       │
│                                        │     ✓      │       │ Size: 56px
│                                        │  Check In  │       │ Icon: 24px
│                                        └────────────┘       │ Radius: 16px
│                                                             │
│  Background: var(--server-*-primary)                        │
│  Shadow: 0 4px 16px rgba(0,0,0,0.15)                        │
│  Text: #FFFFFF                                              │
│  Position: Fixed, 20px from right, 24px from bottom         │
│                                                             │
│  Animation: Scale 1.0 → 1.05 on hover                       │
│  Haptic: Medium on tap                                      │
└─────────────────────────────────────────────────────────────┘
```

---

### 8.2 Card Variants

#### Card — Loading State (Skeleton)

```
┌─────────────────────────────────────────────────────────────┐
│ ▌                                                           │
│ ▌  ┌───────────────────────────────────────────────────┐    │
│ ▌  │                                                   │    │
│ ▌  │  ┌─────────────────────────────────────────────┐  │    │
│ ▌  │  │  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │  │    │ ← Shimmer block
│ ▌  │  │  ░░░░░░░░░░░░░░░░░░░░░░░                    │  │    │
│ ▌  │  └─────────────────────────────────────────────┘  │    │
│ ▌  │                                                   │    │
│ ▌  │  ┌─────────────────────────────────────────────┐  │    │
│ ▌  │  │  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │  │    │
│ ▌  │  │  ░░░░░░░░░░░░░░░░░░░░                       │  │    │
│ ▌  │  └─────────────────────────────────────────────┘  │    │
│ ▌  │                                                   │    │
│ ▌  │  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░                   │    │
│ ▌  │                                                   │    │
│ ▌  └───────────────────────────────────────────────────┘    │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Skeleton blocks:
├── Background: var(--surface-card)
├── Shimmer: Linear gradient sweep, left to right
├── Shimmer color: rgba(0,0,0,0.04) → rgba(0,0,0,0.08) → rgba(0,0,0,0.04)
├── Animation: 1.5s infinite, ease-in-out
└── Block radius: 6px

Block sizes match actual content layout for smooth transition.
```

#### Card — Error State

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │                  ⚠️                                 │   │  Icon: 32px
│   │                                                     │   │
│   │   Something went wrong                              │   │  H4 style
│   │   We couldn't load this profile.                    │   │  Body style
│   │                                                     │   │
│   │              [ Try Again ]                          │   │  Secondary button
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Error card styling:
├── Background: var(--surface-card)
├── Border: 1px solid var(--feedback-warning) at 30% opacity
├── Icon: Warning, in var(--feedback-warning)
├── Text: Normal ink colors (not red)
└── CTA: Secondary style

⚠️ Avoid red backgrounds or alarming language.
   Errors should feel recoverable, not catastrophic.
```

#### Card — Connection Lost (Terminated)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │   ┌─────┐                                           │   │
│   │   │ 👋  │  This connection has ended               │   │  Muted styling
│   │   └─────┘                                           │   │
│   │                                                     │   │
│   │   You were on Day 6 together.                       │   │  Caption
│   │   Keep going — your next connection is waiting.     │   │
│   │                                                     │   │
│   │              [ Dismiss ]                            │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Terminated card:
├── Background: var(--status-reset-bg)
├── Text: var(--unora-ink-tertiary)
├── No server accent bar
└── Encouraging, not mournful tone
```

---

### 8.3 Streak State Variants (Expanded)

#### State: Active — Mutual Check-In Complete

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Day 8 of 15                                               │
│   ═════════════════════════════░░░░░░░░░░░                  │  Server color bar
│                                                             │
│   ┌─────────────────────┐  ┌─────────────────────┐          │
│   │                     │  │                     │          │
│   │   ✓ You             │  │   ✓ Partner         │          │
│   │   Today at 9:42 AM  │  │   Today at 2:15 PM  │          │  Both green
│   │                     │  │                     │          │
│   └─────────────────────┘  └─────────────────────┘          │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │  ✦ Connection deepening...                            │   │  Trust status
│   │     Something is forming...                            │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Both backgrounds: var(--status-active-bg)
Both borders: 1px solid var(--status-active)
Checkmarks: Filled, var(--status-active)
```

#### State: At Risk — Partner Missing

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Day 7                                                     │
│   ═════════════════════════════░░░░░░░░░░░                  │  Still server color
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │   ⚠️  Your streak is at risk                        │   │  Amber alert
│   │                                                     │   │
│   │   You checked in. Your partner hasn't yet today.    │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────┐  ┌─────────────────────┐          │
│   │                     │  │                     │          │
│   │   ✓ You             │  │   ⏳ Partner        │          │
│   │   Today at 9:42 AM  │  │   Waiting...        │          │  You=green
│   │                     │  │                     │          │  Partner=amber
│   └─────────────────────┘  └─────────────────────┘          │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │               🔔 Nudge Your Partner                 │   │  Nudge CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Alert card: var(--status-at-risk-bg), border var(--status-at-risk)
Partner box: dashed border, var(--status-at-risk)
Nudge button: Secondary style with bell icon
```

#### State: Payment Window (Day N+1)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Day 7 of 15    ⚠️ Streak needs recovery                   │
│   ═══════════════════════════░░░░░░░░░░░░░                  │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │   You missed yesterday's check-in.                  │   │
│   │                                                     │   │
│   │   Your partner showed up. Would you like            │   │
│   │   to continue this streak?                          │   │
│   │                                                     │   │
│   │   ─────────────────────────────────────────         │   │
│   │                                                     │   │
│   │   ┌─────────────────────────────────────────────┐   │   │
│   │   │         Continue Streak — ₹99              │   │   │  Primary CTA
│   │   └─────────────────────────────────────────────┘   │   │
│   │                                                     │   │
│   │   If this connection ends within 24h, this          │   │  Microcopy
│   │   amount will convert to credits.                   │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                  [ Let it reset instead ]                   │  Tertiary link
│                                                             │
└─────────────────────────────────────────────────────────────┘

Payment card: var(--status-payment-bg)
Border: var(--status-payment)
Primary CTA: Full server color (not error red)
Reset option: Tertiary text style, no pressure
```

---

### 8.4 Progress Indicators

#### Linear Progress Bar (Streak)

```
Day Progress Bar:

         Day 1        Day 7         Day 15
           │            │              │
           ▼            ▼              ▼
┌──────────════════════════════░░░░░░░░░──────────┐
│  ●       ●       ◐       ○       ○       ⭐     │
│  R1      R2      R3      R4      R5      Chat   │
└─────────────────────────────────────────────────┘

Bar height: 6px
Filled color: var(--server-*-primary)
Unfilled: var(--border-subtle)
Milestone dots: 12px
Current milestone: Subtle pulse animation (scale 1.0 → 1.1)

● = Earned (filled, server color)
◐ = In progress (half-filled)
○ = Locked (outline only)
⭐ = Day 15 Chat (special icon)
```

#### Circular Progress (Cooldown Timer)

```
┌───────────────────────────────────────┐
│                                       │
│          ┌───────────┐                │
│         ╱     4h     ╲               │  
│        │     23m      │              │  Text: H2 style
│         ╲            ╱               │
│          └───────────┘                │
│                                       │
│        until next refresh             │  Caption
│                                       │
└───────────────────────────────────────┘

Ring: 4px stroke
Progress: Server color (clockwise fill)
Background ring: var(--border-subtle)
Size: 96px diameter
Animation: Smooth decrement (not ticking)
```

---

## 9. Onboarding Flow Visuals

### 9.1 Onboarding Philosophy

Onboarding sets the tone. It should feel:

- **Welcoming**: Warm, not clinical
- **Educational**: Explain the "why" behind friction
- **Quick**: 5-7 screens max before verification gate

### 9.2 Welcome Screen

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│                                                             │
│                       [ Unora Logo ]                        │
│                                                             │
│                                                             │
│            Connection worth waiting for.                    │  Tagline: H2
│                                                             │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                   Get Started                       │   │  Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                   Already have an account?                  │  Caption + link
│                        Sign in →                            │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Background: Subtle gradient (brand sand → white)
Logo: Centered, 64px
Tagline: Outfit, H2 weight
CTA: Brand primary (not server-specific yet)
```

### 9.3 Philosophy Introduction (Carousel)

3 screens explaining core concepts:

```
Screen 1: "No photos. Just hobbies."
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                      [ Illustration:                        │
│                        Anonymous cards                      │
│                        with hobby icons ]                   │
│                                                             │
│                                                             │
│           Discover people through                           │
│           what they do, not how they look.                  │
│                                                             │
│                                                             │
│                   ● ○ ○                                     │  Pagination dots
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                     Next                            │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Screen 2: "15 days of showing up."
(Illustration: Calendar with streak visualization)
"Chat unlocks after 15 days of mutual consistency."

Screen 3: "Earned, not given."
(Illustration: Reveals unlocking progressively)
"Identity is revealed gradually as trust builds."
```

### 9.4 Verification Gate Screen

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                       🛡️                                    │  Shield icon, 48px
│                                                             │
│              Everyone here is verified.                     │  H2
│                                                             │
│   Unora uses government-backed verification                 │
│   to ensure every person you meet is real.                  │
│   You can add trust signals now or later from                │
│   Settings.                                                  │
│                                                             │
│   ─────────────────────────────────────────                 │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │        Continue with Google   [G]                   │   │  Trust Booster
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │        Continue with Apple    []                   │   │  Trust Booster
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                    Skip for now →                           │  Tertiary link
│                                                             │
└─────────────────────────────────────────────────────────────┘

Shield icon: var(--feedback-success)
Tone: Trust as earned benefit, not barrier
"Skip for now": Proceeds to next step, optional verification
```

### 9.5 Server Selection (First Time)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│               What brings you to Unora?                     │  H2
│                                                             │
│   You can switch between servers anytime.                   │  Caption
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │   🔥  Looking for a Partner                         │   │ 
│   │       Find someone for a meaningful relationship    │   │  Server card 1
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │   👋  Friend / Companion                            │   │
│   │       Connect with someone who gets you             │   │  Server card 2
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │   🎯  Accountability / Growth Buddy                 │   │
│   │       Build habits and goals with a partner         │   │  Server card 3
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Server cards:
├── Left border: 4px, server color
├── Background: var(--server-*-surface)
├── On tap: Scale 0.98, highlight border
├── Selected: Server-primary border all around
```

---

## 10. Error, Loading & Empty States

### 10.1 Loading States

#### Full Screen Loading

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                      [ Unora Logo ]                         │  Logo: 48px, subtle pulse
│                                                             │
│                                                             │
│                 Loading your connections...                 │  Caption, muted
│                                                             │
│                                                             │
│                                                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Background: var(--surface-background)
Logo animation: Opacity 0.6 → 1.0, 1s infinite
No spinner — logo pulse is more branded
```

#### Inline Loading (Button)

```
Button — Loading State:

┌─────────────────────────────────────────────┐
│                                             │
│           ◐  Connecting...                  │  Spinner replaces icon
│                                             │
└─────────────────────────────────────────────┘

Spinner: 16px, white (on primary button)
Text: Faded to 80% opacity
Button: Non-interactive during load
Animation: Spinner rotates 360° in 800ms, linear
```

#### Pull-to-Refresh (Where Applicable)

**Note:** Pull-to-refresh is disabled on Discovery (no infinite scroll). Only used on Streak list and Settings.

```
Pull indicator:

        ↓ Release to refresh

        ┌──────────┐
        │    ↻     │
        └──────────┘

Icon: ArrowsClockwise, 20px
Text: Caption style
Threshold: 80px pull distance
Animation: Icon rotates on release
```

---

### 10.2 Error States

#### Network Error — Full Screen

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│                                                             │
│                        📡                                   │  Icon: 48px
│                                                             │
│                                                             │
│              Couldn't connect to Unora                      │  H3
│                                                             │
│   Check your internet connection and try again.             │  Body, muted
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                    Try Again                        │   │  Primary
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Icon: WifiSlash or similar, var(--unora-ink-tertiary)
Language: Simple, no technical jargon
CTA: Primary button (retry is the main action)
```

#### Server Error — Inline Toast

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   [ Normal screen content ]                                 │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  ⚠️  Something went wrong. We're on it.   [Retry]   │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Toast position: Bottom of screen, above navigation
Background: var(--feedback-warning) at 10% + border
Icon: Warning, var(--feedback-warning)
Text: var(--unora-ink-primary)
Duration: Persistent until dismissed or retried
```

#### Validation Error (Inline)

```
Input field with error:

┌─────────────────────────────────────────────────────────────┐
│   Email                                                     │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  notanemail                                         │   │
│   └─────────────────────────────────────────────────────┘   │
│   ⚠️  Please enter a valid email address                    │
└─────────────────────────────────────────────────────────────┘

Border: 1.5px solid var(--feedback-error)
Error text: var(--feedback-error), Caption size
Icon: Warning, inline with text
Animation: Subtle shake on error (150ms, 2 oscillations)
```

---

### 10.3 Empty States

#### Empty Discovery (No Cards Yet)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Partner                                                   │
│                                                             │
│                                                             │
│                   [ Illustration:                           │
│                     Hourglass or                            │
│                     seeds being planted ]                   │
│                                                             │
│                                                             │
│              Ready to discover someone new?                 │  H3
│                                                             │
│   Tap Refresh to see profiles that match                    │
│   your interests and filters.                               │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │               ↻ Refresh Discovery                   │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Illustration: Custom, brand-aligned, 120px
Tone: Inviting, not empty
CTA: Primary refresh action
```

#### Empty Streaks (No Active Connections)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Your Streaks                                              │
│                                                             │
│                                                             │
│                   [ Illustration:                           │
│                     Two paths converging ]                  │
│                                                             │
│                                                             │
│              No active connections yet                      │  H3
│                                                             │
│   When you match with someone, your                         │
│   shared streak will appear here.                           │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │               Go to Discovery                       │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Empty states should never feel lonely or like failure.
Language emphasizes "yet" — potential, not absence.
```

#### Empty Inbox (After Chat Unlock, No Messages)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Chat with [Partner Name]                                  │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                   ┌─────────────┐                           │
│                   │     💬      │                           │
│                   └─────────────┘                           │
│                                                             │
│           You did it! 15 days together.                     │
│                                                             │
│           Say hello to [First Name].                        │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│ ┌─────────────────────────────────────────────────────────┐ │
│ │  Type your first message...                             │ │
│ └─────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘

Celebration moment — the culmination of 15 days.
Input field ready; gentle prompt to break the ice.
No "no messages" language — this is a fresh start.
```

---

## 11. Modal & Sheet Patterns

### 11.0 Card Detail Modal (Discovery)

When a user taps a **Discovery Teaser** in the vertical stack, a **Card Detail Modal** appears as a centered overlay. This modal contains the full **Detail Card** with all information needed to make a connection decision.

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   [ Discovery list dimmed behind overlay ]                  │
│                                                             │
│       ┌─────────────────────────────────────────────┐       │
│       │                                         ✕   │       │  Close button
│       │                                             │       │
│       │  ┌─────────────────────────────────────┐    │       │
│       │  │  🏋️ Gym                              │    │       │  ← Hobby Anchor 1
│       │  │  "Consistency over intensity"        │    │       │  ← Micro-copy
│       │  └─────────────────────────────────────┘    │       │
│       │                                             │       │
│       │  ┌─────────────────────────────────────┐    │       │
│       │  │  🎨 Painting                         │    │       │  ← Hobby Anchor 2
│       │  │  "Late nights, slow strokes"         │    │       │  ← Micro-copy
│       │  └─────────────────────────────────────┘    │       │
│       │                                             │       │
│       │  ─────────────────────────────────────────  │       │
│       │                                             │       │
│       │  ┌────────────┐                             │       │
│       │  │ ▓▓▓▓▓▓░░░ │  Highly Consistent          │       │  ← Consistency Band
│       │  └────────────┘                             │       │
│       │                                             │       │
│       │  "Seeking someone who values slow           │       │  ← Intent Statement
│       │   mornings and deep conversations."         │       │
│       │                                             │       │
│       │  ─────────────────────────────────────────  │       │
│       │                                             │       │
│       │  ┌─────────────────────────────────────┐    │       │
│       │  │              🤝 Connect             │    │       │  ← Primary CTA
│       │  └─────────────────────────────────────┘    │       │
│       │                                             │       │
│       └─────────────────────────────────────────────┘       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Card Detail Modal Specifications:**

| Property | Value |
|----------|-------|
| **Overlay** | `var(--surface-overlay)` — rgba(26,26,26,0.6) |
| **Modal background** | `var(--surface-elevated)` — #FFFFFF |
| **Modal radius** | 20px |
| **Modal shadow** | 0 8px 32px rgba(0,0,0,0.15) |
| **Modal width** | Min(90vw, 360px) |
| **Modal max-height** | 85vh (scrollable if content exceeds) |
| **Padding** | 24px |
| **Position** | Centered horizontally and vertically |

**Close/X Button:**

```
┌────────┐
│   ✕    │   Position: Top-right corner of modal
└────────┘   Size: 44px × 44px touch target
             Icon: 20px, var(--unora-ink-secondary)
             Background: Transparent (hover: subtle highlight)
```

**Modal Content (Detail Card):**

1. **Close Button** — Top-right, allows dismissal
2. **Hobby Anchors** — Full display with icon, title, and micro-copy
3. **Consistency Band** — Full visual bar + label
4. **Intent Statement** — Complete text, not truncated
5. **Connect Button** — Primary CTA at bottom

**Dismiss Behavior:**

- **Close/X tap**: Closes modal, returns to Discovery list
- **Overlay tap**: Closes modal
- **Swipe down**: Closes modal (gesture-based dismiss)
- **Back button/gesture**: Closes modal

**Animation (Scale Up + Fade In):**

```
Open:
├── Overlay: Fade in, 0 → 1 opacity, 150ms
├── Modal: Scale 0.9 → 1.0, opacity 0 → 1, 250ms
└── Easing: ease-out

Close:
├── Modal: Scale 1.0 → 0.95, opacity 1 → 0, 200ms
├── Overlay: Fade out, 1 → 0 opacity, 150ms
└── Easing: ease-in
```

**After "Connect" Action:**

When the user taps "Connect":
1. Button enters loading state (spinner replaces icon)
2. On success: Modal closes automatically
3. Toast appears: "Interest sent"
4. Discovery list updates (teaser shows "Interest Sent" state)

---

### 11.1 Bottom Sheet (Standard)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   [ Content behind, dimmed overlay ]                        │
│                                                             │
├─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┤
│                         ───                                 │  Handle: 32px × 4px
│                                                             │
│   Filter by                                                 │  H3
│                                                             │
│   [ Filter content... ]                                     │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                  Apply Filters                      │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │  Safe area
└─────────────────────────────────────────────────────────────┘

Overlay: var(--surface-overlay)
Sheet background: var(--surface-elevated)
Corner radius: 20px (top only)
Handle: var(--border-medium), centered
Animation: Slide up from bottom, 250ms
Dismissible: Swipe down or tap overlay
```

### 11.2 Confirmation Modal (Destructive Action)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   [ Content behind, dimmed overlay ]                        │
│                                                             │
│       ┌─────────────────────────────────────────────┐       │
│       │                                             │       │
│       │          End this connection?               │       │  H3
│       │                                             │       │
│       │   This will permanently end your streak     │       │
│       │   with this person. You were on Day 6.      │       │
│       │                                             │       │
│       │   ─────────────────────────────────────     │       │
│       │                                             │       │
│       │   ┌─────────────────────────────────────┐   │       │
│       │   │       End Connection               │   │       │  Destructive
│       │   └─────────────────────────────────────┘   │       │
│       │                                             │       │
│       │              Keep Going                     │       │  Text link
│       │                                             │       │
│       └─────────────────────────────────────────────┘       │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Modal background: var(--surface-elevated)
Radius: 20px
Shadow: 0 8px 32px rgba(0,0,0,0.15)
Destructive button: Outline style, var(--feedback-error) border
Cancel: Tertiary text link (not a button)

⚠️ Destructive action is NEVER the most prominent button.
   "Keep Going" should be the visually dominant option.
```

### 11.3 Toast Notifications

```
Success Toast:
┌─────────────────────────────────────────────────────────────┐
│   ✓  Nudge sent                                        ✕    │
└─────────────────────────────────────────────────────────────┘

Info Toast:
┌─────────────────────────────────────────────────────────────┐
│   ℹ️  Your next refresh is in 4 hours                   ✕    │
└─────────────────────────────────────────────────────────────┘

Warning Toast:
┌─────────────────────────────────────────────────────────────┐
│   ⚠️  Check-in window closes at midnight               ✕    │
└─────────────────────────────────────────────────────────────┘

Toast Styling:
├── Position: Bottom center, 20px above navigation
├── Background: var(--surface-elevated)
├── Shadow: 0 4px 16px rgba(0,0,0,0.1)
├── Radius: 12px
├── Padding: 12px 16px
├── Icon: 16px, status color
├── Text: Body size
├── Dismiss: X icon, 16px, tap target 44px
├── Auto-dismiss: 4 seconds (except errors)
├── Animation: Fade + slide up 20px, 200ms
```

---

## 12. Illustration Guidelines

### 12.1 Illustration Style

Unora illustrations should be:

- **Abstract, not literal**: Represent concepts, not specific people
- **Warm and inviting**: Use brand color palette
- **Minimal detail**: Clean lines, limited elements
- **Gender-neutral**: Avoid specific human representations where possible

### 12.2 Illustration Color Palette

```css
/* Illustration palette (derived from brand) */
--illust-primary: #C4A77D;      /* Golden tan */
--illust-secondary: #E8DED5;    /* Warm sand */
--illust-accent: #D4714A;       /* Terracotta pop */
--illust-neutral: #7A7A7A;      /* Muted gray */
--illust-background: #FAF8F5;   /* Warm off-white */
```

### 12.3 Key Illustrations Needed

| Context | Concept | Visual Metaphor |
|---------|---------|-----------------|
| Onboarding 1 | Anonymous discovery | Floating cards with hobby icons |
| Onboarding 2 | 15-day journey | Path with milestones |
| Onboarding 3 | Gradual reveal | Layers peeling back |
| Empty Discovery | Potential | Seeds/sapling/hourglass |
| Empty Streaks | Connection waiting | Two paths converging |
| Match Created | Celebration | Gentle confetti/sparkle |
| Reveal Unlocked | Discovery | Eye opening/curtain lifting |
| Verification | Security | Shield/lock in friendly style |

### 12.4 Illustration Sizing

| Placement | Size | Notes |
|-----------|------|-------|
| Onboarding carousel | 200px height | Centered, full bleed optional |
| Empty states | 120px height | Centered above text |
| Inline cards | 64px | Contained within card |
| Celebration modals | 80px | Centered, animated optional |

---

## Appendix: Extended Token Reference

```css
/* Extended Copy-Paste Token Reference */

:root {
  /* ===== LIGHT MODE ===== */
  
  /* Brand Core */
  --unora-primary: #E8DED5;
  --unora-primary-accent: #C4A77D;
  --unora-primary-hover: #B08D5B;
  
  /* Ink */
  --unora-ink-primary: #1A1A1A;
  --unora-ink-secondary: #4A4A4A;
  --unora-ink-tertiary: #7A7A7A;
  --unora-ink-muted: #A3A3A3;
  
  /* Surfaces */
  --surface-background: #FAFAF8;
  --surface-card: #FFFFFF;
  --surface-elevated: #FFFFFF;
  --surface-overlay: rgba(26, 26, 26, 0.6);
  
  /* Borders */
  --border-subtle: #E8E8E6;
  --border-medium: #D4D4D0;
  --border-strong: #A3A3A3;
  
  /* Servers */
  --server-partner-primary: #D4714A;
  --server-partner-secondary: #F2E0D5;
  --server-partner-accent: #8B3A2F;
  --server-partner-surface: #FAF5F2;
  
  --server-friend-primary: #4A9B8C;
  --server-friend-secondary: #D9EDE8;
  --server-friend-accent: #2D6B5E;
  --server-friend-surface: #F5FAF8;
  
  --server-growth-primary: #5B6ABF;
  --server-growth-secondary: #E0E3F5;
  --server-growth-accent: #3D4785;
  --server-growth-surface: #F7F8FC;
  
  /* Status */
  --status-active: #4A9B8C;
  --status-active-bg: #E8F5F2;
  --status-at-risk: #E6A43A;
  --status-at-risk-bg: #FDF4E3;
  --status-payment: #D4714A;
  --status-payment-bg: #F9EAE3;
  --status-reset: #7A7A7A;
  --status-reset-bg: #F2F2F2;
  
  /* Feedback */
  --feedback-success: #4A9B8C;
  --feedback-warning: #E6A43A;
  --feedback-error: #C94A4A;
  --feedback-info: #5B6ABF;
  
  /* Typography */
  --font-display: 'Outfit', sans-serif;
  --font-body: 'Inter', sans-serif;
  
  /* Spacing */
  --space-unit: 4px;
  --padding-card: 16px;
  --padding-screen: 20px;
  
  /* Radius */
  --radius-sm: 8px;
  --radius-md: 12px;
  --radius-lg: 16px;
  --radius-xl: 24px;
  
  /* Shadows */
  --shadow-card: 0 2px 12px rgba(0,0,0,0.06);
  --shadow-elevated: 0 8px 32px rgba(0,0,0,0.12);
  --shadow-button: 0 2px 8px rgba(0,0,0,0.08);
  
  /* ===== DARK MODE ===== */
  
  --dm-surface-background: #121212;
  --dm-surface-card: #1E1E1E;
  --dm-surface-elevated: #252525;
  --dm-ink-primary: #F5F5F3;
  --dm-ink-secondary: #C4C4C0;
  --dm-ink-tertiary: #8A8A86;
  --dm-border-subtle: #2A2A2A;
  --dm-border-medium: #3D3D3D;
}
```

---

## 12. Premium Dark Visual Language

### 12.1 Philosophy

The premium dark mode is the **primary visual expression** of Unora. It should feel:

- **Premium**: Deep, layered surfaces with subtle luminosity
- **Alive**: Gentle glows and micro-animations create organic warmth
- **Intentional**: Every glow, shadow, and motion serves a purpose
- **Warm**: Logo-derived amber accents prevent cold, sterile appearance

> [!IMPORTANT]
> Dark mode is the default. Light mode exists for accessibility but dark mode defines the brand.

---

### 12.2 Logo-Derived Accent Palette

Colors extracted from the Unora brand logo for accents, glows, and highlights:

```css
/* Logo-Derived Premium Accents */
--accent-gold: #C4A77D;           /* Primary accent — buttons, focus, glows */
--accent-gold-bright: #D4B88D;    /* Hover/highlight state */
--accent-gold-deep: #B08D5B;      /* Pressed/active state */
--accent-gold-glow: rgba(196, 167, 125, 0.25);  /* Outer glow */
--accent-gold-subtle: rgba(196, 167, 125, 0.08); /* Inner ambient */

/* Warm Sand Tints */
--accent-sand: #E8DED5;           /* Soft highlights */
--accent-sand-glow: rgba(232, 222, 213, 0.15);  /* Ambient warmth */
```

---

### 12.3 Premium Dark Surface System

```css
/* Premium Dark Surfaces — Layered Depth */
--pdm-background: #0D0D0F;        /* Deepest layer — screen bg */
--pdm-surface-1: #141416;         /* Base cards */
--pdm-surface-2: #1A1A1E;         /* Elevated cards */
--pdm-surface-3: #222226;         /* Modals, sheets */
--pdm-surface-4: #2A2A2F;         /* Highest elevation (rare) */

/* Premium Dark Text */
--pdm-ink-primary: #F5F5F3;       /* Headlines — high contrast */
--pdm-ink-secondary: #C4C4C0;     /* Body text */
--pdm-ink-tertiary: #8A8A86;      /* Captions */
--pdm-ink-muted: #5A5A58;         /* Placeholders, disabled */

/* Premium Dark Borders */
--pdm-border-subtle: #2A2A2E;     /* Card edges */
--pdm-border-medium: #3D3D42;     /* Input borders */
--pdm-border-glow: rgba(196, 167, 125, 0.3); /* Focus ring */
```

---

### 12.4 Depth & Shadow System

Layered shadows create premium depth without flatness:

```css
/* Card Shadows — Multi-Layer */
--pdm-shadow-card: 
  0 2px 8px rgba(0, 0, 0, 0.3),
  0 8px 24px rgba(0, 0, 0, 0.2);

/* Elevated Shadows — Modals/Sheets */
--pdm-shadow-elevated: 
  0 4px 16px rgba(0, 0, 0, 0.4),
  0 16px 48px rgba(0, 0, 0, 0.3);

/* Button Shadows */
--pdm-shadow-button: 
  0 2px 8px rgba(0, 0, 0, 0.25),
  0 0 20px var(--accent-gold-glow);

/* Inner Ambient Glow (Cards) */
--pdm-inner-glow: inset 0 1px 0 rgba(255, 255, 255, 0.03);
```

**Glassmorphism for Overlays:**

```css
.modal-backdrop {
  background: rgba(13, 13, 15, 0.85);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
}
```

---

### 12.5 Glow & Highlight Treatments

#### Primary CTA Button (Premium)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │   │ ← Subtle gradient
│   │ ░░░░░░░░░░░░░░  🤝 Connect  ░░░░░░░░░░░░░░░░░░░░░░ │   │   (gold-deep → gold)
│   │ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │   │
│   └─────────────────────────────────────────────────────┘   │
│                    ╲ ░ ░ glow ░ ░ ╱                         │ ← Outer glow halo
│                                                             │
│   Background: linear-gradient(135deg,                       │
│               var(--accent-gold-deep) 0%,                   │
│               var(--accent-gold) 100%)                      │
│   Shadow: 0 0 24px rgba(196, 167, 125, 0.35)                │
│   Border: 1px solid rgba(255, 255, 255, 0.1)                │
│   Text: #FFFFFF                                             │
│                                                             │
│   Hover: Glow intensity +20%, subtle lift                   │
│   Pressed: Scale 0.97, glow dims slightly                   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

#### Focus State Ring

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐   │
│   ╎ ┌─────────────────────────────────────────────────┐ ╎   │
│   ╎ │                                                 │ ╎   │ ← 2px gold ring
│   ╎ │            [ Input Field ]                      │ ╎   │   with soft glow
│   ╎ │                                                 │ ╎   │
│   ╎ └─────────────────────────────────────────────────┘ ╎   │
│   └ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┘   │
│                                                             │
│   Focus ring: 2px solid var(--accent-gold)                  │
│   Ring glow: 0 0 12px var(--accent-gold-glow)               │
│   Animation: Fade in 150ms ease-out                         │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

#### Server Accent Glows (Active State)

```css
/* Server-colored outer glow on active elements */
--server-partner-glow: 0 0 16px rgba(212, 113, 74, 0.3);
--server-friend-glow: 0 0 16px rgba(74, 155, 140, 0.3);
--server-growth-glow: 0 0 16px rgba(91, 106, 191, 0.3);
```

---

### 12.6 Micro-Interaction Animation Library

All animations respect `prefers-reduced-motion`. Durations and easings are calibrated for premium feel.

#### Button Interactions

```css
/* Primary Button Tap */
@keyframes button-tap {
  0%   { transform: scale(1.0); }
  40%  { transform: scale(0.97); }
  100% { transform: scale(1.0); }
}
Duration: 180ms
Easing: cubic-bezier(0.34, 1.56, 0.64, 1) /* Elastic overshoot */

/* Glow Pulse on Release */
@keyframes button-glow-pulse {
  0%   { box-shadow: 0 0 20px rgba(196, 167, 125, 0.25); }
  50%  { box-shadow: 0 0 32px rgba(196, 167, 125, 0.45); }
  100% { box-shadow: 0 0 20px rgba(196, 167, 125, 0.25); }
}
Duration: 400ms
Trigger: On tap release
```

#### Card Interactions

```css
/* Card Lift Effect */
.card:active {
  transform: translateY(-2px) scale(0.99);
  box-shadow: 
    0 4px 16px rgba(0, 0, 0, 0.35),
    0 12px 32px rgba(0, 0, 0, 0.25);
}
Duration: 120ms
Easing: ease-out

/* Card Hover (Desktop) */
.card:hover {
  transform: translateY(-1px);
  box-shadow: var(--pdm-shadow-elevated);
  border-color: var(--pdm-border-medium);
}
Duration: 200ms
```

#### Modal Transitions

```css
/* Modal Open */
@keyframes modal-enter {
  0%   { opacity: 0; transform: scale(0.95); }
  100% { opacity: 1; transform: scale(1.0); }
}
Duration: 250ms
Easing: cubic-bezier(0.16, 1, 0.3, 1) /* Smooth deceleration */

/* Modal Close */
@keyframes modal-exit {
  0%   { opacity: 1; transform: scale(1.0); }
  100% { opacity: 0; transform: scale(0.97); }
}
Duration: 180ms
Easing: ease-in

/* Backdrop Blur Animation */
@keyframes backdrop-blur-in {
  0%   { backdrop-filter: blur(0px); opacity: 0; }
  100% { backdrop-filter: blur(16px); opacity: 1; }
}
Duration: 250ms
```

#### Tab & Navigation

```css
/* Tab Switch Indicator */
.tab-indicator {
  transition: transform 280ms cubic-bezier(0.4, 0.0, 0.2, 1);
  /* Spring-like physics feel */
}

/* Tab Icon Transition */
.tab-icon {
  transition: color 180ms ease, transform 180ms ease;
}
.tab-icon.active {
  transform: scale(1.05);
}
```

#### Progress & Loading

```css
/* Premium Shimmer (Skeleton Loading) */
@keyframes premium-shimmer {
  0%   { background-position: -200% 0; }
  100% { background-position: 200% 0; }
}
Background: linear-gradient(
  90deg,
  var(--pdm-surface-1) 0%,
  rgba(196, 167, 125, 0.08) 50%,  /* Gold-tinted highlight */
  var(--pdm-surface-1) 100%
);
Background-size: 200% 100%;
Duration: 1.8s infinite

/* Progress Bar Fill */
@keyframes progress-fill {
  0%   { width: var(--progress-start); }
  100% { width: var(--progress-end); }
}
Duration: 400ms
Easing: ease-out

/* Milestone Unlock Pulse */
@keyframes milestone-pulse {
  0%   { transform: scale(1.0); box-shadow: 0 0 0 rgba(196, 167, 125, 0); }
  50%  { transform: scale(1.15); box-shadow: 0 0 16px rgba(196, 167, 125, 0.5); }
  100% { transform: scale(1.0); box-shadow: 0 0 0 rgba(196, 167, 125, 0); }
}
Duration: 600ms
Trigger: On milestone unlock
```

#### Success & Celebration

```css
/* Checkmark Draw Animation */
@keyframes checkmark-draw {
  0%   { stroke-dashoffset: 24; }
  100% { stroke-dashoffset: 0; }
}
Duration: 350ms
Easing: ease-out
Overshoot: Scale 1.0 → 1.1 → 1.0

/* Success Glow Burst */
@keyframes success-glow-burst {
  0%   { opacity: 1; transform: scale(0.8); }
  40%  { opacity: 0.8; transform: scale(1.3); }
  100% { opacity: 0; transform: scale(1.5); }
}
Duration: 500ms
Element: Circular glow ring behind checkmark

/* Reveal Unlock Animation */
@keyframes reveal-unlock {
  0%   { opacity: 0; transform: scale(0.9); filter: blur(4px); }
  60%  { opacity: 1; transform: scale(1.02); filter: blur(0); }
  100% { opacity: 1; transform: scale(1.0); filter: blur(0); }
}
Duration: 600ms
Accompaniment: Golden ring expands outward
```

---

### 12.7 Reduced Motion Fallbacks

```css
@media (prefers-reduced-motion: reduce) {
  *, *::before, *::after {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
  
  /* Preserve essential visual feedback */
  .button:active { transform: none; }
  .card:active { transform: none; }
  .modal { animation: none; opacity: 1; transform: none; }
}
```

---

### 12.8 Premium Dark Card Example

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   Background: var(--pdm-background) #0D0D0F                 │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ ▌                                                   │   │ Card: #1A1A1E
│   │ ▌  ┌─────────────────────────────────────────────┐  │   │ Border: 1px #2A2A2E
│   │ ▌  │  🏋️ Gym                                     │  │   │ Left accent: server color
│   │ ▌  │  "Consistency over intensity"               │  │   │   with subtle glow
│   │ ▌  └─────────────────────────────────────────────┘  │   │
│   │ ▌                                                   │   │ Text: #F5F5F3 headline
│   │ ▌     ●●●●●●░░░░  Highly Consistent               │   │       #8A8A86 caption
│   │ ▌                                                   │   │
│   │ ▌  ┌─────────────────────────────────────────────┐  │   │ CTA: Gold gradient
│   │ ▌  │ ░░░░░░░░░░  🤝 Connect  ░░░░░░░░░░░░░░░░░░ │  │   │      + outer glow
│   │ ▌  └─────────────────────────────────────────────┘  │   │
│   │ ▌                                                   │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   Shadow: 0 2px 8px rgba(0,0,0,0.3),                        │
│           0 8px 24px rgba(0,0,0,0.2)                        │
│   Inner glow: inset 0 1px 0 rgba(255,255,255,0.03)          │
│   Server bar glow: 0 0 8px rgba(server-color, 0.25)         │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

**Document maintained by:** Design Systems Team  
**Last updated:** January 2026  
**Review cycle:** Monthly  
**Version:** 1.3

