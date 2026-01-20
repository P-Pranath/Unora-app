# Unora UI Specification — Server Selection

**Document ID:** Spec-06  
**Screen Name:** Server Selection (Onboarding)  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Server Selection (Onboarding)** — Intent segmentation screen

### 1.2 User Flow Reference
**Phase 1 → Phase 2 Transition** — This screen bridges onboarding and discovery. Users arrive after completing Profile Enrichment and Profile Review.

**Sequence:**
```
Profile Enrichment → Profile Review → [Server Selection] → Discovery
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.B.1](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Segment user intent before entering the main app. Users choose their initial focus (Romantic, Platonic, or Growth) to receive relevant matches.

### 1.4 Primary Action
**Select a server** and tap "Enter Server" to begin Discovery.

---

## 2. Low-Fidelity Wireframe (ASCII)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│                                                             │
│         What brings you to Unora?                           │  ← Headline (H1)
│                                                             │
│         Choose your focus.                                  │  ← Subtitle
│         You can switch anytime.                             │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  🔥                                                 │   │
│   │                                                     │   │
│   │  Looking for a Partner                              │   │  ← Server Card 1
│   │  Find someone for a meaningful relationship.        │   │     (Terracotta)
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  👋                                                 │   │
│   │                                                     │   │
│   │  Friend / Companion                                 │   │  ← Server Card 2
│   │  Connect with someone who gets you.                 │   │     (Teal)
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  🎯                                                 │   │
│   │                                                     │   │
│   │  Growth Buddy                                       │   │  ← Server Card 3
│   │  Build habits and goals together.                   │   │     (Indigo)
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │  ← Action Area
│   │                   Enter Server                      │   │     (Fixed bottom)
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.1 Selected State

```
┌─────────────────────────────────────────────────────────────┐
│   ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓   │
│   ┃  🔥  ✓                                              ┃   │  ← SELECTED
│   ┃                                                     ┃   │     Border: 2px
│   ┃  Looking for a Partner                              ┃   │     Accent color
│   ┃  Find someone for a meaningful relationship.        ┃   │     Checkmark visible
│   ┃                                                     ┃   │
│   ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  👋                                                 │   │  ← Unselected
│   │  Friend / Companion                                 │   │
│   │  Connect with someone who gets you.                 │   │
│   └─────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
SERVER SELECTION CONTAINER
├── Position: fixed, 100vw × 100vh
├── Background: var(--surface-background) → #FAFAF8
├── Display: flex, column
│
├── [HEADER AREA]
│   ├── Padding: var(--padding-screen) → 20px
│   ├── Padding-top: env(safe-area-inset-top) + 24px
│   └── Text-align: center
│
├── [CARDS AREA] — flex: 1, scrollable
│   ├── Padding: 20px horizontal
│   ├── Gap: var(--space-4) → 16px (between cards)
│   └── Padding-bottom: 120px (space for action area)
│
└── [ACTION AREA] — fixed bottom
    ├── Position: fixed, bottom: 0
    ├── Width: 100%
    ├── Padding: 20px
    ├── Padding-bottom: env(safe-area-inset-bottom) + 20px
    └── Background: var(--surface-background) with blur

Premium Dark Mode (Default):
├── Background: var(--pdm-background) → #0D0D0F
├── Cards: Elevated surfaces with inner glow
├── Selected card: Server color border with outer glow
└── Action area: Glassmorphism blur backdrop
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Card default** | Surface `#1A1A1E`, border `#2A2A2E`, inner glow |
| **Card hover** | Border brightens to `#3D3D42`, subtle lift |
| **Card selected** | Server accent border (2px) + outer glow |
| **Server icon** | Filled with server color, subtle icon glow |
| **Checkmark** | Server color with glow: `0 0 8px rgba(server-color, 0.4)` |
| **Action area** | Glass backdrop: `blur(16px)`, semi-transparent bg |

**Selected Card Glow (by server):**
```css
/* Partner server selection glow */
.card.partner.selected {
  border: 2px solid #E07D5A;
  box-shadow: 0 0 16px rgba(224, 125, 90, 0.3),
              0 4px 16px rgba(0, 0, 0, 0.3);
}

/* Friend server selection glow */
.card.friend.selected {
  border: 2px solid #4A9B8C;
  box-shadow: 0 0 16px rgba(74, 155, 140, 0.3),
              0 4px 16px rgba(0, 0, 0, 0.3);
}

/* Growth server selection glow */
.card.growth.selected {
  border: 2px solid #7B8AD9;
  box-shadow: 0 0 16px rgba(123, 138, 217, 0.3),
              0 4px 16px rgba(0, 0, 0, 0.3);
}
```



