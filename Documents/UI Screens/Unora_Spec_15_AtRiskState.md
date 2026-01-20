# Unora UI Specification — At-Risk State (Partner Missed)

**Document ID:** Spec-15  
**Screen Name:** At-Risk State (Partner Missed)  
**Version:** 1.0  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**At-Risk State (Partner Missed)** — Warning state when partner hasn't checked in

### 1.2 User Flow Reference
**Phase 4 (Streak Loop) → Warning State** — Displayed when User A checks in but User B has not by end of day.

**Sequence:**
```
User Check-In → Partner Fails to Check-In → [At-Risk State] → Nudge → Wait
                                                    ↓
                                           (Next day: Payment State)
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.D.3](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Alert the user that their streak is in danger and offer a supportive tool (Nudge) to remind their partner. The tone balances **urgency with support** — not punitive or aggressive.

### 1.4 Key Constraints

| Rule | Details |
|------|---------|
| **No Payment** | Payment is only requested on Day N+1 if partner doesn't recover |
| **Nudge Limits** | Limited nudges per day (2 for Free, 5 for Plus, Unlimited for Pro) |
| **Visual Theme** | Amber/Warning — NOT Error/Red |

---

## 2. Low-Fidelity Wireframes (ASCII)

### 2.1 At-Risk Full View

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   ←  A****  🔥                                    ⚙️  ⋮    │  ← Header
│                                                             │
│                                                             │
│                        Day 7                                │  ← Day counter
│                      ───────                                │     (Amber tint)
│                    ⚠️ At Risk                               │  ← Status
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │   ⚠️  Streak at risk                                │   │  ← Alert Banner
│   │       Your partner hasn't checked in yet.           │   │     (Amber)
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                                                     │   │
│   │   ┌────────────────┐    ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐   │   │
│   │   │                │    ╎                       ╎   │   │
│   │   │   ✓ You        │    ╎    ⏳ Partner         ╎   │   │  ← Status Split
│   │   │   Checked In   │    ╎    Waiting...         ╎   │   │
│   │   │                │    ╎                       ╎   │   │
│   │   │   (Green)      │    ╎    (Amber/Dashed)     ╎   │   │
│   │   │                │    ╎                       ╎   │   │
│   │   └────────────────┘    └ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┘   │   │
│   │                                                     │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│       They have until midnight to check in.                 │  ← Helper text
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │          🔔  Nudge Partner                          │   │  ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │     (Amber)
│                                                             │
│       Nudges remaining: 2 today                             │  ← Limit indicator
│                                                             │
│   ─────────────────────────────────────────────────────     │
│                                                             │
│   CONNECTION CORE                                           │
│   ✦ Trust at risk... (dimmed pulse)                          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Status Split Detail

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌────────────────────┐    ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐    │
│   │                    │    ╎                         ╎    │
│   │   ✓                │    ╎   ⏳                    ╎    │
│   │                    │    ╎                         ╎    │
│   │   You              │    ╎   Partner               ╎    │
│   │   Checked In       │    ╎   Waiting...            ╎    │
│   │                    │    ╎                         ╎    │
│   │   Upper Body       │    ╎   ○ ○ ○ (pulsing)       ╎    │
│   │                    │    ╎                         ╎    │
│   └────────────────────┘    └ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┘    │
│                                                             │
│        ✓ Solid Green            ⏳ Dashed Amber             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.3 Nudge Sent State

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │      ✓ Nudge Sent                                   │   │  ← Disabled
│   └─────────────────────────────────────────────────────┘   │     (Muted Amber)
│                                                             │
│       You can nudge again in 2 hours                        │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
AT-RISK VIEW CONTAINER
├── Same as Streak Detail (Spec-12)
├── Theme Override: Amber replaces Server color for status
│
├── [HEADER] — Standard
│
├── [HERO] — 120px
│   ├── Day counter: Amber tint
│   └── Status: "⚠️ At Risk"
│
├── [ALERT BANNER] — 60px
│   ├── Background: var(--feedback-warning) @ 10%
│   ├── Border: 1px solid var(--feedback-warning)
│   └── Margin: 0 20px
│
├── [STATUS SPLIT CARD] — 140px
│   ├── Two columns: You / Partner
│   └── Margin: 20px
│
├── [ACTION AREA] — 100px
│   ├── Nudge button
│   └── Limit indicator
│
└── [CONNECTION CORE] — Dimmed Trust Orb (at-risk visual state)

Premium Dark Mode (Default):
├── Background: var(--pdm-background) → #0D0D0F
├── Alert banner: Amber with subtle glow border
├── User column: Green success with gentle glow
└── Partner column: Amber dashed with pulsing glow
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Day counter** | Amber-tinted with subtle glow |
| **Alert banner** | Amber border with glow: `0 0 8px rgba(230,164,58,0.2)` |
| **User column (✓)** | Green @ 10% bg, solid green border with subtle glow |
| **Partner column (⏳)** | Amber @ 8% bg, dashed amber border with pulse |
| **Nudge button** | Amber bg with outer glow: `0 0 12px rgba(230,164,58,0.3)` |
| **Pulsing animation** | Amber dots with glow intensity variation |

**At-Risk Alert Glow:**
```css
/* Amber alert banner with soft glow */
.alert-banner.warning {
  border: 1px solid var(--feedback-warning);
  box-shadow: inset 0 0 12px rgba(230, 164, 58, 0.08),
              0 0 8px rgba(230, 164, 58, 0.15);
}

