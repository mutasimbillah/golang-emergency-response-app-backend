createPostgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
runPostgres:
	docker start postgres15
startdb:
	docker start postgres15

createDb:
	docker exec -it postgres15 createdb --username=root --owner=root er_app

dropDb:
	docker exec -it postgres15 dropdb er_app

migrateUp:
	migrate --path db/migration --database "postgres://root:secret@localhost:5432/er_app?sslmode=disable" --verbose up

migrateDown:
	migrate --path db/migration --database "postgres://root:secret@localhost:5432/er_app?sslmode=disable" --verbose down

server:
	go run main.go

.PHONY: createPostgres runPostgres startDb createDb dropDb migrateUp migrateDown server