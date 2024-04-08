.PHONY: default run build test clean docs

APP_NAME = "openings"

default: run

run: 
	@swag init
	@go run main.go
build:
	@go build -o ./bin/$(APP_NAME) main.go
test:
	@go test -v ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -Rf ./docs