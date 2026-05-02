# Movie Reseller (Practice Project)

A **practice project** to sharpen Go skills by building a micro‑service based movie reseller platform.  
The service exposes a **REST API** (Gin) and **gRPC** endpoints for communication with other micro‑services, uses **GORM** as the ORM, and stores data in **PostgreSQL**.

## Features

- CRUD operations for movies, users, and orders via Gin REST endpoints  
- gRPC services for inter‑service communication  
- PostgreSQL persistence with GORM  
- Dockerized for easy deployment  
- Live‑reload development using **air**

## Prerequisites

- Go 1.22+  
- Docker & Docker Compose  
- [air](https://github.com/cosmtrek/air) (optional, for hot‑reloading)  
- PostgreSQL (local or via Docker)

## Project Structure (high‑level)

```
.
├── cmd/
│   └── main.go                # Application entry point
├── config/
│   └── config.go              # Configuration loader
├── config.json                # Example configuration file
├── internal/
│   ├── handler/               # Gin HTTP handlers
│   │   └── test.go
│   ├── model/                 # GORM model definitions
│   ├── proto/                 # Protobuf definitions & generated code
│   │   ├── moviepb/
│   │   │   ├── movie.pb.go
│   │   │   └── movie_grpc.pb.go
│   │   ├── userpb/
│   │   │   ├── user.pb.go
│   │   │   └── user_grpc.pb.go
│   │   ├── movie.proto
│   │   ├── user.proto
│   │   └── server.go
│   ├── repository/            # Data‑access layer
│   ├── route/
│   │   └── routes.go          # Gin route definitions
│   └── service/
│       └── user_service.go    # Business logic
├── Dockerfile                 # Container image definition
├── docker-compose.yml         # Docker Compose configuration
├── go.mod
├── go.sum
└── README.md
```

## Development

### 1. Start PostgreSQL (Docker)

```bash
docker compose up -d postgres
```

### 2. Run with live reload (air)

```bash
# Install air if not already installed
go install github.com/cosmtrek/air@latest

# Start the application; air watches for file changes and restarts automatically
air
```

The API will be available at `http://localhost:8080` and gRPC on `localhost:50051`.

## Testing

Run unit tests across the project:

```bash
go test ./... -v
```

## Building & Running with Docker

### Build the image

```bash
docker build -t movie-reseller .
```

### Run the container

```bash
docker run -d \
  -p 8080:8080 -p 50051:50051 \
  --name movie-reseller \
  --env-file .env \
  movie-reseller
```

## Deploy with Docker Compose

```bash
docker compose up --build
```

This brings up both the **movie‑reseller** service and a PostgreSQL instance defined in `docker-compose.yml`. The APIs are exposed on the same ports as above.

## Usage

- **REST API**: `GET /movies`, `POST /orders`, etc. (see `internal/handler` for route implementations)  
- **gRPC**: Use the generated protobuf stubs in `internal/proto` to call services from other micro‑services.

## TODO

- [ ] Web socket for real‑time support  
- [ ] Comprehensive testing (unit, integration, e2e)