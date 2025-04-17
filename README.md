# 🔐 go-vault

A secure file management system built with Go. Features user authentication with JWT, structured file uploads, folder organization, and per-user access control. Designed with clean architecture and best backend practices in mind.

---

## 🚀 Features

- ✅ User signup and login with JWT
- 📂 File upload with optional folder structure
- 🔐 Authentication middleware for route protection
- 📥 Download files securely
- 🗑️ Delete files owned by the user
- 📁 Create, list, and manage folders
- 👤 User isolation (each user accesses only their files)
- 💾 Local file storage (optional S3 integration in future)
- 🧱 Clean architecture with handler/service/repository separation

---

## 🧰 Tech Stack

- [Go](https://golang.org)
- [Gin](https://github.com/gin-gonic/gin) – Web framework
- [GORM](https://gorm.io) – ORM for database access
- JWT for authentication
- Bcrypt for password hashing
- SQLite (for development) or PostgreSQL (for production)
- Docker + Makefile for environment setup and build

---

## 📁 Project Structure

```
go-vault/
├── cmd/                # Application entry point
├── internal/
│   ├── auth/           # Auth logic (JWT, login, signup)
│   ├── user/           # User models, services, repositories
│   ├── file/           # File upload, download, delete
│   ├── folder/         # Folder management
│   ├── middleware/     # Auth middleware
│   ├── config/         # Environment config
│   ├── db/             # Database connection and migrations
│   └── utils/          # Utility functions
├── uploads/            # Local file storage directory
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
├── .env
└── README.md
```

---

## 🏁 Getting Started

### Prerequisites
- Go 1.21+
- Docker & Docker Compose (optional)

### Run locally
```bash
git clone https://github.com/your-username/go-vault.git
cd go-vault
make run
```

### Run with Docker
```bash
docker-compose up --build
```

---

## 📌 License

This project is licensed under the MIT License.