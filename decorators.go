package efficientgo

import (
	"fmt"
	"io"
)

// Adder is the interfacce for our Adder function and middleware.
type Adder interface {
	Add(x, y int) int
}

// AdderMiddleware exists to support middleware functions on our Adder implementation.
type AdderMiddleware func(Adder) Adder

// CalcAdd is our implementation of the Adder interface.
type CalcAdd func(x, y int) int

// Add adds the given two numbers together and returns the result.
func (c CalcAdd) Add(x, y int) int {
	return c(x, y)
}

// WithLogging adds logging to our Adder implementation.
func WithLogging(out io.Writer) AdderMiddleware {
	return func(a Adder) Adder {
		fn := func(x, y int) int {
			res := a.Add(x, y)
			fmt.Fprintf(out, "Result of Add(%d, %d) = %d", x, y, res)
			return res
		}
		return CalcAdd(fn)
	}
}

// WithEvent sends an event on the given channel whenever Add() has been called.
func WithEvent(events chan string) AdderMiddleware {
	return func(a Adder) Adder {
		fn := func(x, y int) int {
			res := a.Add(x, y)
			events <- "CalcAdd called!"
			return res
		}
		return CalcAdd(fn)
	}
}
