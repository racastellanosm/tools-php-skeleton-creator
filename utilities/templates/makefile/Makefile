SHELL				:= /bin/bash
PROJECT_NAME		:= project_name_placeholder
DOCKER_COMPOSE		:= docker compose -p $(PROJECT_NAME)

help:
	@echo "${GREEN}-------------- USAGE  --------------------------------------${RESET}"
	@echo "$ make ${GREEN}target${RESET} [options] "
	@echo "${GREEN}-------------- Available Targets ---------------------------${RESET}"
	@grep -E '^[a-zA-Z_0-9%-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "${GREEN}%-20s${RESET} %s\n", $$1, $$2}'

build: ## Build the images
	@echo "🔄 Building images..."
	@$(DOCKER_COMPOSE) down --remove-orphans -v --rmi local || $(MAKE) .cleanup
	@$(DOCKER_COMPOSE) build --pull || $(MAKE) .cleanup
	@echo "✅ Images built successfully!"
start: ## Start the stack
	@echo "🔄 Starting stack..."
	@$(DOCKER_COMPOSE) up -d || $(MAKE) .cleanup
	@echo "🚀 Stack started successfully!"
stop: ## Stop the stack
	@echo "🔄 Stopping stack..."
	@$(DOCKER_COMPOSE) down --remove-orphans -v || $(MAKE) .cleanup
	@echo "✅ Stack stopped successfully!"
unit-test:	## Execute Unit Tests within the docker image
	@echo "👀 Running Unit Tests..."
	@$(DOCKER_COMPOSE) run --rm --no-deps app /bin/sh -c "php bin/phpunit" || $(MAKE) .cleanup
	@$(MAKE) stop
integration-test: ## Execute Integration Tests with all the stack running in docker
	@$(MAKE) start
	@echo "👀 Running Functional Tests..."
	@$(DOCKER_COMPOSE) exec -it app /bin/sh -c "php bin/phpunit" || $(MAKE) .cleanup
	@$(MAKE) stop
functional-test: ## Execute Functional Tests with all the stack running in docker
	@$(MAKE) start
	@echo "👀 Running Functional Tests..."
	@$(DOCKER_COMPOSE) exec -it app /bin/sh -c "php bin/phpunit" || $(MAKE) .cleanup
	@$(MAKE) stop

## Helpers Functions
.cleanup:
	@$(DOCKER_COMPOSE) down --remove-orphans -v
	@echo "‼️ Stack stopped with error..."
	@exit 1