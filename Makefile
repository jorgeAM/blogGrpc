server/run:
	@go run cmd/server/main.go

client/run:
	@go run cmd/client/main.go

generate/proto:
	@protoc --go_out=plugins=grpc:. ./blogpb/blog.proto