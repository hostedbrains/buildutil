UTIL_BINARY=buildutil

## build: builds the buildutil binary for current OS
.PHONY: build
build:
	@echo "Building buildutil binary..."
	@go build -o bin/${UTIL_BINARY}
	@echo "Done!"

.PHONY: test
test:
	@echo "Testing buildutil binary..."
	@go test -coverprofile cover.out -v ./...
