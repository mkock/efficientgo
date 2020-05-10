package efficientgo

import (
	"sort"
	"testing"
)

var sorted []uint32 // Prevents compiler from optimizing out the function calls which assign to it.

func TestSortingMethods(t *testing.T) {
	objs := genObjs(100)

	// Bubble sort.
	sorted := bubbleSort(objs)
	if hasZero(sorted) {
		t.Error("bubbleSort is missing some numbers:", sorted)
	}
	if !sort.SliceIsSorted(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j]
	}) {
		t.Error("bubbleSort did not return sorted data:", sorted)
	}

	// Builtin sort.
	sorted = builtinSort(objs)
	if hasZero(sorted) {
		t.Error("builtinSort is missing some numbers:", sorted)
	}
	if !sort.SliceIsSorted(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j]
	}) {
		t.Error("builtinSort did not return sorted data:", sorted)
	}
}

func BenchmarkSorts100(b *testing.B) {
	// Using the same input for both benchmarks.
	b.StopTimer()
	objs := genObjs(100)
	b.StartTimer()

	b.Run("BenchmarkBubbleSort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sorted = bubbleSort(objs)
		}
	})

	b.Run("BenchmarkBuiltinSort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sorted = builtinSort(objs)
		}
	})
}

func BenchmarkSorts1000(b *testing.B) {
	// Using the same input for both benchmarks.
	b.StopTimer()
	objs := genObjs(1000)
	b.StartTimer()

	b.Run("BenchmarkBubbleSort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sorted = bubbleSort(objs)
		}
	})

	b.Run("BenchmarkBuiltinSort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sorted = builtinSort(objs)
		}
	})
}

func BenchmarkSorts10000(b *testing.B) {
	// Using the same input for both benchmarks.
	b.StopTimer()
	objs := genObjs(10000)
	b.StartTimer()

	b.Run("BenchmarkBubbleSort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sorted = bubbleSort(objs)
		}
	})

	b.Run("BenchmarkBuiltinSort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sorted = builtinSort(objs)
		}
	})
}
