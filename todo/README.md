Install GRPC gateway. https://github.com/grpc-ecosystem/grpc-gateway
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

Please read Makefile for how to start the service and run tests
```
// Run unit tests
make test 

// Run integration tests at consumer side
make integration-test-consumer

// Run integration tests at provider side
make integration-test

// Start the service
make run

```

Please read [Pact Go](https://github.com/pact-foundation/pact-go) for more details about contract testing
