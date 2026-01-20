# Unora UI Specification — Discovery Feed

**Document ID:** Spec-07  
**Screen Name:** Discovery Feed  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Discovery Feed** — Main home screen for browsing daily connections

### 1.2 User Flow Reference
**Phase 2: Active Discovery** — This is the primary screen where users review their daily batch of 5 potential connections.

**Sequence:**
```
Server Selection → [Discovery Feed] ←→ Detail Modal
                        ↓
                 Active Streak → Reveal Journey
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.C](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Present the daily batch of 5 potential connections in a compact, scannable format. Users can view details, express interest, or refresh for new cards (subject to tier cooldowns and capacity locks).

### 1.4 Navigation Architecture

> [!IMPORTANT]
> **This screen is Contextual (Server-Specific).** Shows potential matches for the CURRENTLY SELECTED SERVER only. Server Switcher is VISIBLE on this screen.

**Reference:** [UserFlow Section 1.1 — Navigation Architecture](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

| Aspect | Behavior |
|--------|----------|
| **Tab Position** | Center (Tab 2) — Cards/Logo icon |
| **Tab Status** | Active/Highlighted when on this screen |
| **Server Switcher** | ✓ **VISIBLE** — Top-Left dropdown |
| **Theme** | Accent colors match currently selected server |
| **Bottom Navigation** | 3-Tab Floating Glass bar (Streaks, Discovery, Profile) |

```
HEADER STRUCTURE (Server Switcher PRESENT)
├── Top-Left: Server Switcher dropdown (🔥 Partner ▼)
├── Center: Screen title "Discovery"
├── Right: Filter + Refresh buttons
└── Subline: Filter summary

BOTTOM NAVIGATION (Tab 2 Active)
├── Streaks: 🔥 Inactive
├── Discovery: 🃏 Active (highlighted, golden accent)
└── Profile: 👤 Inactive
```

### 1.5 Primary Action
- **View Detail** — Tap a card to open the Detail Modal
- **Refresh** — Generate new batch (if available)

---

## 2. Low-Fidelity Wireframe (ASCII)

### 2.1 Main Discovery Feed

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   Partner Discovery              🔍 Filter    🔄 Refresh    │  ← Header
│   ─────────────────────────────────────────────────────────│
│   Showing: Age 25-30 • High Consistency                     │  ← Filter summary
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ █ │  🏋️ Gym   ●●●○○                                │   │  ← Card 1
│   └─────────────────────────────────────────────────────┘   │
│                                                             │  ← 8px gap
│   ┌─────────────────────────────────────────────────────┐   │
│   │ █ │  📚 Reading   ●●●●○                             │   │  ← Card 2
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ █ │  🧘 Meditation   ●●○○○                          │   │  ← Card 3
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ █ │  ✈️ Travel   ●●●●●                              │   │  ← Card 4
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │ █ │  🎸 Music   ●●●○○                               │   │  ← Card 5
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│    🔥 Streaks      🃏 Discovery      👤 Profile             │  ← 3-Tab Bottom Nav
│                    ━━━━━━━━━━━                              │     Discovery active
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Card Detail (Left Border Accent)

```
┌─────────────────────────────────────────────────────────────┐
│ █ │  🏋️  Gym, Reading                    ●●●○○             │
│ █ │       +2 more                                          │
└─────────────────────────────────────────────────────────────┘
  ↑
  Server-colored left border (4px)
  Partner = Terracotta
  Friend = Teal
  Growth = Indigo
```

### 2.3 Refresh Button States

```
ACTIVE:           COOLDOWN:           LOCKED:
┌────────┐        ┌────────────┐      ┌────────────┐
│   🔄   │        │  ⏱️ 11:42  │      │   🔒  1/1  │
└────────┘        └────────────┘      └────────────┘
 Enabled           Timer visible       Capacity full
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
DISCOVERY FEED CONTAINER
├── Position: fixed, 100vw × 100vh
├── Background: var(--surface-background) → #FAFAF8
├── Display: flex, column
│
├── [HEADER AREA] — 100px including safe area
│   ├── Padding: var(--padding-screen) → 20px horizontal
│   ├── Padding-top: env(safe-area-inset-top) + 16px
│   ├── Row 1: Title + Filter + Refresh
│   └── Row 2: Filter summary caption
│
├── [CARDS AREA] — flex: 1, scrollable
│   ├── Padding: 20px horizontal
│   ├── Gap: var(--space-2) → 8px (tight stack)
│   ├── Padding-top: 8px
│   └── Padding-bottom: 100px (space for bottom nav)
│
└── [BOTTOM NAV] — 80px + safe area
    ├── Position: fixed, bottom: 0
    ├── Width: 100%
    ├── Height: 80px + env(safe-area-inset-bottom)
    └── Background: var(--surface-card) with subtle top border

