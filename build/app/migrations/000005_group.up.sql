CREATE TABLE experiment_groups (
                                   id BIGSERIAL PRIMARY KEY,

                                   experiment_id BIGINT NOT NULL
                                       REFERENCES experiments(id)
                                           ON DELETE CASCADE,

                                   rolling_percentage INT NOT NULL
                                       CHECK (rolling_percentage BETWEEN 0 AND 100),

                                   name TEXT NOT NULL,

                                   device_ids TEXT[] NOT NULL DEFAULT '{}'
);

CREATE INDEX idx_experiment_groups_experiment_id
    ON experiment_groups(experiment_id);