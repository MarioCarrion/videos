# Go Tools: For database schema migrations

_Published Video - TBD_

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

igrate down

```
migrate -path migrations/ -database postgres://user:password@localhost:5432/dbname?sslmode=disable down
```

## Problem

_TBD_
