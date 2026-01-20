# Unora — UI Specification: Screen #19

## Identity Reveal (Day 15)

**Version:** 1.0  
**Last Updated:** January 2026  
**Status:** Final  
**Author:** Product Design Team

---

## 1. Document Metadata

| Attribute | Value |
|-----------|-------|
| **Screen Name** | Identity Reveal (Day 15) |
| **Screen ID** | `SCREEN_19_IDENTITY_REVEAL` |
| **Spec ID** | Spec-19 |
| **User Flow Reference** | [Section 2.E.3 — Day 15: Identity Reveal + Chat Unlock](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) |
| **PRD Reference** | [Section 15.2 — Reveal 5: Identity + Chat Unlock](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md), [Section 15.3 — Reveal Timing by Tier](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) |
| **DSD Reference** | [Section 12 — Premium Dark Visual Language](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) |
| **Trigger** | Streak counter increments to Day 15 |
| **Visibility** | Both users in the connection (simultaneously) |

---

## 2. The Grand Finale — Experience Philosophy

### 2.1 The Moment of Truth

This is **the most important screen in Unora**. 

After 15 days of mutual consistency, anonymous connection transforms into known identity. This is the culmination of the Unora philosophy: **earned trust becomes real connection**.

**This screen must feel:**

| Quality | Implementation |
|---------|----------------|
| **Cinematic** | Full-screen takeover, dramatic lighting, gold accents |
| **Earned** | Copy emphasizes the 15-day journey, not just the outcome |
| **Premium** | Dark, rich surfaces with golden warmth |
| **Revelatory** | Animation dissolves the abstract into the real |
| **Celebratory** | Confetti, haptics, emotional copy |
| **Transitional** | Clear call-to-action to begin chatting |

> [!IMPORTANT]
> This is NOT a notification or a card update. This is a **full-screen cinematic event** that demands attention and celebrates the milestone.

### 2.2 Critical Logic Constraints

> [!CAUTION]
> These rules are **absolute and non-negotiable**.

| Constraint | Rule |
|------------|------|
| **Day 15 is UNIVERSAL** | All tiers reach identity reveal at Day 15. No exceptions. |
| **NOT Accelerable** | No tier or payment can speed up Day 15. It cannot be purchased. |
| **NOT Purchasable** | This reveal is earned only. Money cannot buy it. |
| **Simultaneous Unlock** | Both users see this screen at the same time (within the check-in window). |
| **Mutual Achievement** | Both users contributed equally (7.5 check-ins each = 15 total days). |
| **Data from Verification** | Full legal name is fetched from government verification, not user input. |

### 2.3 What Gets Revealed

| Data | Source | Display |
|------|--------|---------|
| **Full Legal Name** | User-provided (onboarding) | H1 headline text |
| **Face Photos** | User-uploaded during onboarding | Hero image gallery |
| **Chat Functionality** | System unlock | "Say Hello" button becomes active |

---

## 3. Visual Theme & Styling

### 3.1 Premium Dark Cinematic Foundation

This screen uses the **deepest, most premium dark mode** styling in the entire app.

```css
/* Cinematic Dark Foundation */
--pdm-background: #0D0D0F;        /* Deepest layer — almost black */
--pdm-surface-1: #141416;         /* Base layer */
--pdm-surface-2: #1A1A1E;         /* Elevated elements */

/* Text — Light on Dark, Maximum Contrast */
--pdm-ink-primary: #FFFFFF;       /* Pure white headlines */
--pdm-ink-secondary: #E5E5E3;     /* Off-white body */
--pdm-ink-tertiary: #8A8A86;      /* Muted captions */

/* Gold Accents — Celebratory Warmth */
--accent-gold: #C4A77D;           /* Primary accent */
--accent-gold-bright: #D4B88D;    /* Highlight state */
--accent-gold-deep: #B08D5B;      /* Deep gold for gradients */
--accent-gold-glow: rgba(196, 167, 125, 0.25);  /* Outer glow */
--accent-gold-ambient: rgba(196, 167, 125, 0.08); /* Background warmth */

/* Celebratory Spotlight Effect */
--spotlight-gradient: radial-gradient(
  ellipse 50% 35% at 50% 30%,
  rgba(196, 167, 125, 0.15) 0%,
  rgba(196, 167, 125, 0.05) 40%,
  transparent 70%
);
```

