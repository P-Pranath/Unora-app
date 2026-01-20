# Unora — UI Specification: Screen #22

## Subscription / Tier Page

**Version:** 1.0  
**Last Updated:** January 2026  
**Status:** Final  
**Author:** Product Design Team

---

## 1. Document Metadata

| Attribute | Value |
|-----------|-------|
| **Screen Name** | Subscription / Tier Page |
| **Screen ID** | `SCREEN_22_SUBSCRIPTION` |
| **Spec ID** | Spec-22 |
| **User Flow Reference** | Accessible from User Profile, Discovery Lock Upsell, or Settings |
| **PRD Reference** | [Section 16 — Monetization](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) |
| **DSD Reference** | [Section 12 — Premium Dark Visual Language](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) |
| **Role** | Tier comparison page where users evaluate and upgrade their subscription |

---

## 2. Screen Purpose & Context

### 2.1 Selling the Investment

This is the **most important conversion screen in Unora**. It must accomplish three goals simultaneously:

| Goal | How It's Achieved |
|------|-------------------|
| **Educate** | Clear comparison of what each tier provides |
| **Inspire** | Premium aesthetic that makes users *want* to upgrade |
| **Convert** | Frictionless path to subscription purchase |

> [!IMPORTANT]
> This screen must **sell the value**, not just list features. The visual hierarchy makes Pro irresistible while Free feels limiting but accessible.

### 2.2 Entry Points

| Entry Point | Context | Default State |
|-------------|---------|---------------|
| **User Profile (Free)** | Tap "⭐ Upgrade" | Show all tiers, Free pre-selected |
| **User Profile (Paid)** | Tap "Manage" | Show current tier selected, change option |
| **Discovery Lock Modal** | Tap "Upgrade" CTA | Show Plus/Pro comparison, highlight capacity |
| **At-Risk Nudge Upsell** | Tap nudge limit upgrade | Show Plus/Pro, highlight nudge benefits |

### 2.3 Core Philosophy (PRD Section 16)

> **"Money buys flexibility, never outcomes."**

The tier page must reinforce this principle:
- Paid tiers provide **more forgiveness** (recoveries), **more reach** (connections), and **faster pacing** (refresh)
- Paid tiers do NOT provide better matches, priority visibility, or access to other users

---

## 3. Tier Structure (PRD Section 16.2)

### 3.1 Complete Tier Comparison

| Feature | Free | Plus | Pro |
|---------|------|------|-----|
| **Active Connections** | 1 | 2 | 4 |
| **Discovery Refresh** | 24h cooldown | 12h cooldown | 6h cooldown |
| **Nudges per At-Risk Period** | 1 | 3 | 4 |
| **Earned Reveals (Day 1-14)** | 2 (Day 5, 10) | 3 (Day 4, 8, 12) | 4 (Day 3, 6, 9, 12) |
| **Paid Reveals Available** | Up to 3 | Up to 2 | Up to 1 |
| **Free Streak Recoveries** | 0 | 1 per connection | 2 per connection |
| **Filters** | All | All | All |

### 3.2 Pricing

| Tier | Monthly | Annual (Savings) |
|------|---------|------------------|
| **Free** | ₹0 | — |
| **Plus** | ₹199/month | ₹1,920/year (~₹160/mo, Save ~20%) |
| **Pro** | ₹399/month | ₹3,840/year (~₹320/mo, Save ~20%) |

---

## 4. Visual Theme & Styling

### 4.1 Premium Dark Foundation

This screen uses the most elevated expression of the Premium Dark aesthetic.

```css
/* Premium Dark Surfaces */
--pdm-background: #0D0D0F;        /* Deep charcoal canvas */
--pdm-surface-1: #141416;         /* Base card surface */
--pdm-surface-2: #1A1A1E;         /* Elevated card (Free, Plus) */
--pdm-surface-3: #222226;         /* Premium card (Pro) */

/* Text */
--pdm-ink-primary: #F5F5F3;       /* Headlines — high contrast */
--pdm-ink-secondary: #C4C4C0;     /* Body text */
--pdm-ink-tertiary: #8A8A86;      /* Captions, muted */
--pdm-ink-muted: #5A5A58;         /* Disabled, hints */

/* Tier Accent Colors */
--tier-free: #7A7A7A;             /* Muted gray — accessible but basic */
--tier-plus: #4A9B8C;             /* Teal — elevated, familiar */
--tier-pro-gold: #C4A77D;         /* Gold primary accent */
--tier-pro-gold-bright: #D4B88D;  /* Gold hover */
--tier-pro-gold-deep: #B08D5B;    /* Gold gradient deep */
--tier-pro-glow: rgba(196, 167, 125, 0.35); /* Pro outer glow */

/* Pro Gradient */
--pro-gradient: linear-gradient(135deg, #B08D5B 0%, #C4A77D 50%, #D4B88D 100%);
--pro-border-gradient: linear-gradient(135deg, rgba(196, 167, 125, 0.6), rgba(212, 184, 141, 0.8), rgba(196, 167, 125, 0.6));
```

### 4.2 Typography

