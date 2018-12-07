package numgener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumGenerator(t *testing.T) {
	gen := NewNumGenerator()
	n := gen.Num()
	assert.Equal(t, uint64(1), n, "they should be equal")
}
