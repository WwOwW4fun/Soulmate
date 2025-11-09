package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"planner/api"
	"planner/middleware"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server := http.NewServeMux()

	server.Handle("/chatguided", middleware.Middleware(api.ChatGuidedHandler()))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("CLIENT_URL")}, // http://localhost:5173
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
	})

	handler := c.Handler(server)

	fmt.Printf("Listening at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}