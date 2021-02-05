GO ?= go
GOOS ?= $(shell uname -s | tr A-Z a-z)
GOARCH ?= amd64
CGO_ENABLED ?= 0
GOPROXY ?= direct
LDFLAGS = -w -s

GO_ENV = env GOPROXY=$(GOPROXY) GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED)
GO_BUILD ?= $(GO_ENV) $(GO) build -ldflags "$(LDFLAGS)"
GO_RUN ?= $(GO_ENV) $(GO) run
GO_TEST ?= $(GO_ENV) $(GO) test -count=1 -failfast
GOLANGCI_LINT ?= golangci-lint

PROJECT = go-click

binary:
	$(GO_BUILD) -o $(PROJECT) cmd/$(PROJECT)/*.go

run:
	$(GO_RUN) cmd/$(PROJECT)/*.go

lint:
	$(GOLANGCI_LINT) run

clean:
	@rm $(PROJECT) || true
