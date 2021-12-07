# Learning Relational Databases in Go using MySQL

[Video](https://youtu.be/gT2ztW4j66A)

Using standard library via [database/sql](stdlib/) with the [go-sql-driver/mysql/](https://github.com/go-sql-driver/mysql/)

Database Schema used:

```sql
CREATE TABLE users(
  name       VARCHAR(10) NOT NULL,
  birth_year SMALLINT NULL DEFAULT 0,
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

## Using docker for running MySQL

```
docker run \
  --rm \
  -d \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=dbname \
  -p 3306:3306 \
  mysql:8.0.27
```
