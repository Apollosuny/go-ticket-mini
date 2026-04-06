# go-ticket-mini

A minimal ticket support system built with Go using a service-oriented architecture.

## Overview

This project provides a simple backend for managing support tickets and comments. It is organized with a public HTTP gateway and an internal gRPC service.

## Architecture

```text
Client
  -> HTTP/JSON
API Gateway (Gin)
  -> gRPC
Ticket Service
  -> PostgreSQL
```

## Core Features

- create a ticket
- get a ticket by id
- list tickets
- update ticket status
- add a comment to a ticket

## Infrastructure

- **Go** for backend services
- **Gin** for the API Gateway
- **gRPC** with **Protocol Buffers** for internal service communication
- **PostgreSQL** for persistent storage
- **GORM** for database access
- **Docker Compose** for local infrastructure setup
- **Makefile** for common development commands

## Project Structure

```text
go-ticket-mini/
├── api/
│   └── proto/
├── cmd/
│   ├── api-gateway/
│   └── ticket-service/
├── internal/
│   └── ticket/
│       ├── domain/
│       ├── repository/
│       ├── service/
│       ├── endpoint/
│       └── transport/
├── pkg/
│   ├── config/
│   └── database/
├── migrations/
├── docker-compose.yml
├── Makefile
└── go.mod
```

## Requirements

- Go 1.23+
- Docker
- protoc
- protoc-gen-go
- protoc-gen-go-grpc

## Setup

### 1. Clone the repository

```bash
git clone git@github.com:Apollosuny/go-ticket-mini.git
cd go-ticket-mini
```

### 2. Create the environment file

```bash
cp .env.example .env
```

### 3. Start PostgreSQL

```bash
docker compose up -d
```

### 4. Run the migration

```bash
psql -h localhost -U postgres -d go_ticket_mini -f migrations/001_init.sql
```

### 5. Generate protobuf code

```bash
make proto
```

### 6. Run the ticket service

```bash
make run-ticket
```

## Common Commands

```bash
make proto
make run-ticket
make run-gateway
go mod tidy
docker compose up -d
docker compose down
```