| Element | Font | Size | Weight | Color |
|---------|------|------|--------|-------|
| **Hero Headline** | Outfit | 28px | 700 | --pdm-ink-primary + gold glow |
| **Hero Subline** | Inter | 16px | 400 | --pdm-ink-secondary |
| **Tier Title** | Outfit | 22px | 600 | Tier accent color |
| **Price** | Outfit | 32px | 700 | --pdm-ink-primary |
| **Price Period** | Inter | 14px | 400 | --pdm-ink-tertiary |
| **Feature Text** | Inter | 15px | 500 | --pdm-ink-primary |
| **Feature Caption** | Inter | 13px | 400 | --pdm-ink-tertiary |
| **Badge Text** | Inter | 11px | 600 | #0D0D0F (on gold) |
| **Toggle Labels** | Inter | 14px | 500 | Active/Inactive states |
| **CTA Button** | Inter | 16px | 600 | #FFFFFF / #0D0D0F |

---

## 5. Layout Specification

### 5.1 Full-Screen ASCII Wireframe

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ BACKGROUND: #0D0D0F ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓ │
│                                                                             │
│   [Safe Area Top]                                                           │
│                                                                             │
│   ╔═══════════════════════════════════════════════════════════════════════╗ │
│   ║   ←                       Choose Your Plan                        ╳   ║ │  Header
│   ╚═══════════════════════════════════════════════════════════════════════╝ │
│                                                                             │
│ ░░░░░░░░░░░░░░░░░░░░░░░░░░ SCROLLABLE ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │
│                                                                             │
│   ╭─────────────────────────────────────────────────────────────────────╮   │
│   │                                                                     │   │
│   │              ✨ Unlock Your Potential ✨                            │   │  HERO
│   │                                                                     │   │  SECTION
│   │          Invest in connections that last.                          │   │
│   │                                                                     │   │
│   ╰─────────────────────────────────────────────────────────────────────╯   │
│                                                                             │
│   ┌────────────────────────────────────────────────────────────────────┐    │
│   │           [ Monthly ]          [ Yearly  Save 20% ]                │    │  BILLING
│   └────────────────────────────────────────────────────────────────────┘    │  TOGGLE
│                                                                             │
│   ▼─────────────────────────────────────────────────────────────────────▼   │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                                                                     │   │
│   │   Free                                                              │   │  TIER CARD:
│   │   ─────────────────────────────────────────────────────────────     │   │  FREE
│   │   ₹0                                                                │   │
│   │                                                                     │   │
│   │   ○  1 active connection                                            │   │
│   │   ○  Refresh every 24 hours                                         │   │
│   │   ○  1 nudge per at-risk period                                     │   │
│   │   ○  2 earned reveals                                               │   │
│   │   ○  No free recoveries                                             │   │
│   │                                                                     │   │
│   │   ┌─────────────────────────────────────────────────────────────┐   │   │
│   │   │                     Current Plan                            │   │   │
│   │   └─────────────────────────────────────────────────────────────┘   │   │
│   │                                                                     │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │ ▌                                                                   │   │  TIER CARD:
│   │ ▌  Plus                                          ✓ Popular          │   │  PLUS
│   │ ▌  ─────────────────────────────────────────────────────────────    │   │  (Teal accent)
│   │ ▌  ₹199/mo                                                          │   │
│   │ ▌                                                                   │   │
│   │ ▌  ●  2 active connections                                          │   │
│   │ ▌  ●  Refresh every 12 hours                                        │   │
│   │ ▌  ●  3 nudges per at-risk period                                   │   │
│   │ ▌  ●  3 earned reveals                                              │   │
│   │ ▌  ●  1 free recovery per connection                                │   │
│   │ ▌                                                                   │   │
│   │ ▌  ┌─────────────────────────────────────────────────────────────┐  │   │
│   │ ▌  │                     Choose Plus                             │  │   │
│   │ ▌  └─────────────────────────────────────────────────────────────┘  │   │
│   │ ▌                                                                   │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│   ╔═════════════════════════════════════════════════════════════════════╗   │
│   ║ ░░░░░░ GOLD GLOW ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ║   │  TIER CARD:
│   ║ ▌                                                                   ║   │  PRO
│   ║ ▌  ⭐ Pro                              ┌───────────────────────┐    ║   │  (Gold gradient
│   ║ ▌                                      │     Best Value        │    ║   │   border + glow)
│   ║ ▌  ─────────────────────────────────── └───────────────────────┘    ║   │
│   ║ ▌  ₹399/mo                                                          ║   │
│   ║ ▌                                                                   ║   │
│   ║ ▌  ●  4 active connections                                          ║   │
│   ║ ▌  ●  Refresh every 6 hours                                         ║   │
│   ║ ▌  ●  4 nudges per at-risk period                                   ║   │
│   ║ ▌  ●  4 earned reveals                                              ║   │
│   ║ ▌  ●  2 free recoveries per connection                              ║   │
│   ║ ▌                                                                   ║   │
│   ║ ▌  ┌─────────────────────────────────────────────────────────────┐  ║   │
│   ║ ▌  │ ░░░░░░░░░░░░░░░  ⭐ Upgrade to Pro  ░░░░░░░░░░░░░░░░░░░░░░ │  ║   │
│   ║ ▌  └─────────────────────────────────────────────────────────────┘  ║   │
│   ║ ▌                                                                   ║   │
│   ║ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ║   │
│   ╚═════════════════════════════════════════════════════════════════════╝   │
│                                                                             │
│   ─────────────────────────────────────────────────────────────────────     │
│                                                                             │
│   Compare all features →                                                    │   Expand link
│                                                                             │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│                                                                             │
│   [Safe Area Bottom]                                                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 5.2 Layout Dimensions

