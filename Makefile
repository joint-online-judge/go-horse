.PHONY: dev clean pre-commit test build swag pre-dev

APP_NAME = go-horse
BUILD_DIR = $(PWD)/build

dev:
	DEBUG=1 fiber dev -p "make pre-dev" -D docs -e "env"

pre-dev: swag
	go mod tidy

clean:
	rm -rf ./build

pre-commit:
	pre-commit run --all-files

test: clean critic
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

swag:
	swag init
	pre-commit run --files ./docs/swagger.* || true
