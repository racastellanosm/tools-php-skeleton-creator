# CLI tool to create a PHP Skeleton

[![Pipelines](https://github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/actions/workflows/pull_request.yaml/badge.svg)](https://github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/actions/workflows/pull_request.yaml) [![Release](https://github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/actions/workflows/release.yaml/badge.svg)](https://github.com/Equation-Labs-I-O/eqlabs-tools-php-skeleton-creator/actions/workflows/release.yaml)

helps developers manage php projects scaffolding from local to production environments, it uses `symfony/skeleton` os
`slim/slim-skeleton` as base layout and includes DDD + CQRS basic configuration and files.

## Installation

This CLI tool can be installed with:

- `MacOS` & `Linux` via Homebrew

 ```bash
brew install equationlabs-php-cli
```

> CLI tools is made using Go and can be used on any platform regardless of the operating system with the GoReleaser
> Tool[1].

Every release of the tool is published on GitHub and Brew automatically with the pipeline, so you can always get the
latest version.

## Usage

After installing the CLI tool, you can create a new PHP skeleton project by running the following command in your
terminal:

> Usage: `equationlabs-cli [global options] <command> [command options] [arguments]`

```bash
$ equationlabs-php-cli help
```

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
Makefile
README.md
```

[1]: https://goreleaser.com

[2]: https://github.com/SigStore/cosign

[3]: https://www.sigstore.dev/

[4]: https://github.com/sigstore/cosign/blob/main/KEYLESS.md