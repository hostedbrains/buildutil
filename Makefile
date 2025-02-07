UTIL_BINARY=buildutil

export GOBIN ?= $(shell pwd)/bin

GO_FILES := $(shell \
	find . '(' -path '*/.*' -o -path './vendor' ')' -prune \
	-o -name '*.go' -print | cut -b3-)

GOLINT = $(GOBIN)/golint
STATICCHECK = $(GOBIN)/staticcheck

## build: builds the buildutil binary for current OS
.PHONY: build
build: test lint scan build_linux build_mac
	@echo "Building buildutil binary..."
	@env CGO_ENABLED=0 buildutil --build --output bin/${UTIL_BINARY} --withLDFlags
	@echo "Done!"

## build: builds the buildutil binary for Linux OS
.PHONY: build_linux
build_linux:
	@echo "Building buildutil binary..."
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 buildutil --build --output bin/${UTIL_BINARY}-amd64-linux --withLDFlags
	@echo "Done!"

## build: builds the buildutil binary for Mac OS
.PHONY: build_mac
build_mac:
	@echo "Building buildutil binary..."
	@env GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 buildutil --build --output bin/${UTIL_BINARY} --withLDFlags
	@echo "Done!"

## test: executes tests
.PHONY: test
test: scan
	@echo "Testing buildutil application..."
	@go test -coverprofile cover.out -v ./...

## Create test coverage report
.PHONY: testreport
testreport:
	@echo "Generating test report..."
	@go tool cover -html=cover.out

## Vulnerability scanning
.PHONY: scan
scan:
	@echo "Doing vulnerability scanning"
	@govulncheck ./...

## Download required modules.
.PHONY: install
install:
	go mod download

$(GOLINT): tools/go.mod
	cd tools && go install golang.org/x/lint/golint

$(STATICCHECK): tools/go.mod
	cd tools && go install honnef.co/go/tools/cmd/staticcheck@2024.1.1

.PHONY: lint
lint: install $(GOLINT) $(STATICCHECK)
	@rm -rf lint.log
	@echo "Checking gofmt"
	@gofmt -d -s $(GO_FILES) 2>&1 | tee lint.log
	@echo "Checking go vet"
	@go vet ./... 2>&1 | tee -a lint.log
	@echo "Checking golint"
	@$(GOLINT) ./... | tee -a lint.log
	@echo "Checking staticcheck"
	@$(STATICCHECK) ./... 2>&1 |  tee -a lint.log
	@echo "Checking for license headers..."
	@./.build/check_license.sh | tee -a lint.log
	@[ ! -s lint.log ]