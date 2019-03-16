package main

import "fmt"

type stream_t struct {
	N uint64
	Q uint64
	a uint64
	b uint64
	x uint64
}

type usage_t struct {
	pow   uint64
	index uint64
}

type query_t struct {
	B uint64
	E uint64
	K uint64
}

func (stream stream_t) generator() ([]usage_t, []query_t) {
	var i uint64
	x := stream.x
	usage := make([]usage_t, stream.N)
	query := make([]query_t, stream.Q)

	for i = 0; i < stream.N; i++ {
		x = (x*stream.a + stream.b) % (10E9 + 7)
		usage[i].pow = x
		usage[i].index = i
	}
	for i = 0; i < stream.Q; i++ {
		x = (x*stream.a + stream.b) % (10E9 + 7)
		query[i].B = x
		x = (x*stream.a + stream.b) % (10E9 + 7)
		query[i].E = x
		if query[i].E < query[i].B {
			swp := query[i].E
			query[i].E = query[i].B
			query[i].B = swp
		}
		x = (x*stream.a + stream.b) % (10E9 + 7)
		if x > query[i].E-query[i].B+1 {
			query[i].K = x % (query[i].E - query[i].B - 1)
		} else {
			query[i].K = x
		}
	}
	return usage, query
}

func main() {
	var T uint8
	var N, Q, a, b, x uint64
	var i uint64

	fmt.Scanf("%d", &T)
	for i = 0; i < 1; i++ {
		var stream stream_t

		fmt.Scanf("%d%d%d%d%d", &N, &Q, &a, &b, &x)
		stream.N = N
		stream.Q = Q
		stream.a = a
		stream.b = b
		stream.x = x
		usage, query := stream.generator()
	}
}
