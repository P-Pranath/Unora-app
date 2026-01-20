# Unora UI Specification — Filter Sheet

**Document ID:** Spec-09  
**Screen Name:** Filter Sheet  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Filter Sheet** — Bottom sheet overlay for refining matches

### 1.2 User Flow Reference
**Phase 2 (Discovery) → Filter Adjustment** — Accessible from the Discovery Feed header.

**Sequence:**
```
Discovery Feed → [Filter Sheet] → Apply → Back to Discovery
                                    ↓
                          (User manually taps Refresh to apply)
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.C](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Refine the pool of potential connections based on **compatibility signals** (consistency, intent, lifestyle) rather than just demographics. Filters are server-specific to match different relationship intents.

### 1.4 Primary Action
- **Apply Filters** — Save filter state and close sheet
- **Reset** — Clear all non-mandatory filters

---

## 2. Low-Fidelity Wireframe (ASCII)

### 2.1 Filter Sheet Structure

```
┌─────────────────────────────────────────────────────────────┐
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│  ← Discovery Feed
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│     (dimmed)
│▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒│
├─────────────────────────────────────────────────────────────┤
│                         ───────                             │  ← Drag handle
│                                                             │
│   Filter Partner Matches                              ✕     │  ← Header + Close
│                                                             │
│   ─────────────────────────────────────────────────────     │
│                                                             │
│   BASICS                                                    │  ← Section header
│                                                             │
│   Age Range                                                 │
│   ○─────────────●═══════●─────────○                         │  ← Dual slider
│   18            25      32        45                        │
│                                                             │
│   Gender                                                    │
│   ┌────────┐  ┌────────┐  ┌────────┐  ┌──────────────┐      │
│   │ Women  │  │  Men   │  │  All   │  │ Non-binary   │      │  ← Chips
│   └────────┘  └────────┘  └────────┘  └──────────────┘      │
│                                                             │
│   ─────────────────────────────────────────────────────     │
│                                                             │
│   COMPATIBILITY                                             │  ← Section header
│                                                             │
│   Consistency Level                                         │
│   ┌────────┐  ┌──────────┐  ┌────────┐                      │
│   │  Calm  │  │ Regular  │  │  High  │                      │  ← Chips
│   └────────┘  └──────────┘  └────────┘                      │
│                                                             │
│   Availability                                              │
│   ┌───────────┐  ┌───────────┐  ┌───────────┐               │
│   │  Morning  │  │  Evening  │  │  Anytime  │               │  ← Chips
│   └───────────┘  └───────────┘  └───────────┘               │
│                                                             │
│   ─────────────────────────────────────────────────────     │
│                                                             │
│   RELATIONSHIP                                              │  ← Partner-specific
│                                                             │
│   Relationship Intent                                       │
│   ┌──────────────┐  ┌────────────────┐                      │
│   │  Long-term   │  │  Open to both  │                      │
│   └──────────────┘  └────────────────┘                      │
│                                                             │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   Reset                              [  Apply Filters  ]    │  ← Fixed footer
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Chip States

