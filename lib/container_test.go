package qinv

import (
    "testing"

    "github.com/docker/docker/api/types"
    "github.com/stretchr/testify/assert"
)

// Test the instanciation of a PhysicalMachine
func TestNewContainer(t *testing.T) {
    m, _ := NewMachine("hostname","4.7.10")
    d := DockerEngine{
        DockerHost: "unix:///var/run/docker.sock",
        Machine: m,
    }
    c := types.Container{
        ID: "container_id",
        Names: []string{"/container_name"},
    }
    exp := Container{
        Container: c,
        Engine: d,
    }
    got, _ := NewContainer(c, d)
    assert.Equal(t, exp, got)
}
