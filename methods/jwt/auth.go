package jwt

import (
	"context"
	"time"

	"github.com/mat285/auth"
)

const (
	AuthType      = "jwt_auth"
	AuthValueType = "jwt_token_pair"
)

const (
	UserRefreshTokenDuration    = 24 * time.Hour
	UserAccessTokenDuration     = 5 * time.Minute
	ServiceRefreshTokenDuration = 24 * time.Hour
	ServiceAccessTokenDuration  = 1 * time.Hour
)

type Authenticator struct {
	Config Config
}

func New(cfg Config) *Authenticator {
	return &Authenticator{
		Config: cfg,
	}
}

func (*Authenticator) Type() string {
	return AuthType
}

func (a *Authenticator) Extract(ctx context.Context, v auth.AuthValue) (string, error) {
	if v.Type() != AuthValueType {
		return "", auth.ErrInvalidAuthValueType
	}
	typed, ok := v.(*TokenPair)
	if !ok {
		return "", auth.ErrInvalidAuth
	}
	claims, err := ExtractClaims(ctx, a.Config, typed.Access)
	if err != nil {
		return "", err
	}
	return claims.Identity.Identifier, nil
}

func (a *Authenticator) Auth(ctx context.Context, ident auth.Identity) (auth.AuthValue, error) {
	return GenerateTokenPair(ctx, a.Config, ident)
}

func (a *Authenticator) Validate(ctx context.Context, v auth.AuthValue) error {
	if v.Type() != AuthValueType {
		return auth.ErrInvalidAuthValueType
	}
	typed, ok := v.(*TokenPair)
	if !ok {
		return auth.ErrInvalidAuth
	}
	return ValidateAccessToken(ctx, a.Config, typed.Access)
}

func (a *Authenticator) Refresh(ctx context.Context, v auth.AuthValue) (auth.AuthValue, error) {
	if v.Type() != AuthValueType {
		return nil, auth.ErrInvalidAuthValueType
	}
	typed, ok := v.(*TokenPair)
	if !ok {
		return nil, auth.ErrInvalidAuth
	}
	claims, err := ExtractClaims(ctx, a.Config, typed.Refresh)
	if err != nil {
		return nil, err
	}
	return a.Auth(ctx, claims.Identity)
}
