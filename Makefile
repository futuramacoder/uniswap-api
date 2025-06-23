# message wrappers
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
APP_NAME=b-wallet-api

run:
	go run cmd/main.go
swagger:
	echo -e "$(OK_COLOR)->> Generating swagger spec$(NO_COLOR)"; \
    	swag init -d "./app/api" -g "api.go" -o "docs"