/* Partner waiting dots pulse with glow */
@keyframes waiting-pulse {
  0%, 100% { opacity: 0.4; filter: drop-shadow(0 0 2px rgba(230, 164, 58, 0.3)); }
  50%      { opacity: 1.0; filter: drop-shadow(0 0 6px rgba(230, 164, 58, 0.5)); }
}
```



### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Alert banner padding | `var(--space-4)` | 16px |
| Alert margin-bottom | `var(--space-4)` | 16px |
| Status card padding | `var(--space-5)` | 20px |
| Column gap | `var(--space-4)` | 16px |
| Button margin-top | `var(--space-5)` | 20px |

---

## 4. Component Inventory

### 4.1 Amber/Warning Theme

**Reference:** [Unora_DesignSystem.md — Section 2.1, 8.3](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)

| Token | Value | Usage |
|-------|-------|-------|
| `var(--feedback-warning)` | #E6A43A (Amber) | Borders, icons, buttons |
| `var(--feedback-warning) @ 10%` | rgba(230,164,58,0.1) | Backgrounds |
| `var(--feedback-success)` | #4A9B8C (Teal) | User's "Checked In" state |

> **Important:** Amber signals urgency but NOT error. It's supportive, not punitive.

### 4.2 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Day counter | Outfit | 700 | 64px | `--feedback-warning` @ 80% |
| "At Risk" status | Inter | 600 | 14px | `--feedback-warning` |
| Alert headline | Inter | 600 | 16px | `--feedback-warning` |
| Alert body | Inter | 400 | 14px | `--unora-ink-secondary` |
| Status label | Inter | 600 | 14px | (Green or Amber) |
| Status sublabel | Inter | 400 | 12px | `--unora-ink-tertiary` |
| Helper text | Inter | 400 | 12px | `--unora-ink-tertiary` |
| Nudge limit | Inter | 500 | 12px | `--unora-ink-tertiary` |

### 4.3 Alert Banner

```
ALERT BANNER
├── Width: 100% - padding
├── Background: var(--feedback-warning) @ 10%
├── Border: 1px solid var(--feedback-warning)
├── Border radius: var(--radius-md) → 12px
├── Padding: 16px
│
├── [ICON]
│   ├── ⚠️ Warning triangle
│   ├── Size: 20px
│   └── Color: var(--feedback-warning)
│
└── [TEXT]
    ├── Headline: "Streak at risk"
    └── Body: "Your partner hasn't checked in yet."
