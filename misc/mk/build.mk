
build: bin
	go build -o bin/ ./...

bin:
	rm -rf bin
	mkdir bin

protoc:
	protoc -I=api/proto/ \
	       --go_out=api/ \
	       --go_opt=paths=source_relative \
	       --go-grpc_out=api/ \
	       --go-grpc_opt=paths=source_relative \
	       --experimental_allow_proto3_optional \
	       api/proto/*.proto
