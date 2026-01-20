# Unora UI Specification — Auth / Login

**Document ID:** Spec-03  
**Screen Name:** Auth / Login  
**Version:** 1.0  
**Date:** January 2026  
**Classification:** Internal / Engineering / Design  
**Status:** Developer Ready

---

## 1. Metadata & Overview

### 1.1 Screen Name

**Auth / Login** (also: Sign In, Returning User Authentication)

### 1.2 User Flow Reference

**Phase 1: Verified Onboarding** — The Auth/Login screen uses **Phone Number + OTP** as the primary authentication method for both new and returning users. This supports Unora's "Verification-First" architecture for the India market.

**Sequence Position (Returning Users):**
```
Splash Screen → Welcome Landing → [Phone + OTP Auth] → Discovery/Home (if verified)
                     ↓                                           ↓
               Get Started →                            → Resume Onboarding (if incomplete)
```

**Sequence Position (New Users):**
```
Splash → Welcome → Get Started → [Phone + OTP] → Photo & Trust Setup → Profile Creation
                                       ↓
                              (Progressive Verification)
```

**Reference:** [Unora_UserFlow_Logic.md — Section 2.A (Onboarding & Progressive Verification)](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 1.3 Purpose

The Auth/Login screen enables users to securely authenticate using their phone number and OTP verification. This phone-anchored identity system supports downstream verification requirements and ensures a consistent, trustworthy user base.

### 1.4 Primary Action

**Enter phone number and verify via OTP** to authenticate and access the user's account or begin the verification flow.

---

## 2. Low-Fidelity Wireframe (ASCII)

### 2.1 Login Screen — Phone + OTP Method (Primary)

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │   ← System status bar
│                                                             │
│   ←                                                         │   ← Back navigation
│                                                             │
│                                                             │
│                                                             │
│                   ┌───────────────┐                         │
│                   │               │                         │
│                   │   [ UNORA ]   │                         │   ← App Logo (48px)
│                   │               │                         │
│                   └───────────────┘                         │
│                                                             │
│               Welcome back.                                 │   ← Headline (H2)
│                                                             │
│                                                             │
│   Phone number                                              │   ← Input label
│   ┌─────────────────────────────────────────────────────┐   │
│   │  +91  |                                             │   │   ← Phone input
│   └─────────────────────────────────────────────────────┘   │     (+91 prefix)
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                     Get OTP                         │   │   ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
│                                                             │
│               Don't have an account?                        │   ← Caption
│                    Get Started →                            │   ← Link to onboarding
│                                                             │
│                                                             │
│                 Use Password Login →                        │   ← Tertiary fallback
│                                                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │   ← Safe area (bottom)
└─────────────────────────────────────────────────────────────┘
```

### 2.2 OTP Verification Screen

```
┌─────────────────────────────────────────────────────────────┐
│                     [Status Bar]                            │
│                                                             │
│   ←                                                         │   ← Back navigation
│                                                             │
│                                                             │
│                                                             │
│                   ┌───────────────┐                         │
│                   │               │                         │
│                   │   [ UNORA ]   │                         │   ← App Logo (48px)
│                   │               │                         │
│                   └───────────────┘                         │
│                                                             │
│               Enter verification code                       │   ← Headline (H2)
│                                                             │
│   We sent a 6-digit code to                                 │   ← Body text
│   +91 98765 ••••5                                           │   ← Masked phone
│                                                             │         (specific number)
│                                                             │
│       ┌───┐  ┌───┐  ┌───┐  ┌───┐  ┌───┐  ┌───┐              │
│       │ 4 │  │ 2 │  │ 7 │  │   │  │   │  │   │              │   ← OTP input boxes
│       └───┘  └───┘  └───┘  └───┘  └───┘  └───┘              │
│                                                             │
│                                                             │
│                  Resend code (0:45)                         │   ← Tertiary (disabled)
│                                                             │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                    Verify                           │   │   ← Primary CTA
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│                                                             │
│              Use Email Login →                              │   ← Tertiary link
│                                                             │
│                                                             │
├─────────────────────────────────────────────────────────────┤
└─────────────────────────────────────────────────────────────┘
```

### Layout Constraint Check

| Validation | Result |
|------------|--------|
| **Layout Type** | Vertical Stack (per DSD v1.2 Section 1.2) |
| **DSD Alignment** | Matches onboarding screen philosophy (DSD Section 9) |
| **Single Focus** | ✓ One primary CTA per screen |
| **No Modal** | ✓ Full-screen form, not modal overlay |
| **Input Focused** | ✓ Minimal inputs, tap-based authentication options |

---

## 3. Layout & Spacing Specs

### 3.1 Container Structure

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  AUTH SCREEN CONTAINER                                                        │
│                                                                               │
│  ├── Position: fixed                                                          │
│  ├── Width: 100vw                                                             │
│  ├── Height: 100vh                                                            │
│  ├── Display: flex                                                            │
│  ├── Flex-direction: column                                                   │
│  ├── Background: var(--surface-background) → #FAFAF8                          │
│  │                                                                            │
│  ├── [HEADER AREA]                                                            │
│  │   ├── Height: 56px                                                         │
│  │   ├── Padding: var(--space-4) → 16px (horizontal)                          │
│  │   ├── Contains: Back button (left-aligned)                                 │
│  │   └── Safe area: env(safe-area-inset-top)                                  │
│  │                                                                            │
│  ├── [CONTENT AREA]                                                           │
│  │   ├── Flex: 1 (scrollable if needed)                                       │
│  │   ├── Padding: var(--padding-screen) → 20px (horizontal)                   │
│  │   ├── Padding-top: var(--space-6) → 24px                                   │
│  │   ├── Display: flex                                                        │
│  │   ├── Flex-direction: column                                               │
│  │   └── Gap: var(--space-6) → 24px (between sections)                        │
│  │                                                                            │
│  ├── [ACTION AREA]                                                            │
│  │   ├── Padding: var(--padding-screen) → 20px (all sides)                    │
│  │   ├── Padding-bottom: var(--space-8) → 40px                                │
│  │   └── Gap: var(--gap-stack) → 12px                                         │
│  │                                                                            │
│  └── [SAFE AREA BOTTOM]                                                       │
│      └── Height: env(safe-area-inset-bottom)                                  │
│                                                                               │
│  Premium Dark Mode (Default):                                                 │
│  ├── Background: var(--pdm-background) → #0D0D0F                              │
│  ├── Input surfaces: var(--pdm-surface-2) → #1A1A1E with subtle border glow   │
│  └── Focus states: Gold ring with ambient glow                                │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 3.4 Premium Dark Visual Treatment

| Element | Treatment |
|---------|-----------|
| **Background** | Deep charcoal `#0D0D0F` |
| **Input fields** | Surface `#1A1A1E`, border `#2A2A2E`, inner glow on focus |
| **Focus ring** | 2px gold `#C4A77D` with glow `0 0 12px rgba(196,167,125,0.3)` |
| **CTA Button** | Gold gradient with outer glow (DSD Section 12.5) |
| **OTP boxes** | Elevated surface with gold border on fill, pulse on current |

**Input Focus Animation (Premium):**
```css
/* Premium focus with glow ring */
.input:focus {
  border: 2px solid var(--accent-gold);
  box-shadow: 0 0 12px rgba(196, 167, 125, 0.25);
  transition: all 150ms ease-out;
}
```

**OTP Box States:**
```css
/* Current OTP box pulse */
@keyframes otp-pulse {
  0%, 100% { box-shadow: 0 0 0 rgba(196, 167, 125, 0); }
  50%      { box-shadow: 0 0 8px rgba(196, 167, 125, 0.4); }
}
```

### 3.2 Spacing Tokens

| Element | Token | Value |
|---------|-------|-------|
| Screen horizontal padding | `var(--padding-screen)` | 20px |
| Header height | — | 56px |
| Logo size | — | 48px height |
| Logo margin-bottom | `var(--space-6)` | 24px |
| Headline margin-bottom | `var(--space-2)` | 8px |
| Input group gap | `var(--space-5)` | 20px |
| Input label margin-bottom | `var(--space-2)` | 8px |
| Forgot password margin-top | `var(--space-3)` | 12px |
| CTA margin-top | `var(--space-6)` | 24px |
| Divider margin | `var(--space-6)` | 24px (top & bottom) |
| Footer link margin-top | `var(--space-8)` | 40px |

### 3.3 Z-Index Layers

| Layer Name | Z-Index | Contents |
|------------|---------|----------|
| **Background Layer** | 0 | Solid background color |
| **Content Layer** | 1 | Logo, form inputs, buttons |
| **Overlay Layer** | 10 | Error toasts, loading overlays |
| **Keyboard Layer** | 50 | Virtual keyboard (system) |
| **System Layer** | 100+ | iOS/Android status bar |

---

## 4. Component Inventory

### 4.1 Header / Navigation

**Component Name:** Back Navigation Header

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  BACK NAVIGATION HEADER                                                       │
│                                                                               │
│  ┌────────┐                                                                   │
│  │   ←    │                                                                   │
│  └────────┘                                                                   │
│                                                                               │
│  Specifications:                                                              │
│  ├── Touch target: 44px × 44px                                                │
│  ├── Icon: ArrowLeft (Phosphor), 24px                                         │
│  ├── Icon color: var(--unora-ink-primary) → #1A1A1A                           │
│  ├── Position: Left-aligned, 16px from edge                                   │
│  └── Tap: Navigate back to Welcome Landing                                    │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 4.2 Logo Component

| Property | Value |
|----------|-------|
| **Asset Type** | SVG (preferred) or PNG @3x |
| **Size** | 48px height |
| **Position** | Centered horizontally |
| **Color** | `var(--unora-primary-accent)` → #C4A77D |
| **Dark Mode** | Light variant (#F5F5F3) |

### 4.3 Typography Specifications

| Element | Font Family | Weight | Size | Line Height | Color |
|---------|-------------|--------|------|-------------|-------|
| **Headline (H2)** | `var(--font-display)` → Outfit | 600 | 22px | 1.25 | `var(--unora-ink-primary)` → #1A1A1A |
| **Body Text** | `var(--font-body)` → Inter | 400 | 16px | 1.5 | `var(--unora-ink-secondary)` → #4A4A4A |
| **Input Label** | `var(--font-body)` → Inter | 500 | 14px | 1.4 | `var(--unora-ink-secondary)` → #4A4A4A |
| **Input Text** | `var(--font-body)` → Inter | 400 | 16px | 1.5 | `var(--unora-ink-primary)` → #1A1A1A |
| **Input Placeholder** | `var(--font-body)` → Inter | 400 | 16px | 1.5 | `var(--unora-ink-muted)` → #A3A3A3 |
| **Error Text** | `var(--font-body)` → Inter | 500 | 12px | 1.4 | `var(--feedback-error)` → #C94A4A |
| **Caption** | `var(--font-body)` → Inter | 400 | 14px | 1.5 | `var(--unora-ink-tertiary)` → #7A7A7A |
| **Tertiary Link** | `var(--font-body)` → Inter | 500 | 14px | 1.5 | `var(--unora-primary-accent)` → #C4A77D |

### 4.4 Input Components

#### Phone Input

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  PHONE INPUT — DEFAULT                                                        │
│                                                                               │
│  Phone number                                                       ← Label   │
│  ┌─────────────────────────────────────────────────────────────┐              │
│  │  +91  |  98765 12345                                        │              │
│  └─────────────────────────────────────────────────────────────┘              │
│                                                                               │
│  Specifications:                                                              │
│  ├── Height: 52px                                                             │
│  ├── Border: 1.5px solid var(--border-medium) → #D4D4D0                       │
│  ├── Border radius: var(--radius-md) → 12px                                   │
│  ├── Background: var(--surface-card) → #FFFFFF                                │
│  ├── Padding: 16px horizontal                                                 │
│  ├── Font: Inter 16px / 400                                                   │
│  ├── Keyboard: Numeric (.numberPad on iOS, inputType="phone" on Android)     │
│  ├── Prefix: "+91" (non-editable, color: var(--unora-ink-tertiary))          │
│  ├── Separator: "|" divider after prefix (1px, var(--border-subtle))         │
│  ├── Auto-formatting: Adds space after 5th digit (XXXXX XXXXX)                │
│  ├── Max length: 10 digits (India standard)                                   │
│  └── Validation: Real-time on blur (must be exactly 10 digits)                │
└───────────────────────────────────────────────────────────────────────────────┘
```

**Input States:**

| State | Visual Changes |
|-------|----------------|
| **Default** | Border: `var(--border-medium)`, prefix visible |
| **Focused** | Border: 2px solid `var(--unora-primary-accent)` |
| **Filled** | Border: `var(--border-medium)`, number visible with formatting |
| **Error** | Border: 1.5px solid `var(--feedback-error)`, error text below |
| **Disabled** | Opacity: 0.5, non-interactive |

#### Password Input

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  PASSWORD INPUT                                                               │
│                                                                               │
│  Password                                                           ← Label   │
│  ┌─────────────────────────────────────────────────────────────┐              │
│  │  ••••••••                                         👁        │              │
│  └─────────────────────────────────────────────────────────────┘              │
│                                                                               │
│  Show/Hide Toggle:                                                            │
│  ├── Icon: Eye (hidden) / EyeSlash (visible)                                  │
│  ├── Size: 20px                                                               │
│  ├── Touch target: 44px × 44px                                                │
│  ├── Color: var(--unora-ink-tertiary)                                         │
│  └── Position: Right-aligned, 12px from edge                                  │
└───────────────────────────────────────────────────────────────────────────────┘
```

#### OTP Input

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  OTP INPUT (6 digits)                                                         │
│                                                                               │
│      ┌───┐  ┌───┐  ┌───┐  ┌───┐  ┌───┐  ┌───┐                                 │
│      │ 4 │  │ 2 │  │ 7 │  │   │  │   │  │   │                                 │
│      └───┘  └───┘  └───┘  └───┘  └───┘  └───┘                                 │
│                                                                               │
│  Specifications:                                                              │
│  ├── Box size: 48px × 56px                                                    │
│  ├── Gap: var(--space-2) → 8px                                                │
│  ├── Border: 1.5px solid var(--border-medium)                                 │
│  ├── Border radius: var(--radius-sm) → 8px                                    │
│  ├── Font: Outfit 24px / 600, centered                                        │
│  ├── Filled box: Border var(--unora-primary-accent)                           │
│  ├── Current box: Border 2px var(--unora-primary-accent), subtle pulse        │
│  ├── Auto-advance: Focus moves on digit entry                                 │
│  └── Auto-submit: Triggers verify on 6th digit                                │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 4.5 Button Components

#### Primary Button ("Get OTP" / "Verify")

| Property | Value | Reference |
|----------|-------|-----------|
| **Height** | 52px | DSD Section 3.1 |
| **Width** | Full width - 40px | — |
| **Border Radius** | `var(--radius-md)` → 12px | DSD Section 2.4 |
| **Background** | `var(--unora-primary-accent)` → #C4A77D | — |
| **Text Color** | #FFFFFF | — |
| **Font** | Inter 16px / 600 | DSD Section 3.1 |
| **Shadow** | 0 2px 8px rgba(0,0,0,0.08) | DSD Section 3.1 |

**Button States:**

| State | Appearance |
|-------|------------|
| **Default** | Full opacity |
| **Pressed** | Scale 0.98, opacity 0.9 |
| **Disabled** | Opacity 0.4, non-interactive |
| **Loading** | Spinner replaces text, spinner 16px white |

#### Tertiary Link ("Use Password Login" / "Get Started")

| Property | Value |
|----------|-------|
| **Height** | 44px (touch target) |
| **Background** | Transparent |
| **Text Color** | `var(--unora-primary-accent)` → #C4A77D |
| **Font** | Inter 14px / 500 |
| **Alignment** | Centered or right-aligned per context |

> **Note:** "Use Password Login" is an optional legacy fallback. Password Input component is not included in primary component inventory but may be implemented for fallback flow if needed.

### 4.6 Color Tokens Summary

| Token | Light Mode | Dark Mode | Usage |
|-------|------------|-----------|-------|
| `--surface-background` | #FAFAF8 | #121212 | Screen background |
| `--surface-card` | #FFFFFF | #1E1E1E | Input backgrounds |
| `--unora-primary-accent` | #C4A77D | #C4A77D | Primary CTA, focused inputs |
| `--unora-ink-primary` | #1A1A1A | #F5F5F3 | Headlines, input text |
| `--unora-ink-secondary` | #4A4A4A | #C4C4C0 | Labels, body text |
| `--unora-ink-tertiary` | #7A7A7A | #8A8A86 | Captions, icons |
| `--unora-ink-muted` | #A3A3A3 | #5A5A58 | Placeholders |
| `--border-medium` | #D4D4D0 | #3D3D3D | Input borders |
| `--feedback-error` | #C94A4A | #C94A4A | Error states |

---

## 5. Interaction & Logic Specification

### 5.1 Triggers

| Trigger | Element | Action |
|---------|---------|--------|
| **Tap** | Back arrow | Navigate to Welcome Landing |
| **Tap** | Phone number input | Focus input, show numeric keyboard |
| **Tap** | "Get OTP" button | Send OTP to entered phone number |
| **Tap** | "Use Password Login" | Navigate to Email/Password login screen (legacy fallback) |
| **Tap** | "Get Started" | Navigate to Profile Creation (new user) |
| **Tap** | OTP digit box | Focus that position |
| **Tap** | "Verify" button | Submit OTP for verification |
| **Tap** | "Resend code" | Trigger OTP resend (when enabled after 60s) |
| **Tap** | "Use Email Login" | Navigate to email/password authentication (fallback) |

### 5.2 Behaviors

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  PHONE + OTP AUTHENTICATION BEHAVIOR FLOW                                     │
│                                                                               │
│  USER ARRIVES FROM WELCOME LANDING ("Sign in →")                              │
│        │                                                                      │
│        ▼                                                                      │
│  ┌─────────────────────────────────────────────────────────────────────────┐  │
│  │  PHONE LOGIN SCREEN DISPLAYED                                           │  │
│  │                                                                         │  │
│  │  Initial state:                                                         │  │
│  │  ├── Phone input: Empty, not focused, +91 prefix visible               │  │
│  │  ├── Get OTP button: Disabled (no input)                                │  │
│  │  ├── "Use Password Login" link: Visible at bottom                       │  │
│  │  └── All elements visible and accessible                                │  │
│  └─────────────────────────────────────────────────────────────────────────┘  │
│        │                                                                      │
│        ├── [User taps phone input]                                            │
│        │         │                                                            │
│        │         ▼                                                            │
│        │   Numeric keyboard appears (.numberPad on iOS)                       │
│        │   Focus on input field                                               │
│        │         │                                                            │
│        ├── [User enters phone number]                                         │
│        │         │                                                            │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  REAL-TIME VALIDATION & FORMATTING                           │  │
│        │   │                                                               │  │
│        │   │  Phone format check:                                          │  │
│        │   │  ├── Auto-format: +91 XXXXX XXXXX (space after 5th digit)    │  │
│        │   │  ├── Validate: Exactly 10 digits required                     │  │
│        │   │  ├── During typing: No error shown                            │  │
│        │   │  └── On blur: Show error if invalid                           │  │
│        │   │                                                               │  │
│        │   │  Button state logic:                                          │  │
│        │   │  ├── < 10 digits: Button DISABLED                             │  │
│        │   │  └── = 10 digits: Button ENABLED                              │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │         │                                                            │
│        ├── [User taps "Get OTP"]                                              │
│        │         │                                                            │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  OTP GENERATION FLOW                                         │  │
│        │   │                                                               │  │
│        │   │  1. Button enters LOADING state                               │  │
│        │   │     └── Spinner appears, text: "Sending OTP..."               │  │
│        │   │  2. Phone input becomes non-interactive                       │  │
│        │   │  3. API call: POST /auth/send-otp                             │  │
│        │   │     Body: { "phone": "+919876512345" }                        │  │
│        │   │  4. Haptic: Success feedback                                  │  │
│        │   │  5. Navigate to OTP Verification Screen                       │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │         │                                                            │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  OTP VERIFICATION SCREEN                                     │  │
│        │   │                                                               │  │
│        │   │  Display:                                                     │  │
│        │   │  ├── "We sent a 6-digit code to"                              │  │
│        │   │  ├── "+91 98765 ••••5" (mask middle 5 digits)                 │  │
│        │   │  ├── 6 OTP input boxes                                        │  │
│        │   │  └── "Resend code (0:60)" countdown                           │  │
│        │   │                                                               │  │
│        │   │  Behavior:                                                    │  │
│        │   │  ├── Auto-focus first box                                     │  │
│        │   │  ├── Auto-advance on digit entry                              │  │
│        │   │  ├── Auto-submit on 6th digit                                 │  │
│        │   │  └── Resend enabled after 60 seconds                          │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │         │                                                            │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  [AUTO-SUBMIT] Verify button triggers                        │  │
│        │   │                                                               │  │
│        │   │  1. Button enters LOADING state                               │  │
│        │   │  2. OTP boxes become non-interactive                          │  │
│        │   │  3. API call: POST /auth/verify-otp                           │  │
│        │   │     Body: { "phone": "+919876512345", "otp": "123456" }      │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │         │                                                            │
│        │         ▼                                                            │
│        │   ┌───────────────────────────────────────────────────────────────┐  │
│        │   │  [DECISION] OTP verification result?                         │  │
│        │   └───────────────────────────────────────────────────────────────┘  │
│        │         │                           │                                │
│        │         ▼                           ▼                                │
│        │   ┌─────────────────┐      ┌─────────────────────────────────────┐   │
│        │   │    ✓ SUCCESS    │      │            ✗ FAILURE               │   │
│        │   │                 │      │                                     │   │
│        │   │ SMART ROUTING   │      │ Display inline error:               │   │
│        │   │ (See 5.2.1)     │      │ "Incorrect code. Try again."        │   │
│        │   │                 │      │                                     │   │
│        │   │ Route based on: │      │ Clear OTP boxes                     │   │
│        │   │ ├── New user?   │      │ Refocus first box                   │   │
│        │   │ └── Returning?  │      │ Haptic: error feedback              │   │
│        │   │                 │      │ Enable "Resend code" immediately    │   │
│        │   └─────────────────┘      └─────────────────────────────────────┘   │
│        │                                                                      │
│        └── [User taps "Use Password Login"]                                   │
│                  │                                                            │
│                  └──→ Navigate to Email/Password screen (legacy fallback)      │
│                                                                               │
└───────────────────────────────────────────────────────────────────────────────┘
```

**Reference:** 
- [Unora_UserFlow_Logic.md — Section 2.A (Onboarding)](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md)

### 5.2.1 Smart Routing Logic

```
┌───────────────────────────────────────────────────────────────────────────────┐
│  POST-OTP VERIFICATION SMART ROUTING                                          │
│                                                                               │
│  OTP VERIFIED SUCCESSFULLY ✓                                                  │
│        │                                                                      │
│        ▼                                                                      │
│  ┌────────────────────────────────────────────────────────────────────────┐   │
│  │  [DECISION] Is this a new user?                                        │   │
│  │  (Check: Phone number exists in database?)                             │   │
│  └────────────────────────────────────────────────────────────────────────┘   │
│        │                           │                                          │
│        ▼                           ▼                                          │
│  ┌──────────────────────┐   ┌─────────────────────────────────────────────┐   │
│  │    NEW USER          │   │        RETURNING USER                       │   │
│  │                      │   │                                             │   │
│  │  Account created     │   │  [DECISION] Is profile complete?            │   │
│  │  with phone number   │   │                                             │   │
│  │                      │   │  Query user_profiles table for:             │   │
│  │  Route to:           │   │  ├── identity_verified (Boolean)            │   │
│  │  ┌────────────────┐  │   │  ├── profile_complete (Boolean)             │   │
│  │  │ Identity       │  │   │  └── last_incomplete_step (String)          │   │
│  │  │ Verification   │  │   │                                             │   │
│  │  │ (Hard Gate)    │  │   │        │                    │                │   │
│  │  └────────────────┘  │   │        ▼                    ▼                │   │
│  │        ↓             │   │  ┌──────────────┐    ┌──────────────────┐   │   │
│  │  DigiLocker/         │   │  │   COMPLETE   │    │   INCOMPLETE     │   │   │
│  │  Aadhaar             │   │  │              │    │                  │   │   │
│  │        ↓             │   │  │ [DECISION]   │    │ Route to:        │   │   │
│  │  Auto-fill Profile   │   │  │ Verified?    │    │ Resume Profile   │   │   │
│  │        ↓             │   │  │    │     │   │    │ Creation         │   │   │
│  │  Profile Enrichment  │   │  │    ▼     ▼   │    │                  │   │   │
│  │                      │   │  │  ┌───┐ ┌───┐ │    │ (Step where      │   │   │
│  └──────────────────────┘   │  │  │YES│ │NO │ │    │  user left off)  │   │   │
│                             │  │  │   │ │   │ │    │                  │   │   │
│                             │  │  │ Go│ │ Go│ │    └──────────────────┘   │   │
│                             │  │  │ to│ │ to│ │                           │   │
│                             │  │  │Dis│ │Ver│ │                           │   │
│                             │  │  │cov│ │ifi│ │                           │   │
│                             │  │  │ery│ │cat│ │                           │   │
│                             │  │  │   │ │ion│ │                           │   │
│                             │  │  └───┘ └───┘ │                           │   │
│                             │  └──────────────┘                           │   │
│                             └─────────────────────────────────────────────┘   │
│                                                                               │
│  **PROGRESSIVE VERIFICATION PRINCIPLE:**                                      │
│  User accesses Discovery after Photos + Profile. Trust grows over time.       │
│                                                                               │
│  **Reference:** PRD Section 10 (Progressive Verification)                     │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 5.3 Post-Authentication Routing

> **Note:** Post-OTP verification routing logic is fully defined in **Section 5.2.1 (Smart Routing Logic)** above.

**Summary:**
- **New users** → Photo & Trust Setup → Profile Creation → Server Selection → Discovery
- **Returning users (complete profile)** → Discovery/Home (if verified) or Verification (if not)
- **Returning users (incomplete profile)** →  Resume Profile Creation (at last incomplete step)

**Progressive Model:** Users access Discovery after completing Photo & Trust Setup and Profile Creation. Trust verification continues in background.

### 5.4 Transitions

#### Screen Entry (from Welcome)

| Property | Value | Reference |
|----------|-------|-----------|
| **Animation Type** | Slide from right | DSD Section 5.4 |
| **Duration** | 300ms | DSD Section 5.4 |
| **Easing** | ease-out | — |

#### Screen Exit (to Discovery)

| Property | Value |
|----------|-------|
| **Animation Type** | Cross-fade |
| **Duration** | 200ms |
| **Easing** | ease-out |

#### Input Focus Animation

```
Input Focus Transition:
[0ms]      Input tap detected
[0ms]      Border color animates: medium → primary-accent
[0ms]      Border width: 1.5px → 2px
[150ms]    Transition complete
[+]        Keyboard slides up (system)
```

#### Error Animation

```
Error Shake Animation:
[0ms]      Error detected
[0ms]      Haptic: error
[0-150ms]  Input shakes: translateX(0 → 4px → -4px → 2px → 0)
[150ms]    Error text fades in below input
```

---

## 6. State Definitions

### 6.1 State Matrix

| State | Visual Appearance | Conditions | Duration |
|-------|-------------------|------------|----------|
| **Default** | Empty phone input, Get OTP disabled | Fresh screen load | Until input |
| **Input Active** | Focused phone input, numeric keyboard visible | User typing | Until blur/submit |
| **Form Ready** | Get OTP enabled | Valid 10-digit phone number entered | Until submit |
| **Loading** | Spinner on button, phone input disabled | OTP generation in progress | Until response |
| **OTP Screen** | 6 OTP boxes, countdown timer | Successful OTP send | Until verification |
| **Error** | Error message, affected input highlighted | Invalid phone or OTP failure | Until retry |
| **Success** | Brief confirmation, navigate away | OTP verification success | ~500ms |

### 6.2 Default State

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ←                                                         │
│                                                             │
│                   [ UNORA LOGO ]                            │
│                                                             │
│               Welcome back.                                 │
│                                                             │
│                                                             │
│   Phone number                                              │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  +91  |                                             │   │   ← Placeholder
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                     Get OTP                         │   │   ← DISABLED
│   └─────────────────────────────────────────────────────┘   │       (opacity 0.4)
│                                                             │
│               Don't have an account?                        │
│                    Get Started →                            │
│                                                             │
│                 Use Password Login →                        │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Button disabled until valid 10-digit phone number entered.
Border colors: var(--border-medium)
```

### 6.3 Loading State (OTP Generation)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   (Back button disabled during load)                        │
│                                                             │
│                   [ UNORA LOGO ]                            │
│                                                             │
│               Welcome back.                                 │
│                                                             │
│                                                             │
│   Phone number                                              │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  +91  |  98765 12345                                │   │   ← Non-interactive
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                 ◐  Sending OTP...                   │   │   ← Loading
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Spinner: 16px, white
Text: Faded to 80% opacity
Phone input: pointer-events: none
```

### 6.4 Error State (Invalid Phone)

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   ←                                                         │
│                                                             │
│                   [ UNORA LOGO ]                            │
│                                                             │
│               Welcome back.                                 │
│                                                             │
│                                                             │
│   Phone number                                              │
│   ┌─────────────────────────────────────────────────────┐   │
│   │  +91  |  98765 123                                  │   │   ← Border: error
│   └─────────────────────────────────────────────────────┘   │
│   ⚠️  Invalid phone number. Must be 10 digits.             │   ← Error message
│                                                             │
│   ┌─────────────────────────────────────────────────────┐   │
│   │                     Get OTP                         │   │   ← Disabled
│   └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Error styling:
├── Input border: var(--feedback-error) → #C94A4A
├── Error text: var(--feedback-error), 12px, 500 weight
├── Error icon: Warning, inline
└── Animation: Subtle shake (150ms)
```

### 6.5 OTP Error State

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│               Enter verification code                       │
│                                                             │
│   We sent a 6-digit code to                                 │
│   +91 98765•••43                                            │
│                                                             │
│                                                             │
│       ┌───┐  ┌───┐  ┌───┐  ┌───┐  ┌───┐  ┌───┐              │
│       │ 4 │  │ 2 │  │ 7 │  │ 8 │  │ 9 │  │ 1 │              │   ← All red border
│       └───┘  └───┘  └───┘  └───┘  └───┘  └───┘              │
│                                                             │
│              ⚠️ Incorrect code. Try again.                   │   ← Error message
│                                                             │
│                  Resend code →                              │   ← Enabled after error
│                                                             │
└─────────────────────────────────────────────────────────────┘

OTP boxes: All borders turn var(--feedback-error)
Shake animation: 150ms
Boxes: Clear after shake, refocus first box
```

---

## 7. Content & Copy Guidelines

### 7.1 Text Strings — Phone Login Screen

| Element | Copy | Notes |
|---------|------|-------|
| **Headline** | "Welcome back." | Warm, personal |
| **Phone Label** | "Phone number" | Simple, clear |
| **Phone Prefix** | "+91" | Non-editable, India default |
| **Primary CTA** | "Get OTP" | — |
| **CTA Loading** | "Sending OTP..." | — |
| **Footer Caption** | "Don't have an account?" | — |
| **Footer Link** | "Get Started →" | — |
| **Fallback Link** | "Use Password Login →" | Legacy option, bottom |

### 7.2 Text Strings — OTP Screen

| Element | Copy | Notes |
|---------|------|-------|
| **Headline** | "Enter verification code" | — |
| **Body** | "We sent a 6-digit code to" | — |
| **Phone Display** | "+91 98765 ••••5" | Masked (show first 5, last 1 digit) |
| **Resend (disabled)** | "Resend code (0:45)" | Countdown from 60s |
| **Resend (enabled)** | "Resend code →" | After 60s |
| **Primary CTA** | "Verify" | — |
| **CTA Loading** | "Verifying..." | — |
| **Alt Method** | "Use Email Login →" | Fallback option |

### 7.3 Error Messages

| Scenario | Message | Tone |
|----------|---------|------|
| **Invalid phone format** | "Invalid phone number. Must be 10 digits." | Clear, instructive |
| **Phone not registered** | "This phone number is not registered. Get Started →" | Helpful redirect |
| **OTP send failure** | "Couldn't send OTP. Check your network and try again." | Action-oriented |
| **Invalid OTP** | "Incorrect code. Try again." | Simple |
| **OTP expired** | "This code has expired. Request a new one." | Action-oriented |
| **Too many attempts** | "Too many attempts. Try again in 5 minutes." | Firm but fair |
| **Network error** | "Couldn't connect. Check your internet." | — |

### 7.4 Tone Guidelines

| Principle | Application |
|-----------|-------------|
| **Presence over Performance** | No "fast login" or speed language |
| **Clarity over Clutter** | Minimal inputs, clear labels |
| **Firmness without Friction** | Errors are calm, solutions are suggested |
| **Anti-Tinder** | No gamification, no urgency |

**Reference:** [Unora_PRD.md — Section 7 (Core Philosophy)](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md)

### 7.5 Avoid

| ❌ Don't Use | ✓ Instead |
|-------------|-----------|
| "Login" | "Sign In" / "Get OTP" |
| "Enter mobile" | "Phone number" |
| "Wrong OTP!" | "Incorrect code. Try again." |
| "That didn't work" | "Incorrect code. Try again." |
| "Invalid number!!" | "Invalid phone number. Must be 10 digits." |
| "Hurry back to your matches" | "Welcome back." |

---

## 8. Accessibility Specifications

### 8.1 Screen Reader Support

```
Login Screen:
├── Announce: "Sign In to Unora. Welcome back."
├── Focus order: Back → Logo (skip) → Headline → Email input → Password input 
│               → Forgot password → Sign In → Phone OTP → Get Started
├── Input labels: Connected via aria-labelledby
├── Password toggle: aria-label="Show password" / "Hide password"
└── Error states: aria-live="polite" for error announcements

OTP Screen:
├── Announce: "Verification. Enter the 6-digit code sent to your phone."
├── OTP inputs: Group as aria-label="6-digit verification code"
├── Individual boxes: aria-label="Digit 1 of 6", etc.
└── Auto-read errors: aria-live="assertive"
```

### 8.2 Keyboard Navigation

| Key | Action |
|-----|--------|
| **Tab** | Move to next input/button |
| **Shift+Tab** | Move to previous input/button |
| **Enter** | Submit form (when on input) |
| **Space** | Activate button |
| **Backspace** | OTP: Move to previous box and clear |

### 8.3 Color Contrast

| Element | Foreground | Background | Ratio | WCAG |
|---------|------------|------------|-------|------|
| Headline | #1A1A1A | #FAFAF8 | 16:1 | ✓ AAA |
| Input text | #1A1A1A | #FFFFFF | 21:1 | ✓ AAA |
| Placeholder | #A3A3A3 | #FFFFFF | 2.6:1 | ✓ AA (placeholder exempt) |
| Error text | #C94A4A | #FAFAF8 | 5.2:1 | ✓ AA |
| CTA text | #FFFFFF | #C4A77D | 3.2:1 | ✓ AA (large) |

### 8.4 Touch Targets

| Component | Touch Size | Visual Size |
|-----------|------------|-------------|
| Back button | 44px × 44px | 24px icon |
| Text inputs | Full width × 52px | — |
| Password toggle | 44px × 44px | 20px icon |
| Primary CTA | Full width × 52px | — |
| Tertiary links | 44px height | Text size |
| OTP boxes | 48px × 56px | — |

---

## 9. Platform-Specific Notes

### 9.1 iOS

- **Keyboard**: `keyboardType: .numberPad` for phone input
- **Phone**: `textContentType: .telephoneNumber` for autofill
- **OTP**: `textContentType: .oneTimeCode` for SMS autofill (critical for UX)
- **Face ID/Touch ID**: Consider biometric login for returning users (Phase 2)
- **Back gesture**: Disable edge swipe during OTP verification

### 9.2 Android

- **Autofill**: Use autofill hints for phone number
- **Keyboard**: `inputType="phone"` for phone input
- **OTP autofill**: SMS Retriever API for auto-read (critical for UX)
- **Back button**: Navigate to Welcome Landing

### 9.3 Web (PWA)

- **Autocomplete**: `autocomplete="tel"` for phone number input
- **Input mode**: `inputmode="numeric"` for phone input
- **Focus management**: Auto-focus phone number field on load
- **OTP**: Consider Web OTP API for autofill (experimental)

---

## 10. Security Considerations

### 10.1 Security Requirements

| Requirement | Implementation |
|-------------|----------------|
| **Password not stored** | Client never stores password in plaintext |
| **HTTPS only** | All auth requests over TLS 1.3 |
| **Rate limiting** | Max 5 attempts per 15 minutes |
| **OTP expiry** | 5-minute validity |
| **Session management** | Secure, httpOnly cookies |
| **Brute force protection** | Progressive delays on failed attempts |

### 10.2 Password Visibility

- Password is masked by default (•••••••)
- Toggle reveals actual characters
- Auto-hide after 3 seconds of visibility (optional)
- Always revert to hidden on screen blur

---

## 11. Implementation Checklist

| Requirement | Priority | Status |
|-------------|----------|--------|
| Phone input with +91 prefix and formatting | Critical | ☐ |
| Get OTP button with loading state | Critical | ☐ |
| Error state display (inline) | Critical | ☐ |
| Phone number validation (10 digits) | Critical | ☐ |
| OTP input (6 boxes) with auto-advance | High | ☐ |
| OTP auto-submit on 6th digit | High | ☐ |
| OTP SMS autofill (iOS/Android) | High | ☐ |
| Resend OTP timer (60s countdown) | High | ☐ |
| Smart routing logic (new vs returning users) | Critical | ☐ |
| Post-OTP routing to verification/discovery | Critical | ☐ |
| Back navigation | High | ☐ |
| Use Password Login fallback link | Medium | ☐ |
| Dark mode support | Medium | ☐ |
| Keyboard handling (numeric) | High | ☐ |
| Accessibility (screen reader) | Medium | ☐ |
| Touch targets (44px minimum) | Medium | ☐ |
| Rate limiting UI feedback | Medium | ☐ |

---

## 12. Related Documents

| Document | Relevance |
|----------|-----------|
| [Unora_PRD.md](file:///c:/Unora/Founder_Product_docs/Unora_PRD.md) | Core philosophy, verification requirements |
| [Unora_DesignSystem.md](file:///c:/Unora/Founder_Product_docs/Unora_DesignSystem.md) | Input specs, error states (Section 10), tokens |
| [Unora_UserFlow_Logic.md](file:///c:/Unora/Founder_Product_docs/Unora_UserFlow_Logic.md) | Onboarding flow, verification gates |
| [Unora_Spec_02_WelcomeCarousel.md](file:///c:/Unora/Founder_Product_docs/UI%20Screens/Unora_Spec_02_WelcomeCarousel.md) | Previous screen (welcome) |
| Unora_Spec_04_PasswordReset.md (planned) | Related flow |
| Unora_Spec_05_ProfileCreation.md (planned) | Next step for new users |

---

**Document maintained by:** Product Design Team  
**Last updated:** January 2026  
**Next review:** Upon design system updates

---

*This specification is developer-ready and should be implemented as defined. Any deviations require design review.*
