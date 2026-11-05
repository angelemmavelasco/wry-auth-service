package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"auth-service/internal/auth"
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/handlers"
	"auth-service/internal/service"
)

func main() {
	fmt.Println("Initializing WRY service...")
	dbBaseConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error while loading db environment variables: %v", err)
	}

	ctx := context.Background()

	dbConnector := database.NewConnector(dbBaseConfig)
	pool, err := dbConnector.Connect(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer pool.Close()

	log.Println("Database connection established successfully")

	queries := database.New(pool)
	userService := service.NewUserService(queries)

	tokenMgr, err := auth.NewTokenManager(
		"certs/private.pem",
		"certs/public.pem",
		"key-2026-v1",
		"auth-service",
	)
	if err != nil {
		log.Fatalf("error while initializing token service: %v", err)
	}

	h := &handlers.Handler{
		TokenMgr:    tokenMgr,
		UserService: userService,
	}

	mux := http.NewServeMux()

	// Auth & JWKS routes
	mux.HandleFunc("POST /api/v1/login", h.Login)
	mux.HandleFunc("GET /.well-known/jwks.json", h.JWKS)

	// User routes
	mux.HandleFunc("GET /api/v1/users", h.ListUsers)
	mux.HandleFunc("POST /api/v1/users", h.CreateUser)
	mux.HandleFunc("GET /api/v1/users/email/{email}", h.GetUserByEmail)
	mux.HandleFunc("GET /api/v1/users/username/{username}", h.GetUserByUsername)
	mux.HandleFunc("GET /api/v1/users/{id}", h.GetUserByID)
	mux.HandleFunc("DELETE /api/v1/users/{id}", h.DeleteUser)

	srv := &http.Server{
		Addr:         "127.0.0.1:8081",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Print("Auth service running on 127.0.0.1:8081")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("internal server error: %v", err)
	}
}
