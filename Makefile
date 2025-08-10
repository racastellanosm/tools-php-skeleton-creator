SHELL			:= /bin/bash
APP_NAME		:= equationlabs-php-cli
APP_MODULE		:= github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator
BUILD_VERSION	:= development # It will be overridden by release process
BUILD_DATE		:= $(shell date +%Y-%m-%dT%H:%M:%S)

help:
	@echo "${GREEN}-------------- USAGE  --------------------------------------${RESET}"
	@echo "$ make ${GREEN}target${RESET} [options] "
	@echo "${GREEN}-------------- Available Targets ---------------------------${RESET}"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build: ##	Build the project
	@echo "${GREEN}Building the project...${RESET}"
	@go build \
		-ldflags "-X main.buildVersion=${BUILD_VERSION} -X main.buildDate=${BUILD_DATE}" \
		-o bin/${APP_NAME} main.go
	@echo "${GREEN}Build completed successfully!${RESET}"

test: ##	Run the tests with coverage
	@echo "${GREEN}Running tests...${RESET}"
	@go test -cover ./... | tee coverage.txt || make .failed
	@echo "${GREEN}Tests completed! Check coverage.txt for details.${RESET}"


# Helper target to handle build failure
.failed:
	@echo "${RED}Build failed! Please check the output for errors.${RESET}"
	exit 1