### 3.2 Background Treatment

The background is not flat — it has a **subtle animated golden spotlight** that creates depth and draws attention to the hero image.

```css
.identity-reveal-screen {
  background: 
    /* Spotlight effect behind photo */
    radial-gradient(
      ellipse 60% 40% at 50% 35%,
      rgba(196, 167, 125, 0.12) 0%,
      rgba(196, 167, 125, 0.04) 50%,
      transparent 80%
    ),
    /* Base gradient */
    linear-gradient(
      180deg,
      #0D0D0F 0%,
      #121214 50%,
      #0D0D0F 100%
    );
}

/* Optional: Subtle animated shimmer */
.identity-reveal-screen::before {
  content: '';
  position: absolute;
  inset: 0;
  background: 
    radial-gradient(
      circle at 50% 30%,
      rgba(196, 167, 125, 0.06) 0%,
      transparent 50%
    );
  animation: spotlight-pulse 4s ease-in-out infinite;
}

@keyframes spotlight-pulse {
  0%, 100% { opacity: 0.6; transform: scale(1); }
  50% { opacity: 1; transform: scale(1.05); }
}
```

### 3.3 Typography

| Element | Font | Size | Weight | Color |
|---------|------|------|--------|-------|
| **Day Badge** | Outfit | 14px | 600 | --accent-gold |
| **Celebration Header** | Outfit | 32px | 700 | #FFFFFF |
| **Partner Name (H1)** | Outfit | 36px | 700 | #FFFFFF |
| **"Verified Identity" Badge** | Inter | 12px | 500 | --accent-gold |
| **Journey Text** | Inter | 16px | 400 | --pdm-ink-secondary |
| **CTA Button** | Inter | 18px | 600 | #1A1A1A |
| **Caption** | Inter | 13px | 400 | --pdm-ink-tertiary |

---

## 4. Layout Specification

