#!/bin/bash

GOLANGCI="${GOPATH}/bin/golangci-lint"
opt="run --verbose -E gofmt -E goimports  -E asciicheck -E unparam  -E typecheck -E unconvert"

$GOLANGCI $opt ./...