import os
import shutil


# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

# Commands executed (in order) once the project tree has been pruned.
COMMANDS = [
   # 1. Make sure that the go.mod file matches the source code in the module
   "echo Running go mod tidy...",
   "go mod tidy -v",

   # 2. Format Go source code according to the official Go formatting guidelines
   "echo Running go fmt...",
   "gofmt -l -s -w .",
]


def remove_database():
   """Removes the database layer (connection + concrete repositories) if unused."""
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "internal/db"))
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "internal/repository/mysql"))
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "internal/repository/oracle"))


def remove_sftp():
   """Removes the sftp implementation if it isn't going to be used."""
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "internal/sftp"))


def remove_smtp():
   """Removes the smtp implementation if it isn't going to be used."""
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "internal/smtp"))


def remove_cache():
   """Removes the redis cache implementation if it isn't going to be used."""
   shutil.rmtree(os.path.join(PROJECT_DIRECTORY, "internal/repository/redis"))


def main():
   if '{{ cookiecutter.use_database }}'.upper() != 'Y':
      remove_database()

   if '{{ cookiecutter.use_sftp }}'.upper() != 'Y':
      remove_sftp()

   if '{{ cookiecutter.use_smtp }}'.upper() != 'Y':
      remove_smtp()

   if '{{ cookiecutter.use_cache }}'.upper() != 'Y':
      remove_cache()

   # Execute commands
   for command in COMMANDS:
      os.system(command)


if __name__ == "__main__":
   main()
