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
	@go test -race -cover -json -coverprofile=.coverage/cover.out.tmp ./... | grep -Ev "diodepb" | tparse -format=markdown > .coverage/test-report.md
	@cat .coverage/cover.out.tmp | grep -Ev "diodepb" > .coverage/cover.out
	@go tool cover -func=.coverage/cover.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' > .coverage/coverage.txt