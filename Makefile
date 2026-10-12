PROTO_DIR := api

THIRD_PARTY_DIR := third_party

GEN_DIR := internal/generated

OPENAPI_DIR := api

PROTO_FILES := api/go_load.proto

.PHONY: proto check

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
check:
	gofmt -l .
	go vet ./...
