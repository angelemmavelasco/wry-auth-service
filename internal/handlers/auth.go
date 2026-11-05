package handlers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"net/http"

	"auth-service/internal/auth"
	"auth-service/internal/service"
)

const maxRequestBodyBytes = 1 << 20 // 1 MB

type Handler struct {
	TokenMgr    *auth.TokenManager
	UserService *service.UserService
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppID    string `json:"app_id"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)
	defer r.Body.Close()

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json payload")
		return
	}

	if req.AppID == "" {
		writeError(w, http.StatusBadRequest, "app_id is required")
		return
	}

	// TODO(security): Enforce per-IP and per-account rate limiting on login attempts.
	user, err := h.UserService.Authenticate(r.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) || errors.Is(err, service.ErrInactiveUser) {
			writeError(w, http.StatusUnauthorized, "invalid credentials")
			return
		}
		log.Printf("login error: %v", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	// TODO(security): Load user roles dynamically per app_id from user_app_roles table.
	roles := []string{"user"}
	tokenString, err := h.TokenMgr.GenerateToken(user.ID, req.AppID, roles)
	if err != nil {
		log.Printf("token generation error: %v", err)
		writeError(w, http.StatusInternalServerError, "error while generating token")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"access_token": tokenString,
		"token_type":   "Bearer",
	})
}

func (h *Handler) JWKS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	pubKey := h.TokenMgr.PublicKey

	nStr := base64.RawURLEncoding.EncodeToString(pubKey.N.Bytes())
	eBytes := big.NewInt(int64(pubKey.E)).Bytes()
	eStr := base64.RawURLEncoding.EncodeToString(eBytes)
	resp := map[string]interface{}{
		"keys": []map[string]string{
			{
				"kty": "RSA",
				"use": "sig",
				"kid": h.TokenMgr.KeyID,
				"alg": "RS256",
				"n":   nStr,
				"e":   eStr,
			},
		},
	}

	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
