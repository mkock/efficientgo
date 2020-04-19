package efficientgo

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	p.New = func() interface{} {
		return new(random)
	}
}

const size = 500_000

var p sync.Pool

type random struct {
	a, b uint32
}

func unpooledRandomizer() []uint32 {
	out := make([]uint32, size)

	for i := 0; i < size; i++ {
		r := random{rand.Uint32(), rand.Uint32()}
		out[i] = r.a + r.b
	}

	return out
}

func pooledRandomizer() []uint32 {
	out := make([]uint32, size)

	for i := 0; i < size; i++ {
		r := p.Get().(*random)
		r.a, r.b = rand.Uint32(), rand.Uint32()
		out[i] = r.a + r.b

		p.Put(r)
	}

	return out
}
