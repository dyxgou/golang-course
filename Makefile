.DEFAULT_GOAL := run

run: build
	./bin/main

build:
	go build -o ./bin/main main.go
