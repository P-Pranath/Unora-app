# TABLE_DEFINITIONS.md — Part 1: Types & Core Tables

**Version:** 1.0  
**Date:** January 2026  
**Database:** PostgreSQL 15+  
**Executable:** Yes (run in order)

---

## Enum Types

```sql
-- Server types (intent-based)
CREATE TYPE server_type_enum AS ENUM ('partner', 'friend', 'growth');

-- Subscription tiers
CREATE TYPE subscription_tier_enum AS ENUM ('free', 'plus', 'pro');

-- Verification status (progressive)
CREATE TYPE verification_status_enum AS ENUM (
  'pending', 'phone_verified', 'photos_submitted', 
  'photo_quality_verified', 'gov_id_verified'
);

-- Verification signal types
CREATE TYPE verification_signal_type_enum AS ENUM (
  'google_oauth', 'apple_oauth', 'linkedin_oauth', 
  'photo_quality', 'gov_id_token'
);



-- Discovery batch status
CREATE TYPE batch_status_enum AS ENUM ('active', 'consumed', 'expired');

-- Interest status
CREATE TYPE interest_status_enum AS ENUM ('pending', 'matched', 'expired', 'wiped');

-- Connection status
CREATE TYPE connection_status_enum AS ENUM ('active', 'terminated');

-- Streak states (explicit state machine)
CREATE TYPE streak_state_enum AS ENUM (
  'active', 'at_risk', 'payment_window', 
  'reset', 'completed', 'terminated'
);

-- Reveal status
CREATE TYPE reveal_status_enum AS ENUM ('locked', 'earned', 'purchased', 'unlocked');

-- Payment types
CREATE TYPE payment_type_enum AS ENUM ('subscription', 'recovery', 'reveal', 'nudge');

-- Credit transaction types
CREATE TYPE credit_transaction_type_enum AS ENUM ('conversion', 'spend', 'refund', 'adjustment');

-- Safety flag types
CREATE TYPE safety_flag_type_enum AS ENUM ('fake_consistency', 'manipulation', 'behavioral_shift');

-- Report types
CREATE TYPE report_type_enum AS ENUM ('harassment', 'fake', 'spam', 'other');

-- Report status
CREATE TYPE report_status_enum AS ENUM ('pending', 'reviewed', 'actioned', 'dismissed');
```

---

## Module 1: Identity & Verification

```sql
-- Core user account
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  phone_number VARCHAR(20) NOT NULL,
  phone_country_code VARCHAR(5) NOT NULL DEFAULT '+91',
  verification_status verification_status_enum NOT NULL DEFAULT 'pending',
  subscription_tier subscription_tier_enum NOT NULL DEFAULT 'free',
  
  -- Global inventory (spans all servers)
  active_connection_count INTEGER NOT NULL DEFAULT 0,
  last_global_refresh_at TIMESTAMP WITH TIME ZONE,
  refresh_available_at TIMESTAMP WITH TIME ZONE,
  
  -- Credit balance
  credit_balance INTEGER NOT NULL DEFAULT 0,
  
  -- Behavioral trust (computed by AI, stored for queries)
  behavioral_trust_score NUMERIC(3,2) CHECK (behavioral_trust_score BETWEEN 0 AND 1),
  
  -- Timestamps
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  CONSTRAINT chk_active_connections CHECK (active_connection_count >= 0)
);

COMMENT ON TABLE users IS 'Core user account with global inventory tracking';
COMMENT ON COLUMN users.active_connection_count IS 'Global across ALL servers, not per-server';
COMMENT ON COLUMN users.last_global_refresh_at IS 'Single refresh timer for all servers';

-- Verification signals (OAuth, photo quality, gov ID tokens)
CREATE TABLE verification_signals (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  signal_type verification_signal_type_enum NOT NULL,
  
  -- For OAuth: provider user ID (encrypted)
  external_id_encrypted BYTEA,
  
  -- For gov ID: tokenized reference only (no raw data)
  token_reference VARCHAR(255),
  
  verified_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  expires_at TIMESTAMP WITH TIME ZONE,
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE
);

COMMENT ON TABLE verification_signals IS 'Progressive verification: OAuth links, photo quality, gov ID tokens';
COMMENT ON COLUMN verification_signals.external_id_encrypted IS 'AES-256 encrypted OAuth provider ID';
COMMENT ON COLUMN verification_signals.token_reference IS 'Gov ID provider token, no raw document stored';
```

