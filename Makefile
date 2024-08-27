.PHONY: deps
deps:
	@go mod tidy

.PHONY: lint
lint:
	@golangci-lint run ./... --config .github/golangci.yaml

.PHONY: test
test:
	@go test -race ./...

.PHONY: test-coverage
test-coverage:
	@mkdir -p .coverage
	@go test `go list ./... | grep -Ev "diodepb|examples|internal"` -race -cover -json -coverprofile=.coverage/cover.out.tmp ./... | tparse -format=markdown > .coverage/test-report.md
	@cat .coverage/cover.out.tmp | grep -v "ingester" > .coverage/cover.out
	@go tool cover -func=.coverage/cover.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' > .coverage/coverage.txt

.PHONY: codegen
codegen:
	@go run internal/cmd/codegen/main.go | gofmt > ./diode/ingester.go
