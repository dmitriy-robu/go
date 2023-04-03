APP_NAME=api
ROOT_PATH=/app
DOCKER_EXE=docker compose
DOCKER_NETWORK=go-rust-drop-network

.PHONY: docker

build:
	@cd "$(ROOT_PATH)/cmd/$(APP_NAME)" && go build -buildvcs=false -o "$(APP_NAME)"
	@chmod +x "$(ROOT_PATH)/cmd/$(APP_NAME)/$(APP_NAME)"
	@mv "$(ROOT_PATH)/cmd/$(APP_NAME)/$(APP_NAME)" "$(ROOT_PATH)/tmp/$(APP_NAME)"

docker: check-network
	@echo "Starting the Docker Compose stack..."
	$(DOCKER_EXE) up -d

check-network:
	@echo "Checking if the network exists..."
	@docker network inspect $(DOCKER_NETWORK) > /dev/null 2>&1 || (echo "Network does not exist. Please create it." && docker network create $(DOCKER_NETWORK))
	@echo "Network exists."


