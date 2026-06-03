CREATE TABLE experiments (
                             id BIGSERIAL PRIMARY KEY,

                             name TEXT NOT NULL,

                             rollout_percentage INT NOT NULL
                                 CHECK (rollout_percentage BETWEEN 0 AND 100),

                             start_date TIMESTAMP NULL,
                             end_date TIMESTAMP NULL,

                             status TEXT NOT NULL
);