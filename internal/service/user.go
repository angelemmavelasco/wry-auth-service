package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"auth-service/internal/auth"
	"auth-service/internal/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("username or email already exists")
	ErrInvalidInput       = errors.New("invalid user input")
	ErrInvalidUUID        = errors.New("invalid user id format")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInactiveUser       = errors.New("user account is inactive")

	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9._-]{3,32}$`)
)

// UserDTO represents the safe public representation of a user without sensitive fields.
type UserDTO struct {
	ID              string    `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	IsActive        bool      `json:"is_active"`
	IsEmailVerified bool      `json:"is_email_verified"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// CreateUserInput represents the payload required to register/create a new user.
type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService struct {
	queries *database.Queries
}

func NewUserService(queries *database.Queries) *UserService {
	return &UserService{
		queries: queries,
	}
}

func (s *UserService) ListUsers(ctx context.Context, limit, offset int32) ([]UserDTO, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	users, err := s.queries.ListUsers(ctx, database.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	result := make([]UserDTO, 0, len(users))
	for _, u := range users {
		result = append(result, toUserDTO(u))
	}
	return result, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*UserDTO, error) {
	var pgUUID pgtype.UUID
	if err := pgUUID.Scan(strings.TrimSpace(id)); err != nil {
		return nil, ErrInvalidUUID
	}

	user, err := s.queries.GetUserByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	dto := toUserDTO(user)
	return &dto, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*UserDTO, error) {
	normalizedEmail, err := validateEmail(email)
	if err != nil {
		return nil, err
	}

	user, err := s.queries.GetUserByEmail(ctx, normalizedEmail)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	dto := toUserDTO(user)
	return &dto, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*UserDTO, error) {
	cleaned := strings.TrimSpace(username)
	if !usernameRegex.MatchString(cleaned) {
		return nil, fmt.Errorf("%w: username must be 3-32 characters (letters, numbers, '.', '_', '-')", ErrInvalidInput)
	}

	user, err := s.queries.GetUserByUsername(ctx, cleaned)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	dto := toUserDTO(user)
	return &dto, nil
}

func (s *UserService) CreateUser(ctx context.Context, input CreateUserInput) (*UserDTO, error) {
	username := strings.TrimSpace(input.Username)
	if !usernameRegex.MatchString(username) {
		return nil, fmt.Errorf("%w: username must be 3-32 characters (letters, numbers, '.', '_', '-')", ErrInvalidInput)
	}

	email, err := validateEmail(input.Email)
	if err != nil {
		return nil, err
	}

	passwordHash, err := auth.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	user, err := s.queries.CreateUser(ctx, database.CreateUserParams{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, ErrUserAlreadyExists
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	dto := toUserDTO(user)
	return &dto, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	var pgUUID pgtype.UUID
	if err := pgUUID.Scan(strings.TrimSpace(id)); err != nil {
		return ErrInvalidUUID
	}

	if _, err := s.queries.GetUserByID(ctx, pgUUID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFound
		}
		return fmt.Errorf("failed to verify user before deletion: %w", err)
	}

	if err := s.queries.DeleteUser(ctx, pgUUID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	// TODO(security): Revoke active refresh tokens/sessions associated with the deleted user once the refresh_tokens table is implemented.
	return nil
}

// Authenticate verifies a user's credentials (by username or email) against the stored Argon2id/Bcrypt hash.
// TODO(security): Implement MFA and account lockout after consecutive failed login attempts.
func (s *UserService) Authenticate(ctx context.Context, identifier, password string) (*UserDTO, error) {
	identifier = strings.TrimSpace(identifier)
	if identifier == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	var user database.User
	var err error
	if strings.Contains(identifier, "@") {
		user, err = s.queries.GetUserByEmail(ctx, identifier)
	} else {
		user, err = s.queries.GetUserByUsername(ctx, identifier)
	}

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("failed to query user for authentication: %w", err)
	}

	valid, err := auth.VerifyPassword(password, user.PasswordHash)
	if err != nil || !valid {
		return nil, ErrInvalidCredentials
	}

	if !user.IsActive {
		return nil, ErrInactiveUser
	}

	dto := toUserDTO(user)
	return &dto, nil
}

func validateEmail(raw string) (string, error) {
	cleaned := strings.TrimSpace(raw)
	if cleaned == "" || len(cleaned) > 255 {
		return "", fmt.Errorf("%w: invalid email length", ErrInvalidInput)
	}
	parsed, err := mail.ParseAddress(cleaned)
	if err != nil || parsed.Address != cleaned {
		return "", fmt.Errorf("%w: invalid email format", ErrInvalidInput)
	}
	return strings.ToLower(parsed.Address), nil
}

func toUserDTO(u database.User) UserDTO {
	return UserDTO{
		ID:              uuidToString(u.ID),
		Username:        u.Username,
		Email:           u.Email,
		IsActive:        u.IsActive,
		IsEmailVerified: u.IsEmailVerified,
		CreatedAt:       u.CreatedAt.Time,
		UpdatedAt:       u.UpdatedAt.Time,
	}
}

func uuidToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	b := u.Bytes
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

