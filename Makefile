
# Build commands
build:
	@go build -o bin/api backend/cmd/api/main.go

run:
	@cd backend && go run cmd/api/main.go 
	@cd frontend && npm run dev

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