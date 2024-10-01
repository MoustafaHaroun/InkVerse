# build commands
build:
	@go build -o bin/api cmd/api/main.go

run:
	@go run cmd/api/main.go

# docker commands
docker-up:
	docker compose -f ./docker-compose.yaml up --detach

docker-down:
	docker compose down

# database commands
db-start:
	docker compose start db

db-stop:
	docker compose stop db 