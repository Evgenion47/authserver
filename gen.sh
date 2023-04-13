protoc -I ./api --go_out ./pkg/api --go_opt paths=source_relative \
                --go-grpc_out ./pkg/api --go-grpc_opt paths=source_relative \
                --grpc-gateway_out ./pkg/api --grpc-gateway_opt paths=source_relative \
                --grpc-gateway_opt generate_unbound_methods=true \
                --openapiv2_out ./pkg/api \
                api/api.proto