Premium Dark Mode (Default):
├── Background: var(--pdm-background) → #0D0D0F
├── Cards: Elevated surfaces with server-colored left border + glow
├── Bottom nav: Glass backdrop with inner glow highlight
└── Active tab: Server color with subtle icon glow
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Connection cards** | Surface `#1A1A1E`, border `#2A2A2E`, server color left bar with glow |
| **Card left border** | 4px with matching outer glow: `0 0 8px rgba(server-color, 0.25)` |
| **Consistency band** | Server color dots, active dots have subtle glow |
| **Refresh button** | Gold accent, glow pulse on tap |
| **Bottom nav** | Glass backdrop `blur(16px)`, inner highlight top edge |
| **Active tab** | Server color icon with glow: `0 0 8px rgba(server-color, 0.4)` |

**Card Stack Premium Styling:**
```css
/* Premium connection card */
.connection-card {
  background: var(--pdm-surface-2);
  border: 1px solid var(--pdm-border-subtle);
  box-shadow: var(--pdm-shadow-card),
              inset 0 1px 0 rgba(255, 255, 255, 0.03);
}

/* Server-colored left accent with glow */
.card-server-accent::before {
  width: 4px;
  background: var(--server-color);
  box-shadow: 0 0 8px rgba(var(--server-color-rgb), 0.3);
}
```



### 3.2 Discovery Card Dimensions

| Property | Value |
|----------|-------|
| Width | 100% (full container width - padding) |
| Height | 90px |
| Padding | 16px |
| Border radius | `var(--radius-md)` → 12px |
| Left border | 4px solid [server accent] |
| Touch target | Full card area |

### 3.3 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen padding | `var(--padding-screen)` | 20px |
| Card gap | `var(--space-2)` | 8px |
| Header title margin-bottom | `var(--space-1)` | 4px |
| Filter summary margin-bottom | `var(--space-3)` | 12px |
| Bottom nav icon gap | `var(--space-8)` | 40px |

### 3.4 Z-Index Layers

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Background | 0 | Screen bg |
| Cards | 1 | Card list |
| Header | 10 | Fixed header |
| Bottom Nav | 10 | Tab bar |
| Modal | 50 | Detail modal overlay |
| System | 100+ | Status bar |

---

## 4. Component Inventory

### 4.1 Dynamic Server Theming

**All accent colors change based on active server:**

| Server | Token | Hex | Usage |
|--------|-------|-----|-------|
| **Partner** | `var(--server-partner-primary)` | #C9785D | Left border, icons, active tab |
| **Friend** | `var(--server-friend-primary)` | #4A9B8C | Left border, icons, active tab |
| **Growth** | `var(--server-growth-primary)` | #6B5B95 | Left border, icons, active tab |

**Affected Elements:**
- Card left border accent
- Header title icon
- Filter/Refresh button icons
- Bottom nav active tab
- Consistency dots (filled)

### 4.2 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Header title | Outfit | 600 | 20px | `--unora-ink-primary` |
| Filter summary | Inter | 400 | 12px | `--unora-ink-tertiary` |
| Card hobby text | Inter | 500 | 16px | `--unora-ink-primary` |
| Card secondary | Inter | 400 | 14px | `--unora-ink-secondary` |
| Tab label | Inter | 500 | 12px | `--unora-ink-tertiary` / active: server color |

### 4.3 Discovery Teaser Card (Variant A)

