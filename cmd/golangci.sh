#!/bin/bash

GOLANGCI="${GOPATH}/bin/golangci-lint --version"
opt="run -E gofmt --timeout=100s -E goimports -E golint  -E asciicheck -E unparam -E gosec -E typecheck -E unconvert"

$GOLANGCI $opt ./...