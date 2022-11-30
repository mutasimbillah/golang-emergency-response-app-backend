createPostgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

startDb:
	docker start postgres15

stopDb:
	docker stop postgres15

createDb:
	docker exec -it postgres15 createdb --username=root --owner=root er_app

dropDb:
	docker stop postgres15
	docker start postgres15
	docker exec -it postgres15 dropdb er_app

migrateUp:
	migrate --path db/migrations --database "postgres://root:secret@localhost:5432/er_app?sslmode=disable" --verbose up

migrateDown:
	migrate --path db/migrations --database "postgres://root:secret@localhost:5432/er_app?sslmode=disable" --verbose down

server:
	go run main.go

.PHONY: createPostgres runPostgres stopDb startDb createDb dropDb migrateUp migrateDown buildContainer server
#A phony target is one that is not really the name of a file. 
#It is just a name for some commands to be executed when you make an explicit request. 
#There are two reasons to use a phony target: to avoid a conflict with a file of the same name, and to improve performance.