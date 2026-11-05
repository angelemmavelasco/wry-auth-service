package handlers

import (
	"encoding/base64"
	"encoding/json"
	"math/big"
	"net/http"

	"auth-service/internal/auth"
)

type Handler struct {
	TokenMgr *auth.TokenManager
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppID    string `json:"app_id"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Username != "test_user" || req.Password != "secret123" {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	roles := []string{"admin", "editor"}
	tokenString, err := h.TokenMgr.GenerateToken("user_1", req.AppID, roles)
	if err != nil {
		http.Error(w, "error while token generation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"acess_token": tokenString,
		"token_type":  "Bearer",
	})

}

func (h *Handler) JWKS(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
