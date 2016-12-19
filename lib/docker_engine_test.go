package qinv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewDockerEngine checks if constructor works
func TestNewDockerEngine(t *testing.T) {
	exp := DockerEngine{
		DockerHost: "unix:///var/run/docker.sock",
	}
	got, _ := NewDockerEngine("unix:///var/run/docker.sock")
	assert.Equal(t, exp, got)
}
