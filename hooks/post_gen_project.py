import os

COMMANDS = [
    "ls -l -R -1",
    # 1. Format the project
    "gofmt -s -w .",

    # 2. Install dependencies
    "echo Installing dependencies...",
    "go mod download",

    # 3. Making sure that the go.mod file matches the source code in the module
    "echo Running go mod tidy...",
    "go mod tidy",

    "echo Project ready"
]

for command in COMMANDS:
    os.system(command)
