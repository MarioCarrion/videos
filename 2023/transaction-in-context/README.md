# README

## Install

## Running PostgreSQL container

```
docker run \
    -d \
    -e POSTGRES_DB=dbname \
    -e POSTGRES_HOST_AUTH_METHOD=trust \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_USER=user \
    -p 5432:5432 \
    postgres:16.0-alpine3.18
```
