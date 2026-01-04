# Load environment variables from .env file
include .env
export

# --- Database ---

postgres:
	docker compose up -d

createdb:
	docker exec -it pokedex_db createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DB)

dropdb:
	docker exec -it pokedex_db dropdb $(POSTGRES_DB)

# --- Migrations ---

migrateup:
	migrate -path sql/schema -database "$(DB_SOURCE)" -verbose up

migratedown:
	migrate -path sql/schema -database "$(DB_SOURCE)" -verbose down

# --- Application ---

run:
	go run main.go