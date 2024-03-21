AuthProto:
	@cd services/auth && \
	PATH=$$PATH:$(shell go env GOPATH)/bin && \
	protoc --go_out=. --go-grpc_out=. auth.proto

JWTProto:
	@cd services/JWT && \
	PATH=$$PATH:$(shell go env GOPATH)/bin && \
	protoc --go_out=. --go-grpc_out=. jwt.proto