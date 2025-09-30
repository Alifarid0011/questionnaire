
# Questionnaire Back-End

A **production-ready backend service** for managing questionnaires, equipped with robust **authentication, authorization, and semantic validation features**.  
This project is designed with clean architecture principles, dependency injection (Wire), and containerized deployment (Docker + Docker Compose).

---

## 🚀 Features

- **Authentication**:  
  - Secure JWT-based authentication system.
  - Access and refresh tokens for session management.

- **Authorization**:  
  - Role-Based Access Control (**RBAC**) powered by **Casbin**.
  - Flexible policies with `sub`, `obj`, and `act` definitions:
    - **sub** → Subject (role or user ID).  
    - **obj** → Object (typically API endpoints).  
    - **act** → Action (`GET`, `POST`, `PUT`, `DELETE`, or `*`).  
    - **eft** → Effect (`allow` or `deny`).  
  - Example matcher rule:  
    ```text
    m = r.sub == "super_admin" || (g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && (r.act == p.act || p.act == "*"))
    ```

- **Semantic Answer Validation**:  
  - Integrates with [Semantic Answer Validator](https://github.com/Alifarid0011/semantic-answer-validator).  
  - Used to validate **short answers** against expected responses with semantic similarity.  

  Example request:
```http
  POST http://127.0.0.1:8020/check_answer
  Content-Type: application/json

  {
    "student_answer": "Artificial intelligence improves efficiency in many industries.",
    "accepted_answers": [
      "AI enhances productivity across different sectors.",
      "Smart systems reduce human errors."
    ]
  }
```

Example response:

```json
{
  "similarity_score": 0.9,
  "accepted": true,
  "details": {
    "sbert_similarity": 0.7707,
    "nli_entail_count": 1,
    "nli_contra_count": 0,
    "nli_entail_avg_prob": 0.4999,
    "nli_contra_avg_prob": 0.0003,
    "keywords_present": true,
    "student_negation": false
  }
}
```

---

## 📂 Project Structure

```
questionnaire-back-end/
│── casbin/                # Casbin model configuration
│   └── model.conf
│
│── config/                # Configuration layer (YAML-based)
│   └── environment/       # Environment configs
│       ├── casbin.go
│       ├── config.go
│       └── mongo.go
│
│── internal/              # Core business logic
│   ├── controller/        # HTTP controllers
│   ├── dto/               # Data transfer objects
│   ├── middleware/        # Custom middlewares
│   ├── models/            # Database models
│   ├── repository/        # Database repositories
│   ├── service/           # Service layer
│   └── validation/        # Request validation
│
│── routers/               # API routes
│── utils/                 # Utilities and helpers
│── wire/                  # Dependency injection (Google Wire)
│   ├── provider/
│   ├── wire.go
│   └── wire_gen.go
│
│── .gitignore
│── docker-compose.yml
│── Dockerfile
│── go.mod
│── main.go
│── README.md
```

---

## 🛠️ Tech Stack

* **Language**: [Go (Golang)](https://golang.org/)
* **Dependency Injection**: [Google Wire](https://github.com/google/wire)
* **Authorization**: [Casbin](https://casbin.org/) (RBAC model)
* **Authentication**: JWT tokens
* **Database**: MongoDB (via configuration layer)
* **Validation**: Custom request validators + Semantic Answer Validator
* **Containerization**: Docker & Docker Compose

---

## ⚡ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/questionnaire-back-end.git
cd questionnaire-back-end
```

### 2. Environment Configuration

* Place your environment-specific YAML configs under `config/environment/`.

### 3. Run with Docker

```bash
docker-compose up --build
```

### 4. Run Locally (without Docker)

```bash
go mod tidy
go run main.go
```

---

## 📖 API Authorization (Casbin)

* **Objects (obj)**: API endpoints
* **Actions (act)**: HTTP methods (`GET`, `POST`, `PUT`, `DELETE`)
* **Subjects (sub)**: Users or roles (`admin`, `teacher`, `student`, etc.)
* **Effects (eft)**: Permission outcome (`allow` or `deny`)

Policies can be configured directly via Casbin adapters.

---

## 🧪 Example: Semantic Answer Validator Integration

```http
POST http://127.0.0.1:8020/check_answer
Content-Type: application/json

{
  "student_answer": "Sports are dangerous and harm your health.",
  "accepted_answers": [
    "Physical activity improves fitness and mental health.",
    "Exercise is important for maintaining overall wellbeing."
  ],
  "keywords": ["Exercise","activity"]
}
```

Response:

```json
{
  "similarity_score": 0.0,
  "accepted": false,
  "details": {
    "sbert_similarity": 0.5345,
    "nli_entail_count": 0,
    "nli_contra_count": 2,
    "nli_entail_avg_prob": 0.0028,
    "nli_contra_avg_prob": 0.8741,
    "keywords_present": false,
    "student_negation": false
  }
}
```

---

## 📦 Ready for Use

* ✅ Fully containerized (Docker & Docker Compose)
* ✅ Clean architecture (controllers, services, repositories, DTOs)
* ✅ Role-based authorization (Casbin RBAC)
* ✅ JWT authentication
* ✅ Semantic short-answer validation

---
