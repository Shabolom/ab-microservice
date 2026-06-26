CREATE TABLE feature_toggles (
                                 id BIGSERIAL PRIMARY KEY,
                                 namespace_id BIGINT NOT NULL,
                                 name TEXT NOT NULL,
                                 rollout_percentage INT NOT NULL,
                                 status TEXT NOT NULL,
                                 created_at TIMESTAMP NOT NULL DEFAULT now(),
                                 updated_at TIMESTAMP NOT NULL DEFAULT now(),
                                 deleted_at TIMESTAMP NULL,

                                 CONSTRAINT fk_feature_toggles_namespace_id
                                     FOREIGN KEY (namespace_id)
                                         REFERENCES namespaces (id)
                                         ON DELETE CASCADE,

                                 CONSTRAINT uq_feature_toggles_namespace_id_name
                                     UNIQUE (namespace_id, name),

                                 CONSTRAINT chk_feature_toggles_rollout_percentage
                                     CHECK (rollout_percentage >= 0 AND rollout_percentage <= 100),

                                 CONSTRAINT chk_feature_toggles_status
                                     CHECK (status IN ('draft', 'active', 'disabled', 'archived'))
);