### 4.1 Full-Screen Immersive Overlay

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░ PREMIUM DARK BACKGROUND ░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░ #0D0D0F ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░                                                  ░░░░░░░░░░░   │
│   ░░░░░░░░░░              ✦ DAY 15 ACHIEVED ✦                 ░░░░░░░░░░░   │   Day Badge
│   ░░░░░░░░░░                                                  ░░░░░░░░░░░   │
│   ░░░░░░░░░░                                                  ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╭ ─ ─ ─ ─ ─ GOLD SPOTLIGHT ─ ─ ─ ─ ─ ╮       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎                                     ╎       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎    ╔═══════════════════════════╗    ╎       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎    ║                           ║    ╎       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎    ║                           ║    ╎       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎    ║      [ HERO PHOTO ]       ║    ╎       ░░░░░░░░░░░   │   Hero Image
│   ░░░░░░░░░░    ╎    ║     (Face Revealed)       ║    ╎       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎    ║                           ║    ╎       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎    ║                           ║    ╎       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎    ╚═══════════════════════════╝    ╎       ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ╎              ◉ ○ ○ ○                ╎       ░░░░░░░░░░░   │   Photo Dots
│   ░░░░░░░░░░    ╰ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ╯       ░░░░░░░░░░░   │
│   ░░░░░░░░░░                                                  ░░░░░░░░░░░   │
│   ░░░░░░░░░░                Priya Sharma                      ░░░░░░░░░░░   │   Name (H1)
│   ░░░░░░░░░░             ✓ Verified Identity                  ░░░░░░░░░░░   │   Verified Badge
│   ░░░░░░░░░░                                                  ░░░░░░░░░░░   │
│   ░░░░░░░░░░     You did it. 15 days of showing up.           ░░░░░░░░░░░   │   Journey Text
│   ░░░░░░░░░░          Here's who they really are.             ░░░░░░░░░░░   │
│   ░░░░░░░░░░                                                  ░░░░░░░░░░░   │
│   ░░░░░░░░░░    ┌─────────────────────────────────────────┐   ░░░░░░░░░░░   │
│   ░░░░░░░░░░    │ ░░░░░░░░░░░ 💬 Say Hello ░░░░░░░░░░░░░░ │   ░░░░░░░░░░░   │   Primary CTA
│   ░░░░░░░░░░    └─────────────────────────────────────────┘   ░░░░░░░░░░░   │   (Gold Gradient)
│   ░░░░░░░░░░                                                  ░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│   ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│                                                                             │
│   ─────────────────── 🎊 CONFETTI LAYER 🎊 ───────────────────              │   Z-110
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Layer Stack (Z-Index):
├── Z-100: Full-screen overlay (blocks all interaction below)
├── Z-101: Background with spotlight gradient
├── Z-102: Content (photo, name, button)
└── Z-110: Confetti particles (top layer)
```

### 4.2 Layout Tokens

| Property | Value |
|----------|-------|
| **Screen Type** | Full-screen immersive overlay |
| **Background** | `var(--pdm-background)` with spotlight gradient |
| **Z-Index** | 100 (above all app content) |
| **Content Alignment** | Center (flex, justify-content: center, align-items: center) |
| **Content Max Width** | 360px |
| **Horizontal Padding** | 24px |

### 4.3 Vertical Spacing

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   [Safe Area Top]                                                           │
│                                                                             │
│   [Variable Flex Space — pushes content to vertical center]                 │
│                                                                             │
│           ✦ DAY 15 ACHIEVED ✦                                               │
│                                                                             │
│   [24px]                                                                    │
│                                                                             │
│           ╔═══════════════════════════════════════╗                         │
│           ║                                       ║                         │
│           ║           [ HERO PHOTO ]              ║    280px × 280px        │
│           ║                                       ║                         │
│           ╚═══════════════════════════════════════╝                         │
│                         ◉ ○ ○ ○                                             │
│                                                                             │
│   [24px]                                                                    │
│                                                                             │
│                      Priya Sharma                                           │   36px H1
│                   ✓ Verified Identity                                       │   12px badge
│                                                                             │
│   [16px]                                                                    │
│                                                                             │
│          You did it. 15 days of showing up.                                 │   16px
│              Here's who they really are.                                    │   16px
│                                                                             │
│   [32px]                                                                    │
│                                                                             │
│           ┌─────────────────────────────────────────┐                       │
│           │           💬 Say Hello                  │    56px height        │
│           └─────────────────────────────────────────┘                       │
│                                                                             │
│   [Variable Flex Space]                                                     │
│                                                                             │
│   [Safe Area Bottom]                                                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 5. Component Specifications

### 5.1 Day Badge

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│                          ✦ DAY 15 ACHIEVED ✦                                │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Specs:
├── Text: "✦ DAY 15 ACHIEVED ✦"
├── Font: Outfit 14px / 600 / uppercase / letter-spacing: 0.15em
├── Color: var(--accent-gold) / #C4A77D
├── Alignment: Center
├── Stars (✦): Character or decorative icons, 10px
├── Animation: Fade in + slight glow pulse (continuous)
```

### 5.2 Hero Photo Container

The hero moment — the face revealed for the first time.

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                                                                              ║
║      ╭ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ╮           ║
║      ╎                    GOLD GLOW AURA                        ╎           ║
║      ╎                                                          ╎           ║
║      ╎    ╔════════════════════════════════════════════════╗    ╎           ║
║      ╎    ║                                                ║    ╎           ║
║      ╎    ║                                                ║    ╎           ║
║      ╎    ║                                                ║    ╎           ║
║      ╎    ║                                                ║    ╎           ║
║      ╎    ║             [ FACE PHOTO ]                     ║    ╎           ║
║      ╎    ║                                                ║    ╎           ║
║      ╎    ║                                                ║    ╎           ║
║      ╎    ║                                                ║    ╎           ║
║      ╎    ║                                                ║    ╎           ║
║      ╎    ╚════════════════════════════════════════════════╝    ╎           ║
║      ╎                                                          ╎           ║
║      ╰ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ╯           ║
║                                                                              ║
║                              ◉   ○   ○   ○                                  ║   Swipe dots
║                                                                              ║
╚══════════════════════════════════════════════════════════════════════════════╝

