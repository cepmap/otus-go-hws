BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

run: build
	$(BIN) --config ./configs/config.yaml
build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)
run-postgres:
	docker run --rm -p 15432:5432  --name  go-calendar-postgres -e POSTGRES_PASSWORD=postgres -d postgres
stop-postgres:
	docker stop go-calendar-postgres

version: build
	$(BIN) version

test:
	go test -race ./internal/... ./pkg/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.63.4

lint: install-lint-deps
	golangci-lint run ./...

migrate-up:
	goose --dir=./migrations postgres "host=127.0.0.1 port=15432 user=postgres password=postgres dbname=postgres sslmode=disable" up
migrate-down:
	goose --dir=./migrations postgres "host=127.0.0.1 port=15432 user=postgres password=postgres dbname=postgres sslmode=disable" down

.PHONY: build run build-img run-img version test lint
