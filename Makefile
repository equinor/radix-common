.PHONY: test
test:
	go test -cover `go list ./...`

.PHONY: lint
lint: bootstrap
	golangci-lint run --max-same-issues 0

HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)

bootstrap:
ifndef HAS_GOLANGCI_LINT
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@1.55.2
endif
