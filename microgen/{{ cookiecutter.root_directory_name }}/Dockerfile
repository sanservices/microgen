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

RUN apt update && apt install openssl -y && apt install ca-certificates

# Get the certificate from the server
RUN openssl s_client -connect www.sandals.com:443 < /dev/null 2> /dev/null | \ 
    openssl x509 -inform pem -text | sed -n -e '/BEGIN/,$p' > /etc/ssl/certs/cert.pem

RUN update-ca-certificates

COPY --from=build /{{ cookiecutter.root_directory_name }}/{{ cookiecutter.root_directory_name }} .

EXPOSE 8080

CMD ["./{{ cookiecutter.root_directory_name }}"]