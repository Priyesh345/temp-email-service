package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"temp-email-service/internal/handler"
	"temp-email-service/internal/service"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Root welcome route
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Temp Email Service API is running."}`))
	})

	r.Post("/email", handler.CreateEmailHandler)
	r.Get("/email/{emailID}", handler.GetInboxHandler)
	r.Post("/email/{emailID}/message", handler.AddMessageHandler)

	// Background task to auto-delete expired emails
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			service.CleanExpiredEmails()
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started at :" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {

		fmt.Println("Error starting server:", err)
	}
}
