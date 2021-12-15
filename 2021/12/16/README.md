# Relational Databases - Bulk Insertions in MySQL

[Video](https://youtu.be/FaXtQy98V7M)

Using the third party package [go-sql-driver/mysql/](https://github.com/go-sql-driver/mysql/).

Database Schema used:

```sql
CREATE TABLE users(
  first_name VARCHAR(10) NOT NULL,
  last_name  VARCHAR(10) NOT NULL
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