```

### 4.4 Status Split Card

```
STATUS SPLIT CARD
├── Width: 100% - padding
├── Background: var(--surface-card)
├── Border: 1px solid var(--border-subtle)
├── Border radius: var(--radius-lg) → 16px
├── Padding: 20px
├── Display: flex, 2 columns, gap 16px
│
├── [USER COLUMN] — "You"
│   ├── Background: var(--feedback-success) @ 10%
│   ├── Border: 2px solid var(--feedback-success)
│   ├── Border radius: var(--radius-md)
│   ├── Padding: 16px
│   ├── Icon: ✓ checkmark (24px, green)
│   ├── Label: "You"
│   ├── Sublabel: "Checked In"
│   └── Activity: "Upper Body"
│
└── [PARTNER COLUMN] — "Partner"
    ├── Background: var(--feedback-warning) @ 5%
    ├── Border: 2px dashed var(--feedback-warning)  ← DASHED
    ├── Border radius: var(--radius-md)
    ├── Padding: 16px
    ├── Icon: ⏳ hourglass (24px, amber)
    ├── Label: "Partner"
    ├── Sublabel: "Waiting..."
    └── Animation: Pulsing dots ○ ○ ○
```

### 4.5 Nudge Button

#### State A: Available
| Property | Value |
|----------|-------|
| Height | 52px |
| Width | 100% |
| Background | `var(--feedback-warning)` |
| Text | White, Inter 16px/600 |
| Icon | 🔔 Bell (16px, left of text) |
| Border radius | `var(--radius-md)` → 12px |

#### State B: Sent (Disabled)
| Property | Value |
|----------|-------|
| Background | `var(--feedback-warning)` @ 30% |
| Text | `var(--feedback-warning)`, Inter 16px/600 |
| Icon | ✓ checkmark |
| Label | "Nudge Sent" |
| Interaction | Disabled |

#### State C: Limit Reached (Upsell)
| Property | Value |
|----------|-------|
| Background | `var(--surface-disabled)` |
| Text | `var(--unora-ink-tertiary)` |
| Label | "No nudges left" |
| Action | Opens upgrade modal |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Nudge button (available) | Send nudge notification |
| Tap | Nudge button (limit) | Open upgrade modal |
| Timer | Midnight | Transition to Payment State |

### 5.2 Nudge Flow

```
USER taps "Nudge Partner"
    │
    ▼
┌─────────────────────────────────────────┐
│ [CHECK] Nudges remaining today?         │
└─────────────────────────────────────────┘
    │                    │
    ▼                    ▼
  YES                   NO
    │                    │
    ▼                    ▼