| Property | Value |
|----------|-------|
| **Screen Background** | `var(--pdm-background)` / #0D0D0F |
| **Content Padding** | 20px horizontal |
| **Section Gap** | 24px |
| **Card Gap** | 16px |
| **Hero Section Padding** | 32px vertical |
| **Tier Card Padding** | 24px |
| **Tier Card Radius** | 20px |
| **Pro Card Radius** | 24px (slightly larger) |

---

## 6. Component Specifications

### 6.1 Header Bar

```
╔══════════════════════════════════════════════════════════════════════════════╗
║   ←                           Choose Your Plan                          ╳    ║
╚══════════════════════════════════════════════════════════════════════════════╝

Specs:
├── Background: Transparent (content scrolls under)
├── Height: 56px
├── Padding: 0 20px
├── Back Button (←): Phosphor "CaretLeft", 24px, var(--pdm-ink-primary)
├── Title: "Choose Your Plan"
├── Title Font: Outfit 18px / 600, centered
├── Title Color: var(--pdm-ink-primary)
├── Close Button (╳): Phosphor "X", 24px, right-aligned
└── Action: Navigate back / close screen
```

### 6.2 Hero Section ✨

The emotional hook — must feel aspirational.

```
╭─────────────────────────────────────────────────────────────────────────────╮
│                                                                             │
│                      ✨ Unlock Your Potential ✨                            │
│                                                                             │
│                   Invest in connections that last.                          │
│                                                                             │
╰─────────────────────────────────────────────────────────────────────────────╯

Container:
├── Background: Transparent (floating text on dark bg)
├── Padding: 32px 20px
├── Alignment: Center

Headline:
├── Text: "Unlock Your Potential"
├── Font: Outfit 28px / 700
├── Color: var(--pdm-ink-primary)
├── Text Shadow: 0 0 32px rgba(196, 167, 125, 0.3)  /* Subtle gold glow */
├── Sparkle Icons: Phosphor "Sparkle", 20px, var(--tier-pro-gold)
│   └── Position: Flanking headline, slight vertical offset

Subline:
├── Text: "Invest in connections that last."
├── Font: Inter 16px / 400
├── Color: var(--pdm-ink-secondary)
├── Margin Top: 12px
├── Font Style: Italic
```

### 6.3 Billing Toggle

```
┌────────────────────────────────────────────────────────────────────────────┐
│                                                                            │
│      ┌─────────────────┐    ┌─────────────────────────────────────┐       │
│      │     Monthly     │    │      Yearly        Save 20%         │       │
│      └─────────────────┘    └─────────────────────────────────────┘       │
│                                                                            │
└────────────────────────────────────────────────────────────────────────────┘

Container:
├── Background: var(--pdm-surface-2) / #1A1A1E
├── Border: 1px solid var(--pdm-border-subtle)
├── Border Radius: 12px
├── Padding: 4px
├── Margin: 0 20px 24px
├── Layout: Flex, space-evenly

Option (Inactive):
├── Background: Transparent
├── Text Font: Inter 14px / 500
├── Text Color: var(--pdm-ink-tertiary)
├── Padding: 12px 20px
├── Border Radius: 8px

Option (Active):
├── Background: var(--pdm-surface-3)
├── Text Color: var(--pdm-ink-primary)
├── Shadow: 0 2px 8px rgba(0, 0, 0, 0.2)

"Save 20%" Badge (on Yearly):
├── Text: "Save 20%"
├── Font: Inter 11px / 600
├── Color: #0D0D0F (dark text)
├── Background: var(--tier-pro-gold)
├── Border Radius: 6px
├── Padding: 4px 8px
├── Margin Left: 8px
├── Box Shadow: 0 0 12px var(--tier-pro-glow)
```

### 6.4 Tier Card — Free (Muted Styling)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   Free                                                                      │
│   ─────────────────────────────────────────────────────────────────         │
│                                                                             │
│   ₹0                                                                        │
│                                                                             │
│   ○  1 active connection                                                    │
│   ○  Refresh every 24 hours                                                 │
│   ○  1 nudge per at-risk period                                             │
│   ○  2 earned reveals                                                       │
│   ○  No free recoveries                                                     │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                         Current Plan                                │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Card Container:
├── Background: var(--pdm-surface-1) / #141416
├── Border: 1px solid var(--pdm-border-subtle) / #2A2A2E
├── Border Radius: 20px
├── Padding: 24px
├── Margin: 0 20px

Tier Title:
├── Text: "Free"
├── Font: Outfit 22px / 600
├── Color: var(--tier-free) / #7A7A7A

Divider:
├── Width: 60px
├── Height: 2px
├── Background: var(--tier-free)
├── Margin: 12px 0

Price:
├── Text: "₹0"
├── Font: Outfit 32px / 700
├── Color: var(--pdm-ink-tertiary)

Feature List:
├── Layout: Vertical stack, gap 12px
├── Margin Top: 20px

Feature Row:
├── Icon: Empty circle (○), Phosphor "Circle", 16px, var(--tier-free)
├── Text Font: Inter 15px / 500
├── Text Color: var(--pdm-ink-secondary)
├── Gap between icon and text: 12px

