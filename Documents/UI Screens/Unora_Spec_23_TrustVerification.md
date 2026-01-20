# Screen #23: Trust & Verification

**Version:** 1.0  
**Date:** January 2026  
**Location:** Profile Tab → Trust & Verification  
**Design System Reference:** [Section 4.X: Verification UI](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md)  
**PRD Reference:** [Section 10: Progressive Verification](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)

---

## 1. Screen Purpose

Centralize all trust-related verification actions in a dedicated, calm, non-urgent settings screen. This screen allows users to view their verification status, complete optional verifications, and understand how trust is earned through consistency.

**Key Principles:**
- Zero urgency or fear language
- No red warnings or countdown timers
- Verification feels optional but valuable
- Matches Profile tab aesthetics

---

## 2. Entry Points

| Entry Point | Navigation |
|-------------|------------|
| Profile Tab → Settings Section | Tap "Trust & Verification" row |
| Onboarding Completion (Optional) | Post-onboarding optional prompt |

---

## 3. Screen Layout

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  ← Back                   Trust & Verification                              │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌───────────────────────────────────────────────────────────────────────┐  │
│  │                                                                       │  │
│  │   "Trust builds through consistency.                                  │  │
│  │    You're always in control."                                         │  │
│  │                                                                       │  │
│  └───────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│  ─────────────────────────────────────────────────────────────────────────  │
│                                                                             │
│  YOUR TRUST STATUS                                                          │
│                                                                             │
│  ✓ Phone verified                                                           │
│    └── Verified on Jan 9, 2026                                              │
│                                                                             │
│  ✓ Photos submitted                                                          │
│    └── 4 of 6 photos uploaded                                               │
│    └── [Add more photos →]                                                  │
│                                                                             │
│  ✓ Photo quality verified                                                   │
│    └── AI verified your photos meet quality standards                       │
│                                                                             │
│  ⏳ Consistency                                                              │
│    └── Day 7 of first streak                                                │
│    └── Earned through completing a 15-day streak                            │
│                                                                             │
│  ○ Government ID (Optional)                                                 │
│    └── Unlocks profile flexibility features                                 │
│    └── Available after streak completion or 30 days                         │
│    └── [Learn more →]                                                       │
│                                                                             │
│  ─────────────────────────────────────────────────────────────────────────  │
│                                                                             │
│  TRUST BOOSTERS (Optional)                                                  │
│                                                                             │
│  ✓ Google Account                   [Connected]                             │
│  ○ Apple Account                    [Connect →]                             │
│  ○ LinkedIn                         [Connect →]                             │
│                                                                             │
│  ─────────────────────────────────────────────────────────────────────────  │
│                                                                             │
│  [Privacy & Data →]                                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 4. Component Specifications

### 4.1 Header

| Property | Value |
|----------|-------|
| Height | 56px |
| Background | `var(--surface-primary)` |
| Left | Back arrow (← icon), tap returns to Profile |
| Center | "Trust & Verification" in `var(--h3-size)` 20px |

### 4.2 Hero Quote

| Property | Value |
|----------|-------|
| Background | `var(--surface-secondary)` |
| Padding | 24px |
| Border Radius | 12px |
| Text | Body-1, 16px, center-aligned |
| Color | `var(--unora-ink-secondary)` |
| Line Height | 1.6 |

**Copy:**
> "Trust builds through consistency. You're always in control."

### 4.3 Trust Status Section

Each verification row follows this structure:

```
┌─────────────────────────────────────────────────────────────────┐
│ [Icon]  [Title]                                                 │
│         └── [Subtext line 1]                                    │
│         └── [CTA if applicable]                                 │
└─────────────────────────────────────────────────────────────────┘
```

**Icon States:**

| State | Icon | Color |
|-------|------|-------|
| Complete | ✓ | `var(--semantic-success)` #51CF66 |
| In Progress | ⏳ | `var(--unora-ink-secondary)` #4A4A4A |
| Optional/Pending | ○ | `var(--unora-ink-muted)` #A3A3A3 |

**Row Specifications:**

| Property | Value |
|----------|-------|
| Row Height | Auto (min 64px) |
| Icon Size | 24px |
| Title | Body-1, 16px, `var(--unora-ink-primary)` |
| Subtext | Body-2, 14px, `var(--unora-ink-tertiary)` |
| CTA | Body-2, 14px, `var(--unora-primary-accent)` |
| Vertical Spacing | 8px between rows |

