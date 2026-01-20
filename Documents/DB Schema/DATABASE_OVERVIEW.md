# DATABASE_OVERVIEW.md

**Version:** 1.0  
**Date:** January 2026  
**Database:** PostgreSQL 15+  
**Classification:** Internal / Engineering

---

## Quick Navigation

| Module | Purpose | Key Tables |
|--------|---------|------------|
| [Identity & Verification](#module-1-identity--verification) | User accounts, progressive trust | `users`, `verification_signals` |
| [Profile & Hobbies](#module-2-profile--hobbies) | User-provided data, photos | `profiles`, `photos`, `hobbies` |
| [Discovery & Caching](#module-3-discovery--caching) | Server-scoped batches, cards | `discovery_batches`, `discovery_cards` |
| [Interest & Matching](#module-4-interest--matching) | Mutual interest, connection formation | `interests`, `connections` |
| [Streak Engine](#module-5-streak-engine) | Daily check-ins, state machine | `streaks`, `check_ins`, `nudges` |
| [Reveals](#module-6-reveals) | Milestone unlocks, content delivery | `reveals`, `reveal_milestones` |
| [Monetization & Credits](#module-7-monetization--credits) | Subscriptions, payments, protection | `subscriptions`, `payments`, `credits` |
| [Audit & Safety](#module-8-audit--safety) | Logging, safety flags | `audit_logs`, `safety_flags` |
| [Personality Intelligence](#module-9-personality-intelligence) | Optional signals, numeric scores | `personality_scores`, `personality_answers` |

---

## System Mental Model

```
┌─────────────────────────────────────────────────────────────────────────────────┐
│                              USER JOURNEY FLOW                                   │
├─────────────────────────────────────────────────────────────────────────────────┤
│                                                                                 │
│   ONBOARDING          DISCOVERY           MATCHING           STREAK             │
│   ──────────          ─────────           ────────           ──────             │
│                                                                                 │
│   ┌────────┐         ┌──────────┐        ┌──────────┐       ┌──────────┐        │
│   │ users  │────────▶│ discovery│───────▶│ interests│──────▶│ streaks  │        │
│   │profiles│         │ _batches │        │          │       │ check_ins│        │
│   │ photos │         │ _cards   │        │connections│      │ reveals  │        │
│   └────────┘         └──────────┘        └──────────┘       └──────────┘        │
│       │                   │                   │                  │              │
│       │              SERVER-SCOPED        SERVER-SCOPED     CONNECTION-SCOPED   │
│       │                                                                         │
│   GLOBAL INVENTORY:                                                             │
│   ├── active_connection_count (spans all servers)                               │
│   └── last_global_refresh_at (single timer for all servers)                     │
│                                                                                 │
└─────────────────────────────────────────────────────────────────────────────────┘
```

---

## Module 1: Identity & Verification

### Purpose
Manages user accounts, authentication, and the progressive verification model. Trust is built through behavior (phone → photos → photo quality → streaks → optional Gov ID).

### Key Tables

| Table | Description |
|-------|-------------|
| `users` | Core account record. Contains auth, tier, global inventory counters |
| `verification_signals` | OAuth links, Gov ID tokens, trust boosters |

### Ownership Boundaries
- `users` owns all verification records
- Verification signals are user-scoped, never shared with connections

### What This Module Guarantees
- ✓ Phone number uniqueness
- ✓ Minimum 3 photos before Discovery access
- ✓ AI photo quality verification (passive, automatic)
- ✓ Gov ID tokens stored, raw documents never retained

---

## Module 2: Profile & Hobbies

### Purpose
Stores user-provided profile data and photos. All photos are private until Day 15 reveal. Hobbies are the primary discovery anchor.

### Key Tables

| Table | Description |
|-------|-------------|
| `profiles` | Name, DOB, gender, city, bio, optional fields |
| `photos` | 3-6 photos per user, private by default |
| `hobbies` | User's selected hobbies with micro-descriptions |
| `hobby_options` | Master list of available hobbies |

### Ownership Boundaries
- Each user has exactly one `profiles` record
- Photos belong to user, linked to reveals for Day 15
- Hobbies are user-selected from `hobby_options`

### What This Module Guarantees
- ✓ Minimum 2 hobbies required before Discovery
- ✓ At least one face photo flagged during upload
- ✓ Profile data editable until Gov ID locks fields
- ✓ Photos never exposed until reveal milestone

---

## Module 3: Discovery & Caching

### Purpose
Manages the server-scoped discovery experience. Each server maintains independent cached card batches. Global refresh timer controls availability.

### Key Tables

| Table | Description |
|-------|-------------|
| `servers` | Reference table for server types: partner, friend, growth |
| `discovery_batches` | User's batch per server, tracks refresh state |
| `discovery_cards` | Individual cards within a batch (exactly 5 per batch) |
| `filters` | User's saved filters per server |

### Ownership Boundaries
- One active batch per user per server (enforced by unique constraint)
- Cards are generated from candidate pool, cached until next refresh
- Filters are server-scoped but stored globally for the user

### What This Module Guarantees
- ✓ Exactly 5 cards per batch (tier-independent)
- ✓ Cards cached until explicit refresh action
- ✓ Refresh timer is GLOBAL (not per-server)
- ✓ Filters applied only at refresh, not real-time

### Read Path: Discovery Screen
```sql
-- 1. Check if user can access Discovery (capacity lock)
SELECT active_connection_count, tier FROM users WHERE id = $user_id;
-- Compare against tier limits: Free=1, Plus=2, Pro=4

-- 2. Check if refresh is available (time lock)
SELECT refresh_available_at FROM users WHERE id = $user_id;
-- Compare against NOW()

-- 3. Get current batch for server
SELECT * FROM discovery_batches
WHERE user_id = $user_id AND server_type = $server AND deleted_at IS NULL;

-- 4. Get cards for batch
SELECT dc.* FROM discovery_cards dc
WHERE dc.batch_id = $batch_id AND dc.deleted_at IS NULL;
```

---

## Module 4: Interest & Matching

### Purpose
Handles mutual interest detection and connection formation. Implements "Total Wipe" rule when user reaches capacity.

### Key Tables

| Table | Description |
|-------|-------------|
| `interests` | Unilateral interest expression (pending until mutual) |
| `connections` | Formed when mutual interest detected, server-bound |
| `interest_wipes` | Audit log for Total Wipe executions |

### Ownership Boundaries
- Interests are directional: sender → receiver
- Connections require both parties (user_a, user_b, ordered by ID)
- Total Wipe affects only the user who triggered capacity fill

### What This Module Guarantees
- ✓ Mutual interest required for connection
- ✓ First-to-Lock protocol (server timestamp ordering)
- ✓ Total Wipe executed when final slot fills
- ✓ Wiped interests never shown to recipients

### Matching Logic Flow
```
Interest sent → Stored pending
                    │
                    ▼
            Recipient also interested?
                    │
           ┌────────┴────────┐
           │                 │
          NO                YES
           │                 │
     Stays pending      Check sender slots
                             │
                    ┌────────┴────────┐
                    │                 │
                 OPEN              FULL
                    │                 │
           Create connection      Match fails
           Check final slot?     (Dynamic invisibility)
                    │
           If final → Total Wipe
```

---

## Module 5: Streak Engine

### Purpose
The core behavioral loop. Manages daily check-ins, state transitions, recovery windows, and nudges. Streaks are first-class entities with explicit state machines.

### Key Tables

| Table | Description |
|-------|-------------|
| `streaks` | State machine for each connection's streak |
| `check_ins` | Daily check-in records (immutable) |
| `nudges` | Nudge history with tier limits |
| `recovery_payments` | Links streak recovery to payment |

### Ownership Boundaries
- One active streak per connection at any time
- Check-ins belong to streak, never deleted
- Nudges are per at-risk period, counted per tier

### What This Module Guarantees
- ✓ Only breaker can make recovery payment
- ✓ No late check-ins after day ends
- ✓ Streak continues from previous count on recovery
- ✓ Reset does NOT terminate connection
- ✓ Day 15 completion unlocks identity reveal + chat

### Streak States

| State | Entry Condition | Exit Conditions |
|-------|-----------------|-----------------|
| `active` | Both check in | One misses → `at_risk`, Both miss → `reset` |
| `at_risk` | One user missed (Day N) | End of Day N → `payment_window` |
| `payment_window` | Day N+1 begins | Payment → `active`, Deadline passes → `reset` |
| `reset` | No recovery payment | New check-in cycle begins |
| `completed` | Day 15 mutual check-in | Terminal state |
| `terminated` | User explicitly ends | Terminal state |

### Read Path: Streaks Dashboard
```sql
-- Get ALL active connections across ALL servers (global aggregator)
SELECT
  c.id, c.server_type,
  s.streak_state, s.current_day, s.breaker_user_id,
  partner.id as partner_id
FROM connections c
JOIN streaks s ON s.connection_id = c.id AND s.deleted_at IS NULL
JOIN users partner ON partner.id = CASE
  WHEN c.user_a_id = $user_id THEN c.user_b_id
  ELSE c.user_a_id
END
WHERE (c.user_a_id = $user_id OR c.user_b_id = $user_id)
  AND c.connection_status = 'active'
  AND c.deleted_at IS NULL;
```

---

## Module 6: Reveals

### Purpose
Manages the gradual information disclosure system. Reveals are tier-specific (earned vs purchased) with Day 15 identity reveal universal.

### Key Tables

| Table | Description |
|-------|-------------|
| `reveal_milestones` | Configuration: which reveals unlock at which days per tier |
| `reveals` | Per-user-per-streak reveal state (locked/earned/purchased/unlocked) |
| `reveal_contents` | JSONB storage for AI-framed reveal content |

### Ownership Boundaries
- Reveal milestones are global configuration
- Each user has independent reveal progression (strict tier isolation)
- Content is personalized per user (AI-framed)

### What This Module Guarantees
- ✓ Tier isolation: Pro user's reveals don't affect Free partner
- ✓ Day 15 identity cannot be accelerated by any tier or payment
- ✓ Maximum 5 reveals total (earned + purchased)
- ✓ Reveal content is never raw data dump (AI-framed)

### Read Path: Check Reveal Status
```sql
SELECT
  r.reveal_number,
  r.reveal_status,
  r.unlocked_at,
  rm.reveal_day_required, -- Internal: user never sees this
  CASE
    WHEN r.reveal_status = 'unlocked' THEN rc.content_json
    ELSE NULL
  END as content
FROM reveals r
JOIN reveal_milestones rm ON rm.reveal_number = r.reveal_number
  AND rm.tier = $user_tier
LEFT JOIN reveal_contents rc ON rc.reveal_id = r.id
WHERE r.streak_id = $streak_id AND r.user_id = $user_id;
```

---

## Module 7: Monetization & Credits

### Purpose
Handles subscriptions, one-time payments, and credit protection. Implements 24-hour credit conversion for bad-faith terminations.

### Key Tables

| Table | Description |
|-------|-------------|
| `subscriptions` | User tier: free/plus/pro, billing dates |
| `payments` | All payment transactions (recovery, reveals, nudges) |
| `credits` | In-app credit balance and transactions |
| `credit_protections` | 24-hour windows for bad-faith detection |

### Ownership Boundaries
- One active subscription per user
- Payments are immutable transaction records
- Credits are closed-loop currency (no withdrawal)

### What This Module Guarantees
- ✓ Credits auto-convert on bad-faith termination within 24h
- ✓ Recovery payments linked to specific streak
- ✓ Free recovery allowances tracked per connection
- ✓ Credits cannot pay for subscriptions

### Credit Protection Logic
```
Recovery payment made → 24-hour protection window opens
                             │
                             ▼
                   Partner terminates?
                             │
                    ┌────────┴────────┐
                    │                 │
              WITHIN 24h         AFTER 24h
                    │                 │
              Auto-convert        No conversion
              to credits          Payment consumed
```

---

## Module 8: Audit & Safety

### Purpose
Logging for compliance, behavior pattern detection, and silent safety interventions.

### Key Tables

| Table | Description |
|-------|-------------|
| `audit_logs` | Append-only log of significant events |
| `safety_flags` | AI-generated flags for suspicious patterns |
| `blocks` | User-to-user blocks |
| `reports` | User-submitted safety reports |

### Ownership Boundaries
- Audit logs are system-owned, immutable
- Safety flags are internal (never shown to users)
- Blocks and reports are user-initiated

### What This Module Guarantees
- ✓ All state transitions logged
- ✓ Payment events have full audit trail
- ✓ Safety interventions are silent (no public shaming)
- ✓ Flagged accounts can have visibility reduced

---

## Module 9: Personality Intelligence

### Purpose
Optional, gradual personality signal collection. Stores ONLY numeric scores and confidence—no labels, no summaries, no generated text. AI generates summaries ON DEMAND from these scores.

### Key Tables

| Table | Description |
|-------|-------------|
| `personality_questions` | Static, versioned questions (system-defined, never AI-generated) |
| `personality_answers` | User responses to questions (one per question, optional) |
| `personality_scores` | Computed dimension scores + confidence (JSONB, numeric only) |

### Ownership Boundaries
- Questions are system-owned reference data
- Answers belong to individual users
- Scores are 1:1 with users (nullable if never participated)

### What This Module Guarantees
- ✓ Participation is OPTIONAL (no penalty for skipping)
- ✓ Partial completion supported (scores improve over time)
- ✓ NO generated summaries stored (stateless on demand)
- ✓ Users NEVER see their own personality summary
- ✓ NO labels, rankings, or visible scores
- ✓ Does NOT influence matching or discovery order

---

## Read Paths: Key User Flows

### Daily Check-In Flow

```sql
-- 1. Get active streaks for user
SELECT s.id, s.streak_state, s.current_day, c.id as connection_id
FROM streaks s
JOIN connections c ON c.id = s.connection_id
WHERE (c.user_a_id = $user_id OR c.user_b_id = $user_id)
  AND s.streak_state IN ('active', 'at_risk')
  AND s.deleted_at IS NULL;

-- 2. Check if already checked in today
SELECT 1 FROM check_ins
WHERE streak_id = $streak_id
  AND user_id = $user_id
  AND check_in_date = CURRENT_DATE;

-- 3. Record check-in (if not exists)
INSERT INTO check_ins (id, streak_id, user_id, check_in_date, check_in_at, effort_signal)
VALUES (gen_random_uuid(), $streak_id, $user_id, CURRENT_DATE, NOW(), $effort_signal);

-- 4. Check if both checked in → update streak
-- (Application logic evaluates both users' check-in status)
```

---

## What the Database Guarantees vs Application Layer

| Guarantee | Enforced By |
|-----------|-------------|
| No duplicate check-ins same day | Database (UNIQUE constraint) |
| One active streak per connection | Database (UNIQUE partial index) |
| Tier connection limits | Database (CHECK) + Application (atomic increment) |
| No reveal 5 before Day 15 | Database (CHECK) + Application |
| Recovery payment by breaker only | Database (FK to breaker) + Application (auth) |
| State machine transitions | Application (with trigger validation) |
| Global refresh timer | Application (computes from `last_global_refresh_at`) |
| Total Wipe on capacity fill | Application (atomic transaction) |
| Credit conversion within 24h | Application (scheduled job) + Database (window tracking) |

---

*This document provides the mental model for understanding the Unora database. For relationship details, see ERD_RELATIONSHIPS.md. For exact schema, see TABLE_DEFINITIONS.md.*
