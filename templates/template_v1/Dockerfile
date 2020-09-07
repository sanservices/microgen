# project argument is set outside
# to be able to use it on both phases, builder and final
ARG project=my_profject

# ------ BUILD STAGE -------
FROM golang:alpine AS builder
ARG project

RUN apk add --no-cache git alpine-sdk

WORKDIR /${project}

# Install go-swagger
RUN dir=$(mktemp -d) && \
    git clone https://github.com/go-swagger/go-swagger "$dir" && \
    cd "$dir" && \
    go install ./cmd/swagger

# Install statik
RUN go get github.com/rakyll/statik

# Install pacakge dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download -x
COPY . .

# Generate swagger file
RUN swagger generate spec -o files/swaggerui/v1/swagger.yml

# Generate swagger binary bundle
RUN statik -f -src=./files/swaggerui/v1 -dest internal/api/v1 -p swagger

# Build application
RUN go build -o ./${project}

# ---------------------------
# ------- FINAL STAGE -------
FROM alpine:3.9
ARG project
ENV entry=${project}

RUN apk update
RUN apk --no-cache add ca-certificates \
    curl \
    gcc

WORKDIR /${project}

# Copy files from builder
COPY --from=builder /${project}/settings.yaml .
COPY --from=builder /${project}/${project} .
COPY --from=builder /${project}/files/swaggerui/v1/index.html ./files/swaggerui/v1/index.html

# Create non-root user for app
RUN adduser -D -g 'gouser' gouser && \
    chown -R gouser:gouser /${project}

USER gouser

HEALTHCHECK --interval=30s \
    --timeout=5s \
    --retries=3 \
    --start-period=1m \
    CMD curl \
    -f http://localhost:8080/healthcheck || exit 1

EXPOSE 8080
ENTRYPOINT ./${entry}

## build image and remove temporary images:
#docker build -t ${project} . && docker image prune --filter label=stage=builder
