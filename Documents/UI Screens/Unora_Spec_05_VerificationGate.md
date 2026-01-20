# Unora UI Specification — Photo & Trust Setup

**Document ID:** Spec-05  
**Screen Name:** Photo & Trust Setup  
**Version:** 2.0 (Progressive Verification)  
**Date:** January 2026  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name
**Photo & Trust Setup** (formerly: Verification Gate)

### 1.2 User Flow Reference
**Phase 1: Progressive Onboarding** — This screen collects mandatory photos and offers optional trust boosters. There is **no hard gate**. Users must upload 3+ photos to proceed.

**Sequence:**
```
Phone Auth → [Photo & Trust Setup] → Profile Creation → Server Selection → Discovery
                     ↓
              (Retry if photos invalid)
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.A.2](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose
Collect mandatory photos for personhood verification and Day 15 reveal. Offer optional trust boosters (Google/Apple/LinkedIn) as lightweight identity signals.

### 1.4 Primary Action
**Upload 3–6 photos** (mandatory) and optionally connect trust boosters.

---

## 2. Low-Fidelity Wireframe (ASCII)

### 2.1 Photo Upload Screen

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│  ← Back                                                     │
│                                                             │
│                                                             │
│              Add your photos                                │  ← H2
│                                                             │
│         Your photos stay hidden until Day 15.               │  ← Body
│         Only you can see them until then.                   │
│                                                             │
│                                                             │
│   ┌───────────┐  ┌───────────┐  ┌───────────┐               │
│   │           │  │           │  │           │               │
│   │    +      │  │    +      │  │    +      │               │  ← Photo grid
│   │           │  │           │  │           │               │     3 required
│   │           │  │           │  │           │               │
│   └───────────┘  └───────────┘  └───────────┘               │
│                                                             │
│   ┌───────────┐  ┌───────────┐  ┌───────────┐               │
│   │           │  │           │  │           │               │
│   │    +      │  │    +      │  │    +      │               │  ← Optional
│   │           │  │           │  │           │               │     slots
│   │           │  │           │  │           │               │
│   └───────────┘  └───────────┘  └───────────┘               │
│                                                             │
│         Minimum 3 photos • Include at least one             │  ← Caption
│         clear photo of your face                            │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │              Continue                                │   │  ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │     (disabled until 3+)
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.2 Photo Validation Error

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│              Add your photos                                │
│                                                             │
│         Your photos stay hidden until Day 15.               │
│                                                             │
│                                                             │
│   ┌───────────┐  ┌───────────┐  ┌───────────┐               │
│   │   [✓]     │  │   [⚠️]    │  │    +      │               │
│   │  img 1    │  │  img 2    │  │           │               │
│   │           │  │ (blurry)  │  │           │               │
│   └───────────┘  └───────────┘  └───────────┘               │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  ⚠️ Photo 2 appears blurry. Try a clearer shot?    │   │  ← Inline warning
│   │                                             [Retry] │   │     (yellow, not red)
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │              Continue                                │   │  ← Still enabled
│   └─────────────────────────────────────────────────────┘   │     if 3+ valid
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### 2.3 Optional Trust Boosters

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│  ← Back                                                     │
│                                                             │
│              ✓ Photos uploaded                              │  ← Progress indicator
│                                                             │
│              Add trust signals                              │  ← H2
│              (optional)                                     │
│                                                             │
│         Connect accounts to build trust.                    │  ← Body
│         These are visible only to you.                      │
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  [G]  Continue with Google                          │   │  ← OAuth button
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  []  Continue with Apple                           │   │  ← OAuth button
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  [in] Connect LinkedIn                              │   │  ← OAuth button
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
│         Skip for now →                                      │  ← Skip link
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │              Continue                                │   │  ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 3. Typography & Styling

| Element | Token | Size | Weight | Color |
|---------|-------|------|--------|-------|
| Headline | `--h2-size` | 24px | 600 | `--unora-ink-primary` |
| Body | `--body-1` | 16px | 400 | `--unora-ink-secondary` |
| Caption | `--caption` | 12px | 400 | `--unora-ink-tertiary` |
| CTA | `--button-lg` | 16px | 600 | White on `--unora-primary-accent` |
| Skip Link | `--body-2` | 14px | 400 | `--unora-primary-accent` |

---

## 4. Component Specifications

### 4.1 Photo Grid

| Property | Value |
|----------|-------|
| Grid | 3 columns, 2 rows |
| Slot Size | (screen width - 48px padding - 16px gaps) / 3 |
| Aspect Ratio | 1:1 (square) |
| Border Radius | 12px |
| Empty State | Dashed border, + icon centered |
| Filled State | Photo thumbnail, subtle checkmark overlay |
| Error State | Yellow border, ⚠️ icon overlay |

### 4.2 Photo Upload Behavior

```
USER TAPS EMPTY SLOT:
  │
  ▼
OPEN SYSTEM PICKER:
  ├── Camera
  └── Photo Library
  │
  ▼
