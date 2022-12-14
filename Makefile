clean-pb:
	rm -rf pb/*

.PHONY: test
test:
	go test ./... -race -cover

.PHONY: buildup
buildup:
	docker-compose up --build

.PHONY: stop
stop:
	docker-compose stop

.PHONY: up
up:
	docker-compose up

.PHONY: down
down:
	docker-compose down

.PHONY: client
client:
	go run cmd/client/main.go -address 0.0.0.0:50051

.PHONY: proto
proto:
	protoc --go_out=pb \
	--go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--proto_path=internal/port/grpc/proto internal/port/grpc/proto/*.proto