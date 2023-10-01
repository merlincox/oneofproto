protoc:
	protoc --proto_path=proto --go_opt=paths=source_relative --go_out=./testapi --go-grpc_out=./testapi --go-grpc_opt=paths=source_relative ./proto/testapi.proto