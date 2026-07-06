# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repository is

This is **not** a Go application — it is a collection of four [Cookiecutter](https://cookiecutter.readthedocs.io/) templates that *generate* Go microservices for Sanservices. Every `.go`, `.yml`, `.mod`, `Dockerfile`, etc. under a `{{ cookiecutter.root_directory_name }}/` directory is a **template file**, not source code. They contain Jinja2 syntax (`{{ cookiecutter.module_name }}`, `{% if cookiecutter.use_database == 'y' %}...{% endif %}`) and will **not** compile, lint, or pass `gofmt` in place. Do not try to `go build`/`go test` them from this repo.

The four generators:

| Dir | Generates | Notes |
|-----|-----------|-------|
| `microgen/` | REST microservice (Echo + uber/fx) | Most complete; layered domain architecture |
| `microgrpcgen/` | gRPC + gRPC-Gateway microservice | Same as microgen plus protobuf/`buf` codegen and a REST gateway |
| `micromcpgen/` | gRPC + REST + MCP in one module | Based on `microgrpcgen`; adds an MCP server (`internal/mcp`) exposing the service as MCP tools |
| `nanogen/` | Knative serverless function (`func`) | Lighter; requires the `func` CLI; module is hardcoded `function` |

Each generator is independent and self-contained: its own `cookiecutter.json` (prompts), `hooks/` (Python pre/post-generation scripts), and `{{ cookiecutter.root_directory_name }}/` (the templated project tree).

## Generating a project (how the templates are "run")

```bash
# From a release tag, picking the generator with --directory
cookiecutter https://github.com/sanservices/microgen --checkout v3.4.0 --directory="microgrpcgen"

# From a local clone (run from the parent of the repo)
cookiecutter microgen --directory="microgen"
cookiecutter microgen --directory="microgrpcgen"
cookiecutter microgen --directory="micromcpgen"
cookiecutter microgen --directory="nanogen"
```

Cookiecutter prompts for the keys in that generator's `cookiecutter.json`. The `use_*` flags accept `y`/`n` and drive both Jinja conditionals in the templates and folder deletion in the post-gen hook.

### cookiecutter.json variables per generator

| Generator | Variables |
|-----------|-----------|
| `microgen` / `microgrpcgen` | `root_directory_name`, `module_name`, `service_name`, `use_database`, `use_cache`, `use_kafka` |
| `micromcpgen` | same as above + `use_mcp` |
| `nanogen` | `root_directory_name`, `handler`, `use_database`, `use_cache`, `use_smtp`, `use_sftp` |

nanogen's `handler` key sets the function handler name; it has `use_smtp`/`use_sftp` instead of `use_kafka`.

### Generation hooks (where the real logic lives)

- `*/hooks/post_gen_project.py` — runs in the **generated** project dir. It `shutil.rmtree`s feature folders that weren't selected (`db/`, `internal/kafka`, `redis/`, `smtp/`, `sftp/`), prettifies `config.yml`, then shells out to `go mod tidy`, `gofmt`, and (microgrpcgen) `buf dep update && buf generate`.
- `nanogen/hooks/pre_gen_project.py` — runs `func create --language go ...` first, so the `func` CLI **must** be installed before generating a nanogen project.

When editing a template, remember the matching hook may delete or keep the file based on a `use_*` flag — keep the Jinja conditionals in the template and the deletion logic in the hook in sync.

## Prerequisites for working with the templates

- `cookiecutter` (`pip install cookiecutter` or `brew install cookiecutter`)
- For **microgrpcgen** / **micromcpgen**: `protobuf` + `buf` (`brew install protobuf buf`), plus Go plugins installed by their post-gen hooks: `swag`, `protoc-gen-go`, `protoc-gen-go-grpc`, `protoc-gen-grpc-gateway`, `protoc-gen-openapiv2`
- For **nanogen**: the Knative `func` CLI (see the [kn-functions](https://github.com/sanservices/kn-functions) repo)

## Generated-service architecture (microgen / microgrpcgen)

The generated services use **uber/fx** for dependency injection. `main.go` is the composition root: `fx.Provide(...)` wires the layers and `fx.Invoke(...)` registers routes and `OnStart`/`OnStop` lifecycle hooks. Layers are stitched together purely by fx — there is no manual wiring.

Layered, interface-driven domain design under `internal/{{ service_name }}/`:

- `{{ service_name }}.go` — the **domain contract**: declares `Service`, `Repository`, and (if cache) `Cache` interfaces. This is the seam everything else depends on.
- `service/` — business logic; depends on the `Repository` (and `Cache`) interfaces, never on concrete DB code.
- `repository/` — implementations of `Repository`, one subpackage per backend: `mysql/`, `oracle/`, `sqlite/`, `redis/`. Unselected backends are deleted by the post-gen hook.
- `entity/` — structs mapping raw data from repositories.
- `internal/api/` — Echo server setup (`api.go`), routing, and versioned handlers under `api/v1/` (handler, DTOs, swagger). `healthcheck/` is always present.
- `internal/kafka/` — optional Kafka listener (topics/routes), started in an fx `OnStart` hook.
- `config/` — reads `settings.yml` into a `Settings` struct via `gopkg.in/yaml.v2`; sections are conditional on the `use_*` flags.

Sanservices internal libraries the generated code relies on: `github.com/sanservices/apilogger/v2` (logging), `github.com/sanservices/apicore` (middleware), `github.com/sanservices/kit` (database, kafkalistener).

**microgrpcgen specifics:** `main.go` starts both a gRPC server (`StartGRPCServer`, with reflection enabled) and a REST gRPC-Gateway (`setupGrpcGatewayHandler`) bound to the Echo server. Proto lives in `internal/api/proto/*.proto`; `buf.yaml` + `buf.gen.yaml` (buf v2) drive code generation into `internal/` (pb) and OpenAPI yaml. The handler implements the generated gRPC service interface.

**micromcpgen specifics:** everything microgrpcgen has, plus a third surface — MCP — gated behind a `use_mcp` flag (like `use_database`/`use_cache`/`use_kafka`). `internal/mcp/` holds an MCP server built on the official [`github.com/modelcontextprotocol/go-sdk`](https://github.com/modelcontextprotocol/go-sdk): `mcp.go` constructs the server and serves it over the Streamable HTTP transport on its own port (`MCP.port`, default `8081`); `tools.go` registers each operation as a tool in `registerTools` and the handlers delegate to the shared `service.Service`. It is wired into fx via `mcpserver.New` in `main.go` and started with `go m.StartServer(ctx)` in the `OnStart` hook, so gRPC, REST, and MCP all run concurrently off one `service` layer. All MCP code (import, fx provide, lifecycle param, `config.go` `MCP` struct, `config.yml` section, Dockerfile `EXPOSE`) is wrapped in `{% raw %}{% if cookiecutter.use_mcp == 'y' %}{% endraw %}` conditionals; when `use_mcp != y` the post-gen hook's `remove_mcp()` deletes `internal/mcp` and `go mod tidy` drops the SDK dependency. The template also fixes microgrpcgen's duplicate `log` import and missing `GRPC`/`MCP` `config.yml` sections. The MCP SDK requires Go ≥ 1.23, so `go.mod` and the Dockerfile pin Go 1.23. To add a tool: add a handler in `tools.go` and register it in `registerTools`.

A standalone example MCP client is generated at `internal/mcp/client/main.go`. Run the service first (`make run`), then `go run ./internal/mcp/client` to list tools and call the sample `get_user` tool. Accepts `-endpoint` and `-user` flags.

## Working in a *generated* project (from its README)

These commands apply to an output project, not this repo.

### Makefile targets (run `make help` to see them)

| Target | What it does |
|--------|-------------|
| `make run` | `go run .` |
| `make build` | Build the binary |
| `make test` | `go test ./...` |
| `make lint` | `golangci-lint run` |
| `make fmt` | `gofmt -l -s -w .` |
| `make tidy` | `go mod tidy` |
| `make docker-build` | Build the Docker image |
| `make docker-run` | Run the Docker image on port 8080 |

### Runtime notes

- Requires `settings.yml` (and `settings_test.yml` for tests). Override the path with the `SETTINGS_PATH` env var.
- Datadog APM tracing is pre-wired; configure with the standard `DD_*` environment variables.
- Docker (simplest): `docker build -t <name>:latest .` then `docker run -p 8080:8080 <name>:latest` — the Dockerfile also generates swagger.
- Local hot-reload: install `modd`.
- Healthcheck: `GET localhost:8080/healthcheck`; swagger at `http://localhost:8080/v1/docs`.

## Linting (template `.golangci.yml`)

The generated `.golangci.yml` deliberately relaxes complexity linters (`cyclop`/`gocyclo` max 70, `dupl` threshold 300, `nestif` 4) and **excludes test files** (`run.tests: false`). The `issues.exclude` list suppresses `undefined:` errors for the optional packages (redis, kafka, echo, etc.) that may be removed by the hook.

## Releases

Pushing a tag matching `v*` triggers `.github/workflows/pre_release.yml`, which opens a **draft** GitHub release. Templates are consumed by tag (`--checkout vX.Y.Z`), so version bumps are tag-driven.
