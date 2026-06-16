CREATE TABLE layers (
                        id BIGSERIAL PRIMARY KEY,
                        namespace_id BIGINT NOT NULL
                            REFERENCES namespaces(id)
                                ON DELETE CASCADE,

                        name TEXT NOT NULL,
                        description TEXT NOT NULL
);

CREATE INDEX idx_layers_namespace_id
    ON layers(namespace_id);