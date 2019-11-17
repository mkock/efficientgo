package efficientgo

import (
	"bytes"
	"testing"
)

func TestCalcAdd(t *testing.T) {
	a := CalcAdd(func(x, y int) int {
		return x + y
	})
	actual := a.Add(10, 20)
	if actual != 30 {
		t.Fatalf("expected a.Add(10, 20) to equal 30, got %d", actual)
	}
}

func TestCalcAddWithLogging(t *testing.T) {
	a := CalcAdd(func(x, y int) int {
		return x + y
	})

	var buf bytes.Buffer
	b := WithLogging(&buf)(a)

	actual := b.Add(5, 8)
	if actual != 13 {
		t.Fatalf("expected a.Add(5, 8) to equal 13, got %d", actual)
	}
	log := buf.String()
	expectLog := "Result of Add(5, 8) = 13"
	if log != expectLog {
		t.Fatalf("expected buffer to contain %q, got %q", expectLog, log)
	}
}

func TestCalcAddWithEvent(t *testing.T) {
	a := CalcAdd(func(x, y int) int {
		return x + y
	})

	events := make(chan string, 1)
	defer close(events)

	b := WithEvent(events)(a)
	actual := b.Add(30, 15)
	if actual != 45 {
		t.Fatalf("expected a.Add(30, 15) to equal 45, got %d", actual)
	}
	evt := <-events
	if evt != "CalcAdd called!" {
		t.Fatalf("expected to receive an event containing %q, got %v", "CalcAdd called!", evt)
	}
}
