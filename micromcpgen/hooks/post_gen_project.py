import os
import shutil
import yaml


# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

# Commands executed (in order) once the project tree has been pruned.
COMMANDS = [
   # 1. Install the protobuf / gRPC / gateway / OpenAPI code generators
   "echo Installing protobuf code generators...",
   "go install github.com/swaggo/swag/cmd/swag@latest",
   "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest",
   "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
   "go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest",
   "go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest",

   # 2. Generate protobuf, gRPC, gateway and OpenAPI code.
   #    Run before `go mod tidy` so the generated imports can be resolved.
   "echo Generating protobuf code...",
   "buf dep update",
   "buf generate",

   # 3. Make sure that the go.mod file matches the source code in the module
   "echo Running go mod tidy...",
   "go mod tidy -v",

   # 4. Format Go source code according to the official Go formatting guidelines
   "echo Running go fmt...",
   "gofmt -l -s -w .",
]


def remove_dirs_named(names):
   """Recursively delete every directory whose name is in `names`."""
   for dirpath, dirnames, _ in os.walk(PROJECT_DIRECTORY):
      for name in list(dirnames):
         if name in names:
            shutil.rmtree(os.path.join(dirpath, name))
            dirnames.remove(name)  # don't descend into the deleted directory


def remove_database():
   """Removes the database layer (connection + concrete repositories) if unused."""
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "db"))
   remove_dirs_named(("mysql", "oracle", "sqlite"))


def remove_cache():
   """Removes the redis cache implementation if it isn't going to be used."""
   remove_dirs_named(("redis",))


def remove_kafka():
   """Removes the kafka implementation if it isn't going to be used."""
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "internal/kafka"))


def remove_mcp():
   """
   Removes the MCP server if it isn't going to be used. The MCP SDK dependency
   is dropped automatically by `go mod tidy` once the import is gone.
   """
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "internal/mcp"))


def prettify_config():
   """Prettify the settings.yml config file to a standard format."""
   with open("settings.yml", "r") as file:
      data = yaml.safe_load(file)

   with open("settings.yml", "w") as file:
      yaml.dump(data, file, sort_keys=False, indent=2)


def main():
   # 1. Remove database implementation if it is not going to be used
   if '{{ cookiecutter.use_database }}'.upper() != 'Y':
      remove_database()

   # 2. Remove cache implementation if it is not going to be used
   if '{{ cookiecutter.use_cache }}'.upper() != 'Y':
      remove_cache()

   # 3. Remove kafka implementation if it is not going to be used
   if '{{ cookiecutter.use_kafka }}'.upper() != 'Y':
      remove_kafka()

   # 4. Remove MCP server if it is not going to be used
   if '{{ cookiecutter.use_mcp }}'.upper() != 'Y':
      remove_mcp()

   # 5. Prettify the settings.yml file
   prettify_config()

   # 6. Execute commands
   for command in COMMANDS:
      os.system(command)


if __name__ == "__main__":
   main()
