package efficientgo

import (
	"io/ioutil"
	"testing"
)

func BenchmarkMarshalDirect(b *testing.B) {
	// unmarshalDirect(os.Stdout)

	for i := 0; i < b.N; i++ {
		marshalDirect(ioutil.Discard)
	}
}

func BenchmarkMarshalWriter(b *testing.B) {
	// unmarshalWriter(os.Stdout)

	for i := 0; i < b.N; i++ {
		marshalWriter(ioutil.Discard)
	}
}
