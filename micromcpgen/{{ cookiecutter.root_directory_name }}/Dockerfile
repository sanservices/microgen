# syntax=docker/dockerfile:1

# ---------- build stage ----------
FROM golang:1.23 AS build

WORKDIR /app

# Download dependencies first so this layer is cached unless go.mod/go.sum change.
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Build the application. Generated protobuf code is committed, so no codegen here.
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /app/bin/{{ cookiecutter.root_directory_name }} .

# ---------- final stage ----------
# Distroless provides CA certificates and a non-root user, and contains no shell
# or package manager. cc-debian12 includes the C runtime needed by CGO backends
# such as sqlite.
FROM gcr.io/distroless/cc-debian12:nonroot AS final

WORKDIR /

# Application binary and default configuration. Override the config at runtime
# with the SETTINGS_PATH env var or by mounting a different settings.yml.
COPY --from=build /app/bin/{{ cookiecutter.root_directory_name }} /{{ cookiecutter.root_directory_name }}
COPY --from=build /app/settings.yml /settings.yml

EXPOSE 8080
EXPOSE 50051
{% if cookiecutter.use_mcp == 'y' %}EXPOSE 8081{% endif %}

USER nonroot:nonroot

ENTRYPOINT ["/{{ cookiecutter.root_directory_name }}"]
