# Unora UI Specification — Streaks Dashboard

**Document ID:** Spec-11  
**Screen Name:** Streaks Dashboard  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Streaks Dashboard** — Default home screen for existing users

### 1.2 User Flow Reference
**Phase 4 (The Streak Loop)** — This is where users manage their daily commitment and view progress across all active connections.

**Sequence:**
```
Login → [Streaks Dashboard] ←→ Streak Detail
              ↓
         Discovery (if slots available)
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.D](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Display all active connections with their streak progress, daily status, and upcoming reveals. Enable quick actions for at-risk or recovery states.

### 1.4 Navigation Architecture

> [!IMPORTANT]
> **This screen is a Global Aggregator.** The Streaks tab shows ALL active connections from ALL servers in one unified list. Server Switcher is HIDDEN on this screen.

**Reference:** [UserFlow Section 1.1 — Navigation Architecture](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

| Aspect | Behavior |
|--------|----------|
| **Tab Position** | Left (Tab 1) — Flame icon |
| **Tab Status** | Active/Highlighted when on this screen |
| **Server Switcher** | ✗ **HIDDEN** — filtering by server is unnecessary |
| **Connection List** | Aggregated from ALL servers (mixed Terracotta/Teal/Indigo borders) |
| **Bottom Navigation** | 3-Tab Floating Glass bar (Streaks, Discovery, Profile) |

```
HEADER STRUCTURE (Server Switcher ABSENT)
├── Left: "My Streaks" title
├── Right: Current date
└── NO server switcher dropdown (global view)

BOTTOM NAVIGATION (Tab 1 Active)
├── Streaks: 🔥 Active (highlighted, golden accent)
├── Discovery: 🃏 Inactive
└── Profile: 👤 Inactive
```

### 1.5 Tier Constraints

| Tier | Max Active Streaks |
|------|--------------------|
| Free | 1 |
| Plus | 2 |
| Pro | 4 |

---

## 2. Low-Fidelity Wireframes (ASCII)

### 2.1 Streaks Dashboard (Active State)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   My Streaks                               Monday, Jan 6    │  ← Header
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  ⚠️  Check-in window closes in 4 hours              │   │  ← Urgent banner
│   └─────────────────────────────────────────────────────┘   │     (if applicable)
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ █ │  👤 A****                          Day 7        │   │  ← Streak Card 1
│   │ █ │      Partner is waiting for you… ···────────    │   │     (Partner server)
│   │ █ │      ✦ Connection strengthening    [Check In]   │   │     At Risk state
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ █ │  👤 S****                          Day 12       │   │  ← Streak Card 2
│   │ █ │      Both checked in ✓              ────────    │   │     (Friend server)
│   │ █ │      ✦ Trust building...                        │   │     Active state
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ █ │  👤 P****                          Day 5        │   │  ← Streak Card 3
│   │ █ │      You missed yesterday ⚠️        ···─────    │   │     (Growth server)
│   │ █ │      🎯 Streak at risk!            [Recover ₹]  │   │     Payment state
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│    📊 Streaks      🔍 Discover       👤 Profile             │  ← Bottom Nav
│    ━━━━━━━━━━                                              │     Streaks active
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Streak Card Detail

```
┌─────────────────────────────────────────────────────────────┐
│ █ │  ┌────────┐                                             │
│ █ │  │  👤    │   A****                        Day 7       │  ← Avatar + Name
│ █ │  │ (blur) │   Partner is waiting...                    │     Big day counter
│ █ │  └────────┘                                             │
│ █ │                                                         │
│ █ │  ●●●●●●●○○○○○○○○                                       │  ← Progress dots
│ █ │                                                         │
│ █ │  ✦ Trust building... (pulse)           [Check In]     │  ← Status + Action
│   │                                                         │
└─────────────────────────────────────────────────────────────┘
  ↑
  Server-colored left border (4px)
  Partner = Terracotta, Friend = Teal, Growth = Indigo
