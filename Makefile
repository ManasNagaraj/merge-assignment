BUILD_DIR ?= build/bin
DOCKER_COMPOSE_FILE ?= ./docker/docker-compose.yml



.PHONY: build run
build:
	rm -rf $(BUILD_DIR)
	mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/apiserver ./cmd/server
run: build
	./$(BUILD_DIR)/apiserver

migrate-up:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up
