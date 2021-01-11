# Go Tools: For database schema migrations

* [Published Video](https://youtu.be/EavdaeUmn64)
* [Blog post](https://mariocarrion.com/2021/01/10/golang-tools-for-database-schema-migrations.html)

1. Install [`direnv`](https://github.com/direnv/direnv)

## Setup

Using docker run

```
docker run \
  --rm \
  --name postgres-migration \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_USER=user \
  -e POSTGRES_DB=dbname \
  postgres:12.5-alpine
```
Then

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
```

Migrate up

```
migrate -path migrations/ -database postgres://user:password@localhost:5432/dbname?sslmode=disable up
```

Migrate down

```
migrate -path migrations/ -database postgres://user:password@localhost:5432/dbname?sslmode=disable down
```

Create migration

```
migrate create -ext sql -dir migrations/ books_add_column_pages
```

## Problem

How to implement Database Schema Migrations?
