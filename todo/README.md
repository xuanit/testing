# Requirements
* Install GRPC gateway. https://github.com/grpc-ecosystem/grpc-gateway
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```
* [Mockery](https://github.com/vektra/mockery)
```
// The command should be run with go module disabled
export GO111MODULE=off
go get github.com/vektra/mockery/.../

// Restore go module
export GO111MODULE=auto

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

// Generate files such as mocks, *.pb.go, *.pb.gw.go and *.swagger.json
make generate

```

Please read [Pact Go](https://github.com/pact-foundation/pact-go) for more details about contract testing
