.PHONY: all dev pre-dev clean pre-commit install-dev test build swag

APP_NAME = go-horse
BUILD_DIR = $(PWD)/build

all: dev

dev:
	mkdir -p build
	air

pre-dev: swag
	go mod tidy
	make build
	docker cp ./build/go-horse joj2-go-horse:/go-horse
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

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) .

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

swag:
	swag init
