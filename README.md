# 🚀 Go MongoDB API with Frontend

A full-stack web application built with Go, MongoDB, and vanilla JavaScript. This project provides a complete user management system with a RESTful API backend and a beautiful responsive frontend.

## ✨ Features

- **RESTful API** - Complete CRUD operations for user management
- **MongoDB Integration** - Modern MongoDB driver with proper connection handling
- **Beautiful Frontend** - Responsive web interface with modern design
- **CORS Support** - Cross-origin resource sharing enabled
- **Real-time Updates** - Frontend updates immediately after API operations
- **Input Validation** - Form validation and error handling
- **Mobile Responsive** - Works perfectly on all devices

## 🛠️ Tech Stack

### Backend
- **Go 1.24+** - Main programming language
- **MongoDB** - NoSQL database
- **httprouter** - Fast HTTP router
- **CORS** - Cross-origin support

### Frontend
- **HTML5** - Structure
- **CSS3** - Styling with gradients and animations
- **Vanilla JavaScript** - Interactive functionality
- **Responsive Design** - Mobile-first approach

## 📋 Prerequisites

Before running this application, make sure you have:

- [Go 1.24+](https://golang.org/dl/) installed
- [MongoDB](https://www.mongodb.com/try/download/community) installed and running
- [Git](https://git-scm.com/) for version control

## 🚀 Quick Start

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/go-mongodb-api.git
cd go-mongodb-api
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Start MongoDB
Make sure MongoDB is running on your system:
```bash
# On macOS with Homebrew
brew services start mongodb/brew/mongodb-community

# On Linux
sudo systemctl start mongod

# On Windows
net start MongoDB
```

### 4. Run the Application
```bash
go run main.go
```

### 5. Access the Application
- **Frontend**: http://localhost:9000
- **API**: http://localhost:9000/api/users

## 📚 API Documentation

### Base URL
```
http://localhost:9000/api
```

### Endpoints

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| `GET` | `/users` | Get all users | - |
| `GET` | `/users/:id` | Get user by ID | - |
| `POST` | `/users` | Create new user | `{"name": "string", "age": number, "gender": "string"}` |
| `DELETE` | `/users/:id` | Delete user | - |

### Example Requests

#### Get All Users
```bash
curl -X GET http://localhost:9000/api/users
```

#### Create User
```bash
curl -X POST http://localhost:9000/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","age":30,"gender":"male"}'
```

#### Delete User
```bash
curl -X DELETE http://localhost:9000/api/users/USER_ID
```

## 🗂️ Project Structure

```
go-mongodb-api/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── .gitignore             # Git ignore rules
├── README.md              # Project documentation
├── controllers/
│   └── user.go            # User controller with CRUD operations
├── models/
│   └── user.go            # User data model
└── frontend/
    └── index.html         # Frontend web interface
```

## 🎨 Frontend Features

- **Modern UI** - Beautiful gradient design with smooth animations
- **User Management** - Add, view, and delete users through web interface
- **Real-time Updates** - Automatic refresh after operations
- **Form Validation** - Client-side validation with error messages
- **Responsive Design** - Works on desktop, tablet, and mobile
- **Success/Error Messages** - User feedback for all operations

## 🔧 Configuration

### Database Configuration
The application connects to MongoDB on `localhost:27017` by default. To change this, modify the connection string in `main.go`:

```go
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
```

### Server Port
The server runs on port `9000` by default. To change this, modify the port in `main.go`:

```go
log.Fatal(http.ListenAndServe(":9000", handler))
```

## 🧪 Testing

### Test API Endpoints
```bash
# Test if server is running
curl http://localhost:9000/api/users

# Create a test user
curl -X POST http://localhost:9000/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","age":25,"gender":"other"}'
```

### Frontend Testing
1. Open http://localhost:9000 in your browser
2. Try adding a new user through the form
3. Verify the user appears in the list
4. Test the delete functionality

## 🚀 Deployment

### Local Development
```bash
go run main.go
```

### Build for Production
```bash
go build -o go-mongodb-api main.go
./go-mongodb-api
```

### Docker (Optional)
Create a `Dockerfile`:
```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/frontend ./frontend
EXPOSE 9000
CMD ["./main"]
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Go](https://golang.org/) - The Go programming language
- [MongoDB](https://www.mongodb.com/) - The database platform
- [httprouter](https://github.com/julienschmidt/httprouter) - HTTP router
- [MongoDB Go Driver](https://go.mongodb.org/mongo-driver) - Official MongoDB driver

## 📞 Support

If you have any questions or run into issues, please open an issue on GitHub.

---

**Happy Coding!** 🎉 