VERSION = 0.2.0
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
###   Custom Installs   #####################################################
#############################################################################

GOLANGCI_LINT = $(DEFAULT_GOBIN)/golangci-lint
TEST_RUNNER = $(DEFAULT_GOBIN)/gotestsum
GODA = $(DEFAULT_GOBIN)/goda

$(GOLANGCI_LINT):
	@echo ">> Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v1.50.1

$(TEST_RUNNER):
	@echo ">> Couldn't find $(TEST_RUNNER); installing ..."
	@GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get -u get gotest.tools/gotestsum && \
	$(GO) get -u install gotest.tools/gotestsum

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
	@$(GOLANGCI_LINT) \
	--enable=typecheck \
	--enable=revive \
	--enable=gocritic \
	--enable=misspell \
	--enable=nakedret \
	--enable=unparam \
	--enable=lll \
	--enable=goconst \
	--out-format=colored-line-number \
	run ./...

lint:
	@echo '>> Linting source code'
	@GO111MODULE=on $(MAKE) lint-silent

test: $(TEST_RUNNER)
	@echo '>> Running all tests'
	@GO111MODULE=on $(TEST_RUNNER) --format testname -- ./...

test-nocolor:
	@echo '>> Running all tests'
	@GO111MODULE=on $(GO) test ./... -v

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
