GO := $(shell which go)

.PHONY: gen generate
gen: generate
generate:
	$(GO) generate ./...