# Unora UI Specification — Profile Enrichment

**Document ID:** Spec-04  
**Screen Name:** Profile Creation  
**Version:** 2.0 (Progressive Verification)  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Profile Creation** — Multi-step form for user-provided profile data

### 1.2 User Flow Reference
**Phase 1: Progressive Onboarding** — Profile Creation occurs AFTER Photo & Trust Setup. All fields are user-provided and editable.

**Sequence:**
```
Phone Auth → Photo & Trust Setup → [Profile Creation] → Server Selection → Discovery
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.A.3](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Collect all profile information (Name, DOB, Gender, City, Education, Hobbies) to enable accurate matching and meaningful gradual reveals. All fields are user-provided and editable.

### 1.4 Primary Action
**Complete each step** and submit enrichment data via "Continue" buttons.

---

## 2. Low-Fidelity Wireframe (ASCII)

### 2.1 Step Layout (Generic)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   ←                              Step 1 of 4                │  ← Header
│                                                             │
│   ████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░               │  ← Progress bar
│                                                             │
│         What's your name?                                   │  ← Headline (H2)
│                                                             │
│         This is how you'll appear on Day 15.                │  ← Subtitle
│                                                             │
│   Name                                                      │  ← Label
│   ┌─────────────────────────────────────────────────────┐   │
│   │  [First name + Last name]                           │   │  ← Input
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                    Continue                         │   │  ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
├─────────────────────────────────────────────────────────────┤
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Step Layout (Generic)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   ←                              Step 1 of 4                │  ← Header
│                                                             │
│   ████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░               │  ← Progress bar
│                                                             │
│         Where are you based?                                │  ← Headline (H2)
│                                                             │
│         We'll show you people nearby.                       │  ← Subtitle
│                                                             │
│   City                                                      │  ← Label
│   ┌─────────────────────────────────────────────────────┐   │
│   │  Bangalore                                          │   │  ← Input
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                    Continue                         │   │  ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
├─────────────────────────────────────────────────────────────┤
└─────────────────────────────────────────────────────────────┘
```

### 2.3 Profile Creation Steps Overview

| Step | Title | Data Collected | Notes |
|------|-------|----------------|-------|
| 1 | Identity | Name, DOB, Gender | All user-provided, all editable |
| 2 | Location | City, Locality | Required |
| 3 | Background | Education, Profession, Religion (opt) | Profession optional |
| 4 | Hobbies | Hobby anchors + micro-descriptions | Minimum 2 required |
| 4.5 | Personality Context | Situational responses | ❌ **Optional** — Can skip |

> **Note:** Photos are collected in the preceding Photo & Trust Setup screen (Spec-05).

### 2.4 Optional Personality Context Step

> [!IMPORTANT]
> This step is **completely optional**. Skipping does not block onboarding, discovery, or any core functionality.

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   ←                                                         │  ← Back button
│                                                             │
│                                                             │
│         Add personality context                             │  ← Headline
│                (optional)                                   │  ← Subhead
│                                                             │
│         Help others understand how you connect.             │  ← Description
│                                                             │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                  Get Started →                      │   │  ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                    I'll do this later                       │  ← Secondary link
│                                                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Layout Specs:**

| Element | Specification |
|---------|---------------|
| **Headline** | "Add personality context" — Outfit 22px / 600 |
| **Subhead** | "(optional)" — Inter 14px / 400, tertiary color |
| **Description** | "Help others understand how you connect." — Inter 16px / 400 |
| **Primary CTA** | "Get Started →" — Full-width primary button |
| **Secondary link** | "I'll do this later" — Tertiary text link, centered |

**Behavior:**

| Action | Result |
|--------|--------|
| Tap "Get Started →" | Opens Personality Setup Screen (Spec-24) |
| Tap "I'll do this later" | Proceeds directly to Server Selection |
| Tap ← (Back) | Returns to Hobbies step |

**Critical Rules:**
- ❌ No progress bar on this step (it is optional, not part of required flow)
- ❌ No urgency language ("complete your profile", "don't miss out")
- ❌ No penalty for skipping (discovery fully accessible)
- ✓ Always skippable, always calm



### 2.5 Photo Upload Step

