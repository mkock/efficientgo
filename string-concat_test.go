package efficientgo

import "testing"

func BenchmarkConcatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatString()
	}
}

func BenchmarkConcatStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatStringBuilder()
	}
}
