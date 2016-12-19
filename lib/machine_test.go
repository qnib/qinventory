package qinv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the instanciation of a PhysicalMachine
func TestNewMachine(t *testing.T) {
	exp := Machine{
		ID:       "host1",
		Hostname: "host1",
		Kernel:   "4.7.10",
	}
	got, _ := NewMachine("host1", "4.7.10")
	assert.Equal(t, exp, got)
}

// GetID test if the id is returned correctly
func TestMachineGetID(t *testing.T) {
	m, _ := NewMachine("host1", "4.7.10")
	got := m.GetID()
	assert.Equal(t, "host1", got)
}

// Test GetMe call, which returns a casted version
func TestMachineGetMe(t *testing.T) {
	m, _ := NewMachine("host1", "4.7.10")
	got := m.GetMe()
	assert.Equal(t, m.GetID(), got.GetID())
}