Photo Container Specs:
├── Size: 280px × 280px (square, aspect-ratio: 1)
├── Border Radius: 24px
├── Object Fit: cover
├── Border: 3px solid var(--accent-gold) / #C4A77D
├── Box Shadow:
│   ├── 0 8px 32px rgba(0, 0, 0, 0.4)              /* Depth shadow */
│   ├── 0 0 60px var(--accent-gold-glow)          /* Gold outer glow */
│   └── inset 0 0 0 1px rgba(255, 255, 255, 0.1)  /* Inner light edge */

Gold Glow Aura (Behind Photo):
├── Size: 320px × 320px (extends 20px beyond photo)
├── Background: radial-gradient(circle, var(--accent-gold-glow) 0%, transparent 70%)
├── Blur: filter: blur(40px)
├── Animation: Subtle pulse (scale 1.0 → 1.02 → 1.0, 3s infinite)
├── Z-Index: -1 (behind photo)

Photo Indicator Dots:
├── Count: Matches number of photos (typically 3-6)
├── Dot Size: 8px
├── Gap: 12px
├── Active Dot: var(--accent-gold), filled
├── Inactive Dots: rgba(255, 255, 255, 0.3), filled
├── Margin Top: 16px
```

### 5.3 Photo Gallery Interaction

Users can swipe through multiple face photos.

```
SWIPE BEHAVIOR:
├── Horizontal swipe: Navigate between photos
├── Snap-to-center: Photos snap to center position
├── Velocity-based: Fast swipe advances even small distances
├── Dot indicator: Updates to show current photo
├── Loop: NO (linear navigation, first to last only)

GESTURE SPECS:
├── Threshold: 50px horizontal movement to trigger change
├── Animation: 300ms ease-out slide
├── Haptic: .light on photo change
```

### 5.4 Identity Block (Name + Verified Badge)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│                           Priya Sharma                                      │   Name (H1)
│                                                                             │
│                       ✓ Verified Identity                                   │   Badge
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Name:
├── Text: "[Full Legal Name]" (from verification)
├── Font: Outfit 36px / 700
├── Color: #FFFFFF
├── Alignment: Center
├── Max Width: 320px
├── Truncation: None (name always fits, can wrap to 2 lines)

Verified Badge:
├── Icon: Phosphor "CheckCircle" or "SealCheck", 16px
├── Text: "Verified Identity"
├── Font: Inter 12px / 500 / uppercase / letter-spacing: 0.1em
├── Color: var(--accent-gold) / #C4A77D
├── Gap (icon to text): 6px
├── Margin Top: 8px
├── Layout: Inline-flex, centered
```

### 5.5 Journey Text (Celebration Copy)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│               You did it. 15 days of showing up.                            │
│                   Here's who they really are.                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Specs:
├── Text: "You did it. 15 days of showing up.\nHere's who they really are."
├── Font: Inter 16px / 400
├── Color: var(--pdm-ink-secondary) / #E5E5E3
├── Line Height: 1.5
├── Alignment: Center
├── Max Width: 280px

Alternative (personalized):
├── Text: "You did it. 15 days of showing up.\nMeet [FirstName]."
```

### 5.6 Primary CTA — "Say Hello"

The button that unlocks the chat and begins the new chapter.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │   │
│   │ ░░░░░░░░░░░░░░░░░░░░ 💬 Say Hello ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │   │
│   │ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                    ╲ ░ ░ ░ gold glow ░ ░ ░ ╱                                │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Button Specs:
├── Width: 100% (max 320px)
├── Height: 56px
├── Background: linear-gradient(135deg, var(--accent-gold-deep) 0%, var(--accent-gold) 100%)
│               (#B08D5B → #C4A77D)
├── Border: 1px solid rgba(255, 255, 255, 0.15)
├── Border Radius: 16px
├── Box Shadow:
│   ├── 0 4px 16px rgba(0, 0, 0, 0.3)
│   └── 0 0 32px var(--accent-gold-glow)

Text & Icon:
├── Icon: Phosphor "ChatCircle", 22px, left of text
├── Text: "Say Hello"
├── Font: Inter 18px / 600
├── Color: #1A1A1A (dark on gold)
├── Gap: 10px

States:
├── Default: As specified
├── Hover: Glow intensity +30%, slight lift (translateY(-2px))
├── Pressed: scale(0.97), glow reduces slightly
├── Animation: Subtle continuous glow pulse when idle
```

