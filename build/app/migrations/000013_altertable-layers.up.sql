ALTER TABLE layers
    ADD CONSTRAINT uq_layers_namespace_id_name
        UNIQUE (namespace_id, name);