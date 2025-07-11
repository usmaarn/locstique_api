package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/usmaarn/locstique_api/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
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

	//Routes
	v1Router.Get("/health-check", handlers.HealthCheckHandler)

	router.Mount("/v1", v1Router)

	fmt.Println("App running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
