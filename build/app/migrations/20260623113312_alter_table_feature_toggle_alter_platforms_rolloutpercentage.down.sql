UPDATE feature_toggles
SET rollout_percentage = 0
WHERE rollout_percentage IS NULL;

ALTER TABLE feature_toggles
    ALTER COLUMN rollout_percentage SET NOT NULL;