┌─────────────────┐  ┌─────────────────────────────────────┐
│ SEND NUDGE      │  │ **LIMIT REACHED**                   │
│                 │  │                                     │
│ Haptic: Success │  │ Show Modal:                         │
│                 │  │ "No nudges left today"              │
│ Animation:      │  │                                     │
│ Bell wobble     │  │ Options:                            │
│                 │  │ ├── "Upgrade to Plus" (more nudges) │
│ Toast:          │  │ └── "Wait until tomorrow"           │
│ "Nudge sent!"   │  │                                     │
│                 │  │                                     │
│ Button →        │  │                                     │
│ "Nudge Sent"    │  │                                     │
│ (disabled)      │  │                                     │
└─────────────────┘  └─────────────────────────────────────┘
```

### 5.3 Nudge Limits by Tier

| Tier | Cooldown Period | Max Per At-Risk Period |
|------|-----------------|------------------------|
| Free | 1 per 24 hours | Max 1 per at-risk period |
| Plus | 1 per ~7 hours | Max 3 per at-risk period |
| Pro | 1 per ~5 hours | Max 4 per at-risk period |

> **Note:** Even Pro users are bounded to prevent harassment/spam. Limits are enforced per at-risk period (from partner miss until resolution).

### 5.4 Bell Animation

```
BELL WOBBLE ANIMATION
├── Trigger: On nudge send
├── Animation: Rotate -15° → 15° → -10° → 10° → 0°
├── Duration: 400ms
├── Easing: ease-out
└── Accompanies: Haptic feedback
```

### 5.5 Transitions

| Event | Transition |
|-------|------------|
| Button press | Scale 0.95, 100ms |
| Nudge sent | Bell wobble + button state change |
| Toast | Slide up from bottom, 3s display |
| Partner checks in | Transition to Success (State C) |
| Midnight without check-in | Transition to Payment State |

---

## 6. State Definitions

### 6.1 State Matrix

| State | Button | Indicator | Condition |
|-------|--------|-----------|-----------|
| Default At-Risk | "Nudge Partner" | "Nudges: 2" | Nudges available |
| Nudge Sent | "Nudge Sent" (disabled) | "Nudge again in 2h" | Cooldown active |
| Limit Reached | "No nudges left" | "Upgrade for more" | Daily limit hit |

### 6.2 Default At-Risk State

```
Alert: Visible (amber)
Status split: You ✓ / Partner ⏳
Button: "🔔 Nudge Partner" (amber, enabled)
Indicator: "Nudges remaining: X today"
```

### 6.3 Nudge Sent State

```
Alert: Visible
Button: "✓ Nudge Sent" (muted, disabled)
Indicator: "You can nudge again in X hours"
Cooldown: Active timer
```

### 6.4 Limit Reached State

```
Alert: Visible
Button: "No nudges left" (disabled)
Indicator: "Upgrade for unlimited nudges"
Action: Tap opens upgrade modal
```

---

## 7. Content & Copy Guidelines

### 7.1 Alert Copy

| Element | Copy |
|---------|------|
| Headline | "Streak at risk" |
| Body | "Your partner hasn't checked in yet." |
| Helper | "They have until midnight to check in." |

### 7.2 Status Labels

| Column | Label | Sublabel |
|--------|-------|----------|
| User | "You" | "Checked In" |
| Partner | "Partner" | "Waiting..." |

### 7.3 Button Labels

| State | Label |
|-------|-------|
| Available | "🔔 Nudge Partner" |
| Sent | "✓ Nudge Sent" |
| Limit | "No nudges left" |

### 7.4 Nudge Indicators

| Context | Text |
|---------|------|
| Available | "Nudges remaining: X today" |
| Cooldown | "You can nudge again in X hours" |
| Limit | "Upgrade to Plus for more nudges" |

### 7.5 Toast Messages

| Action | Message |
|--------|---------|
| Nudge sent | "Nudge sent! 🔔" |
| Partner checks in | "Partner checked in! Streak extended! ✓" |

### 7.6 Tone Guidelines

| ❌ Avoid | ✓ Use |
|---------|-------|
| "Partner FAILED" | "Partner hasn't checked in" |
| "Streak WILL break" | "Streak at risk" |
| "Their fault" | "You showed up" (supportive) |

---

## 8. Accessibility

### 8.1 Screen Reader
- Alert: "Alert. Streak at risk. Your partner hasn't checked in yet."
- Status: "Your status: Checked in, Upper Body. Partner status: Waiting."
- Button: "Nudge Partner button. Nudges remaining: 2 today."

### 8.2 Touch Targets
- Nudge button: 52px height, full width

### 8.3 Color Contrast
- Amber (#E6A43A) meets WCAG AA on white background

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Alert banner (amber) | Critical | ☐ |
| Status split card | Critical | ☐ |
| Nudge button 3 states | Critical | ☐ |
| Nudge limit logic | Critical | ☐ |
| Cooldown timer | High | ☐ |
| Bell animation | Medium | ☐ |
| Upgrade modal trigger | High | ☐ |
| Toast notifications | High | ☐ |
| Pulsing animation | Medium | ☐ |
| Dark mode | Medium | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 14 — Streak System, Nudges |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Section 2.1, 8.3 — Warning Colors |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.D — Streak States |
| [Unora_Spec_12_StreakDetail.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_12_StreakDetail.md) | Parent screen (State D) |
| Unora_Spec_16_PaymentState.md (planned) | Next state if partner doesn't recover |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
