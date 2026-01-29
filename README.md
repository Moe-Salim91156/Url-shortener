# URL Shortener & Pastebin (Go)

A simple but **properly-architected URL Shortener and Pastebin** written in Go, built as a learning project to understand **backend system design, clean layering, authentication, and authorization** — not just "making it work".

This project includes:

* User authentication (register / login / logout)
* Session-based authorization
* URL ownership enforcement
* **Pastebin service for sharing text/code snippets**
* Web dashboard for managing URLs and pastes
* Clean service → store → handler architecture

---

## ✨ Features

### 🔐 Authentication & Authorization

* Register new users
* Login using username & password
* Session-based auth using cookies
* Protected routes using middleware
* Ownership checks (users can only delete their own URLs)

### 🔗 URL Shortening

* Create short URLs
* Resolve short URLs publicly
* Each URL is owned by a user

### 📋 Pastebin

* Create text/code pastes with optional titles
* View pastes in a clean, formatted display
* Copy paste content to clipboard
* Each paste is owned by a user
* Public paste viewing via short codes

### 📊 Dashboard

* List all URLs created by the logged-in user
* List all pastes created by the logged-in user
* Delete URLs and pastes you own
* Clean, minimal HTML pages (no JS frameworks)

### 🧱 Architecture

* Clear separation of concerns
* In-memory stores (easy to replace with DB)
* Dependency injection via `main.go`
* No global state hacks

---

## 📂 Project Structure

```
Url-shortener/
├── cmd/
│   └── main.go            # Application entry point
├── internal/
│   ├── handlers/          # HTTP handlers (no business logic)
│   ├── middleware/        # Auth middleware
│   ├── models/            # Domain models
│   ├── services/          # Business logic layer
│   └── store/             # Data layer (in-memory stores)
├── go.mod
├── go.sum
└── README.md
```

### Layer Responsibilities

| Layer      | Responsibility                         |
| ---------- | -------------------------------------- |
| Store      | Data persistence (currently in-memory) |
| Service    | Business rules & validation            |
| Handler    | HTTP glue (request/response)           |
| Middleware | Cross-cutting concerns (auth)          |

---

## 🚀 Running the Project

### Requirements

* Go 1.21+ (or compatible)

### Run

```bash
go run cmd/main.go
```

Server will start at:

```
http://localhost:8000
```

---

## 🌍 Application Routes

### Public Routes

| Route         | Description              |
| ------------- | ------------------------ |
| `/register`   | Register a new user      |
| `/login`      | Login                    |
| `/{code}`     | Redirect to original URL |
| `/paste/{code}` | View a paste           |

### Protected Routes (Login Required)

| Route           | Description                |
| --------------- | -------------------------- |
| `/dashboard`    | List your URLs and pastes  |
| `/shorten`      | Create a new short URL     |
| `/delete/{code}` | Delete a URL you own      |
| `/create-paste` | Create a new paste         |
| `/delete-paste/{code}` | Delete a paste you own |
| `/logout`       | Logout                     |

---

## 🧪 Authorization Rules

* A user **can only see their own URLs and pastes**
* A user **cannot delete URLs or pastes owned by others**
* Pastes can be viewed publicly by anyone with the link
* Unauthorized requests are redirected to `/login`

---

## 🛠 Design Decisions

### Why Services?

Services encapsulate business logic so that:

* Handlers stay thin
* Logic is testable
* Storage can be swapped without rewriting logic

### Why Middleware for Auth?

Authentication is a **cross-cutting concern**.
Middleware ensures:

* No duplication
* Clear protected boundaries
* Handlers assume a valid user

### Why In-Memory Stores?

* Focus on architecture first
* Easy to replace with SQLite / Postgres later

---

## 🔮 Possible Extensions

* Replace in-memory stores with SQLite/Postgres
* Add CSRF protection
* Add expiration for short URLs and pastes
* Syntax highlighting for code pastes
* Private pastes (password-protected)
* REST API version
* Unit tests for services

---

## 🎯 Learning Goals of This Project

This project was built to:

* Develop a **mental model for backend systems**
* Practice clean Go project structure
* Understand auth flows end-to-end
* Build something that can be rewritten confidently

> "I want to own this code — not just make it pass."

---

## 🧑‍💻 Author

**Mohammad Salim**
Software Engineer | Cloud & Backend Engineering

---

## 📜 License

This project is for learning and experimentation purposes.
Feel free to fork, modify, and extend it.

