build:
	docker-compose build denet-backend

run:
	docker-compose up denet-backend --force-recreate

test:
	go test -v ./...

migrate_up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' down

swag:
	swag init -g cmd/main.go

database:
	sudo docker run --name=denet-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres