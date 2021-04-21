# project argument is set outside
# to be able to use it on both phases, builder and final
ARG project=app

# ------ BUILD STAGE -------
FROM golang:1.15.7-alpine3.12 AS builder
ARG project

LABEL stage=builder

# Basic libraries
RUN apk add --no-cache git alpine-sdk

WORKDIR /${project}

# Download and install go-swagger
RUN dir=$(mktemp -d) && \
    git clone https://github.com/go-swagger/go-swagger "$dir" && \
    cd "$dir" && \
    go build cmd/swagger/swagger.go && \
    cp swagger /usr/local

# Install statik
RUN go get github.com/rakyll/statik

# Copy project list of dependencies
COPY ./go.mod .
COPY ./go.sum .

# Install pacakge dependencies
RUN go get -d -v ./...

# Copy project files
COPY . .

# Generate swagger file
RUN /usr/local/swagger generate spec -w ./internal/api/v1 -w ./ -o ./files/swaggerui/v1/swagger.yml --scan-models

# Generate swagger binary bundle
RUN statik -dest internal/api/v1 -p swagger -src=./files/swaggerui/v1 -f

# Build application
RUN go build -o ./${project}

# --------------------------------------------------
# ------- FINAL STAGE -------
FROM alpine:3.10
ARG project
ENV entry=${project}

LABEL maintainer=devops@sandals.com name=${project}

RUN apk update
RUN apk --no-cache add ca-certificates \
    curl

# Create non-root user for app
RUN adduser -D -g 'gouser' gouser && \
    mkdir -p /${project} && \
    chown -R gouser:gouser /${project}

USER gouser

HEALTHCHECK --interval=30s \
    --timeout=5s \
    --retries=3 \
    --start-period=1m \
    CMD curl \
    -f http://localhost:8080/healthcheck || exit 1

# Copy files from builder
# settings should be passed as docker-config
# COPY --from=builder /${project}/settings/settings.yml ${project}/settings/settings.yml
COPY --from=builder /${project}/${project} ${project}/
COPY --from=builder --chown=gouser:gouser /${project}/files/ /${project}/files/

EXPOSE 8080
WORKDIR /${project}
ENTRYPOINT ./${entry}

## build image and remove temporary images:
# docker build -t ${project} . && docker image prune --filter label=stage=builder