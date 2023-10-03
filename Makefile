PHONY: generate-v1
generate-v1:
	mkdir -p protocol/
	protoc candlestick/v1/*.proto \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=./protocol \
	--go-grpc_out=./protocol

	protoc prices/v1/*.proto \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=./protocol \
	--go-grpc_out=./protocol

