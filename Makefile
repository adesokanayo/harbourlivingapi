postgres:
	docker run -p 5432:5432 --name docker-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it docker-postgres createdb --username=postgres --owner=postgres harbour

dropdb:
	docker exec -it docker-postgres dropdb harbour

migrateup:
	migrate -path ./db/migration -database "postgres://postgres:secret@localhost:5432/harbour?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migration -database "postgres://postgres:secret@localhost:5432/harbour?sslmode=disable" -verbose down

test:
	go test -v -cover ./...
sqlc:
	sqlc generate
gq:
	gqlgen generate

start:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server gq start