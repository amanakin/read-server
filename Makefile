all: bin/read-server ;

bin/read-server:
	go build -mod=vendor -v -o ./bin/read-server ./cmd/read-server

test:
	go test -mod=vendor -v ./cmd/read-server

vendor:
	go mod vendor

clean:
	rm -fv ./bin/read-server

.PHONY: all bin/* test vendor clean
