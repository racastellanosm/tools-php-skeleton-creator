SHELL				:= /bin/bash
PROJECT_NAME		:= project_name_placeholder
DOCKER_COMPOSE		:= docker compose -p $(PROJECT_NAME)
DEBUG_PARAMETERS	:= -dxdebug.mode=debug -dxdebug.start_with_request=yes -dxdebug.discover_client_host=0 -dxdebug.client_port=9003 -dxdebug.client_host=host.docker.internal -dxdebug.log=build/xdebug.log
COVERAGE_PARAMETERS	:= -dxdebug.mode=coverage

help:
	@echo "${GREEN}-------------- USAGE  --------------------------------------${RESET}"
	@echo "$ make ${GREEN}target${RESET} [options] "
	@echo "${GREEN}-------------- Available Targets ---------------------------${RESET}"
	@grep -E '^[a-zA-Z_0-9%-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "${GREEN}%-20s${RESET} %s\n", $$1, $$2}'

build: ## Build the images
	@echo "🔄 Building image..."
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
	@if [ -z "$(debug)" ]; then \
		echo "👀 Running Unit Tests..."; \
		$(DOCKER_COMPOSE) run --rm --no-deps app sh -c "php ${COVERAGE_PARAMETERS} bin/phpunit --testsuite unit" || $(MAKE) .cleanup; \
	else \
		echo "🐛 Running Unit Tests with Xdebug..."; \
		$(DOCKER_COMPOSE) run --rm --no-deps app sh -c "php ${DEBUG_PARAMETERS} bin/phpunit --no-coverage --group debug --testsuite unit" || $(MAKE) .cleanup; \
	fi
integration-test: ## Execute Integration Tests with all the stack running in docker
	@$(MAKE) start
	@$(MAKE) migrations
	@if [ -z "$(debug)" ]; then \
  		echo "👀 Running Integration Tests..."; \
  		$(DOCKER_COMPOSE) exec app sh -c "php ${COVERAGE_PARAMETERS} bin/phpunit --testsuite integration" || $(MAKE) .cleanup; \
  	else \
  	  	echo "🐛 Running Integration Tests with Xdebug..."; \
  		$(DOCKER_COMPOSE) exec app sh -c "php ${DEBUG_PARAMETERS} bin/phpunit --no-coverage --group debug --testsuite integration" || $(MAKE) .cleanup; \
  	fi
	@$(MAKE) stop
functional-test: ## Execute Functional Tests with all the stack running in docker
	@$(MAKE) start
	@$(MAKE) migrations
	@if [ -z "$(debug)" ]; then \
  		echo "👀 Running Functional Tests..."; \
		$(DOCKER_COMPOSE) exec -e APP_ENV=test api sh -c "php ${COVERAGE_PARAMETERS} bin/phpunit --testsuite functional" || $(MAKE) .cleanup; \
	else \
	  	echo "🐛 Running Functional Tests with Xdebug..."; \
		$(DOCKER_COMPOSE) exec -e APP_ENV=test api sh -c "php ${DEBUG_PARAMETERS} bin/phpunit --no-coverage --group debug --testsuite functional" || $(MAKE) .cleanup; \
	fi
	@$(MAKE) stop
migrations: ## Execute doctrine migrations
	@echo "🔄 Executing migrations..."
	@$(DOCKER_COMPOSE) run --rm --no-deps app sh -c "bin/console d:m:m --no-interaction" || $(MAKE) .cleanup
	@echo "✅ Migrations executed successfully!"

## Helpers Functions
.cleanup:
	@$(DOCKER_COMPOSE) down --remove-orphans -v
	@echo "‼️ Stack finished with errors..."
	@exit 1

.PHONY: build start stop unit-test integration-test functional-test migrations