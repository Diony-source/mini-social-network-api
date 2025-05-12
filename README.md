# Mini Social Network API 🧑‍🤝‍🧑

A clean, modular backend API for a mini social network.  
Built with scalability, security, and observability in mind using Go, PostgreSQL, Logrus logging, and Docker.

---

## 🚀 Features

- User registration and login (JWT-based)
- Authenticated post creation
- Follow/unfollow user functionality
- PostgreSQL with raw SQL migrations
- Input sanitization against XSS (custom `sanitize` package)
- Centralized structured logging with Logrus
- Versioned API (`/v1/`)
- Fully Dockerized development environment
- Postman collection included

---

## 🧱 Tech Stack

- **Go** 1.21+
- **PostgreSQL** 15
- **Chi Router** (for routing)
- **JWT** for authentication
- **Logrus** for logging
- **Docker + Docker Compose**
- **godotenv** for environment management

---

## 📁 Project Structure

```
.
├── cmd/                    # Entry point
├── config/                 # Loads .env into Config struct
├── db/migrations/          # SQL migration files
├── internal/
│   ├── http/               # Router & middleware
│   └── v1/                 # Versioned business logic
│       ├── user/           # Register/Login
│       ├── post/           # Create Post
│       └── follow/         # Follow user
├── pkg/
│   ├── auth/               # JWT & password helpers
│   ├── db/                 # PostgreSQL connector
│   ├── logger/             # Logrus setup
│   └── sanitize/           # Input sanitizer
├── .env.example
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## ⚙️ Setup

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/mini-social-network-api.git
cd mini-social-network-api
```

### 2. Configure environment

Copy the sample and fill in your values:

```bash
cp .env.example .env
```

### 3. Build & Run (Docker)

```bash
docker-compose up --build
```

App: [http://localhost:8080](http://localhost:8080)

---

## 🧪 API Endpoints

| Method | Path        | Description            | Auth Required |
|--------|-------------|------------------------|---------------|
| POST   | /v1/register | Register user          | ❌
| POST   | /v1/login    | Login + get token      | ❌
| GET    | /v1/profile  | Auth user profile      | ✅
| POST   | /v1/posts    | Create post            | ✅
| POST   | /v1/follow   | Follow another user    | ✅

---

## 🔍 API Testing (Postman)

1. Import `mini-social-network-api.postman_collection.json`
2. Set environment:
   - `base_url` = http://localhost:8080
   - `jwt_token` = *(auto-filled on login)*

---

## 🧼 Input Sanitization

All user input (username, email, content) is passed through the custom `sanitize` package to prevent script injections.

---

## 📜 Logging

All handlers and services use [Logrus](https://github.com/sirupsen/logrus) for structured, context-aware logs.

Logs include:
- Request inputs
- Errors from decoding, validation, DB
- Successful actions

---

## 🐳 Dockerized

- Multistage Dockerfile (builder + minimal alpine)
- PostgreSQL & API services via Docker Compose
- Volumes for persistent db storage

---

## 📄 License

MIT