# Build stage
FROM golang:1.17.2 AS build

WORKDIR /{{.ServiceName}}
ARG APP_PORT
COPY . .
RUN set -ex; \
    go mod tidy && \
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o {{.ServiceName}} ./cmd/app

# Final image with no source code
FROM alpine:3.14

WORKDIR /
COPY --from=build /{{.ServiceName}}/{{.ServiceName}} .
EXPOSE ${APP_PORT}
ENTRYPOINT /{{.ServiceName}}