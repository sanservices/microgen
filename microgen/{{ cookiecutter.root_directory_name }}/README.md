# Service

## Overview

This is a REST microservice generated from the Sanservices Golang microservice template (Echo + uber/fx).

## Local setup

### Docker

The easiest way to run this service is to build and run its container image. The multi-stage `Dockerfile` compiles the
binary and packages it into a small, non-root [distroless](https://github.com/GoogleContainerTools/distroless) image
together with its default `settings.yml`.

```
make docker-build   # docker build -t {{ cookiecutter.root_directory_name }}:latest .
make docker-run     # run the image and publish the service port(s)
```

Or directly:

```
docker build -t {{ cookiecutter.root_directory_name }}:latest .
docker run --rm -p 8080:8080 {{ cookiecutter.root_directory_name }}:latest
```

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

## Observability

The service emits [Datadog](https://docs.datadoghq.com/tracing/) APM traces out of the box. The tracer is started at
boot with the service name and version from `settings.yml`, and incoming HTTP requests are traced automatically via
Datadog's echo middleware. Configure it through the standard Datadog environment variables — no code changes required:

| Variable | Purpose | Default |
|----------|---------|---------|
| `DD_ENV` | Deployment environment (e.g. `prod`, `staging`) | _(unset)_ |
| `DD_SERVICE` | Service name override | `Service.name` from `settings.yml` |
| `DD_VERSION` | Service version override | `Service.version` from `settings.yml` |
| `DD_AGENT_HOST` | Datadog agent host | `localhost` |
| `DD_TRACE_AGENT_PORT` | Datadog agent trace port | `8126` |

If no Datadog agent is reachable the tracer simply no-ops, so it is safe to leave enabled in local development.

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
