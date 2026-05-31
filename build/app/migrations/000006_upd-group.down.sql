ALTER TABLE experiment_groups
    ALTER COLUMN device_ids DROP DEFAULT;

ALTER TABLE experiment_groups
ALTER COLUMN device_ids TYPE TEXT[]
USING device_ids::TEXT[];

ALTER TABLE experiment_groups
    ALTER COLUMN device_ids SET DEFAULT '{}'::TEXT[];