# Microservice generator using cookiecutter-golang

## Installation

First, get Cookiecutter.

```console
$ pip install cookiecutter
```

Alternatively, you can install `cookiecutter` with homebrew:

```console
$ brew install cookiecutter
```

# ***Requirements***


**Before even beginning to use the template you must install the following tools:**


# Protobuffers
- [protobuf](https://protobuf.dev/getting-started/gotutorial/)
   ``` 
      brew install protobuf
   ```
- [buf](https://github.com/bufbuild/buf) ***protobuf manager***
   ``` 
      brew install buf
   ```

## Swagger generation
Install the following to generate swagger.
   ``` 
      go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
   ```



## Run

Finally, to run it based on this template, type:

```console
$ cookiecutter https://github.com/sanservices/microgen --checkout v3.4.0 --directory="microgrpcgen"
```

or if you cloned the repository

```console
$ cookiecutter microgen --directory="nanogen | microgen"
```

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.
