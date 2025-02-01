UTIL_BINARY=buildutil

## build: builds the buildutil binary for current OS
.PHONY: build
build:
	@echo "Building buildutil binary..."
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ./bin/buildutil --build --output bin/${UTIL_BINARY}-amd64-linux --withLDFlags
	@echo "Done!"

.PHONY: test
test:
	@echo "Testing buildutil binary..."
	@go test -coverprofile cover.out -v ./...
