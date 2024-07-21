package jwt

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/mat285/auth"
)

type Claims struct {
	Type     ClaimType     `json:"type"`
	Identity auth.Identity `json:"identity"`
	Usage    string        `json:"usage"`

	jwt.RegisteredClaims
}

type ServiceIdentity struct {
	Name string `json:"name"`
}

type ClaimType string

const (
	ClaimTypeUnknown ClaimType = ""
	ClaimTypeUser    ClaimType = "user"
)
