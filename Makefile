clean:
	rm -rf grpc/*

.PHONY: server
server:
	go run cmd/server/*.go -port 50051

.PHONY: client
client:
	go run cmd/client/main.go -address 0.0.0.0:50051

.PHONY: proto
proto:
	protoc --go_out=grpc \
	--go_opt=paths=source_relative \
	--go-grpc_out=grpc --go-grpc_opt=paths=source_relative \
	--proto_path=proto proto/*.proto