# INDEXES_PERFORMANCE.md

**Version:** 1.0  
**Date:** January 2026  
**Database:** PostgreSQL 15+

---

## 1. Index Strategy Overview

### Design Principles

| Principle | Implementation |
|-----------|----------------|
| **Read-optimized for discovery** | Heavy indexing on discovery path |
| **Partial indexes for soft deletes** | All tables have `WHERE deleted_at IS NULL` |
| **Composite indexes for state lookups** | Multi-column indexes for streak/discovery |

### Assumptions

- **India-first scale**: High concurrency during local midnight
- **Burst patterns**: Check-ins spike at 23:00-23:59 IST
- **Target**: <50ms p99 for discovery reads, <100ms for check-in writes

---

## 2. High-Traffic Table Indexes

### users

```sql
CREATE UNIQUE INDEX idx_users_phone ON users (phone_number) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_refresh_available ON users (refresh_available_at) 
WHERE deleted_at IS NULL AND refresh_available_at IS NOT NULL;
```

### profiles

```sql
CREATE UNIQUE INDEX idx_profiles_user ON profiles (user_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_profiles_city ON profiles (city) WHERE deleted_at IS NULL;
CREATE INDEX idx_profiles_dob ON profiles (date_of_birth) WHERE deleted_at IS NULL;
```

### discovery_batches

```sql
-- One active batch per user per server
CREATE UNIQUE INDEX idx_discovery_batches_user_server_active 
ON discovery_batches (user_id, server_type) 
WHERE batch_status = 'active' AND deleted_at IS NULL;
```

### discovery_cards

```sql
CREATE INDEX idx_discovery_cards_batch ON discovery_cards (batch_id, display_order) 
WHERE deleted_at IS NULL;
CREATE INDEX idx_discovery_cards_candidate ON discovery_cards (candidate_user_id) 
WHERE deleted_at IS NULL;
```

### interests

```sql
-- For Total Wipe
CREATE INDEX idx_interests_sender_pending ON interests (sender_user_id) 
WHERE interest_status = 'pending' AND deleted_at IS NULL;

-- For mutual interest check
CREATE UNIQUE INDEX idx_interests_pair_pending 
ON interests (sender_user_id, receiver_user_id, server_type) 
WHERE interest_status = 'pending' AND deleted_at IS NULL;
```

### connections

```sql
CREATE INDEX idx_connections_user_a_active ON connections (user_a_id) 
WHERE connection_status = 'active' AND deleted_at IS NULL;
CREATE INDEX idx_connections_user_b_active ON connections (user_b_id) 
WHERE connection_status = 'active' AND deleted_at IS NULL;
CREATE UNIQUE INDEX idx_connections_pair ON connections (user_a_id, user_b_id, server_type) 
WHERE deleted_at IS NULL;
```

### streaks

```sql
-- One active streak per connection
CREATE UNIQUE INDEX idx_streaks_connection_active ON streaks (connection_id) 
WHERE streak_state NOT IN ('reset', 'completed', 'terminated') AND deleted_at IS NULL;

-- State-based queries
CREATE INDEX idx_streaks_at_risk ON streaks (streak_state, updated_at) 
WHERE streak_state = 'at_risk' AND deleted_at IS NULL;
CREATE INDEX idx_streaks_payment_window ON streaks (recovery_deadline_at) 
WHERE streak_state = 'payment_window' AND deleted_at IS NULL;
```

### check_ins

```sql
-- Prevents double check-in same day
CREATE UNIQUE INDEX idx_check_ins_streak_user_date ON check_ins (streak_id, user_id, check_in_date);
CREATE INDEX idx_check_ins_streak ON check_ins (streak_id, check_in_date DESC);
```

### reveals

```sql
CREATE INDEX idx_reveals_streak_user ON reveals (streak_id, user_id, reveal_number) 
WHERE deleted_at IS NULL;
```

### personality_scores / personality_answers

```sql
CREATE UNIQUE INDEX idx_personality_scores_user ON personality_scores (user_id);
CREATE UNIQUE INDEX idx_personality_answers_user_question ON personality_answers (user_id, question_id);
CREATE INDEX idx_personality_questions_dimension ON personality_questions (dimension) WHERE is_active = true;
```

---

## 3. Concurrency Controls

### Preventing Double Check-Ins

```sql
INSERT INTO check_ins (id, streak_id, user_id, check_in_date, check_in_at, effort_signal)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (streak_id, user_id, check_in_date) DO NOTHING
RETURNING id;
-- If RETURNING is NULL, check-in already exists
```

### Preventing Multiple Active Streaks

The UNIQUE partial index `idx_streaks_connection_active` prevents concurrent creation.

### Atomic Connection Slot Increment

```sql
BEGIN;
SELECT active_connection_count, subscription_tier FROM users WHERE id = $1 FOR UPDATE;
-- Check if under limit, then:
UPDATE users SET active_connection_count = active_connection_count + 1 WHERE id = $1;
INSERT INTO connections (...);
COMMIT;
```

---

## 4. Hot Tables & Monitoring

| Rank | Table | Critical Indexes |
|------|-------|------------------|
| 1 | `users` | `idx_users_phone`, PK |
| 2 | `discovery_batches` | `idx_discovery_batches_user_server_active` |
| 3 | `discovery_cards` | `idx_discovery_cards_batch` |
| 4 | `streaks` | `idx_streaks_connection_active` |
| 5 | `check_ins` | `idx_check_ins_streak_user_date` |

---

## 5. Partitioning Strategy (Future)

| Table | Partition Trigger | Strategy |
|-------|-------------------|----------|
| `check_ins` | >100M rows | Range by `check_in_date` (monthly) |
| `discovery_cards` | >50M rows | Range by `created_at` (weekly) |
| `audit_logs` | >1B rows | Range by `created_at` (daily) |

---

*For exact DDL, see TABLE_DEFINITIONS.md.*
