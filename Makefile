.PHONY: grpcgen
grpcgen: ## generate protobuf files
	 protoc pkg/grpc/proto/*.proto --go_out=plugins=grpc:.

.PHONY: swaggergen
swaggergen: ## generate swagger api documentation
	 swagger generate spec -m -o ./swagger.json

.PHONY: up
up: ## build and run service in docker
	docker-compose up --build -d

.PHONY: dev-build-up
dev-build-up: ## build and run service in docker
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/crm ./cmd/main.go
	docker build -f Dockerfile.dev -t p1hub/qilin-crm-api:latest .
	docker-compose up
