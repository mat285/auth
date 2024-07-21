package auth

import (
	"context"
)

type Manager struct {
	Config      Config
	Persist     Persist
	LoginMethod LoginMethod
	Auth        Authenticator
}

func New(opts ...Option) (*Manager, error) {
	m := &Manager{}

	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *Manager) Register(ctx context.Context, reg Registration, idata interface{}) (*Identity, error) {
	sc, err := m.LoginMethod.New(reg)
	if err != nil {
		return nil, err
	}
	ident := NewIdentity(reg.Identifier, *sc, idata)
	return m.Persist.Create(ctx, ident)
}

func (m *Manager) Login(ctx context.Context, creds Creds) (*Identity, AuthValue, error) {
	ident, err := m.Persist.Get(ctx, creds.Identifier)
	if err != nil {
		return nil, nil, err
	}
	err = m.LoginMethod.Verify(ctx, ident.Creds, creds)
	if err != nil {
		return nil, nil, err
	}

	tokens, err := m.Auth.Auth(ctx, *ident)
	if err != nil {
		return nil, nil, err
	}
	return ident, tokens, nil
}

func (m *Manager) Identify(ctx context.Context, v AuthValue) (*Identity, error) {
	id, err := m.Auth.Extract(ctx, v)
	if err != nil {
		return nil, err
	}
	return m.Persist.Get(ctx, id)
}
