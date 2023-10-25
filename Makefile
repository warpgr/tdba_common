PHONY: generate-v1
generate-v1:
	mkdir -p go/protocol/
	protoc candlestick/v1/*.proto \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=./go/protocol \
	--go-grpc_out=./go/protocol

	protoc prices/v1/*.proto \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=./go/protocol \
	--go-grpc_out=./go/protocol

	protoc orders/v1/*.proto \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=./go/protocol \
	--go-grpc_out=./go/protocol

	protoc wallet/v1/*.proto \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=./go/protocol \
	--go-grpc_out=./go/protocol



