package efficientgo

import (
	"math/rand"
	"sort"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type obj struct {
	id   uint32
	data string
}

func genObjs(nr int) []obj {
	nrs := make([]obj, nr)
	for i := 0; i < nr; i++ {
		nrs[i] = obj{uint32(rand.Intn(999) + 1), strings.Repeat("xyz123", rand.Intn(100))}
	}
	return nrs
}

func bubbleSort(in []obj) []uint32 {
	if len(in) == 0 {
		return []uint32{} // Edge case.
	}
	out := make([]uint32, len(in))
	out[0] = in[0].id

	for i := 1; i < len(in); i++ {
		out[i] = in[i].id
		for j := i; j > 0 && out[j] > out[j-1]; j-- {
			out[j], out[j-1] = out[j-1], out[j] // Bubble down.
		}
	}

	return out
}

func builtinSort(in []obj) []uint32 {
	ids := make([]uint32, 0, len(in))

	for _, o := range in {
		ids = append(ids, o.id)
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[i] > ids[j]
	})

	return ids
}

func hasZero(in []uint32) bool {
	for _, nr := range in {
		if nr == 0 {
			return true
		}
	}
	return false
}