---

## Module 2: Profile & Hobbies

```sql
-- User profile (one per user)
CREATE TABLE profiles (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  
  -- Identity fields (locked after gov ID verification)
  first_name VARCHAR(50) NOT NULL,
  date_of_birth DATE NOT NULL,
  gender VARCHAR(20) NOT NULL,
  
  -- Location
  city VARCHAR(100) NOT NULL,
  
  -- Bio and intent
  bio TEXT,
  intent_statement TEXT,
  
  -- Optional fields (JSONB for flexibility)
  optional_fields JSONB DEFAULT '{}',
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  CONSTRAINT chk_dob CHECK (date_of_birth < CURRENT_DATE - INTERVAL '18 years')
);

COMMENT ON TABLE profiles IS 'User-provided profile data, 1:1 with users';
COMMENT ON COLUMN profiles.optional_fields IS 'JSON: height, education, languages, etc.';

-- User photos (3-6 required)
CREATE TABLE photos (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  
  storage_key VARCHAR(255) NOT NULL, -- S3/cloud storage reference
  display_order INTEGER NOT NULL,
  is_face_photo BOOLEAN NOT NULL DEFAULT false,
  
  -- AI validation
  ai_validated_at TIMESTAMP WITH TIME ZONE,
  ai_validation_passed BOOLEAN,
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  CONSTRAINT chk_display_order CHECK (display_order BETWEEN 1 AND 6)
);

COMMENT ON TABLE photos IS 'User photos, private until Day 15 reveal';
COMMENT ON COLUMN photos.is_face_photo IS 'At least one face photo required';

-- Master hobby list
CREATE TABLE hobby_options (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  category VARCHAR(50) NOT NULL,
  icon_name VARCHAR(50),
  micro_descriptions JSONB NOT NULL DEFAULT '[]',
  is_active BOOLEAN NOT NULL DEFAULT true,
  sort_order INTEGER NOT NULL DEFAULT 0
);

COMMENT ON TABLE hobby_options IS 'Master list of available hobbies';
COMMENT ON COLUMN hobby_options.micro_descriptions IS 'Array of pre-written descriptions users can select';

-- User's selected hobbies (2+ required)
CREATE TABLE hobbies (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  hobby_option_id UUID NOT NULL REFERENCES hobby_options(id),
  
  micro_description VARCHAR(200), -- User's selected/custom description
  display_order INTEGER NOT NULL,
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  UNIQUE (user_id, hobby_option_id) 
);

COMMENT ON TABLE hobbies IS 'User hobby selections, primary discovery anchor';
```

---

## Module 3: Discovery & Caching