```
DEFAULT:                    SELECTED:
┌────────────┐              ┏━━━━━━━━━━━━┓
│   Calm     │              ┃ ✓  Calm    ┃   ← Server accent bg
└────────────┘              ┗━━━━━━━━━━━━━┛
 Gray border                 Filled + checkmark
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
FILTER SHEET CONTAINER
├── Type: Bottom Sheet
├── Position: fixed, bottom: 0
├── Width: 100%
├── Height: dynamic, max 85vh
├── Background: var(--surface-card) → #FFFFFF
├── Border-top-left-radius: var(--radius-xl) → 20px
├── Border-top-right-radius: var(--radius-xl) → 20px
├── Z-Index: 50
│
├── [DRAG HANDLE]
│   ├── Width: 40px, Height: 4px
│   ├── Background: var(--border-medium)
│   ├── Border radius: var(--radius-full)
│   ├── Margin: 12px auto
│   └── Purpose: Visual affordance for drag-to-dismiss
│
├── [HEADER AREA]
│   ├── Padding: 0 24px 16px
│   ├── Display: flex, space-between
│   └── Sticky: Yes (doesn't scroll)
│
├── [SCROLL AREA] — flex: 1
│   ├── Padding: 0 24px
│   ├── Overflow-y: auto
│   └── Padding-bottom: 100px (space for footer)
│
└── [FOOTER AREA] — fixed bottom
    ├── Position: sticky, bottom: 0
    ├── Padding: 16px 24px
    ├── Padding-bottom: env(safe-area-inset-bottom) + 16px
    └── Background: var(--surface-card)

Premium Dark Mode (Default):
├── Sheet surface: var(--pdm-surface-3) → #222226
├── Backdrop: Glassmorphism blur with 85% opacity
├── Selected chips: Server accent with outer glow
└── Slider thumb: Gold accent with glow
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Sheet surface** | Elevated charcoal `#222226` with top border highlight |
| **Backdrop** | Glass blur: `blur(16px)`, bg `rgba(13,13,15,0.85)` |
| **Drag handle** | Lighter charcoal `#3D3D42` |
| **Section headers** | Gold-tinted text `rgba(196,167,125,0.7)` |
| **Chip (default)** | Border `#2A2A2E`, surface `#1A1A1E` |
| **Chip (selected)** | Server accent bg + glow: `0 0 8px rgba(server,0.3)` |
| **Slider track** | Charcoal `#2A2A2E`, fill in server color with glow |
| **Slider thumb** | White with server-colored ring + subtle glow |
| **Apply button** | Server gradient + outer glow |



### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Sheet padding | — | 24px horizontal |
| Section gap | `var(--space-6)` | 24px |
| Section header margin-bottom | `var(--space-3)` | 12px |
| Filter group margin-bottom | `var(--space-4)` | 16px |
| Chip gap | `var(--space-2)` | 8px |
| Chip row wrap gap | `var(--space-2)` | 8px |

### 3.3 Z-Index Stack

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Discovery Feed | 1 | Background content |
| Scrim | 49 | Dimmed overlay |
| Filter Sheet | 50 | Sheet content |
| Footer | 51 | Sticky buttons |

---

## 4. Component Inventory

### 4.1 Dynamic Server Theming

**Selected chip accent changes per active server:**

| Server | Token | Hex | Usage |
|--------|-------|-----|-------|
| **Partner** | `var(--server-partner-primary)` | #C9785D | Selected chip bg |
| **Friend** | `var(--server-friend-primary)` | #4A9B8C | Selected chip bg |
| **Growth** | `var(--server-growth-primary)` | #6B5B95 | Selected chip bg |

### 4.2 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Sheet title | Outfit | 600 | 18px | `--unora-ink-primary` |
| Section header | Inter | 600 | 12px | `--unora-ink-tertiary` (uppercase) |
| Filter label | Inter | 500 | 14px | `--unora-ink-secondary` |
| Chip text | Inter | 500 | 14px | `--unora-ink-primary` / White (selected) |
| Slider label | Inter | 400 | 12px | `--unora-ink-tertiary` |
| Button text | Inter | 600 | 16px | White / Accent |

### 4.3 Filter Chip (Multi-Select)

