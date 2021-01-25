PROJECT = go-click

OS ?= $(shell uname -s | tr A-Z a-z)
ARCH ?= amd64

GO ?= go
GOOS ?= $(OS)
GOARCH ?= $(ARCH)
CGO_ENABLED ?= 1
GOPROXY ?= direct
BASE_LDFLAGS = -w -s -X main.Version=$(VERSION)
ifneq ($(OS),darwin)
LDFLAGS = $(BASE_LDFLAGS) -linkmode external -extldflags -static
else
LDFLAGS = $(BASE_LDFLAGS)
endif
GO_ENV = env GOPROXY=$(GOPROXY) GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED)
GO_BUILD = $(GO_ENV) $(GO) build -ldflags "$(LDFLAGS)" -tags timetzdata
GO_TEST = $(GO_ENV) $(GO) test -count=1 -failfast

binary: $(PROJECT)

$(PROJECT):
	$(GO_BUILD) -o $(PROJECT) cmd/$(PROJECT)/main.go
