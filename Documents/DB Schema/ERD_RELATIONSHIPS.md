# ERD_RELATIONSHIPS.md

**Version:** 1.0  
**Date:** January 2026  
**Database:** PostgreSQL 15+  
**Notation:** `1───*` (One-to-Many), `*───*` (Many-to-Many via join table)

---

## Entity Relationship Diagram

### Notation Key

```
1───* : One-to-Many        (Parent has many Children)
*───1 : Many-to-One        (Child belongs to Parent)
*───* : Many-to-Many       (via join table)
1───1 : One-to-One
○───○ : Optional reference (nullable FK)
●───● : Required reference (NOT NULL FK)
```

---

## Core Identity & Profile

```
users
│
├──1───*── profiles                    (1 user has 1 profile, but profile references user)
│           │
│           └──1───*── profile_edit_history
│
├──1───*── photos                      (1 user has 3-6 photos)
│           │
│           └──*───1── reveals         (photo revealed at Day 15)
│
├──1───*── hobbies                     (1 user has 2+ hobbies)
│           │
│           └──*───1── hobby_options   (reference to master hobby list)
│
├──1───*── verification_signals        (1 user has 0+ trust signals)
│
└──1───1── subscriptions               (1 user has exactly 1 subscription record)
```

---

## Discovery System

```
servers (reference table: partner, friend, growth)
│
├──*───*── users (via discovery_batches)
│
└──1───*── discovery_batches
            │
            ├──*───1── users                  (batch belongs to user)
            │
            ├──1───*── discovery_cards        (1 batch has exactly 5 cards)
            │           │
            │           ├──*───1── users     (card represents a candidate user)
            │           │
            │           └──○───1── interests (card may become interest)
            │
            └──*───1── filters               (batch uses applied filters)

filters
│
├──*───1── users                            (filter belongs to user)
│
└──*───1── servers                          (filter is server-scoped)
```

---

## Interest & Matching

```
interests
│
├──●───1── users (as sender_user_id)         (interest has sender)
│
├──●───1── users (as receiver_user_id)       (interest has receiver)
│
├──●───1── servers                           (interest is server-scoped)
│
├──○───1── discovery_cards                   (interest originated from card)
│
└──○───1── connections                       (if matched, links to connection)

connections
│
├──●───1── users (as user_a_id)              (connection has user A, lower ID)
│
├──●───1── users (as user_b_id)              (connection has user B, higher ID)
│
├──●───1── servers                           (connection is server-bound)
│
├──1───*── streaks                           (1 connection has 0+ streaks over time)
│
└──1───*── interest_wipes                    (audit: total wipe events)

interest_wipes
│
├──●───1── users (as wiped_user_id)          (user who triggered wipe)
│
├──●───1── connections (as triggering_connection_id)
│
└──1───*── wiped_interests                  (list of interests that were wiped)
```

---

## Streak Engine

```
streaks
│
├──●───1── connections                       (streak belongs to connection)
│
├──○───1── users (as breaker_user_id)        (NULL when active, set when one misses)
│
├──1───*── check_ins                         (1 streak has many check-ins)
│
├──1───*── nudges                            (1 streak has many nudges)
│
├──1───*── reveals                           (1 streak has up to 5 reveals per user)
│           │
│           └──*───1── users                 (reveal is per-user within streak)
│
└──○───1── recovery_payments                 (if recovered, links to payment)

check_ins
│
├──●───1── streaks                           (check-in belongs to streak)
│
├──●───1── users                             (check-in by specific user)
│
└── UNIQUE (streak_id, user_id, check_in_date)   -- No double check-ins

nudges
│
├──●───1── streaks                           (nudge for specific streak)
│
├──●───1── users (as sender_user_id)         (nudge sent by maintaining user)
│
├──●───1── users (as receiver_user_id)       (nudge received by missing user)
│
└── COUNT limited by tier per at-risk period
```

---

## Reveals

```
reveal_milestones (configuration table)
│
├── tier (free/plus/pro)
│
├── reveal_number (1-5)
│
├── reveal_day_required (internal, never shown to users)
│
└── reveal_content_schema_version

reveals
│
├──●───1── streaks                           (reveal within streak context)
│
├──●───1── users                             (reveal for specific user)
│
├──●───1── reveal_milestones                 (which milestone this is)
│
├──○───1── payments                          (if purchased, links to payment)
│
└──1───1── reveal_contents                   (AI-framed content if unlocked)

reveal_contents
│
├──●───1── reveals                           (content for specific reveal)
│
└── content_json (JSONB: AI-framed, versioned)
```

---

## Monetization & Credits

```
subscriptions
│
├──1───1── users                             (exactly 1 subscription per user)
│
├── tier (free/plus/pro)
│
├── billing_cycle_start
│
└── billing_cycle_end

payments
│
├──●───1── users                             (payment by user)
│
├──○───1── streaks (for recovery payments)
│
├──○───1── reveals (for reveal purchases)
│
├──○───1── nudges (for instant nudge purchases)
│
└── payment_type ENUM: 'subscription', 'recovery', 'reveal', 'nudge'

credits
│
├──●───1── users                             (credit balance for user)
│
└──1───*── credit_transactions               (credit history)

credit_transactions
│
├──●───1── credits                           (transaction on credit account)
│
├──○───1── payments                          (if conversion from payment)
│
├──○───1── credit_protections                (if triggered by protection window)
│
└── transaction_type ENUM: 'conversion', 'spend', 'refund', 'adjustment'

credit_protections
│
├──●───1── payments                          (protection window for payment)
│
├──●───1── connections                       (monitored connection)
│
├── window_starts_at
│
├── window_ends_at
│
└──○───1── credit_transactions               (if conversion triggered)
```