```sql
-- Server reference table
CREATE TABLE servers (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  server_type server_type_enum NOT NULL UNIQUE,
  display_name VARCHAR(50) NOT NULL,
  icon_name VARCHAR(50) NOT NULL,
  sort_order INTEGER NOT NULL DEFAULT 0
);

-- Insert reference data
INSERT INTO servers (server_type, display_name, icon_name, sort_order) VALUES
  ('partner', 'Looking for Partner', 'HeartStraight', 1),
  ('friend', 'Friend / Companion', 'HandWaving', 2),
  ('growth', 'Accountability / Growth', 'Target', 3);

-- User's saved filters per server
CREATE TABLE filters (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  server_type server_type_enum NOT NULL,
  
  filter_config JSONB NOT NULL DEFAULT '{}',
  is_pending BOOLEAN NOT NULL DEFAULT false, -- Applied on next refresh
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  
  UNIQUE (user_id, server_type)
);

COMMENT ON TABLE filters IS 'Server-scoped filters, applied on refresh only';
COMMENT ON COLUMN filters.is_pending IS 'True if edited but not yet applied';

-- Discovery batches (one active per user per server)
CREATE TABLE discovery_batches (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  server_type server_type_enum NOT NULL,
  
  batch_status batch_status_enum NOT NULL DEFAULT 'active',
  filter_snapshot JSONB, -- Filters at time of generation
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  expires_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE
);

COMMENT ON TABLE discovery_batches IS 'Cached discovery batch per user per server';

-- Individual discovery cards (exactly 5 per batch)
CREATE TABLE discovery_cards (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  batch_id UUID NOT NULL REFERENCES discovery_batches(id),
  candidate_user_id UUID NOT NULL REFERENCES users(id),
  
  display_order INTEGER NOT NULL CHECK (display_order BETWEEN 1 AND 5),
  -- NOTE: NO compatibility_score or match_vector fields.
  -- Matching is RULE-BASED (mutual interest + filters), not AI-driven.
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE
);

COMMENT ON TABLE discovery_cards IS 'Individual cards in discovery batch, exactly 5';
```

---

## Module 4: Interest & Matching

```sql
-- Interest expressions (pending until mutual)
CREATE TABLE interests (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  sender_user_id UUID NOT NULL REFERENCES users(id),
  receiver_user_id UUID NOT NULL REFERENCES users(id),
  server_type server_type_enum NOT NULL,
  
  discovery_card_id UUID REFERENCES discovery_cards(id),
  interest_status interest_status_enum NOT NULL DEFAULT 'pending',
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  matched_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  CONSTRAINT chk_not_self CHECK (sender_user_id != receiver_user_id)
);

COMMENT ON TABLE interests IS 'Unilateral interest, pending until mutual';

-- Connections (formed on mutual interest)
CREATE TABLE connections (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_a_id UUID NOT NULL REFERENCES users(id),
  user_b_id UUID NOT NULL REFERENCES users(id),
  server_type server_type_enum NOT NULL,
  
  connection_status connection_status_enum NOT NULL DEFAULT 'active',
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  terminated_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  -- Enforce ordered pair (A < B) to prevent duplicates
  CONSTRAINT chk_user_order CHECK (user_a_id < user_b_id)
);

COMMENT ON TABLE connections IS 'Server-bound connection formed on mutual interest';

-- Audit log for Total Wipe events
CREATE TABLE interest_wipes (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  wiped_user_id UUID NOT NULL REFERENCES users(id),
  triggering_connection_id UUID NOT NULL REFERENCES connections(id),
  wiped_count INTEGER NOT NULL,
  wiped_interest_ids UUID[] NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE interest_wipes IS 'Audit: Total Wipe events when user fills final slot';
```

---

## Module 5: Streak Engine

