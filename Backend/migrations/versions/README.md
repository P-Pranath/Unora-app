# Migrations

This directory will contain SQL migration files managed by Goose.

## Usage

### Create a new migration
```bash
make migration-create
```

### Run all pending migrations
```bash
make migration-up
```

### Rollback the last migration
```bash
make migration-down
```

### Check migration status
```bash
make migration-status
```
