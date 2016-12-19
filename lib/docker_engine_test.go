package qinv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewDockerEngine checks if constructor works
func TestNewDockerEngine(t *testing.T) {
	exp := DockerEngine{
		ID:         "id1",
		DockerHost: "unix:///var/run/docker.sock",
	}
	got, _ := NewDockerEngine("id1", "unix:///var/run/docker.sock")
	assert.Equal(t, exp, got)
}

// GetID test if the id is returned correctly
func TestEngineGetID(t *testing.T) {
	e, _ := NewDockerEngine("id1", "unix:///var/run/docker.sock")
	got := e.GetID()
	assert.Equal(t, "id1", got)
}
