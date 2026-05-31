ALTER TABLE experiment_groups
    ALTER COLUMN device_ids DROP DEFAULT;

ALTER TABLE experiment_groups
ALTER COLUMN device_ids TYPE BIGINT[]
USING device_ids::BIGINT[];

ALTER TABLE experiment_groups
    ALTER COLUMN device_ids SET DEFAULT '{}'::BIGINT[];