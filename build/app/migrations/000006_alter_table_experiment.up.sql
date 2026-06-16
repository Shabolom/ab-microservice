ALTER TABLE experiments
    ADD COLUMN namespace TEXT NOT NULL DEFAULT 'default';

ALTER TABLE experiments
    ADD CONSTRAINT uq_experiments_namespace_name
        UNIQUE (namespace, name);