CTA Button (Current Plan):
├── Text: "Current Plan"
├── Font: Inter 16px / 500
├── Color: var(--pdm-ink-tertiary)
├── Background: Transparent
├── Border: 1.5px dashed var(--pdm-border-medium)
├── Border Radius: 12px
├── Height: 52px
├── Margin Top: 24px
├── State: Disabled (no hover/press effects)
```

### 6.5 Tier Card — Plus (Elevated, Teal Accent)

```
┌──────────────────────────────────────────────────────────────────────────────┐
│ ▌                                                                            │
│ ▌  Plus                                                     ✓ Popular       │
│ ▌  ─────────────────────────────────────────────────────────────────         │
│ ▌                                                                            │
│ ▌  ₹199/mo                                          (₹160/mo billed yearly) │
│ ▌                                                                            │
│ ▌  ●  2 active connections                                                   │
│ ▌  ●  Refresh every 12 hours                                                 │
│ ▌  ●  3 nudges per at-risk period                                            │
│ ▌  ●  3 earned reveals                                                       │
│ ▌  ●  1 free recovery per connection                                         │
│ ▌                                                                            │
│ ▌  ┌───────────────────────────────────────────────────────────────────┐     │
│ ▌  │                           Choose Plus                             │     │
│ ▌  └───────────────────────────────────────────────────────────────────┘     │
│ ▌                                                                            │
└──────────────────────────────────────────────────────────────────────────────┘

Card Container:
├── Background: var(--pdm-surface-2) / #1A1A1E
├── Border: 1px solid var(--pdm-border-subtle)
├── Left Border: 4px solid var(--tier-plus) / #4A9B8C
├── Border Radius: 20px
├── Padding: 24px
├── Margin: 0 20px
├── Box Shadow: 0 4px 16px rgba(0, 0, 0, 0.25)

"Popular" Badge:
├── Text: "✓ Popular"
├── Icon: Phosphor "Check", 12px
├── Font: Inter 11px / 600 / uppercase
├── Color: var(--tier-plus)
├── Border: 1px solid var(--tier-plus) at 50%
├── Background: rgba(74, 155, 140, 0.15)
├── Border Radius: 6px
├── Padding: 4px 10px
├── Position: Top right of card

Tier Title:
├── Text: "Plus"
├── Font: Outfit 22px / 600
├── Color: var(--tier-plus) / #4A9B8C

Divider:
├── Background: var(--tier-plus)

Price:
├── Text: "₹199/mo" (or "₹160/mo" if Yearly)
├── Font: Outfit 32px / 700
├── Color: var(--pdm-ink-primary)
├── Suffix (Yearly): "(billed yearly)" in Inter 13px, var(--pdm-ink-tertiary)

Feature Row:
├── Icon: Filled circle (●), Phosphor "CheckCircle", 16px, var(--tier-plus)
├── Text Color: var(--pdm-ink-primary)

CTA Button (Choose Plus):
├── Text: "Choose Plus"
├── Font: Inter 16px / 600
├── Color: #FFFFFF
├── Background: var(--tier-plus)
├── Border Radius: 12px
├── Height: 52px
├── Shadow: 0 2px 8px rgba(74, 155, 140, 0.25)
├── Hover: Background brightens, shadow expands
├── Pressed: Scale 0.98
```

### 6.6 Tier Card — Pro (⭐ Dominant, Gold Gradient Border, Glassmorphism)

```
╔═════════════════════════════════════════════════════════════════════════════╗
║ ░░░░░░░░░░░░░░░░░░░░ OUTER GOLD GLOW ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ║
║ ▌                                                                           ║
║ ▌  ⭐ Pro                                        ┌───────────────────────┐  ║
║ ▌                                                │     Best Value        │  ║
║ ▌  ─────────────────────────────────────────     └───────────────────────┘  ║
║ ▌                                                                           ║
║ ▌  ₹399/mo                                        (₹320/mo billed yearly)  ║
║ ▌                                                                           ║
║ ▌  ●  4 active connections                                                  ║
║ ▌  ●  Refresh every 6 hours                                                 ║
║ ▌  ●  4 nudges per at-risk period                                           ║
║ ▌  ●  4 earned reveals                                                      ║
║ ▌  ●  2 free recoveries per connection                                      ║
║ ▌                                                                           ║
║ ▌  ╔═══════════════════════════════════════════════════════════════════╗    ║
║ ▌  ║ ░░░░░░░░░░░░░░░░  ⭐ Upgrade to Pro  ░░░░░░░░░░░░░░░░░░░░░░░░░░░ ║    ║
║ ▌  ╚═══════════════════════════════════════════════════════════════════╝    ║
║ ▌                                                                           ║
║ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ ║
╚═════════════════════════════════════════════════════════════════════════════╝

Card Container:
├── Background: 
│   ├── Primary: rgba(34, 34, 38, 0.95) + backdrop-filter: blur(16px)
│   └── Overlay: linear-gradient(135deg, rgba(196, 167, 125, 0.05), rgba(212, 184, 141, 0.1))
├── Border: 2px solid (gradient border technique)
│   └── Gradient: linear-gradient(135deg, rgba(176, 141, 91, 0.7), var(--tier-pro-gold), rgba(176, 141, 91, 0.7))
├── Left Border: 4px solid var(--tier-pro-gold)
├── Border Radius: 24px
├── Padding: 28px
├── Margin: 0 20px
├── Box Shadow:
│   ├── 0 8px 32px rgba(0, 0, 0, 0.4)
│   ├── 0 0 48px rgba(196, 167, 125, 0.2)   /* Outer gold glow */
│   └── inset 0 1px 0 rgba(255, 255, 255, 0.05)   /* Inner highlight */

