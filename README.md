# Mini Social Network API 🧑‍🤝‍🧑

A clean, modular backend API for a mini social network.  
Designed with extensibility, security, and production practices in mind.  
This project was built during my Go backend learning journey and reflects both beginner-level practice and senior-level architecture.

---

## 🚀 Features

- User registration and login (JWT authentication)
- Post creation by authenticated users
- Follow system between users
- Request sanitization to prevent injection attacks
- Versioned API (`/v1/`)
- PostgreSQL with raw SQL migrations
- Dockerized setup with Compose support
- Clean architecture with modular design
- Postman collection for API testing

---

## 🛠️ Tech Stack

- Go 1.21+
- PostgreSQL
- Docker & Docker Compose
- Chi Router
- JWT + Middleware
- Godotenv (.env support)
- Custom sanitize input module

---

## 📂 Project Structure

```
.
├── cmd/                    # Entry point: main.go
├── config/                 # Loads .env
├── db/                     # SQL migrations
├── internal/
│   ├── http/               # Router and dependencies
│   ├── middleware/         # JWT Auth middleware
│   └── v1/                 # Versioned domains
│       ├── user/
│       ├── post/
│       └── follow/
├── pkg/
│   ├── auth/               # JWT & Password utils
│   ├── db/                 # Postgres connector
│   └── sanitize/           # Input sanitizer for safety
├── Dockerfile
├── docker-compose.yml
├── .env.example
└── README.md
```

---

## ⚙️ Setup

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/mini-social-network-api.git
cd mini-social-network-api
```

### 2. Create and configure `.env`

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=mini_social_network_api
PORT=8080
JWT_SECRET=your_super_secret_key
```

> **Note:** Use `.env.example` as a template.

### 3. Run with Docker 🐳

```bash
docker-compose up --build
```

> This will:
> - Spin up PostgreSQL
> - Run migrations
> - Launch the Go API

Server runs at `http://localhost:8080`

---

## 🔎 API Endpoints

All routes are prefixed with `/v1`

| Method | Path        | Description               | Auth Required |
|--------|-------------|---------------------------|----------------|
| POST   | /register   | Register new user         | ❌
| POST   | /login      | Login and get JWT         | ❌
| GET    | /profile    | View your profile         | ✅
| POST   | /posts      | Create a new post         | ✅
| POST   | /follow     | Follow another user       | ✅

---

## 📫 API Testing

📦 Postman collection: `mini-social-network-api.postman_collection.json`

1. Import the collection in Postman  
2. Create an environment with:
   - `base_url` = `http://localhost:8080`
   - `jwt_token` = *(auto set after login)*

---

## 🧼 Input Sanitization

To protect against XSS or injection:
- All user inputs are sanitized via `pkg/sanitize`
- Applied on handlers for registration, login, posts, and follows

---

## 🤝 Contribution

This project is a part of my backend development growth.  
Upcoming improvements:
- Swagger documentation
- Redis caching
- Full test coverage
- CI/CD pipeline

---

## 📄 License

MIT