---

## Audit & Safety

```
audit_logs (append-only)
│
├──●───1── users (as actor_user_id)          (who performed action)
│
├──○───*── (polymorphic reference to any entity)
│
├── event_type ENUM: 'check_in', 'interest', 'match', 'recovery', 'terminate', etc.
│
└── event_data (JSONB: full context)

safety_flags
│
├──●───1── users (as flagged_user_id)        (user with flag)
│
├── flag_type ENUM: 'fake_consistency', 'manipulation', 'behavioral_shift'
│
├── severity (1-10)
│
└── reviewed_at (NULL if pending review)

blocks
│
├──●───1── users (as blocker_user_id)
│
├──●───1── users (as blocked_user_id)
│
└── UNIQUE (blocker_user_id, blocked_user_id)

reports
│
├──●───1── users (as reporter_user_id)
│
├──●───1── users (as reported_user_id)
│
├──○───1── connections                       (optional context)
│
├──○───1── streaks                           (optional context)
│
├── report_type ENUM: 'harassment', 'fake', 'spam', 'other'
│
└── status ENUM: 'pending', 'reviewed', 'actioned', 'dismissed'
```

---

## Personality Intelligence

```
users
│
├──1───1── personality_scores       (1 user has 0..1 score record, optional)
│
└──1───*── personality_answers      (1 user has 0+ answers, optional)
            │
            └──*───1── personality_questions  (answer references question)

personality_questions (reference table)
├── id: UUID
├── dimension: VARCHAR (e.g., 'openness', 'conscientiousness')
├── question_text: TEXT
├── options: JSONB (array of {id, text, weight})
├── version: INTEGER
└── is_active: BOOLEAN

personality_scores
├── id: UUID
├── user_id: UUID (UNIQUE, 1:1 with users)
├── dimension_scores: JSONB (numeric scores + confidence ONLY)
├── last_asked_dimension: VARCHAR (for rotation)
└── NO SUMMARIES OR LABELS STORED

personality_answers
├── id: UUID
├── user_id: UUID
├── question_id: UUID
├── selected_option_id: VARCHAR
├── context: 'onboarding' | 'streak_checkin'
└── streak_id: UUID (optional, if during streak)
```

**Critical Constraints:**

| Constraint | Enforcement |
|------------|-------------|
| Optional participation | No FK requirement, nullable score record |
| One answer per question | UNIQUE (user_id, question_id) |
| No summaries stored | Schema has no text generation fields |
| No labels or rankings | Only numeric scores + confidence |

---

## Join Tables (Many-to-Many)

### users ←→ hobby_options

```
hobbies (join table)
│
├──●───1── users
│
├──●───1── hobby_options
│
├── micro_description (user's selected description variant)
│
└── display_order
```

---

## Non-Obvious Constraints

### One Active Streak Per Connection
```
UNIQUE INDEX ON streaks (connection_id)
WHERE streak_state NOT IN ('reset', 'completed', 'terminated')
  AND deleted_at IS NULL
```

### One Active Batch Per User Per Server
```
UNIQUE INDEX ON discovery_batches (user_id, server_type)
WHERE batch_status = 'active' AND deleted_at IS NULL
```

### No Duplicate Interests Same Pair
```
UNIQUE INDEX ON interests (sender_user_id, receiver_user_id, server_type)
WHERE interest_status = 'pending' AND deleted_at IS NULL
```

### Connection User Ordering (A < B)
```
CHECK (user_a_id < user_b_id)
-- This ensures no duplicate pairs (A,B) vs (B,A)
```

### Photo Count Constraints
```
-- Enforced at application layer with validation
-- Minimum 3 photos before Discovery access
-- Maximum 6 photos per user
```

---

## Reference Tables (Static Data)

```
servers
├── id: UUID
├── server_type: ENUM ('partner', 'friend', 'growth')
├── display_name: VARCHAR
├── icon_name: VARCHAR (Phosphor icon reference)
└── sort_order: INTEGER

hobby_options
├── id: UUID
├── name: VARCHAR
├── category: VARCHAR
├── icon_name: VARCHAR
├── micro_descriptions: JSONB (array of pre-written options)
└── is_active: BOOLEAN

reveal_milestones
├── id: UUID
├── tier: ENUM ('free', 'plus', 'pro')
├── reveal_number: INTEGER (1-5)
├── reveal_day_required: INTEGER (internal logic, never shown)
├── reveal_type: ENUM ('personality', 'lifestyle', 'social', 'human', 'identity')
└── content_schema_version: INTEGER
```

---

*This document defines all entity relationships. For exact DDL, see TABLE_DEFINITIONS.md. For index strategy, see INDEXES_PERFORMANCE.md.*
