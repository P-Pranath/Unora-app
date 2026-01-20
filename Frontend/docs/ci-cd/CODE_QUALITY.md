# Git Hooks & Code Quality Setup

This document describes the Git hooks and code quality tools configured in this project.

## Overview

The project uses:

- **Husky** - Git hooks management
- **ESLint** - JavaScript/TypeScript linting
- **Prettier** - Code formatting
- **Commitlint** - Commit message validation

## Git Hooks

### Pre-push Hook

Runs automatically before `git push`. The push will be blocked if any check fails.

**Checks performed:**

1. **ESLint** - Validates code against linting rules
2. **Prettier** - Ensures code is properly formatted
3. **TypeScript** - Verifies there are no type errors

### Commit-msg Hook

Validates commit messages against conventional commit format.

## Commit Message Format

All commit messages must follow the [Conventional Commits](https://www.conventionalcommits.org/) specification.

### Format

```
type: subject
```

### Valid Types

| Type       | Description                                       |
| ---------- | ------------------------------------------------- |
| `feat`     | New feature                                       |
| `fix`      | Bug fix                                           |
| `docs`     | Documentation changes                             |
| `style`    | Code style changes (formatting, semicolons, etc.) |
| `refactor` | Code refactoring (no feature or bug fix)          |
| `perf`     | Performance improvements                          |
| `test`     | Adding or updating tests                          |
| `build`    | Build system or external dependencies             |
| `ci`       | CI/CD configuration changes                       |
| `chore`    | Other changes (maintenance, tooling)              |
| `revert`   | Revert a previous commit                          |

### Rules

- Subject must not be empty
- Type must not be empty
- Subject should be lowercase
- Header max length: 100 characters

### Examples

```bash
# ✅ Valid commits
feat: add user authentication
fix: resolve crash on app startup
docs: update readme with setup instructions
style: format code with prettier
refactor: simplify login logic
chore: update dependencies

# ❌ Invalid commits
Add new feature          # Missing type
feat:                    # Empty subject
FEAT: Add feature        # Type should be lowercase
feat: Add Feature        # Subject should be lowercase
```

## Available Scripts

```bash
# Linting
yarn lint           # Check for linting errors
yarn lint:fix       # Auto-fix linting errors

# Formatting
yarn format         # Format all files with Prettier
yarn check-format   # Check if files are formatted

# Type checking
yarn typecheck      # Run TypeScript type check
```

## Configuration Files

| File                   | Purpose                      |
| ---------------------- | ---------------------------- |
| `.eslintrc.js`         | ESLint configuration         |
| `.eslintignore`        | Files to ignore for ESLint   |
| `.prettierrc.js`       | Prettier configuration       |
| `.prettierignore`      | Files to ignore for Prettier |
| `commitlint.config.js` | Commit message rules         |
| `.husky/pre-push`      | Pre-push hook script         |
| `.husky/commit-msg`    | Commit message hook script   |

## Troubleshooting

### Bypassing Hooks (Not Recommended)

In rare cases where you need to bypass hooks:

```bash
# Bypass pre-push hook
git push --no-verify

# Bypass commit-msg hook
git commit --no-verify -m "message"
```

⚠️ **Warning:** Only use `--no-verify` in exceptional circumstances. The hooks exist to maintain code quality.

### Fixing Issues

```bash
# Fix linting errors
yarn lint:fix

# Fix formatting issues
yarn format

# Fix type errors
# Manually fix errors shown by:
yarn typecheck
```

## Setup for New Team Members

Git hooks are automatically installed when running `yarn install` due to the `prepare` script in `package.json`.

If hooks are not working, run:

```bash
npx husky install
```
