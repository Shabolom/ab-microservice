CREATE TABLE IF NOT EXISTS parametergroup (
    id BIGSERIAL PRIMARY KEY,
    percent INTEGER NOT NULL DEFAULT 100,
    experiment_id BIGINT NOT NULL REFERENCES experiments(id) ON DELETE CASCADE,

    CHECK (percent >= 0 AND percent <= 100)
    );