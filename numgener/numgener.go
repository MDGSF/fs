package numgener

import (
	"sync"
	"sync/atomic"
)

type TNumGenerator struct {
	sync.Mutex
	i uint64
}

func NewNumGenerator() *TNumGenerator {
	return &TNumGenerator{i: 0}
}

func (g *TNumGenerator) Num() uint64 {
	return atomic.AddUint64(&g.i, 1)
}
