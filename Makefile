all: server ;

client:
	go build -mod=vendor -v -o ./bin/client ./cmd/client

server:
	go build -mod=vendor -v -o ./bin/read-server ./cmd/read-server

test:
	go test -mod=vendor -v ./cmd/read-server
	go test -mod=vendor -v ./cmd/client

vendor:
	go mod vendor

clean:
	rm -fv ./bin/read-server
	rm -fv ./bin/client

.PHONY: all bin/* test vendor clean
