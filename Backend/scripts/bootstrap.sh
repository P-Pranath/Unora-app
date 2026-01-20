#!/bin/bash

echo -e "\033[1;34m=== Installing development dependencies ===\033[0m"

echo "  Fetching missing Go dependencies for Ent..."
go get github.com/spf13/cobra@v1.6.1
go get github.com/olekukonko/tablewriter@v0.0.5
go get entgo.io/ent/cmd/ent@v0.14.4

echo "  Installing air for hot reload..."
go install github.com/air-verse/air@latest

echo "  Installing goose for database migrations..."
go install github.com/pressly/goose/v3/cmd/goose@latest

echo "  Installing Wire for dependency injection..."
go install github.com/google/wire/cmd/wire@latest

echo "  Installing Swag for Swagger documentation..."
go install github.com/swaggo/swag/cmd/swag@latest

echo "  Installing goimports for code formatting..."
go install golang.org/x/tools/cmd/goimports@latest

echo "  Installing gci for better import grouping..."
go install github.com/daixiang0/gci@latest

echo "  Installing golangci-lint for linting..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

echo "  Installing pre-commit (Python)..."
if command -v pip >/dev/null 2>&1; then
    pip install pre-commit
else
    echo "  pip not found, skipping pre-commit install"
fi

echo "  Installing commitlint (requires Node.js)..."
npm install --save-dev @commitlint/cli @commitlint/config-conventional

echo -e "\033[1;34m=== Installing pre-commit ===\033[0m"
if command -v pre-commit >/dev/null 2>&1; then
    pre-commit install
else
    echo "  pre-commit not found, skipping hook installation"
fi

echo -e "\033[1;34m=== Installing security tools ===\033[0m"
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install golang.org/x/vuln/cmd/govulncheck@latest

echo -e "\033[1;32m=== Dependencies installed successfully ===\033[0m"
echo -e "  \033[1;33mNote: You may need to add \$GOPATH/bin to your PATH\033[0m"
echo -e "  Current GOPATH: \033[90m${GOPATH:-not set}\033[0m"
