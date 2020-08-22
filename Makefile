all: build

default: build

build:
	go build -o bin/go-brocade go-brocade.go
	go build -o bin/svg-to-svgo svg-to-svgo.go
