package main

import (
	"context"
	"fmt"

	"github.com/blend/go-sdk/uuid"
	"github.com/mat285/auth"
	"github.com/mat285/auth/login/basic"
	"github.com/mat285/auth/methods/jwt"
	"github.com/mat285/auth/persist/mem"
	"github.com/mat285/auth/pkg/crypto"
)

func main() {
	ctx := context.Background()
	manager, err := New(auth.Config{})
	if err != nil {
		panic(err)
	}

	r1 := auth.Registration{
		Creds: auth.Creds{
			Identifier: uuid.V4().String(),
			Password:   []byte("pass1"),
		},
	}

	r2 := auth.Registration{
		Creds: auth.Creds{
			Identifier: uuid.V4().String(),
			Password:   []byte("pass2"),
		},
	}

	id1, err := manager.Register(ctx, r1, nil)
	if err != nil {
		panic(err)
	}

	out1, av, err := manager.Login(ctx, r1.Creds)
	if err != nil {
		panic(err)
	}
	if out1.Identifier != id1.Identifier {
		panic(fmt.Sprintf("Expected %s but got %s ident", id1.Identifier, out1.Identifier))
	}
	if av.Type() != jwt.AuthValueType {
		panic(fmt.Sprintf("Unknown auth value returned %q", av))
	}
	fmt.Println(manager.Identify(ctx, av))

	_, err = manager.Register(ctx, r2, nil)
	if err != nil {
		panic(err)
	}

}

func jwtConfig() jwt.Config {
	var err error
	cfg := jwt.Config{}
	cfg.PrivateKey, err = crypto.ParsePemEd25519PrivateKey([]byte(privateKey))
	if err != nil {
		panic(err)
	}
	cfg.PublicKey, err = crypto.ParsePemEd25519PublicKey([]byte(publicKey))
	if err != nil {
		panic(err)
	}
	return cfg
}

func New(cfg auth.Config) (*auth.Manager, error) {
	return auth.New(
		auth.OptConfig(cfg),
		auth.OptPersist(mem.New()),
		auth.OptLogin(basic.New()),
		auth.OptAuthenticator(jwt.New(jwtConfig())),
	)
}

const (
	privateKey = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIJ5tnQ+IGntkQbTtX2+8Mqz9WiWkHicZcGP79RrvUztZ
-----END PRIVATE KEY-----
`

	publicKey = `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEAwF0E/+gNtw181XUMHenb+fjlXCJgP9womllqu4fKrlo=
-----END PUBLIC KEY-----

`
)
