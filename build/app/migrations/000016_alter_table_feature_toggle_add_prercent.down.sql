ALTER TABLE feature_toggles
DROP CONSTRAINT IF EXISTS chk_feature_toggles_rollout_percentage;

ALTER TABLE feature_toggles
DROP COLUMN IF EXISTS rollout_percentage;