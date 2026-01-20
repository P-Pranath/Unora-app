# Unora UI Specification — Locked Discovery State

**Document ID:** Spec-10  
**Screen Name:** Locked Discovery State  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Locked Discovery State** — Empty state replacing the card stack when browsing is restricted

### 1.2 User Flow Reference
**Phase 2 (Discovery Loop) → Double Lock System** — This state appears when either the Time Lock or Capacity Lock is triggered.

**Sequence:**
```
Discovery Feed → Refresh → [Lock Check] → Locked Discovery State
                                ↓
                    Lock 1: Time Lock (Cooldown)
                    Lock 2: Capacity Lock (Slots Full)
```

**Reference:** 
- [Unora_UserFlow_Logic.md — Section 2.B.2](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)
- [Unora_PRD.md — Section 12.5](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)

### 1.3 Purpose
Enforce intentionality by limiting infinite browsing. The lock system ensures users focus on existing connections rather than endless swiping, reinforcing the core Unora philosophy of **Presence over Performance**.

### 1.4 The Double Lock System

> **CRITICAL:** Both locks are **GLOBAL** — a lock in Partner server also blocks Friend and Growth servers.

| Lock Type | Trigger | Resolution |
|-----------|---------|------------|
| **Time Lock** | User has used their refresh | Wait for cooldown to expire |
| **Capacity Lock** | Active connection slots are full | Complete/end a streak, or upgrade |

---

## 2. Low-Fidelity Wireframes (ASCII)

### 2.1 Variant A: Time Lock (Scarcity)

```
┌─────────────────────────────────────────────────────────────┐
│   Partner Discovery              🔍 Filter    ⏱️ 11:42     │  ← Timer in header
│   ─────────────────────────────────────────────────────────│
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                        ┌───────────┐                        │
│                       ╱             ╲                       │
│                      │   ⏱️         │                       │  ← Timer icon (48px)
│                      │    11:42     │                       │     inside ring
│                       ╲             ╱                       │
│                        └───────────┘                        │
│                    ████████░░░░░░░░░░░░                     │  ← Progress ring
│                                                             │
│                                                             │
│              Your next refresh is brewing.                  │  ← Headline
│                                                             │
│              New connections available in                   │  ← Subtext
│                      11 hours 42 minutes                    │
│                                                             │
│                                                             │
│              ┌─────────────────────────────────┐            │
│              │        View Your Streaks        │            │  ← Primary CTA
│              └─────────────────────────────────┘            │
│                                                             │
│                    Switch servers? →                        │  ← Tertiary (disabled)
│                    (Also locked)                            │     Shows global scope
│                                                             │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│    🔥 Streaks      🃏 Discovery      👤 Profile             │  ← 3-Tab Bottom Nav
│                    ━━━━━━━━━━━                              │     Discovery active
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Variant B: Capacity Lock (Focus)

```
┌─────────────────────────────────────────────────────────────┐
│   Partner Discovery              🔍 Filter    🔒  1/1      │  ← Lock in header
│   ─────────────────────────────────────────────────────────│
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                                                             │
│                           🎯                                │  ← Target icon (48px)
│                                                             │
│                                                             │
│                    You are focused.                         │  ← Headline
│                                                             │
│              You have 1 active connection.                  │  ← Subtext
│              Maintain it, complete it, or upgrade           │
│              to browse more.                                │
│                                                             │
│                                                             │
│              ┌───────────────────────────────────────┐      │
│              │   ●●●●●●●●○○○○○○○     Day 8 of 15    │      │  ← Streak preview
│              └───────────────────────────────────────┘      │
│                                                             │
│              ┌─────────────────────────────────┐            │
│              │       View Your Streak          │            │  ← Primary CTA
│              └─────────────────────────────────┘            │
│                                                             │
│                      Upgrade Plan →                         │  ← Secondary
│                                                             │
│                    Switch servers? →                        │  ← Tertiary (disabled)
│                    (Also locked)                            │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│    🔥 Streaks      🃏 Discovery      👤 Profile             │  ← 3-Tab Bottom Nav
│                    ━━━━━━━━━━━                              │     Discovery active
└─────────────────────────────────────────────────────────────┘
```

### 2.3 Global Lock Indicator

```
When user taps another server tab:

