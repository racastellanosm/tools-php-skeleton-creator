# CLI tool to create a PHP Skeleton

[![Pipelines](https://github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/actions/workflows/pull_request.yaml/badge.svg)](https://github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/actions/workflows/pull_request.yaml) [![Release](https://github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/actions/workflows/release.yaml/badge.svg)](https://github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/actions/workflows/release.yaml)

helps developers manage php projects scaffolding from local to production environments, it uses `symfony/skeleton` or
`slim/slim-skeleton` as base layout and includes DDD + CQRS basic configuration and files.

## Installation & Usage

### Via Homebrew

This CLI tool can be installed with the folling command (works for `MacOS` & `Linux` ):

 ```bash
brew install equationlabs-php-cli
equationlabs-php-cli help
equationlabs-php-cli <command> [command options] [arguments]
```

### Via Docker

If you don't want to install anything in your machine, you can use the ready-to-use `docker` image to work with the
tool.

> Remember to mount the volume on the path where you want to create the skeleton.

 ```bash
docker pull ghcr.io/equationlabs/equationlabs-php-cli:${version}
docker run --rm -w /app --rm ghcr.io/equationlabs/equationlabs-php-cli:${version} ${command} ${options} ${arguments}
```

> CLI tool brew and docker package is managed with the GoReleaser [Tool][1].

## Signature Verification

EquationLabs PHP CLI binaries are signed using [cosign][2], which is part of [sigstore][3].
Signatures can be verified as follows (OS and architecture omitted for clarity):

```console
$ COSIGN_EXPERIMENTAL=1 cosign verify-blob --signature equationlabs-php-cli.sig equationlabs-php-cli
tlog entry verified with uuid: "$uuid" index: "$index"
Verified OK
```

The above uses the (currently experimental) [keyless signing][4] method.
Alternatively, one can verify the signature by also providing the certificate:

```console
$ cosign verify-blob --cert equationlabs-php-cli.pem --signature equationlabs-php-cli.sig equationlabs-php-cli
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
composer.json
composer.lock
Dockerfile
docker-compose.yml
Makefile
README.md
````

[1]: https://goreleaser.com

[2]: https://github.com/SigStore/cosign

[3]: https://www.sigstore.dev/

[4]: https://github.com/sigstore/cosign/blob/main/KEYLESS.md