# Unora — UI Specification: Screen #16

## Payment / Recovery Window

**Version:** 1.0  
**Last Updated:** January 2026  
**Status:** Final  
**Author:** Product Design Team

---

## 1. Document Metadata

| Attribute | Value |
|-----------|-------|
| **Screen Name** | Payment / Recovery Window |
| **Screen ID** | `SCREEN_16_RECOVERY_WINDOW` |
| **User Flow Reference** | [Section 2.C.4 — Recovery Logic (Day N+1)](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) |
| **PRD Reference** | [Section 13.4 — Streak Flow (Day N+1)](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md), [Section 16.2 — Tier Structure](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md), [Section 17 — Credits & Payment Protection](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) |
| **DSD Reference** | [Section 2.1 — Color Palette](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md), [Section 8.3 — Streak State Variants](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) |
| **Visibility** | Breaker Only (User who missed Day N) |
| **Trigger** | Day N+1 begins after a single user missed check-in on Day N |

---

## 2. Screen Placement & Purpose

### 2.1 Flow Context

```
Day N: User B misses check-in → Streak moves to AT RISK
                │
                ▼
         End of Day N: Miss is recorded
                │
                ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│                       DAY N+1: PAYMENT WINDOW OPENS                         │
│                                                                             │
│   ┌─────────────────────────┐     ┌───────────────────────────────────────┐ │
│   │  USER A: MAINTAINER     │     │  USER B: BREAKER                      │ │
│   │  (Checked in on Day N)  │     │  (Missed on Day N)                    │ │
│   ├─────────────────────────┤     ├───────────────────────────────────────┤ │
│   │                         │     │                                       │ │
│   │  Sees: Waiting screen   │     │  ★ SEES: RECOVERY WINDOW (Screen 16) │ │
│   │  "Waiting for partner's │     │                                       │ │
│   │   decision"             │     │  Options:                             │ │
│   │                         │     │  ├── Recover (Pay or Use Allowance)   │ │
│   │  Can still send nudges  │     │  └── Let it Reset                     │ │
│   │                         │     │                                       │ │
│   │  ⚠️ NEVER PAYS          │     │  ⚠️ NO CHECK-IN OPTION                │ │
│   │                         │     │                                       │ │
│   └─────────────────────────┘     └───────────────────────────────────────┘ │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 2.2 Purpose

This screen serves as the **single recovery interface** for users who missed a check-in. It provides:

1. **Clear accountability** — The breaker understands they missed.
2. **Recovery path** — An option to continue the streak.
3. **Graceful exit** — An option to let the streak reset without penalty language.
4. **Payment protection assurance** — Microcopy confirming credit conversion if the connection ends.

### 2.3 Critical Logic Constraints

> [!CAUTION]
> The following rules are **non-negotiable** and define the screen's behavior.

| Constraint | Implementation |
|------------|----------------|
| **Breaker Only** | This screen is NEVER shown to the Maintaining User |
| **No Check-In Option** | There is NO way to check in on Day N+1 — the day is lost |
| **Payment = Recovery** | Recovery is achieved by payment or using a free allowance, not by action |
| **Maintaining User Never Pays** | The user who showed up is never prompted for payment |
| **Credit Protection Required** | Microcopy MUST include the 24-hour credit conversion assurance |

---

## 3. Visual Theme & Styling

### 3.1 Theme: Payment / Terracotta

This screen uses the **Payment / Terracotta** theme, NOT the standard server color. This is intentional — it visually distinguishes the payment state from normal streak interactions.

```css
/* Payment Theme Tokens */
--status-payment: #D4714A;        /* Terracotta — Primary accent */
--status-payment-bg: #F9EAE3;     /* Light terracotta surface */

