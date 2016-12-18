package qinv
import (
    "github.com/docker/docker/api/types"
)
// Container holds informations about a container
type Container struct {
    types.Container
    Machine Machine
}

// NewContainer returns a container instance
func NewContainer(c types.Container, m Machine) (Container, error) {
    return Container{
        Container: c,
        Machine: m,
    }, nil
}
