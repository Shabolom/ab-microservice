CREATE TABLE experiment_groups (
   id BIGSERIAL PRIMARY KEY,
   experiment_id TEXT NOT NULL
       REFERENCES experiments(id)
           ON DELETE CASCADE,
   name TEXT NOT NULL,
   rolling_percentage INT NOT NULL
       CHECK (rolling_percentage BETWEEN 0 AND 100)
);

CREATE INDEX idx_experiment_groups_experiment_id
    ON experiment_groups(experiment_id);