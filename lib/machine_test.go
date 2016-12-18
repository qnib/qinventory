package qinv

import (
    "testing"
    "github.com/stretchr/testify/assert"
)
// Test the instanciation of a PhysicalMachine
func TestNewMachine(t *testing.T) {
    exp := Machine{
        Hostname: "host1",
        Kernel: "4.7.10",
    }
    got, _ := NewMachine("host1","4.7.10")
    assert.Equal(t, exp, got)
}
