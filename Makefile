.PHOHY: build, clear

build:
	go build -v ./cmd/apiserver


.PHONY: test
test :
	go test -v -race -timeout 30s ./...

.PHONY: clean
clean:
	rm -f apiserver

.DEFAULT_GOAL := build