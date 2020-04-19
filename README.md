# Efficient Go

This is just a bunch of code experiments in Go, some of which are benchmarked and documented here for educational and self-awareness purposes.

## JSON Marshal

Marshalling a simple struct into JSON using two strategies:

1. `MarshalDirect`: the usual approach of marshalling into a byte slice and then printing that.
2. `MarshalWriter`: avoiding the temporary buffer by writing directly to a `Writer`.

Metrics:

```bash
goos: linux
goarch: amd64
pkg: github.com/mkock/efficient-go
BenchmarkMarshalDirect-8   	  300000	      4320 ns/op	     480 B/op	       7 allocs/op
BenchmarkMarshalWriter-8   	 1000000	      1961 ns/op	     304 B/op	       5 allocs/op
PASS
coverage: 80.0% of statements
ok  	github.com/mkock/efficient-go	3.331s
Success: Benchmarks passed.
```

## String Concatenation

Building a string from a series of random text snippets using two strategies:

1. `ConcatString`: the usual approach of constructing a new string from the concatenation of two others.
2. `ConcatStringBuilder`: avoiding the overhead of allocation by leveraging the `strings.Builder` from the stdlib.

Metrics:

```bash
goos: linux
goarch: amd64
pkg: github.com/mkock/efficient-go
BenchmarkConcatString-8          	       1	23900002999 ns/op	208581365520 B/op	   54598 allocs/op
BenchmarkConcatStringBuilder-8   	     300	    5072790 ns/op	    45183032 B/op	      38 allocs/op
PASS
ok  	github.com/mkock/efficient-go	25.918s
Success: Benchmarks passed.
```

## Decorators

There's a really interesting blog post that shows how decorators/middleware could work in a way that respects the
single responsibility and open/closed principles. Link: [bartfokker.com](https://bartfokker.com/posts/decorators/).

This experiment is mostly a reproduction of the author's examples for my own reference.

Main concepts:
- There must be two interfaces: one for working with the implementation, and one for working with middlewares.
- Middlewares can take custom arguments, but are required to return (a modified version of) a type that satisfies the interface for the implemented solution.
- Middlewares return a function that redefines the implemented solution in order to add additional functionality to it.


## sync.Pool

Iterating over a struct and using it for managing two random numbers to generate
a result, using two strategies:

1. `unpooledRandomizer`: allocates a struct per iteration.
2. `pooledRandomizer`: uses sync.Pool to reduce allocations be recycling them.

Metrics:

```bash
goos: windows
goarch: amd64
pkg: github.com/mkock/efficientgo
BenchmarkUnpooledRandomizer-8                 36          36206583 ns/op         2007138 B/op          1 allocs/op
--- BENCH: BenchmarkUnpooledRandomizer-8
    pool_test.go:13: 500000
    pool_test.go:13: 500000
    pool_test.go:13: 500000
BenchmarkPooledRandomizer-8                   27          46013848 ns/op         2007675 B/op          3 allocs/op
--- BENCH: BenchmarkPooledRandomizer-8
    pool_test.go:22: 500000
    pool_test.go:22: 500000
PASS
ok      github.com/mkock/efficientgo    3.961s
```
