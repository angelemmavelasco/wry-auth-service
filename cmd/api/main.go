package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"auth-service/internal/auth"
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/handlers"
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

	_ = database.New(pool)

	tokenMgr, err := auth.NewTokenManager(
		"certs/private.pem",
		"certs/public.pem",
		"key-2026-v1",
		"auth-service",
	)
	if err != nil {
		log.Fatalf("error while initializing token service: %v", err)
	}
	h := &handlers.Handler{TokenMgr: tokenMgr}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/login", h.Login)
	mux.HandleFunc("/.well-known/jwks.json", h.JWKS)

	log.Print("Auth service running on 8081 port")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("internal server error: %v", err)
	}
}
