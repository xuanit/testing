test:
	go test -v ./...

install:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

set-up:
	export PATH=$(PATH):$(GOPATH)/bin

generate:
	go generate --all

run:
	docker-compose down
	docker-compose up -d
	go run ./server/main.go

integration-test:
	docker-compose down && docker-compose up -d  && go test -count=1 -v -tags=integration ./...
