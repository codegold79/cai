#!/bin/bash
echo "lint with no options"
golangci-lint run

printf "\nlint with options\n"
golangci-lint run -E gosec,goconst,gofmt,goimports,unparam -c golangci.yaml

 printf "\ndone linting\n"