┌─────────────────────────────────────────────────────────────┐
│                                                             │
│             🔒 Discovery is locked                          │
│                                                             │
│     This applies to all servers until                       │
│     your cooldown ends or a slot opens.                     │
│                                                             │
│              [Got it]                                       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
LOCKED DISCOVERY CONTAINER
├── Position: Replaces Discovery card stack
├── Width: 100%
├── Height: Full content area (minus header/nav)
├── Background: var(--surface-background)
├── Display: flex, column
├── Align: center
├── Justify: center
│
├── [ICON AREA]
│   ├── Size: 120px container
│   ├── Margin-bottom: var(--space-6) → 24px
│   └── Contains: Icon + progress ring (Time Lock)
│
├── [TEXT AREA]
│   ├── Text-align: center
│   ├── Max-width: 280px
│   └── Margin-bottom: var(--space-8) → 40px
│
├── [STREAK PREVIEW] (Capacity Lock only)
│   ├── Width: 280px
│   ├── Margin-bottom: var(--space-6) → 24px
│   └── Shows current streak status
│
└── [ACTION AREA]
    ├── Width: 280px
    ├── Gap: var(--space-3) → 12px
    └── Contains: Primary + Secondary CTAs

Premium Dark Mode (Default):
├── Background: var(--pdm-background) → #0D0D0F
├── Progress ring: Gold accent with soft glow
├── Focus icon: Gold with warm ambient glow
└── Streak preview: Elevated surface with subtle inner glow
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Progress ring track** | Border-subtle `#2A2A2E` |
| **Progress ring fill** | Gold gradient with glow: `0 0 16px rgba(196,167,125,0.3)` |
| **Timer icon** | Gold `#C4A77D` with ambient glow |
| **Focus icon** | Gold with warm glow: `0 0 24px rgba(196,167,125,0.25)` |
| **Streak preview card** | Surface `#1A1A1E`, border `#2A2A2E`, inner glow |
| **Progress dots** | Server color, completed dots have subtle glow |
| **Primary CTA** | Gold gradient + outer glow (DSD Section 12.5) |

**Calming Glow Animation (Time Lock):**
```css
/* Gentle breathing glow for timer icon */
@keyframes calm-breathe {
  0%, 100% { filter: drop-shadow(0 0 16px rgba(196, 167, 125, 0.2)); }
  50%      { filter: drop-shadow(0 0 24px rgba(196, 167, 125, 0.35)); }
}
Duration: 3s infinite
/* Slower animation = calmer feeling */
```



### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Icon size | — | 48px icon, 120px container |
| Icon margin-bottom | `var(--space-6)` | 24px |
| Headline margin-bottom | `var(--space-2)` | 8px |
| Subtext margin-bottom | `var(--space-8)` | 40px |
| Button gap | `var(--space-3)` | 12px |
| Content max-width | — | 280px |

### 3.3 Z-Index

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Background | 0 | Screen bg |
| Content | 1 | Lock state content |
| Header | 10 | Fixed header |
| Bottom Nav | 10 | Tab bar |
| Toast/Modal | 50 | Global lock message |

---

## 4. Component Inventory

### 4.1 Visual Tone

> **Important:** Use **calm, neutral colors** — NOT error/warning colors. This is intentional friction, not punishment.

| Element | Color |
|---------|-------|
| Icon | `var(--unora-primary-accent)` → Warm gold |
| Progress ring | `var(--unora-primary-accent)` |
| Text | `var(--unora-ink-primary)` / `--unora-ink-secondary` |
| Background | `var(--surface-background)` |

### 4.2 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Headline | Outfit | 600 | 22px | `--unora-ink-primary` |
| Subtext | Inter | 400 | 16px | `--unora-ink-secondary` |
| Timer display | Outfit | 600 | 32px | `--unora-primary-accent` |
| Link text | Inter | 500 | 14px | `--unora-ink-tertiary` |

### 4.3 Time Lock Components

#### Progress Ring
```
PROGRESS RING
├── Size: 120px diameter
├── Stroke width: 6px
├── Track: var(--border-subtle)
├── Fill: var(--unora-primary-accent)
├── Progress: Based on time remaining / total cooldown
├── Animation: Smooth decrease
│
└── [CENTER CONTENT]
    ├── Timer icon: 32px
    └── Countdown: "11:42" (HH:MM format)
```

