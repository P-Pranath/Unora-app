# Backend API

A Go backend project following clean architecture patterns.

## Tech Stack

- **Go 1.25** - Programming language
- **Gin** - Web framework
- **Ent** - ORM
- **Goose** - Database migrations
- **MySQL 8.0** - Database
- **Redis 7** - Cache & Queue
- **Asynq** - Background job processing
- **MinIO** - Object storage with presigned URLs
- **Docker** - Containerization

## Quick Start

```bash
# Copy environment file
cp .env.example .env

# Install development dependencies
make bootstrap

# Start infrastructure (MySQL, Redis, MinIO)
make docker-infra

# Run migrations
make migration-up

# Start the server
make run
```

## Docker Development

```bash
# Start all services
make docker-up

# View logs
make docker-logs

# Stop all services
make docker-down
```

## Available Commands

Run `make help` to see all available commands.

### Development
- `make run` - Run server locally
- `make run-worker` - Run worker locally
- `make run-cron` - Run cron locally
- `make build` - Build all binaries
- `make test` - Run tests

### Docker
- `make docker` - Interactive Docker menu
- `make docker-up` - Start containers
- `make docker-down` - Stop containers
- `make docker-logs` - View logs
- `make docker-infra` - Start only MySQL, Redis, MinIO

### Database
- `make migration-create` - Create new migration
- `make migration-up` - Run migrations
- `make migration-down` - Rollback migration
- `make migration-status` - View status
- `make backup-db` - Backup database
- `make seed` - Run seeder

### Ent ORM
- `make ent-init` - Create new schema
- `make ent-generate` - Generate Ent code

### Code Quality
- `make lint` - Run linter
- `make fmt` - Format code
- `make security-scan` - Run security tools

### Module Scaffolding
- `make scaffold-module` - Create new module

## Project Structure

```
.
├── cmd/                  # Entry points
│   ├── server/          # HTTP API server
│   ├── worker/          # Background worker
│   ├── cron/            # Cron scheduler
│   └── seeder/          # Database seeder
├── internal/            # Private application code
│   ├── config/          # Configuration
│   ├── router/          # Routes & middleware
│   ├── health/          # Health checks
│   └── shared/          # Shared utilities
├── pkg/                 # Public packages
│   ├── database/        # DB & Ent setup
│   ├── errors/          # Custom errors
│   ├── logger/          # Logging
│   ├── queue/           # Asynq client
│   ├── redis/           # Redis client
│   ├── response/        # API responses
│   ├── storage/         # MinIO client
│   └── validation/      # Request validation
├── ent/                 # Ent schemas
├── migrations/          # SQL migrations
├── docker/              # Docker configs
└── scripts/             # Utility scripts
```

## Storage (MinIO)

The project uses MinIO for file storage with presigned URL support.

### Frontend Integration

```javascript
// Get presigned upload URL
const response = await fetch('/api/v1/storage/presigned-upload', {
  method: 'POST',
  body: JSON.stringify({ fileName: 'document.pdf', contentType: 'application/pdf' })
});
const { presignedUrl } = await response.json();

// Upload file directly to MinIO
await fetch(presignedUrl, {
  method: 'PUT',
  body: file,
  headers: { 'Content-Type': 'application/pdf' }
});

// Get presigned download URL
const downloadResponse = await fetch('/api/v1/storage/presigned-download/document.pdf');
const { presignedUrl: downloadUrl } = await downloadResponse.json();
```

## Environment Variables

See `.env.example` for all configuration options.

## License

Private
