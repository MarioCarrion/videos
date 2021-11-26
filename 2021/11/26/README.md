# Relational Databases - Transactions in PostgreSQL

[Link Episode](https://youtu.be/Bi_VnV4HrFQ)

* Using standard library via [database/sql](stdlib/) with `pgx` and
* Using **only** the [jackc/pgx](thirdparty/) package.

Database Schema used:

```sql
CREATE TABLE users(
  name       VARCHAR  NOT NULL,
  birth_year SMALLINT NULL DEFAULT 0
);
```
