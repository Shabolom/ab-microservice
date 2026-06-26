ALTER TABLE feature_toggles
    ADD COLUMN buckets INT[] NOT NULL DEFAULT '{}',
    ADD COLUMN ios_buckets INT[] NOT NULL DEFAULT '{}',
    ADD COLUMN android_buckets INT[] NOT NULL DEFAULT '{}',
    ADD COLUMN web_buckets INT[] NOT NULL DEFAULT '{}';

ALTER TABLE feature_toggles
    ADD CONSTRAINT chk_feature_toggles_buckets_count
        CHECK (
            array_length(buckets, 1) IS NULL
                OR array_length(buckets, 1) <= 100
            ),
    ADD CONSTRAINT chk_feature_toggles_ios_buckets_count
        CHECK (
            array_length(ios_buckets, 1) IS NULL
                OR array_length(ios_buckets, 1) <= 100
        ),
    ADD CONSTRAINT chk_feature_toggles_android_buckets_count
        CHECK (
            array_length(android_buckets, 1) IS NULL
                OR array_length(android_buckets, 1) <= 100
        ),
    ADD CONSTRAINT chk_feature_toggles_web_buckets_count
        CHECK (
            array_length(web_buckets, 1) IS NULL
                OR array_length(web_buckets, 1) <= 100
        );