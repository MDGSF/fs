package numgener

import "sync"

type TNumGenerator struct {
	sync.Mutex
	i uint64
}

func NewNumGenerator() *TNumGenerator {
	return &TNumGenerator{i: 0}
}

func (g *TNumGenerator) Num() uint64 {
	g.Lock()
	g.i++
	newnum := g.i
	g.Unlock()
	return newnum
}
