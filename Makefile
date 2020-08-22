all: build

default: build

build:
	go build -o bin/go-brocade go-brocade.go
