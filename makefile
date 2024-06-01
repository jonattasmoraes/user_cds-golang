.PHONY: default run build test docs	clean

APP_NAME=app_go

default: build

run:
	@docker compose up
build:
	@go build -o $(APP_NAME) cmd/app/main.go
test:
	@go test ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs