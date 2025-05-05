
# Mini Social Network API 🧑‍🤝‍🧑

A clean, modular backend API for a mini social network.  
Designed with extensibility and production practices in mind.  
This project was built during my Go backend learning journey and reflects both beginner-level practice and senior-level architecture.

---

## 🚀 Features

- User registration and login (JWT authentication)
- Post creation by authenticated users
- Follow system between users
- Versioned API (`/v1/`)
- PostgreSQL with raw SQL migrations
- Clean architecture with modular design
- Postman collection for API testing

---

## 🛠️ Tech Stack

- Go 1.21+
- PostgreSQL
- Chi Router
- JWT + Middleware
- Godotenv (.env support)

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
├── pkg/                    # Shared utils (auth, db)
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
DB_SOURCE=postgres://postgres:yourpassword@localhost:5432/mini_social_network_api?sslmode=disable
PORT=8080
JWT_SECRET=your_secret_key
```

### 3. Run database migrations

```bash
psql -U postgres -d mini_social_network_api -f db/migrations/000001_create_users.sql
psql -U postgres -d mini_social_network_api -f db/migrations/000002_create_posts.sql
psql -U postgres -d mini_social_network_api -f db/migrations/000003_create_follows.sql
```

### 4. Start the server

```bash
go run cmd/main.go
```

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
   - `jwt_token` = *(leave blank — auto set after login)*

---

## 🤝 Contribution

This project is a part of my backend learning roadmap.  
Future plans:
- Swagger documentation
- Docker support
- Redis caching
- Full test coverage

---

## 📄 License

MIT