**Reference:** [Unora_DesignSystem.md — Section 3.4](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

```
FILTER CHIP
├── Height: 40px
├── Padding: 8px 16px
├── Border radius: var(--radius-lg) → 16px
├── Font: Inter 14px / 500
│
├── [DEFAULT STATE]
│   ├── Background: transparent
│   ├── Border: 1.5px solid var(--border-medium)
│   └── Text: var(--unora-ink-primary)
│
├── [SELECTED STATE]
│   ├── Background: [server accent color]
│   ├── Border: none
│   ├── Text: white
│   └── Icon: ✓ checkmark, 14px, left of text
│
└── [PRESSED STATE]
    ├── Scale: 0.97
    └── Duration: 100ms
```

### 4.4 Range Slider (Age)

```
DUAL RANGE SLIDER
├── Track height: 4px
├── Track background: var(--border-subtle)
├── Track fill: [server accent color]
├── Thumb size: 24px
├── Thumb background: white
├── Thumb border: 2px solid [server accent]
├── Thumb shadow: 0 2px 4px rgba(0,0,0,0.1)
│
├── [LABELS]
│   ├── Min/Max values at ends
│   └── Selected range above thumbs
│
└── [RANGE]
    ├── Min: 18
    ├── Max: 55
    └── Step: 1
```

### 4.5 Section Divider

| Property | Value |
|----------|-------|
| Height | 1px |
| Color | `var(--border-subtle)` |
| Margin | 24px 0 |

### 4.6 Footer Buttons

#### Reset Button (Tertiary)
| Property | Value |
|----------|-------|
| Height | 44px |
| Background | transparent |
| Text | [Server accent color], Inter 14px/500 |

#### Apply Button (Primary)
| Property | Value |
|----------|-------|
| Height | 52px |
| Width | flex: 1 |
| Background | [Server accent color] |
| Text | White, Inter 16px/600 |
| Border radius | `var(--radius-md)` → 12px |

---

## 5. Server-Specific Filter Content

### 5.1 Partner Server Filters

**Reference:** [Unora_PRD.md — Section 12.4](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)

| Section | Filter | Type | Options |
|---------|--------|------|---------|
| **Basics** | Age Range | Slider | 18-55 |
| | Gender | Multi-chip | Women, Men, Non-binary, All |
| **Compatibility** | Consistency Level | Multi-chip | Calm, Regular, High |
| | Availability | Multi-chip | Morning, Evening, Anytime |
| **Relationship** | Relationship Intent | Multi-chip | Long-term, Open to both |
| | Family Planning | Multi-chip | Want kids, Don't want, Open |

### 5.2 Friend Server Filters

| Section | Filter | Type | Options |
|---------|--------|------|---------|
| **Basics** | Age Range | Slider | 18-55 |
| **Compatibility** | Consistency Level | Multi-chip | Calm, Regular, High |
| | Availability | Multi-chip | Morning, Evening, Anytime |
| **Friendship Style** | Activity Preference | Multi-chip | Active (sports, hikes), Chill (talk, games), Both |
| | Social Energy | Multi-chip | Introvert-friendly, Extrovert, Ambivert |
| **Interests** | Hobby Categories | Multi-chip | Fitness, Creative, Outdoors, Learning, Music |

### 5.3 Growth Server Filters

| Section | Filter | Type | Options |
|---------|--------|------|---------|
| **Basics** | Age Range | Slider | 18-55 |
| **Compatibility** | Consistency Level | Multi-chip | Calm, Regular, High |
| | Availability | Multi-chip | Morning, Evening, Anytime |
| **Goals** | Goal Category | Multi-chip | Fitness, Career, Habits, Learning, Wellness |
| | Accountability Style | Multi-chip | Gentle check-ins, Structured, Strict push |
| | Time Commitment | Multi-chip | Light (weekly), Moderate (few times), Intense (daily) |

### 5.4 Universal Filters (All Servers)

| Filter | Type | Options |
|--------|------|---------|
| Consistency Level | Multi-chip | Calm, Regular, High |
| Availability | Multi-chip | Morning, Evening, Anytime |
| Age Range | Slider | 18-55 |

---

## 6. Interaction & Logic Specification

### 6.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Filter button (Discovery) | Open filter sheet |
| Tap | Close (✕) | Close sheet without saving |
| Tap | Scrim/backdrop | Close sheet without saving |
| Swipe down | Drag handle | Close sheet without saving |
| Tap | Chip | Toggle selection |
| Drag | Slider thumb | Adjust range |
| Tap | Reset | Clear all filters |
| Tap | Apply Filters | Save state, close sheet |

### 6.2 Apply Filters Logic

```
USER taps "Apply Filters"
    │
    ▼
┌─────────────────────────────────────────┐
│ 1. Save filter state to local storage  │
│ 2. Update "Filter Summary" caption     │
│    on Discovery header                  │
│ 3. Close sheet with slide-down anim    │
│ 4. **DO NOT auto-refresh**             │
│    (User must tap Refresh manually)     │
└─────────────────────────────────────────┘
    │
    ▼
Discovery Feed shows updated filter summary:
"Showing: Age 25-32 • High Consistency"
```

> **Important:** Filters do NOT trigger an automatic refresh. Users must manually tap the Refresh button to see new cards matching their filters.

### 6.3 Reset Logic

```
USER taps "Reset"
    │
    ▼
┌─────────────────────────────────────────┐
│ Clear all selected filters              │
│ Reset sliders to default (18-55)        │
│ All chips return to unselected state    │
│ "Apply Filters" remains available       │
└─────────────────────────────────────────┘
```

### 6.4 Entry/Exit Transitions

| Transition | Animation |
|------------|-----------|
| Open | Slide up from bottom, 300ms ease-out |
| Close (apply/cancel) | Slide down, 250ms ease-in |
| Scrim in | Fade 0 → 50%, 200ms |
| Scrim out | Fade 50% → 0, 150ms |
| Chip select | Scale 0.97 → 1.0, 100ms |

---

## 7. State Definitions

### 7.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Default | All chips unselected, full range | Fresh open |
| Filters Active | Selected chips highlighted | User made selections |
| Apply Ready | Button enabled | Any change from default |

### 7.2 Filter Summary Display

After applying, the Discovery header shows a summary:

| Filters Applied | Summary Text |
|-----------------|--------------|
| Age only | "Age 25-32" |
| Consistency only | "High Consistency" |
| Multiple | "Age 25-32 • High Consistency" |
| None/All | "Showing all" |

---

## 8. Content & Copy Guidelines

### 8.1 Dynamic Header

| Server | Header Text |
|--------|-------------|
| Partner | "Filter Partner Matches" |
| Friend | "Filter Friend Matches" |
| Growth | "Filter Growth Matches" |

### 8.2 Section Headers

| Section | Text |
|---------|------|
| Demographics | "BASICS" |
| Compatibility | "COMPATIBILITY" |
| Partner-specific | "RELATIONSHIP" |
| Friend-specific | "FRIENDSHIP STYLE" |
| Growth-specific | "GOALS" |

### 8.3 Button Labels

| Button | Label |
|--------|-------|
| Reset | "Reset" |
| Apply | "Apply Filters" |

### 8.4 Chip Labels

**Consistency:**
- Calm, Regular, High

**Availability:**
- Morning, Evening, Anytime

**Relationship Intent (Partner):**
- Long-term, Open to both

**Friendship Style (Friend):**
- Active, Chill, Both

**Accountability (Growth):**
- Gentle check-ins, Structured, Strict push

---

## 9. Accessibility

### 9.1 Screen Reader
- Sheet open: "Filter Partner Matches. Close button."
- Section: "Basics section. Age range slider, 25 to 32."
- Chip: "High consistency, selected" / "Calm, not selected"
- Buttons: "Reset filters" / "Apply filters button"

### 9.2 Touch Targets
- Chips: 40px height minimum
- Slider thumbs: 44px touch area
- Buttons: 44px+ height

### 9.3 Drag Handle
- Swipe down gesture to dismiss
- Alternative: Close button for accessibility

---

## 10. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Bottom sheet with drag handle | Critical | ☐ |
| Server-specific filter content | Critical | ☐ |
| Multi-select chips (server themed) | Critical | ☐ |
| Age range slider | High | ☐ |
| Scrollable content area | High | ☐ |
| Fixed footer (Reset + Apply) | Critical | ☐ |
| Save filter state | Critical | ☐ |
| Update filter summary | High | ☐ |
| Swipe to dismiss | Medium | ☐ |
| Dark mode | Medium | ☐ |

---

## 11. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 12.4 — Filter specifications |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 3.4 — Choice Chips |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.C — Discovery Flow |
| [Unora_Spec_07_DiscoveryFeed.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_07_DiscoveryFeed.md) | Parent screen |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
