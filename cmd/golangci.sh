#!/bin/bash

GOLANGCI="${GOPATH}/bin/golangci-lint"
opt="run -E gofmt --timeout=100s -E goimports  -E asciicheck -E unparam  -E typecheck -E unconvert"

$GOLANGCI $opt ./...