#### Timer Icon
| Property | Value |
|----------|-------|
| Icon | Hourglass / Timer |
| Size | 48px |
| Color | `var(--unora-primary-accent)` |
| Style | Outlined, friendly |

### 4.4 Capacity Lock Components

#### Target/Focus Icon
| Property | Value |
|----------|-------|
| Icon | Target / Bullseye / Focus |
| Size | 48px |
| Color | `var(--unora-primary-accent)` |
| Style | Filled, warm |

#### Slot Indicator
```
SLOT INDICATOR (Header)
├── Format: "1/1 Active" or "2/2 Active"
├── Icon: 🔒 lock prefix
├── Color: var(--unora-ink-tertiary)
└── Position: Replaces Refresh button
```

#### Streak Preview Card
```
STREAK PREVIEW
├── Width: 280px
├── Height: 60px
├── Background: var(--surface-card)
├── Border: 1px solid var(--border-subtle)
├── Border radius: var(--radius-md)
├── Padding: 12px 16px
│
├── [PROGRESS DOTS]
│   ├── 15 dots representing streak days
│   ├── Filled: var(--unora-primary-accent)
│   └── Empty: var(--border-subtle)
│
└── [LABEL]
    └── "Day 8 of 15" — Inter 14px/500
```

### 4.5 Action Buttons

#### Primary CTA
| Lock Type | Label |
|-----------|-------|
| Time Lock | "View Your Streaks" |
| Capacity Lock | "View Your Streak" |

| Property | Value |
|----------|-------|
| Height | 52px |
| Width | 280px |
| Background | `var(--unora-primary-accent)` |
| Text | White, Inter 16px/600 |
| Border radius | `var(--radius-md)` |

#### Secondary CTA (Capacity Lock)
| Property | Value |
|----------|-------|
| Label | "Upgrade Plan" |
| Style | Tertiary link |
| Color | `var(--unora-primary-accent)` |

#### Tertiary Link (Global Indicator)
| Property | Value |
|----------|-------|
| Label | "Switch servers?" |
| State | Disabled/dimmed |
| Subtext | "(Also locked)" |
| Color | `var(--unora-ink-muted)` |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Condition | Result |
|---------|-----------|--------|
| Refresh tap | Slots at capacity | Show Capacity Lock |
| Refresh tap | Within cooldown | Show Time Lock |
| Tab switch | Any lock active | Show same lock state (global) |
| Timer expires | Cooldown complete | Return to Discovery |

### 5.2 Lock Evaluation Logic

```
USER taps "Refresh" OR opens Discovery
    │
    ▼
┌─────────────────────────────────────────┐
│ [CHECK 1] Capacity Lock                 │
│ Are Active Connections = Tier Limit?    │
└─────────────────────────────────────────┘
    │                    │
    ▼                    ▼
  YES                   NO
    │                    │
    ▼                    ▼
┌─────────────────┐  ┌─────────────────────────────────────┐
│ **CAPACITY      │  │ [CHECK 2] Time Lock                 │
│ LOCK**          │  │ Is Cooldown Timer Active?           │
│                 │  └─────────────────────────────────────┘
│ Show Variant B  │      │                    │
│ (Focus state)   │      ▼                    ▼
│                 │    YES                   NO
│ GLOBAL:         │      │                    │
│ All servers     │      ▼                    ▼
│ locked          │  ┌─────────────────┐  ┌─────────────────┐
└─────────────────┘  │ **TIME LOCK**   │  │ **UNLOCKED**    │
                     │                 │  │                 │
                     │ Show Variant A  │  │ Show 5 new      │
                     │ (Timer state)   │  │ Discovery cards │
                     │                 │  │                 │
                     │ GLOBAL:         │  │                 │
                     │ All servers     │  │                 │
                     │ locked          │  │                 │
                     └─────────────────┘  └─────────────────┘
```

### 5.3 Cooldown Durations by Tier

| Tier | Cooldown | Capacity |
|------|----------|----------|
| Free | 24 hours | 1 slot |
| Plus | 12 hours | 2 slots |
| Pro | 6 hours | 4 slots |

### 5.4 Global Lock Behavior

When user taps a different server tab while locked:

