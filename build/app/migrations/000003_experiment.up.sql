CREATE TABLE experiments (
                             id BIGSERIAL PRIMARY KEY,

                             name TEXT NOT NULL,

                             rollout_percentage INT NOT NULL
                                 CHECK (rollout_percentage BETWEEN 0 AND 100),

                             bucket INT[] NOT NULL DEFAULT '{}',

                             start_date TIMESTAMP NULL,
                             end_date TIMESTAMP NULL,

                             status TEXT NOT NULL,

                             CHECK (
                                 array_length(bucket, 1) IS NULL
                                     OR array_length(bucket, 1) <= 100
                                 )
);