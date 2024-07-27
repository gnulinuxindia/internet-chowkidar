# Internet Chowkidar
This repository contains the code for internet chowkidar, a web application that allows users to monitor any blocks, bans, or restrictions on the internet by their ISP. The application is built using Go and PostgreSQL.

## Features
- **Logging**: Structured logging with [slog](https://pkg.go.dev/golang.org/x/exp/slog)
- **Configuration**: Configuration management with [koanf](https://pkg.go.dev/github.com/knadh/koanf)
- **HTTP Server**: HTTP server with [net/http](https://pkg.go.dev/net/http)
- **Database**: DB ORM with [ent](https://entgo.io)
- **Task Runner**: Task runner with [just](https://just.systems)

## Requirements
- [Go >= 1.21.5](https://go.dev/doc/install)
- [redocly](https://github.com/Redocly/redocly-cli) (generate API documentation)
- [ent](https://entgo.io/docs/getting-started) (DB ORM)
- [goose](https://github.com/pressly/goose?tab=readme-ov-file#install) (DB migrations tool)
- [ogen](https://ogen.dev/docs/intro) (generate APIs)
- [wire](https://github.com/google/wire) (dependency injection)
- [just](https://just.systems/docs/en/getting-started) (task runner)

## Quickstart
1. Clone the repository
2. Install the dependencies
```bash
go mod tidy
```
4. Set up local environment variables using the `.env.example` file.
```bash
cp .env.example .env
```
5. Run codegen
```bash
just # yes, that's it, nothing else
```
6. Run the application
```bash
just run
```

## Codegen
Necessary for generating APIs, DB schemas, dependency injection, etc. A detailed explanation of the codegen process can be found [here](docs/en/codegen.md).

## Devcontainers
This repository contains a devcontainer configuration for Visual Studio Code. The devcontainer is pre-configured with all the necessary tools and extensions to start developing right away. A detailed explanation of the devcontainer can be found [here](docs/en/devcontainer.md).

## Migrations
Refer to the [migrations](docs/en/migrations.md) documentation for more information on how to create and apply migrations.