```sql
-- Streaks (state machine per connection)
CREATE TABLE streaks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  connection_id UUID NOT NULL REFERENCES connections(id),
  
  streak_state streak_state_enum NOT NULL DEFAULT 'active',
  current_day INTEGER NOT NULL DEFAULT 1 CHECK (current_day BETWEEN 1 AND 15),
  reset_count INTEGER NOT NULL DEFAULT 0,
  
  -- Breaker tracking (NULL when active)
  breaker_user_id UUID REFERENCES users(id),
  
  -- Recovery window
  recovery_deadline_at TIMESTAMP WITH TIME ZONE,
  recovery_payment_id UUID,
  
  -- AI scoring (internal)
  streak_health_score NUMERIC(3,2),
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  completed_at TIMESTAMP WITH TIME ZONE,
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  -- State constraints
  CONSTRAINT chk_at_risk_has_breaker CHECK (
    streak_state NOT IN ('at_risk', 'payment_window') OR breaker_user_id IS NOT NULL
  ),
  CONSTRAINT chk_payment_window_has_deadline CHECK (
    streak_state != 'payment_window' OR recovery_deadline_at IS NOT NULL
  ),
  CONSTRAINT chk_completed_day_15 CHECK (
    streak_state != 'completed' OR current_day = 15
  )
);

COMMENT ON TABLE streaks IS 'First-class streak entity with explicit state machine';

-- Daily check-ins (immutable)
CREATE TABLE check_ins (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  streak_id UUID NOT NULL REFERENCES streaks(id),
  user_id UUID NOT NULL REFERENCES users(id),
  
  check_in_date DATE NOT NULL,
  check_in_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  
  -- Hobby Echo context
  effort_signal VARCHAR(100),
  hobby_context_id UUID REFERENCES hobby_options(id),
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
  -- No deleted_at: check-ins are immutable
);

COMMENT ON TABLE check_ins IS 'Daily check-in records, append-only, never deleted';

-- Nudges (bounded per at-risk period)
CREATE TABLE nudges (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  streak_id UUID NOT NULL REFERENCES streaks(id),
  sender_user_id UUID NOT NULL REFERENCES users(id),
  receiver_user_id UUID NOT NULL REFERENCES users(id),
  
  at_risk_period INTEGER NOT NULL, -- Which at-risk period (increments on reset)
  nudge_variant VARCHAR(100), -- AI-selected pre-written nudge
  
  sent_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  delivered_at TIMESTAMP WITH TIME ZONE,
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE nudges IS 'Tier-limited nudges, bounded per at-risk period';
```

---

## Module 6: Reveals

```sql
-- Reveal milestone configuration
CREATE TABLE reveal_milestones (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tier subscription_tier_enum NOT NULL,
  reveal_number INTEGER NOT NULL CHECK (reveal_number BETWEEN 1 AND 5),
  
  reveal_day_required INTEGER NOT NULL, -- Internal logic, never shown to users
  reveal_type VARCHAR(50) NOT NULL, -- personality, lifestyle, social, human, identity
  content_schema_version INTEGER NOT NULL DEFAULT 1,
  
  UNIQUE (tier, reveal_number)
);

-- Insert tier-specific reveal schedules
INSERT INTO reveal_milestones (tier, reveal_number, reveal_day_required, reveal_type) VALUES
  ('free', 1, 5, 'personality'), ('free', 2, 10, 'lifestyle'),
  ('plus', 1, 4, 'personality'), ('plus', 2, 8, 'lifestyle'), ('plus', 3, 12, 'social'),
  ('pro', 1, 3, 'personality'), ('pro', 2, 6, 'lifestyle'), ('pro', 3, 9, 'social'), ('pro', 4, 12, 'human');
-- Day 15 identity reveal is universal (handled separately)

COMMENT ON TABLE reveal_milestones IS 'Config: tier-specific reveal day requirements (internal)';

-- User's reveal state per streak
CREATE TABLE reveals (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  streak_id UUID NOT NULL REFERENCES streaks(id),
  user_id UUID NOT NULL REFERENCES users(id),
  reveal_number INTEGER NOT NULL CHECK (reveal_number BETWEEN 1 AND 5),
  
  reveal_status reveal_status_enum NOT NULL DEFAULT 'locked',
  unlocked_at TIMESTAMP WITH TIME ZONE,
  payment_id UUID, -- If purchased
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE,
  
  UNIQUE (streak_id, user_id, reveal_number)
);

COMMENT ON TABLE reveals IS 'Per-user reveal state within streak (tier isolation)';

-- AI-framed reveal content
CREATE TABLE reveal_contents (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  reveal_id UUID NOT NULL REFERENCES reveals(id) UNIQUE,
  
  content_json JSONB NOT NULL, -- AI-framed, versioned
  schema_version INTEGER NOT NULL DEFAULT 1,
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE reveal_contents IS 'AI-framed reveal content, never raw data dump';
```

---

## Module 7: Monetization & Credits

