include app.env
export 

postgres:
	docker run --name postgres12 -p $(DB_PORT):5432 -d -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	docker exec -it postgres12 dropdb $(DB_NAME)

migrateup:
	migrate -path db/migrations -database="$(DB_SOURCE)" up

migratedown:
	migrate -path db/migrations -database="$(DB_SOURCE)" down

sqlc:
	sqlc generate

test:
	go test -count=1 -v ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Yarik7610/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock