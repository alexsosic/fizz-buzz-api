GO ?= go
GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell $(GO) list ./...)
GOFILES := $(shell find . -name "*.go" -type f -not -path './vendor/*')
DOCKER_COMPOSE_FILES = -f docker-compose.yaml
DOCKER_IMAGE = alexsosic/fizz-buzz-api
VERSION = master

.PHONY: all
all: fmt lint vet test

.PHONY: build
build:
	$(GO) build -o bin/fizz-buzz-api .

.PHONY: test
test:
	$(GO) test -cover -coverprofile=coverage.txt $(PACKAGES) && echo "\n==>\033[32m Ok\033[m\n" || exit 1

.PHONY: test-short
test-short:
	$(GO) test -cover -coverprofile=coverage.txt --short $(PACKAGES) && echo "\n==>\033[32m Ok\033[m\n" || exit 1

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: lint
lint:
	for pkg in ${PACKAGES}; do \
		golint -set_exit_status $$pkg || GOLINT_FAILED=1; \
	done; \
	[ -z "$$GOLINT_FAILED" ]

.PHONY: vet
vet:
	$(GO) vet $(PACKAGES)

.PHONY: clean
clean:
	$(GO) clean -modcache -x -i ./...
	find . -name coverage.txt -delete
	rm bin/*

.PHONY: docker-build
docker-build:
	docker build . -t $(DOCKER_IMAGE):$(VERSION)

.PHONY: docker-pull
docker-pull:
	docker-compose $(DOCKER_COMPOSE_FILES) pull

.PHONY: docker-start
docker-start:
	docker-compose $(DOCKER_COMPOSE_FILES) up -d

.PHONY: docker-stop
docker-stop:
	docker-compose $(DOCKER_COMPOSE_FILES) down
