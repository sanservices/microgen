# Overview

microgen is a `CLI` tool to help create micro services or just a rest api in go. It uses an opiniated structure that tries to implement clean architecture.

# Dependencies
In order to properlly work with the generated code, you will need to install some dependencies. Follow the links to get installation instructions.

- [Go-swagger](https://goswagger.io/)
- [Statik](https://github.com/rakyll/statik)
- [Modd](https://github.com/cortesi/modd)

# Installation

Run `go get -u github.com/san-services/microgen`
Then you can restart your terminal and start using it.

# Usage

To create a new project go to the directory where you want to save it and run:

```sh
microgen new project --name "my awesome project" --module path/to/your/repository
```

After that a new folder will be generated on your current location. Navigate into it, and run `modd` it will download dependencie, generate the swagger files and start the server, you can check the generated swagger at [http://localhost:8080/v1/docs](http://localhost:8080/v1/docs)
