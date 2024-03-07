default: docs

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

.PHONY: test
test:
	go test ./...

GOPATH := $(shell go env GOPATH)

$(GOPATH)/bin/swagger: 
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5

swagger/public.swagger.json: $(GOPATH)/bin/swagger

internal/client internal/models: swagger/public.swagger.json
	@swagger generate client -f swagger/public.swagger.json -t internal

docs: $(wildcard examples/*)
	@go generate ./...

clean:
	@rm -rf docs internal/compute
