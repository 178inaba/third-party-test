LOCAL_PACKAGE_PREFIX := github.com/178inaba/third-party-test

.PHONY: all fmt fmt-diff ci-lint ci-lint-fix lint vet test run install-tools go-get-tools

all: ci-lint-fix fmt ci-lint lint vet test

fmt:
	goimports -local '$(LOCAL_PACKAGE_PREFIX)' -w .

fmt-diff:
	test -z $$(goimports -local '$(LOCAL_PACKAGE_PREFIX)' -l .) || (goimports -local '$(LOCAL_PACKAGE_PREFIX)' -d . && exit 1)

ci-lint:
	golangci-lint run

ci-lint-fix:
	golangci-lint run --fix

lint:
	golint -set_exit_status ./...

vet:
	go vet ./...

test:
	go test -race -count 1 -cover ./...

run:
	go run main.go

install-tools: go-get-tools
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.31.0

go-get-tools:
	GO111MODULE=off go get -u golang.org/x/lint/golint golang.org/x/tools/cmd/goimports
