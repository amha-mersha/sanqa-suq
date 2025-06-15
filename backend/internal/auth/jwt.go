package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey []byte
	issuer    string
}

type CustomClaims struct {
	UserID     string `json:"user_id"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Provider   string `json:"provider,omitempty"`
	ProviderID string `json:"provider_id,omitempty"`
	jwt.RegisteredClaims
}

type contextKey string

const UserClaimsKey contextKey = "user_claims"

func NewJWTService(secretKey, issuer string) *JWTService {
	return &JWTService{
		secretKey: []byte(secretKey),
		issuer:    issuer,
	}
}
func (s *JWTService) GenerateToken(userID, role, email, provider, providerID string) (string, error) {
	claims := CustomClaims{
		UserID:     userID,
		Role:       role,
		Email:      email,
		Provider:   provider,
		ProviderID: providerID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *JWTService) ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return s.secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}
	return nil, jwt.ErrTokenMalformed
}

func GetClaimsFromContext(ctx context.Context) (*CustomClaims, error) {
	claims, ok := ctx.Value(UserClaimsKey).(*CustomClaims)
	if !ok {
		return nil, errors.New("no claims found in context")
	}
	return claims, nil
}
