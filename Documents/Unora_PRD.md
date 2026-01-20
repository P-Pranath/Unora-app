# Unora — Product Requirement Document

**Version:** 1.0  
**Date:** January 2026  
**Classification:** Internal / Founding Team / Early Investors  
**Author:** Product Leadership

---

## Table of Contents

1. [Executive Summary](#1-executive-summary)
2. [What Unora Is](#2-what-unora-is)
3. [What Unora Is Not](#3-what-unora-is-not)
4. [Category Positioning](#4-category-positioning)
5. [Core Problem Being Solved](#5-core-problem-being-solved)
6. [Why Existing Solutions Fail Structurally](#6-why-existing-solutions-fail-structurally)
7. [Core Philosophy (Locked)](#7-core-philosophy-locked)
8. [Target User](#8-target-user)
9. [System Architecture Overview](#9-system-architecture-overview)
10. [Onboarding & Identity Verification](#10-onboarding--identity-verification)
11. [Server Selection (Intent Segmentation)](#11-server-selection-intent-segmentation)
12. [Discovery & Matching](#12-discovery--matching)
13. [Streak System](#13-streak-system)
14. [Nudge System](#14-nudge-system)
15. [Reveal System](#15-reveal-system)
16. [Monetization](#16-monetization)
17. [Credits & Payment Protection](#17-credits--payment-protection)
18. [AI Role](#18-ai-role)
19. [Trust, Safety & Ethics](#19-trust-safety--ethics)
20. [Metrics & Success Signals](#20-metrics--success-signals)
21. [Edge Cases & Failure Modes](#21-edge-cases--failure-modes)
22. [Roadmap](#22-roadmap)
23. [Technical Stack (Recommended)](#23-technical-stack-recommended)
24. [Appendix: Investor Q&A Reference](#24-appendix-investor-qa-reference)

---

## 1. Executive Summary

Unora is a commitment-first connection platform where consistency, safety, and anticipation precede communication. Unlike conventional dating applications optimized for rapid matching and instant gratification, Unora deliberately introduces friction as a filter for intent.

Users begin anonymous. They discover others through hobbies and consistency signals rather than photographs. When mutual interest occurs, they do not chat. They build a shared streak over 15 days, with real identity details revealed gradually as milestones are reached. Conversation unlocks only after something real has been built.

This document defines the complete product specification for Unora's minimum viable product and subsequent phases, intended for use by engineers, designers, AI/ML engineers, investors, and the founding team.

---

## 2. What Unora Is

Unora is a platform for building human connection through:

- **Earned communication** — Chat is unlocked only after 15 days of mutual consistency
- **Anonymous discovery** — Users browse real people via hobbies, not appearances
- **Mutual choice** — Connections require both parties to independently express interest
- **Streak-based progression** — Daily check-ins create shared investment
- **Gradual reveals** — Identity information is disclosed incrementally as trust builds
- **Intent-first design** — The system structurally filters for users who are willing to invest effort

The core user journey:

1. Profile creation with phone verification and mandatory photos
2. Progressive trust building through consistency and optional verification
3. Server selection based on intent (partner, friend, accountability buddy)
4. Anonymous discovery via hobby-anchored cards
5. Mutual interest expression
6. Streak initiation upon mutual match
7. Daily check-ins with recovery mechanisms
8. Gradual reveals unlocked by consistency
9. Chat unlocked after 15-day completion

---

## 3. What Unora Is Not

**Unora is not a dating app in the conventional sense.**

It does not optimize for:

- Speed of matching
- Volume of conversations
- Swipe-based discovery
- Photo-first attraction
- Instant messaging access
- Gamified attention capture

It does not offer:

- Infinite scroll
- Algorithmic feed optimization for engagement
- Public profiles
- Social graph connectivity
- Super likes, roses, or attention-buying mechanics
- Pay-to-message or pay-to-match features

**Unora is not for everyone.** That is intentional. Users who want instant gratification already have infinite alternatives. Unora serves a distinct audience that existing products structurally cannot reach.

---

## 4. Category Positioning

**Primary Category:** Lifestyle / Intentional Connection  
**Secondary Positioning:** Commitment-based social platform

Unora occupies a distinct category separate from:

| Category        | Optimization Target           | Unora Distinction                  |
| --------------- | ----------------------------- | ---------------------------------- |
| Dating apps     | Match volume, chat initiation | Earned communication, delayed chat |
| Social networks | Attention, engagement time    | Bounded interaction, no feed       |
| Habit apps      | Individual streaks            | Shared streaks with another person |
| Messaging apps  | Communication frequency       | Communication as earned outcome    |

Dating is an outcome users may experience on Unora — it is not the category. Many connections will be romantic; some will be friendships or accountability partnerships. The product does not position itself around any single outcome.

---

## 5. Core Problem Being Solved

**Modern connection platforms have optimized the same swipe-and-chat model for over a decade, but the underlying interaction design creates structural problems:**

### The Paradox of Infinite Choice

Unlimited options create paralysis, reduced commitment, and shallow evaluation. Users optimize for quantity of matches rather than quality of connection.

### The Attention Asymmetry Problem

Platforms that allow instant messaging create unequal dynamics where some users receive overwhelming attention while others receive none. This leads to degraded experience on both sides.

### The Effort-Trust Gap

When communication costs nothing, messages carry no signal of genuine interest. Users cannot distinguish between casual curiosity and meaningful intent.

### The Safety Burden

Identity-first platforms expose users (particularly women) to harassment, judgment, and unwanted contact before any trust has been established.

### The Stigma Barrier

In markets like India, many potential users have never used dating apps due to social stigma. Products positioned as "dating apps" exclude this audience by default.

### The Ghosting Equilibrium

Low switching costs and abundant alternatives make abandoning conversations the path of least resistance. Platforms have no mechanism to make continuation more attractive than exit.

---

## 6. Why Existing Solutions Fail Structurally

| Platform Approach         | Structural Flaw                                                                   |
| ------------------------- | --------------------------------------------------------------------------------- |
| Swipe-based matching      | Rewards superficial evaluation, creates match inflation, devalues each connection |
| Chat-first interaction    | No barrier to entry means no signal of intent                                     |
| Photo-first discovery     | Optimizes for appearance over compatibility                                       |
| Unlimited messaging       | Creates harassment vectors, particularly for women                                |
| Algorithmic engagement    | Incentivizes platform addiction over successful outcomes                          |
| Premium boosts/visibility | Turns attention into a purchasable commodity                                      |
| Infinite scroll           | Prevents focus, encourages comparison, reduces commitment                         |

**The fundamental insight:** These platforms have made it easy to match, but hard to trust, hard to stay consistent, and hard to build anything real. Faster matching did not fix ghosting, distraction, or burnout. It amplified them.

---

## 7. Core Philosophy (Locked)

These principles are non-negotiable and inform every product decision:

| Principle                                     | Implication                                                                 |
| --------------------------------------------- | --------------------------------------------------------------------------- |
| **Connection should not be instant**          | Chat is earned, not given                                                   |
| **Consistency > Speed**                       | Users who return daily receive more than users who binge                    |
| **Anticipation builds attachment**            | Delayed reveals create emotional investment                                 |
| **Effort builds trust**                       | Actions that require investment carry meaning                               |
| **Money buys flexibility, never outcomes**    | Users cannot pay for another person's attention                             |
| **Filters give agency, reveals give emotion** | Control and surprise serve different purposes                               |
| **Reset ≠ Rejection**                         | A broken streak is a pause, not an ending                                   |
| **AI supports depth, never replaces effort**  | AI frames, enhances, and protects — it does not substitute for human action |

---

## 8. Target User

### Primary Market

- **Geography:** India-first, initially Bangalore only
- **Age:** 22–35
- **Demographics:** Urban, English-first
- **Psychographics:** College students and working professionals; thoughtful users tired of swipe culture

### Key User Characteristics

1. **Connection-seeking but platform-skeptical** — Many have never used dating apps due to stigma or distaste for existing experiences
2. **Consistency-capable** — Willing to invest daily effort for delayed reward
3. **Privacy-valuing** — Prefer gradual disclosure over public exposure
4. **Intent-clear** — Know what they are looking for and want filtered experiences
5. **Quality over quantity** — Prefer fewer, more meaningful interactions to many shallow ones

### Users We Are Not Targeting

- Users seeking instant gratification
- Users who prioritize volume of matches
- Users uncomfortable with delayed communication
- Users unwilling to verify identity

### Market Expansion Path

1. **Phase 1:** Bangalore only (highest density of early adopters, comfort with subscriptions, strong creator and startup culture)
2. **Phase 2:** Other metro cities (Delhi NCR, Mumbai, Hyderabad, Chennai, Pune)
3. **Phase 3:** Tier-2 cities with sufficient density
4. **Phase 4:** Age-specific cohorts (18–22 college-specific, 35–45 separate cohort)

---

## 9. System Architecture Overview

Unora is composed of five tightly coupled layers. Each layer is intentionally opinionated. Removing friction from any one layer breaks the system.

### Layer 1: Identity & Trust Layer

Real-person verification without public identity exposure. Users provide complete information but most remains hidden until earned reveals.

### Layer 2: Discovery & Matching Layer

Constrained choice combined with mutual intent. Anonymous cards, hobby-based attraction, and required bilateral interest before any connection forms.

### Layer 3: Streak & Progression Engine

The core behavioral loop. Daily check-ins, at-risk states, recovery windows, and milestone-based progression.

### Layer 4: AI Mediation Layer

Pacing, safety, framing, and scoring. AI operates quietly in service of user experience without replacing human effort.

### Layer 5: Monetization & Control Layer

Ethical monetization via time, flexibility, and control. Money affects pace and forgiveness, never outcomes or access to other people.

---

## 10. Onboarding & Progressive Verification

Unora uses a **Progressive Verification Model** where trust is earned through presence and consistency, not paperwork. There is no hard gate at onboarding. Instead, personhood is verified through mandatory steps, and identity is optionally confirmed over time through behavioral and documentary signals.

### 10.1 Progressive Verification Overview

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  VERIFICATION LAYERS (Progressive Trust Model)                              │
│                                                                             │
│  MANDATORY (Personhood):                                                    │
│  ├── Phone number + OTP authentication                                      │
│  ├── 3–6 photos uploaded during onboarding                                  │
│  └── Device continuity tracking                                             │
│                                                                             │
│  OPTIONAL TRUST BOOSTERS (Free Integrations):                               │
│  ├── Google Sign-In                                                         │
│  ├── Apple Sign-In                                                          │
│  └── LinkedIn Connect                                                       │
│                                                                             │
│  PROGRESSIVE IN-APP VERIFICATION:                                           │
│  ├── AI-powered media quality evaluation (automatic)                        │
│  ├── Optional voice presence                                                │
│  ├── Behavioral consistency (15-day streak completion)                      │
│  └── Consistency patterns over time                                         │
│                                                                             │
│  OPTIONAL IDENTITY ESCALATION (Post-Trust):                                 │
│  └── Government ID verification (token-based, no raw ID storage)            │
│      → Unlocks flexibility features, never gates outcomes                   │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 10.2 Onboarding Flow Overview

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  ONBOARDING SEQUENCE (Progressive Model)                                    │
│                                                                             │
│  Step 1: Auth         → Phone number + OTP authentication                   │
│  Step 2: Photos       → 3–6 photos uploaded (mandatory, private by default) │
│  Step 3: Profile      → User provides Name, DOB, Gender, City, Hobbies      │
│  Step 4: Trust Boost  → Optional: Google/Apple/LinkedIn (one-tap)           │
│  Step 5: Server       → User selects intent server                          │
│  Step 6: Discovery    → User enters Discovery                               │
│                                                                             │
│  IN-APP (Async):                                                            │
│  ├── AI media quality filtering runs automatically                          │
│  ├── Behavioral trust earned through streak completion                      │
│  └── Optional Gov ID verification available post-trust                      │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 10.3 Mandatory Verification (Personhood)

Users must complete these steps to enter Discovery. This is non-negotiable.

#### Phone + OTP Authentication

- Standard SMS OTP flow
- Device binding via phone number
- Required for account creation

#### Mandatory Photo Upload

Users must upload 3–6 photos during onboarding. This serves multiple purposes:

| Purpose | Explanation |
|---------|-------------|
| **Personhood Signal** | Photos provide visual evidence of a real person |
| **Liveness Preparation** | Photos enable future liveness challenges |
| **Reveal Foundation** | Photos are the content of the Day 15 reveal |
| **Commitment Filter** | Effort required filters out low-intent users |

**Photo Requirements:**

- Minimum 3 photos, maximum 6
- At least one clear face photo required
- AI-assisted photo validation (blur detection, face presence)
- All photos private by default — revealed only at Day 15 milestone

**Privacy Messaging:**

> "Your photos stay hidden until Day 15. Only you can see them until then."

#### Device Continuity

- Device fingerprinting for account security
- Pattern detection for suspicious account behavior
- Session management across devices

### 10.4 Optional Trust Boosters

Free integrations that enhance trust signals without creating gates.

| Integration | Trust Signal | Implementation |
|-------------|--------------|----------------|
| **Google Sign-In** | Established Google account | OAuth 2.0 |
| **Apple Sign-In** | Apple ID verification | Sign in with Apple |
| **LinkedIn Connect** | Professional identity signal | OAuth 2.0 |

**Rules:**

- Trust boosters are **never gates** — users can skip without penalty
- Visible only to the user (not shown to potential matches)
- One-tap integration during onboarding or later in settings
- Multiple boosters can be connected

**Why these matter:**

Established accounts on major platforms provide lightweight identity signals without the friction of government ID verification.

### 10.5 Progressive In-App Verification

Trust builds through behavior over time. These verifications happen naturally during the user journey.

#### AI-Powered Media Quality Evaluation

AI continuously evaluates uploaded photos for quality and authenticity.

| Aspect | Detail |
|--------|--------|
| **When** | Automatically during photo upload |
| **What** | AI evaluates image quality, face presence, usability |
| **Outcome** | Usable images pass; low-quality images prompt re-upload |
| **UX** | Non-interruptive — runs in background |

**Why passive:** Users are never asked to perform verification actions. AI maintains quality silently.

#### Optional Voice Presence

Future enhancement for humanization signal.

- Short voice recording during check-in
- AI verification that recording is live (not pre-recorded)
- Never shared with matches — for verification only

#### Behavioral Trust (Streak Completion)

The strongest trust signal: completing a 15-day mutual streak.

| Behavior | Trust Implication |
|----------|-------------------|
| Consistent daily check-ins | Real person with routine |
| Recovery payments made | Financial stake in identity |
| Multiple streaks completed | Pattern of genuine engagement |

**System tracks:**

- Check-in timing consistency
- Recovery payment history
- Connection termination patterns
- Device continuity

### 10.6 Optional Identity Escalation

Government ID verification available after behavioral trust is established.

#### When Available

Government ID verification becomes available after:
- At least one streak completion, OR
- 30 days of active account history, OR
- User explicitly requests via settings

#### Implementation

| Aspect | Detail |
|--------|--------|
| **Methods** | Government-backed digital identity (token-based) |
| **Storage** | Verification tokens only — no raw ID retention |
| **Encryption** | All PII encrypted at rest (AES-256) |
| **Visibility** | Zero visibility of verification method to other users |

#### What Gov ID Unlocks

Government ID is a **flexibility feature**, not an access gate:

| Benefit | Explanation |
|---------|-------------|
| **Profile Name Lock** | Option to lock name as verified (cannot be edited) |
| **Travel Mode** | Extended location flexibility for verified travelers |
| **Enhanced Recovery** | Faster support escalation for verified accounts |
| **Future Features** | Reserved for premium features requiring verified identity |

> [!IMPORTANT]
> **Government ID never unlocks access to other users.** It unlocks personal flexibility features only.

### 10.7 Why Unora Verifies Through Consistency

> [!NOTE]
> This section explains the philosophy behind progressive verification for investors and team members.

**Trust is earned through presence, not paperwork.**

Traditional dating apps face two flawed options:
1. **No verification** — Catfishing, fake profiles, harassment
2. **Hard-gate verification** — High friction, privacy concerns, excludes users

Unora introduces a third path: **Progressive behavioral verification.**

| Verification Type | Signal Strength | User Friction | Privacy Impact |
|-------------------|-----------------|---------------|----------------|
| Gov ID at signup | High initial | High | High |
| Behavioral over time | High cumulative | Low | Low |

**Why behavioral signals work:**

- A user who maintains a 15-day streak is demonstrably real
- Consistency patterns are hard to fake at scale
- AI media filtering prevents unusable profile photos
- Bad actors are filtered by the effort required

**What we explicitly do NOT claim:**

- ❌ "All users are government-verified" (unless they complete optional escalation)
- ❌ "Names are legal names" (unless gov ID verified)

**What we CAN claim:**

- ✓ "All users are real people" (phone + photos + behavioral signals)
- ✓ "Trust is earned through consistency" (behavioral signals)
- ✓ "Photos meet quality standards" (AI-powered filtering)

### 10.8 Profile Data (User-Provided)

All profile fields are user-provided and editable (unless Gov ID verification is completed and user opts to lock fields).

**Profile fields:**

| Field | Required | Notes |
|-------|----------|-------|
| Name | Yes | User-provided, editable |
| Date of Birth | Yes | User-provided, editable |
| Gender | Yes | User-provided, editable |
| City | Yes | User-provided, editable |
| Education | Optional | User-provided, editable |
| Profession | Optional | User-provided, editable |
| Religion/Culture | Optional | User-provided, editable |
| Photos | Yes (3-6) | Private until Day 15 reveal |
| Hobbies | Yes (min 2) | With micro-descriptions |
| Bio/Intent | Optional | User-provided, editable |

**Privacy Principle:** Privacy through delayed exposure, not through missing data. Complete information enables accurate matching and meaningful reveals while maintaining anonymity until trust is established.

### 10.9 Edge Cases & Handling

| Scenario | Handling |
|----------|----------|
| **User skips optional trust boosters** | No penalty. Trust builds through behavior. |
| **Photos rejected (blur/no face)** | Must resubmit before continuing. Clear, helpful guidance. |
| **User wants Gov ID but hasn't earned it** | Shown unlock criteria. Encouraged to complete streak. |
| **Repeated photo upload failures** | Support escalation available. No permanent block. |

### 10.10 Verification States

Users have a verification status visible only to themselves (never to matches):

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  VERIFICATION STATUS (User's View Only)                                     │
│                                                                             │
│  ✓ Phone verified                                                           │
│  ✓ Photos submitted (3/6)                                                   │
│  ✓ Photo quality verified                                                   │
│  ⏳ Consistency building (Day 4 of first streak)                            │
│  ○ Government ID (optional)                                                 │
│                                                                             │
│  "Trust builds through consistency. You're always in control."              │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 11. Server Selection (Intent Segmentation)

After verification, users can browse any of the three intent servers:

| Server                            | Description                     |
| --------------------------------- | ------------------------------- |
| **Looking for a partner**         | Romantic connection intent      |
| **Friend / companion**            | Platonic connection intent      |
| **Accountability / growth buddy** | Goal-oriented connection intent |

### 11.1 Universal Server Switching

**All users (Free, Plus, Pro) can switch between servers at any time, instantly.** There is no cooldown, no tier restriction, and no switching fee.

| Rule | Implementation |
|------|---------------|
| **Switch Anytime** | Users tap a server tab and immediately enter that server's discovery view |
| **No Cooldown** | Switching is instant; no waiting period between server changes |
| **No Tier Gating** | Free, Plus, and Pro users all have identical switching capabilities |
| **Matches Never Cross Servers** | A connection made in the Friend server stays in the Friend server |
| **Intent Visibility** | Server selection is explicit and visible to potential matches |

### 11.2 The Global Inventory (Critical Constraint)

While server switching is unrestricted, **all resources are shared globally across servers**. This is the core constraint that enforces intentional resource budgeting.

#### Global Refresh Timer

The Refresh cooldown is **account-wide**, not per-server.

| Scenario | Behavior |
|----------|----------|
| User refreshes in Dating server | Refresh timer starts (e.g., 12 hours for Plus) |
| User switches to Friend server | Refresh button is **LOCKED** — same global timer applies |
| User switches to Growth server | Refresh button is **LOCKED** — same global timer applies |
| Timer expires | Refresh becomes available in **whichever server the user is currently viewing** |

**Implication:** A Plus user gets 1 refresh every 12 hours **for the entire account**, not per server. Users must choose which server to invest their refresh in.

#### Global Connection Slots

Connection slots (Free: 1, Plus: 2, Pro: 4) are a **single shared pool** across all servers.

| Scenario | Behavior |
|----------|----------|
| Free user sends interest in Friend server → match created | Slot 1/1 filled |
| Same user switches to Dating server | Can browse cached cards, but **Connect button is DISABLED** |
| Same user switches to Growth server | Can browse cached cards, but **Connect button is DISABLED** |

**Implication:** Users cannot maintain separate connection pools per server. A Free user with 1 active streak in any server has 0 slots available in all servers.

### 11.3 Server-Specific Cached Batches

Each server maintains its own **cached card batch** from the last refresh in that server.

| Action | Result |
|--------|--------|
| User refreshes in Dating server | 5 new Dating cards generated and cached |
| User switches to Friend server | Sees Friend server's last cached batch (if any) |
| User returns to Dating server | Sees the same 5 Dating cards (still cached) |

**When switching to a server with no cached batch:**
- If Refresh is available: Show empty state with Refresh button
- If Refresh is locked: Show lock screen with global timer countdown

### 11.4 Why This Architecture Works

| Principle | Implementation |
|-----------|---------------|
| **Freedom of Intent** | Users can explore different connection types without artificial barriers |
| **Resource Scarcity** | Global inventory forces users to budget their limited actions across servers |
| **Focus Enforcement** | Users cannot spray interests across all servers simultaneously |
| **Clean UX** | Switching is instant; resource state is always clearly communicated |

### 11.5 Why Servers Matter

1. **Reduces expectation mismatch** — Users know what the other person is seeking
2. **Prevents romantic pressure in non-romantic contexts** — Friendship seekers are not exposed to romantic pursuit
3. **Improves matching context** — Intent signal enables better filter-based discovery

---

## 12. Discovery & Matching

### 12.1 Anonymous Cards

The main discovery screen displays anonymous cards that intentionally avoid identity signals.

**Visible on cards:**

- Hobby anchors with micro-descriptions
- Consistency score (banded: Calm / Regular / Highly consistent — not numeric)
- Availability level
- Personality cues (tri-state selectors)
- Intent framing statement

**Not visible:**

- Name
- Face / photos
- Social links
- Exact profession
- Location specifics beyond city

**Visible in Card Detail Modal (on tap):**

- All card content above (expanded)
- Hobby anchors with full micro-descriptions
- Consistency band (full visual)
- Intent statement (full text)
- Personality Intelligence Summary (if available) — see Section 18.7

> [!NOTE]
> The Personality Intelligence Summary is a short, third-person, AI-generated paragraph derived from the viewed user's personality signals. It appears only when sufficient signals exist; otherwise the section is omitted entirely.

### 12.2 Hobby Presentation

Hobbies are displayed as imagination anchors, not simple lists.

**Format:**

```
🏋️ Gym — "Consistency over intensity"
🎨 Painting — "Late nights, slow strokes"
🧘 Yoga — "Morning ritual"
```

These micro-descriptions are pre-written options, not free text. Users select from curated descriptions that convey personality without requiring creative writing.

### 12.3 Discovery Logic (Critical Distinction)

**Users are matched by mutual attraction to hobbies, which encompasses two distinct vectors:**

1. **Shared Passion (Identical):** Users connecting because they love the same activity (e.g., Runner matching with Runner).

2. **Complementary Interest (Distinct):** Users connecting because they value the other's specific discipline (e.g., High-energy Gym-goer matching with Calm Painter).

**Key Principle:** The system does not force "sameness," but it enthusiastically supports it when it drives mutual interest.

**Example (Complementary):**

- User A hobbies: Gym, Hiking, Biking
- User B hobbies: Painting, Cooking, Dancing

**Discovery experience:**

- User A sees: "Anonymous profile — Loves painting"
- User B sees: "Anonymous profile — Consistent gymmer"

**Example (Shared):**

- User A hobbies: Running, Photography
- User B hobbies: Running, Cooking

**Discovery experience:**

- User A sees: "Anonymous profile — Avid runner"
- User B sees: "Anonymous profile — Hit the trails daily"

Both vectors are valid. The attraction is to what the other person does — whether that mirrors the user's own passion or complements it entirely.

### 12.4 Filters

**All filters are available to ALL users regardless of tier.** There is no filter gating.

---

#### 12.4.1 Universal Filters (All Servers)

These filters apply across all servers regardless of intent:

**Discovery & Interest Filters:**

- Interests / hobbies (broad categories)
- Intent alignment
- Lifestyle preference

**Behavioral & Commitment Filters:**

- Consistency preference (desired partner consistency level: Calm / Regular / Highly consistent)
- Time commitment level (light / moderate / dedicated)
- Availability window (morning / afternoon / evening / flexible)
- Communication style (async-friendly / real-time preferred / balanced)
- Response pattern preference (quick responder / thoughtful pacer / flexible)
- Streak history signal (first-timer / experienced / veteran — based on completed streaks)

**Long-Term Alignment Filters:**

- Commitment horizon (exploring / open to long-term / seeking long-term)
- Conflict resolution style (direct discussion / space first / mediator-friendly)
- Growth mindset indicator (stability-seeking / growth-oriented / balanced)
- Energy match preference (high-energy / calm presence / adaptable)

---

#### 12.4.2 Server-Specific Filters

Each server has filters aligned strictly with its intent. Filters that are inappropriate for a server's purpose are deliberately excluded.

---

##### Looking for a Partner (Romantic Connection Intent)

**Eligibility Constraint Filters:**

> [!IMPORTANT]
> Gender and age function as **eligibility constraints only** — they determine who appears in discovery, not how profiles are ranked or prioritized. These are pass/fail gates, not scoring signals.

- **Gender identity** (self-declared by user during onboarding)
- **Genders open to connecting with** (user specifies which genders they wish to see)
- **User age** (required, exact age collected during onboarding)
- **Preferred age range** (min–max range selector; profiles outside this range are excluded from discovery)

**Romantic Compatibility Filters:**

- Relationship intent clarity (casually exploring / open to commitment / seeking committed relationship)
- Physical proximity preference (same city / willing to relocate / open to distance)
- Family planning alignment (wants children / does not want children / undecided / prefer not to say)
- Living situation compatibility (lives alone / with family / with roommates / flexible)
- Lifestyle alignment: dietary preferences (vegetarian / non-vegetarian / vegan / no preference)
- Lifestyle alignment: substance use (non-smoker / social drinker / no preference)
- Relationship experience (first serious relationship / experienced / divorced or separated / prefer not to say)
- Attachment style self-assessment (secure / anxious / avoidant / unsure — optional, educational framing)
- Love language preference (quality time / words of affirmation / acts of service / physical touch / gifts — optional)
- Long-term vision alignment (career-focused / family-focused / balanced / evolving)

---

##### Friend / Companion (Platonic Connection Intent)

> [!NOTE]
> Gender and age filters are deliberately excluded from this server. Friendship discovery is based on shared interests, consistency, and compatibility — not demographic constraints.

**Friendship Compatibility Filters:**

- Friendship style (activity-based / conversation-based / support-based / all-rounder)
- Social energy level (introvert-friendly / extrovert-friendly / ambivert)
- Hangout preference (in-person meetups / virtual connection / mixed)
- Conversation depth preference (light and fun / deep discussions / both)
- Shared interest intensity (casual hobbyist / serious enthusiast / professional level)
- Friendship availability (looking for one close friend / open to multiple friendships / building a circle)
- Communication frequency preference (daily check-ins / weekly catchups / occasional but meaningful)
- Trust-building pace (fast trusters / gradual openers / context-dependent)
- Support style (listener / advisor / cheerleader / practical helper)
- Boundary clarity (clear boundaries preferred / flexible / context-dependent)

---

##### Accountability / Growth Buddy (Goal-Oriented Connection Intent)

> [!NOTE]
> Gender and age filters are deliberately excluded from this server. Accountability partnerships are based on goal alignment, consistency, and mutual commitment — not demographic traits.

**Goal Alignment Filters:**

- Primary goal category (fitness / career / learning / creative / financial / mental health / habit-building / other)
- Goal specificity level (broad direction / specific targets / milestone-driven)
- Accountability style (gentle encouragement / direct feedback / structured check-ins)
- Check-in frequency preference (daily / every few days / weekly)
- Preferred accountability format (text updates / voice notes / shared progress tracking)
- Expertise exchange interest (willing to mentor / seeking mentorship / peer-to-peer only)
- Timeline alignment (short-term sprint / medium-term project / long-term lifestyle change)
- Failure response preference (compassionate reset / analytical review / immediate restart)
- Progress sharing comfort (private milestones / shared celebrations / full transparency)
- Commitment level signal (experimenting / moderately committed / fully dedicated)

---

#### 12.4.3 Filter Design Principles

**Behavioral signals over demographic traits:** Filters prioritize consistency, communication patterns, and commitment indicators. These predict long-term compatibility more reliably than static demographic information.

**Eligibility constraints are binary gates:** Where demographic filters exist (gender/age in Looking for a Partner only), they function as inclusion/exclusion criteria — not as ranking, scoring, or prioritization signals. A user either appears in discovery or does not.

**No appearance-based filtering:** Physical appearance, attractiveness ratings, photo-based filtering, or body-type screening are not available as filters. Discovery is based on behavioral and compatibility signals.

**No swipe mechanics:** Filters reduce the pool to compatible candidates; discovery within that pool is not gamified with swipe-based interfaces.

**Optional filters respect privacy:** Sensitive filters (attachment style, family planning, relationship history) are always optional. Users who skip these filters see all profiles; users who set them see filtered results.

**Why no filter gating:** Filters give agency and reduce mismatch. Restricting them creates frustration and bad matches. Monetization comes from other levers, not from withholding basic control.

### 12.5 Discovery Logic: The Global Double-Lock System

Discovery in Unora is governed by two independent locks that work together to enforce the **Focus** and **Scarcity** philosophies. Both locks must be clear for a user to access the Refresh function.

> [!IMPORTANT]
> **Both locks are GLOBAL across all servers.** A refresh used in the Dating server locks the refresh in Friend and Growth servers. A slot filled in any server reduces available slots in all servers.

---

#### Lock 1: Time-Based Lock (Scarcity) — GLOBAL

This lock controls **when** a user can refresh their discovery cards. It enforces intentional pacing and prevents infinite browsing.

**Batch Size:** Every refresh generates exactly **5 cards**. This is fixed across all tiers.

**Mechanism:** All tiers use a **Rolling Cooldown**. The timer starts immediately when the user taps Refresh.

| Tier | Cooldown Duration |
|------|-------------------|
| **Free** | **24 hours** from the moment of last refresh |
| **Plus** | **12 hours** from the moment of last refresh |
| **Pro** | **6 hours** from the moment of last refresh |

**Global Timer Behavior:**
- The cooldown timer starts from the moment the user taps Refresh **in any server**
- The same timer applies **across all three servers** — Dating, Friend, and Growth
- Users see a countdown timer showing time until next refresh is available
- No refresh accumulation — if a user waits 48 hours, they still get only 1 refresh
- Purpose: Creates anticipation and prevents binge-browsing behavior across servers

**Cross-Server Impact:**

| Action | Result |
|--------|--------|
| Plus user refreshes in Dating server | 12-hour timer starts |
| User switches to Friend server | Refresh button **LOCKED** — shows same 12-hour countdown |
| User switches to Growth server | Refresh button **LOCKED** — shows same 12-hour countdown |
| Timer expires | Refresh available in **current server only** (next refresh starts new global timer) |

---

#### Lock 2: Capacity-Based Lock (Focus) — GLOBAL

This lock controls **whether** a user can access discovery at all. It enforces focused attention on existing connections.

**Rule:** If a user's **Active Connections equal their Tier Limit**, the Discovery/Refresh button is **Hard Locked** (disabled) **in all servers**.

| Tier | Active Connection Limit | Discovery State When At Limit |
|------|------------------------|-------------------------------|
| **Free** | 1 | Hard Locked (all servers) |
| **Plus** | 2 | Hard Locked (all servers) |
| **Pro** | 4 | Hard Locked (all servers) |

**Global Slot Behavior:**
- Connection slots are a **single shared pool** across all servers
- A Free user with 1 active connection in the Friend server has 0 slots available in Dating and Growth servers
- The Connect/Interest button is **DISABLED** in all servers when at capacity
- Users can still browse cached cards, but cannot express new interest

**Cross-Server Impact:**

| Action | Result |
|--------|--------|
| Free user matches in Friend server | Slot 1/1 filled |
| User switches to Dating server | Can view cached cards; **Connect button DISABLED** |
| User's Friend streak ends | Slot freed; Connect button **re-enabled in all servers** |

**Hard Lock Behavior:**
- The Refresh button is visually disabled (greyed out, non-interactive)
- Users cannot express interest or generate new cards
- The lock lifts automatically when an Active Connection slot becomes available (through streak completion, termination, or tier upgrade)

**UX — Locked State Message:**

When a user at capacity attempts to access Discovery in any server, they see:

| User Tier | Message Displayed |
|-----------|------------------|
| **Free** | "You are focused on your current streak. Maintain it, upgrade to Plus or Pro, or end it to browse again." |
| **Plus** | "You are focused on your current streaks. Maintain them, upgrade to Pro, or end one to browse again." |
| **Pro** | "You are focused on your current streaks. Maintain them or end one to browse again." |

---

#### Why the Global Double-Lock System Works

| Lock | Philosophy | User Behavior Impact |
|------|------------|---------------------|
| **Lock 1 (Time)** | Scarcity | Prevents endless browsing; creates anticipation; forces server prioritization |
| **Lock 2 (Capacity)** | Focus | Forces investment in current connections before seeking new ones |
| **Global Scope** | Resource Budgeting | Users must choose which server to invest limited actions in |

**Combined effect:** Users cannot browse infinitely (Lock 1) AND must focus on existing connections before seeking new ones (Lock 2). The global scope ensures users cannot circumvent limits by switching servers. This eliminates the "infinite options" problem that plagues other platforms.

> [!IMPORTANT]
> Both locks operate independently and globally. A user may have Lock 1 clear (cooldown expired) but Lock 2 engaged (at capacity). In this case, discovery remains locked in all servers until capacity frees up. Conversely, a user may have open slots (Lock 2 clear) but still be on cooldown (Lock 1 engaged) across all servers.

### 12.6 Discovery Refresh & Filter Application (Locked)

> [!CAUTION]
> The following behavior is locked and defines the single source of truth for discovery updates.

---

#### Card Generation Rules

- Discovery cards are generated **only when the user explicitly taps the Refresh button**.
- There is no infinite scroll.
- There is no automatic loading of new cards.
- The Refresh button is tier-controlled and may be locked based on the user's subscription tier and cooldown rules (see 12.5).

---

#### Filter Application Rules

| Action                              | Effect                                                     |
| ----------------------------------- | ---------------------------------------------------------- |
| User modifies filters               | No immediate change to displayed cards                     |
| User taps Refresh                   | Filters are evaluated and applied; new cards are generated |
| User views cards without refreshing | Cards remain from most recent refresh                      |

- Filters are evaluated and applied **only at the moment the user taps Refresh**.
- Users may modify filters before refreshing or reuse their existing filters.
- Changing filters alone does not update cards.
- Cards shown always correspond exactly to the filters applied at the most recent refresh action.

---

#### Filter Transparency

A short summary paragraph is displayed at the top of the discovery screen:

> "Showing profiles matching: [active filters summary]"

**Summary behavior:**

- The paragraph explains which filters were used to generate the current set of cards.
- The summary updates only after a successful refresh.
- Users always know exactly why they are seeing a specific set of cards.

**Example:**

> "Showing profiles matching: Looking for a partner · Ages 25–32 · Highly consistent · Evening availability"

---

#### Design Principles

| Principle                                             | Implication                                                           |
| ----------------------------------------------------- | --------------------------------------------------------------------- |
| **Discovery is intentional, not passive**             | Users must take explicit action to see new profiles                   |
| **Transparency over mystery**                         | Users always understand their current filter state                    |
| **Paid tiers unlock speed, not quality**              | Faster/more frequent refreshes, not better filters or better profiles |
| **No visibility advantages beyond refresh frequency** | All tiers see the same quality of profiles                            |

---

#### Constraints (Locked)

| Constraint                              | Reason                                                 |
| --------------------------------------- | ------------------------------------------------------ |
| ❌ No infinite browsing                 | Prevents endless scrolling and attention fragmentation |
| ❌ No auto-refresh                      | User maintains control; no surprise content changes    |
| ❌ No background filter re-evaluation   | Cards remain stable until explicit refresh             |
| ✓ Refresh is the single source of truth | All discovery updates flow through the Refresh action  |

---

#### Tier Consistency

This logic applies consistently across Free, Plus, and Pro tiers. The only differences are:

| Tier     | Refresh Availability     | Cooldown Duration           |
| -------- | ------------------------ | --------------------------- |
| **Free** | No manual refresh        | N/A (daily profiles only)   |
| **Plus** | Manual refresh available | ~12 hours between refreshes |
| **Pro**  | Manual refresh available | ~6 hours between refreshes  |

Filter access, filter transparency, and card generation rules are identical across all tiers.

### 12.7 Mutual Connection Flow

1. **Browse** — User A views anonymous cards
2. **Express Interest** — User A requests to connect with User B
3. **Independent Interest** — User B (without knowing A's interest) independently requests to connect with User A
4. **Mutual Lock** — Only when both have expressed interest does the system reveal the match

Copy example at match:

> "You've found a streak partner. Show up daily to unlock who they really are."

There is no chat. There is no profile reveal. The streak begins.

### 12.8 Active Connection Limits

| Tier     | Active Connections |
| -------- | ------------------ |
| **Free** | 1                  |
| **Plus** | 2                  |
| **Pro**  | 4                  |

**Why cap Pro at 4:** Beyond this threshold, attention fragments and connection quality drops. This limit serves the product thesis, not monetization constraints.

### 12.9 Concurrency & The "Total Wipe" Rule

This section defines how the system handles the edge case where a user sends interest to multiple profiles (up to 5 in a batch) and one or more recipients match back.

---

#### The Concurrency Problem

**Scenario:** User A (Free tier, 1 slot limit) expresses interest in 5 profiles from their batch. Multiple recipients may independently confirm mutual interest. Who gets the slot, and what happens to the others?

---

#### Protocol: First-to-Lock Rule

The system uses a **First-to-Lock** protocol based on server-side timestamp of match confirmation:

1. **Match Confirmation Timestamp:** When a recipient confirms mutual interest, the system records the exact server timestamp
2. **Slot Claim:** The first mutual match to confirm (lowest timestamp) immediately claims the open slot and becomes an **Active Connection**
3. **Streak Begins:** The Active Connection initiates a streak immediately

---

#### The "Total Wipe" (Commitment Rule)

**Immediately upon the start of an Active Streak, ALL other pending outgoing interests sent by that user are permanently deleted.**

| Event | System Action |
|-------|---------------|
| Mutual match confirmed → Streak begins | All other pending outgoing interests from this user are **wiped** |
| User at capacity (all slots filled) | Discovery is Hard Locked (Lock 2 engaged) |

**What gets deleted:**
- All unmatched outgoing interest expressions from the user's most recent batch
- Any older pending interests that were never reciprocated

**What is NOT affected:**
- The newly Active Connection (protected)
- Any other existing Active Connections (if user has multiple slots)
- Inbound interest from others (handled separately)

---

#### Rationale: Why Total Wipe?

> [!IMPORTANT]
> A user cannot maintain "backup options" while in a streak. Focus must be absolute.

| Principle | Implication |
|-----------|-------------|
| **No Hedging** | Users cannot keep 4 potential matches "warm" while pursuing 1 |
| **Commitment Signal** | Starting a streak means full commitment to that connection |
| **Fairness to Recipients** | Other users are not left waiting indefinitely for a response |
| **System Simplicity** | No queue management, no pending states, no stale matches |

---

#### Outcome for Other Recipients

**The other recipients (who did not match fast enough) will never see the request.**

- The interest expression is deleted before they can act on it
- They receive no notification of a "missed connection"
- From their perspective, it is as if the interest was never sent

**Why this is acceptable:**
- They never knew about the interest, so there is no disappointment
- No false hope or dangling notifications
- Clean system state

---

#### Recycling Strategy: Fresh Intent, Not Stale Likes

**Since pending requests are deleted (not rejected), the algorithm is free to re-suggest those same profiles to the user in future Discovery cycles.**

| State | Algorithm Behavior |
|-------|-------------------|
| Interest deleted via Total Wipe | Profile eligible for re-suggestion in future refreshes |
| User explicitly passed/rejected | Profile excluded from future suggestions |

**Why this matters:**
- If the user connects with that profile later, it is based on **fresh, current intent**
- Not a stale "like" from weeks ago when circumstances were different
- Both parties are evaluating each other in the present moment
- Prevents the accumulation of outdated interest signals

---

#### Edge Case: User Has Multiple Slots (Capacity-Triggered Wipe)

For Plus (2 slots) and Pro (4 slots) users, the Total Wipe is **capacity-triggered**, not match-triggered. Pending interests remain active until the user reaches full capacity.

| Scenario | Slots State | Match Event | System Behavior |
|----------|-------------|-------------|----------------|
| **A: Slots Open** | 0/2 or 1/4 Active | First match confirms | Match fills an open slot; **pending interests remain active** (no wipe yet) |
| **B: One Slot Left** | 1/2 or 3/4 Active | Next match confirms | Match fills the **final slot**; **Total Wipe immediately executes** — all remaining pending interests deleted |
| **C: Full** | 2/2 or 4/4 Active | N/A | Discovery is Hard Locked; no new interests can be sent |

**Example Flow (Plus user with 2 slots):**

1. User has 0 Active Connections, sends interest to 5 profiles
2. Profile B matches → fills Slot 1 → **pending interests to C, D, E remain active**
3. Profile D matches → fills Slot 2 (final slot) → **Total Wipe executes** → pending interests to C, E deleted
4. Discovery is now Hard Locked until a slot frees up

**Why Capacity-Triggered:**
- Users with open slots should be able to fill them naturally
- Focus is enforced only when the user is **fully committed** (all slots occupied)
- No artificial restrictions while the user still has room for more connections

---

#### Why This System Matters

| Principle | Implementation |
|-----------|---------------|
| **Capacity-Based Focus** | Focus is enforced only when the user is fully committed (all slots filled) |
| **Clean State** | No pending queues, no stale matches, no maintenance burden |
| **Fresh Connections** | Future matches are based on current intent, not old signals |
| **Respect for All Users** | No one is left waiting indefinitely; clean experience for everyone |

---

## 13. Streak System (Locked)

The streak is the core behavioral loop. It is not gamification. It is alignment.

---

### 13.1 Streak Definition

A streak represents **shared consistency** between two connected users. It advances only when both users check in on the same day.

- The streak belongs to both parties equally
- Progress requires mutual action
- Partial completion does not advance the streak

---

### 13.2 Daily Check-In Mechanism

**The check-in is:**

- One tap
- With a micro-prompt (question-based)
- No free text required
- Same effort for all tiers

**Rules:**

- Each user has one daily check-in
- Both users must check in on the same day for the streak to increment
- Partial completion (only one user checking in) does not advance the streak

**Purpose:** Show up, not perform.

The check-in is intentionally low-friction to complete but high-meaning to accumulate. It asks nothing complex — only presence.

---

#### 13.2.1 The Hobby Echo (Contextual Progress)

Instead of a static "Check-In Complete" screen, Unora uses AI to share activity progress without sharing identity.

**Input:** Users answer a dynamic, tap-based question about their specific hobby (e.g., "Leg day or cardio?" → "Leg day").

**Processing:** AI sanitizes this input and frames it as a progress update.

**Output:** Upon mutual check-in, users see their partner's progress summary (e.g., "Your partner crushed Leg Day today").

**Privacy Rule:** Strictly limits information to the activity domain. No location, names, or specific details are shared.

**Purpose:** Proves that the partner is real and active, creating connection through shared effort before the conversation begins.

---

### 13.3 Streak States

```
┌─────────────────────────────────────────────────────────────────┐
│                         STREAK STATES                           │
├─────────────┬───────────────────────────────────────────────────┤
│ ACTIVE      │ Both users checked in today. Streak counter       │
│             │ increments. Normal progression.                   │
├─────────────┼───────────────────────────────────────────────────┤
│ AT RISK     │ One user has not checked in by end of day.        │
│             │ Nudges enabled. No payment yet.                   │
├─────────────┼───────────────────────────────────────────────────┤
│ PAYMENT     │ Day N+1. Breaker sees payment option only.        │
│ WINDOW      │ No late check-in available. Nudges remain.        │
├─────────────┼───────────────────────────────────────────────────┤
│ RESET       │ Payment window expired without payment.           │
│             │ Streak returns to 0. Connection remains active.   │
└─────────────┴───────────────────────────────────────────────────┘
```

---

### 13.4 Streak Flow (Locked)

#### Day N — Missed Check-In

When one user fails to check in by the end of Day N:

- Streak moves to **AT RISK** state
- The maintaining user (who checked in) is notified: "Your streak is at risk"
- Nudges are enabled for the maintaining user
- **No payment is requested on Day N**
- **No late or recovery check-in is allowed**

The breaker cannot fix the situation on Day N. The day ends with the miss recorded.

#### Day N+1 — Payment Window (Breaker Only)

On Day N+1:

- Only the user who missed the check-in (the breaker) is prompted to pay to continue the streak
- This is the **only recovery action available**
- The maintaining user is never asked to pay
- Nudges remain enabled for the maintaining user
- **No check-in option exists on this day** — recovery is payment-only

#### End of Day N+1

At the end of the payment window:

- **If payment is completed:** Streak continues from its previous count
- **If payment is not completed:** Streak resets to 0

> [!IMPORTANT]
> A reset does not terminate the connection. The connection remains active; only the streak counter returns to zero.

---

### 13.5 Case: Both Users Miss on the Same Day

When both users fail to check in on the same day:

- Streak resets to 0 immediately
- No payment window opens (neither user is the sole breaker)
- No blame assigned to either party
- Connection remains active
- Both users may continue or terminate at any time

---

### 13.6 Reset Behavior

When a streak resets:

1. Streak counter returns to 0
2. The connection remains active unless explicitly ended by a user
3. Neither user is blamed or punished
4. After reset, users may continue building a new streak or terminate the connection at any time

**Reset is a pause, not an ending.** The relationship continues; only the counter restarts.

---

### 13.7 Termination

**Either user may terminate the connection at any time.**

Termination rules:

- Termination permanently ends the streak and the connection
- Both users return to discovery
- No penalty language or blame is surfaced

**Credit Conversion on Bad-Faith Termination:**

If a paid recovery is followed by termination by the non-paying user within 24 hours, credit conversion rules apply (as defined in Section 17: Credits & Payment Protection). This protects breakers from paying for recovery only to have the other user immediately leave.

---

### 13.8 Who Pays and Why

**The breaker pays. The maintaining user never pays.**

Rationale:

- The person who broke the streak is responsible for its recovery
- The person who showed up did nothing wrong — they should not bear cost
- This creates clear accountability without blame language
- Payment is for continuation, not for forgiveness

**What payment buys:**

- Immediate streak continuation from the previous count
- Not time to check in — payment itself is the recovery action

---

### 13.9 Language Principles

The system never uses:

- "You broke the streak"
- "Your partner failed"
- "Penalty" or "punishment"
- Accusatory framing

Instead:

- "The streak is at risk"
- "Your partner hasn't checked in yet"
- "Would you like to continue?"
- "This connection has ended"

All copy maintains dignity for both parties.

---

### 13.10 Locked Principles

> [!CAUTION]
> The following principles are non-negotiable and define the streak system's integrity.

| Principle                                   | Implication                                               |
| ------------------------------------------- | --------------------------------------------------------- |
| **Only the breaker pays**                   | The maintaining user is never prompted or expected to pay |
| **Recovery is payment-only**                | No late check-ins; missed days cannot be "made up"        |
| **No grace periods on Day N**               | The miss is recorded; action happens on Day N+1           |
| **Reset ≠ Termination**                     | A reset returns the counter to 0; the connection persists |
| **Money buys continuation, never progress** | Payment restores the streak count; it does not add days   |
| **Either user may leave at any time**       | Consent is preserved; no hostage dynamics                 |

---

## 14. Nudge System

### 14.1 What a Nudge Is

A nudge is:

- One tap to send
- No typing required
- No obligation created
- No embarrassment caused

Think of it as a **tap on the shoulder**, not a message. It signals "I showed up — I hope you do too" without pressure or accusation.

### 14.2 When Nudges Appear

Nudges are available when ALL of the following are true:

- Person A has checked in
- Person B has not checked in
- The streak is AT RISK
- Person A has not exhausted their nudge allowance for this at-risk period

Person A sees:

> "Your streak is at risk. Want to nudge your partner?"

### 14.3 What Happens When a Nudge Is Sent

Person B receives:

> "Your partner checked in today and would like to keep the streak going."

The nudge contains:

- No name
- No accusation
- No urgency language
- No read receipt

### 14.4 Nudge Limits by Tier

| Tier     | Nudge Frequency | Max Nudges per At-Risk Period |
| -------- | --------------- | ----------------------------- |
| **Free** | 1 per 24 hours  | 1 total                       |
| **Plus** | 1 per ~7 hours  | 3 total                       |
| **Pro**  | 1 per ~5 hours  | 4 total                       |

**Nudges are bounded for all tiers.** Even Pro users cannot spam. This protects trust and prevents pressure.

### 14.5 Copy Principles

All nudge copy is:

- **Pre-written** — Users select from options, not compose
- **Non-judgmental** — No blame, guilt, or shame
- **Contextual** — Appropriate to the situation
- **AI-selected** — System chooses which variant based on context

**Examples of nudge variants:**

| Type              | Example Copy                                       |
| ----------------- | -------------------------------------------------- |
| **Neutral**       | "Your partner checked in today."                   |
| **Encouragement** | "Something is building between you. Keep showing up." |
| **Commitment**    | "They're waiting to continue the streak with you." |

### 14.6 Why Nudges Work Psychologically

1. **Reframes the streak as shared** — Both parties own it together
2. **Triggers reciprocity** — "They showed up; I should too"
3. **Avoids shame** — No accusation, only invitation
4. **Creates just enough social pull** — Enough to motivate, not enough to pressure

### 14.7 Why Free Users Are Not Powerless

Free users have limited nudges but they are not without tools:

- They can still nudge once per at-risk period
- Their nudge carries the same weight as any other
- The core mechanic (showing that they showed up) works regardless of tier
- The limitation affects frequency, not effectiveness

### 14.8 Why Nudges Are Bounded

Unlimited nudges would:

- Create pressure dynamics
- Enable harassment-adjacent behavior
- Undermine the value of each nudge
- Shift responsibility to the non-missing user

**Bounded tools create trust. Unlimited tools create pressure.**

---

## 15. Reveal System

### 15.1 Reveal Philosophy

Reveals are **earned, not purchased outright.** Consistency unlocks information. Money can accelerate timing within tier-specific bounds but cannot bypass the requirement for effort.

Maximum total reveals per connection: **5 (hard cap for all tiers)**

---

### 15.1.1 The Mystery Reveal Mechanic (Critical UX Principle)

> [!IMPORTANT]
> **Users never see countdown timers, specific day markers, or visible roadmaps for reveals.** The reveal system operates on **Variable Reward** psychology — users know consistency leads to rewards, but they never know exactly when.

**Internal Logic vs. User Experience:**

| Layer | Behavior |
|-------|----------|
| **Backend (Internal)** | System tracks milestone days (Day 5, 10, 15 etc.) and triggers reveals on schedule |
| **Frontend (User-facing)** | Users see only an abstract "Connection Strength" visual that grows/glows with streak consistency |

**What Users See:**
- A premium, abstract visual element ("Trust Orb" / "Connection Core") that intensifies with consistency
- Vague, intriguing copy: "Connection deepening..." / "Trust building..." / "Something is forming..."
- **NO countdown text** ("Next reveal in 2 days")
- **NO numbered progress bars** ("Day 5 of 15")
- **NO visible milestone markers** ("R1, R2, R3, R4")

**What Happens:**
- When a milestone day is reached internally, a **Surprise Reveal Modal** appears unexpectedly
- The reveal feels like a gift / discovery, not a predictable reward
- This unpredictability maximizes dopamine and emotional investment

**Psychological Rationale:**

| Principle | Implication |
|-----------|-------------|
| **Variable Reward Loops** | Unpredictability creates "Slot Machine" psychology — users stay engaged because the next reward could be imminent |
| **Anticipation > Expectation** | Knowing *something* will happen is more engaging than knowing *exactly when* |
| **Surprise Amplifies Emotion** | An unexpected reveal creates stronger emotional response than a predicted one |
| **Reduces Gaming Behavior** | Users cannot optimize around visible milestones; they simply maintain consistency |

### 15.2 Finalized Reveal Sequence

The reveal sequence is designed to be non-colliding (each reveal contains distinct categories) and emotionally progressive (building from context to identity).

---

#### Reveal 1 — Personality Spark

**Unlock:** First reveal milestone (varies by tier)

**Contains:**

- Age (exact)
- Height (range)
- City area (neighborhood/district)
- Primary hobby depth (beginner/intermediate/serious)
- One personality signal

**Purpose:** Establishes basic demographic grounding and first personality impression. Just enough to anchor imagination without revealing identity.

**Why this is not a filter:** Age and location can be filtered in discovery; this reveal provides exact values and adds personality context that filtering cannot capture.

---

#### Reveal 2 — Lifestyle Reality

**Unlock:** Second reveal milestone

**Contains:**

- Profession domain (not exact title)
- Education level
- Body-type descriptor (optional, user-controlled opt-in)
- Daily routine type (morning person / night owl / variable)

**Purpose:** Creates expectation alignment around lifestyle compatibility. Career context, education level, and physical self-description provide substance to the anonymous connection.

**Why this is not a filter:** Profession domain can be broadly filtered; this provides more specific indication. Body type and routine are not filterable.

---

#### Reveal 3 — Social & Emotional Context

**Unlock:** Third reveal milestone

**Contains:**

- Religion (optional, opt-in only)
- Ethnicity/cultural background (broad category, optional)
- Language background
- One emotional/social signal chosen by user

**Purpose:** Provides cultural and social grounding. Users who prioritize cultural alignment receive signal; users who do not prioritize it see what their partner chose to share.

**Why this is not a filter:** This information is sensitive and opt-in. Revealing it after consistency is established reduces judgment and increases openness.

---

#### Reveal 4 — Human Presence

**Unlock:** Fourth reveal milestone

**Contains:**

- Voice note (short, recorded during onboarding)
- One moment snapshot (candid, non-face photo)
- One immutable fact they chose to share

**Purpose:** Humanizes the connection. Voice conveys emotion, cadence, and personality that text cannot. The non-face photo creates visual context without full identity exposure.

**Why this is not a filter:** Voice and candid presence are inherently unfilter able. They exist to create emotional response, not to screen.

---

#### Reveal 5 — Identity + Chat Unlock

**Unlock:** 15-day milestone (NOT accelerable by any tier or payment)

**Contains:**

- Full name
- Photos (face photos submitted during onboarding)
- Chat functionality unlocked

**Purpose:** This is the culmination of the streak. Identity is revealed only after 15 days of mutual consistency. Both parties have invested; both have earned this moment.

**Why identity is never accelerated:** The 15-day identity reveal is the core product promise. If it could be purchased, the entire system loses meaning. Paid acceleration applies to Reveals 1–4 only.

---

### 15.3 Reveal Timing by Tier (Internal Logic — Never Surfaced to Users)

> [!CAUTION]
> The following timing table is **internal system logic only**. Users never see these specific day numbers. They experience reveals as surprise events triggered by their consistent behavior.

Reveals 1–4 occur on cadence determined by tier. Reveal 5 (identity + chat) always requires 15 days.

| Tier     | Internal Reveal Cadence | Earned Reveals |
| -------- | ----------------------- | -------------- |
| **Free** | Day 5, 10, n/a, n/a     | 2 earned       |
| **Plus** | Day 4, 8, 12, n/a       | 3 earned       |
| **Pro**  | Day 3, 6, 9, 12         | 4 earned       |

**Note:** Free users earn 2 reveals; Plus earns 3; Pro earns 4. The 5th reveal (identity) is consistent across all tiers at Day 15.

**User Experience:** Instead of seeing "Next reveal in 3 days", users see their abstract Connection Strength visual grow brighter. When the internal milestone is reached, a Surprise Reveal Modal appears.

### 15.4 Paid Reveal Logic

Users may purchase additional reveals within caps:

| Tier     | Max Paid Reveals |
| -------- | ---------------- |
| **Free** | Up to 3 paid     |
| **Plus** | Up to 2 paid     |
| **Pro**  | Up to 1 paid     |

**Total reveals (earned + paid) = 5 maximum for everyone**

This creates symmetry: all users can access all 5 reveals, but tiers differ in how many are earned vs purchased.

**Pricing:**

- Paid reveals: ₹79–110 per reveal

### 15.5 AI-Enhanced Reveal Framing

Raw profile data is never dumped directly. AI reframes factual truth into sensory, human language.

**Example:**
| Raw Data | AI-Framed Reveal |
|----------|------------------|
| "Mechanical Engineer, Bangalore" | "They build things that move — in a city that wakes up to filter coffee." |

This increases emotional resonance without deception. Same truth, more dopamine.

---

## 16. Monetization

### 16.1 Core Monetization Philosophy

Unora monetizes:

- **Time** — Streak recovery buys time to check in
- **Flexibility** — Multiple connections allow parallel exploration
- **Speed** — Reveal acceleration reduces waiting (within bounds)
- **Control** — Faster refresh, more nudges

Unora does NOT monetize:

- **Another person's attention** — Cannot pay for visibility to specific users
- **Forced visibility** — No "boost" mechanics that increase exposure unilaterally
- **Messaging rights** — Cannot pay to message someone who hasn't earned chat
- **Bypass of effort** — Cannot skip streak requirements for identity reveal

**Money buys flexibility, never outcomes.**

### 16.2 Tier Structure

#### Free Tier — The Philosophy Trial

**Active connections:** 1
**Discovery refresh:** Limited daily profiles, no manual refresh
**Filters:** All filters available
**Nudges:** 1 per 24 hours (1 total per at-risk period)
**Reveals:** 2 earned (Day 5, 10)
**Paid reveals:** Up to 3 purchasable
**Streak recovery:** Must pay every time

Free tier is generous on access (filters, reveals possible) but constrained on forgiveness and parallelism. This is intentional:

- One mistake returns them to zero
- One mismatch means restart
- No hedging across multiple connections
- They feel the weight of each decision

**Free lets you try the philosophy. Paid lets you practice it comfortably.**

---

#### Plus Tier — ₹199/month (₹1,920/year → ~₹160/month)

**Active connections:** 2
**Discovery refresh:** Higher daily cap, refresh every ~12 hours
**Filters:** All filters available
**Nudges:** 1 per ~7 hours (max 3 per at-risk period)
**Reveals:** 3 earned (Day 4, 8, 12)
**Paid reveals:** Up to 2 purchasable
**Streak recovery:** 1 free recovery per connection, then paid

**Value proposition:**

- More than 1 connection → comparison without chaos
- Fewer resets / less stress → 1 free recovery per connection
- Faster clarity → accelerated reveal schedule
- More control → faster refresh

**User justification:** "I don't want to lose something I've started."

---

#### Pro Tier — ₹399/month (₹3,840/year → ~₹320/month)

**Active connections:** 4
**Discovery refresh:** Highest daily cap, refresh every ~6 hours (still finite)
**Filters:** All filters available
**Nudges:** 1 per ~5 hours (max 4 per at-risk period)
**Reveals:** 4 earned (Day 3, 6, 9, 12)
**Paid reveals:** Up to 1 purchasable
**Streak recovery:** 2 free recoveries per connection, then paid

**Additional Pro features:**

- Priority nudge delivery
- Context-rich nudge variants

**Value proposition:**

- Maximum parallel connections (still capped at 4)
- Maximum safety net (2 free streak recoveries)
- Fastest emotional payoff (earliest reveal schedule)
- Highest agency

**User justification:** "I care about this. I want control, not chance."

---

### 16.3 Additional Monetization Levers

| Item                          | Price        |
| ----------------------------- | ------------ |
| Paid reveals                  | ₹79–110 each |
| One-time streak recovery      | ₹99          |
| Future: Founding member badge | ₹999/year    |

### 16.4 Why This Pricing

**₹199 / ₹399 is intentional, not arbitrary:**

1. **Not a volume swipe app** — Unora competes for attention and emotional effort, not with Tinder's ₹99 promos
2. **Indian psychology sweet spot:**
   - ₹99 feels disposable → attracts unserious users
   - ₹199 feels considered but affordable
   - ₹399 feels premium but not elite
3. **Pricing as intent filter** — The product needs friction to filter intent; pricing is part of that filter

### 16.5 Why Monetization Reinforces Trust

| Design Choice                       | Trust Implication                 |
| ----------------------------------- | --------------------------------- |
| Cannot buy attention                | Users know interest is genuine    |
| Cannot skip streak                  | Identity reveal is earned by both |
| Breaker pays, not victim            | Responsibility is clear and fair  |
| Free has all filters                | Control is not gated              |
| Recovery buys time, not forgiveness | Still requires action             |

---

## 17. Credits & Payment Protection

### 17.1 Credit Conversion Rule

**If a paid streak recovery is followed by the other user ending the connection within 24 hours, the payment is automatically converted to in-app credits for future use.**

This is automatic:

- No manual action required
- No support tickets
- No ambiguity

### 17.2 When Credits Are Triggered

**Credits ARE issued if:**

- User A pays for streak recovery
- AND the other user (B) terminates the connection
- AND this termination happens within 24 hours of A's payment

**Credits are NOT issued if:**

- User A pays and A themselves ends the connection (paying user chose to leave)
- The streak resets naturally without explicit termination
- The connection ends after 24 hours
- Both users mutually drift without explicit termination
- The user just regrets paying

### 17.3 Who Gets Credits

Only the person who paid. The non-paying user never benefits from credits.

This prevents:

- Coordinated abuse
- Infinite credit loops
- Confusion about responsibility

### 17.4 What Credits Can Be Used For

**Allowed:**

- Paid reveals
- Future streak recoveries
- Other in-app purchases (excluding subscriptions)

**Not allowed:**

- Subscription payments
- Cash withdrawal
- Gifting to other users

**Credits are a closed-loop currency.**

### 17.5 UX Implementation

Payment confirmation includes microcopy:

> "If this connection ends within 24 hours, this amount will be converted to credits for future use."

No banners. No heavy explanation. Just reassurance at point of payment.

### 17.6 Purpose

- **Protect good-faith users** — Prevents resentment after paying
- **Prevent infinite credit loops** — Constrained rules prevent gaming
- **Maintain trust** — Users know they won't be "cheated"
- **Keep money in ecosystem** — Credits encourage continued engagement

---

## 18. AI Role

> [!IMPORTANT]
> **Unora uses intelligence to moderate pacing and translate signals, not to predict human compatibility.**

AI in Unora operates quietly, where it matters. It is not a feature to market. It is infrastructure for better outcomes.

### 18.0 AI Philosophy

AI operates as core infrastructure throughout Unora — powering pacing, safety, and contextual intelligence without exposing mechanics to users.

**AI Continuously Enables:**

- **Adaptive pacing** — Reveals occur when connections are healthy
- **Behavioral pattern recognition** — Risk detection and trust evaluation
- **Signal translation** — Structured data becomes human-readable context
- **Media quality maintenance** — Profiles remain genuine and usable

**System Design Principles:**

| Principle | Implementation |
|-----------|----------------|
| Matching is structural | Rule-based: mutual interest + filters + availability |
| Reveal order is fixed | AI modulates timing; sequence is invariant |
| AI enhances, never replaces | Human effort drives the streak; AI provides context |

---

### 18.1 Reveal Timing Modulation

AI influences **when** reveals occur, never **what order** they occur in.

**Inputs:**

- Streak health signals (check-in consistency, recovery frequency)
- Behavioral patterns (engagement strength, response latency)

**Output:**

- **Proceed** — Reveal occurs at scheduled milestone
- **Pause** — Reveal is held briefly (streak health is weak)
- **Hold** — Reveal is delayed until engagement stabilizes

**Constraints:**

| Constraint | Implementation |
|------------|----------------|
| **No reordering** | Reveal sequence is FIXED (1→2→3→4→5); AI cannot change order |
| **No acceleration** | AI cannot make reveals happen faster than tier-defined schedule |
| **Timing only** | AI affects delay/hold decisions, never content or sequence |

**Purpose:** Protects emotional investment by ensuring reveals happen when the connection is healthy, not when it's at risk.

---

### 18.2 Behavioral Risk & Trust Evaluation

A unified backend system that combines behavioral signals for pacing, safety, and review decisions.

**Input Signals (Merged):**

- Check-in timing consistency
- Recovery frequency and patterns
- Nudge response latency
- Past streak failures
- Device fingerprinting anomalies
- Bot-like behavioral patterns
- Manipulation indicators (post-chat unlock)
- Verification result anomalies

**Output:**

- Hidden risk/trust assessment (never shown to users)
- Pacing recommendations for reveal timing
- Safety flags for human review
- Nudge selection context

**Actions:**

- Adjust reveal timing (via 18.1)
- Select appropriate nudge variants (pre-written, not AI-generated)
- Silent throttling for flagged accounts
- Escalation to human review

**Constraints:**

| Constraint | Implementation |
|------------|----------------|
| **Fully backend** | Users never see scores, assessments, or flags |
| **No matching influence** | Risk signals do NOT affect who appears in discovery |
| **No ranking** | Users are not ordered or prioritized by trust score |
| **No visible punishment** | Safety interventions are quiet and protective |

---

### 18.3 AI-Framed Language (Translation Layer)

AI converts structured numeric or state signals into calm, human language. This is a **stateless translation function** with no memory or judgment.

#### 18.3.1 Reveal Framing

Raw profile data is never dumped directly. AI reframes factual truth into sensory, human language.

**Example:**

| Raw Data | AI-Framed Reveal |
|----------|------------------|
| "Mechanical Engineer, Bangalore" | "They build things that move — in a city that wakes up to filter coffee." |

This increases emotional resonance without deception. Same truth, more dopamine.

#### 18.3.2 The Contextual Echo Engine (Hobby Echo)

**Function:** Translates the user's tap-based check-in answer into a motivational summary for their partner.

**Process:**

1. User selects generic option (e.g., "High Intensity", "Rest")
2. AI sanitizes and frames this as a progress update
3. Partner sees: "Your partner had a High Intensity session today"

**Constraints:**

- Strictly limits information to the activity domain
- No location, names, or specific details shared
- Pre-written framing patterns, not free-generated text

#### 18.3.3 Personality Intelligence Summary

**Function:** Translates structured personality signals (numeric scores + confidence levels) into a short, third-person, observational paragraph visible to other users in the Card Detail Modal.

**Inputs:**

- Probabilistic personality dimension scores (0–100)
- Confidence values for each dimension
- No text history, no cached summaries, no previous outputs

**Output:**

- A 50–90 word third-person paragraph
- Calm, observational, non-judgmental tone
- No labels, no scores, no categories

**Generation Rules:**

- Generated fresh on demand (stateless)
- Never stored as text
- Visible only to OTHER users (never to self)
- Displayed inside Discovery Card Detail Modal

**What AI Does NOT Do (Personality-Specific):**

| AI Does Not...                        | Rationale                                            |
| ------------------------------------- | ---------------------------------------------------- |
| Generate or ask personality questions | Questions are system-defined, not AI-generated       |
| Assign scores or categories           | AI converts existing scores to language only         |
| Influence matching or discovery order | Summary is for human context, not algorithmic input  |
| Remember previous summaries           | Each generation is stateless from numeric source     |

#### 18.3.4 Personality Signal Collection

Personality signals are collected gradually through two optional mechanisms:

1. **Initial Setup (Optional):** During or after profile creation, users may complete a brief personality context exercise
2. **Streak Check-In Micro-Questions (Optional):** Occasionally, a single lightweight question may be injected during daily check-ins

**Collection Rules:**

| Rule | Implementation |
|------|----------------|
| **Always optional** | Users may skip setup or micro-questions without penalty |
| **Non-blocking** | Skipping does NOT affect discovery visibility or matching |
| **Accumulative** | Numeric state grows over time; summaries improve with more signals |
| **System-defined** | AI NEVER generates the questions — all questions are pre-written |

---

### 18.4 AI-Powered Media Quality & Authenticity Filtering

AI continuously evaluates uploaded media to maintain profile quality and authenticity.

**Capabilities:**

| Capability | Description |
|------------|-------------|
| **Quality Evaluation** | Detects blur, low-resolution, or unusable images |
| **Face Presence Detection** | Verifies at least one photo contains a clear human face |
| **Authenticity Signals** | Filters out visual noise that degrades discovery quality |

**How It Works:**

- **Passive** — Runs automatically on all uploaded media
- **Non-interruptive** — No user action required
- **Always-on** — Operates in the background throughout the platform

**Purpose:** Maintains baseline authenticity and usable profiles without disrupting user experience.

---

### 18.5 AI System Boundaries

> [!IMPORTANT]
> These boundaries define where AI operates and where it does not.

| AI Operates On... | AI Does Not Touch... |
|-------------------|----------------------|
| Reveal timing modulation | Who matches with whom |
| Risk pattern detection | Compatibility scoring |
| Personality signal interpretation | User rankings or priority ordering |
| Media quality filtering | Raw user-to-user messaging |
| Contextual language framing | Consent-based milestones |

**Core Invariant:** AI enhances system intelligence; matching and connection remain user-driven.

---

## 19. Trust, Safety & Ethics

### 19.1 Anonymity-First Design

- Users begin completely anonymous
- No username or display name visible initially
- Photos are held until Reveal 5 (Day 15)
- Discovery is based on hobbies and consistency, not appearance

### 19.2 Gradual Disclosure

- Information is revealed in stages over 15 days
- Each reveal is a milestone earned through consistency
- Users cannot "front-load" sharing to accelerate connection
- The pace is system-controlled, not user-controlled

### 19.3 Consent at Every Stage

- Mutual interest required before streak begins
- Either party can exit at any time without penalty
- Reveals are based on data users provided during onboarding; no surprises
- Chat is unlocked only after both parties complete the 15-day journey

### 19.4 No Public Profiles

- Users do not have profiles visible to anyone outside an active connection
- There is no "viewing" of profiles outside the discovery/match context
- No follower counts, popularity signals, or exposure metrics

### 19.5 No Exposure-Based Ranking

- Discovery is not based on who is "most viewed" or "most liked"
- There are no leaderboards, popularity rankings, or visibility scores
- Users cannot pay to be "more visible" in general

### 19.6 Silent Safety Interventions

When safety concerns arise:

- AI throttles visibility quietly
- Flagged accounts are reviewed without public notification
- Terminations (when necessary) are handled without shaming language
- No user is told they were "reported" unless required by process

### 19.7 Ethical Monetization Stance

| Ethical Commitment                 | Implementation                                                      |
| ---------------------------------- | ------------------------------------------------------------------- |
| Money never buys access to people  | No "super likes" or attention-purchasing                            |
| Outcomes are earned, not purchased | 15-day identity reveal cannot be accelerated                        |
| Free users are not second-class    | All filters available; reveals possible via purchase                |
| The breaker pays, never the victim | Streak recovery is only the responsibility of the person who missed |
| Commitments are protected          | Credits issued if connection ends within 24h of payment             |

### 19.8 Progressive Trust Verification

Trust is built through behavior, not bureaucracy:

| Verification Layer | Purpose | User Experience |
|--------------------|---------|-----------------|
| **Phone + Photos** | Establish personhood | Upfront, mandatory, low friction |
| **Liveness Challenge** | Confirm photo authenticity | In-streak, optional, non-alarming |
| **Behavioral Signals** | Prove consistent engagement | Passive, earned through actions |
| **Government ID** | Optional identity confirmation | Post-trust, unlocks flexibility |

**Safety through progressive verification:**

- Photo liveness prevents profile catfishing
- Behavioral patterns detect bad actors over time
- Consistency requirements filter out low-intent users
- No claims of "verified identity" without government ID completion

### 19.9 Verification Copy Principles

> [!IMPORTANT]
> Language around verification must be calm and trust-forward:

| Never Say | Instead Say |
|-----------|-------------|
| "KYC" | "Trust verification" |
| "Government verification required" | "Optional identity confirmation" |
| "You must verify" | "Build your trust profile" |
| "Verification failed" | "Let's try again" |
| "Unverified account" | "Trust building in progress" |

### 19.10 Personality Intelligence Privacy

The Personality Intelligence Summary adheres to strict privacy principles:

| Principle | Implementation |
|-----------|----------------|
| **Self-visibility prohibited** | Users never see their own personality summary |
| **No raw data exposure** | No answers, scores, or dimension values are shown to any user |
| **Stateless generation** | Summaries are regenerated from numeric state only; no paragraph memory |
| **Optional collection** | Personality signals are entirely opt-in; skipping has no penalty |
| **No filtering or ranking** | Summaries do not influence discovery order or matching algorithms |
| **Third-person only** | All summaries use observational "they/their" language |

> [!IMPORTANT]
> Personality Intelligence exists to add depth to discovery, not to categorize, score, or rank users. It complements hobby-based discovery without replacing it.

---

## 20. Metrics & Success Signals

### 20.1 Core Success Metrics

| Metric                     | Description                                    | Target (MVP)                                    |
| -------------------------- | ---------------------------------------------- | ----------------------------------------------- |
| **D7 Retention**           | Users active 7 days after registration         | Track, no target yet                            |
| **Streak Completion Rate** | % of matches that reach Day 15                 | >15% (indicates right users finding each other) |
| **Reset-to-Restart Ratio** | % of resets where users choose to restart      | >30% (indicates emotional investment)           |
| **Reveal Engagement**      | User actions taken within 24h of reveal        | Track engagement velocity                       |
| **Recovery Payment Rate**  | % of at-risk streaks where payment is made     | Track without optimizing                        |
| **Chat Initiation Rate**   | % of users who send first message after unlock | >60% (indicates streak built genuine interest)  |

### 20.2 What NOT to Optimize Early

| Metric                            | Why Not to Optimize                             |
| --------------------------------- | ----------------------------------------------- |
| **Match volume**                  | More matches ≠ better outcomes; quality matters |
| **Time in app**                   | Engagement should be purposeful, not addictive  |
| **DAU/MAU ratio**                 | Users should return to continue, not browse     |
| **Revenue per user (short-term)** | Trust-building matters more than extraction     |
| **Conversion from Free to Paid**  | Conversion should be pull-based, not push-based |

### 20.3 Leading Indicators of Product-Market Fit

1. **Organic word-of-mouth** — Users referring friends without incentive
2. **High streak completion among completers** — The users who complete are deeply satisfied
3. **Re-engagement after break** — Users who step away return voluntarily
4. **Low support volume on core mechanics** — Users understand the system intuitively
5. **Qualitative feedback themes** — Users describe it as "different," "refreshing," "worth it"

---

## 21. Edge Cases & Failure Modes

### 21.1 One-Sided Effort

**Scenario:** User A checks in every day; User B frequently misses but recovers with payment.

**System behavior:**

- AI detects pattern through streak score
- Reveals may slow down for this connection
- Choice window after each reset gives User A opportunity to exit without guilt
- System does not force continuation

**Design principle:** Behavior resolves intent over time. The system does not judge, but it creates natural exit points.

### 21.2 Repeated Breakers

**Scenario:** A user repeatedly breaks streaks across multiple connections.

**System behavior:**

- AI flags pattern
- No direct punishment visible to user
- May affect matching priority (reduced visibility score)
- Extreme cases trigger human review

**Design principle:** We let behavior decide, not rules. Repeat offenders self-select out through natural consequences.

### 21.3 Free Users Losing Connections

**Scenario:** Free user's one connection ends; they must start over with discovery limits.

**System behavior:**

- Immediate return to discovery
- Clear next action always available ("Find a new connection")
- No messaging that makes them feel "stuck"
- Monetization touchpoint (Plus offers 2 connections)

**Design principle:** Every state has a next step. We never leave users stranded.

### 21.4 Paid User / Free User Dynamics (Locked: Strict Tier Isolation)

> [!CAUTION]
> All features, speeds, limits, and entitlements are strictly bound to the user's own subscription tier. No user ever receives features, pacing, or benefits from another user's tier.

**Scenario:** Pro user matched with Free user. Both progress through the same streak timeline.

---

#### Core Rules

| Shared                                 | Per-User (Tier-Specific)   |
| -------------------------------------- | -------------------------- |
| Streak progress                        | Reveal unlock timing       |
| Streak states (Active, At Risk, Reset) | Paid reveal eligibility    |
| Day count toward milestones            | Streak recovery allowances |
| Chat unlock milestone (Day 15)         | Nudge frequency and limits |
| Identity reveal milestone (Day 15)     | Discovery refresh rates    |

- **Streak progress is shared.** Both users advance together; the day count is identical for both.
- **Reveal milestones are shared.** When Day 5, Day 10, etc. is reached, both users hit the same milestone.
- **Feature access is always tier-specific.** Each user unlocks only what their tier permits.

---

#### System Behavior

When a reveal milestone is reached, the system evaluates each user independently:

| User Tier | Behavior                                                                                   |
| --------- | ------------------------------------------------------------------------------------------ |
| **Free**  | Unlocks only Free-tier reveals (Day 5, Day 10). Remaining reveals require purchase.        |
| **Plus**  | Unlocks only Plus-tier reveals (Day 4, Day 8, Day 12). Remaining reveals require purchase. |
| **Pro**   | Unlocks all Pro-tier reveals (Day 3, Day 6, Day 9, Day 12). May purchase remaining if any. |

**No cross-tier benefit occurs:**

- A Free user matched with a Pro user does **not** receive faster reveals.
- A Plus user matched with a Pro user does **not** receive Pro reveal timing.
- Each user's reveal cadence, paid reveal eligibility, streak recovery allowances, and nudge limits are evaluated per user, not per connection.

---

#### Chat & Identity Rules

- Chat unlock occurs only after full streak completion (Day 15).
- Full identity reveal occurs only at Day 15.
- **No tier accelerates identity or chat unlock.** This milestone is universal and non-bypassable.
- Both Free and Pro users reach identity and chat at the same moment — the 15-day mark.

---

#### What This Looks Like in Practice

**Example: Pro user + Free user connection**

| Day    | Pro User Experience    | Free User Experience           |
| ------ | ---------------------- | ------------------------------ |
| Day 3  | Earns Reveal 1 (auto)  | No reveal unlocked             |
| Day 5  | —                      | Earns Reveal 1 (auto)          |
| Day 6  | Earns Reveal 2 (auto)  | No reveal unlocked             |
| Day 9  | Earns Reveal 3 (auto)  | No reveal unlocked             |
| Day 10 | —                      | Earns Reveal 2 (auto)          |
| Day 12 | Earns Reveal 4 (auto)  | No earned reveal; may purchase |
| Day 15 | Identity + Chat unlock | Identity + Chat unlock         |

Both users reach Day 15 identity reveal simultaneously. The Pro user has accumulated more reveals along the way; the Free user has fewer earned reveals but may purchase additional ones.

---

#### Design Principles

| Principle                                | Implication                                                        |
| ---------------------------------------- | ------------------------------------------------------------------ |
| **Connections are shared experiences**   | Streak progress and milestones are mutual                          |
| **Benefits are personal entitlements**   | Feature access is strictly tier-bound                              |
| **No tier leakage**                      | Higher tiers do not "spill over" to partners                       |
| **No cross-tier subsidies**              | Free users never receive Plus/Pro advantages indirectly            |
| **No indirect upgrades through pairing** | Matching with a Pro user does not upgrade a Free user's experience |

---

#### What This Explicitly Prevents

- ❌ A higher-tier user benefiting their partner's reveal timing
- ❌ Shared reveal cadence based on the higher tier
- ❌ Free or Plus users receiving premium advantages through their partner
- ❌ Any scenario where pairing affects feature access

### 21.5 Churn Moments

**High-risk churn moments:**

- First streak reset (emotional deflation)
- Day 8–12 plateau (anticipation fatigue)
- After identity reveal if expectations mismatch
- After chat opens if conversation doesn't flow

**Mitigation strategies:**

- Reset is framed as pause + choice, not failure
- Micro-milestones (3 days, first recovery, etc.) punctuate waiting
- Expectation management through gradual reveals (less mismatch at Day 15)
- Chat is positioned as continuation, not new beginning

### 21.6 Progressive Verification Edge Cases

**Scenario:** User completes onboarding but skips all optional verification.

**System behavior:**

- User proceeds to Discovery normally
- Liveness challenge offered during first streak (Day 5-7)
- No penalty for skipping
- Subtle periodic prompts in settings
- Trust builds through behavioral signals instead

**Design principle:** Optional means optional. Behavioral trust is earned regardless of optional verification completion.

---

**Scenario:** User's uploaded photos fail liveness matching.

**System behavior:**

- Retry available immediately
- Failure not visible to matches
- After 3 failures: suggest re-uploading photos
- Support escalation available
- No account penalty

**Design principle:** Liveness is for photo authenticity verification, not gatekeeping.

---

**Scenario:** User wants government ID verification but hasn't met criteria.

**System behavior:**

- Show clear unlock criteria (streak completion OR 30 days active)
- Encourage streak completion as fastest path
- No exception requests processed

**Design principle:** Gov ID is earned through presence, maintaining the progressive philosophy.

---

## 22. Roadmap

### Phase 1: MVP (Months 1–4)

**Objective:** Validate core behavioral loop with real users in Bangalore

**Scope:**

- Core onboarding flow with phone + mandatory photo upload
- Progressive verification (liveness challenge, behavioral trust)
- Optional trust boosters (Google/Apple/LinkedIn)
- Single intent server (romantic connection only)
- Anonymous discovery with hobby-anchored cards
- All filters available to all users
- Mutual interest → streak initiation
- Full streak mechanics (check-in, at-risk, recovery, reset)
- Nudge system (all tiers)
- 5-reveal structure with tier-based timing
- Day-15 identity + chat unlock
- Free / Plus / Pro tier structure
- Credit system for payment protection
- Basic AI matching (compatibility scoring)
- Basic safety monitoring
- Trust & Verification settings screen

**Success criteria:**

- 500+ verified users
- 50+ completed streaks (Day 15)
- Streak completion rate >10%
- Qualitative feedback indicating differentiated value

---

### Phase 2: AI Depth (Months 5–8)

**Objective:** Improve reveal experience and safety through AI investment

**Scope:**

- Dynamic reveal timing modulation based on streak health
- Improved nudge selection (context-aware, pre-written variants)
- Enhanced behavioral risk detection (manipulation patterns)
- AI-framed reveals (emotional language, not raw data)
- Unified trust & safety monitoring (proactive intervention opportunities)

**Success criteria:**

- Improved streak completion rate (>15%)
- Reduced time-to-first-mutual-interest
- Lower safety incident rate
- Higher post-chat engagement scores

---

### Phase 3: Expansion (Months 9–12+)

**Objective:** Expand beyond initial market and romantic outcomes

**Scope:**

- Geographic expansion (other metro cities)
- Additional intent servers (friend, accountability buddy)
- Age-specific cohorts (college-specific, 35+)
- Advanced monetization features (founding member, premium nudges)
- Group streaks (exploratory: accountability groups)
- Integration with external verification (international markets)
- Platform maturity (moderation tools, appeals process)

**Success criteria:**

- Multi-city presence
- Non-romantic connection success stories
- Sustainable unit economics
- Clear path to profitability

---

## 23. Technical Stack (Recommended)

### 23.1 Frontend

- **Framework:** React Native (cross-platform iOS/Android)
- **State management:** State-driven UI (streak states must be explicit and predictable)
- **Offline handling:** Check-in should queue if offline, sync when available

### 23.2 Backend

- **Framework:** Node.js with NestJS (structured, modular)
- **Architecture:** Event-driven (streak events, nudge events, reveal events as first-class concepts)
- **API:** RESTful with consideration for GraphQL for complex queries

### 23.3 Database

- **Primary:** PostgreSQL (core relational data: users, connections, profiles)
- **Cache/timers:** Redis (streak timers, recovery windows, session management)
- **Search:** Elasticsearch (if discovery volume requires)

### 23.4 AI Layer

- **Feature extraction:** Custom services for compatibility scoring
- **LLM usage:** Rule-constrained (no free-text autonomy)
- **Model hosting:** Cloud-based inference with local fallback for latency-sensitive operations

### 23.5 Security

- **Data at rest:** AES-256 encryption for all PII
- **Data in transit:** TLS 1.3 minimum
- **Identity verification:** Tokenized storage (no raw documents retained)
- **Access control:** Role-based with audit logging
- **Compliance:** Prepared for India's data protection requirements

### 23.6 Infrastructure

- **Cloud provider:** AWS (primary) with consideration for regional requirements
- **Deployment:** Container-based (ECS/EKS)
- **Monitoring:** Comprehensive logging with streak-state visibility
- **Backup:** Daily backups with point-in-time recovery

---

## 24. Appendix: Investor Q&A Reference

### "Do you really think people will patiently maintain a streak?"

**Short answer:** No — most people won't. And that's the point.

**Full answer:** We intentionally design for intent, not volume. Friction filters out low-effort users early and improves experience and retention for those who stay. Streaks already work in products like Duolingo and fitness apps; here the emotional stakes are higher because another human is involved.

**One-liner:** "We don't optimize for activity — we optimize for outcome."

---

### "What happens if someone breaks the streak?"

**Short answer:** A missed day puts the streak at risk, not broken.

**Full answer:** If one person misses, there's a recovery window. If both miss, the streak resets. A reset doesn't end the match — it triggers a mutual decision to restart or move on within a fixed time window.

**One-liner:** "Reset ≠ rejection. Reset creates a pause and a choice."

---

### "Won't this create frustration and churn, especially for free users?"

**Short answer:** Only if users are left idle — which we don't do.

**Full answer:** Every reset immediately presents a clear next action: restart or move on. Free users are never stranded. There's always forward motion, which preserves dignity and reduces churn.

**One-liner:** "We never leave users stuck — every state has a next step."

---

### "What if one person keeps missing repeatedly?"

**Short answer:** The system doesn't judge — behavior resolves it naturally.

**Full answer:** Each reset triggers a mutual choice. Over repeated breaks, one user will usually opt out. We don't force breakups or impose hard limits; human intent self-selects.

**One-liner:** "We let behavior decide, not rules."

---

### "How do users motivate each other without chat?"

**Short answer:** Through bounded, non-intrusive signals — not messages.

**Full answer:** We use one-tap nudges that signal presence ("I showed up") without pressure or guilt. This preserves safety, attraction, and anonymity while enabling accountability.

**One-liner:** "We don't let users chase each other — we let them signal commitment."

---

### "Is nudging really enough?"

**Short answer:** Yes — anything more becomes pressure.

**Full answer:** Free-text reminders, multiple alerts, or escalation create imbalance and resentment. Nudges work because they are rare, voluntary, and respectful. The rest of the motivation comes from system design: progress visibility, shared milestones, and anticipation.

**One-liner:** "Bounded tools create trust; unlimited tools create pressure."

---

### "Why won't this just feel slow compared to dating apps?"

**Short answer:** Because slowness is the feature, not the bug.

**Full answer:** Existing apps optimize for speed and attention, which leads to distraction and low-value interactions. We optimize for focus and progression. Users who want instant gratification already have infinite options — we're unlocking a different audience.

**One-liner:** "Instant access feels cheap; earned access feels valuable."

---

### "Isn't this just another dating app?"

**Short answer:** No — it's a connection platform with delayed communication.

**Full answer:** Dating is an outcome, not the category. We don't optimize for matches or chat volume; we optimize for consistency and intent. Many users form romantic connections, but that's a result, not the hook.

**One-liner:** "We don't sell matches — we sell commitment."

---

### "What about women's safety and stigma?"

**Short answer:** Safety is embedded structurally, not marketed loudly.

**Full answer:** Delayed chat, earned identity reveals, bounded actions, and mutual consent dramatically reduce harassment and pressure. This design naturally makes the platform safer and more comfortable, especially in conservative contexts — without positioning it as a "women-only" product.

**One-liner:** "We design for safety through structure, not slogans."

---

### "How does monetization fit without breaking trust?"

**Short answer:** We monetize time and flexibility — never identity.

**Full answer:** Users can pay to protect streaks, reduce waiting time, or maintain multiple connections. We never sell unilateral access to another person. All monetization preserves mutual consent.

**One-liner:** "We sell control over time, not control over people."

---

### "Why now? What's changed?"

**Short answer:** People are exhausted by attention-driven apps.

**Full answer:** There's growing fatigue with distraction, endless swiping, and shallow interaction. At the same time, users are more comfortable with progress-based systems (streaks, habits) and delayed rewards. AI allows us to pace and moderate this responsibly at scale.

**One-liner:** "We're building for the post-swipe, attention-fatigued user."

---

### "Why will this retain users long-term?"

**Short answer:** Because progress creates attachment.

**Full answer:** Retention comes from visible progression, shared milestones, anticipation, and emotional investment — not infinite choice. Users don't come back to browse; they come back to continue.

**One-liner:** "People don't abandon things they've invested effort into."

---

### "What if this only works for a small group?"

**Short answer:** That's acceptable — and intentional.

**Full answer:** We don't need everyone. A smaller, high-intent user base with strong retention and monetization is more sustainable than mass low-intent usage.

**One-liner:** "Depth beats breadth in connection products."

---

## Document End

**Last updated:** January 2026  
**Document owner:** Product Leadership  
**Distribution:** Internal, Investors, Founding Team

---

_Friction is not a tax. It is the filter. Unora is intentionally not for everyone. That is what makes it defensible._