**Reference:** [Unora_DesignSystem.md — Section 10.2](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

```
DISCOVERY TEASER CARD
├── Width: 100%
├── Height: 90px
├── Background: var(--surface-card) → #FFFFFF
├── Border: 1px solid var(--border-subtle)
├── Border-left: 4px solid [server accent color]
├── Border radius: var(--radius-md) → 12px
├── Padding: 16px
├── Shadow: 0 2px 4px rgba(0,0,0,0.04)
│
├── [LEFT] Hobby Icon
│   ├── Size: 40px container, 24px icon
│   ├── Background: [server accent] @ 10%
│   └── Border radius: var(--radius-sm) → 8px
│
├── [CENTER] Text Content
│   ├── Primary: Hobby anchors (comma separated)
│   ├── Secondary: "+X more" if > 2 hobbies
│   └── Align: left
│
└── [RIGHT] Consistency Indicator
    ├── 5 dots, 8px each
    ├── Gap: 4px
    ├── Filled: [server accent color]
    └── Empty: var(--border-subtle)
```

### 4.4 Header Actions

#### Filter Button
| Property | Value |
|----------|-------|
| Icon | Filter/funnel, 20px |
| Color | `var(--unora-ink-secondary)` |
| Touch target | 44px × 44px |
| Action | Open Filter Sheet |

#### Refresh Button (3 States)

**State: Active (Available)**
| Property | Value |
|----------|-------|
| Icon | Refresh/rotate, 20px |
| Color | [Server accent color] |
| Touch target | 44px × 44px |
| Action | Generate new batch |

**State: Cooldown (Timer Active)**
| Property | Value |
|----------|-------|
| Display | ⏱️ + "HH:MM" countdown |
| Color | `var(--unora-ink-tertiary)` |
| Touch target | 44px × 44px |
| Action | Show toast "Refresh available in X" |

**State: Locked (Capacity Full)**
| Property | Value |
|----------|-------|
| Display | 🔒 + "X/X" capacity indicator |
| Color | `var(--unora-ink-muted)` |
| Touch target | 44px × 44px |
| Action | **Open Global Lock Modal** (critical upsell/focus moment) |

### 4.5 Bottom Navigation Bar (3-Tab Model)

**Reference:** [DesignSystem Section 4.2 — Bottom Navigation](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

```
3-TAB BOTTOM NAVIGATION (Floating Glass)
├── Height: 64px + margin (floating)
├── Margin: 20px bottom, 16px sides
├── Background: rgba(13, 13, 15, 0.85) + blur(20px)
├── Border: 1px solid rgba(255, 255, 255, 0.08)
├── Border radius: 20px
│
├── [TAB 1] Streaks
│   ├── Icon: 🔥 Flame (24px)
│   ├── Label: "Streaks" (11px)
│   └── State: Inactive (tertiary color, outlined)
│
├── [TAB 2] Discovery ← ACTIVE ON THIS SCREEN
│   ├── Icon: 🃏 Cards / Unora Logo (24px)
│   ├── Label: "Discovery" (11px)
│   ├── State: Active (golden accent, filled, indicator bar)
│   └── Indicator: 3px × 24px bar below icon
│
└── [TAB 3] Profile
    ├── Icon: 👤 Avatar (24px)
    ├── Label: "Profile" (11px)
    └── State: Inactive (tertiary color, outlined)
```

### 4.6 Consistency Dots

| Property | Value |
|----------|-------|
| Count | 5 dots (representing Days 1-5 milestone) |
| Size | 8px diameter each |
| Gap | 4px |
| Filled color | [Server accent color] |
| Empty color | `var(--border-subtle)` → #E8E8E6 |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Discovery Card | Open Detail Modal (Screen #08) |
| Tap | Filter Button | Open Filter Bottom Sheet |
| Tap | Refresh Button | Execute refresh logic |
| Tap | Bottom Nav Tab | Switch server context |
| Scroll | Card list | Standard scroll behavior |

### 5.2 Card Tap Behavior

```
USER taps a Discovery Card
    │
    ▼
┌─────────────────────────────────────────┐
│ Card: Scale 1.0 → 0.98 → 1.0 (100ms)    │
│ Haptic: Light impact                     │
└─────────────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────────────┐
│ Open Detail Modal (Screen #08)          │
│ Animation: Card expands to modal        │
│ Duration: 300ms ease-out                │
└─────────────────────────────────────────┘
```

### 5.3 Refresh Logic

```
USER taps Refresh Button
    │
    ▼
┌─────────────────────────────────────────┐
│ [CHECK] Are Active Connection Slots     │
│         at Tier Limit?                  │
└─────────────────────────────────────────┘
    │                              │
    ▼                              ▼
  YES (Locked)                   NO (Open)
    │                              │
    ▼                              ▼
┌───────────────────────────┐  ┌─────────────────────────────────────┐
│ **GLOBAL LOCK MODAL**     │  │ [STEP] Evaluate and apply           │
│                           │  │        PENDING filters              │
│ Open modal with CTAs:     │  └─────────────────────────────────────┘
│                           │      │
│  ┌─────────────────────┐  │      ▼
│  │    [View Streak]    │  │  ┌─────────────────────────────────────┐
│  └─────────────────────┘  │  │ [CHECK] Is Cooldown Timer Active?   │
│                           │  └─────────────────────────────────────┘
│  ┌─────────────────────┐  │      │                    │
│  │      [Upgrade]      │  │      ▼                    ▼
│  └─────────────────────┘  │    YES                   NO
│                           │      │                    │
│ Button shows lock icon +  │      ▼                    ▼
│ capacity (e.g., 1/1)      │  ┌─────────────────┐  ┌─────────────────┐
└───────────────────────────┘  │ **COOLDOWN**    │  │ **REFRESH**     │
                               │                 │  │                 │
                               │ Show toast:     │  │ Generate 5 new  │
                               │ "Refresh in X"  │  │ cards           │
                               │                 │  │                 │
                               │ Button shows    │  │ Haptic: Success │
                               │ timer countdown │  │                 │
                               └─────────────────┘  │ Animation:      │
                                                    │ Cards fade out  │
                                                    │ New cards slide │
                                                    │ in from bottom  │
                                                    └─────────────────┘
```

**Cooldown Times by Tier:**
| Tier | Cooldown |
|------|----------|
| Free | 24 hours |
| Plus | 12 hours |
| Pro | 6 hours |

### 5.4 Server Switch Behavior

```
USER taps Bottom Nav Tab (different server)
    │
    ▼
┌─────────────────────────────────────────┐
│ 1. Update active tab visual             │
│ 2. Apply server color theme             │
│ 3. Load cards for that server           │
│ 4. Update header title                  │
└─────────────────────────────────────────┘
```

### 5.5 Transitions

| Transition | Animation |
|------------|-----------|
| Card tap | Scale 0.98, 100ms |
| Modal open | Expand from card position, 300ms |
| Refresh cards | Fade out old, slide up new, 400ms |
| Tab switch | Cross-fade, 200ms |
| Header color change | Color transition, 200ms |

---

## 6. State Definitions

### 6.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Default | 5 cards visible | Normal browsing |
| Empty | Empty state message | No cards available |
| Refresh Available | Refresh button active | Slots + cooldown OK |
| Cooldown Active | Timer on refresh button | Within cooldown period |
| Capacity Locked | Lock icon on refresh | Active slots at limit |
| Loading | Skeleton cards | Fetching new batch |

### 6.2 Default State (5 Cards)

```
Header: "Partner Discovery" (or Friend/Growth)
Filter summary: "Showing: [filters]"
Cards: 5 teaser cards visible
Refresh: Based on cooldown/capacity
Bottom Nav: Active server tab highlighted
```

### 6.3 Empty State (Initial Load Only)

> [!IMPORTANT]
> **Static Batch Behavior:** Discovery cards form a **Static Batch** and do **not** disappear after viewing. Once generated, all 5 cards remain visible until the user performs a refresh action.
>
> This empty state should **only appear on a fresh server load** before the first-ever refresh—i.e., when the server has never generated a batch for this user.

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                                                             │
│                         📭                                  │
│                                                             │
│           Ready to discover someone new?                    │  ← Per Design System 10.3
│                                                             │
│                   [ 🔄 Refresh ]                            │  ← Primary CTA
│                                                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 6.4 Loading State

```
5 skeleton cards (shimmer animation)
├── Left border: gray
├── Content: gray blocks
└── Animation: shimmer left-to-right, 1.5s loop
```

---

## 7. Content & Copy Guidelines

### 7.1 Header Copy

| Server | Title |
|--------|-------|
| Partner | "Partner Discovery" |
| Friend | "Friend Discovery" |
| Growth | "Growth Discovery" |

### 7.2 Filter Summary Examples

| Filter Applied | Display |
|----------------|---------|
| Age range | "Age 25-30" |
| Consistency | "High Consistency" |
| Multiple | "Age 25-30 • High Consistency" |
| None | "Showing all" |

### 7.3 Card Content

| Element | Content |
|---------|---------|
| Primary hobbies | "Gym, Reading" (max 2 shown) |
| Overflow | "+3 more" |
| Consistency | ●●●○○ (visual only) |

### 7.4 Toast Messages

| State | Message |
|-------|---------|
| Cooldown | "Refresh available in X hours" |
| Locked | "Complete a connection to unlock refresh" |
| Refresh success | "5 new connections ready" |

### 7.5 Empty State Copy

| Context | Message |
|---------|---------|
| Cooldown active | "No cards left. Refresh available in [time]." |
| Capacity locked | "No cards left. Complete a connection to see more." |

---

## 8. Accessibility

### 8.1 Screen Reader
- Header: "Partner Discovery. Filter button. Refresh button."
- Card: "Card 1 of 5. Hobbies: Gym, Reading. Consistency 3 of 5. Tap to view details."
- Refresh states: "Refresh available" / "Refresh in 11 hours" / "Refresh locked"

### 8.2 Touch Targets
- Cards: Full 90px height
- Header buttons: 44px × 44px
- Bottom nav tabs: 60px × 60px minimum

### 8.3 Color Contrast
- All text meets WCAG AA
- Consistency dots have sufficient contrast

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| 5 discovery cards in vertical stack | Critical | ☐ |
| Dynamic server theming | Critical | ☐ |
| Card tap → Detail Modal | Critical | ☐ |
| Refresh button 3 states | Critical | ☐ |
| Cooldown timer logic | Critical | ☐ |
| Capacity lock logic | Critical | ☐ |
| Bottom navigation | Critical | ☐ |
| Server switch | High | ☐ |
| Filter button | High | ☐ |
| Empty state | High | ☐ |
| Loading skeleton | Medium | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 12 — Discovery Logic |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 10 — Card Variants, Server Tokens |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.C — Discovery Flow |
| [Unora_Spec_06_ServerSelect.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_06_ServerSelect.md) | Previous screen |
| Unora_Spec_08_DetailModal.md (planned) | Card detail overlay |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
