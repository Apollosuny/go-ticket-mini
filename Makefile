PROTO_DIR=api/proto

proto:
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/ticket.proto

run-ticket:
	go run ./cmd/ticket-service

run-gateway:
	go run ./cmd/api-gateway