---

## 6. Confetti Celebration

### 6.1 Confetti Particle System

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   🎊              ✨           🎉            ✦            🎊               │
│         ★                            ▪                    ✨                │
│              ▫        🎊        ★              ▪                            │
│    ✦                       ✨                       ▫            🎉         │
│         ▪         ★                  🎊                    ✦                │
│                                                                             │
│   (Particles fall with gravity + air resistance)                            │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Confetti Specs:
├── Trigger: Auto-play on screen entry (after reveal animation)
├── Particle Count: 80-120
├── Colors:
│   ├── var(--accent-gold) / #C4A77D (40%)
│   ├── var(--accent-gold-bright) / #D4B88D (30%)
│   ├── #FFFFFF (20%)
│   └── var(--accent-sand) / #E8DED5 (10%)
├── Shapes: Rectangles (4px × 8px), Squares (6px), Circles (4px)
├── Origin: Full width, top of screen
├── Spread: 180° arc (full horizontal)
├── Initial Velocity: Random upward + outward
├── Gravity: 0.5 (gentle fall)
├── Duration: 3000ms (particles fade and fall)
├── Fade: Particles fade to 0 opacity over last 1000ms

CSS Reference (using canvas or CSS animation):

@keyframes confetti-fall {
  0% {
    transform: translateY(-100vh) rotate(0deg);
    opacity: 1;
  }
  70% {
    opacity: 1;
  }
  100% {
    transform: translateY(100vh) rotate(720deg);
    opacity: 0;
  }
}

Accessibility:
├── prefers-reduced-motion: Disable confetti entirely
├── Confetti is decorative only (no interaction)
```

---

## 7. Reveal Animation Sequence

### 7.1 Entry Animation — Abstract to Identity

The most important animation in the app. The photo "dissolves" from blurred/abstract to clear.

```
Timeline: Screen Opens
══════════════════════════════════════════════════════════════════════════════

[0ms]            SCREEN FADE IN
                 ├── Background fades from black → visible
                 └── Duration: 400ms, ease-out

[200ms]          DAY BADGE FADE IN
                 ├── Opacity: 0 → 100%
                 ├── Transform: translateY(-10px) → translateY(0)
                 └── Duration: 300ms, ease-out

[400ms]          GOLD AURA FADE IN
                 ├── Opacity: 0 → 100%
                 ├── Scale: 0.8 → 1.0
                 └── Duration: 500ms, ease-out

[500ms - 1500ms] ⭐ PHOTO REVEAL (THE HERO MOMENT) ⭐
                 │
                 ├── PHASE 1: Abstract State (0ms)
                 │   ├── Shows abstract avatar (gradient + minimal face outline)
                 │   └── Same dimensions as final photo
                 │
                 ├── PHASE 2: Dissolve Transition (500ms - 1200ms)
                 │   ├── Abstract avatar scales to 0.98
                 │   ├── Blur filter: 0 → 20px → 0 (gaussian blur)
                 │   ├── Opacity crossfade: abstract 100% → 0%, photo 0% → 100%
                 │   ├── Golden particles briefly emanate from edges
                 │   └── Haptic: .medium at transition midpoint
                 │
                 └── PHASE 3: Revealed State (1200ms+)
                     ├── Photo fully visible, crisp, no blur
                     ├── Scale: 0.98 → 1.0 (subtle settle)
                     └── Gold border glow intensifies

[1300ms]         HAPTIC: .success (double tap pattern)

[1400ms]         NAME FADE IN
                 ├── Opacity: 0 → 100%
                 ├── Transform: translateY(15px) → translateY(0)
                 └── Duration: 400ms, ease-out

[1500ms]         VERIFIED BADGE FADE IN
                 └── Duration: 200ms

[1600ms]         JOURNEY TEXT FADE IN
                 ├── Opacity: 0 → 100%
                 └── Duration: 300ms

[1800ms]         CONFETTI BURST 🎊
                 ├── Particles launch from top
                 └── Duration: 3000ms total

[2000ms]         CTA BUTTON FADE IN
                 ├── Opacity: 0 → 100%
                 ├── Transform: translateY(20px) → translateY(0)
                 ├── Glow begins pulsing
                 └── Duration: 400ms, ease-out
