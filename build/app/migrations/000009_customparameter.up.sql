CREATE TABLE IF NOT EXISTS customparameter (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    namespace_id BIGINT NOT NULL REFERENCES namespaces(id) ON DELETE CASCADE,
    type VARCHAR(9) NOT NULL DEFAULT 'STRING',

    UNIQUE (name, namespace_id),

    CHECK (type IN ('STRING', 'INT', 'SEMVER', 'DATE'))
    );