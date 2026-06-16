ALTER TABLE customparameter
ALTER COLUMN type TYPE VARCHAR(20);

ALTER TABLE customparameter
DROP CONSTRAINT IF EXISTS customparameter_type_check;

ALTER TABLE customparameter
    ADD CONSTRAINT customparameter_type_check
        CHECK (
            type IN (
                     'STRING',
                     'INT',
                     'SEMVER',
                     'DATE',
                     'STRING_ARRAY',
                     'INT_ARRAY',
                     'SEMVER_ARRAY',
                     'DATE_ARRAY'
                )
            );