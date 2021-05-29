all: protoc build run

protoc:
	@echo "Generating protobuf code"
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false  --go-grpc_out=. --go-grpc_opt=paths=source_relative api/manpass.proto

build:
	@echo "Compiling main function"
	@go build cmd/manpass/main.go

run:
	@echo "Running the app"
	@ DB_FILE=manpass_db MIGRATIONS_PATH=file://./migrations ./main