### 3.2 Card Dimensions

| Property | Value |
|----------|-------|
| Width | 100% (full container width - padding) |
| Min-height | 100px |
| Padding | 20px |
| Border radius | `var(--radius-lg)` → 16px |
| Touch target | Full card area |

### 3.3 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen padding | `var(--padding-screen)` | 20px |
| Card gap | `var(--space-4)` | 16px |
| Headline margin-bottom | `var(--space-2)` | 8px |
| Subtitle margin-bottom | `var(--space-8)` | 40px |
| Icon margin-bottom | `var(--space-3)` | 12px |

### 3.3 Z-Index Layers

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Background | 0 | Screen bg |
| Content | 1 | Cards |
| Action Area | 10 | Fixed CTA |
| System | 100+ | Status bar |

---

## 4. Component Inventory

### 4.1 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Headline (H1) | Outfit | 600 | 28px | `--unora-ink-primary` |
| Subtitle | Inter | 400 | 16px | `--unora-ink-secondary` |
| Card Title | Outfit | 600 | 18px | `--unora-ink-primary` |
| Card Description | Inter | 400 | 14px | `--unora-ink-secondary` |
| Button Text | Inter | 600 | 16px | White |

### 4.2 Server Card Component (Selectable)

```
SELECTABLE SERVER CARD
├── Width: 100%
├── Min-height: 100px
├── Padding: 20px
├── Border radius: var(--radius-lg) → 16px
├── Background: var(--surface-card) → #FFFFFF
├── Border: 1.5px solid var(--border-medium)
│
├── [ICON AREA]
│   ├── Size: 32px
│   ├── Margin-bottom: var(--space-3) → 12px
│   └── Color: Server accent color
│
├── [TEXT AREA]
│   ├── Title: Outfit 18px / 600
│   ├── Description: Inter 14px / 400
│   └── Gap: 4px
│
└── [CHECKMARK] — visible only when selected
    ├── Position: Top-right, 16px from edges
    ├── Size: 24px
    └── Color: Server accent color
```

### 4.3 Card States

#### Default State
| Property | Value |
|----------|-------|
| Background | `var(--surface-card)` → #FFFFFF |
| Border | 1.5px solid `var(--border-medium)` → #D4D4D0 |
| Icon | Outlined, server accent color |
| Checkmark | Hidden |

#### Selected State
| Property | Value |
|----------|-------|
| Background | Server accent color @ 8% opacity |
| Border | 2px solid [server accent color] |
| Icon | Filled, server accent color |
| Checkmark | Visible, server accent color |
| Shadow | 0 4px 12px [server color] @ 15% |

#### Pressed State
| Property | Value |
|----------|-------|
| Scale | 0.98 |
| Transition | 100ms ease-out |

### 4.4 Server Color Tokens & Icons

| Server | Token (Light) | Hex (Light) | Token (Dark) | Hex (Dark) | Icon (Phosphor) |
|--------|---------------|-------------|--------------|------------|----------------|
| **Partner** | `--server-partner-primary` | #D4714A | `--server-partner-primary` | #E07D5A | HeartStraight |
| **Friend** | `--server-friend-primary` | #4A9B8C | `--server-friend-primary` | #4A9B8C | HandWaving |
| **Growth** | `--server-growth-primary` | #5B6ABF | `--server-growth-primary` | #7B8AD9 | Target |

### 4.5 Primary Button ("Enter Server")

| Property | Value |
|----------|-------|
| Height | 52px |
| Background | `var(--unora-primary-accent)` → #C4A77D |
| Text | White, Inter 16px/600 |
| Border radius | `var(--radius-md)` → 12px |
| Disabled | Opacity 0.4 (when no card selected) |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Server Card | Select that card (radio behavior) |
| Tap | "Enter Server" | Navigate to Discovery with selected server |

### 5.2 Selection Behavior

