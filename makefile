VERSION=$(shell git describe --tags 2>/dev/null || echo "snapshot")

.PHONY: build
build:
	@CGO_ENABLED=0 go build -o bin/setup cmd/setup/main.go

.PHONY: test
test:
	@go test -v ./...

.PHONY: release
release:
	@cd bin && tar -czf queue-setup.$(VERSION).x64.tar.gz setup
	@rm bin/setup