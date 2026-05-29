CREATE TABLE experiments (
 id BIGSERIAL  PRIMARY KEY,
 name TEXT NOT NULL,
 namespace TEXT NOT NULL,
 rolling_percentage INT NOT NULL
     CHECK (rolling_percentage BETWEEN 0 AND 100)
);