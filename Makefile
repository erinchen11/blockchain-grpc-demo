run-protoc-grpc-go:
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto

proto-gen-go:
	protoc --go_out=. --go_opt=paths=source_relative *.proto

run-client:
	go run client/main.go

run-server:
	go run server/main.go
run:
	go run main.go