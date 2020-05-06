
.PHONY: gqlgen
gqlgen: ## generate graphql api
	 go run github.com/99designs/gqlgen --verbose --config gqlgen.yml

.PHONY: grpcgen
grpcgen: ## generate protobuf files
	 @protoc pkg/grpc/proto/game.proto --go_out=plugins=grpc:.

.PHONY: up
up: ## build and run service in docker
	docker-compose up --build

.PHONY: upd
upd: ## build and run service in docker
	docker-compose up --build -d