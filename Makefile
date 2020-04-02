
.PHONY: gqlgen
gqlgen: ## generate graphql api
	 go run github.com/99designs/gqlgen --verbose --config gqlgen.yml

.PHONY: up
up: ## build and run service in docker
	docker-compose up --build