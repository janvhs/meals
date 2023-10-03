.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	gofumpt -w -l .

.PHONY: generate
generate:
	sqlc generate

.PHONY: run
run:
	. ./.env && go run .
