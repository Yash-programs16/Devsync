# DevSync — Developer Collaboration Platform

> **Status: Under Active Development 🚧**
>
> DevSync is an evolving, microservices-based developer collaboration platform designed to unify communication, engineering knowledge, and development workflows in one workspace.

## Overview

Development teams often rely on disconnected tools for messaging, sharing code, documenting decisions, managing workspaces, and tracking Git activity. This fragmentation leads to lost technical context, repeated discussions, and inefficient collaboration.

DevSync aims to solve this problem by bringing developer communication, knowledge management, and engineering integrations into a single platform built specifically for programmers.

The long-term vision is to create a developer workspace where discussions become searchable knowledge, technical decisions remain accessible, and repository activity connects directly with team collaboration.

## Current Development Status

DevSync is currently under active development. The architecture and feature set described below represent the current design and planned direction. Not every service or feature has been implemented yet.

### Implemented Foundations

- Go backend project structure and configuration.
- PostgreSQL database integration and migrations.
- Authentication service with registration and login.
- Password hashing and email OTP verification.
- PASETO access and refresh token flows.
- Logout and token refresh functionality.
- Workspace creation with transactional database operations.
- Workspace ownership and membership foundations.

### In Progress / Planned

- API Gateway and inter-service gRPC communication.
- Chat and real-time messaging.
- Knowledge management and decision tracking.
- GitHub and GitLab integrations.
- Redis-powered presence, caching, and event distribution.
- Notifications and unified search.
- Frontend integration using Flutter or React.
- Containerization, observability, and deployment.

## Architecture — DevSync v1

The planned architecture follows a service-oriented design using Go, gRPC, PostgreSQL, and Redis.

```text
                 Flutter / React
                        |
                   HTTPS / REST
                        |
                 +--------------+
                 | API Gateway  |
                 | Auth         |
                 | Rate Limiting|
                 | Routing      |
                 +------+-------+
                        |
                  Internal gRPC
                        |
       +----------------+----------------+
       |        |       |       |        |
      Auth  Workspace  Chat  Knowledge Integration
       |        |       |       |        |
       +--------+-------+-------+--------+
                        |
             PostgreSQL + Redis
                        |
          Notifications / Presence / Search
```

The API Gateway will be the public entry point. Frontend clients will communicate through the gateway rather than accessing internal services directly.

Services will communicate through gRPC for structured internal requests, while Redis will support event distribution and fast-changing state where appropriate.

## Planned Services

### 1. Auth Service

**Responsibility:** Identity, authentication, and session management.

Planned capabilities:
- User registration and login.
- Password hashing and email verification.
- PASETO access and refresh tokens.
- Token verification, refresh, and logout.
- OAuth authentication.
- Friend requests and direct-message metadata.

Planned tables:
`users`, `refresh_tokens`, `friendships`, `friend_requests`

The Auth Service will own identity-related data and will not manage chat messages.

### 2. Workspace Service

**Responsibility:** Team and workspace organization.

Planned capabilities:
- Create and delete workspaces.
- Invite and remove members.
- Roles and permissions.
- Channel creation and management.
- Workspace membership and access control.

Planned tables:
`workspaces`, `workspace_members`, `roles`, `permissions`, `channels`, `invites`

### 3. Chat Service

**Responsibility:** Messaging and real-time collaboration.

Planned capabilities:
- Real-time messages.
- Threads and replies.
- Message editing and deletion.
- Emoji reactions.
- Typing indicators.
- Delivery and read receipts.
- Message history and streaming.

Planned communication:
- gRPC unary methods for history and message operations.
- Bidirectional streaming for live chat where appropriate.
- Redis Pub/Sub for message-event distribution.

### 4. Knowledge Service

**Responsibility:** Preserve and organize engineering knowledge.

Planned capabilities:
- Code snippets with syntax highlighting.
- Technical decisions and decision records.
- Polls and voting outcomes.
- Documentation and useful resources.
- Pinned discussions.
- Links between discussions, snippets, and decisions.
- Searchable engineering knowledge.

A key workflow will transform