```

### 2.3 Empty State

```
┌─────────────────────────────────────────────────────────────┐
│   My Streaks                               Monday, Jan 6    │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                           💫                                │  ← Illustration
│                                                             │
│                                                             │
│             No active connections yet.                      │  ← Headline
│                                                             │
│       Your streaks will appear here once you                │  ← Subtext
│       connect with someone in Discovery.                    │
│                                                             │
│                                                             │
│              ┌─────────────────────────────┐                │
│              │      Start Discovery        │                │  ← Primary CTA
│              └─────────────────────────────┘                │
│                                                             │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│    📊 Streaks      🔍 Discover       👤 Profile             │
│    ━━━━━━━━━━                                              │
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
STREAKS DASHBOARD CONTAINER
├── Position: fixed, 100vw × 100vh
├── Background: var(--surface-background)
├── Display: flex, column
│
├── [HEADER AREA] — 80px
│   ├── Padding: 20px horizontal
│   ├── Padding-top: env(safe-area-inset-top) + 16px
│   ├── Left: "My Streaks" title
│   └── Right: Current date
│
├── [URGENT BANNER] — 48px (conditional)
│   ├── Background: var(--feedback-warning) @ 10%
│   ├── Margin: 0 20px 16px
│   └── Only shown if check-in pending < 6 hours
│
├── [CARDS AREA] — flex: 1, scrollable
│   ├── Padding: 0 20px
│   ├── Gap: var(--space-3) → 12px
│   └── Padding-bottom: 100px
│
└── [BOTTOM NAV] — 80px + safe area
    └── Standard tab bar

Premium Dark Mode (Default):
├── Background: var(--pdm-background) → #0D0D0F
├── Streak cards: Elevated surfaces with server-colored left border + glow
├── Day counter: Large, gold-tinted with subtle glow
└── Progress dots: Filled dots glow, empty dots are subtle charcoal
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Streak cards** | Surface `#1A1A1E`, border `#2A2A2E`, server-colored left bar with glow |
| **Card left border** | 4px server color with glow: `0 0 8px rgba(server,0.3)` |
| **Day counter** | Gold-tinted text with subtle glow on milestone days (5, 10, 15) |
| **Progress dots (filled)** | Server color with subtle glow |
| **Progress dots (empty)** | Charcoal `#2A2A2E` |
| **Avatar (blurred)** | Soft inner glow treatment |
| **Check-In button** | Server-colored gradient + glow |
| **Urgent banner** | Warning orange @ 15% with subtle border glow |

**Streak Day Milestone Glow:**
```css
/* Special glow for milestone days */
.day-counter.milestone {
  color: var(--accent-gold);
  text-shadow: 0 0 12px rgba(196, 167, 125, 0.4);
}

/* Filled progress dot with glow */
.progress-dot.filled {
  background: var(--server-color);
  box-shadow: 0 0 4px rgba(var(--server-color-rgb), 0.4);
}
```



### 3.2 Streak Card Dimensions

| Property | Value |
|----------|-------|
| Width | 100% (container - padding) |
| Height | 120px |
| Padding | 16px |
| Border radius | `var(--radius-lg)` → 16px |
| Left border | 4px solid [server color] |

### 3.3 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen padding | `var(--padding-screen)` | 20px |
| Card gap | `var(--space-3)` | 12px |
| Avatar size | — | 48px |
| Day counter size | — | 40px text |
| Progress dots | — | 15 dots, 6px each |

### 3.4 Z-Index Layers

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Background | 0 | Screen bg |
| Cards | 1 | Streak list |
| Urgent Banner | 5 | Pinned notification |
| Header | 10 | Title bar |
| Bottom Nav | 10 | Tab bar |

---

## 4. Component Inventory

### 4.1 Status Colors

