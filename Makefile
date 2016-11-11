PKGS=$(shell glide nv)

all: fmt vet lint test

fmt:
	go fmt .
	go fmt ./cmd/gcode

vet:
	go vet .
	go vet ./cmd/gcode

lint:
	golint .
	golint ./cmd/gcode

test:
	go test -cover .
	go test -cover ./cmd/gcode
