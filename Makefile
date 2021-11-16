.PHONY: all
all: env.check clean templatize.go

.PHONY: templatize.go
templatize.go:
	go run main.go $(SERVICE_NAME)
	chmod +x $(SERVICE_NAME)/build/go-test.sh

.PHONE: clean
clean:
	rm -rf ${SERVICE_NAME}

.PHONY: env.check
env.check:
ifeq ($(and $(SERVICE_NAME)),)
	$(error Env variables not found. The env variables you need to pass are: SERVICE_NAME)
endif
