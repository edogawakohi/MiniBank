postgres:
	docker run --name minibank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:18.1-alpine
createdb:
	docker exec -it minibank createdb --username=root --owner=root mini_bank
dropdb:
	docker exec -it minibank dropdb mini_bank
migrateup:
	migrate -path db/migration -database "postgres://root:123@localhost:5432/mini_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://root:123@localhost:5432/mini_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server