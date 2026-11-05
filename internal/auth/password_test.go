package auth

import (
	"strings"
	"testing"
)

func TestHashAndVerifyPasswordArgon2id(t *testing.T) {
	password := "SuperSecretPassword123!"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	if !strings.HasPrefix(hash, "$argon2id$v=19$m=65536,t=3,p=4$") {
		t.Fatalf("expected argon2id PHC prefix, got: %s", hash)
	}

	ok, err := VerifyPassword(password, hash)
	if err != nil {
		t.Fatalf("unexpected error verifying password: %v", err)
	}
	if !ok {
		t.Fatal("expected password verification to succeed")
	}

	ok, err = VerifyPassword("WrongPassword123!", hash)
	if err != nil {
		t.Fatalf("unexpected error verifying wrong password: %v", err)
	}
	if ok {
		t.Fatal("expected wrong password verification to fail")
	}
}

func TestHashAndVerifyPasswordBcrypt(t *testing.T) {
	password := "Password123!"

	hash, err := HashPasswordBcrypt(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password with bcrypt: %v", err)
	}

	if !strings.HasPrefix(hash, "$2a$12$") && !strings.HasPrefix(hash, "$2b$12$") {
		t.Fatalf("expected bcrypt cost 12 prefix, got: %s", hash)
	}

	ok, err := VerifyPassword(password, hash)
	if err != nil {
		t.Fatalf("unexpected error verifying bcrypt password: %v", err)
	}
	if !ok {
		t.Fatal("expected bcrypt password verification to succeed")
	}
}

func TestPasswordStrengthValidation(t *testing.T) {
	if _, err := HashPassword("short"); err == nil {
		t.Fatal("expected error for password shorter than 8 characters")
	}
}


