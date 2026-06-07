ALTER TABLE experiments
    ADD COLUMN passing_cities VARCHAR[] NOT NULL DEFAULT '{}',
    ADD COLUMN excluded_cities VARCHAR[] NOT NULL DEFAULT '{}',
    ADD COLUMN passing_stores VARCHAR[] NOT NULL DEFAULT '{}',
    ADD COLUMN excluded_stores VARCHAR[] NOT NULL DEFAULT '{}';