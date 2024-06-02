.PHONY: default run build test docs	clean

APP_NAME=app_go/cmd/app

default: build

run:
	@docker compose up
run-with-docs:
	@swag init -dir cmd/app
	@docker compose up --build
build:
	@go build -o $(APP_NAME) cmd/app/main.go
test:
	@go test ./...
docs:
	@swag init -dir cmd/app
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs