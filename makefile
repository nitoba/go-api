# prisma-migrate-dev:
# 	go run github.com/steebchen/prisma-client-go migrate dev


# prisma-generate:
# 	go run github.com/steebchen/prisma-client-go generate


.PHONY: default run build tests docs clean
# Variables
APP_NAME=apis
APP_ENTRY_POINT=./cmd/main.go

# Tasks
default: run

run:
	@go run $(APP_ENTRY_POINT)
run-with-docs:
	@swag init
	@go run $(APP_ENTRY_POINT)
build:
	@go build -o $(APP_NAME) $(APP_ENTRY_POINT)
tests:
	@go test -v ./...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs