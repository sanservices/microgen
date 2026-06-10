# Microservice generator using cookiecutter-golang

## Installation

First, get Cookiecutter.

```console
$ pip install cookiecutter
```

Alternatively, you can install `cookiecutter` with homebrew:

```console
$ brew install cookiecutter
```

# ***Requirements***


**Before even beginning to use the template you must install the following tools:**


# Protobuffers
- [protobuf](https://protobuf.dev/getting-started/gotutorial/)
   ``` 
      brew install protobuf
   ```
- [buf](https://github.com/bufbuild/buf) ***protobuf manager***
   ``` 
      brew install buf
   ```

## Swagger generation
Install the following to generate swagger.
   ``` 
      go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
   ```



## Run

Finally, to run it based on this template, type:

```console
$ cookiecutter https://github.com/sanservices/microgen --checkout v3.4.0 --directory="microgrpcgen"
```

or if you cloned the repository

```console
$ cookiecutter microgen --directory="nanogen | microgen | microgrpcgen | micromcpgen"
```

The available generators are:

- `microgen` — REST microservice (Echo + uber/fx)
- `microgrpcgen` — gRPC + REST (grpc-gateway) microservice
- `micromcpgen` — gRPC + REST + MCP in a single module (based on `microgrpcgen`)
- `nanogen` — Knative serverless function

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

## What the generated projects include

The microservice generators (`microgen`, `microgrpcgen`, `micromcpgen`) scaffold a project with:

- a layered, interface-driven domain package wired together with [uber/fx](https://github.com/uber-go/fx);
- structured logging via `apilogger`, plus startup config validation (`settings.yml`, overridable with the `SETTINGS_PATH` env var);
- basic [Datadog](https://docs.datadoghq.com/tracing/) APM tracing, configured through the standard `DD_*` environment variables;
- a `Makefile` (run `make help`) and a multi-stage, non-root [distroless](https://github.com/GoogleContainerTools/distroless) `Dockerfile`.

Optional features — database, cache, Kafka, and (for `micromcpgen`) MCP — are toggled by the prompts, and only the code for the selected features is generated.
