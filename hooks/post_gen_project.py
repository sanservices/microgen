import os
import shutil
import yaml

from cookiecutter.main import cookiecutter


# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

COMMANDS = [
   # 1. Make sure that the go.mod file matches the source code in the module
   "echo Running go mod tidy...",
   "go mod tidy -v"

   # 2. Is used to format Go source code according to the official Go formatting guidelines.
   "&& echo Running go fmt...",
   "gofmt -l -s -w ."
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
   
   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "cache"
   ))

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
