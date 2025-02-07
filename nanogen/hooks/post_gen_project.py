import os
import shutil


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
   COMMANDS.append("gopls imports -w internal/repository/repository.go")

   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "internal/db"
   ))
   
   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "internal/repository/mysql"
   ))
   
   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "internal/repository/oracle"
   ))

def remove_sftp():
   """
   Add kn function files
   """
   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "internal/sftp"
   ))

def remove_smtp():
   """
   Add kn function files
   """
   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "internal/smtp"
   ))
   
def remove_cache():
   """
   Add kn function files
   """
   shutil.rmtree(os.path.join(
      PROJECT_DIRECTORY, "internal/repository/redis"
   ))

def main():   
   if '{{ cookiecutter.use_database }}'.upper() != 'Y':
      remove_database()
   
   if '{{ cookiecutter.use_sftp }}'.upper() != 'Y':
      remove_sftp()
   
   if '{{ cookiecutter.use_smtp }}'.upper() != 'Y':
      remove_smtp()
      
   if '{{ cookiecutter.use_cache }}'.upper() != 'Y':
      remove_cache()
      
   
   # 6. Execute commands
   for command in COMMANDS:
      os.system(command)
   
if __name__ == "__main__":
   main()
