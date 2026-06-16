CREATE TABLE IF NOT EXISTS customparametercondition (
    id BIGSERIAL PRIMARY KEY,
    parameter_id BIGINT NOT NULL REFERENCES customparameter(id) ON DELETE CASCADE,
    value TEXT NOT NULL,
    condition VARCHAR(16) NOT NULL,
    parameter_group_id BIGINT NOT NULL REFERENCES parametergroup(id) ON DELETE CASCADE,

    CHECK (
              condition IN (
              '=',
              '<>',
              '>',
              '<',
              '>=',
              '<=',
              'IN',
              'NOT IN',
              'CONTAINS',
              'NOT CONTAINS',
              'BETWEEN',
              'NOT BETWEEN',
              'ONE OF',
              'NOT ONE OF'
                           )
    )
    );