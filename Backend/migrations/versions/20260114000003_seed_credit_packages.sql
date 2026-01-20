-- +goose Up

-- ============================================================================
-- SEED DATA: CREDIT PACKAGES
-- Tiered credit packages with bonus credits for larger purchases
-- ============================================================================

INSERT INTO credit_packages (id, name, description, credit_amount, bonus_credits, price_amount, currency, discount_percent, badge_text, is_popular, is_active, sort_order) VALUES
-- Starter Pack
(UUID(), 'Starter Pack', 'Perfect for trying out premium features', 50, 0, 9900, 'INR', 0, NULL, FALSE, TRUE, 1),

-- Popular Pack
(UUID(), 'Popular Pack', 'Most chosen by our users', 150, 25, 24900, 'INR', 10, 'Most Popular', TRUE, TRUE, 2),

-- Value Pack
(UUID(), 'Value Pack', 'Great value for regular users', 350, 75, 49900, 'INR', 15, 'Best Value', FALSE, TRUE, 3),

-- Premium Pack
(UUID(), 'Premium Pack', 'Ultimate pack for power users', 800, 200, 99900, 'INR', 20, 'Premium', FALSE, TRUE, 4);

-- +goose Down

DELETE FROM credit_packages;
