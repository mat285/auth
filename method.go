package auth

import (
	"context"
	"errors"
)

var (
	ErrInvalidAuthType      = errors.New("invalid auth type")
	ErrInvalidAuthValueType = errors.New("invalid auth value type")
	ErrInvalidAuth          = errors.New("invalid auth")
)

type LoginMethod interface {
	Type() string
	New(Registration) (*StoredCreds, error)
	Verify(context.Context, StoredCreds, Creds) error
}

type Authenticator interface {
	Type() string
	Extract(context.Context, AuthValue) (string, error)
	Auth(context.Context, Identity) (AuthValue, error)
	Validate(context.Context, AuthValue) error
}

type Refresher interface {
	Auth(context.Context, AuthValue) (AuthValue, error)
}

type AuthValue interface {
	Type() string
}

type TokenPair struct {
	Access  string
	Refresh string
}
