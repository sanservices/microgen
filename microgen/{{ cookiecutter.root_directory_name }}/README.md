# Service

## Overview

This is a service which is based on goproposal for version 2 of the Sanservices boilerplate Golang microservice project.

## Local setup

### Docker

The easiest way to run this application is to build a docker image from it and run that image as a container. This would
handle everything from compiling the executable to generating the swagger documentation without much need to know how
any of that works.

To do that run

```
docker build -t {{ cookiecutter.module_name }}-service:latest .
```

and then

```
docker run --name {{ cookiecutter.module_name }}-service -p 8080:8080 {{ cookiecutter.module_name }}-service:latest
```

in the project's root directory.

The downside to using this method of compiling and executing the app is that it makes debugging a little more
complicated. If you're using VSCode for development, you can find information on how to get that right over here

- https://code.visualstudio.com/docs/containers/debug-common.

### Non-Docker setup

The project ships with a `Makefile` and a `settings.yml` pre-filled with sensible
local defaults, so you can build and run it without any extra tooling.

#### Setup steps:

1. Review `settings.yml` and adjust it for your environment (ports, database, cache,
   ...). The config file path can be overridden at runtime with the `SETTINGS_PATH`
   environment variable.

2. Use the `Makefile` for the common tasks (run `make help` to list them all):

   ```
   make run     # run the service locally (go run .)
   make build   # build the binary
   make test    # run the test suite
   make lint    # run golangci-lint
   make tidy    # sync go.mod / go.sum with the source
   ```

   Or run it directly without make:

   ```
   go run .
   ```

3. Try out the endpoints. The service exposes a health check at
   `GET localhost:8080/healthcheck`. Use an API client such as
   [Postman](https://www.postman.com/downloads/) to exercise the rest of the API.

## Project architecture

`settings.yml`

*️ Service configuration, read at startup and mapped onto typed structs by the `config` package. The file path can be
overridden with the `SETTINGS_PATH` environment variable.

`/internal`

*️ The main package of the project, where everything specific to its domain and functionality lives.

`/internal/api`

*️ Everything related to the service API — router, routes, handlers, middleware (filters), etc.

`/internal/api/v1/swagger`

*️ The embedded swagger UI assets, served by the docs handler via `//go:embed`.

`/internal/{{ cookiecutter.service_name }}`

*️ All business logic and data-repository interaction for the {{ cookiecutter.service_name }} domain.

`/internal/{{ cookiecutter.service_name }}/entity`

*️ Structs matching the raw data fetched from each repository, before it is transformed/enriched in the API layer.

`/internal/{{ cookiecutter.service_name }}/{{ cookiecutter.service_name }}.go`

*️ The domain contract — the `Service`, `Repository` (and optional `Cache`) interfaces implemented by the packages within.

`/internal/{{ cookiecutter.service_name }}/repository`

*️ Implementations of the `Repository` interface, one sub-package per backend (mysql/oracle/sqlite/redis).

`/internal/{{ cookiecutter.service_name }}/service`

*️ The business logic: responsible for data transformation/enrichment and choosing where to fetch data from.
