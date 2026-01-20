-- +goose Up

-- ============================================================================
-- UNORA DATABASE SCHEMA
-- Version: 1.0 | Date: January 2026
-- Converted from PostgreSQL to MySQL
-- ============================================================================

-- ============================================================================
-- MODULE 1: IDENTITY & VERIFICATION
-- ============================================================================

-- Core user account
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    
    -- Auth fields (supports both email/OAuth and phone)
    email VARCHAR(255),
    phone_number VARCHAR(20),
    phone_country_code VARCHAR(5) DEFAULT '+91',
    
    -- OAuth provider info
    provider VARCHAR(20),
    provider_user_id VARCHAR(255),
    
    -- Basic profile (from OAuth or user input)
    name VARCHAR(100),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    picture VARCHAR(500),
    
    -- Verification and tier
    verification_status ENUM('pending', 'phone_verified', 'photos_submitted', 'photo_quality_verified', 'gov_id_verified') NOT NULL DEFAULT 'pending',
    subscription_tier ENUM('free', 'plus', 'pro') NOT NULL DEFAULT 'free',
    
    -- Global inventory (spans all servers)
    active_connection_count INT NOT NULL DEFAULT 0,
    last_global_refresh_at DATETIME(3),
    refresh_available_at DATETIME(3),
    
    -- Credit balance
    credit_balance INT NOT NULL DEFAULT 0,
    
    -- Behavioral trust (computed by AI, stored for queries)
    behavioral_trust_score DECIMAL(3,2),
    
    -- Timestamps
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT chk_active_connections CHECK (active_connection_count >= 0),
    
    INDEX idx_users_deleted (deleted_at),
    UNIQUE INDEX idx_users_email (email),
    INDEX idx_users_phone (phone_number),
    INDEX idx_users_provider (provider, provider_user_id)
);

-- Verification signals (OAuth, photo quality, gov ID tokens)
CREATE TABLE verification_signals (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    signal_type ENUM('google_oauth', 'apple_oauth', 'linkedin_oauth', 'photo_quality', 'gov_id_token') NOT NULL,
    
    -- For OAuth: provider user ID (encrypted)
    external_id_encrypted BLOB,
    
    -- For gov ID: tokenized reference only (no raw data)
    token_reference VARCHAR(255),
    
    verified_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    expires_at DATETIME(3),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT fk_verification_signals_user FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_verification_signals_user (user_id)
);

-- ============================================================================
-- MODULE 2: PROFILE & HOBBIES
-- ============================================================================

-- User profile (one per user)
CREATE TABLE profiles (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    
    -- Identity fields (locked after gov ID verification)
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    date_of_birth DATE,
    gender VARCHAR(20),
    
    -- Location
    city VARCHAR(100),
    
    -- Bio and intent
    bio TEXT,
    intent_statement TEXT,
    
    -- Optional fields (JSON for flexibility)
    optional_fields JSON,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT fk_profiles_user FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE INDEX idx_profiles_user (user_id),
    INDEX idx_profiles_city (city),
    INDEX idx_profiles_dob (date_of_birth)
);

-- User photos (3-6 required)
CREATE TABLE photos (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    
    storage_key VARCHAR(255) NOT NULL,
    display_order INT NOT NULL,
    is_face_photo BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- AI validation
    ai_validated_at DATETIME(3),
    ai_validation_passed BOOLEAN,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT chk_display_order CHECK (display_order BETWEEN 1 AND 6),
    CONSTRAINT fk_photos_user FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_photos_user (user_id)
);

-- Master hobby list
CREATE TABLE hobby_options (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(50) NOT NULL,
    icon_name VARCHAR(50),
    micro_descriptions JSON NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    sort_order INT NOT NULL DEFAULT 0,
    
    INDEX idx_hobby_options_category (category),
    INDEX idx_hobby_options_active (is_active)
);

-- User's selected hobbies (2+ required)
CREATE TABLE hobbies (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    hobby_option_id CHAR(36) NOT NULL,
    
    micro_description VARCHAR(200),
    display_order INT NOT NULL,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT fk_hobbies_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_hobbies_option FOREIGN KEY (hobby_option_id) REFERENCES hobby_options(id),
    UNIQUE INDEX idx_hobbies_user_option (user_id, hobby_option_id)
);

-- ============================================================================
-- MODULE 3: DISCOVERY & CACHING
-- ============================================================================

