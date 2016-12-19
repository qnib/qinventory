package qinv

import (
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
)

// Test the instanciation of a PhysicalMachine
func TestNewContainer(t *testing.T) {
	d := DockerEngine{
		DockerHost: "unix:///var/run/docker.sock",
	}
	c := types.Container{
		ID:    "container_id",
		Names: []string{"/container_name"},
	}
	exp := Container{
		Container: c,
	}
	got, _ := NewContainer(c, d)
	assert.Equal(t, exp, got)
}