/* Dark Mode Equivalents */
--dm-status-payment: #E07D5A;     /* Softer terracotta */
--dm-status-payment-bg: #2A1F1A;  /* Dark warm surface */
```

### 3.2 Tone

The visual and copy tone is **supportive, not punitive**:

| ❌ Avoid | ✅ Use Instead |
|---------|---------------|
| "You broke the streak" | "Your streak needs recovery" |
| "Pay your penalty" | "Continue your streak" |
| "You failed" | "You missed yesterday's check-in" |
| "Punishment" | "Recovery" |
| Error-red backgrounds | Warm terracotta surfaces |

---

## 4. Layout Specification

### 4.1 Screen Structure (Modal Overlay)

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   [ Dimmed Background: Streak Screen Visible at 60% opacity ]               │
│                                                                             │
│       ┌─────────────────────────────────────────────────────────────────┐   │
│       │                                                                 │   │
│       │   ┌─────────────────────────────────────────────────────────┐   │   │
│       │   │                                                         │   │   │
│       │   │                   💫                                    │   │   │  Icon: 48px
│       │   │                                                         │   │   │
│       │   │            Your streak needs recovery                   │   │   │  H3: Outfit 18px/600
│       │   │                                                         │   │   │
│       │   │   You missed yesterday's check-in.                      │   │   │  Body: Inter 14px/400
│       │   │   Your partner showed up and would love                 │   │   │
│       │   │   to keep going with you.                               │   │   │
│       │   │                                                         │   │   │
│       │   └─────────────────────────────────────────────────────────┘   │   │
│       │                                                                 │   │
│       │   ─────────────────────────────────────────────────────────     │   │  Divider
│       │                                                                 │   │
│       │   ┌─────────────────────────────────────────────────────────┐   │   │
│       │   │                                                         │   │   │
│       │   │              RECOVERY CTA (Dynamic)                     │   │   │  Primary Button
│       │   │                                                         │   │   │
│       │   └─────────────────────────────────────────────────────────┘   │   │
│       │                                                                 │   │
│       │   If this connection ends within 24h,                           │   │  Microcopy: Inter 12px/400
│       │   this amount converts to credits.                              │   │  Color: --unora-ink-tertiary
│       │                                                                 │   │
│       │                                                                 │   │
│       │                 Let it reset instead →                          │   │  Tertiary Link
│       │                                                                 │   │
│       └─────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 4.2 Layout Tokens

| Property | Value |
|----------|-------|
| **Overlay Type** | Modal (centered) |
| **Overlay Background** | `rgba(26, 26, 26, 0.6)` — Dimmed backdrop |
| **Modal Background** | `var(--status-payment-bg)` / `#F9EAE3` |
| **Modal Border** | `1px solid var(--status-payment)` / `#D4714A` |
| **Modal Radius** | `20px` / `var(--radius-xl)` |
| **Modal Shadow** | `0 8px 32px rgba(0, 0, 0, 0.15)` |
| **Modal Width** | `min(340px, 90vw)` |
| **Modal Padding** | `24px` / `var(--space-6)` |
| **Z-Index** | `1000` (above all content) |

### 4.3 Content Spacing

```
Modal Internal Layout:
├── Icon:              48px, centered
├── Gap:               16px
├── Headline (H3):     Outfit 18px/600, centered
├── Gap:               12px
├── Body Text:         Inter 14px/400, centered, max-width 280px
├── Gap:               24px
├── Divider:           1px solid var(--border-subtle), full width
├── Gap:               24px
├── Primary CTA:       Full width, 52px height
├── Gap:               12px
├── Credit Assurance:  Inter 12px/400, centered, tertiary color
├── Gap:               20px
└── Tertiary Link:     Inter 14px/500, centered
```

---

## 5. Component Specifications

### 5.1 Recovery Header

**Purpose:** Establishes context without blame language.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│                            💫                                               │
│                                                                             │
│                 Your streak needs recovery                                  │
│                                                                             │
│      You missed yesterday's check-in. Your partner showed up               │
│      and would love to keep going with you.                                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Icon: Custom "spark" or "restart" icon (not warning icon)
├── Size: 48px
├── Color: var(--status-payment)
└── Alternative: Phosphor "ArrowCounterClockwise" or custom

Headline:
├── Font: Outfit (var(--font-display))
├── Size: 18px (var(--h3-size))
├── Weight: 600
├── Color: var(--unora-ink-primary)
├── Alignment: Center

Body:
├── Font: Inter (var(--font-body))
├── Size: 14px (var(--body-size))
├── Weight: 400
├── Line Height: 1.5
├── Color: var(--unora-ink-secondary)
├── Alignment: Center
├── Max Width: 280px
```

### 5.2 Primary Action: Recovery CTA (Dynamic)

**This button changes based on tier and remaining allowance.**

#### State A: Payable (Standard ₹99)

Applies to: **Free Tier (always)**, Plus/Pro after allowance exhausted.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                                                                     │   │
│   │                    Recover Streak — ₹99                             │   │
│   │                                                                     │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Button Specs:
├── Background: var(--status-payment) / #D4714A
├── Text: #FFFFFF
├── Font: Inter 16px / 600
├── Height: 52px
├── Radius: 12px (var(--radius-md))
├── Width: 100%
├── Shadow: 0 2px 8px rgba(0, 0, 0, 0.08)

States:
├── Default: Full opacity
├── Pressed: Scale 0.98, opacity 0.9
├── Loading: Spinner replaces icon, text fades to 60%
```

