import os
from pathlib import Path
import shutil

COMMANDS = [
    # 2. Install dependencies
    "echo Installing dependencies...",
    "go mod download",

    # 3. Making sure that the go.mod file matches the source code in the module
    "echo Running go mod tidy...",
    "go mod tidy"
]


def check_if_kafka_is_needed():
    """
    Checks if the project has conditional packages.
    """

    include_kafka = '{{ cookiecutter.include_kafka }}'

    if include_kafka == 'Yes':
        return

    path = os.getcwd() + "/internal/kafka"

    shutil.rmtree(path)


for command in COMMANDS:
    os.system(command)

check_if_kafka_is_needed()
os.system("gofmt -s -w .")
print("Project ready")
