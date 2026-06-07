ALTER TABLE experiments
DROP COLUMN IF EXISTS passing_cities,
    DROP COLUMN IF EXISTS excluded_cities,
    DROP COLUMN IF EXISTS passing_stores,
    DROP COLUMN IF EXISTS excluded_stores;