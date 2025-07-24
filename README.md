# CLI tool to create a PHP Skeleton

helps developers manage php projects scaffolding from local to production environments, it uses symfony/skeleton as base layout and includes DDD + CQRS basic configuration and files.


## Requirements
You need to have the following tools installed on your machine to use this CLI tool:

- Docker & Docker Compose. (Preferred version is 20.10.0 or higher)
- composer (PHP Package Manager).

## Installation
This CLI tool can be installed with:

- `MacOS` & `Linux` via Homebrew: `brew install equationlabs/tap/equationlabs-cli`

> CLI tools is made using Go and can be used on any platform regardless of the operating system.

Every release of the tool is published on GitHub and Brew automatically with the pipeline, so you can always get the latest version.

## Usage
After installing the CLI tool, you can create a new PHP skeleton project by running the following command in your terminal:

> Usage: `equationlabs-cli [global options] <command> [command options] [arguments]`

```bash
$ equationlabs-cli help
$ equationlabs-cli new <application-name>
```

After installation the skeleton provides a makefile to run dockerized command to ease the development process, you can run the following command to start the project:

```bash
$ make build-dev # Install dependencies and build the project for development
$ make build # Install dependencies and build the project for production
$ make start # Start the project
$ make unit-test  # Run the unit tests
$ make integration-test  # Run the integration tests
$ make functional-test  # Run the functional tests
$ make stop # Stop the project
```

## File Structure
After completion, you will have a basic PHP project structure ready to go using DDD & Hexagonal structure and different symfony components apart from the basic ones.

After running the CLI tool, your project will have the following structure, so you can start developing your application right away, without wasting time setting up the basic structure or configuring the project:

```yaml
.github/
├── workflows/
├── ci.yml
├── release.yml
bin/
├── console
config/
├── packages/
├── services.php
├── routes.php
public/
├── index.php
src/
├── Application/
├── Domain/
├── Infrastructure/
tests/
├── Unit/
├── Integration/
├── Functional/
.gitignore
.dockerignore
composer.json
composer.lock
Dockerfile
docker-compose.yml
docker-compose.override.yml
Makefile
README.md
```