# Ent Generated Code

This directory contains generated Ent code. Do not edit files in this directory manually.

To regenerate code after modifying schemas:

```bash
make ent-generate
```

Or directly:

```bash
go run entgo.io/ent/cmd/ent generate ./ent/schema --target ./ent/generated --feature sql/upsert
```
