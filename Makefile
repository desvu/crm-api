.PHONY: grpcgen
grpcgen: ## generate protobuf files
	 @protoc pkg/grpc/proto/game.proto --go_out=plugins=grpc:.

.PHONY: up
up: ## build and run service in docker
	docker-compose up --build -d