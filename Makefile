SHELL			:= /bin/bash
APP_NAME		:= equationlabs-php-cli
APP_MODULE		:= github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator
BUILD_VERSION	:= development # It will be overridden by release process
BUILD_DATE		:= $(shell date +%Y-%m-%dT%H:%M:%S)
BUILD_ARCH		:= $(shell uname -m)
BUILD_OS		:= $(shell uname -s | tr A-Z a-z)

help:
	@echo "${GREEN}-------------- USAGE  --------------------------------------${RESET}"
	@echo "$ make ${GREEN}target${RESET} [options] "
	@echo "${GREEN}-------------- Available Targets ---------------------------${RESET}"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build: ##	Build the project
	@echo "${GREEN}Building the project...${RESET}"
	@echo "${GREEN}Building for OS ${BUILD_OS} and ARCH ${BUILD_ARCH}...${RESET}"
	@go mod tidy
	@GOOS=${BUILD_OS} GOARCH=${BUILD_ARCH} go build \
		-ldflags "-X main.buildVersion=${BUILD_VERSION} -X main.buildDate=${BUILD_DATE}" \
		-o ${APP_NAME} main.go
	@echo "${GREEN}Build completed successfully!${RESET}"

images: ##	BUild docker image for development
	@echo "${GREEN}Building docker image for OS ${BUILD_OS} and ARCH ${BUILD_ARCH}...${RESET}"
	@docker build -t ${APP_NAME}:latest -f Dockerfile .

test: ##	Run the tests with coverage
	@echo "${GREEN}Running tests...${RESET}"
	@go test -cover ./... | tee coverage.txt || make .failed
	@echo "${GREEN}Tests completed! Check coverage.txt for details.${RESET}"


# Helper target to handle build failure
.failed:
	@echo "${RED}Build failed! Please check the output for errors.${RESET}"
	exit 1