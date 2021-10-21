# README

This folder includes three examples covering:

1. SQL Injection example, 
1. Preventing SQL Injection example with static queries, and
1. Preventing SQL Injection example when creating dynamic queries.

Each example uses the `pgx` API directly, only for PostgreSQL, and the `database/sql` API that should be applicable to any database driver.

## PostgreSQL Instructions

Run the PostgreSQL Docker container:

```
docker run \
  -d \
  --name security-dbs \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5432:5432 \
  postgres:12.5-alpine
```

Using `psql`:

```
docker exec -it security-dbs psql -U user dbname
```

Create the following table:

```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
  id         UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  name       VARCHAR NOT NULL,
  birth_year INT NOT NULL DEFAULT 1900,
  is_admin   BOOL NOT NULL DEFAULT false
);
```

Insert a few demo users:

```
INSERT INTO users(name, birth_year, is_admin) VALUES  ('Mario', 1901, true), ('Wario', 1950, false), ('Xario', 1970, false);
```