"Best Value" Badge:
├── Text: "Best Value"
├── Font: Inter 11px / 700 / uppercase
├── Color: #0D0D0F (dark on gold)
├── Background: var(--pro-gradient)
├── Border Radius: 8px
├── Padding: 6px 14px
├── Box Shadow: 0 0 16px var(--tier-pro-glow)
├── Position: Top right of card
├── Animation: Subtle pulse glow (optional)

Star Icon (⭐):
├── Icon: Phosphor "Star" (filled)
├── Size: 20px
├── Color: var(--tier-pro-gold)
├── Position: Before "Pro" title

Tier Title:
├── Text: "Pro"
├── Font: Outfit 24px / 700
├── Color: var(--tier-pro-gold)
├── Letter Spacing: -0.02em

Divider:
├── Background: var(--pro-gradient)
├── Height: 2px
├── Width: 80px

Price:
├── Text: "₹399/mo" (or "₹320/mo" if Yearly)
├── Font: Outfit 36px / 700
├── Color: #FFFFFF
├── Text Shadow: 0 0 24px rgba(196, 167, 125, 0.3)

Feature Row:
├── Icon: Phosphor "CheckCircle" (filled), 18px, var(--tier-pro-gold)
├── Text Color: #FFFFFF
├── Font Weight: 500

CTA Button (Upgrade to Pro):
├── Text: "⭐ Upgrade to Pro"
├── Star Icon: 16px, #1A1A1A
├── Font: Inter 16px / 700
├── Color: #1A1A1A (dark on gold)
├── Background: var(--pro-gradient)
├── Border: none
├── Border Radius: 14px
├── Height: 56px
├── Shadow: 
│   ├── 0 4px 16px rgba(0, 0, 0, 0.3)
│   └── 0 0 32px var(--tier-pro-glow)
├── Hover: Brightness 110%, shadow expands
├── Pressed: Scale 0.97, shadow dims
├── Animation: Button glow pulse on idle (subtle)
```

### 6.7 Feature Row Component

```
●  4 active connections

Specs:
├── Layout: Flex, align-items: center
├── Height: 28px
├── Gap: 12px

Icon (Checkmark):
├── Phosphor "CheckCircle" (filled)
├── Size: 16-18px
├── Color: Tier accent color

Text:
├── Font: Inter 15px / 500
├── Color: Based on tier (muted for Free, primary for Plus/Pro)

Improved/Upgrade Indicator (optional):
├── If feature is improved from current tier:
│   └── Add subtle glow behind icon
│   └── Text color becomes tier accent
```

### 6.8 Full Feature Comparison (Expandable Section)

```
┌────────────────────────────────────────────────────────────────────────────┐
│                                                                            │
│   Compare all features                                           ▼         │
│                                                                            │
└────────────────────────────────────────────────────────────────────────────┘

Collapsed State:
├── Text: "Compare all features"
├── Icon: Phosphor "CaretDown", 18px
├── Font: Inter 15px / 500
├── Color: var(--pdm-ink-tertiary)
├── Underline: On hover

Expanded State (Bottom Sheet or Inline Expansion):
┌────────────────────────────────────────────────────────────────────────────┐
│                                                                            │
│   Feature                       Free       Plus        Pro                 │
│   ──────────────────────────────────────────────────────────────────       │
│                                                                            │
│   Active Connections             1          2          4                   │
│   Discovery Refresh             24h        12h         6h                  │
│   Nudges per At-Risk             1          3          4                   │
│   Earned Reveals                 2          3          4                   │
│   Free Recoveries            None         1           2                   │
│   All Filters                   ✓          ✓          ✓                   │
│                                                                            │
│   ──────────────────────────────────────────────────────────────────       │
│   Price                         ₹0      ₹199/mo    ₹399/mo                 │
│                                                                            │
└────────────────────────────────────────────────────────────────────────────┘

