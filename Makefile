PROJECT_NAME := "go-lambda-boilerplate-api"
PKG := "${PROJECT_NAME}"

build: dep gen ## Build the binary file
	@go build -i -v ./.../${PROJECT_NAME}

deploy: ## Generate build and zip
	@bin/deploy.sh
