CREATE TABLE books (
  isbn        VARCHAR   PRIMARY KEY,
  author_uuid uuid      NOT NULL,
  created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),

  FOREIGN KEY (author_uuid) REFERENCES authors (uuid)
);
