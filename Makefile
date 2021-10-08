.PHONY: all generate build test vet clean

all: clean artictl

generate:
	go generate ./...

build: generate
	go build ./...

test: build
	go test ./...

vet: test
	go vet ./...

artictl: build
	go build -o artictl ./cmd/artictl/main.go

clean:
	rm -f artictl
