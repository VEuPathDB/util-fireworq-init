VERSION  = $(shell git describe --tags 2>/dev/null || echo "snapshot")
BIN_NAME = setup-queues

.PHONY: build
build: bin/$(BIN_NAME) bin/entrypoint

.PHONY: test
test:
	go test -v ./...

.PHONY: release
release: bin/queue-setup.$(VERSION).x64.tar.gz
	rm bin/entrypoint bin/$(BIN_NAME)

bin/queue-setup.$(VERSION).x64.tar.gz: bin/$(BIN_NAME) bin/entrypoint
	cd bin && tar -czf queue-setup.$(VERSION).x64.tar.gz $(BIN_NAME) entrypoint

bin/$(BIN_NAME): $(shell find . -type f -name '*.go')
	CGO_ENABLED=0 go build -o bin/$(BIN_NAME) cmd/setup/main.go

bin/entrypoint: entrypoint
	cp entrypoint bin/entrypoint
