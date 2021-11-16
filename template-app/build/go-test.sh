#! /bin/sh

GOFILES=./...
COVERAGE_THRESHOLD=$1

go mod tidy
go test -count=1 -p=1 -covermode=atomic -coverprofile=coverage.out $GOFILES
if [ $? -ne 0 ]; then
  echo "Tests failed"; \
  exit 1; \
fi;

go tool cover -html coverage.out -o coverage.html
coverage=$(go tool cover -func coverage.out | grep 'total:' | awk '{print int($3)}')

echo "The overall coverage is $coverage%. Look at coverage.html for details."

if [ $coverage -lt $COVERAGE_THRESHOLD ]; then
  echo "The coverage $coverage% is below the accepted threshold $COVERAGE_THRESHOLD%."; \
  exit 1; \
fi;