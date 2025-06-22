# Variables
SCHEMA_DIR := .
OUTPUT_FILE := router/supergraph.graphql
CONFIG_FILE := router/supergraph-config.yaml
ROVER_IMAGE := worksome/rover:latest

# Default target
all: stop compose-supergraph start

# Compose the supergraph using Rover CLI in Docker
compose-supergraph:
	@echo "🧩 Composing supergraph with Rover CLI..."
	docker run --rm \
		-e APOLLO_ELV2_LICENSE=accept \
		-v "$(SCHEMA_DIR)":/workdir \
		-w /workdir \
		$(ROVER_IMAGE) \
		supergraph compose --config ./$(CONFIG_FILE) > $(OUTPUT_FILE)
	@echo "✅ Supergraph composed and written to $(OUTPUT_FILE)"

# Start all docker containers
start:
	@echo "🧩 Starting all containers..."
	docker-compose up -d --build
	@echo "✅ All containers up!"

# Stopping all docker containers
stop:
	@echo "🧩 Stopping all containers..."
	docker-compose down
	@echo "✅ All containers down!"
