-- +goose Up

-- ============================================================================
-- SEED DATA: REVEAL MILESTONES
-- Three progressive reveal levels unlocked at different streak days
-- ============================================================================

INSERT INTO reveal_milestones (id, reveal_number, day_required, reveal_type, title, description, icon_name, credit_cost, is_active) VALUES
-- Reveal 1: Personality (Day 5)
(UUID(), 1, 5, 'personality', 'Personality Reveal', 'Discover personality insights based on your conversation patterns and compatibility analysis.', 'sparkles', 0, TRUE),

-- Reveal 2: Values (Day 10)
(UUID(), 2, 10, 'values', 'Values Reveal', 'Uncover shared values and life perspectives that create deeper connections.', 'heart', 25, TRUE),

-- Reveal 3: Lifestyle (Day 15 - Completion)
(UUID(), 3, 15, 'lifestyle', 'Lifestyle Reveal', 'Full compatibility assessment with actionable insights for your relationship.', 'crown', 50, TRUE);

-- +goose Down

DELETE FROM reveal_milestones;