```sql
-- Subscriptions (one per user)
CREATE TABLE subscriptions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) UNIQUE,
  
  tier subscription_tier_enum NOT NULL DEFAULT 'free',
  billing_cycle_start DATE,
  billing_cycle_end DATE,
  
  -- Free recovery allowances
  free_recoveries_used INTEGER NOT NULL DEFAULT 0,
  free_recoveries_limit INTEGER NOT NULL DEFAULT 0, -- Plus: 1, Pro: 2 per connection
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE subscriptions IS 'User subscription tier and billing';

-- All payments (immutable transaction log)
CREATE TABLE payments (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  
  payment_type payment_type_enum NOT NULL,
  amount_inr INTEGER NOT NULL, -- Amount in paise
  
  -- Context references
  streak_id UUID REFERENCES streaks(id),
  reveal_id UUID REFERENCES reveals(id),
  
  -- External payment gateway
  gateway_reference VARCHAR(255),
  gateway_status VARCHAR(50),
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITH TIME ZONE -- Soft delete for audit
);

COMMENT ON TABLE payments IS 'All payment transactions, immutable for audit';

-- Credit balances
CREATE TABLE credits (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) UNIQUE,
  
  balance INTEGER NOT NULL DEFAULT 0 CHECK (balance >= 0),
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Credit transactions
CREATE TABLE credit_transactions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  credit_account_id UUID NOT NULL REFERENCES credits(id),
  
  transaction_type credit_transaction_type_enum NOT NULL,
  amount INTEGER NOT NULL, -- Positive for credits, negative for debits
  
  -- Source references
  source_payment_id UUID REFERENCES payments(id),
  protection_window_id UUID,
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- 24-hour credit protection windows
CREATE TABLE credit_protections (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  payment_id UUID NOT NULL REFERENCES payments(id),
  connection_id UUID NOT NULL REFERENCES connections(id),
  
  window_starts_at TIMESTAMP WITH TIME ZONE NOT NULL,
  window_ends_at TIMESTAMP WITH TIME ZONE NOT NULL,
  
  triggered BOOLEAN NOT NULL DEFAULT false,
  triggered_at TIMESTAMP WITH TIME ZONE,
  credit_transaction_id UUID REFERENCES credit_transactions(id),
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE credit_protections IS '24h window: auto-convert payment to credits if partner terminates';
```

---

## Module 8: Audit & Safety

```sql
-- Audit logs (append-only)
CREATE TABLE audit_logs (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  actor_user_id UUID REFERENCES users(id),
  
  event_type VARCHAR(50) NOT NULL,
  entity_type VARCHAR(50),
  entity_id UUID,
  
  event_data JSONB NOT NULL DEFAULT '{}',
  ip_address INET,
  user_agent TEXT,
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
  -- No deleted_at: append-only
);

COMMENT ON TABLE audit_logs IS 'Append-only audit trail for all significant events';

-- Safety flags (AI-generated)
CREATE TABLE safety_flags (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  flagged_user_id UUID NOT NULL REFERENCES users(id),
  
  flag_type safety_flag_type_enum NOT NULL,
  severity INTEGER NOT NULL CHECK (severity BETWEEN 1 AND 10),
  details JSONB,
  
  reviewed_at TIMESTAMP WITH TIME ZONE,
  reviewed_by UUID REFERENCES users(id),
  action_taken VARCHAR(100),
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE safety_flags IS 'AI-generated safety flags, silent interventions';

-- User blocks
CREATE TABLE blocks (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  blocker_user_id UUID NOT NULL REFERENCES users(id),
  blocked_user_id UUID NOT NULL REFERENCES users(id),
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  
  CONSTRAINT chk_not_self_block CHECK (blocker_user_id != blocked_user_id),
  UNIQUE (blocker_user_id, blocked_user_id)
);

-- User reports
CREATE TABLE reports (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  reporter_user_id UUID NOT NULL REFERENCES users(id),
  reported_user_id UUID NOT NULL REFERENCES users(id),
  
  connection_id UUID REFERENCES connections(id),
  streak_id UUID REFERENCES streaks(id),
  
  report_type report_type_enum NOT NULL,
  description TEXT,
  status report_status_enum NOT NULL DEFAULT 'pending',
  
  reviewed_at TIMESTAMP WITH TIME ZONE,
  reviewed_by UUID REFERENCES users(id),
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

---

## Module 9: Personality Intelligence

> **CRITICAL:** This module stores ONLY numeric scores + confidence. NO labels, NO summaries, NO generated text.

```sql
-- Personality questions (static, versioned)
CREATE TABLE personality_questions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  dimension VARCHAR(50) NOT NULL, -- e.g., 'openness', 'conscientiousness'
  question_text TEXT NOT NULL,
  options JSONB NOT NULL, -- Array of {id, text, weight}
  version INTEGER NOT NULL DEFAULT 1,
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE personality_questions IS 'Static, versioned personality questions. AI NEVER generates these.';
COMMENT ON COLUMN personality_questions.options IS 'Pre-written options with numeric weights for dimension scoring';

