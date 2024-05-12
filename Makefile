GOMODULE  = GO111MODULE=on
GO        = $(GOMODULE) go

.PHONY: deps
deps:
	$(GO) mod tidy -compat=1.22.2 && $(GO) mod vendor

.PHONY: test
test:
	$(GO) test ./... -p 4 -v