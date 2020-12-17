proto:
	protoc --go-grpc_out=. user/delivery/grpc/user.proto
	protoc --go_out=. user/delivery/grpc/user.proto