```
┌─────────────────────────────────────────────────────────────┐
│   ←                              Step 3 of 4                │
│   ██████████████████████████████████████████████████████    │
│                                                             │
│         Add your photos                                     │
│                                                             │
│         Photos stay hidden until Day 15.                    │
│         They're only for your future reveal.                │
│                                                             │
│   ┌─────────┐  ┌─────────┐  ┌─────────┐                     │
│   │         │  │         │  │         │                     │
│   │   📷    │  │   📷    │  │   📷    │                     │  ← Photo slots
│   │   +     │  │   +     │  │   +     │                     │
│   │         │  │         │  │         │                     │
│   └─────────┘  └─────────┘  └─────────┘                     │
│                                                             │
│   ┌─────────┐  ┌─────────┐  ┌─────────┐                     │
│   │         │  │         │  │         │                     │
│   │   +     │  │   +     │  │   +     │                     │
│   │         │  │         │  │         │                     │
│   └─────────┘  └─────────┘  └─────────┘                     │
│                                                             │
│         Minimum 3 photos required                           │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                    Continue                         │   │
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
PROFILE CREATION CONTAINER
├── Position: fixed, 100vw × 100vh
├── Background: var(--surface-background) → #FAFAF8
├── Display: flex, column
│
├── [HEADER] — 56px height
│   ├── Back button: left-aligned
│   └── Step indicator: right-aligned
│
├── [PROGRESS BAR] — 4px height
│   └── Fill: var(--unora-primary-accent)
│
├── [CONTENT AREA] — flex: 1
│   ├── Padding: 20px horizontal
│   ├── Padding-top: 24px
│   └── Gap: 24px
│
└── [ACTION AREA]
    ├── Padding: 20px
    └── Padding-bottom: 40px + safe-area

Premium Dark Mode (Default):
├── Background: var(--pdm-background) → #0D0D0F
├── Progress fill: Gold gradient with glow
├── Input surfaces: Elevated charcoal with inner glow
└── Photo slots: Subtle border glow on hover
```

### 3.2 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Progress bar** | Gold fill with glow: `0 0 8px rgba(196,167,125,0.4)` |
| **Input fields** | Surface `#1A1A1E`, border `#2A2A2E`, gold focus ring |
| **Photo slots (empty)** | Dashed border `#2A2A2E`, hover: gold outline |
| **Photo slots (filled)** | Subtle inner glow `inset 0 1px 0 rgba(255,255,255,0.03)` |
| **Verified badge** | Teal with subtle glow for "Verified" indicator |

**Progress Bar Glow:**
```css
/* Premium progress bar with gold glow */
.progress-fill {
  background: linear-gradient(90deg, var(--accent-gold-deep), var(--accent-gold));
  box-shadow: 0 0 12px rgba(196, 167, 125, 0.35);
  transition: width 300ms ease-out;
}
```



### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen padding | `var(--padding-screen)` | 20px |
| Progress bar height | — | 4px |
| Section gap | `var(--space-6)` | 24px |
| Input group gap | `var(--space-5)` | 20px |
| Label margin-bottom | `var(--space-2)` | 8px |
| Photo grid gap | `var(--space-3)` | 12px |

### 3.3 Z-Index Layers

| Layer | Z-Index | Contents |
|-------|---------|----------|
| Background | 0 | Screen background |
| Content | 1 | Form elements |
| Overlay | 10 | Date picker, photo modal |
| System | 100+ | Status bar |

---

## 4. Component Inventory

### 4.1 Progress Bar

```
PROGRESS BAR
├── Width: 100%
├── Height: 4px
├── Background: var(--border-subtle) → #E8E8E6
├── Fill: var(--unora-primary-accent) → #C4A77D
├── Border radius: var(--radius-full)
└── Animation: width transition 300ms ease-out
```

### 4.2 Typography

| Element | Font | Weight | Size | Color |
|---------|------|--------|------|-------|
| Step indicator | Inter | 500 | 14px | `--unora-ink-tertiary` |
| Headline (H2) | Outfit | 600 | 22px | `--unora-ink-primary` |
| Subtitle | Inter | 400 | 16px | `--unora-ink-secondary` |
| Input label | Inter | 500 | 14px | `--unora-ink-secondary` |
| Input text | Inter | 400 | 16px | `--unora-ink-primary` |
| Helper/privacy note | Inter | 400 | 12px | `--unora-ink-tertiary` |

