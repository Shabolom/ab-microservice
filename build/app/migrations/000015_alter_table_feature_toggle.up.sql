ALTER TABLE feature_toggles
DROP CONSTRAINT IF EXISTS chk_feature_toggles_rollout_percentage;

ALTER TABLE feature_toggles
DROP COLUMN rollout_percentage;

ALTER TABLE feature_toggles
    ADD COLUMN buckets INT[] NOT NULL DEFAULT '{}';

ALTER TABLE feature_toggles
    ADD CONSTRAINT chk_feature_toggles_buckets_count
        CHECK (
            array_length(buckets, 1) IS NULL
                OR array_length(buckets, 1) <= 100
            );