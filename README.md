# 📱 Social Media API (Go)

A simple **REST API** built with **Go** for managing social media posts, users, and interactions.

## 🚀 Features  
- User authentication (JWT)  
- Create, read, update, and delete posts  
- Like and comment on posts  
- Follow/unfollow users  

## 🔧 Installation  

```sh
git clone https://github.com/lahiruudayakumara/social-media-api.git
cd social-media-api
go mod tidy
go run main.go
```

## 🛠 Environment Variables  

Create a `.env` file and add the following:  

```env
JWT_SECRET=JWT-Secret-Key
DATABASE_URL=postgres://postgres:Your-Password@localhost:5432/Your-DB-Name?sslmode=disable
```

## ⚙️ API Endpoints  

| Method | Endpoint         | Description              |
|--------|-----------------|--------------------------|
| POST   | `/auth/signup`  | Register a new user      |
| POST   | `/auth/login`   | User login (JWT token)   |
| GET    | `/posts`        | Fetch all posts         |
| POST   | `/posts`        | Create a new post       |
| PUT    | `/posts/{id}`   | Update a post           |
| DELETE | `/posts/{id}`   | Delete a post           |
| POST   | `/posts/{id}/like` | Like a post         |
| POST   | `/posts/{id}/comment` | Comment on a post |

## 🛠 Tech Stack  
- **Go (Golang)**  
- **Gin** (for routing)  
- **GORM** (ORM for database)  
- **PostgreSQL** (Database)  
- **JWT** (Authentication)  