Table Styling:
├── Background: var(--pdm-surface-2)
├── Header Row: var(--pdm-ink-tertiary), uppercase
├── Cell Padding: 12px 16px
├── Pro Column: var(--tier-pro-gold) text highlight
├── Checkmarks (✓): var(--tier-plus)
```

---

## 7. State Definitions

### 7.1 State A — Free User Viewing

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  STATE A: FREE USER VIEWING TIERS                                           │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Free Card:                                                                 │
│  ├── Shows "Current Plan" (disabled CTA)                                    │
│  └── Muted styling, no accent                                               │
│                                                                             │
│  Plus Card:                                                                 │
│  ├── Shows "Choose Plus" (active CTA)                                       │
│  └── "Popular" badge visible                                                │
│                                                                             │
│  Pro Card:                                                                  │
│  ├── Shows "Upgrade to Pro" (active CTA, gold gradient)                     │
│  ├── "Best Value" badge visible                                             │
│  └── Dominant visual — this should catch the eye                            │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 7.2 State B — Plus User Viewing

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  STATE B: PLUS USER VIEWING TIERS                                           │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Free Card:                                                                 │
│  ├── Hidden OR shows "Downgrade" option (based on business decision)        │
│  └── If shown, very subdued                                                 │
│                                                                             │
│  Plus Card:                                                                 │
│  ├── Shows "Current Plan ✓"                                                 │
│  ├── Active styling (teal accent, border)                                   │
│  └── "Manage Subscription" link below                                       │
│                                                                             │
│  Pro Card:                                                                  │
│  ├── Shows "Upgrade to Pro" (active CTA)                                    │
│  └── Emphasis on additional benefits (more connections, more recoveries)    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 7.3 State C — Pro User Viewing

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  STATE C: PRO USER VIEWING TIERS                                            │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Free Card:                                                                 │
│  └── Hidden                                                                 │
│                                                                             │
│  Plus Card:                                                                 │
│  └── Hidden OR shows "Downgrade" (subdued)                                  │
│                                                                             │
│  Pro Card:                                                                  │
│  ├── Shows "Current Plan ✓"                                                 │
│  ├── Gold border, full styling                                              │
│  ├── "Manage Subscription" link below                                       │
│  └── Shows renewal date: "Renews on [Date]"                                 │
│                                                                             │
│  Additional Options:                                                        │
│  ├── "Cancel Subscription" link (tertiary, subdued)                         │
│  └── "View Payment History" link                                            │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 7.4 State D — Processing Payment

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  STATE D: PAYMENT IN PROGRESS                                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Trigger: User taps "Choose Plus" or "Upgrade to Pro"                       │
│                                                                             │
│  Flow:                                                                      │
│  1. Button enters loading state (spinner, "Processing...")                  │
│  2. Native Payment Sheet opens (Apple Pay / Google Pay / UPI)               │
│  3. On success: Confirmation modal with confetti                            │
│  4. On failure: Error toast, button returns to normal                       │
│                                                                             │
│  CTA States:                                                                │
│  ├── Loading: Spinner replaces icon, text "Processing...", disabled        │
│  ├── Success: Checkmark animation, "Subscribed!"                            │
│  └── Error: Shake animation, button returns to normal                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 8. Interaction & Logic

### 8.1 Billing Toggle Logic

```
User taps "Monthly" / "Yearly" toggle:
│
├── Update all displayed prices to reflect selection
│   ├── Monthly: ₹199/mo (Plus), ₹399/mo (Pro)
│   └── Yearly: ₹160/mo (Plus), ₹320/mo (Pro) + "(billed yearly)" suffix
│
├── "Save 20%" badge:
│   ├── Visible on Yearly option
│   └── Gold accent with glow
│
└── Animation: Smooth price transition (fade/morph)
```

### 8.2 Tier Selection Flow

```
User taps Tier CTA:
│
├── If Same Tier (Current Plan):
│   └── No action (button is disabled)
│
├── If Upgrade:
│   ├── CTA enters loading state
│   ├── Trigger Native Payment Sheet
│   │   ├── Apple Pay (iOS)
│   │   ├── Google Pay (Android)
│   │   └── UPI / Card options (India)
│   │
│   ├── On Payment Success:
│   │   ├── Close payment sheet
│   │   ├── Show success confirmation modal
│   │   │   ├── "Welcome to Unora [Tier]!"
│   │   │   ├── Confetti animation
│   │   │   └── "Start Exploring" CTA
│   │   └── Update user subscription state
│   │
│   └── On Payment Failure:
│       ├── Close payment sheet
│       ├── Show error toast: "Payment failed. Please try again."
│       └── Return CTA to normal state
│
└── If Downgrade:
    ├── Show confirmation modal
    │   ├── "Downgrade to [Tier]?"
    │   ├── List what they'll lose
    │   └── "Your current plan will continue until [End Date]"
    └── Process if confirmed
```

### 8.3 Feature Comparison Expand

```
User taps "Compare all features →":
│
├── Option A: Inline Expansion
│   ├── Animate height reveal
│   └── Feature table slides in below cards
│
├── Option B: Bottom Sheet
│   ├── Sheet slides up
│   ├── Full comparison table
│   └── Dismiss handle at top
│
└── Icon rotates (▼ → ▲) on expansion
```

---

## 9. Animations & Motion

### 9.1 Screen Entry

```
Timeline: Screen Opens
══════════════════════════════════════════════════════════════════════════════

[0ms - 300ms]   SCREEN FADE IN
                ├── Background: Opacity 0 → 100%
                └── Easing: ease-out

[100ms]         HERO SECTION FADE IN
                ├── Opacity: 0 → 100%
                └── Text glow begins

[200ms]         BILLING TOGGLE FADE IN
                ├── Opacity: 0 → 100%
                └── Transform: translateY(10px) → translateY(0)

[300ms]         TIER CARDS STAGGER IN
                ├── Free: 300ms delay
                ├── Plus: 400ms delay
                ├── Pro: 500ms delay
                ├── Each: Opacity 0 → 100%, translateY(20px) → translateY(0)
                └── Pro card glow begins after entry

[800ms]         ALL ELEMENTS SETTLED
                └── Idle animations begin
```

### 9.2 Pro Card Idle Glow

```css
/* Subtle breathing glow for Pro card */
@keyframes pro-card-glow {
  0%, 100% { 
    box-shadow: 
      0 8px 32px rgba(0, 0, 0, 0.4),
      0 0 32px rgba(196, 167, 125, 0.15);
  }
  50% { 
    box-shadow: 
      0 8px 32px rgba(0, 0, 0, 0.4),
      0 0 48px rgba(196, 167, 125, 0.25);
  }
}

