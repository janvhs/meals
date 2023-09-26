.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	gofumpt -w -l .

.PHONY: run
run:
	. ./.env && go run .