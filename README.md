# MyTodo - Production-Grade Task Management System

A production-ready, Jira-like task management system built with **Go** (backend) and **Next.js** (frontend), following **Clean Architecture** and **Domain-Driven Design** principles.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

---

## 🌟 Features

### Core Features
- ✅ **User Authentication** - JWT-based auth with OAuth support (Google, GitHub)
- ✅ **Organizations & Projects** - Multi-tenancy support
- ✅ **Issues/Tasks** - Full CRUD with advanced metadata
- ✅ **Kanban Boards** - Visual task management
- ✅ **Sprints** - Scrum/Agile workflow support
- ✅ **Comments & Attachments** - Collaboration features
- ✅ **Real-time Notifications** - WebSocket & email notifications
- ✅ **Custom Workflows** - Configurable status transitions
- ✅ **Automation Rules** - Event-driven task automation
- ✅ **Integrations** - GitHub, Slack, and webhook support
- ✅ **Advanced Search** - Full-text search with Elasticsearch

### Technical Highlights
- 🏗️ **Clean Architecture** - Separation of concerns across layers
- 🎯 **Domain-Driven Design** - Business logic in domain layer
- 🔄 **Event-Driven** - Async processing with message queues
- 📊 **Highly Scalable** - Supports 5M+ users, 1B+ issues
- 🚀 **Production-Ready** - Docker, Kubernetes, monitoring included
- 🧪 **Well-Tested** - Unit, integration, and E2E tests
- 📝 **Well-Documented** - Comprehensive documentation

---

## 📁 Project Structure

```
mytodo/
├── apps/
│   ├── api/              # Go backend (300+ folders)
│   │   ├── cmd/          # Application entrypoints
│   │   ├── internal/     # Domain modules (DDD)
│   │   │   ├── auth/
│   │   │   ├── issues/
│   │   │   ├── projects/
│   │   │   └── ...
│   │   ├── pkg/          # Shared libraries
│   │   └── tests/        # Integration & E2E tests
│   │
│   └── web/              # Next.js frontend
│       ├── app/          # App router pages
│       ├── components/   # React components
│       ├── hooks/        # Custom hooks
│       └── services/     # API clients
│
├── packages/             # Shared packages
│   ├── ui/              # UI component library
│   ├── types/           # TypeScript types
│   └── sdk/             # API SDK
│
├── infrastructure/       # DevOps & deployment
│   ├── docker/          # Docker configs
│   ├── kubernetes/      # K8s manifests
│   ├── terraform/       # Infrastructure as Code
│   └── scripts/         # Utility scripts
│
├── docs/                # Documentation
│   ├── architecture/    # Architecture docs
│   ├── api/            # API documentation
│   └── database/       # Database schema
│
├── plan/               # Development planning
│   ├── development-roadmap.md       # Layer-by-layer implementation plan
│   ├── progress-tracker.md          # Task checklist
│   ├── implementation-guide.md      # Step-by-step guide
│   └── schema-planning.md           # Database design
│
└── tools/              # CLI tools
    ├── migrations/     # Database migrations
    ├── seed/          # Data seeding
    └── codegen/       # Code generators
```

---

## 🚀 Quick Start

### Prerequisites

- **Go 1.21+**
- **Node.js 18+**
- **PostgreSQL 15+**
- **Redis 7+**
- **Docker & Docker Compose** (optional)

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/mytodo.git
cd mytodo
```

### 2. Set Up Environment Variables

```bash
cp .env.example .env
# Edit .env with your configuration
```

### 3. Start Development Environment

#### Option A: Using Docker (Recommended)
```bash
make docker-up        # Start PostgreSQL, Redis, etc.
make migrate-up       # Run database migrations
make seed             # Seed test data
make dev              # Start API server
```

#### Option B: Manual Setup
```bash
# Start PostgreSQL and Redis manually
# Then:
make deps             # Install dependencies
make migrate-up       # Run migrations
make seed             # Seed data
make dev              # Start server
```

### 4. Verify Installation

```bash
# Check API health
curl http://localhost:8080/health

# Or use make command
make health
```

---

## 🛠️ Development

### Available Commands

```bash
make help              # Show all available commands
make dev               # Start development server with hot reload
make test              # Run all tests
make test-unit         # Run unit tests only
make test-integration  # Run integration tests
make build             # Build production binary
make lint              # Run linters
make format            # Format code
```

### Running Tests

```bash
# All tests
make test

# Unit tests only
make test-unit

# Integration tests
make test-integration

# E2E tests
make test-e2e

# With coverage
make coverage
```

### Database Migrations

```bash
# Create new migration
make migrate-create name=add_users_table

# Run migrations
make migrate-up

# Rollback last migration
make migrate-down

# Check migration status
make migrate-status
```

---

## 📖 Documentation

### Planning & Architecture
- **[Development Roadmap](plan/development-roadmap.md)** - Layer-by-layer implementation plan
- **[Progress Tracker](plan/progress-tracker.md)** - Task checklist with status
- **[Implementation Guide](plan/implementation-guide.md)** - Step-by-step build guide
- **[Folder Structure](plan/folder%20strcuture%20and%20project%20architecture.md)** - Architecture overview
- **[Database Schema](plan/schema-planning.md)** - Complete database design

### API Documentation
- **[OpenAPI Spec](docs/api/openapi.yaml)** - REST API specification
- **[API Endpoints](docs/api/endpoints/)** - Endpoint documentation
- **[Postman Collection](docs/api/postman_collection.json)** - Import into Postman

### Architecture Documentation
- **[System Overview](docs/architecture/overview.md)** - High-level architecture
- **[Decision Records](docs/architecture/decision-records/)** - ADRs
- **[Diagrams](docs/architecture/diagrams/)** - System diagrams

---

## 🏗️ Architecture

### Clean Architecture Layers

```
┌─────────────────────────────────────────┐
│         Interface Layer (HTTP)          │
│  Controllers, DTOs, Routes, Validators  │
└─────────────┬───────────────────────────┘
              │
