# ARCHITECTURAL_DECISIONS.md

**Version:** 1.0  
**Date:** January 2026  
**Database:** PostgreSQL 15+  
**Classification:** Internal / Engineering / Architecture Review

---

## Table of Contents

1. [Core Architectural Principles](#1-core-architectural-principles)
2. [Why Streaks Are First-Class Entities](#2-why-streaks-are-first-class-entities)
3. [Server-Scoped vs Global Data](#3-server-scoped-vs-global-data)
4. [Progressive Verification Architecture](#4-progressive-verification-architecture)
5. [Lifecycle State Machines](#5-lifecycle-state-machines)
6. [Soft Deletes & Audit Trail](#6-soft-deletes--audit-trail)
7. [Global Timers & Recovery Windows](#7-global-timers--recovery-windows)
8. [Safety-First Design Tradeoffs](#8-safety-first-design-tradeoffs)
9. [Future-Proofing Considerations](#9-future-proofing-considerations)

---

## 1. Core Architectural Principles

The Unora database schema is designed around five immutable principles derived from the product philosophy:

### 1.1 Intentional Friction as Data Model

> **Principle:** Every constraint in the product philosophy must be enforceable at the database level, not deferred to application code.

| Philosophy | Database Enforcement |
|------------|---------------------|
| Limited connections per tier | CHECK constraint on active connection count |
| One active streak per connection | UNIQUE partial index on connection + active state |
| 15-day identity reveal requirement | Streak day counter, NOT NULL reveal constraints |
| Breaker-only recovery payments | FK constraint linking recovery to breaker user |
| Mutual interest before match | Bidirectional interest records with timestamp ordering |

### 1.2 Earned Trust Over Instant Access

> **Principle:** Trust levels are computed from behavioral signals stored as immutable events, not derived on-the-fly.

The schema stores trust-building events as append-only records:
- Check-in completions (timestamped, never deleted)
- Recovery payments (immutable transaction log)
- Streak completions (permanent achievement record)
- Photo quality verifications (AI-evaluated, stored)

### 1.3 No Hidden State

> **Principle:** Every state transition must be explicitly recorded. Boolean flags are avoided in favor of enum states.

Instead of:
```sql
-- ANTI-PATTERN
is_active BOOLEAN,
is_at_risk BOOLEAN,
is_in_payment_window BOOLEAN
```

We use:
```sql
-- CORRECT PATTERN
streak_state streak_state_enum NOT NULL
-- Values: 'active', 'at_risk', 'payment_window', 'reset', 'completed', 'terminated'
```

### 1.4 Server-Time as Source of Truth

> **Principle:** All time-sensitive operations use server-generated timestamps. Client timestamps are never trusted.

- `created_at`: Server-generated, immutable
- `updated_at`: Server-generated on every mutation
- `check_in_at`: Server-generated when check-in recorded
- `refresh_available_at`: Server-computed based on tier rules

### 1.5 Global Resource Pools

> **Principle:** Connection slots and refresh timers are global across all servers. The schema must enforce this atomically.

The `users` table maintains global counters:
- `active_connection_count`: Incremented/decremented across all servers
- `last_global_refresh_at`: Single timestamp governing all server refreshes

---

## 2. Why Streaks Are First-Class Entities

### 2.1 Streaks Are Not Derived State

A common anti-pattern would be to compute streak status from check-in history:

```sql
-- ANTI-PATTERN: Derived streak
SELECT COUNT(*) as streak_days
FROM check_ins
WHERE connection_id = $1 AND created_at > last_reset_at;
```

**Why this fails:**
- Race conditions during at-risk transitions
- Recovery payment timing becomes ambiguous
- No atomic state machine enforcement
- Performance degrades with check-in history growth

### 2.2 Streak as State Machine

The `streaks` table maintains explicit state:

| Column | Purpose |
|--------|---------|
| `streak_state` | Enum: `active`, `at_risk`, `payment_window`, `reset`, `completed`, `terminated` |
| `current_day` | INTEGER (1-15), increments only on mutual check-in |
| `breaker_user_id` | NULL when active, set when one user misses |
| `recovery_deadline_at` | Computed when entering payment window |
| `reset_count` | Tracks resets for AI scoring |

### 2.3 State Transition Integrity

State transitions are enforced via CHECK constraints:

```sql
-- At-risk MUST have a breaker identified
CHECK (
  (streak_state != 'at_risk' AND streak_state != 'payment_window')
  OR breaker_user_id IS NOT NULL
)

-- Payment window MUST have deadline
CHECK (
  streak_state != 'payment_window' OR recovery_deadline_at IS NOT NULL
)

-- Completed streaks MUST have reached day 15
CHECK (
  streak_state != 'completed' OR current_day = 15
)
```

---

## 3. Server-Scoped vs Global Data

### 3.1 The Fundamental Split

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                            GLOBAL (User-Level)                               │
│  ├── User identity & verification                                            │
│  ├── Profile data & photos                                                   │
│  ├── Subscription tier                                                       │
│  ├── Active connection count (across all servers)                            │
│  ├── Global refresh timer (last_global_refresh_at)                           │
│  └── Credit balance                                                          │
├─────────────────────────────────────────────────────────────────────────────┤
│                         SERVER-SCOPED (Context-Specific)                     │
│  ├── Discovery batches (per server)                                          │
│  ├── Discovery cards (cached per server)                                     │
│  ├── Interests (server-specific, never cross)                                │
│  ├── Connections (server-bound)                                              │
│  └── Streaks (belong to server-scoped connections)                           │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 3.2 Why This Split Matters

**Global resources enforce scarcity:**
- A Free user's single connection slot spans all servers
- Refreshing in Dating locks refresh in Friend and Growth
- Tier limits apply to total active connections, not per-server

**Server-scoped data maintains intent integrity:**
- A match in Dating stays in Dating
- Filters and batches are server-specific
- Server theme/context is preserved per connection

### 3.3 Schema Enforcement

```sql
-- Connections are ALWAYS server-scoped
CREATE TABLE connections (
  id UUID PRIMARY KEY,
  server_type server_type_enum NOT NULL, -- 'partner', 'friend', 'growth'
  user_a_id UUID NOT NULL REFERENCES users(id),
  user_b_id UUID NOT NULL REFERENCES users(id),
  -- Connection cannot cross servers (enforced by application + audit)
);

-- Discovery batches are server-scoped
CREATE TABLE discovery_batches (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id),
  server_type server_type_enum NOT NULL,
  -- UNIQUE: One active batch per user per server
  UNIQUE (user_id, server_type) WHERE deleted_at IS NULL
);
```

---

## 4. Progressive Verification Architecture

### 4.1 Verification Layers

The PRD defines a progressive model with four distinct layers:

| Layer | Storage | Immutability |
|-------|---------|--------------|
| **Mandatory** (Phone + Photos) | `users`, `photos` | Append-only |
| **Trust Boosters** (OAuth links) | `verification_signals` | Soft delete only |
| **Behavioral** (Streaks, check-ins) | `check_ins`, `streaks` | Immutable events |
| **Identity** (Gov ID) | `verification_signals` (tokenized) | Token stored, no raw data |

### 4.2 Why Separate Tables

```sql
-- ANTI-PATTERN: Embedding verification in users table
ALTER TABLE users ADD COLUMN google_verified BOOLEAN;
-- ...leads to schema sprawl and unclear states
```

**Correct pattern:** `verification_signals` table with:
- Polymorphic signal types via enum
- Result storage (pass/fail/pending)
- Timestamp of verification
- Token reference (for Gov ID)

---

## 5. Lifecycle State Machines

### 5.1 Why Enums Over Booleans

Every major entity has an explicit lifecycle:

| Entity | State Enum | Values |
|--------|------------|--------|
| `users` | `verification_status` | `pending`, `phone_verified`, `photos_submitted`, `photo_quality_verified`, `gov_id_verified` |
| `discovery_batches` | `batch_status` | `active`, `consumed`, `expired` |
| `interests` | `interest_status` | `pending`, `matched`, `expired`, `wiped` |
| `connections` | `connection_status` | `active`, `terminated` |
| `streaks` | `streak_state` | `active`, `at_risk`, `payment_window`, `reset`, `completed`, `terminated` |
| `reveals` | `reveal_status` | `locked`, `earned`, `purchased`, `unlocked` |

### 5.2 Transition Enforcement

State transitions are validated at the application layer with database triggers for critical paths:

```sql
-- Streak cannot transition directly from 'active' to 'reset'
-- Must go through 'at_risk' → 'payment_window' first (unless both miss)
CREATE OR REPLACE FUNCTION validate_streak_transition()
RETURNS TRIGGER AS $$
BEGIN
  IF OLD.streak_state = 'active' AND NEW.streak_state = 'reset' THEN
    -- Only valid if BOTH users missed (no breaker identified)
    IF NEW.breaker_user_id IS NOT NULL THEN
      RAISE EXCEPTION 'Invalid state transition: active->reset requires dual miss';
    END IF;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
```

---

## 6. Soft Deletes & Audit Trail

### 6.1 Mandatory Soft Deletes

**All user-generated and trust-related data uses soft deletes.**

| Table | Soft Delete Required | Rationale |
|-------|---------------------|-----------|
| `users` | ✓ | Account recovery, dispute resolution |
| `profiles` | ✓ | Data integrity for existing connections |
| `photos` | ✓ | Reveal content preservation |
| `connections` | ✓ | Streak history, credit protection windows |
| `streaks` | ✓ | AI behavior analysis, pattern detection |
| `check_ins` | ✓ | Immutable behavioral record |
| `interests` | ✓ | Total Wipe audit trail |
| `payments` | ✓ | Financial compliance |
| `discovery_batches` | Optional | Performance (can hard delete old) |

### 6.2 Implementation Pattern

```sql
-- Standard soft delete columns on EVERY table
deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,

-- Partial indexes exclude soft-deleted rows
CREATE INDEX idx_users_active ON users(id) WHERE deleted_at IS NULL;

-- All queries MUST filter by deleted_at
SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL;
```

### 6.3 Cascading Behavior

Soft deleting a parent cascades to children via application logic:
- Deleting a connection soft-deletes its streak and check-ins
- Deleting a user soft-deletes all related records

**No ON DELETE CASCADE** — all cascades are explicit for audit trail.

---

## 7. Global Timers & Recovery Windows

### 7.1 Refresh Timer Architecture

**Single global timer per user, not per server:**

```sql
-- users table
last_global_refresh_at TIMESTAMP WITH TIME ZONE,
refresh_available_at TIMESTAMP WITH TIME ZONE, -- Computed

-- Application computes based on tier:
-- Free: 24 hours, Plus: 12 hours, Pro: 6 hours
```

**Why not per-server timers:**
- PRD explicitly requires global resource scarcity
- Prevents circumventing limits by switching servers
- Simplifies read path for discovery availability

### 7.2 Recovery Window Schema

```sql
-- streaks table
recovery_deadline_at TIMESTAMP WITH TIME ZONE,
recovery_payment_id UUID REFERENCES payments(id),

-- When entering payment_window:
-- recovery_deadline_at = end_of_day(current_day + 1)

-- Recovery payment made:
-- streak_state = 'active', current_day preserved
```

### 7.3 Timezone Handling

All timestamps are **UTC** in the database. Timezone conversion happens at the application layer.

```sql
-- Check-in window is evaluated in user's local timezone
-- But stored as UTC
check_in_at TIMESTAMP WITH TIME ZONE NOT NULL,
check_in_date DATE NOT NULL, -- Computed in user's timezone, stored for indexing
```

---

## 8. Safety-First Design Tradeoffs

### 8.1 Tradeoffs Made for Safety

| Decision | Flexibility Lost | Safety Gained |
|----------|------------------|---------------|
| Soft deletes mandatory | Storage overhead | Full audit trail, dispute resolution |
| Explicit state enums | Schema changes for new states | No ambiguous boolean combinations |
| Global inventory | Per-server independence | Focus enforcement, no limit circumvention |
| Breaker-only recovery | Flexibility in who pays | Fair accountability, no victim burden |
| 24h credit protection window | Immediate termination | Payment protection for good-faith users |
| No raw Gov ID storage | Direct verification display | Privacy protection, token-only |

### 8.2 What the Schema Prevents

| Prevented Scenario | Mechanism |
|--------------------|-----------|
| Double check-in same day | UNIQUE constraint on (user_id, streak_id, check_in_date) |
| Multiple active streaks per connection | UNIQUE partial index on connection + active state |
| Exceeding tier connection limit | CHECK constraint + trigger on active_connection_count |
| Late check-in after day ends | Application + recovery window boundary check |
| Reveal 5 before Day 15 | NOT NULL constraint requiring current_day = 15 |
| Interest spam across slots | Total Wipe trigger on capacity fill |

### 8.3 Trust Assumptions

The schema trusts:
- **Server clock accuracy** — All timestamps are server-generated
- **Application layer validation** — For complex business rules not expressible as constraints
- **Idempotency tokens** — For double-submit protection

The schema does NOT trust:
- **Client timestamps** — Never stored or used for decisions
- **User-provided verification claims** — Verified by external services, tokens stored
- **Implicit state** — All state is explicit

---

## 9. Future-Proofing Considerations

### 9.1 Changes That DO NOT Require Migration

| Change | Why Safe |
|--------|----------|
| New server types (e.g., 'family') | Add to `server_type_enum`, no table changes |
| New verification signals | Add to `verification_signal_type_enum` |
| New reveal content types | `reveal_content` is JSONB with schema version |
| Tier limit adjustments | Application constants, not schema |
| New streak milestones | `reveal_milestones` is JSONB configuration |

### 9.2 Changes That WILL Require Migration

| Change | Migration Complexity |
|--------|---------------------|
| Photo count max change | Index rebuild if limit-based indexes exist |
| Additional streak states | Enum modification + backfill |
| Cross-server connections | Fundamental schema change (not planned) |
| Multi-currency credits | New column + computation changes |

### 9.3 JSONB Usage Policy

JSONB is used ONLY for:
1. **AI scoring outputs** — Structure may evolve rapidly
2. **Reveal content** — Personalized per tier, versioned
3. **Nudge variants** — Selected by AI, pre-written content
4. **Device fingerprints** — Vendor-dependent structure

JSONB is NOT used for:
- Core entity relationships
- State that needs indexing
- Data with known, stable structure

### 9.4 Partitioning Considerations (Future)

High-volume tables that may require partitioning:

| Table | Partition Strategy | Trigger Point |
|-------|-------------------|---------------|
| `check_ins` | Range by `created_at` (monthly) | >100M rows |
| `discovery_cards` | Range by `created_at` (weekly) | >50M rows |
| `interests` | Range by `created_at` (monthly) | >50M rows |
| `audit_logs` | Range by `created_at` (daily) | >1B rows |

---

## Appendix A: Decision Log

| Decision | Date | Rationale | Alternatives Considered |
|----------|------|-----------|------------------------|
| UUID v7 for primary keys | Jan 2026 | Time-ordered for index efficiency, globally unique | UUID v4 (random, less index-friendly) |
| Enum types over VARCHAR | Jan 2026 | Type safety, smaller storage | VARCHAR with CHECK constraint |
| Soft delete via `deleted_at` | Jan 2026 | Audit trail, recovery capability | Boolean `is_deleted` (loses timestamp) |
| Global inventory counters | Jan 2026 | PRD requirement, atomic enforcement | Per-server tracking (violates philosophy) |
| Separate `streaks` from `connections` | Jan 2026 | Streak resets don't terminate connection | Embedded streak in connection (conflates concepts) |

---

*This document is the authoritative reference for architectural decisions made in the Unora database schema. All changes to core architecture must be appended to this document with rationale.*
