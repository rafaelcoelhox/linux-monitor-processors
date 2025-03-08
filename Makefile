.PHONY: build run test clean setup

build:
	go build -o bin/monitor cmd/monitor/main.go

run: build
	./bin/monitor

test:
	go test ./...

clean:
	rm -rf bin/

setup:
	./scripts/setup-ollama 