protoc:
	protoc --proto_path=proto --go_opt=paths=source_relative --go_out=./testapi --go-grpc_out=./testapi --go-grpc_opt=paths=source_relative ./proto/testapi.proto

bench:
	go test -bench=. -benchmem -benchtime 20s

lint:
	golangci-lint run

godoc:
	godoc -http=:6060