#### State B: Free Recovery Available (Plus/Pro Allowance)

Applies to: **Plus Tier (1st recovery)**, **Pro Tier (1st or 2nd recovery)**.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   ┌─────────────────────────────────────────────────────────────────────┐   │
│   │                                                                     │   │
│   │               ✨  Use Free Recovery                                 │   │
│   │                                                                     │   │
│   └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│   Your Plus plan includes 1 free recovery per connection.                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Button Specs:
├── Background: var(--status-payment) / #D4714A (same as paid)
├── Text: #FFFFFF
├── Font: Inter 16px / 600
├── Height: 52px
├── Radius: 12px
├── Width: 100%
├── Icon: Sparkles (✨) 20px, left of text

Sub-label (below button):
├── Font: Inter 12px / 400
├── Color: var(--unora-ink-tertiary)
├── Alignment: Center
├── Content varies by tier:
│   ├── Plus: "Your Plus plan includes 1 free recovery per connection."
│   └── Pro: "Your Pro plan includes 2 free recoveries per connection."
```

#### Dynamic CTA Logic Table

| User Tier | Recoveries Used This Connection | CTA Text | Shows Credit Microcopy |
|-----------|--------------------------------|----------|------------------------|
| **Free** | N/A (no allowance) | "Recover Streak — ₹99" | ✓ Yes |
| **Plus** | 0 of 1 used | "✨ Use Free Recovery" | ✗ No |
| **Plus** | 1 of 1 used | "Recover Streak — ₹99" | ✓ Yes |
| **Pro** | 0 of 2 used | "✨ Use Free Recovery" | ✗ No |
| **Pro** | 1 of 2 used | "✨ Use Free Recovery" | ✗ No |
| **Pro** | 2 of 2 used | "Recover Streak — ₹99" | ✓ Yes |

### 5.3 Credit Assurance Microcopy

**Required whenever payment is involved. Hidden when using free allowance.**

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│   If this connection ends within 24h,                                       │
│   this amount converts to credits.                                          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Specs:
├── Font: Inter 12px / 400 (var(--caption-size))
├── Color: var(--unora-ink-tertiary) / #7A7A7A
├── Alignment: Center
├── Line Height: 1.4
├── Max Width: 280px
├── Margin Top: 12px (below CTA)
```

> [!IMPORTANT]
> This microcopy is **mandatory** for payment transactions. It implements the Credit Conversion Rule from PRD Section 17.

### 5.4 Tertiary Action: Let It Reset

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│                      Let it reset instead →                                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Specs:
├── Type: Tertiary link (text-only)
├── Font: Inter 14px / 500
├── Color: var(--status-payment) / #D4714A
├── Alignment: Center
├── Touch Target: 44px height minimum
├── No border, no background

States:
├── Default: Full opacity
├── Hover: Underline
├── Pressed: Opacity 0.8

Action: Dismisses modal → Streak resets to 0 → Connection remains active
```

---

## 6. State Definitions

### 6.1 State A: Payable (Standard ₹99)

**Trigger:** Free tier user OR Plus/Pro user with exhausted allowance.

```
┌─────────────────────────────────────────────────────────────────┐
│  STATE A: PAYABLE                                               │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│                        💫                                       │
│                                                                 │
│         Your streak needs recovery                              │
│                                                                 │
│   You missed yesterday's check-in. Your partner                 │
│   showed up and would love to keep going with you.              │
│                                                                 │
│   ─────────────────────────────────────────────────             │
│                                                                 │
│   ┌─────────────────────────────────────────────────────────┐   │
│   │            Recover Streak — ₹99                         │   │
│   └─────────────────────────────────────────────────────────┘   │
│                                                                 │
│   If this connection ends within 24h,                           │
│   this amount converts to credits.                              │
│                                                                 │
│                                                                 │
│               Let it reset instead →                            │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

