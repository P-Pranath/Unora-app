# Unora — UI Specification: Screen #17

## Milestone Reveal Event

**Version:** 1.0  
**Last Updated:** January 2026  
**Status:** Final  
**Author:** Product Design Team

---

## 1. Document Metadata

| Attribute | Value |
|-----------|-------|
| **Screen Name** | Milestone Reveal Event |
| **Screen ID** | `SCREEN_17_REVEAL_EVENT` |
| **User Flow Reference** | [Section 2.E — Reveal Progression & Tier Isolation](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) |
| **PRD Reference** | [Section 15 — Reveal System](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) |
| **DSD Reference** | [Section 4.3 — Reveal Milestones](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md), [Section 4.4 — Blur Logic](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) |
| **Visibility** | User who reaches milestone day (evaluated independently per user) |
| **Trigger** | Streak reaches a tier-specific reveal milestone |

---

## 2. Experience Philosophy

### 2.1 The Cinematic Moment

This is not a notification. This is not an inline update. This is a **cinematic event** — a reward for patience and consistency.

The UI must feel:

| Quality | Implementation |
|---------|----------------|
| **Premium** | Rich dark surfaces, gold accents, elegant typography |
| **Immersive** | Full-screen takeover with dimmed backdrop |
| **Rewarding** | Celebration copy, haptic feedback, smooth animations |
| **Anticipatory** | "Tap to Reveal" mechanic builds one final moment of suspense |
| **Earned** | Copy emphasizes trust built through consistency |

> [!IMPORTANT]
> This screen uses **Dark Mode styling exclusively**, regardless of user's system preference. The dark backdrop makes the revealed content "pop" and creates a premium, cinematic atmosphere.

### 2.2 Design Tone

