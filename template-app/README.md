# {{.ServiceName}}

## Libraries Used

| Library                             | Version | Use
|:------------------------------------|:--------|:-----------
| github.com/gorilla/mux              |v1.7.4   | Router
| github.com/prometheus/client_golang |v1.6.0   | Prometheus
| github.com/prometheus/client_model  |v0.2.0   | Prometheus
| github.com/rs/zerolog               |v1.18.0  | Logger
| github.com/stretchr/testify         |v1.4.0   | Test
| github.com/golangci/golangci-lint   |v1.26.0  | Go linter

## Makefile
This repository utilises Makefile to automate testing, building and packaging the service code. Following are the commands 
you can run using the Makefile:
* `make setup`: Installs all the dependencies required to build, test and package the service.
* `make build`: Builds the go server code and outputs a linux binary.
* `make test`: Builds and starts a local docker instance of the service. Runs the unit and integration tests, outputs the 
coverage both to the cli and as a html file. The task is wired to fail if the test coverage is below a set COVERAGE_THRESHOLD.  
* `make teardown`: Stops the running instances of the service.
* `make docker.push`: Push the build docker image to ECR registry.
* `make docker.login`: Login to AWS ECR registry.
* `make fix`: Runs the default linter(golangci-lint) with a --fix flag.
* `make lint`: Runs the defaut linter(golangci-lint)
* `make all`: Runs the linter, buids the code, run tests and package the service.

## Endpoints
* `/health`: Standard http health endpoint 
* `/metrics`: Prometheus metrics endpoint

## Tests
This repository has necessary unit and integration tests. The tests run in a docker environment using [docker-compose](docker-compose.yaml). 
The external dependencies like db / queue can also be brought up in the same [docker-compose](docker-compose.yaml) for integration tests. 
The test stage also outputs a coverage report which is created using the [atomic](https://blog.golang.org/cover) cover mode. 
The test stage is wired to fail if the total coverage of all the packages fall below a certain threshold (_COVERAGE_THRESHOLD_), 
which can be set in the [Makefile](Makefile).