FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /{{ cookiecutter.root_directory_name }}

# Install go-swagger
RUN dir=$(mktemp -d) && \
    git clone https://github.com/go-swagger/go-swagger "$dir" && \
    cd "$dir" && \
    git checkout v0.27.0 && \
    go install \
    -ldflags "-X github.com/go-swagger/go-swagger/cmd/swagger/commands.Version=$(git describe --tags) -X github.com/go-swagger/go-swagger/cmd/swagger/commands.Commit=$(git rev-parse HEAD)" \
    ./cmd/swagger

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Generate swagger file
RUN swagger generate spec --compact -o internal/api/handlers/handler/swagger/swagger.yml

# Build the Go application
RUN go build -o {{ cookiecutter.root_directory_name }}

#Final
FROM debian:bookworm-slim AS Final

COPY --from=build /{{ cookiecutter.root_directory_name }}/{{ cookiecutter.root_directory_name }} .

EXPOSE 8080

CMD ["./{{ cookiecutter.root_directory_name }}"]