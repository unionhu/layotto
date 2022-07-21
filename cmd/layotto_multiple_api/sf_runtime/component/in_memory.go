package component_sf

import (
	"context"
	"mosn.io/layotto/components/custom"
)

type inMemoryHello struct {
	ctx    context.Context
	config *custom.Config
}

func (i *inMemoryHello) Initialize(ctx context.Context, config custom.Config) error {
	i.ctx = ctx
	i.config = &config
	return nil
}

func (i *inMemoryHello) SayHello1(name string) (string, error) {
	return "Hello " + name, nil
}

func NewInMemoryHello() custom.Component {
	return &inMemoryHello{}
}
