# Simplebank

It should be a big project, but i was doing it by watching tutorials. And i found that tutorials aren't free from some point, so i decided to stop this project.

## Features that have been touched

- Dockerizing postgres
- Use migration scripts for fast db start seed
- Using `gin` as backend framework
- Using `sqlc` package
- Writing unit tests for database methods, mocking database
- Add CI/CD workflow
- Work on Transactions both in theory and practice, make transactional methods for money exchange

## How to run

First, add `.env` file in the root of project

```.env
DB_DRIVER=postgres
DB_PORT=5433
DB_USER=root
DB_PASSWORD=secret
DB_NAME=simple_bank
DB_SOURCE=postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable
HTTP_SERVER_ADDRESS=0.0.0.0:8080
```

Then, call next commands in order

```bash
make postgres
make createdb
make migrateup
make server
```
