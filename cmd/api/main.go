package main

import (
	"log"
	"net/http"

	"auth-service/internal/auth"
	"auth-service/internal/handlers"
)

func main() {
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
		log.Fatalf("internal error server: %v", err)
	}
}