┌─────────────▼───────────────────────────┐
│        Application Layer (CQRS)         │
│  Commands, Queries, Handlers, UseCases  │
└─────────────┬───────────────────────────┘
              │
┌─────────────▼───────────────────────────┐
│         Domain Layer (Business)         │
│   Entities, Services, Repositories*     │
│         (*interfaces only)              │
└─────────────┬───────────────────────────┘
              │
┌─────────────▼───────────────────────────┐
│   Infrastructure Layer (Implementation) │
│  DB, Cache, Search, Queue, External API │
└─────────────────────────────────────────┘
```

### Technology Stack

**Backend:**
- **Language:** Go 1.21+
- **Framework:** Gin / Echo (HTTP)
- **Database:** PostgreSQL 15
- **Cache:** Redis 7
- **Search:** Elasticsearch 8
- **Queue:** NATS / Kafka
- **ORM:** SQLC / Ent

**Frontend:**
- **Framework:** Next.js 14 (App Router)
- **UI Library:** React 18
- **Styling:** Tailwind CSS
- **State:** Zustand / Jotai
- **API Client:** React Query

**DevOps:**
- **Containers:** Docker
- **Orchestration:** Kubernetes
- **IaC:** Terraform
- **CI/CD:** GitHub Actions
- **Monitoring:** Prometheus + Grafana
- **Logging:** ELK Stack

---

## 🧪 Testing Strategy

### Test Pyramid

```
     /\
    /E2E\        E2E Tests (10%)
   /──────\
  /  Intg  \     Integration Tests (30%)
 /──────────\
/    Unit    \   Unit Tests (60%)
──────────────
```

### Running Tests

```bash
# Unit tests (fast, no external dependencies)
make test-unit

# Integration tests (with real DB)
make test-integration

# E2E tests (full system)
make test-e2e

# All tests with coverage
make coverage
```

---

## 📊 Database Schema

The system uses **60+ tables** organized into logical modules:

- **Users & Auth** - users, sessions, oauth_accounts
- **Organizations** - organizations, members, roles
- **Projects** - projects, project_members
- **Issues** - issues, issue_types, statuses, priorities
- **Boards** - boards, columns, cards
- **Sprints** - sprints, sprint_issues
- **Comments** - comments, mentions
- **Attachments** - attachments, file_metadata
- **Notifications** - notifications, preferences
- **Workflows** - workflows, transitions
- **Automation** - rules, actions, triggers
- **Integrations** - integrations, webhooks

See [schema-planning.md](plan/schema-planning.md) for complete schema.

---

## 🚢 Deployment

### Docker Deployment

```bash
# Build images
make build-docker

# Run with Docker Compose
docker-compose up -d
```

### Kubernetes Deployment

```bash
# Apply manifests
kubectl apply -f infrastructure/kubernetes/

# Check status
kubectl get pods
```

### Production Checklist

- [ ] Environment variables configured
- [ ] SSL certificates installed
- [ ] Database backups scheduled
- [ ] Monitoring alerts configured
- [ ] Log aggregation set up
- [ ] Rate limiting configured
- [ ] Security headers enabled
- [ ] CORS whitelist configured

---

## 🔒 Security

- **Authentication:** JWT with refresh tokens
- **Authorization:** Role-based access control (RBAC)
- **Password Hashing:** bcrypt with salt
- **SQL Injection:** Parameterized queries
- **XSS Protection:** Input sanitization
- **CSRF Protection:** Token-based
- **Rate Limiting:** Per IP and per user
- **Security Headers:** HSTS, CSP, X-Frame-Options

---

## 📈 Performance

### Scale Targets

| Metric | Target |
|--------|--------|
| Users | 5M+ |
| Issues | 1B+ |
| Projects | 500K |
| Requests/sec | 50K+ |
| Response Time | < 100ms (p95) |
| Uptime | 99.9% |

### Optimization Strategies

- **Database:** Indexes, query optimization, connection pooling
- **Cache:** Redis for hot data, cache-aside pattern
- **Search:** Elasticsearch for full-text search
- **CDN:** Static asset caching
- **Load Balancing:** Horizontal scaling

---

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Clean Architecture principles
- Write tests for new features
- Update documentation
- Follow Go and TypeScript style guides
- Keep PRs focused and small

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- Inspired by Jira, Linear, Notion, and other great task management tools
- Built with best practices from the Go and React communities
- Architecture patterns from Uncle Bob's Clean Architecture

---

## 📞 Support

- **Documentation:** [docs/](docs/)
- **Issues:** [GitHub Issues](https://github.com/yourusername/mytodo/issues)
- **Discussions:** [GitHub Discussions](https://github.com/yourusername/mytodo/discussions)

---

## 🗺️ Roadmap

See [development-roadmap.md](plan/development-roadmap.md) for the complete implementation plan.

### Current Status: Phase 0 - Foundation Setup

**Next Milestones:**
1. ✅ Phase 0: Foundation Setup
2. ⚪ Phase 1: Authentication & Users
3. ⚪ Phase 2: Organizations & Projects
4. ⚪ Phase 3: Issues Core
5. ⚪ Phase 4: Boards & Sprints

---

**Built with ❤️ using Go and Next.js**