```
USER taps "Friend" tab (while Partner is locked)
    │
    ▼
┌─────────────────────────────────────────┐
│ Show same lock state                    │
│ Update header: "Friend Discovery"       │
│ Lock message remains                    │
│                                         │
│ Toast (first time):                     │
│ "Discovery is locked across all         │
│  servers until your cooldown ends       │
│  or a slot opens."                      │
└─────────────────────────────────────────┘
```

### 5.5 Unlock Transitions

| Event | Transition |
|-------|------------|
| Timer expires | Auto-refresh, show new cards with fade-in |
| Streak completes | Slot freed, show toast "Slot opened! Refresh to browse" |
| Upgrade | Immediately unlock if new capacity > active |

---

## 6. State Definitions

### 6.1 State Matrix

| State | Lock Type | Visual | Resolution |
|-------|-----------|--------|------------|
| Time Locked | Scarcity | Timer + ring | Wait for cooldown |
| Capacity Locked | Focus | Target + streak | Complete streak or upgrade |
| Unlocked | — | Normal Discovery | — |

### 6.2 Time Lock State

```
Icon: Timer with progress ring
Headline: "Your next refresh is brewing."
Subtext: "New connections available in X hours Y minutes"
Primary CTA: "View Your Streaks"
Tertiary: "Switch servers? (Also locked)"
```

### 6.3 Capacity Lock State

```
Icon: Target / Focus
Headline: "You are focused."
Subtext: "You have X active connection(s). Maintain it, complete it, 
         or upgrade to browse more."
Streak preview: Visual dots showing current progress
Primary CTA: "View Your Streak"
Secondary: "Upgrade Plan →"
Tertiary: "Switch servers? (Also locked)"
```

---

## 7. Content & Copy Guidelines

### 7.1 Time Lock Copy

| Element | Copy |
|---------|------|
| Headline | "Your next refresh is brewing." |
| Subtext | "New connections available in [X hours Y minutes]" |
| Primary CTA | "View Your Streaks" |
| Tertiary | "Switch servers? (Also locked)" |

### 7.2 Capacity Lock Copy (By Tier)

**Free Tier (1 slot):**
| Element | Copy |
|---------|------|
| Headline | "You are focused." |
| Subtext | "You have 1 active connection. Maintain it, complete it, or upgrade to browse more." |
| Primary CTA | "View Your Streak" |
| Secondary | "Upgrade Plan →" |

**Plus Tier (2 slots):**
| Element | Copy |
|---------|------|
| Headline | "You are focused." |
| Subtext | "You have 2 active connections. Complete one or upgrade to browse more." |

**Pro Tier (4 slots):**
| Element | Copy |
|---------|------|
| Headline | "You are at capacity." |
| Subtext | "All 4 connection slots are active. Complete a streak to browse again." |

### 7.3 Global Lock Toast

| Trigger | Message |
|---------|---------|
| Tab switch while locked | "Discovery is locked across all servers" |
| First lock | "Focus on your current connections. New cards soon." |

### 7.4 Tone Guidelines

| Principle | Application |
|-----------|-------------|
| Calm, not punitive | "You are focused" not "Limit reached" |
| Warm | "Brewing" not "Waiting" |
| Encouraging | Redirect to streaks, not dead-end |
| Honest | Show clear resolution paths |

---

## 8. Accessibility

### 8.1 Screen Reader
- Time Lock: "Discovery paused. Next refresh in 11 hours 42 minutes. Button: View your streaks."
- Capacity Lock: "You are focused. 1 of 1 connection slots active. Day 8 of 15. Button: View your streak."
- Global: "This lock applies to all servers."

### 8.2 Touch Targets
- Primary button: 52px × 280px
- Secondary links: 44px height

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Time Lock variant (timer + ring) | Critical | ☐ |
| Capacity Lock variant (focus + streak) | Critical | ☐ |
| Progress ring animation | High | ☐ |
| Countdown timer (live update) | High | ☐ |
| Streak preview card | High | ☐ |
| Global lock behavior | Critical | ☐ |
| Tab switch handling | Critical | ☐ |
| Auto-unlock on timer expire | Critical | ☐ |
| Tier-specific copy | High | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 12.5 — Double Lock System |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.B.2 — Lock Logic |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 4.1 — Empty States |
| [Unora_Spec_07_DiscoveryFeed.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_07_DiscoveryFeed.md) | Parent screen |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
