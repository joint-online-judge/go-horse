.PHONY: all dev pre-dev clean pre-commit install-dev test build swag

APP_NAME = go-horse
BUILD_DIR = $(PWD)/build

all: dev

dev:
	mkdir -p build
	air

pre-dev:
	go mod tidy
	make build
	docker cp ./build/go-horse joj2-go-horse:.
	docker restart joj2-go-horse
	docker logs -f joj2-go-horse

clean:
	rm -rf ./build

pre-commit:
	pre-commit run --all-files

install-dev:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest
	python3 -m pip install -U pre-commit
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	pre-commit install

test: clean
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

BUILT_AT := $(shell date +'%F %T %z')
GO_VERSION := $(shell go version)
GIT_AUTHOR := $(shell git show -s --format='format:%aN <%ae>' HEAD)
GIT_COMMIT := $(shell git log --pretty=format:"%H" -1)
VERSION := dev
FLAGS := "-s -w \
-X 'github.com/joint-online-judge/go-horse/pkg/configs.BuiltAt=$(BUILT_AT)' \
-X 'github.com/joint-online-judge/go-horse/pkg/configs.GoVersion=$(GO_VERSION)' \
-X 'github.com/joint-online-judge/go-horse/pkg/configs.GitAuthor=$(GIT_AUTHOR)' \
-X 'github.com/joint-online-judge/go-horse/pkg/configs.GitCommit=$(GIT_COMMIT)' \
-X 'github.com/joint-online-judge/go-horse/pkg/configs.Version=$(VERSION)'"
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags=$(FLAGS) -o $(BUILD_DIR)/$(APP_NAME) .

swag:
	swag init