-- Server reference table
CREATE TABLE servers (
    id CHAR(36) PRIMARY KEY,
    server_type ENUM('partner', 'friend', 'growth') NOT NULL UNIQUE,
    display_name VARCHAR(50) NOT NULL,
    icon_name VARCHAR(50) NOT NULL,
    sort_order INT NOT NULL DEFAULT 0
);

-- Insert reference data
INSERT INTO servers (id, server_type, display_name, icon_name, sort_order) VALUES
    (UUID(), 'partner', 'Looking for Partner', 'HeartStraight', 1),
    (UUID(), 'friend', 'Friend / Companion', 'HandWaving', 2),
    (UUID(), 'growth', 'Accountability / Growth', 'Target', 3);

-- User's saved filters per server
CREATE TABLE filters (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    server_type ENUM('partner', 'friend', 'growth') NOT NULL,
    
    filter_config JSON NOT NULL,
    is_pending BOOLEAN NOT NULL DEFAULT FALSE,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_filters_user FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE INDEX idx_filters_user_server (user_id, server_type)
);

-- Discovery batches (one active per user per server)
CREATE TABLE discovery_batches (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    server_type ENUM('partner', 'friend', 'growth') NOT NULL,
    
    batch_status ENUM('active', 'consumed', 'expired') NOT NULL DEFAULT 'active',
    filter_snapshot JSON,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    expires_at DATETIME(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT fk_discovery_batches_user FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_discovery_batches_user_server (user_id, server_type, batch_status)
);

-- Individual discovery cards (exactly 5 per batch)
CREATE TABLE discovery_cards (
    id CHAR(36) PRIMARY KEY,
    batch_id CHAR(36) NOT NULL,
    candidate_user_id CHAR(36) NOT NULL,
    
    display_order INT NOT NULL,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT chk_card_display_order CHECK (display_order BETWEEN 1 AND 5),
    CONSTRAINT fk_discovery_cards_batch FOREIGN KEY (batch_id) REFERENCES discovery_batches(id),
    CONSTRAINT fk_discovery_cards_candidate FOREIGN KEY (candidate_user_id) REFERENCES users(id),
    INDEX idx_discovery_cards_batch (batch_id, display_order),
    INDEX idx_discovery_cards_candidate (candidate_user_id)
);

-- ============================================================================
-- MODULE 4: INTEREST & MATCHING
-- ============================================================================

-- Interest expressions (pending until mutual)
CREATE TABLE interests (
    id CHAR(36) PRIMARY KEY,
    sender_user_id CHAR(36) NOT NULL,
    receiver_user_id CHAR(36) NOT NULL,
    server_type ENUM('partner', 'friend', 'growth') NOT NULL,
    
    discovery_card_id CHAR(36),
    interest_status ENUM('pending', 'matched', 'expired', 'wiped') NOT NULL DEFAULT 'pending',
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    matched_at DATETIME(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT chk_not_self CHECK (sender_user_id != receiver_user_id),
    CONSTRAINT fk_interests_sender FOREIGN KEY (sender_user_id) REFERENCES users(id),
    CONSTRAINT fk_interests_receiver FOREIGN KEY (receiver_user_id) REFERENCES users(id),
    CONSTRAINT fk_interests_card FOREIGN KEY (discovery_card_id) REFERENCES discovery_cards(id),
    INDEX idx_interests_sender_pending (sender_user_id, interest_status),
    INDEX idx_interests_receiver (receiver_user_id)
);

-- Connections (formed on mutual interest)
CREATE TABLE connections (
    id CHAR(36) PRIMARY KEY,
    user_a_id CHAR(36) NOT NULL,
    user_b_id CHAR(36) NOT NULL,
    server_type ENUM('partner', 'friend', 'growth') NOT NULL,
    
    connection_status ENUM('active', 'terminated') NOT NULL DEFAULT 'active',
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    terminated_at DATETIME(3),
    deleted_at DATETIME(3),
    
    -- Enforce ordered pair (A < B) to prevent duplicates
    CONSTRAINT chk_user_order CHECK (user_a_id < user_b_id),
    CONSTRAINT fk_connections_user_a FOREIGN KEY (user_a_id) REFERENCES users(id),
    CONSTRAINT fk_connections_user_b FOREIGN KEY (user_b_id) REFERENCES users(id),
    INDEX idx_connections_user_a_active (user_a_id, connection_status),
    INDEX idx_connections_user_b_active (user_b_id, connection_status),
    UNIQUE INDEX idx_connections_pair (user_a_id, user_b_id, server_type)
);

-- Audit log for Total Wipe events
CREATE TABLE interest_wipes (
    id CHAR(36) PRIMARY KEY,
    wiped_user_id CHAR(36) NOT NULL,
    triggering_connection_id CHAR(36) NOT NULL,
    wiped_count INT NOT NULL,
    wiped_interest_ids JSON NOT NULL,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_interest_wipes_user FOREIGN KEY (wiped_user_id) REFERENCES users(id),
    CONSTRAINT fk_interest_wipes_connection FOREIGN KEY (triggering_connection_id) REFERENCES connections(id),
    INDEX idx_interest_wipes_user (wiped_user_id)
);

-- ============================================================================
-- MODULE 5: STREAK ENGINE
-- ============================================================================

-- Streaks (state machine per connection)
CREATE TABLE streaks (
    id CHAR(36) PRIMARY KEY,
    connection_id CHAR(36) NOT NULL,
    
    streak_state ENUM('active', 'at_risk', 'payment_window', 'reset', 'completed', 'terminated') NOT NULL DEFAULT 'active',
    current_day INT NOT NULL DEFAULT 1,
    reset_count INT NOT NULL DEFAULT 0,
    
    -- Breaker tracking (NULL when active)
    breaker_user_id CHAR(36),
    
    -- Recovery window
    recovery_deadline_at DATETIME(3),
    recovery_payment_id CHAR(36),
    
    -- AI scoring (internal)
    streak_health_score DECIMAL(3,2),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    completed_at DATETIME(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT chk_current_day CHECK (current_day BETWEEN 1 AND 15),
    CONSTRAINT fk_streaks_connection FOREIGN KEY (connection_id) REFERENCES connections(id),
    CONSTRAINT fk_streaks_breaker FOREIGN KEY (breaker_user_id) REFERENCES users(id),
    INDEX idx_streaks_connection (connection_id),
    INDEX idx_streaks_state (streak_state),
    INDEX idx_streaks_at_risk (streak_state, updated_at)
);

-- Daily check-ins (immutable)
CREATE TABLE check_ins (
    id CHAR(36) PRIMARY KEY,
    streak_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    
    check_in_date DATE NOT NULL,
    check_in_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    -- Hobby Echo context
    effort_signal VARCHAR(100),
    hobby_context_id CHAR(36),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    -- No deleted_at: check-ins are immutable
    
    CONSTRAINT fk_check_ins_streak FOREIGN KEY (streak_id) REFERENCES streaks(id),
    CONSTRAINT fk_check_ins_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_check_ins_hobby FOREIGN KEY (hobby_context_id) REFERENCES hobby_options(id),
    UNIQUE INDEX idx_check_ins_unique (streak_id, user_id, check_in_date),
    INDEX idx_check_ins_streak_date (streak_id, check_in_date DESC)
);

-- Nudges (bounded per at-risk period)
CREATE TABLE nudges (
    id CHAR(36) PRIMARY KEY,
    streak_id CHAR(36) NOT NULL,
    sender_user_id CHAR(36) NOT NULL,
    receiver_user_id CHAR(36) NOT NULL,
    
    at_risk_period INT NOT NULL,
    nudge_variant VARCHAR(100),
    
    sent_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    delivered_at DATETIME(3),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_nudges_streak FOREIGN KEY (streak_id) REFERENCES streaks(id),
    CONSTRAINT fk_nudges_sender FOREIGN KEY (sender_user_id) REFERENCES users(id),
    CONSTRAINT fk_nudges_receiver FOREIGN KEY (receiver_user_id) REFERENCES users(id),
    INDEX idx_nudges_streak (streak_id, at_risk_period)
);

-- ============================================================================
-- MODULE 6: REVEALS
-- ============================================================================

-- Reveal milestone configuration
CREATE TABLE reveal_milestones (
    id CHAR(36) PRIMARY KEY,
    tier ENUM('free', 'plus', 'pro') NOT NULL,
    reveal_number INT NOT NULL,
    
    reveal_day_required INT NOT NULL,
    reveal_type VARCHAR(50) NOT NULL,
    content_schema_version INT NOT NULL DEFAULT 1,
    
    CONSTRAINT chk_reveal_number CHECK (reveal_number BETWEEN 1 AND 5),
    UNIQUE INDEX idx_reveal_milestones_tier_number (tier, reveal_number)
);

-- Insert tier-specific reveal schedules
INSERT INTO reveal_milestones (id, tier, reveal_number, reveal_day_required, reveal_type) VALUES
    (UUID(), 'free', 1, 5, 'personality'),
    (UUID(), 'free', 2, 10, 'lifestyle'),
    (UUID(), 'plus', 1, 4, 'personality'),
    (UUID(), 'plus', 2, 8, 'lifestyle'),
    (UUID(), 'plus', 3, 12, 'social'),
    (UUID(), 'pro', 1, 3, 'personality'),
    (UUID(), 'pro', 2, 6, 'lifestyle'),
    (UUID(), 'pro', 3, 9, 'social'),
    (UUID(), 'pro', 4, 12, 'human');

-- User's reveal state per streak
CREATE TABLE reveals (
    id CHAR(36) PRIMARY KEY,
    streak_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    reveal_number INT NOT NULL,
    
    reveal_status ENUM('locked', 'earned', 'purchased', 'unlocked') NOT NULL DEFAULT 'locked',
    unlocked_at DATETIME(3),
    payment_id CHAR(36),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT chk_reveals_number CHECK (reveal_number BETWEEN 1 AND 5),
    CONSTRAINT fk_reveals_streak FOREIGN KEY (streak_id) REFERENCES streaks(id),
    CONSTRAINT fk_reveals_user FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE INDEX idx_reveals_streak_user_number (streak_id, user_id, reveal_number)
);

-- AI-framed reveal content
CREATE TABLE reveal_contents (
    id CHAR(36) PRIMARY KEY,
    reveal_id CHAR(36) NOT NULL UNIQUE,
    
    content_json JSON NOT NULL,
    schema_version INT NOT NULL DEFAULT 1,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_reveal_contents_reveal FOREIGN KEY (reveal_id) REFERENCES reveals(id)
);

-- ============================================================================
-- MODULE 7: MONETIZATION & CREDITS
-- ============================================================================

-- Subscriptions (one per user)
CREATE TABLE subscriptions (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL UNIQUE,
    
    tier ENUM('free', 'plus', 'pro') NOT NULL DEFAULT 'free',
    billing_cycle_start DATE,
    billing_cycle_end DATE,
    
    -- Free recovery allowances
    free_recoveries_used INT NOT NULL DEFAULT 0,
    free_recoveries_limit INT NOT NULL DEFAULT 0,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_subscriptions_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- All payments (immutable transaction log)
CREATE TABLE payments (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    
    payment_type ENUM('subscription', 'recovery', 'reveal', 'nudge') NOT NULL,
    amount_inr INT NOT NULL,
    
    -- Context references
    streak_id CHAR(36),
    reveal_id CHAR(36),
    
    -- External payment gateway
    gateway_reference VARCHAR(255),
    gateway_status VARCHAR(50),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    deleted_at DATETIME(3),
    
    CONSTRAINT fk_payments_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_payments_streak FOREIGN KEY (streak_id) REFERENCES streaks(id),
    CONSTRAINT fk_payments_reveal FOREIGN KEY (reveal_id) REFERENCES reveals(id),
    INDEX idx_payments_user (user_id),
    INDEX idx_payments_type (payment_type)
);

-- Update streaks to reference payments
ALTER TABLE streaks
    ADD CONSTRAINT fk_streaks_recovery_payment FOREIGN KEY (recovery_payment_id) REFERENCES payments(id);

-- Update reveals to reference payments
ALTER TABLE reveals
    ADD CONSTRAINT fk_reveals_payment FOREIGN KEY (payment_id) REFERENCES payments(id);

-- Credit balances
CREATE TABLE credits (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL UNIQUE,
    
    balance INT NOT NULL DEFAULT 0,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    
    CONSTRAINT chk_balance CHECK (balance >= 0),
    CONSTRAINT fk_credits_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Credit transactions
CREATE TABLE credit_transactions (
    id CHAR(36) PRIMARY KEY,
    credit_account_id CHAR(36) NOT NULL,
    
    transaction_type ENUM('conversion', 'spend', 'refund', 'adjustment') NOT NULL,
    amount INT NOT NULL,
    
    -- Source references
    source_payment_id CHAR(36),
    protection_window_id CHAR(36),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_credit_transactions_account FOREIGN KEY (credit_account_id) REFERENCES credits(id),
    CONSTRAINT fk_credit_transactions_payment FOREIGN KEY (source_payment_id) REFERENCES payments(id),
    INDEX idx_credit_transactions_account (credit_account_id)
);

-- 24-hour credit protection windows
CREATE TABLE credit_protections (
    id CHAR(36) PRIMARY KEY,
    payment_id CHAR(36) NOT NULL,
    connection_id CHAR(36) NOT NULL,
    
    window_starts_at DATETIME(3) NOT NULL,
    window_ends_at DATETIME(3) NOT NULL,
    
    triggered BOOLEAN NOT NULL DEFAULT FALSE,
    triggered_at DATETIME(3),
    credit_transaction_id CHAR(36),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_credit_protections_payment FOREIGN KEY (payment_id) REFERENCES payments(id),
    CONSTRAINT fk_credit_protections_connection FOREIGN KEY (connection_id) REFERENCES connections(id),
    CONSTRAINT fk_credit_protections_transaction FOREIGN KEY (credit_transaction_id) REFERENCES credit_transactions(id),
    INDEX idx_credit_protections_window (window_ends_at, triggered)
);

-- Update credit_transactions to reference protection windows
ALTER TABLE credit_transactions
    ADD CONSTRAINT fk_credit_transactions_protection FOREIGN KEY (protection_window_id) REFERENCES credit_protections(id);

-- ============================================================================
-- MODULE 8: AUDIT & SAFETY
-- ============================================================================

-- Audit logs (append-only)
CREATE TABLE audit_logs (
    id CHAR(36) PRIMARY KEY,
    actor_user_id CHAR(36),
    
    event_type VARCHAR(50) NOT NULL,
    entity_type VARCHAR(50),
    entity_id CHAR(36),
    
    event_data JSON NOT NULL,
    ip_address VARCHAR(45),
    user_agent TEXT,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    -- No deleted_at: append-only
    
    CONSTRAINT fk_audit_logs_actor FOREIGN KEY (actor_user_id) REFERENCES users(id),
    INDEX idx_audit_logs_actor (actor_user_id),
    INDEX idx_audit_logs_event (event_type),
    INDEX idx_audit_logs_entity (entity_type, entity_id),
    INDEX idx_audit_logs_created (created_at)
);

-- Safety flags (AI-generated)
CREATE TABLE safety_flags (
    id CHAR(36) PRIMARY KEY,
    flagged_user_id CHAR(36) NOT NULL,
    
    flag_type ENUM('fake_consistency', 'manipulation', 'behavioral_shift') NOT NULL,
    severity INT NOT NULL,
    details JSON,
    
    reviewed_at DATETIME(3),
    reviewed_by CHAR(36),
    action_taken VARCHAR(100),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT chk_severity CHECK (severity BETWEEN 1 AND 10),
    CONSTRAINT fk_safety_flags_user FOREIGN KEY (flagged_user_id) REFERENCES users(id),
    CONSTRAINT fk_safety_flags_reviewer FOREIGN KEY (reviewed_by) REFERENCES users(id),
    INDEX idx_safety_flags_user (flagged_user_id),
    INDEX idx_safety_flags_pending (reviewed_at)
);

-- User blocks
CREATE TABLE blocks (
    id CHAR(36) PRIMARY KEY,
    blocker_user_id CHAR(36) NOT NULL,
    blocked_user_id CHAR(36) NOT NULL,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT chk_not_self_block CHECK (blocker_user_id != blocked_user_id),
    CONSTRAINT fk_blocks_blocker FOREIGN KEY (blocker_user_id) REFERENCES users(id),
    CONSTRAINT fk_blocks_blocked FOREIGN KEY (blocked_user_id) REFERENCES users(id),
    UNIQUE INDEX idx_blocks_pair (blocker_user_id, blocked_user_id)
);

-- User reports
CREATE TABLE reports (
    id CHAR(36) PRIMARY KEY,
    reporter_user_id CHAR(36) NOT NULL,
    reported_user_id CHAR(36) NOT NULL,
    
    connection_id CHAR(36),
    streak_id CHAR(36),
    
    report_type ENUM('harassment', 'fake', 'spam', 'other') NOT NULL,
    description TEXT,
    status ENUM('pending', 'reviewed', 'actioned', 'dismissed') NOT NULL DEFAULT 'pending',
    
    reviewed_at DATETIME(3),
    reviewed_by CHAR(36),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_reports_reporter FOREIGN KEY (reporter_user_id) REFERENCES users(id),
    CONSTRAINT fk_reports_reported FOREIGN KEY (reported_user_id) REFERENCES users(id),
    CONSTRAINT fk_reports_connection FOREIGN KEY (connection_id) REFERENCES connections(id),
    CONSTRAINT fk_reports_streak FOREIGN KEY (streak_id) REFERENCES streaks(id),
    CONSTRAINT fk_reports_reviewer FOREIGN KEY (reviewed_by) REFERENCES users(id),
    INDEX idx_reports_status (status),
    INDEX idx_reports_reported (reported_user_id)
);

-- ============================================================================
-- MODULE 9: PERSONALITY INTELLIGENCE
-- ============================================================================

-- Personality questions (static, versioned)
CREATE TABLE personality_questions (
    id CHAR(36) PRIMARY KEY,
    dimension VARCHAR(50) NOT NULL,
    question_text TEXT NOT NULL,
    options JSON NOT NULL,
    version INT NOT NULL DEFAULT 1,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    INDEX idx_personality_questions_dimension (dimension, is_active)
);

-- User responses (one per question, optional)
CREATE TABLE personality_answers (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    question_id CHAR(36) NOT NULL,
    selected_option_id VARCHAR(50) NOT NULL,
    
    context VARCHAR(50) NOT NULL,
    streak_id CHAR(36),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_personality_answers_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_personality_answers_question FOREIGN KEY (question_id) REFERENCES personality_questions(id),
    CONSTRAINT fk_personality_answers_streak FOREIGN KEY (streak_id) REFERENCES streaks(id),
    UNIQUE INDEX idx_personality_answers_user_question (user_id, question_id)
);

-- Computed dimension scores (numeric only)
CREATE TABLE personality_scores (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL UNIQUE,
    
    -- Dimension scores as JSON: {"openness": {"score": 72, "confidence": 0.6}}
    dimension_scores JSON NOT NULL,
    
    last_computed_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    last_asked_dimension VARCHAR(50),
    
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    
    CONSTRAINT fk_personality_scores_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down

-- Drop tables in reverse order of creation (respecting foreign keys)

-- Module 9: Personality Intelligence
DROP TABLE IF EXISTS personality_scores;
DROP TABLE IF EXISTS personality_answers;
DROP TABLE IF EXISTS personality_questions;

-- Module 8: Audit & Safety
DROP TABLE IF EXISTS reports;
DROP TABLE IF EXISTS blocks;
DROP TABLE IF EXISTS safety_flags;
DROP TABLE IF EXISTS audit_logs;

-- Module 7: Monetization & Credits (remove FKs first)
ALTER TABLE credit_transactions DROP FOREIGN KEY IF EXISTS fk_credit_transactions_protection;
ALTER TABLE streaks DROP FOREIGN KEY IF EXISTS fk_streaks_recovery_payment;
ALTER TABLE reveals DROP FOREIGN KEY IF EXISTS fk_reveals_payment;
DROP TABLE IF EXISTS credit_protections;
DROP TABLE IF EXISTS credit_transactions;
DROP TABLE IF EXISTS credits;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS subscriptions;

-- Module 6: Reveals
DROP TABLE IF EXISTS reveal_contents;
DROP TABLE IF EXISTS reveals;
DROP TABLE IF EXISTS reveal_milestones;

-- Module 5: Streak Engine
DROP TABLE IF EXISTS nudges;
DROP TABLE IF EXISTS check_ins;
DROP TABLE IF EXISTS streaks;

-- Module 4: Interest & Matching
DROP TABLE IF EXISTS interest_wipes;
DROP TABLE IF EXISTS connections;
DROP TABLE IF EXISTS interests;

-- Module 3: Discovery & Caching
DROP TABLE IF EXISTS discovery_cards;
DROP TABLE IF EXISTS discovery_batches;
DROP TABLE IF EXISTS filters;
DROP TABLE IF EXISTS servers;

-- Module 2: Profile & Hobbies
DROP TABLE IF EXISTS hobbies;
DROP TABLE IF EXISTS hobby_options;
DROP TABLE IF EXISTS photos;
DROP TABLE IF EXISTS profiles;

-- Module 1: Identity & Verification
DROP TABLE IF EXISTS verification_signals;
DROP TABLE IF EXISTS users;
