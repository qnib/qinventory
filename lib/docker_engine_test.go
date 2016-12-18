package qinv

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// TestNewDockerEngine checks if constructor works
func TestNewDockerEngine(t *testing.T) {
    m := Machine{
        Hostname: "host1",
        Kernel: "4.7.10",
    }
    exp := DockerEngine{
        DockerHost: "unix:///var/run/docker.sock",
        Machine: m,
    }
    got, _ := NewDockerEngine("unix:///var/run/docker.sock", m)
    assert.Equal(t, exp, got)
}
