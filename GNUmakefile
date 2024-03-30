default: docs

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: test
test:
	@go test ./...

docs: $(wildcard examples/*)
	@go generate ./...

clean:
	@rm -rf docs