### 4.4 Trust Boosters Section

| Property | Value |
|----------|-------|
| Section Title | Caption, 12px, uppercase, `var(--unora-ink-muted)` |
| Row Height | 56px |
| Left | Platform icon (24px) + Platform name |
| Right | "Connected" (green) or "Connect →" (accent) |

### 4.5 Privacy Link

| Property | Value |
|----------|-------|
| Height | 48px |
| Text | Body-2, 14px, `var(--unora-ink-secondary)` |
| Right | Arrow icon → |
| Tap | Opens Privacy & Data settings |

---

## 5. Interaction Logic

### 5.1 Photo Quality Status

Photo quality verification is **passive and automatic**. Users do NOT need to take any action.

```
SYSTEM LOGIC:   AI evaluates photo quality during upload
                │
                ▼
OUTCOME:        Pass → Row displays ✓ "Photo quality verified"
                Fail → User prompted to re-upload during upload flow

NOTE: No user action on this screen. Status is read-only.
```

### 5.2 Government ID Row

```
CONDITION:      Is user eligible for Gov ID verification?
                │
        ┌───────┴───────┐
        │               │
        ▼               ▼
  ✓ ELIGIBLE       ✗ NOT YET
  (streak done      (no streak,
   or 30 days)      <30 days)
        │               │
        ▼               ▼
  Show CTA:       Show text:
  [Complete       "Available after
  verification →]  streak completion
  Opens Gov ID    or 30 days"
  flow            No CTA visible
```

### 5.3 Trust Boosters

```
USER ACTION:    User taps "Connect →" for Apple/LinkedIn
                │
                ▼
SYSTEM LOGIC:   Open OAuth flow for selected platform
                │
                ▼
OUTCOME:        On success → Row updates to ✓ "Connected"
                On cancel → No change
                On error → Toast: "Connection failed. Try again."
```

---

## 6. States

### 6.1 New User (Post-Onboarding)

```
✓ Phone verified
✓ Photos submitted (3 of 6)
✓ Photo quality verified
⏳ Consistency (Day 1)
○ Government ID (Available after streak)
○ Trust Boosters (None connected)
```

### 6.2 Mid-Streak User

```
✓ Phone verified
✓ Photos submitted (6 of 6)
✓ Photo quality verified
⏳ Consistency (Day 8)
○ Government ID (Available after streak)
✓ Trust Boosters (Google connected)
```

### 6.3 Fully Verified User

```
✓ Phone verified
✓ Photos submitted (6 of 6)
✓ Photo quality verified
✓ Consistency earned
✓ Government ID verified
✓ Trust Boosters (Google, Apple connected)
```

---

## 7. Copy Guidelines

### 7.1 Approved Copy

| Element | Copy |
|---------|------|
| Hero Quote | "Trust builds through consistency. You're always in control." |
| Photo Quality Subtext | "AI verified your photos meet quality standards" |
| Consistency Subtext | "Earned through completing a 15-day streak" |
| Gov ID Subtext | "Unlocks profile flexibility features" |
| Privacy Link | "Privacy & Data" |

### 7.2 Prohibited Copy

| ❌ Never Use | Why |
|--------------|-----|
| "Verify now" | Creates urgency |
| "Your account is unverified" | Punitive framing |
| "Required to continue" | Implies gate |
| "KYC" | Technical jargon |
| "Complete verification to unlock profiles" | False promise |
| "Liveness check" | Feature does not exist |

---

## 8. Accessibility

| Requirement | Implementation |
|-------------|----------------|
| Screen Reader | All status icons have alt text (e.g., "Complete", "In progress") |
| Touch Targets | All CTAs minimum 44px height |
| Color Contrast | All text meets WCAG AA (4.5:1 for body, 3:1 for large text) |
| Focus Order | Top-to-bottom, logical tab order |

---

## 9. Animation & Transitions

| Element | Animation |
|---------|-----------|
| Screen Entry | Slide in from right, 250ms ease-out |
| Status Update | Fade + slight scale on icon change, 300ms |
| CTA Tap | Subtle press feedback (scale 0.98), 100ms |

---

## 10. Edge Cases

| Scenario | Handling |
|----------|----------|
| No internet | Show cached status, disable CTAs, subtle offline indicator |
| OAuth error | Toast: "Connection failed. Please try again." |
| Photo quality fail | Handled during upload flow, not on this screen |
| Gov ID rejected | Show: "Verification didn't go through. Let's try again." No red |
