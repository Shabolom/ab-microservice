ALTER TABLE feature_toggles
    ADD COLUMN rollout_percentage INT NOT NULL DEFAULT 0;

ALTER TABLE feature_toggles
    ADD CONSTRAINT chk_feature_toggles_rollout_percentage
        CHECK (
            rollout_percentage >= 0
                AND rollout_percentage <= 100
            );