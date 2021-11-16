# http-bootstrapper
This is a code generator that uses `go templates` to generate a bootstrap code for a go http server.

## Usage
Generate go http server code and push it to the specified repository
```shell script
SERVICE_NAME=<service_name> make templatize.go
```

## Makefile
This repository utilises Makefile to automate testing, building and packaging the service code. Following are the commands you can run using the Makefile:
* `make templatize.go`: Runs the go templatizing code to generate a go http server bootstrap code.
repository.
* `make clean`: Removes the generated go http server bootstrap code.
