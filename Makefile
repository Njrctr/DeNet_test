build:
	sudo docker-compose build denet-api

run:
	sudo docker-compose up denet-api --force-recreate

test:
	go test -v ./...

migrate.up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

migrate.down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' down

swag:
	swag init -g cmd/app/main.go

database.up:
	sudo docker run --name=denet-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
database.down:
	sudo docker stop denet-db