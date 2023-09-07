#!/bin/bash

GOLANGCI="golangci-lint"
opt="run -E gofmt --timeout=100s -E goimports -E golint  -E asciicheck -E unparam -E gosec -E typecheck -E unconvert"

$GOLANGCI $opt ./...