package jwt

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/mat285/auth"
)

type TokenPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (t *TokenPair) Type() string {
	return AuthValueType
}

func (t TokenPair) IsZero() bool {
	return len(t.Access) == 0 && len(t.Refresh) == 0
}

func ValidateAccessToken(ctx context.Context, cfg Config, token string) error {
	claims, err := ExtractClaims(ctx, cfg, token)
	if err != nil {
		return err
	}
	if claims.Usage != "access" {
		return fmt.Errorf("Not an access token")
	}
	return nil
}

func ValidateRefreshToken(ctx context.Context, cfg Config, token string) error {
	claims, err := ExtractClaims(ctx, cfg, token)
	if err != nil {
		return err
	}
	if claims.Usage != "refresh" {
		return fmt.Errorf("Not a refresh token")
	}
	return nil
}

func ExtractClaims(ctx context.Context, cfg Config, token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return GetVerifyingKey(cfg)
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, fmt.Errorf("Token is Invalid")
	}
	return claims, nil
}

func GenerateTokenPair(ctx context.Context, cfg Config, ident auth.Identity) (*TokenPair, error) {
	access, err := GenerateAccessToken(ctx, cfg, GenerateClaims(ident), UserAccessTokenDuration)
	if err != nil {
		return nil, err
	}
	refresh, err := GenerateRefreshToken(ctx, cfg, GenerateClaims(ident), UserRefreshTokenDuration)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func GenerateAccessToken(ctx context.Context, cfg Config, claims *Claims, duration time.Duration) (string, error) {
	claims.Usage = "access"
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
	}
	return GenerateJWT(ctx, cfg, claims)
}

func GenerateRefreshToken(ctx context.Context, cfg Config, claims *Claims, duration time.Duration) (string, error) {
	claims.Usage = "refresh"
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
	}
	return GenerateJWT(ctx, cfg, claims)
}

func GenerateClaims(ident auth.Identity) *Claims {
	return &Claims{
		Type:     ClaimTypeUser,
		Identity: ident,
	}
}

func GenerateJWT(ctx context.Context, cfg Config, claims *Claims) (string, error) {
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(new(jwt.SigningMethodEd25519), claims)
	// Create the JWT string
	key, err := GetSigningKey(cfg)
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}

func GetSigningKey(cfg Config) (ed25519.PrivateKey, error) {
	return cfg.GetPrivateKey()
}

func GetVerifyingKey(cfg Config) (ed25519.PublicKey, error) {
	return cfg.GetPublicKey()
}
