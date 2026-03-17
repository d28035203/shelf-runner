# Shelf Runner

Bookstore REST API written in Go. Clean package layout, MySQL via GORM, Gorilla Mux routing, and environment-based configuration.

## Features

- Full CRUD for books
- Layered layout under `pkg/` (config, models, controllers, routes, utils)
- MySQL + GORM with auto-migrate
- JSON error responses
- `.env` configuration via godotenv

## Tech Stack

| Layer | Choice |
|-------|--------|
| Language | Go |
| Router | Gorilla Mux |
| ORM | GORM |
| Database | MySQL |
| Config | godotenv |

## Project Structure

```
shelf-runner/
├── cmd/main/main.go
├── pkg/
│   ├── config/app.go
│   ├── controllers/book-controller.go
│   ├── models/book.go
│   ├── routes/bookstore-routes.go
│   └── utils/utils.go
├── .env.example
├── go.mod
└── README.md
```

## API

Base URL: `http://localhost:9010`

| Method | Path | Description |
|--------|------|-------------|
| GET | `/book/` | List books |
| GET | `/book/{bookId}` | Get one book |
| POST | `/book/` | Create book |
| PUT | `/book/{bookId}` | Update book |
| DELETE | `/book/{bookId}` | Delete book |

### Create example

```bash
curl -X POST http://localhost:9010/book/ \
  -H "Content-Type: application/json" \
  -d '{"name":"The Go Programming Language","author":"Donovan & Kernighan","publication":"Addison-Wesley"}'
```

## Setup

```bash
git clone https://github.com/d28035203/shelf-runner.git
cd shelf-runner
cp .env.example .env
# configure MySQL DSN
go mod tidy
go run ./cmd/main
```

## Environment

```env
SERVER_HOST=0.0.0.0
SERVER_PORT=9010
DB_USER=root
DB_PASS=password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=bookstore
DB_DSN=root:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local
```

## License

MIT
