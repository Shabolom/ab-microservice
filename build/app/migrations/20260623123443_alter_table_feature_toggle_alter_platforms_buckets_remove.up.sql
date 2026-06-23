ALTER TABLE feature_toggles
DROP CONSTRAINT IF EXISTS chk_feature_toggles_buckets_count,
    DROP CONSTRAINT IF EXISTS chk_feature_toggles_ios_buckets_count,
    DROP CONSTRAINT IF EXISTS chk_feature_toggles_android_buckets_count,
    DROP CONSTRAINT IF EXISTS chk_feature_toggles_web_buckets_count;

ALTER TABLE feature_toggles
DROP COLUMN IF EXISTS buckets,
    DROP COLUMN IF EXISTS ios_buckets,
    DROP COLUMN IF EXISTS android_buckets,
    DROP COLUMN IF EXISTS web_buckets;