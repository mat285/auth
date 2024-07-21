package auth

import (
	"context"
)

type Persist interface {
	Get(context.Context, string) (*Identity, error)
	Create(context.Context, Identity) (*Identity, error)
	Update(context.Context, Identity) (*Identity, error)
	Delete(context.Context, string) error
}