Components visible:
├── Recovery Header (Icon + Headline + Body)
├── Divider
├── Primary CTA: "Recover Streak — ₹99"
├── Credit Assurance Microcopy ✓
└── Tertiary Link: "Let it reset instead"
```

### 6.2 State B: Free Allowance Available

**Trigger:** Plus user (1st recovery) OR Pro user (1st/2nd recovery).

```
┌─────────────────────────────────────────────────────────────────┐
│  STATE B: FREE ALLOWANCE                                        │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│                        💫                                       │
│                                                                 │
│         Your streak needs recovery                              │
│                                                                 │
│   You missed yesterday's check-in. Your partner                 │
│   showed up and would love to keep going with you.              │
│                                                                 │
│   ─────────────────────────────────────────────────             │
│                                                                 │
│   ┌─────────────────────────────────────────────────────────┐   │
│   │           ✨  Use Free Recovery                         │   │
│   └─────────────────────────────────────────────────────────┘   │
│                                                                 │
│   Your Plus plan includes 1 free recovery                       │
│   per connection.                                               │
│                                                                 │
│                                                                 │
│               Let it reset instead →                            │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

Components visible:
├── Recovery Header (Icon + Headline + Body)
├── Divider
├── Primary CTA: "✨ Use Free Recovery"
├── Allowance Info (tier-specific sub-label)
├── ⚠️ NO Credit Assurance (no payment = no credit protection needed)
└── Tertiary Link: "Let it reset instead"

Sub-label variants:
├── Plus (1 allowance): "Your Plus plan includes 1 free recovery per connection."
├── Pro (2 allowances): "Your Pro plan includes 2 free recoveries per connection."
└── Pro (1 remaining): "You have 1 free recovery remaining for this connection."
```

### 6.3 State C: Processing

**Trigger:** User taps "Recover" → Payment/action in progress.

```
┌─────────────────────────────────────────────────────────────────┐
│  STATE C: PROCESSING                                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│                        💫                                       │
│                                                                 │
│         Your streak needs recovery                              │
│                                                                 │
│   You missed yesterday's check-in. Your partner                 │
│   showed up and would love to keep going with you.              │
│                                                                 │
│   ─────────────────────────────────────────────────             │
│                                                                 │
│   ┌─────────────────────────────────────────────────────────┐   │
│   │               ◌  Processing...                          │   │
│   └─────────────────────────────────────────────────────────┘   │
│                                                                 │
│   If this connection ends within 24h,                           │
│   this amount converts to credits.                              │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

CTA Button Changes:
├── Text: "Processing..." (replaces original text)
├── Icon: Circular spinner, 20px, white, left of text
├── State: Disabled (non-interactive)
├── Opacity: 0.7

Additional:
├── Tertiary link is HIDDEN during processing
├── Modal cannot be dismissed
├── User cannot tap outside to close
```

### 6.4 State D: Recovery Success (Transition)

**Trigger:** Payment/allowance consumed successfully.

```
┌─────────────────────────────────────────────────────────────────┐
│  STATE D: SUCCESS (Brief transition state)                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│                                                                 │
│                        ✓                                        │
│                                                                 │
│               Streak recovered!                                 │
│                                                                 │
│       Your streak continues from Day 7.                         │
│       Keep the momentum going.                                   │
│                                                                 │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

Specs:
├── Icon: Checkmark, 48px, var(--status-active) / #4A9B8C
├── Headline: "Streak recovered!" — H3, center
├── Body: "Your streak continues from Day X." — Body, center
├── Duration: 2 seconds

