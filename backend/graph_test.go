package backend

import (
	"testing"

	"github.com/mkindahl/gograph/directed"
	"github.com/stretchr/testify/assert"
)

// Test the instanciation of a PhysicalMachine
func TestNewQGraph(t *testing.T) {
	exp := QGraph{
		Graph: directed.New(),
	}
	got, _ := NewQGraph()
	assert.Equal(t, exp, got)
}
