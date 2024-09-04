ENV_FILE := .env
APP_NAME := go-budget-api

# Run the Go application locally (without Docker)
build:
	cd cmd/gin && \
	go mod tidy && \
	go build -o $(PWD)/$(APP_NAME)

run-local: build
	$(PWD)/$(APP_NAME)