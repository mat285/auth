package basic

import (
	"context"
	"crypto/subtle"
	"fmt"

	"github.com/mat285/auth"
	"github.com/mat285/auth/pkg/crypto"
)

type Basic struct {
}

func New() auth.LoginMethod {
	return &Basic{}
}

func (*Basic) Type() string {
	return "basic"
}

func (n *Basic) New(reg auth.Registration) (*auth.StoredCreds, error) {
	if len(reg.Creds.Password) == 0 {
		return nil, fmt.Errorf("no password")
	}

	salt, err := crypto.GenerateSalt()
	if err != nil {
		return nil, err
	}
	return &auth.StoredCreds{
		PasswordHash: crypto.Hash(reg.Creds.Password, salt),
		PasswordSalt: salt,
	}, nil
}

func (n *Basic) Verify(ctx context.Context, stored auth.StoredCreds, cred auth.Creds) error {
	res := subtle.ConstantTimeCompare(stored.PasswordHash, crypto.Hash(cred.Password, stored.PasswordSalt))
	if res != 1 {
		return fmt.Errorf("invalid credentials")
	}
	return nil
}
