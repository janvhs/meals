.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	gofumpt -w -l .
