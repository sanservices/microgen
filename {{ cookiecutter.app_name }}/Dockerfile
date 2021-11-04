# project argument is set outside
# to be able to use it on both phases, builder and final
ARG project=app

# ================================
# ======= BUILD STAGE ============
FROM golang:1.16-alpine AS builder
ARG swagger_host
ARG project

RUN apk add --no-cache git alpine-sdk upx

WORKDIR /${project}

# Install go-swagger
RUN dir=$(mktemp -d) && \
    git clone https://github.com/go-swagger/go-swagger "$dir" && \
    cd "$dir" && \
    go install ./cmd/swagger

# Install pacakge dependencies.
COPY ["go.mod", "go.sum", "./"]
RUN go mod download -x

# Copy the rest of the project,
# if there are files to ignore
# please add them to the .dockerignore.
COPY . .

# Generate swagger file
RUN swagger generate spec -o internal/api/v1/swagger/swagger.yml

# Update the swagger host if needed
RUN if [[ -n "$swagger_host" ]] ; then \
    sed -i "s/localhost:8080/${swagger_host}/g" internal/api/v1/swagger/swagger.yml; \
    fi

RUN go build -o ./${project}

# Compress binary file
RUN upx ./${project}

# ============================
# ======= FINAL STAGE ========
FROM alpine:3.10
ARG project
ENV entry=${project}

RUN apk update
RUN apk add --no-cache ca-certificates curl

WORKDIR /${project}

# Create non-root user for app
RUN adduser -D -g 'appuser' appuser && \
    chown -R appuser:appuser /${project}

# Copy files from builder and bundler
COPY --from=builder /${project}/${project} .

USER appuser

# Make sure that the port correlates
# with what is configured in settings.yml.
HEALTHCHECK --interval=30s \
    --timeout=5s \
    --retries=3 \
    --start-period=1m \
    CMD curl \
    -f http://localhost:8080/healthcheck || exit 1

ENTRYPOINT ./${entry}