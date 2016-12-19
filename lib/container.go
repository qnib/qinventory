package qinv

import (
	"github.com/docker/docker/api/types"
)

// Container holds informations about a container
type Container struct {
	types.Container
}

// NewContainer returns a container instance
func NewContainer(c types.Container) (Container, error) {
	return Container{
		Container: c,
	}, nil
}

// GetID returns the ID of the container, which will be used as unique identifier in the backends
func (c *Container) GetID() string {
	return c.ID
}