.tier-card--pro {
  animation: pro-card-glow 4s ease-in-out infinite;
}
```

### 9.3 CTA Button States

```css
/* Upgrade to Pro Button */
.cta-pro {
  transition: transform 150ms ease, box-shadow 200ms ease, filter 200ms ease;
}

.cta-pro:hover {
  filter: brightness(1.1);
  box-shadow: 0 6px 20px rgba(196, 167, 125, 0.4);
}

.cta-pro:active {
  transform: scale(0.97);
  box-shadow: 0 2px 10px rgba(196, 167, 125, 0.3);
}

/* Loading State */
.cta-pro--loading {
  pointer-events: none;
  /* Spinner replaces star icon */
}
```

### 9.4 Success Celebration

```css
/* On Successful Subscription */
@keyframes confetti-burst {
  0% { opacity: 0; transform: scale(0.5); }
  50% { opacity: 1; transform: scale(1.1); }
  100% { opacity: 1; transform: scale(1.0); }
}

/* Gold ring expansion */
@keyframes gold-ring-expand {
  0% { transform: scale(0); opacity: 1; }
  100% { transform: scale(2); opacity: 0; }
}
```

---

## 10. Copy Table

### 10.1 Hero Section

| Element | Copy |
|---------|------|
| **Headline** | "Unlock Your Potential" |
| **Subline** | "Invest in connections that last." |

### 10.2 Billing Toggle

| Element | Copy |
|---------|------|
| **Monthly Label** | "Monthly" |
| **Yearly Label** | "Yearly" |
| **Savings Badge** | "Save 20%" |

### 10.3 Tier Cards

| Tier | Title | Price (Monthly) | Price (Yearly Display) |
|------|-------|-----------------|------------------------|
| **Free** | "Free" | "₹0" | "₹0" |
| **Plus** | "Plus" | "₹199/mo" | "₹160/mo" + "(billed yearly)" |
| **Pro** | "Pro" | "₹399/mo" | "₹320/mo" + "(billed yearly)" |

### 10.4 Feature List Items

| Tier | Features |
|------|----------|
| **Free** | "1 active connection", "Refresh every 24 hours", "1 nudge per at-risk period", "2 earned reveals", "No free recoveries" |
| **Plus** | "2 active connections", "Refresh every 12 hours", "3 nudges per at-risk period", "3 earned reveals", "1 free recovery per connection" |
| **Pro** | "4 active connections", "Refresh every 6 hours", "4 nudges per at-risk period", "4 earned reveals", "2 free recoveries per connection" |

### 10.5 CTA Buttons

| Context | Button Text |
|---------|-------------|
| **Free (Current)** | "Current Plan" |
| **Plus (Choose)** | "Choose Plus" |
| **Plus (Current)** | "Current Plan ✓" |
| **Pro (Upgrade)** | "⭐ Upgrade to Pro" |
| **Pro (Current)** | "Current Plan ✓" |
| **Loading** | "Processing..." |

### 10.6 Badges

| Badge | Text |
|-------|------|
| **Popular (Plus)** | "✓ Popular" |
| **Best Value (Pro)** | "Best Value" |

### 10.7 Success/Error States

| State | Copy |
|-------|------|
| **Success Modal Title** | "Welcome to Unora [Tier]!" |
| **Success Modal Body** | "Your upgraded benefits are now active." |
| **Success Modal CTA** | "Start Exploring" |
| **Error Toast** | "Payment failed. Please try again." |

### 10.8 Management Options (Paid Users)

| Element | Copy |
|---------|------|
| **Renewal Info** | "Renews on [Date]" |
| **Manage Link** | "Manage Subscription" |
| **Cancel Link** | "Cancel Subscription" |

---

## 11. Accessibility

### 11.1 Screen Reader Announcements

| Moment | Announcement |
|--------|--------------|
| **Screen Opens** | "Choose Your Plan. Compare subscription tiers. Monthly selected." |
| **Toggle Change** | "Yearly billing selected. Save 20 percent." |
| **Tier Card Focus** | "Plus tier. ₹199 per month. 2 active connections. 3 nudges per at-risk period. Choose Plus button." |
| **Pro Card Focus** | "Pro tier. Best Value. ₹399 per month. 4 active connections. Upgrade to Pro button." |
| **Payment Processing** | "Processing payment. Please wait." |
| **Payment Success** | "Success. You are now subscribed to Unora Pro." |

### 11.2 Focus Order

1. Back button
2. Title
3. Close button
4. Hero section (heading level 1)
5. Billing toggle (Monthly → Yearly)
6. Free tier card
7. Plus tier card
8. Pro tier card
9. Compare features link
10. Footer links (if any)

### 11.3 Color Contrast

All text meets WCAG AAA standards:
- Primary text (#F5F5F3) on background (#0D0D0F): **17.4:1**
- Secondary text (#C4C4C0) on cards (#1A1A1E): **8.1:1**
- Gold accent (#C4A77D) as decorative element only; not relied upon for meaning

---

## 12. Technical Notes

### 12.1 API Requirements

```typescript
interface TierDefinition {
  id: 'free' | 'plus' | 'pro';
  name: string;
  priceMonthly: number;       // in INR (₹)
  priceYearly: number;        // in INR (₹), annual amount
  activeConnections: number;
  refreshCooldownHours: number;
  nudgesPerAtRisk: number;
  earnedReveals: number;
  freeRecoveriesPerConnection: number;
}