### 4.3 Input Components

#### Text Input
| Property | Value |
|----------|-------|
| Height | 52px |
| Border | 1.5px solid `var(--border-medium)` |
| Border radius | `var(--radius-md)` → 12px |
| Background | `var(--surface-card)` |
| Focused | 2px `var(--unora-primary-accent)` |

#### Verified Input Field (Name, Age, Gender)

These fields are auto-populated from the verified ID and are **non-editable**.

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  🔒 Priya Sharma                          ✓ Verified        │
│                                                             │
└─────────────────────────────────────────────────────────────┘

VERIFIED INPUT FIELD
├── Height: 52px
├── Background: var(--surface-disabled) → #F0F0EE (grayed-out)
├── Border: 1px solid var(--border-subtle) → #E8E8E6
├── Border radius: var(--radius-md) → 12px
├── Text color: var(--unora-ink-secondary) → #4A4A4A
├── Font: Inter 16px / 400
├── Lock icon: 🔒 Left-aligned, 16px, var(--unora-ink-tertiary)
├── Verified badge: "✓ Verified" Right-aligned, green (#4A9B8C)
├── Pointer events: none (non-interactive)
└── Cursor: not-allowed
```

**Behavior:**
- Fields are read-only and cannot be edited
- Tapping shows toast: "This is verified from your ID"
- No keyboard appears on tap
- Visual treatment clearly indicates locked state

**Dark Mode:**
- Background: var(--dm-surface-disabled) → #1A1A1A
- Border: var(--dm-border-subtle) → #2D2D2D

#### Photo Upload Grid
| Property | Value |
|----------|-------|
| Grid | 3 columns |
| Gap | 12px |
| Cell size | (screenWidth - 40px - 24px) / 3 |
| Aspect ratio | 1:1 (square) |
| Border radius | `var(--radius-md)` → 12px |
| Empty state | Dashed border, + icon centered |
| Filled | Image cover, X delete button |

### 4.4 Buttons

#### Primary CTA ("Continue")
| Property | Value |
|----------|-------|
| Height | 52px |
| Background | `var(--unora-primary-accent)` |
| Text | White, Inter 16px/600 |
| Disabled | Opacity 0.4 |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| Tap | Back arrow | Previous step or exit confirmation |
| Tap | Input field | Focus, show keyboard |
| Tap | Gender pill | Select option |
| Tap | DOB field | Open date picker |
| Tap | Photo slot | Open camera/gallery picker |
| Tap | Photo X | Remove photo |
| Tap | Continue | Validate & proceed |
| Swipe | Between steps | Navigate (optional) |

### 5.2 Behaviors

```
STEP NAVIGATION FLOW

USER taps "Continue"
    │
    ▼
┌─────────────────────────────────────────┐
│ [VALIDATE] Are required fields filled?  │
└─────────────────────────────────────────┘
    │                    │
    ▼                    ▼
  ✓ YES                ✗ NO
    │                    │
    ▼                    ▼
Save step data      Show error state
Progress bar fills  Shake invalid inputs
Slide to next step  Focus first error
    │
    ▼
[Final step?]
    │         │
    ▼         ▼
  YES        NO
    │         │
    ▼         └──→ Next step screen
Submit all data
Navigate to Profile Review
```

### 5.3 Data Validation

| Field | Validation | Error Message |
|-------|-----------|---------------|
| City | Required | "Please select your city" |
| Profession | Required | "Please add your profession" |
| Photos | Min 3 | "Add at least 3 photos" |
| Hobbies | Min 2 | "Add at least 2 hobbies" |

### 5.4 Transitions

| Transition | Animation |
|------------|-----------|
| Step forward | Slide left, 300ms ease-out |
| Step backward | Slide right, 300ms ease-out |
| Progress bar | Width animate, 300ms |
| Error shake | translateX ±4px, 150ms |
| Photo add | Scale up + fade, 200ms |

---

## 6. State Definitions

### 6.1 State Matrix

| State | Appearance | Condition |
|-------|------------|-----------|
| Default | Empty form | Step loaded |
| In Progress | Partial fill | User entering data |
| Valid | Continue enabled | Required fields complete |
| Error | Red borders, error text | Validation failed |
| Loading | Spinner on button | Submitting data |

### 6.2 Input Error State

```
Error Input:
├── Border: 1.5px solid var(--feedback-error)
├── Label: unchanged
├── Error text: var(--feedback-error), 12px
├── Icon: Warning inline
└── Animation: Shake 150ms
```

### 6.3 Photo States

| State | Visual |
|-------|--------|
| Empty | Dashed border, + icon, "Add" |
| Uploading | Progress indicator, dimmed |
| Filled | Image with delete X |
| Error | Red dashed border |

---

## 7. Content & Copy Guidelines

### 7.1 Step Content

| Step | Headline | Subtitle |
|------|----------|----------|
| Entry | "Let's complete your profile" | "We've got your basic details. Now add a few more." |
| 1 | "Where are you based?" | "We'll show you people nearby." |
| 2 | "A bit about you" | "Optional details that help find your match." |
| 3 | "Add your photos" | "Photos stay hidden until Day 15." |
| 4 | "What are you into?" | "Pick hobbies that define you." |

### 7.2 Verified Data Display

| Element | Copy |
|---------|------|
| Section header | "✓ Verified Details" |
| Fields | Name, Age, Gender with lock icons |
| Badge | "[Verified]" next to each field |
| Helper | "These details are verified and cannot be changed" |

### 7.3 Privacy Reassurance

Key steps include subtle privacy messaging:

| Step | Privacy Note |
|------|--------------|
| Photos | 🔒 "Photos stay hidden until the final reveal" |
| Location | 🔒 "Only your city is shown, never your locality" |

### 7.4 Button Labels

| Context | Label |
|---------|-------|
| Entry screen | "Continue" |
| Steps 1-3 | "Continue" |
| Step 4 (hobbies) | "Complete Profile" |
| Loading | "Saving..." |

### 7.5 Tone Guidelines

| Principle | Application |
|-----------|-------------|
| Presence over Performance | No "quick setup" language |
| Privacy First | Reassurance on every step |
| Clarity | Simple questions, one concept per screen |
| Warmth | "What's your name?" not "Enter name" |

### 7.6 Error Messages

| Error | Message |
|-------|---------|
| No city | "Please select your city" |
| Too few photos | "Add at least 3 photos to continue" |
| Upload failed | "Couldn't upload photo. Try again." |
| Too few hobbies | "Add at least 2 hobbies to continue" |

---

## 8. Accessibility

### 8.1 Screen Reader
- Entry: "Your verified details: [Name], [Age], [Gender]. Let's add a few more details."
- Announce: "Step X of 4. [Headline]"
- Progress bar: aria-valuenow, aria-valuemin, aria-valuemax
- Labels connected via aria-labelledby

### 8.2 Touch Targets
- All interactive elements: 44px minimum
- Photo grid cells: Full cell tappable

### 8.3 Keyboard (Web)
- Tab navigation through inputs
- Enter to submit step

---

## 9. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Verified data display (locked fields) | Critical | ☐ |
| 4-step enrichment flow | Critical | ☐ |
| Progress bar animation | High | ☐ |
| All input types (text, selector, photo, hobbies) | Critical | ☐ |
| Validation per step | Critical | ☐ |
| Privacy messaging | High | ☐ |
| Photo upload (min 3) | Critical | ☐ |
| Hobby picker (min 2) | Critical | ☐ |
| Error states | High | ☐ |
| Dark mode | Medium | ☐ |
| Back navigation with confirmation | High | ☐ |
| Data persistence between steps | Critical | ☐ |

---

## 10. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Section 10.4 - Profile Enrichment |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Input components, tokens |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Section 2.A.4 - Profile Enrichment |
| [Unora_Spec_05_VerificationGate.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_05_VerificationGate.md) | Previous screen (with Auto-fill) |

---

**Last updated:** January 2026

*This specification is developer-ready. Deviations require design review.*