VALIDATE PHOTO:
  ├── Check blur (threshold: 20%)
  ├── Check face detection
  └── Check file size (max 10MB)
  │
  ▼
ON SUCCESS:
  └── Show thumbnail with ✓ overlay
  │
ON FAIL:
  └── Show thumbnail with ⚠️ overlay + inline message
```

### 4.3 Primary CTA Button

| Property | Value |
|----------|-------|
| Width | 100% - 32px horizontal margin |
| Height | 52px |
| Border Radius | 26px (pill) |
| **Disabled State** | `--surface-disabled`, text `--unora-ink-muted` |
| **Enabled State** | `--unora-primary-accent`, text white |
| Condition | Enabled when: 3+ photos valid AND 1+ face detected |

### 4.4 Trust Booster Buttons

| Property | Value |
|----------|-------|
| Width | 100% - 32px horizontal margin |
| Height | 52px |
| Border Radius | 12px |
| Background | `--surface-elevated` |
| Border | 1px solid `--border-subtle` |
| Icon | Platform logo, 24px |
| Text | Body-1, left-aligned |

---

## 5. Interaction Logic

### 5.1 Photo Validation

```
VALIDATION RULES:
├── Minimum: 3 photos
├── Maximum: 6 photos
├── Face Required: At least 1 photo must have detectable face
├── Blur Check: Photos below quality threshold flagged (not blocked)
└── File Size: Max 10MB per photo
```

### 5.2 Continue Flow

```
USER TAPS CONTINUE:
  │
  ▼
CHECK PHOTOS:
  │
  ├─ < 3 photos? → Error toast: "Add at least 3 photos"
  │
  ├─ No face detected? → Error toast: "Include a clear face photo"
  │
  └─ Valid? → Navigate to Trust Boosters step
```

### 5.3 Skip Trust Boosters

```
USER TAPS "Skip for now":
  │
  ▼
SYSTEM:
  └── Navigate to Profile Creation
  └── Trust boosters marked as skipped (can add later)
  └── No penalty, no warning
```

---

## 6. States

### 6.1 Initial State

- Photo grid: All slots empty
- Continue button: Disabled
- Caption: "Minimum 3 photos • Include at least one clear photo of your face"

### 6.2 Partial Upload (1-2 photos)

- Photo grid: Some slots filled
- Continue button: Disabled
- Caption: "2 more photos needed"

### 6.3 Ready State (3+ valid photos)

- Photo grid: 3+ slots filled with ✓
- Continue button: Enabled
- Caption: Updates to "Looking good! Continue when ready"

### 6.4 Error State (photo validation failed)

- Affected slot: Yellow border + ⚠️ icon
- Inline message: "Photo appears blurry. Try a clearer shot?"
- Continue button: Enabled if 3+ other photos valid

### 6.5 No Face Detected

- Continue button: Disabled
- Caption: "Include at least one clear photo of your face"

---

## 7. Copy Guidelines

### 7.1 Approved Copy

| Element | Copy |
|---------|------|
| Headline | "Add your photos" |
| Privacy Reassurance | "Your photos stay hidden until Day 15. Only you can see them until then." |
| Photo Caption | "Minimum 3 photos • Include at least one clear photo of your face" |
| Trust Boosters Headline | "Add trust signals (optional)" |
| Trust Boosters Body | "Connect accounts to build trust. These are visible only to you." |
| Skip Link | "Skip for now" |

### 7.2 Error Messages

| Scenario | Copy |
|----------|------|
| Photo blurry | "Photo appears blurry. Try a clearer shot?" |
| No face detected | "Include at least one clear photo of your face" |
| Upload failed | "Upload didn't complete. Let's try again." |

---

## 8. Accessibility

| Requirement | Implementation |
|-------------|----------------|
| Screen Reader | Photo slots announce state: "Photo 1: empty" or "Photo 1: uploaded" |
| Touch Targets | All slots and buttons minimum 44px |
| Color Contrast | All text meets WCAG AA |
| Focus Order | Top-to-bottom, left-to-right for grid |

---

## 9. Animation & Transitions

| Element | Animation |
|---------|-----------|
| Photo Upload | Fade in 200ms, slight scale up from 0.95 |
| Validation Check | Checkmark draws in, 300ms |
| Error Overlay | Fade in 150ms, gentle shake 2x |
| Screen Transition | Slide left to next step, 250ms ease-out |

---

## 10. Privacy & Security

| Aspect | Implementation |
|--------|----------------|
| Photo Storage | Encrypted at rest (AES-256) |
| Photo Visibility | Private until Day 15 reveal |
| OAuth Tokens | Stored securely, no password access |
| Data Minimization | Only metadata extracted from ID providers |

---

## 11. Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | Jan 2026 | Initial DigiLocker/Aadhaar verification gate |
| 2.0 | Jan 2026 | **Major rewrite**: Removed hard gate, replaced with photo upload + optional trust boosters |