**Reference:** [Unora_DesignSystem.md — Section 6](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

| Status | Color Token | Hex | Usage |
|--------|-------------|-----|-------|
| **Active (Healthy)** | `var(--feedback-success)` | #4A9B8C (Teal) | Check icon, healthy state |
| **At Risk (Warning)** | `var(--feedback-warning)` | #E6A43A (Amber) | Warning icon, nudge state |
| **Payment (Critical)** | `var(--feedback-error)` | #C9785D (Terracotta) | Alert, recovery needed |

### 4.2 Server Border Colors

| Server | Token | Hex |
|--------|-------|-----|
| Partner | `var(--server-partner-primary)` | #C9785D |
| Friend | `var(--server-friend-primary)` | #4A9B8C |
| Growth | `var(--server-growth-primary)` | #6B5B95 |

### 4.3 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Header title | Outfit | 600 | 24px | `--unora-ink-primary` |
| Date | Inter | 400 | 14px | `--unora-ink-tertiary` |
| Alias/Name | Outfit | 600 | 18px | `--unora-ink-primary` |
| Status text | Inter | 400 | 14px | `--unora-ink-secondary` |
| Day counter | Outfit | 700 | 32px | `--unora-primary-accent` |
| "Day X" label | Inter | 500 | 12px | `--unora-ink-tertiary` |
| Reveal countdown | Inter | 500 | 14px | `--unora-ink-secondary` |

### 4.4 Streak Card Component

```
STREAK CARD (Dashboard Variant)
├── Width: 100%
├── Height: 120px
├── Background: var(--surface-card)
├── Border: 1px solid var(--border-subtle)
├── Border-left: 4px solid [server color]
├── Border radius: var(--radius-lg) → 16px
├── Padding: 16px
├── Shadow: 0 2px 8px rgba(0,0,0,0.04)
│
├── [LEFT SECTION] — 60px
│   ├── Avatar: 48px, blurred/abstract
│   └── Name: "A****" (masked)
│
├── [CENTER SECTION] — flex: 1
│   ├── Status text: "Partner is waiting..."
│   ├── Progress dots: 15 dots
│   └── Connection status: "✦ Trust building..." (pulsing glow)
│
└── [RIGHT SECTION] — 80px
    ├── Day counter: "Day 7"
    └── Action button (if applicable)
```

### 4.5 Progress Dots

```
PROGRESS DOTS (15-day streak)
├── Count: 15 dots
├── Dot size: 6px
├── Gap: 4px
├── Filled: var(--unora-primary-accent)
├── Empty: var(--border-subtle)
└── Current day: Pulsing animation (optional)
```

### 4.6 Action Buttons

#### Check-In Button (At Risk)
| Property | Value |
|----------|-------|
| Label | "Check In" |
| Background | `var(--feedback-warning)` |
| Text | White, Inter 14px/600 |
| Height | 36px |
| Border radius | `var(--radius-md)` |

#### Recover Button (Payment)
| Property | Value |
|----------|-------|
| Label | "Recover ₹49" |
| Background | `var(--feedback-error)` |
| Text | White, Inter 14px/600 |
| Height | 36px |
| Border radius | `var(--radius-md)` |

### 4.7 Urgent Banner

```
URGENT CHECK-IN BANNER
├── Background: var(--feedback-warning) @ 10%
├── Border: 1px solid var(--feedback-warning) @ 30%
├── Border radius: var(--radius-md)
├── Padding: 12px 16px
├── Icon: ⚠️ warning, 16px
├── Text: "Check-in window closes in X hours"
└── Visible: When pending check-in && < 6 hours remaining
```

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Streak Card | Navigate to Streak Detail (Screen #12) |
| Tap | "Check In" button | Quick check-in, update state |
| Tap | "Recover ₹" button | Open payment flow |
| Tap | "Nudge" (if visible) | Send nudge to partner |
| Tap | "Start Discovery" | Navigate to Discovery Feed |

### 5.2 Card States & Logic

```
FOR EACH ACTIVE CONNECTION:
    │
    ▼
┌─────────────────────────────────────────┐
│ [CHECK] Did User check in today?        │
└─────────────────────────────────────────┘
    │                    │
    ▼                    ▼
  YES                   NO
    │                    │
    ▼                    ▼
┌───────────────┐  ┌─────────────────────────────────────┐
│ [CHECK]       │  │ **PAYMENT STATE**                   │
│ Did Partner   │  │                                     │
│ check in?     │  │ Status: "You missed yesterday ⚠️"   │
│               │  │ Button: "Recover ₹49"               │
│               │  │ Color: Terracotta/Error             │
│               │  │ Streak at risk of breaking          │
└───────────────┘  └─────────────────────────────────────┘
    │         │
    ▼         ▼
  YES        NO
    │         │
    ▼         ▼
┌─────────────────┐  ┌─────────────────────────────────┐
│ **ACTIVE**      │  │ **AT RISK**                     │
│                 │  │                                 │
│ Status: "Both   │  │ Status: "Partner is waiting…"  │
│ checked in ✓"   │  │ Button: "Check In" or "Nudge"  │
│ No action       │  │ Color: Amber/Warning           │
│ needed          │  │                                 │
│ Color: Teal     │  │                                 │
└─────────────────┘  └─────────────────────────────────┘
```

### 5.3 Check-In Window

| Parameter | Value |
|-----------|-------|
| Window opens | 6:00 AM local time |
| Window closes | 11:59 PM local time |
| Grace period | Until 6:00 AM next day |
| After grace | Streak broken OR payment required |

### 5.4 Transitions

| Transition | Animation |
|------------|-----------|
| Card tap | Scale 0.98, 100ms |
| Navigate to detail | Shared element (card expands) |
| Status update | Fade + subtle scale, 200ms |
| Urgent banner appear | Slide down, 250ms |

---

## 6. State Definitions

### 6.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Active (Healthy) | Green check, no action | Both users checked in |
| At Risk | Amber warning, "Check In" button | Partner waiting |
| Payment Required | Red alert, "Recover" button | User missed check-in |
| Empty | Illustration + CTA | No active streaks |
| Loading | Skeleton cards | Fetching data |

### 6.2 Active Card State

```
Border: Server color
Status icon: ✓ green checkmark
Status text: "Both checked in"
Day counter: Normal color
Action: None (tap to view detail)
```

### 6.3 At Risk Card State

```
Border: Server color
Status icon: ⏳ or ⚠️ amber
Status text: "Partner is waiting for you..."
Day counter: Amber accent
Action: "Check In" button (primary warning)
```

### 6.4 Payment Required State

```
Border: Terracotta/Error
Status icon: ⚠️ red
Status text: "You missed yesterday"
Day counter: Red accent
Action: "Recover ₹49" button
Additional: "Streak at risk!" warning text
```

---

## 7. Content & Copy Guidelines

### 7.1 Header Copy

| Element | Copy |
|---------|------|
| Title | "My Streaks" |
| Date | "Monday, Jan 6" (dynamic) |

### 7.2 Card Status Text

| State | Status Text |
|-------|-------------|
| Active | "Both checked in ✓" |
| User Checked, Waiting | "Waiting for partner..." |
| At Risk | "Partner is waiting for you..." |
| Payment | "You missed yesterday ⚠️" |

### 7.3 Connection Status (Abstract — No Countdowns)

> [!IMPORTANT]
> **Mystery Reveal UX:** Users never see day countdowns. They see abstract status indicators that pulse/glow with streak progress.

| Server | Status Text |
|--------|-------------|
| Partner | "✦ Trust building..." |
| Friend | "✦ Connection deepening..." |
| Growth | "✦ Progress forming..." |

### 7.4 Urgent Banner

| Time Remaining | Copy |
|----------------|------|
| 6 hours | "Check-in window closes in 6 hours" |
| 2 hours | "⚠️ Only 2 hours left to check in!" |
| 1 hour | "🚨 Last hour to save your streak!" |

### 7.5 Empty State Copy

| Element | Copy |
|---------|------|
| Headline | "No active connections yet." |
| Subtext | "Your streaks will appear here once you connect with someone in Discovery." |
| CTA | "Start Discovery" |

### 7.6 Button Labels

| State | Label |
|-------|-------|
| At Risk | "Check In" |
| Payment | "Recover ₹49" |
| Nudge available | "Nudge" |

---

## 8. Accessibility

### 8.1 Screen Reader
- Header: "My Streaks. Monday, January 6."
- Card: "Streak with A, 4 letters hidden. Day 7. Status: Partner is waiting. Trust building. Button: Check in."
- Urgent: "Alert: Check-in window closes in 4 hours."

### 8.2 Touch Targets
- Cards: Full card area (120px height)
- Action buttons: 44px minimum height
- Bottom nav: 60px × 60px per tab

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Streak card list | Critical | ☐ |
| Three card states (Active/At Risk/Payment) | Critical | ☐ |
| Server-colored left borders | High | ☐ |
| Progress dots (15-day) | High | ☐ |
| Urgent banner logic | High | ☐ |
| Empty state | High | ☐ |
| Quick action buttons | Critical | ☐ |
| Tap to detail navigation | Critical | ☐ |
| Loading skeleton | Medium | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 14 — Streak System |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 6 — Status Colors |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.D — Streak Loop |
| Unora_Spec_12_StreakDetail.md (planned) | Detail view |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