```

### 7.2 Photo Flip Alternative

An alternative "flip" animation instead of dissolve:

```
FLIP ANIMATION (Alternative):
├── Card rotates 180° on Y-axis
├── At 90° (edge-on), content swaps from abstract → photo
├── First half: Abstract visible
├── Second half: Photo visible
├── Duration: 800ms
├── Easing: ease-in-out
├── 3D perspective: 1000px

CSS:
.photo-container { perspective: 1000px; }
.photo-card {
  transform-style: preserve-3d;
  animation: card-flip 800ms ease-in-out forwards;
}
@keyframes card-flip {
  0% { transform: rotateY(0deg); }
  100% { transform: rotateY(180deg); }
}
```

---

## 8. State Definitions

### 8.1 State A — Reveal Animation Playing

Initial state when screen opens.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  STATE A: REVEAL ANIMATION PLAYING                                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  User cannot interact — animation is auto-playing                           │
│                                                                             │
│  ░░░░░░░░░░░░░░░░░░░░ PREMIUM DARK BACKGROUND ░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│                                                                             │
│                        ✦ DAY 15 ACHIEVED ✦                                  │
│                                                                             │
│                ╔═════════════════════════════════╗                          │
│                ║                                 ║                          │
│                ║      [ ABSTRACT → PHOTO ]       ║   ← Transition playing   │
│                ║         (dissolving)            ║                          │
│                ║                                 ║                          │
│                ╚═════════════════════════════════╝                          │
│                                                                             │
│                        (name fading in...)                                  │
│                                                                             │
│                    (button not visible yet)                                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Duration: ~2000ms
Interactivity: None (user watches)
```

### 8.2 State B — Identity Revealed (Interactive)

Animation complete, user can interact.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  STATE B: IDENTITY REVEALED — INTERACTIVE                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ░░░░░░░░░░░░░░░░░░░░ PREMIUM DARK BACKGROUND ░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│                                                                             │
│   🎊           ✨         🎉          ★           🎊                        │   Confetti
│                                                                             │
│                        ✦ DAY 15 ACHIEVED ✦                                  │
│                                                                             │
│                ╔═════════════════════════════════╗                          │
│                ║                                 ║                          │
│                ║      [ CLEAR FACE PHOTO ]       ║   ← Swipeable            │
│                ║                                 ║                          │
│                ║                                 ║                          │
│                ╚═════════════════════════════════╝                          │
│                            ◉ ○ ○ ○                                          │
│                                                                             │
│                         Priya Sharma                                        │
│                      ✓ Verified Identity                                    │
│                                                                             │
│            You did it. 15 days of showing up.                               │
│                Here's who they really are.                                  │
│                                                                             │
│       ┌─────────────────────────────────────────────────┐                   │
│       │              💬 Say Hello                       │                   │
│       └─────────────────────────────────────────────────┘                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Interactions Available:
├── Swipe photos horizontally
├── Tap "Say Hello" → Navigate to Chat Screen
├── Screen cannot be dismissed otherwise (must engage with chat)
```

---

## 9. Interaction Flow

### 9.1 Entry → Chat Navigation

```
TRIGGER:        Streak counter increments to Day 15
                │
                ▼
SYSTEM:         Display Identity Reveal screen for BOTH users simultaneously
                │
                ▼
UI:             Auto-play reveal animation (State A)
                │
                ▼
[2000ms]        Animation complete → State B (Interactive)
                │
                ▼
USER ACTION:    (Optional) Swipe through photos
                │
                ▼
USER ACTION:    Tap "Say Hello" button
                │
                ▼
HAPTIC:         .success
                │
                ▼
NAVIGATION:     Screen slides out → Chat Interface opens
                │
                ├── Pre-populate: Empty chat, ready for first message
                └── Show: Typing prompt "Say something..."
