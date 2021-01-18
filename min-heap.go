package efficientgo

import (
	"fmt"
	"time"
)

type event struct {
	name        string
	when        time.Time
	description string
}

func (e event) String() string {
	return fmt.Sprintf("%s: %s (%s)", e.when.Format("15:04:05"), e.name, e.description)
}

type events []event

func (e *events) Push(x interface{}) {
	*e = append(*e, x.(event))
}

func (e *events) Pop() interface{} {
	last := len(*e) - 1
	if last < 0 {
		return nil
	}
	x := (*e)[last]
	*e = (*e)[:last]
	return x
}

func (e *events) Len() int {
	return len(*e)
}

func (e *events) Less(i, j int) bool {
	return (*e)[i].when.Before((*e)[j].when)
}

func (e *events) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}
