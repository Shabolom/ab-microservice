CREATE TABLE layer_experiments (
                                   layer_id BIGINT NOT NULL
                                       REFERENCES layers(id)
                                           ON DELETE CASCADE,

                                   experiment_id BIGINT NOT NULL
                                       REFERENCES experiments(id)
                                           ON DELETE CASCADE,

                                   PRIMARY KEY (layer_id, experiment_id)
);

CREATE INDEX idx_layer_experiments_experiment_id
    ON layer_experiments(experiment_id);