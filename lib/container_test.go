package qinv

import (
    "testing"

    "github.com/docker/docker/api/types"
    "github.com/stretchr/testify/assert"
)

// Test the instanciation of a PhysicalMachine
func TestNewContainer(t *testing.T) {
    c := types.Container{
        ID: "container_id",
        Names: []string{"/container_name"},
    }
    m, _ := NewMachine("hostname","4.7.10")
    exp := Container{
        Container: c,
        Machine: m,
    }
    got, _ := NewContainer(c, m)
    assert.Equal(t, exp, got)
}
