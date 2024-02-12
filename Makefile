# Makefile for managing the Docker Compose setup

# Project variables
PROJECT_NAME=myapp
DOCKER_COMPOSE_FILE=docker-compose.yml

DB_USER=username
DB_PASS=password
DB_NAME=db
DB_PORT=5432 # Default PostgreSQL port


# Docker Compose commands
DC=docker-compose -p $(PROJECT_NAME) -f $(DOCKER_COMPOSE_FILE)

# Makefile targets
.PHONY: up down build rebuild start stop logs ps

# Bring up the containers
up:
	$(DC) up -d

# Take down the containers
down:
	$(DC) down

# Build or rebuild services
build:
	$(DC) build

rebuild:
	$(DC) build --no-cache

# Start services
start:
	$(DC) start

# Stop services
stop:
	$(DC) stop

# View output from containers
logs:
	$(DC) logs

# List containers
ps:
	$(DC) ps

create-db:
	@echo "Creating database $(DB_NAME)..."
	docker exec -it myapp-db createdb -U $(DB_USER) $(DB_NAME)

create-table:
	@echo "Creating messages table if not exists..."
	@docker exec -it myapp-db psql -U $(DB_USER) -d $(DB_NAME) -c "CREATE TABLE IF NOT EXISTS messages (id SERIAL PRIMARY KEY, content TEXT NOT NULL);"
