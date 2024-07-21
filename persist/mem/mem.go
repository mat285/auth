package mem

import (
	"context"
	"fmt"
	"sync"

	"github.com/mat285/auth"
)

type Mem struct {
	sync.Mutex
	data map[string]*auth.Identity
}

func New() auth.Persist {
	return &Mem{
		data: make(map[string]*auth.Identity),
	}
}

func (i *Mem) Get(_ context.Context, id string) (*auth.Identity, error) {
	i.Lock()
	defer i.Unlock()
	v := i.getUnsafe(id)
	if v == nil {
		return nil, fmt.Errorf("not found")
	}
	return v, nil
}

func (i *Mem) Create(_ context.Context, ident auth.Identity) (*auth.Identity, error) {
	i.Lock()
	defer i.Unlock()
	v := i.getUnsafe(ident.Identifier)
	if v != nil {
		return nil, fmt.Errorf("already exists")
	}
	i.data[ident.Identifier] = &ident
	return i.getUnsafe(ident.Identifier), nil
}

func (i *Mem) Update(_ context.Context, ident auth.Identity) (*auth.Identity, error) {
	i.Lock()
	defer i.Unlock()
	v := i.getUnsafe(ident.Identifier)
	if v == nil {
		return nil, fmt.Errorf("not found")
	}
	i.data[ident.Identifier] = &ident
	return i.getUnsafe(ident.Identifier), nil
}

func (i *Mem) Delete(_ context.Context, id string) error {
	i.Lock()
	defer i.Unlock()
	delete(i.data, id)
	return nil
}

func (i *Mem) getUnsafe(id string) *auth.Identity {
	v := i.data[id]
	return v
}
