ALTER TABLE experiments
DROP CONSTRAINT IF EXISTS uq_experiments_namespace_name;

ALTER TABLE experiments
DROP COLUMN IF EXISTS namespace;