```
SELECTION LOGIC (Radio Button Behavior)

USER taps a server card
    │
    ▼
┌─────────────────────────────────────────┐
│ Only ONE card can be selected at a time │
└─────────────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────────────┐
│ Tapped card:                            │
│ ├── Border → 2px [server accent]        │
│ ├── Background → [accent] @ 8%          │
│ ├── Checkmark → Visible                 │
│ └── Haptic: light                       │
└─────────────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────────────┐
│ Previously selected card (if any):      │
│ ├── Border → default                    │
│ ├── Background → white                  │
│ └── Checkmark → Hidden                  │
└─────────────────────────────────────────┘
    │
    ▼
"Enter Server" button → ENABLED
```

### 5.3 Entry Behavior

```
USER taps "Enter Server"
    │
    ▼
┌─────────────────────────────────────────┐
│ [DECISION] Global Capacity Lock Check   │
│ Active Connections == Tier Limit?       │
└─────────────────────────────────────────┘
    │                       │
    ▼                       ▼
  YES                      NO
    │                       │
    ▼                       ▼
┌───────────────┐   ┌─────────────────────────────────┐
│ LOCKED STATE  │   │ SYSTEM:                         │
│               │   │ ├── Save selected server        │
│ Route to      │   │ ├── Apply server color theme    │
│ Capacity Lock │   │ ├── Navigate to Discovery       │
│ Screen        │   │ └── Show empty state or cached  │
│               │   │     cards                        │
└───────────────┘   └─────────────────────────────────┘

**Reference:** PRD Section 12.5 — Lock 2 (Capacity-Based Lock)
```

### 5.4 Transitions

| Transition | Animation |
|------------|-----------|
| Card selection | Border/bg animate 150ms ease-out |
| Button enable | Opacity 0.4 → 1.0, 150ms |
| Exit to Discovery | Slide left, 300ms ease-out |
| Checkmark appear | Scale 0 → 1, 200ms with bounce |

---

## 6. State Definitions

### 6.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Default | No selection, button disabled | Initial load |
| Card Selected | One card highlighted, button enabled | User tapped a card |
| Loading | Button shows spinner | Navigating to Discovery |

### 6.2 Default State (No Selection)

```
All cards: Default styling
Button: "Enter Server" — DISABLED (opacity 0.4)
```

### 6.3 Selection Made

```
Selected card: Accent border, bg tint, checkmark
Other cards: Default styling
Button: "Enter Server" — ENABLED
```

---

## 7. Content & Copy Guidelines

### 7.1 Header Copy

| Element | Copy |
|---------|------|
| Headline | "What brings you to Unora?" |
| Subtitle | "Choose your focus. You can switch anytime." |

### 7.2 Card Content

| Server | Title | Description |
|--------|-------|-------------|
| **Partner** | "Looking for a Partner" | "Find someone for a meaningful relationship." |
| **Friend** | "Friend / Companion" | "Connect with someone who gets you." |
| **Growth** | "Growth Buddy" | "Build habits and goals together." |

### 7.3 Button Labels

| State | Label |
|-------|-------|
| Disabled | "Enter Server" |
| Enabled | "Enter Server" |
| Loading | "Loading..." |

### 7.4 Tone Guidelines

| Principle | Application |
|-----------|-------------|
| Non-judgmental | All three options presented equally |
| Flexible | "You can switch anytime" reduces pressure |
| Clear | Simple descriptions, no jargon |

---

## 8. Accessibility

### 8.1 Screen Reader
- Headline: "What brings you to Unora? Choose your focus."
- Cards: "Looking for a Partner. Find someone for a meaningful relationship. Button."
- Selected: "Looking for a Partner. Selected."

### 8.2 Touch Targets
- Cards: Full card area (min 100px height)
- Button: 52px height

### 8.3 Focus Order
1. Headline (skip)
2. Partner card
3. Friend card
4. Growth card
5. Enter Server button

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| 3 server cards with icons | Critical | ☐ |
| Radio selection behavior | Critical | ☐ |
| Server color theming | High | ☐ |
| Button enable/disable | High | ☐ |
| Selection animation | Medium | ☐ |
| Checkmark animation | Medium | ☐ |
| Navigate to Discovery | Critical | ☐ |
| Save server preference | Critical | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 11 — Server Selection |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Server color tokens |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.B — Server Selection |
| [Unora_Spec_04_ProfileCreation.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_04_ProfileCreation.md) | Previous screen |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
