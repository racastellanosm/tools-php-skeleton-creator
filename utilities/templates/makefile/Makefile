SHELL				:= /bin/bash
PROJECT_NAME		:= project_name_placeholder
DOCKER_COMPOSE		:= docker compose -p $(PROJECT_NAME)

help:
	@echo "${GREEN}-------------- USAGE  --------------------------------------${RESET}"
	@echo "$ make ${GREEN}target${RESET} [options] "
	@echo "${GREEN}-------------- Available Targets ---------------------------${RESET}"
	@grep -E '^[a-zA-Z_0-9%-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "${GREEN}%-20s${RESET} %s\n", $$1, $$2}'

build: ## Build the images
	@$(DOCKER_COMPOSE) down --remove-orphans -v --rmi local || $(MAKE) .cleanup
	@$(DOCKER_COMPOSE) build --pull || $(MAKE) .cleanup
start: ## Start the stack
	@$(DOCKER_COMPOSE) up -d || $(MAKE) .cleanup
stop: ## Stop the stack
	@$(DOCKER_COMPOSE) down --remove-orphans -v
unit-test:	## Execute Unit Tests within the docker image
	@$(DOCKER_COMPOSE) run --rm app /bin/sh -c "php bin/phpunit"
	@$(MAKE) stop || $(MAKE) .cleanup


## Helpers Functions
.cleanup:
	@$(DOCKER_COMPOSE) down --remove-orphans -v
	@echo "${RED}Stack stopped with error code $1${RESET}"
	@exit 1
.is-stack-running:
	@if [ -z "$$(docker ps -q -f name=app)" ]; then echo "${RED}The app is not running...${RESET}"; exit 1; fi