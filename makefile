VERSION=$(shell git describe --tags 2>/dev/null || echo "snapshot")

.PHONY: build
build: bin/setup-queues bin/entrypoint

.PHONY: test
test:
	go test -v ./...

.PHONY: release
release: bin/queue-setup.$(VERSION).x64.tar.gz

bin/queue-setup.$(VERSION).x64.tar.gz: bin/setup-queues bin/entrypoint
	cd bin && tar -czf queue-setup.$(VERSION).x64.tar.gz setup-queues entrypoint

bin/setup-queues: $(shell find . -type f -name '*.go')
	CGO_ENABLED=0 go build -o bin/setup-queues cmd/setup/main.go

bin/entrypoint: entrypoint
	cp entrypoint bin/entrypoint
