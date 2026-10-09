# DevSync

DevSync is a developer collaboration platform built with Go and a microservices-oriented architecture. It aims to bring developers together in organized workspaces for communication, knowledge sharing, and project collaboration.

## Features

- User registration, login, and email verification
- PASETO-based access and refresh tokens
- Workspace creation and membership management
- Chat, messaging, and threaded discussions
- PostgreSQL persistence
- Redis integration for supporting infrastructure

## Tech Stack

- **Backend:** Go
- **Database:** PostgreSQL
- **Authentication:** PASETO
- **Routing:** Gorilla Mux
- **Supporting infrastructure:** Redis
- **Architecture:** Microservices-oriented

## Project Structure

- `cmd/` — Service entry points
- `configs/` — Configuration loading
- `internal/` — Application logic and service components
- `migrations/` — Database migration files
- `docs/` — Project documentation
- `deploy/` — Deployment resources

## Getting Started

1. Install Go and PostgreSQL.
2. Configure environment variables using `.env.example`.
3. Set up the required database and supporting services.
4. Apply the database migrations.
5. Run the required services from their respective entry points under `cmd/`.

Refer to the project configuration and documentation for service-specific setup.

## Status

DevSync is under active development.
