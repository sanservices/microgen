import os
import shutil
import yaml


# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

COMMANDS = [
   # 1. Go get swaggo
   "echo Running go get swaggo...",
   "go install github.com/swaggo/swag/cmd/swag@latest",
   "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest",
   "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
   # 2. Make sure that the go.mod file matches the source code in the module
   "&& echo Running go mod tidy...",
   "go mod tidy -v",

   # 3. Is used to format Go source code according to the official Go formatting guidelines.
   "&& echo Running go fmt...",
   "gofmt -l -s -w .",

   #4. protobuf generation
   "buf dep update",
   "buf generate"


]

def remove_database():
   """
   Removes folder needed for the database layer if it isn't going to be used
   """
   
   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "db"
   ))

def remove_cache():
   """
   Removes folder needed for redis cache if it isn't going to be used
   """

   # Define the directory name you want to search for
   target_dir = "redis"

   # Initialize an empty list to store the paths of matching directories
   found_dirs = []

   # Walk through the directory tree starting from the root_dir
   for dirpath, dirnames, filenames in os.walk(PROJECT_DIRECTORY):
       if target_dir in dirnames:
           found_dirs.append(os.path.join(dirpath, target_dir))
           shutil.rmtree(os.path.join(
               dirpath, target_dir
            ))

   # Print the paths of found directories
   for found_dir in found_dirs:
    print("Found directory:", found_dir)

def remove_kafka():
   """
   Removes folder needed for kafka if it isn't going to be used
   """
   
   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "internal/kafka"
   ))


def prettify_config():
   """
   Prettify yml config file to a standard format
   """

   # Read the YAML file
   with open('config.yml', 'r') as file:
      data = yaml.safe_load(file)

   # Write the YAML file with prettified formatting
   with open('config.yml', 'w') as file:
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

   # 4. Prettify config.yml file
   prettify_config()

   # 4. Execute commands
   for command in COMMANDS:
      os.system(command)

if __name__ == "__main__":
   main()
