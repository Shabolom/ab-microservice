CREATE TABLE layer_experiments (
                                   layer_id BIGINT NOT NULL
                                       REFERENCES layers(id)
                                           ON DELETE CASCADE,

                                   experiment_id BIGINT NOT NULL
                                       REFERENCES experiments(id)
                                           ON DELETE CASCADE,

                                   bucket INT[] NOT NULL DEFAULT '{}',

                                   PRIMARY KEY (layer_id, experiment_id),

                                   CHECK (
                                       array_length(bucket, 1) IS NULL
                                           OR array_length(bucket, 1) <= 100
                                       )
);

CREATE INDEX idx_layer_experiments_experiment_id
    ON layer_experiments(experiment_id);