Transition:
├── Auto-dismiss after 2 seconds
├── Navigate back to Streak Screen (Active state)
├── Modal fades out (200ms)
```

---

## 7. Interaction Model

### 7.1 Modal Behavior

| Interaction | Behavior |
|-------------|----------|
| **Tap Recover CTA** | Initiates payment/allowance consumption → State C |
| **Tap "Let it reset"** | Dismisses modal → Streak resets to 0 → Connection stays active |
| **Tap outside modal** | Modal remains (cannot dismiss by tapping outside) |
| **Swipe down** | Modal remains (no swipe-to-dismiss) |
| **Back button (Android)** | Opens confirmation: "Are you sure? Your streak will reset." |

> [!WARNING]
> The modal is **persistent**. Users must make an explicit choice. This prevents accidental resets.

### 7.2 Haptic Feedback

| Action | Haptic | iOS | Android |
|--------|--------|-----|---------|
| Modal appears | Light | `.light` | `TICK` (20ms) |
| Tap Recover | Medium | `.medium` | `CLICK` (40ms) |
| Recovery Success | Heavy (celebration) | `.heavy` × 2 | `HEAVY_CLICK` × 2 |
| Tap "Let it reset" | Rigid (caution) | `.rigid` | `REJECT` |

### 7.3 Animation Sequences

**Modal Open:**
```
[0ms]    Screen triggered
[0ms]    Overlay fades in (0 → 60% opacity)
[0ms]    Modal scales from 0.9 → 1.0, opacity 0 → 1
[0ms]    Haptic: Light tap
[250ms]  Modal fully visible
```

**Modal Close (Success):**
```
[0ms]    Success state displayed
[2000ms] Auto-dismiss triggered
[2000ms] Modal scales 1.0 → 0.95, opacity 1 → 0
[2000ms] Overlay fades out
[2200ms] User back on Streak Screen (Active state)
```

**Modal Close (Reset):**
```
[0ms]    User taps "Let it reset"
[0ms]    Haptic: Rigid
[0ms]    Modal scales 1.0 → 0.95, opacity 1 → 0
[0ms]    Overlay fades out
[200ms]  User back on Streak Screen (Reset state, Day 0)
```

---

## 8. Copy Table

### 8.1 Static Copy

| Element | Copy | Notes |
|---------|------|-------|
| **Headline** | "Your streak needs recovery" | Supportive, not accusatory |
| **Body** | "You missed yesterday's check-in. Your partner showed up and would love to keep going with you." | Emphasizes partner's commitment |
| **CTA (Paid)** | "Recover Streak — ₹99" | Price is always ₹99 |
| **CTA (Free)** | "✨ Use Free Recovery" | Sparkle emoji adds positivity |
| **Credit Microcopy** | "If this connection ends within 24h, this amount converts to credits." | Required for payment |
| **Tertiary Action** | "Let it reset instead →" | No blame, clear alternative |
| **Success Headline** | "Streak recovered!" | Celebratory |
| **Success Body** | "Your streak continues from Day {X}. Keep the momentum going." | Dynamic day count |

### 8.2 Dynamic Copy (Tier-Specific)

| Tier | Allowance State | Sub-label Copy |
|------|-----------------|----------------|
| Free | N/A | (None — no sub-label shown) |
| Plus | 0/1 used | "Your Plus plan includes 1 free recovery per connection." |
| Plus | 1/1 used | (None — shows credit microcopy instead) |
| Pro | 0/2 used | "Your Pro plan includes 2 free recoveries per connection." |
| Pro | 1/2 used | "You have 1 free recovery remaining for this connection." |
| Pro | 2/2 used | (None — shows credit microcopy instead) |

### 8.3 Error Copy

| Scenario | Copy |
|----------|------|
| **Payment Failed** | "Payment couldn't be processed. Please try again or check your payment method." |
| **Network Error** | "Something went wrong. Check your connection and try again." |
| **Timeout** | "This is taking longer than expected. Please wait or try again." |

---

## 9. Accessibility

### 9.1 Screen Reader Announcements

| State | Announcement |
|-------|--------------|
| **Modal Opens** | "Streak recovery needed. You missed yesterday's check-in. Options: Recover streak for 99 rupees, or let it reset." |
| **Free Allowance** | "Streak recovery needed. You have a free recovery available. Options: Use free recovery, or let it reset." |
| **Processing** | "Processing your recovery request." |
| **Success** | "Streak recovered! Your streak continues from Day 7." |

### 9.2 Focus Management

1. When modal opens, focus moves to **Headline**
2. Tab order: Headline → Primary CTA → Tertiary Link
3. Focus trap is enabled — cannot tab outside modal
4. On close, focus returns to previous location

### 9.3 Touch Targets

| Element | Size | Meets Minimum |
|---------|------|---------------|
| Primary CTA | 100% width × 52px | ✓ (52px > 44px) |
| Tertiary Link | 100% width × 44px | ✓ |

---

## 10. Technical Notes

### 10.1 API Requirements

```typescript
interface RecoveryWindowData {
  userId: string;
  connectionId: string;
  streakDay: number;                // Current day before recovery
  userTier: 'free' | 'plus' | 'pro';
  freeRecoveriesUsed: number;       // 0, 1, or 2
  freeRecoveriesTotal: number;      // 0 (free), 1 (plus), 2 (pro)
  paymentAmount: number;            // Always 99 (₹)
  partnerId: string;
  partnerCheckedInAt: string;       // ISO timestamp
}

