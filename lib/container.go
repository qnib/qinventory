package qinv
import (
    "github.com/docker/docker/api/types"
)
// Container holds informations about a container
type Container struct {
    types.Container
    Engine DockerEngine
}

// NewContainer returns a container instance
func NewContainer(c types.Container, e DockerEngine) (Container, error) {
    return Container{
        Container: c,
        Engine: e,
    }, nil
}