-- User responses (one per question, optional)
CREATE TABLE personality_answers (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id),
  question_id UUID NOT NULL REFERENCES personality_questions(id),
  selected_option_id VARCHAR(50) NOT NULL,
  
  context VARCHAR(50) NOT NULL, -- 'onboarding' or 'streak_checkin'
  streak_id UUID REFERENCES streaks(id), -- Set if answered during streak check-in
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  
  UNIQUE (user_id, question_id) -- One answer per question per user (latest wins)
);

COMMENT ON TABLE personality_answers IS 'User responses to personality questions. OPTIONAL, accumulative.';
COMMENT ON COLUMN personality_answers.context IS 'Where the question was answered: onboarding or streak check-in';

-- Computed dimension scores (numeric only)
CREATE TABLE personality_scores (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(id) UNIQUE,
  
  -- Dimension scores (0-100) + confidence (0-1) as JSONB
  dimension_scores JSONB NOT NULL DEFAULT '{}',
  -- Example: {"openness": {"score": 72, "confidence": 0.6}, "conscientiousness": {"score": 85, "confidence": 0.8}}
  
  last_computed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  last_asked_dimension VARCHAR(50), -- For rotation logic during streak check-ins
  
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE personality_scores IS 'Computed personality dimensions. AI reads these to generate summaries ON DEMAND.';
COMMENT ON COLUMN personality_scores.dimension_scores IS 'Numeric scores + confidence ONLY. NO text, labels, or summaries stored.';
```

**What This Module Does NOT Store:**

| Never Stored | Rationale |
|--------------|-----------|
| Generated personality summaries | Stateless generation on demand |
| Personality labels/categories | No visible labels per PRD |
| User-facing scores | Scores are backend-only |
| Raw answers from other users | Privacy: no cross-user data |

---

## Indexes (See INDEXES_PERFORMANCE.md for details)

```sql
-- Key indexes (abbreviated, see full list in performance doc)
CREATE UNIQUE INDEX idx_users_phone ON users (phone_number) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_profiles_user ON profiles (user_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_discovery_batches_active ON discovery_batches (user_id, server_type) 
  WHERE batch_status = 'active' AND deleted_at IS NULL;
CREATE UNIQUE INDEX idx_interests_pair ON interests (sender_user_id, receiver_user_id, server_type) 
  WHERE interest_status = 'pending' AND deleted_at IS NULL;
CREATE UNIQUE INDEX idx_streaks_active ON streaks (connection_id) 
  WHERE streak_state NOT IN ('reset', 'completed', 'terminated') AND deleted_at IS NULL;
CREATE UNIQUE INDEX idx_check_ins_unique ON check_ins (streak_id, user_id, check_in_date);

-- Personality Intelligence indexes
CREATE UNIQUE INDEX idx_personality_scores_user ON personality_scores (user_id);
CREATE UNIQUE INDEX idx_personality_answers_user_question ON personality_answers (user_id, question_id);
CREATE INDEX idx_personality_questions_dimension ON personality_questions (dimension) WHERE is_active = true;
```

---

*This DDL is executable in PostgreSQL 15+. Run modules in order.*
