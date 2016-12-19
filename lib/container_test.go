package qinv

import (
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/stretchr/testify/assert"
)

// Test the instanciation of a PhysicalMachine
func TestNewContainer(t *testing.T) {
	c := types.Container{
		ID:    "container_id",
		Names: []string{"/container_name"},
	}
	exp := Container{
		Container: c,
	}
	got, _ := NewContainer(c)
	assert.Equal(t, exp, got)
}

// TestGetID checks if the ID is returned correctly
func TestGetID(t *testing.T) {
	c := types.Container{
		ID:    "container_id",
		Names: []string{"/container_name"},
	}
	cnt, _ := NewContainer(c)
	got := cnt.GetID()
	assert.Equal(t, "container_id", got)
}