```

### 9.2 Cannot Dismiss Without Action

> [!IMPORTANT]
> This screen **cannot be closed** without tapping "Say Hello". There is no back button, no swipe-to-dismiss. The user must engage with the milestone.

**Rationale:** Day 15 is the culmination of the Unora journey. Allowing users to dismiss it without action would diminish the moment and potentially leave the chat unused.

---

## 10. Haptic Feedback

| Event | iOS | Android | Timing |
|-------|-----|---------|--------|
| Screen opens | `.light` | `TICK` (20ms) | 0ms |
| Photo reveal transition | `.medium` | `CLICK` (40ms) | ~800ms (midpoint) |
| Animation complete | `.success` | `CONFIRM` × 2 (100ms gap) | 1300ms |
| Swipe photo | `.light` | `TICK` (20ms) | On snap |
| Tap "Say Hello" | `.heavy` | `HEAVY_CLICK` (60ms) | On tap |

---

## 11. Copy Table

### 11.1 Screen Copy

| Element | Copy |
|---------|------|
| **Day Badge** | "✦ DAY 15 ACHIEVED ✦" |
| **Name** | "[Full Legal Name]" (dynamic) |
| **Verified Badge** | "✓ Verified Identity" |
| **Journey Text (Line 1)** | "You did it. 15 days of showing up." |
| **Journey Text (Line 2)** | "Here's who they really are." |
| **Primary CTA** | "💬 Say Hello" |

### 11.2 Alternative Copy Variants

| Variant | Copy |
|---------|------|
| **Personalized** | "You did it. 15 days of showing up. Meet [FirstName]." |
| **Minimal** | "15 days. This is them." |
| **Romantic (Partner Server)** | "After 15 days of patience, here's the face behind the heart." |

---

## 12. Accessibility

### 12.1 Screen Reader Announcements

| Moment | Announcement |
|--------|--------------|
| **Screen Opens** | "Day 15 achieved. Identity reveal. Please wait while the reveal animation plays." |
| **Animation Complete** | "Identity revealed. [Priya Sharma]. Verified identity. You did it. 15 days of showing up. Here's who they really are. Say Hello button available. Swipe left or right to view more photos." |
| **Photo Swipe** | "Photo [2] of [4]." |
| **Button Tap** | "Opening chat with [Priya]." |

### 12.2 Reduced Motion

When `prefers-reduced-motion` is enabled:
- No confetti animation
- No dissolve/flip animation — photo appears immediately
- No glow pulse animations
- Fade transitions reduced to 100ms
- Haptics still fire

### 12.3 Focus Management

- On animation complete, focus moves to **"Say Hello" button**
- Photo gallery is swipeable but not focusable (focus stays on button)

---

## 13. Technical Notes

### 13.1 Data Sources

| Data | Source | Notes |
|------|--------|-------|
| **Full Name** | User-provided (onboarding) | Name provided by user, not nickname |
| **Photos** | User's uploaded photos (onboarding) | Face photos only (flagged during upload) |
| **Connection ID** | Streak system | Used to create chat channel |

### 13.2 API Requirements

```typescript
interface IdentityRevealData {
  connectionId: string;
  partnerId: string;
  partnerName: string;           // Full legal name
  partnerPhotos: string[];       // Array of photo URLs
  chatChannelId: string;         // Pre-created, ready for messages
  revealTimestamp: string;       // ISO timestamp
}

interface IdentityRevealAction {
  type: 'open_chat';
  connectionId: string;
  chatChannelId: string;
}
```

### 13.3 Analytics Events

| Event | Trigger | Properties |
|-------|---------|------------|
| `identity_reveal_viewed` | Screen opens | `connection_id`, `user_tier`, `partner_tier` |
| `identity_reveal_animation_complete` | Animation finishes | `animation_duration_ms` |
| `identity_reveal_photo_swiped` | Photo changed | `photo_index`, `total_photos` |
| `identity_reveal_chat_opened` | "Say Hello" tapped | `time_to_tap_ms` |

---

## 14. Related Screens

| Screen | Relationship |
|--------|--------------|
| [Screen 12: Streak View](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_12_StreakView.md) | Pre-reveal — shows Day 15 milestone approaching |
| [Screen 17: Reveal Event](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_17_RevealEvent.md) | Earlier milestone reveals (Days 3-12) |
| [Screen 20: Chat Interface](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_20_ChatInterface.md) | Post-reveal — where conversation begins |

---

**Document maintained by:** Product Design Team  
**Last updated:** January 2026  
**Review cycle:** With each PRD/DSD update
