package efficientgo

import (
	"container/heap"
	"fmt"
	"testing"
	"time"
)

func TestMinHeap(t *testing.T) {
	e1 := event{
		name:        "Birthday",
		when:        time.Now().Add(-24 * time.Hour),
		description: "Date of birth",
	}
	e2 := event{
		name:        "Graduation",
		when:        time.Now().Add(-20 * time.Hour),
		description: "Graduation from university",
	}
	e3 := event{
		name:        "Wedding",
		when:        time.Now().Add(-15 * time.Hour),
		description: "Getting married",
	}
	e4 := event{
		name:        "Career",
		when:        time.Now().Add(-10 * time.Hour),
		description: "Professional career taking off",
	}
	e5 := event{
		name:        "Funeral",
		when:        time.Now().Add(-6 * time.Hour),
		description: "Sudden disease",
	}
	e6 := event{
		name:        "Resurrection",
		when:        time.Now().Add(-1 * time.Hour),
		description: "Back from the dead",
	}

	ee := events{e5, e2}
	heap.Init(&ee) // Must be called for the initial slice to get correct order in the following printout.
	for ee.Len() > 0 {
		fmt.Println(heap.Pop(&ee).(event))
	}

	heap.Push(&ee, e1)
	heap.Push(&ee, e4)
	heap.Push(&ee, e3)
	heap.Push(&ee, e6)

	for ee.Len() > 0 {
		fmt.Println(heap.Pop(&ee).(event))
	}
}
