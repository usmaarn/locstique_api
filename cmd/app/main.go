package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/usmaarn/locstique_api/internal/config"
	"github.com/usmaarn/locstique_api/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env: ", err)
	}

	config.InitializeValidator()
	db := config.InitializeDatabase()
	ctx := context.Background()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			env := os.Getenv("APP_ENV")
			if env == "production" {
				return true
			}
			return false
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	handler := handlers.NewHandler(ctx, db)

	//Routes
	v1Router.Get("/health-check", handler.HealthCheckHandler)

	//Authentication
	v1Router.Post("/auth/register", handler.RegisterUserHandler)

	router.Mount("/v1", v1Router)

	fmt.Println("App running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
