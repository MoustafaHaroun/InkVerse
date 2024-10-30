
# Build commands
build:
	@go build -o bin/api backend/cmd/api/main.go

run:
	@cd backend && go run cmd/api/main.go 

migration:
	@cd backend && go run cmd/migrate/main.go

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