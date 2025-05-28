package main

import (
	"context"
	"go-mongodb-api/controllers"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := httprouter.New()
	uc := controllers.NewUserController(getClient())

	// API routes
	router.GET("/api/users", uc.GetAllUsers)
	router.GET("/api/users/:id", uc.GetUser)
	router.POST("/api/users", uc.CreateUser)
	router.DELETE("/api/users/:id", uc.DeleteUser)

	// Serve static files (frontend)
	router.ServeFiles("/static/*filepath", http.Dir("frontend/"))

	// Root route - redirect to frontend
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.Redirect(w, r, "/static/index.html", http.StatusMovedPermanently)
	})

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)

	log.Println("Server starting on port 9000...")
	log.Println("API available at: http://localhost:9000/api/users")
	log.Println("Frontend available at: http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", handler))
}

func getClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Successfully connected to MongoDB!")
	return client
}
