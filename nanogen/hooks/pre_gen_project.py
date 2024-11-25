import os
import subprocess

def initialize_func():
    try:

        # Run `func` command to initialize the project with func
        subprocess.run(["func", "create", "--language", "go", "../{{ cookiecutter.root_directory_name }}"], check=True)
    except subprocess.CalledProcessError as e:
        print(f"Failed to initialize func: {e}")
        exit(1)

if __name__ == "__main__":
    initialize_func()