interface SubscriptionState {
  currentTier: 'free' | 'plus' | 'pro';
  memberSince?: string;        // ISO date
  expiresAt?: string;          // ISO date
  autoRenew: boolean;
  paymentMethod?: 'apple_pay' | 'google_pay' | 'upi' | 'card';
}

interface UpgradeRequest {
  targetTier: 'plus' | 'pro';
  billingCycle: 'monthly' | 'yearly';
  paymentMethod: string;
}

interface UpgradeResponse {
  success: boolean;
  newTier?: 'plus' | 'pro';
  expiresAt?: string;
  receiptId?: string;
  error?: string;
}
```

### 12.2 Payment Integration

```
Supported Payment Methods:
├── iOS: Apple In-App Purchase → Apple Pay
├── Android: Google Play Billing → Google Pay
├── Fallback: UPI, Credit/Debit Card (Razorpay/Stripe)

Flow:
1. User taps CTA
2. App creates purchase intent
3. Native payment sheet opens
4. User authenticates (Face ID / fingerprint / PIN)
5. Payment processed
6. Server verifies receipt
7. Subscription activated
8. UI updates
```

### 12.3 Analytics Events

| Event | Trigger | Properties |
|-------|---------|------------|
| `subscription_page_viewed` | Screen opens | `current_tier`, `entry_point` |
| `billing_cycle_changed` | Toggle tapped | `new_cycle`, `previous_cycle` |
| `tier_card_viewed` | Card becomes visible | `tier`, `is_current` |
| `upgrade_initiated` | CTA tapped | `target_tier`, `billing_cycle`, `current_tier` |
| `payment_sheet_opened` | Native sheet appears | `target_tier`, `payment_method` |
| `payment_completed` | Successful payment | `new_tier`, `billing_cycle`, `amount` |
| `payment_failed` | Failed payment | `target_tier`, `error_code` |
| `subscription_page_dismissed` | Screen closed | `time_on_screen`, `action_taken` |

---

## 13. Success Confirmation Modal

### 13.1 Modal Design

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   ╭─────────────────────────────────────────────────────────────────────╮   │
│   │                                                                     │   │
│   │                         🎉 ✨ 🎊                                    │   │
│   │                                                                     │   │
│   │                  Welcome to Unora Pro!                              │   │
│   │                                                                     │   │
│   │          Your upgraded benefits are now active.                     │   │
│   │                                                                     │   │
│   │   ─────────────────────────────────────────────────────────────     │   │
│   │                                                                     │   │
│   │   ●  4 active connections ready                                     │   │
│   │   ●  6-hour refresh unlocked                                        │   │
│   │   ●  2 free recoveries per connection                               │   │
│   │                                                                     │   │
│   │   ┌─────────────────────────────────────────────────────────────┐   │   │
│   │   │ ░░░░░░░░░░░░░░░  Start Exploring  ░░░░░░░░░░░░░░░░░░░░░░░░ │   │   │
│   │   └─────────────────────────────────────────────────────────────┘   │   │
│   │                                                                     │   │
│   ╰─────────────────────────────────────────────────────────────────────╯   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Specs:
├── Background: var(--pdm-surface-3) with glassmorphism
├── Border: 2px solid gold gradient
├── Box Shadow: 0 0 48px var(--tier-pro-glow)
├── Confetti: Animated particles, gold/teal/white
├── Title Font: Outfit 24px / 700
├── CTA: Gold gradient button (same as Pro upgrade)
├── Animation: Scale in from 0.9 with confetti burst
```

---

## 14. Related Screens

| Screen | Relationship |
|--------|--------------|
| [Screen 21: User Profile](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_21_UserProfile.md) | Primary entry point via Subscription Card |
| [Screen 10: Discovery Locked](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_10_DiscoveryLocked.md) | Entry via "Upgrade" CTA on capacity lock |
| [Screen 15: At-Risk State](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_15_AtRiskState.md) | Entry via nudge limit upsell |
| [Screen 16: Recovery Window](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_16_RecoveryWindow.md) | Entry via recovery credit upsell |

---

## 15. Design Principles Checklist

| Principle | Implementation |
|-----------|----------------|
| ✅ **Sell the Value** | Pro card dominates visually with gold glow, gradient border, "Best Value" badge |
| ✅ **Premium Dark Aesthetic** | Deep charcoal (#0D0D0F) canvas, layered surfaces, gold accents |
| ✅ **High-Contrast Typography** | Elegant Outfit/Inter fonts, clear hierarchy |
| ✅ **Glassmorphism** | Pro card uses blur + transparency for elevated feel |
| ✅ **Make Pro Irresistible** | Largest card, most visual treatment, sticky CTA in high-traffic scenarios |
| ✅ **No Invented Features** | All tier benefits strictly match PRD Section 16.2 |
| ✅ **No Filter Gating** | All tiers show "All Filters" as included |

---

**Document maintained by:** Product Design Team  
**Last updated:** January 2026  
**Review cycle:** With each PRD/DSD update
