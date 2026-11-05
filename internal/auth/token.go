package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

)

type TokenManager struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	KeyID      string
	Issuer     string
}

func NewTokenManager(privPath, pubPath, keyID, issuer string) (*TokenManager, error) {
	privBytes, err := os.ReadFile(privPath)
	if err != nil {
		return nil, fmt.Errorf("it was not possible to read private key: %w", err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return nil, fmt.Errorf("it was not possible to parse private key: %w", err)
	}
	pubBytes, err := os.ReadFile(pubPath)
	if err != nil {
		return nil, fmt.Errorf("it was not possible to read public key: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, fmt.Errorf("it was not posible to parse public key: %w", err)
	}

	return &TokenManager{
		PublicKey:  pubKey,
		PrivateKey: privKey,
		KeyID:      keyID,
		Issuer:     issuer,
	}, nil
}

type UserClaims struct {
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

func (m *TokenManager) GenerateToken(userID string, audience string, roles []string) (string, error) {
	now := time.Now().UTC()
	claims := UserClaims{
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: m.Issuer,
			Subject: userID,
			Audience: jwt.ClaimStrings{audience},
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
			IssuedAt: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = m.KeyID

	return token.SignedString(m.PrivateKey)
}
