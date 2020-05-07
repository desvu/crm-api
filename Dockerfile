FROM golang:1.13-alpine AS builder

RUN apk add bash ca-certificates git

WORKDIR /application

COPY go.mod.cache go.mod
RUN go mod download

COPY go.mod go.sum ./
RUN go mod download

# Copy all files in currend directiry into home directory
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/crm ./cmd/main.go

FROM alpine:3.9
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /application
COPY --from=builder /application/bin .

ENTRYPOINT /application/crm