PROTO_DIR := api

THIRD_PARTY_DIR := third_party

GEN_DIR := internal/generated

OPENAPI_DIR := api

PROTO_FILES := api/go_load.proto

WIRE_DIR := internal/wiring

.PHONY: proto check wire

proto:
	protoc \
		-I $(PROTO_DIR) \
		-I $(THIRD_PARTY_DIR) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(GEN_DIR) \
		--grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=$(OPENAPI_DIR) \
		$(PROTO_FILES)
wire:
	wire gen ./$(WIRE_DIR)

check:
	gofmt -l .
	go vet ./...