| ❌ Avoid | ✅ Use Instead |
|---------|---------------|
| Plain white backgrounds | Deep dark surfaces (#121212, #1A1A1A) |
| Grey/muted text | Gold accent (var(--unora-primary-accent)) for headers |
| Flat, static layouts | Layered cards with depth, glow effects, animations |
| Generic "unlocked" messaging | Rewarding, trust-focused copy ("Trust earned...") |
| Auto-reveal without interaction | "Tap to Reveal" / "Unwrap" metaphor |

---

## 3. Reveal Milestones (Source of Truth)

### 3.1 The Five Reveals

| Reveal | Name | Contents | Day (Varies by Tier) |
|--------|------|----------|---------------------|
| **R1** | Personality Spark | Age (exact), Height (range), City area, Hobby depth, Personality signal | Free: Day 5, Plus: Day 4, Pro: Day 3 |
| **R2** | Lifestyle Reality | Profession domain, Education level, Body-type (optional), Routine type | Free: Day 10, Plus: Day 8, Pro: Day 6 |
| **R3** | Social & Emotional Context | Religion (opt-in), Cultural background, Language, Emotional signal | Plus: Day 12, Pro: Day 9 |
| **R4** | Human Presence | Voice note, Candid photo (non-face), Immutable fact | Pro: Day 12 |
| **R5** | Identity + Chat Unlock | Full name, Face photos, Chat functionality | **All Tiers: Day 15 (NOT accelerable)** |

### 3.2 Tier-Specific Reveal Schedules

```
    THE PATH TO DAY 15 — REVEAL TIMELINE
    ═══════════════════════════════════════════════════════════════════════

    Day:  1   2   3   4   5   6   7   8   9  10  11  12  13  14  15
          │   │   │   │   │   │   │   │   │   │   │   │   │   │   │
    ──────┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴──

    FREE          │           │                               │
                  │           │                               │
                [R1]        [R2]                           [R5]
               Day 5       Day 10                         Day 15
            Personality  Lifestyle                   Identity+Chat

    ──────────────────────────────────────────────────────────────────

    PLUS      │           │           │                       │
              │           │           │                       │
            [R1]        [R2]        [R3]                   [R5]
           Day 4       Day 8      Day 12                  Day 15
        Personality  Lifestyle   Social               Identity+Chat

    ──────────────────────────────────────────────────────────────────

    PRO   │       │       │       │                           │
          │       │       │       │                           │
        [R1]    [R2]    [R3]    [R4]                       [R5]
       Day 3   Day 6   Day 9   Day 12                     Day 15
     Personality Lifestyle Social  Human              Identity+Chat
                                  Presence

    ═══════════════════════════════════════════════════════════════════
```

### 3.3 Critical Logic Constraints

> [!CAUTION]
> The following rules are **non-negotiable**.

| Constraint | Implementation |
|------------|----------------|
| **Tier Isolation** | Each user's reveals are based solely on their own tier. A Pro user's faster reveals do NOT transfer to their Free partner. |
| **Day 15 is Universal** | Reveal 5 (Identity + Chat) always requires 15 days. No tier accelerates it. No payment bypasses it. |
| **Manual Reveal** | Content is NOT auto-revealed. User must "Tap to Reveal" to see content. |
| **Independent Evaluation** | System evaluates each user independently at every milestone. |

---

## 4. Visual Theme & Styling

### 4.1 Dark Cinematic Foundation

This screen uses a **permanent dark theme** to create an immersive, premium reveal experience.

```css
/* Cinematic Dark Foundation */
--reveal-bg-deep: #0D0D0D;         /* Almost black — deepest layer */
--reveal-bg-surface: #1A1A1A;      /* Primary dark surface (var(--dm-surface-background)) */
--reveal-bg-card: #252525;         /* Elevated card surface */
--reveal-bg-glow: rgba(196, 167, 125, 0.08); /* Gold ambient glow */

/* Text Hierarchy — Light on Dark */
--reveal-text-primary: #FFFFFF;    /* Pure white for headers */
--reveal-text-secondary: #E5E5E3;  /* Off-white for body */
--reveal-text-muted: #8A8A86;      /* Muted for captions */

/* Accent — Gold / Brand */
--reveal-accent-gold: #C4A77D;     /* var(--unora-primary-accent) — celebratory gold */
--reveal-accent-gold-bright: #D4B88E; /* Brighter gold for emphasis */
--reveal-accent-glow: rgba(196, 167, 125, 0.25); /* Gold glow effect */

/* Reveal-Specific */
--reveal-card-border: rgba(196, 167, 125, 0.3); /* Subtle gold border */
--reveal-overlay: rgba(0, 0, 0, 0.85); /* Deep backdrop dimming */
```

### 4.2 Typography

| Element | Font | Size | Weight | Color |
|---------|------|------|--------|-------|
| **Milestone Badge** | Outfit | 12px | 600 | var(--reveal-accent-gold) |
| **Celebratory Header (H1)** | Outfit | 28px | 700 | #FFFFFF |
| **Subheader** | Outfit | 18px | 500 | var(--reveal-accent-gold) |
| **Reveal Card Title** | Inter | 16px | 600 | #FFFFFF |
| **Reveal Card Content** | Inter | 15px | 400 | var(--reveal-text-secondary) |
| **AI-Framed Text** | Inter | 15px | 400 italic | var(--reveal-text-secondary) |
| **Button Text** | Inter | 16px | 600 | #FFFFFF or var(--reveal-bg-surface) |
| **Caption/Microcopy** | Inter | 12px | 400 | var(--reveal-text-muted) |

---

## 5. Layout Specification

### 5.1 Full-Screen Overlay Structure

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░ DEEP DARK BACKDROP ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░░░░ #0D0D0D ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░                                                    ░░░░░░░░░░   │
│   ░░░░░░░░   ┌──────────────────────────────────────────┐     ░░░░░░░░░░   │
│   ░░░░░░░░   │                                          │     ░░░░░░░░░░   │
│   ░░░░░░░░   │            ✦ DAY 5 UNLOCKED ✦            │     ░░░░░░░░░░   │  Milestone Badge
│   ░░░░░░░░   │                                          │     ░░░░░░░░░░   │
│   ░░░░░░░░   │       Trust earned. Here's a glimpse.    │     ░░░░░░░░░░   │  Celebratory Header
│   ░░░░░░░░   │                                          │     ░░░░░░░░░░   │
│   ░░░░░░░░   │     ╔══════════════════════════════╗     │     ░░░░░░░░░░   │
│   ░░░░░░░░   │     ║                              ║     │     ░░░░░░░░░░   │
│   ░░░░░░░░   │     ║     ┌────────────────────┐   ║     │     ░░░░░░░░░░   │
│   ░░░░░░░░   │     ║     │                    │   ║     │     ░░░░░░░░░░   │  REVEAL CARD
│   ░░░░░░░░   │     ║     │   🎁 Tap to Reveal │   ║     │     ░░░░░░░░░░   │  (Locked State)
│   ░░░░░░░░   │     ║     │                    │   ║     │     ░░░░░░░░░░   │
│   ░░░░░░░░   │     ║     └────────────────────┘   ║     │     ░░░░░░░░░░   │
│   ░░░░░░░░   │     ║                              ║     │     ░░░░░░░░░░   │
│   ░░░░░░░░   │     ║    ✨ Personality Spark ✨   ║     │     ░░░░░░░░░░   │  Reveal Name
│   ░░░░░░░░   │     ║                              ║     │     ░░░░░░░░░░   │
│   ░░░░░░░░   │     ╚══════════════════════════════╝     │     ░░░░░░░░░░   │
│   ░░░░░░░░   │                                          │     ░░░░░░░░░░   │
│   ░░░░░░░░   │          [ Continue Streak ]             │     ░░░░░░░░░░   │  Action Button
│   ░░░░░░░░   │                                          │     ░░░░░░░░░░   │
│   ░░░░░░░░   └──────────────────────────────────────────┘     ░░░░░░░░░░   │
│   ░░░░░░░░                                                    ░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Layer Stack (Z-Index):
├── Z-100: Reveal Overlay (full screen, blocks all interaction below)
├── Z-101: Content Container (centered modal area)
├── Z-102: Reveal Card (elevated, glowing)
└── Z-103: Confetti/particles (Day 15 only, above everything)
```

### 5.2 Layout Tokens

| Property | Value |
|----------|-------|
| **Overlay Type** | Full-screen takeover |
| **Overlay Background** | `var(--reveal-bg-deep)` / `#0D0D0D` with radial gradient |
| **Z-Index** | `100` (above all app content) |
| **Content Area Width** | `min(360px, 90vw)` |
| **Content Area Padding** | `32px` |
| **Vertical Alignment** | Centered (flex, align-items: center) |

### 5.3 Background Treatment

```css
/* Radial gradient with subtle gold ambient glow */
.reveal-overlay {
  background: 
    radial-gradient(
      ellipse 60% 40% at 50% 30%,
      rgba(196, 167, 125, 0.06) 0%,
      transparent 70%
    ),
    linear-gradient(
      180deg,
      #0D0D0D 0%,
      #1A1A1A 100%
    );
}

/* Optional: Subtle animated shimmer for luxury feel */
.reveal-overlay::before {
  content: '';
  position: absolute;
  top: 0; left: 0; right: 0; bottom: 0;
  background: 
    linear-gradient(
      135deg,
      transparent 40%,
      rgba(196, 167, 125, 0.03) 50%,
      transparent 60%
    );
  animation: shimmer 8s infinite ease-in-out;
}
```

---

## 6. Component Specifications

### 6.1 Milestone Badge

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│                         ✦ DAY 5 UNLOCKED ✦                                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Specs:
├── Container: Inline pill
├── Background: transparent
├── Text: Outfit 12px / 600 / uppercase / letter-spacing: 0.15em
├── Color: var(--reveal-accent-gold) / #C4A77D
├── Stars (✦): Part of text, or decorative icons 8px flanking text
├── Margin Bottom: 16px
├── Animation: Fade in 0 → 100% (400ms, ease-out)
```

**Dynamic Badge Copy:**

| Milestone | Badge Text |
|-----------|------------|
| Reveal 1 | "✦ DAY {X} UNLOCKED ✦" |
| Reveal 2 | "✦ DAY {X} UNLOCKED ✦" |
| Reveal 3 | "✦ DAY {X} UNLOCKED ✦" |
| Reveal 4 | "✦ DAY {X} UNLOCKED ✦" |
| Reveal 5 | "✦ DAY 15 ACHIEVED ✦" |

### 6.2 Celebratory Header

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│               Trust earned. Here's a glimpse.                               │
│                                                                             │
│                    — or for Day 15 —                                        │
│                                                                             │
│      You did it. 15 days of showing up.                                     │
│           Here's who they really are.                                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Specs:
├── Font: Outfit (var(--font-display))
├── Size: 28px (var(--h1-size)) for main line
├── Weight: 700
├── Color: #FFFFFF
├── Alignment: Center
├── Line Height: 1.25
├── Max Width: 300px
├── Margin Bottom: 32px
├── Animation: Fade in + slide up from 20px (500ms, ease-out, 100ms delay)
```

### 6.3 Reveal Card — Locked State (State A)

The reveal card in its initial "wrapped" state, inviting the user to tap.

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║    ╭─────────────────────────────────────────────────────────────────────╮   ║
║    │                                                                     │   ║
║    │                                                                     │   ║
║    │                           🎁                                        │   ║   Gift icon: 56px
║    │                                                                     │   ║
║    │                                                                     │   ║
║    │                      Tap to Reveal                                  │   ║   Prompt text
║    │                                                                     │   ║
║    │                                                                     │   ║
║    ╰─────────────────────────────────────────────────────────────────────╯   ║
║                                                                              ║
║                       ✨ Personality Spark ✨                                ║   Reveal name
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝

Card Dimensions:
├── Width: 100% of content area (minus padding)
├── Height: Auto (min 200px)
├── Padding: 40px 24px
├── Border Radius: 20px

Card Styling (Locked):
├── Background: var(--reveal-bg-card) / #252525
├── Border: 1px solid var(--reveal-card-border) / rgba(196, 167, 125, 0.3)
├── Box Shadow: 
│   ├── 0 4px 24px rgba(0, 0, 0, 0.4)                    /* Depth */
│   └── 0 0 40px rgba(196, 167, 125, 0.08)              /* Gold glow */

Gift Icon:
├── Icon: Phosphor "Gift" or custom wrapped gift illustration
├── Size: 56px
├── Color: var(--reveal-accent-gold) / #C4A77D
├── Animation: Subtle float (up 4px, down 4px, 3s infinite)

Tap Prompt:
├── Font: Inter 16px / 500
├── Color: var(--reveal-text-secondary) / #E5E5E3
├── Margin Top: 20px

Reveal Name:
├── Font: Outfit 16px / 600
├── Color: var(--reveal-accent-gold) / #C4A77D
├── Spacing: Centered below card, 16px margin-top
├── Sparkles: Decorative emoji or icon flanking text

Interactive States:
├── Hover: Border brightens to rgba(196, 167, 125, 0.5)
├── Pressed: Scale 0.98, border brightens
├── Focus: Gold outline ring (accessibility)
```

### 6.4 Reveal Card — Revealed State (State B)

After the user taps, the content is revealed with a rich animation.

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║    ╭─────────────────────────────────────────────────────────────────────╮   ║
║    │                                                                     │   ║
║    │   ┌───────────────────────────────────────────────────────────┐     │   ║
║    │   │                                                           │     │   ║
║    │   │   🎂  Age          26                                     │     │   ║
║    │   │   📏  Height       5'6" - 5'8"                            │     │   ║
║    │   │   📍  Area         Indiranagar, Bangalore                 │     │   ║
║    │   │   🏋️  Hobby Depth  Serious (3+ years)                     │     │   ║
║    │   │   ✨  Personality  "Thoughtful listener"                  │     │   ║
║    │   │                                                           │     │   ║
║    │   └───────────────────────────────────────────────────────────┘     │   ║
║    │                                                                     │   ║
║    │         ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─                                  │   ║
║    │                                                                     │   ║
║    │   "They're 26, in the heart of Bangalore's café district,          │   ║   AI-Framed Text
║    │    and they've been at this for years. Not a hobby —               │   ║
║    │    a piece of who they are."                                       │   ║
║    │                                                                     │   ║
║    ╰─────────────────────────────────────────────────────────────────────╯   ║
║                                                                              ║
║                       ✨ Personality Spark ✨                                ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝

Card Styling (Revealed):
├── Background: var(--reveal-bg-card) / #252525
├── Border: 1px solid var(--reveal-accent-gold) at 50% opacity
├── Box Shadow enhanced: 
│   ├── 0 8px 32px rgba(0, 0, 0, 0.5)
│   └── 0 0 60px rgba(196, 167, 125, 0.12)              /* Stronger glow */
├── Height: Auto (expands to fit content)

Data Row Styling:
├── Icon: 18px, var(--reveal-accent-gold)
├── Label: Inter 13px / 500, var(--reveal-text-muted)
├── Value: Inter 15px / 500, #FFFFFF
├── Row Height: 32px
├── Row Gap: 8px
├── Divider: 1px dashed var(--reveal-text-muted) at 30%

AI-Framed Text:
├── Font: Inter 15px / 400 / italic
├── Color: var(--reveal-text-secondary) / #E5E5E3
├── Alignment: Center
├── Max Width: 280px
├── Margin Top: 16px
├── Quotes: Displayed (adds literary feel)

Reveal Animation:
├── Duration: 600ms
├── Easing: ease-out
├── Effect: 
│   ├── Gift icon fades out + scales down (0ms - 200ms)
│   ├── Content fades in from 0% opacity (200ms - 600ms)
│   └── Content scales from 0.95 → 1.0 (200ms - 600ms)
├── Haptic: .medium → .light ("CLICK" → "TICK")
```

### 6.5 Reveal Card — Day 15 (Identity + Photos)

Day 15 is a special reveal with photos, name, and chat unlock.

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║    ╭─────────────────────────────────────────────────────────────────────╮   ║
║    │                                                                     │   ║
║    │              ┌───────────────────────────┐                          │   ║
║    │              │                           │                          │   ║
║    │              │       [ PHOTO 1 ]         │                          │   ║   Primary Photo
║    │              │      (Face Photo)         │                          │   ║   150px × 150px
║    │              │                           │                          │   ║   Round corners
║    │              └───────────────────────────┘                          │   ║
║    │                                                                     │   ║
║    │    ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐               │   ║   Additional photos
║    │    │  📷 2   │  │  📷 3   │  │  📷 4   │  │  📷 5   │               │   ║   60px × 60px each
║    │    └─────────┘  └─────────┘  └─────────┘  └─────────┘               │   ║
║    │                                                                     │   ║
║    │                     ─ ─ ─ ─ ─ ─ ─                                   │   ║
║    │                                                                     │   ║
║    │                     Priya Sharma                                    │   ║   Full Name: H2
║    │                                                                     │   ║
║    │                 "After 15 days, you've earned                       │   ║   AI-Framed
║    │                  the name behind the dedication."                   │   ║
║    │                                                                     │   ║
║    ╰─────────────────────────────────────────────────────────────────────╯   ║
║                                                                              ║
║                      🎉 Identity Revealed 🎉                                ║
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝

Day 15 Special Styling:
├── Primary Photo: 150px × 150px, border-radius 16px
├── Photo Border: 3px solid var(--reveal-accent-gold)
├── Photo Shadow: 0 0 30px rgba(196, 167, 125, 0.3)
├── Secondary Photos: Horizontal scroll or grid, 60px each
├── Name: Outfit 24px / 700 / #FFFFFF, centered
├── Confetti: 300ms particle burst on reveal (optional)
```

### 6.6 Action Button — Continue Streak / Open Chat

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                                                                     │   │
│   │                      Continue Streak                                │   │   Reveals 1-4
│   │                                                                     │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                                                                     │   │
│   │                   💬  Start Chatting                                │   │   Day 15 Only
│   │                                                                     │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Reveals 1-4 Button:
├── Text: "Continue Streak"
├── Background: var(--reveal-accent-gold) / #C4A77D
├── Text Color: var(--reveal-bg-surface) / #1A1A1A (dark on gold)
├── Font: Inter 16px / 600
├── Height: 52px
├── Width: 100%
├── Radius: 12px
├── Margin Top: 32px
├── Action: Dismiss overlay → Return to Streak Screen

Day 15 Button:
├── Text: "💬 Start Chatting"
├── Background: var(--reveal-accent-gold) / #C4A77D
├── Text Color: var(--reveal-bg-surface) / #1A1A1A
├── Icon: ChatCircle, 20px, left of text
├── Action: Dismiss overlay → Navigate to Chat Screen
```

---

## 7. State Definitions

### 7.1 State A — Locked ("Tap to Reveal")

**Trigger:** User reaches milestone day, reveal event screen opens.

```
┌─────────────────────────────────────────────────────────────────┐
│  STATE A: LOCKED                                                │
│  The reveal is earned but not yet opened                        │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │
│  ░░░░░░░░░░░░░░░░ DEEP DARK BACKDROP ░░░░░░░░░░░░░░░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░            ✦ DAY 5 UNLOCKED ✦            ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░      Trust earned. Here's a glimpse.     ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░     ╔════════════════════════════════╗   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║                                ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║            🎁                  ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║                                ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║        Tap to Reveal           ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║                                ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ╚════════════════════════════════╝   ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░          ✨ Personality Spark ✨         ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

Components visible:
├── Milestone Badge: "✦ DAY X UNLOCKED ✦"
├── Celebratory Header: "Trust earned. Here's a glimpse."
├── Reveal Card (Locked): Gift icon + "Tap to Reveal" prompt
├── Reveal Name: "✨ [Reveal Name] ✨"
├── ⚠️ NO action button visible yet (appears after reveal)

Interaction:
├── Tap on card → Transition to State B (Revealed)
├── Cannot dismiss without tapping card
```

### 7.2 State B — Revealed (Content Visible)

**Trigger:** User taps the locked reveal card.

```
┌─────────────────────────────────────────────────────────────────┐
│  STATE B: REVEALED                                              │
│  Content is now visible, user can proceed                       │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │
│  ░░░░░░░░░░░░░░░░ DEEP DARK BACKDROP ░░░░░░░░░░░░░░░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░            ✦ DAY 5 UNLOCKED ✦            ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░      Trust earned. Here's a glimpse.     ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░     ╔════════════════════════════════╗   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║  🎂 Age         26             ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║  📏 Height      5'6" - 5'8"    ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║  📍 Area        Indiranagar    ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║  🏋️ Hobby       Serious        ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║  ✨ Personality "Thoughtful"   ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║                                ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ║  "They're 26, in the heart..." ║   ░░░░░░░░░░░  │
│  ░░░░░░░░     ╚════════════════════════════════╝   ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░          ✨ Personality Spark ✨         ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░     [       Continue Streak        ]     ░░░░░░░░░░░  │
│  ░░░░░░░░                                          ░░░░░░░░░░░  │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

Components visible:
├── Milestone Badge (unchanged)
├── Celebratory Header (unchanged)
├── Reveal Card (Revealed): Data rows + AI-framed text
├── Reveal Name (unchanged)
├── Action Button: "Continue Streak" (gold, full-width)

Interaction:
├── Tap "Continue Streak" → Dismiss overlay → Return to Streak Screen
├── Card is now non-interactive (read-only)
```

---

## 8. Reveal-Specific Content Templates

### 8.1 Reveal 1 — Personality Spark

| Data Field | Example Value | Icon |
|------------|---------------|------|
| Age | 26 | 🎂 |
| Height | 5'6" - 5'8" | 📏 |
| City Area | Indiranagar, Bangalore | 📍 |
| Hobby Depth | Serious (3+ years) | 🏋️ |
| Personality Signal | "Thoughtful listener" | ✨ |

**AI-Framed Example:**
> "They're 26, in the heart of Bangalore's café district, and they've been at this for years. Not a hobby — a piece of who they are."

---

### 8.2 Reveal 2 — Lifestyle Reality

| Data Field | Example Value | Icon |
|------------|---------------|------|
| Profession | Engineering / Technology | 💼 |
| Education | Master's Degree | 🎓 |
| Body Type (optional) | Athletic | 🧘 |
| Routine | Morning person | ☀️ |

**AI-Framed Example:**
> "They build things that move — in a city that wakes up to filter coffee. Mornings are their time."

---

### 8.3 Reveal 3 — Social & Emotional Context

| Data Field | Example Value | Icon |
|------------|---------------|------|
| Religion (opt-in) | Hindu | 🕉️ |
| Cultural Background | South Indian | 🌏 |
| Languages | English, Tamil, Hindi | 💬 |
| Emotional Signal | "Values quality time" | 💜 |

**AI-Framed Example:**
> "Rooted in tradition, fluent in three languages, and someone who shows up when it matters."

---

### 8.4 Reveal 4 — Human Presence

| Data Field | Example Value | Icon |
|------------|---------------|------|
| Voice Note | [Playable audio] | 🎤 |
| Candid Photo | [Non-face image] | 📸 |
| Immutable Fact | "Left-handed" | 🤚 |

**AI-Framed Example:**
> "Their voice carries warmth. This is who they sound like — real, unfiltered."

**Special UI Elements:**
- Voice note: Waveform visualization, play button, duration indicator
- Candid photo: Image with subtle border, rounded corners

---

### 8.5 Reveal 5 — Identity + Chat Unlock (Day 15)

| Data Field | Example Value | Icon |
|------------|---------------|------|
| Full Name | Priya Sharma | — |
| Face Photos | [3-6 photos] | — |
| Chat | Unlocked | 💬 |

**Celebratory Header (Day 15):**
> "You did it. 15 days of showing up. Here's who they really are."

**AI-Framed Example:**
> "After 15 days of consistency, this is the name behind the dedication. Say hello to Priya."

**Special UI:**
- Confetti particle burst (300ms)
- Photos displayed prominently
- Action button: "💬 Start Chatting"

---

## 9. Copy Table

### 9.1 Headers by Milestone

| Reveal | Celebratory Header |
|--------|-------------------|
| R1 | "Trust earned. Here's a glimpse." |
| R2 | "Consistency pays off. Meet their world." |
| R3 | "Roots and rhythms. Here's who they are." |
| R4 | "Real presence. Hear and see." |
| R5 | "You did it. 15 days of showing up. Here's who they really are." |

### 9.2 Badge Copy

| Reveal | Badge Text |
|--------|-----------|
| R1-R4 | "✦ DAY {X} UNLOCKED ✦" |
| R5 | "✦ DAY 15 ACHIEVED ✦" |

### 9.3 Locked Card Prompt

| Reveal | Prompt |
|--------|--------|
| All | "Tap to Reveal" |
| (Alternative) | "Unwrap your reward" |

### 9.4 Reveal Names

| Reveal | Name |
|--------|------|
| R1 | "✨ Personality Spark ✨" |
| R2 | "✨ Lifestyle Reality ✨" |
| R3 | "✨ Social & Emotional Context ✨" |
| R4 | "✨ Human Presence ✨" |
| R5 | "🎉 Identity Revealed 🎉" |

### 9.5 Action Button Copy

| Reveal | Button Text |
|--------|-------------|
| R1-R4 | "Continue Streak" |
| R5 | "💬 Start Chatting" |

---

## 10. Animation Sequences

### 10.1 Screen Entry Animation

```
Timeline: Reveal Screen Opens
──────────────────────────────────────────────────────────────────

[0ms - 300ms]    OVERLAY FADE IN
                 ├── Background fades from transparent → visible
                 ├── Easing: ease-out
                 └── Subtle radial gradient builds

[100ms - 500ms]  BADGE FADE IN
                 ├── Opacity: 0 → 100%
                 ├── Transform: none (static position)
                 └── Easing: ease-out

[200ms - 700ms]  HEADER SLIDE + FADE
                 ├── Opacity: 0 → 100%
                 ├── Transform: translateY(20px) → translateY(0)
                 └── Easing: ease-out

[400ms - 900ms]  REVEAL CARD ENTRY
                 ├── Opacity: 0 → 100%
                 ├── Transform: scale(0.95) → scale(1.0)
                 ├── Box shadow builds progressively
                 └── Easing: ease-out

[700ms - 1000ms] REVEAL NAME FADE
                 └── Opacity: 0 → 100%
```

### 10.2 Reveal Animation (Tap to Reveal)

```
Timeline: User Taps Card
──────────────────────────────────────────────────────────────────

[0ms]            TAP DETECTED
                 ├── Haptic: .medium (iOS) / CLICK (Android)
                 └── Card scales to 0.98

[0ms - 200ms]    GIFT ICON FADE OUT
                 ├── Opacity: 100% → 0%
                 ├── Transform: scale(1.0) → scale(0.8)
                 └── "Tap to Reveal" text fades

[150ms - 200ms]  CARD HEIGHT TRANSITION
                 ├── Height animates to accommodate content
                 └── Easing: ease-out

[200ms - 600ms]  CONTENT FADE IN
                 ├── Opacity: 0 → 100%
                 ├── Transform: scale(0.95) → scale(1.0)
                 ├── Staggered row reveal (optional):
                 │   ├── Row 1: delay 200ms
                 │   ├── Row 2: delay 250ms
                 │   └── (50ms stagger per row)
                 └── Easing: ease-out

[600ms - 650ms]  HAPTIC: .light (iOS) / TICK (Android)

[600ms - 900ms]  ACTION BUTTON SLIDE UP
                 ├── Opacity: 0 → 100%
                 ├── Transform: translateY(20px) → translateY(0)
                 └── Easing: ease-out
```

### 10.3 Day 15 Confetti Burst

```
Timeline: Day 15 Reveal Only
──────────────────────────────────────────────────────────────────

[600ms]          CONFETTI TRIGGER (after content reveals)
                 ├── Particle count: 50-80
                 ├── Colors: Gold (#C4A77D), White, subtle server color
                 ├── Origin: Center-top of card
                 ├── Spread: 120° arc
                 ├── Duration: 2000ms (particles fade and fall)
                 └── Physics: Gravity + air resistance

[600ms]          DOUBLE HAPTIC: .heavy × 2 (iOS) / HEAVY_CLICK × 2 (Android)
```

---

## 11. Haptic Feedback

| Action | iOS | Android | Timing |
|--------|-----|---------|--------|
| Screen opens | `.light` | `TICK` (20ms) | On overlay fade start |
| Tap reveal card | `.medium` | `CLICK` (40ms) | Immediate on tap |
| Content revealed | `.light` | `TICK` (20ms) | After 600ms animation |
| Day 15 celebration | `.heavy` × 2 | `HEAVY_CLICK` × 2 | With confetti burst |
| Tap action button | `.success` | `CONFIRM` (50ms) | On button press |

---

## 12. Accessibility

### 12.1 Screen Reader Announcements

| State | Announcement |
|-------|--------------|
| **Screen Opens (Locked)** | "Milestone unlocked. Day 5. Personality Spark reveal available. Tap to reveal." |
| **Content Revealed** | "Reveal complete. Age: 26. Height: 5'6" to 5'8". Area: Indiranagar, Bangalore. Hobby depth: Serious. Personality: Thoughtful listener. Continue streak button available." |
| **Day 15** | "Day 15 achieved! Identity revealed. Name: Priya Sharma. Photos available. Chat now unlocked. Start chatting button available." |

### 12.2 Focus Management

1. On screen open, focus moves to **Reveal Card** (tappable element)
2. After reveal, focus moves to **Action Button**
3. Focus is trapped within overlay until dismissed

### 12.3 Motion Sensitivity

When `prefers-reduced-motion` is enabled:
- No confetti animation
- Instant transitions (no fade/scale)
- Content appears immediately on tap
- Haptics still fire

---

## 13. Technical Notes

### 13.1 API Requirements

```typescript
interface RevealEventData {
  userId: string;
  connectionId: string;
  streakDay: number;              // The day that triggered the reveal
  revealNumber: 1 | 2 | 3 | 4 | 5;
  revealName: string;             // "Personality Spark", etc.
  revealContent: RevealContent;   // Varies by reveal type
  aiFramedText: string;           // Pre-generated by AI
  partnerName?: string;           // Only for Reveal 5
  partnerPhotos?: string[];       // Only for Reveal 5
}

interface RevealContent {
  fields: Array<{
    icon: string;
    label: string;
    value: string;
  }>;
  voiceNoteUrl?: string;          // Reveal 4 only
  candidPhotoUrl?: string;        // Reveal 4 only
}

interface RevealAction {
  type: 'continue' | 'start_chat';
  connectionId: string;
  revealNumber: number;
}
```

### 13.2 Analytics Events

| Event | Trigger | Properties |
|-------|---------|------------|
| `reveal_event_opened` | Screen opens | `reveal_number`, `streak_day`, `user_tier` |
| `reveal_tapped` | User taps "Tap to Reveal" | `reveal_number`, `time_to_tap_ms` |
| `reveal_viewed` | Content fully visible | `reveal_number`, `animation_complete` |
| `reveal_continued` | User taps "Continue Streak" | `reveal_number` |
| `chat_initiated` | User taps "Start Chatting" (Day 15) | — |

### 13.3 Error Handling

| Error Case | UI Response |
|------------|-------------|
| Reveal data fails to load | Show error state with retry, do not dismiss overlay |
| Voice note fails to play | Show inline error, offer retry |
| Photo fails to load | Show placeholder with retry option |

---

## 14. Related Screens

| Screen | Relationship |
|--------|--------------|
| [Screen 12: Streak View](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_12_StreakView.md) | Source screen — Reveal Event overlays this |
| [Screen 18: Chat Interface](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_18_ChatInterface.md) | Navigated to after Day 15 "Start Chatting" |
| [Screen 13: Reveal History](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_13_RevealHistory.md) | Contains record of all revealed information |

---

**Document maintained by:** Product Design Team  
**Last updated:** January 2026  
**Review cycle:** With each PRD/DSD update
