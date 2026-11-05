-- name: GetUserByID :one
SELECT id, username, email, password_hash, is_active, is_email_verified, created_at, updated_at
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT id, username, email, password_hash, is_active, is_email_verified, created_at, updated_at
FROM users
WHERE LOWER(username) = LOWER($1) LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, is_active, is_email_verified, created_at, updated_at
FROM users
WHERE LOWER(email) = LOWER($1) LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password_hash
) VALUES (
    $1, $2, $3
)
RETURNING id, username, email, password_hash, is_active, is_email_verified, created_at, updated_at;

-- name: ListUsers :many
SELECT id, username, email, password_hash, is_active, is_email_verified, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUserEmailVerification :one
UPDATE users
SET is_email_verified = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, username, email, password_hash, is_active, is_email_verified, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

