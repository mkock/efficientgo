package efficientgo

import "testing"

var out []uint32

func BenchmarkUnpooledRandomizer(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		out = unpooledRandomizer()
	}
	b.Log(len(out))
}

func BenchmarkPooledRandomizer(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		out = pooledRandomizer()
	}
	b.Log(len(out))
}
