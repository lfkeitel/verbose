.PHONE: lint test vet full generate

default: lint vet test

lint:
	golint ./...

test:
	go test -v ./...

vet:
	go vet ./...

generate:
	go generate

full: generate lint vet test
