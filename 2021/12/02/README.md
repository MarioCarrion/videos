# Relational Databases - Bulk Insertions in PostgreSQL

* Using standard library via [database/sql](stdlib/) with `pgx` (but not really bulk insert!) and
* **Actually** doing bulk insertions using [jackc/pgx](pgx/) package.

Using Docker:

```
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5432:5432 \
  postgres:13.5-alpine3.15
```

Database Schema used:

```sql
CREATE TABLE users(
  first_name VARCHAR  NOT NULL,
  last_name  VARCHAR  NOT NULL
);
```
