package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"auth-service/internal/service"
)

// ListUsers handles GET /api/v1/users
// Also supports filtering via query parameters ?email=... or ?username=...
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	if email := r.URL.Query().Get("email"); email != "" {
		h.getUserByEmailValue(w, r, email)
		return
	}
	if username := r.URL.Query().Get("username"); username != "" {
		h.getUserByUsernameValue(w, r, username)
		return
	}

	var limit int32 = 20
	var offset int32 = 0

	if lStr := r.URL.Query().Get("limit"); lStr != "" {
		parsed, err := strconv.ParseInt(lStr, 10, 32)
		if err != nil || parsed <= 0 {
			writeError(w, http.StatusBadRequest, "invalid limit parameter")
			return
		}
		limit = int32(parsed)
	}

	if oStr := r.URL.Query().Get("offset"); oStr != "" {
		parsed, err := strconv.ParseInt(oStr, 10, 32)
		if err != nil || parsed < 0 {
			writeError(w, http.StatusBadRequest, "invalid offset parameter")
			return
		}
		offset = int32(parsed)
	}

	users, err := h.UserService.ListUsers(r.Context(), limit, offset)
	if err != nil {
		log.Printf("error listing users: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to list users")
		return
	}

	writeJSON(w, http.StatusOK, users)
}

// GetUserByEmail handles GET /api/v1/users/email/{email}
func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	email := r.PathValue("email")
	if email == "" {
		email = r.URL.Query().Get("email")
	}

	h.getUserByEmailValue(w, r, email)
}

func (h *Handler) getUserByEmailValue(w http.ResponseWriter, r *http.Request, email string) {
	user, err := h.UserService.GetUserByEmail(r.Context(), email)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			writeError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			writeError(w, http.StatusNotFound, "user not found")
		default:
			log.Printf("error fetching user by email: %v", err)
			writeError(w, http.StatusInternalServerError, "failed to get user")
		}
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// GetUserByUsername handles GET /api/v1/users/username/{username}
func (h *Handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	username := r.PathValue("username")
	if username == "" {
		username = r.URL.Query().Get("username")
	}

	h.getUserByUsernameValue(w, r, username)
}

func (h *Handler) getUserByUsernameValue(w http.ResponseWriter, r *http.Request, username string) {
	user, err := h.UserService.GetUserByUsername(r.Context(), username)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			writeError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			writeError(w, http.StatusNotFound, "user not found")
		default:
			log.Printf("error fetching user by username: %v", err)
			writeError(w, http.StatusInternalServerError, "failed to get user")
		}
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// GetUserByID handles GET /api/v1/users/{id}
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := r.PathValue("id")
	user, err := h.UserService.GetUserByID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidUUID):
			writeError(w, http.StatusBadRequest, "invalid user id format")
		case errors.Is(err, service.ErrUserNotFound):
			writeError(w, http.StatusNotFound, "user not found")
		default:
			log.Printf("error fetching user by id: %v", err)
			writeError(w, http.StatusInternalServerError, "failed to get user")
		}
		return
	}

	writeJSON(w, http.StatusOK, user)
}

// CreateUser handles POST /api/v1/users
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)
	defer r.Body.Close()

	var input service.CreateUserInput
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json payload")
		return
	}

	user, err := h.UserService.CreateUser(r.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidInput):
			writeError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, service.ErrUserAlreadyExists):
			writeError(w, http.StatusConflict, err.Error())
		default:
			log.Printf("error creating user: %v", err)
			writeError(w, http.StatusInternalServerError, "failed to create user")
		}
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

// DeleteUser handles DELETE /api/v1/users/{id}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := r.PathValue("id")
	if id == "" {
		id = r.URL.Query().Get("id")
	}

	if err := h.UserService.DeleteUser(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidUUID):
			writeError(w, http.StatusBadRequest, "invalid user id format")
		case errors.Is(err, service.ErrUserNotFound):
			writeError(w, http.StatusNotFound, "user not found")
		default:
			log.Printf("error deleting user: %v", err)
			writeError(w, http.StatusInternalServerError, "failed to delete user")
		}
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "user deleted successfully",
	})
}

