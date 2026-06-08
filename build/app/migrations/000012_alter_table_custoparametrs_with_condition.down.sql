ALTER TABLE customparameter
DROP CONSTRAINT IF EXISTS customparameter_type_check;

ALTER TABLE customparameter
    ADD CONSTRAINT customparameter_type_check
        CHECK (
            type IN (
                     'STRING',
                     'INT',
                     'SEMVER',
                     'DATE'
                )
            );

ALTER TABLE customparameter
ALTER COLUMN type TYPE VARCHAR(9);