interface RecoveryAction {
  type: 'pay' | 'use_allowance' | 'reset';
  connectionId: string;
  paymentMethodId?: string;         // Required if type === 'pay'
}

interface RecoveryResult {
  success: boolean;
  newStreakDay: number;             // Same as before recovery
  creditsGenerated?: number;        // Only if connection ends within 24h
  error?: string;
}
```

### 10.2 Analytics Events

| Event | Trigger | Properties |
|-------|---------|------------|
| `recovery_window_viewed` | Modal opens | `tier`, `free_allowance_available`, `streak_day` |
| `recovery_initiated` | User taps Recover | `type: 'paid' | 'free'`, `amount`, `streak_day` |
| `recovery_completed` | Recovery succeeds | `type`, `streak_day`, `duration_ms` |
| `recovery_failed` | Recovery fails | `error`, `type` |
| `recovery_declined` | User taps "Let it reset" | `streak_day`, `tier` |

### 10.3 Error Handling

| Error Case | UI Response |
|------------|-------------|
| Payment gateway timeout | Show inline error, keep modal open, retry button |
| Insufficient funds | Show error toast, suggest alternate payment |
| Network failure | Show error state with retry option |
| Concurrent modification | Refresh state, modal may close if connection already ended |

---

## 11. Dark Mode Adaptation

```css
/* Dark Mode Token Mapping */
--status-payment: #E07D5A;          /* Lighter terracotta for dark mode */
--status-payment-bg: #2A1F1A;       /* Dark warm surface */
--unora-ink-primary: #F5F5F3;       /* Light text on dark */
--unora-ink-secondary: #C4C4C0;
--unora-ink-tertiary: #8A8A86;
--surface-overlay: rgba(0, 0, 0, 0.7);
```

### Dark Mode Visual

```
┌─────────────────────────────────────────────────────────────────┐
│  DARK MODE: Recovery Window                                     │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   Background: #121212 (app background, dimmed)                  │
│                                                                 │
│   ┌─────────────────────────────────────────────────────────┐   │
│   │                                                         │   │ Modal: #2A1F1A
│   │                      💫                                 │   │ (warm dark surface)
│   │                                                         │   │
│   │        Your streak needs recovery                       │   │ Text: #F5F5F3
│   │                                                         │   │
│   │   You missed yesterday's check-in...                    │   │ Body: #C4C4C0
│   │                                                         │   │
│   │   [    Recover Streak — ₹99    ]                        │   │ CTA: #E07D5A bg
│   │                                                         │   │
│   │   Credit microcopy...                                   │   │ Caption: #8A8A86
│   │                                                         │   │
│   │          Let it reset instead →                         │   │ Link: #E07D5A
│   │                                                         │   │
│   └─────────────────────────────────────────────────────────┘   │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

Modal border: 1px solid var(--dm-status-payment) at 50% opacity
Modal shadow: 0 8px 32px rgba(0, 0, 0, 0.4)
```

---

## 12. Edge Cases

| Scenario | Handling |
|----------|----------|
| **Both users miss on Day N** | This screen does NOT appear. Streak resets immediately with no payment window. |
| **User upgrades tier during payment window** | Refresh allowance data. If now eligible for free recovery, update CTA dynamically. |
| **Connection ends while modal is open** | Close modal, show toast: "This connection has ended." Navigate to Connections list. |
| **User already recovered today** | Should not be possible — modal only appears once per payment window. |
| **Payment window expires** | Server-side: auto-reset streak. Client: modal should not be shown if window has passed. |

---

## 13. Related Screens

| Screen | Relationship |
|--------|--------------|
| [Screen 12: Streak View](file:///c:/Unora/Founder_Product_docs/Unora_Spec_12_StreakView.md) | Parent screen — Recovery Window overlays this |
| [Screen 14: Payment Flow](file:///c:/Unora/Founder_Product_docs/Unora_Spec_14_PaymentFlow.md) | Invoked if user taps "Recover Streak — ₹99" |
| [Screen 15: Maintainer Waiting View](file:///c:/Unora/Founder_Product_docs/Unora_Spec_15_MaintainerWaiting.md) | What the other user sees during recovery window |

---

**Document maintained by:** Product Design Team  
**Last updated:** January 2026  
**Review cycle:** With each PRD/DSD update
