postgres:
	docker run --name postgres12 -p 5433:5432 -d -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migrations -database="postgres://root:secret@localhost:5433/simple_bank?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database="postgres://root:secret@localhost:5433/simple_bank?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -count=1 -v ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test