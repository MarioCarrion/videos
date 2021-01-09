CREATE EXTENSION "uuid-ossp";

CREATE TABLE authors (
  uuid       uuid      PRIMARY KEY DEFAULT uuid_generate_v4(),
  author     VARCHAR   NOT NULL,
  name       VARCHAR   NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
