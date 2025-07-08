VERSION = 0.3.3
GO ?= go
GOFMT ?= $(GO)fmt

DEFAULT_GOPATH=$(shell echo $$GOPATH|tr ':' '\n'|awk '!x[$$0]++'|sed '/^$$/d'|head -1)
ifeq ($(DEFAULT_GOPATH),)
DEFAULT_GOPATH := ~/go
endif
DEFAULT_GOBIN=$(DEFAULT_GOPATH)/bin
export PATH:=$(PATH):$(DEFAULT_GOBIN)

GOLANGCI_LINT=$(DEFAULT_GOBIN)/golangci-lint

GODOC=godoc -index -links=true -notes="BUG|TODO|XXX|ISSUE|FIXME"

#############################################################################
###   Building   ############################################################
#############################################################################

default: build

build:
	@echo ">> Checking build ..."
	@GO111MODULE=on $(GO) build -v ./...

#############################################################################
###   Custom Installs   #####################################################
#############################################################################

GOLANGCI_LINT = $(DEFAULT_GOBIN)/golangci-lint
TEST_RUNNER = $(DEFAULT_GOBIN)/gotestsum
GODA = $(DEFAULT_GOBIN)/goda

$(GOLANGCI_LINT):
	@echo "Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v2.2.1

$(TEST_RUNNER):
	@echo ">> Couldn't find $(TEST_RUNNER); installing ..."
	GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get gotest.tools/gotestsum && \
	$(GO) install gotest.tools/gotestsum

$(GODA):
	@echo ">> Couldn't find $(GODA); installing ..."
	@GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get -u github.com/loov/goda

#############################################################################
###   Linting & Testing   ###################################################
#############################################################################

show-linter:
	@echo $(GOLANGCI_LINT)

lint-silent: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) run \
	--enable=errcheck \
	--enable=dupl \
	--enable=unparam \
	--enable=wastedassign \
	--enable=ineffassign \
	--enable=revive \
	--enable=gocritic \
	--enable=misspell \
	--enable=unparam \
	--enable=lll \
	--enable=goconst \
	--enable=govet \
	--show-stats \
	./...

lint:
	@echo '>> Linting source code'
	@echo "($$($(GOLANGCI_LINT) --version))"
	@GO111MODULE=on $(MAKE) lint-silent

lint-help: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) --help
	@$(GOLANGCI_LINT) run --help

show-linters: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) linters

test: $(TEST_RUNNER)
	@echo '>> Running all tests'
	@GO111MODULE=on $(TEST_RUNNER) --format testname -- ./...

test-nocolor:
	@echo '>> Running all tests'
	@GO111MODULE=on $(GO) test ./... -v

upgrade-deps:
	@echo ">> Upgrading Go module dependencies ..."
	@go get -u ./...
	@go mod tidy

#############################################################################
###   Release Process   #####################################################
#############################################################################

tag:
	@echo "Tags:"
	@git tag|tail -5
	@git tag "v$(VERSION)"
	@echo "New tag list:"
	@git tag|tail -6

tag-and-push: tag
	@git push --tags

tag-delete: VERSION ?= 0.0.0
tag-delete:
	@git tag --delete v$(VERSION)
	@git push --delete origin v$(VERSION)
