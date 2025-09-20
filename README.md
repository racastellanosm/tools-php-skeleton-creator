# CLI tool to create a PHP Skeleton

[![Pipelines](https://github.com/racastellanosm/tools-php-skeleton-creator/actions/workflows/pull_request.yaml/badge.svg)](https://github.com/racastellanosm/tools-php-skeleton-creator/actions/workflows/pull_request.yaml) [![Release](https://github.com/racastellanosm/tools-php-skeleton-creator/actions/workflows/release.yaml/badge.svg)](https://github.com/racastellanosm/tools-php-skeleton-creator/actions/workflows/release.yaml)

helps developers manage php projects scaffolding from local to production environments, it uses `symfony/skeleton` or
`slim/slim-skeleton` as base layout and includes DDD + CQRS basic configuration and files.

By default, comes with a `Dockerfile` with a base image with [RoadRunner][5] + `MySQL` or `PostgreSQL` extension support depending
on the option you pass when use the tool

## Installation & Usage

### Via Homebrew

This CLI tool can be installed with the following command (works for `MacOS` & `Linux` ):

 ```bash
brew install php-skeleton-creator-cli
php-skeleton-creator-cli help
php-skeleton-creator-cli <command> [command options] [arguments]
```

### Via Docker

If you don't want to install anything in your machine, you can use the ready-to-use `docker` image to work with the
tool.

> Remember to mount the volume on the path where you want to create the skeleton.

 ```bash
docker pull ghcr.io/racastellanosm/php-skeleton-creator-cli:${version}
docker run -v $PWD:/build -w /build -t --rm ghcr.io/racastellanosm/php-skeleton-creator-cli:${version} ${command} ${options} ${arguments}
```

> CLI tool brew and docker package is managed with the GoReleaser [Tool][1].

## Signature Verification

PHP Skeleton Creator CLI binaries are signed using [cosign][2], which is part of [sigstore][3].
Signatures can be verified as follows (OS and architecture omitted for clarity):

```console
$ COSIGN_EXPERIMENTAL=1 cosign verify-blob --signature php-skeleton-creator-cli.sig php-skeleton-creator-cli
tlog entry verified with uuid: "$uuid" index: "$index"
Verified OK
```

The above uses the (currently experimental) [keyless signing][4] method.
Alternatively, one can verify the signature by also providing the certificate:

```console
$ cosign verify-blob --cert php-skeleton-creator-cli.pem --signature php-skeleton-creator-cli.sig php-skeleton-creator-cli
Verified OK
```

## CLI tool skeleton output

### File Structure

After completion, you will have a basic PHP project structure ready to go using DDD & Hexagonal structure and different
symfony components apart from the basic ones.

After running the CLI tool, your project will have the following structure, so you can start developing your application
right away, without wasting time setting up the basic structure or configuring the project:

```yaml
bin/
├── console
config/
├── packages/
├── services.php
├── routes.php`
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
.env.dist
composer.json
composer.lock
Dockerfile
docker-compose.yml
docker-compose.override.yml
Makefile
README.md
````

[1]: https://goreleaser.com

[2]: https://github.com/SigStore/cosign

[3]: https://www.sigstore.dev/

[4]: https://github.com/sigstore/cosign/blob/main/KEYLESS.md

[5]: https://roadrunner.dev