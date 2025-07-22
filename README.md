# EquationLabs CLI tool for PHP Skeleton Creator

This CLI tool helps you create a PHP skeleton for your projects and a basic directory structure (using DDD & Hexagonal principles), it's based in the symfony/cli tool and uses the symfony framework as a base for the project.

Everything runs on Docker, so you don't need to install anything on your machine, just run the commands in your terminal and the tool will take care of the rest.

## Requirements
- Docker & Docker Compose installed on your machine.
- composer 

## Installation
This CLI tool can be installed with:

- MacOS: `brew install equationlabs/tap/equationlabs-skeleton`
- Linux: `curl -sSL https://raw.githubusercontent.com/equationlabs/skeleton/main/install.sh | bash`

> CLI tools is made using Go and can be used on any platform regardless of the operating system.

Every release of the tool is published on GitHub and Brew automatically with the pipeline, so you can always get the latest version.

## Usage
After installing the CLI tool, you can create a new PHP skeleton project by running the following command in your terminal:

```bash
$ equationlans-cli create-php-cqrs-keleton --name <package-name>
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
docker-compose.volumes.yml.dist
Makefile
README.md
```


## Additional Components
After completion, you will have a basic PHP project structure ready to go using DDD & Hexagonal structure and different symfony components apart from the basic ones.

Here you have the additional components that will be installed in the project:

- symfony/messenger
- symfony/cache
- symfony/http-client
- symfony/uuid
- phpunit/phpunit
- symfony/phpunit-bridge


## Features
- Generates a `composer.json` file with basic metadata using the provided package.
- Use symfony framework API skeleton for microservice as base for the project, with CQRS and DDD structure.
- Creates a `README.md` file with the project name and description.
- Creates a basic directory structure with `src`, `tests`, and `config` directories.
- Install dependencies using Composer.
- You can input the version of symfony you want to use, or it will default to the latest stable version.
- A README file with basic information about the project, structure and usage taken from this README file.
