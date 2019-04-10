package main

import "fmt"
import "os"

//import "time"

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

func merge(left, right []usage_t) []usage_t {
	var ri, li, i uint64
	result := make([]usage_t, len(right)+len(left))
	for {
		if ri == uint64(len(right)) && li == uint64(len(left)) {
			break
		}
		if ri == uint64(len(right)) {
			for ; li != uint64(len(left)); li++ {
				result[i] = left[li]
				i++
			}
		} else if li == uint64(len(left)) {
			for ; ri != uint64(len(right)); ri++ {
				result[i] = right[ri]
				i++
			}
		} else {
			if left[li].pow > right[ri].pow {
				result[i] = right[ri]
				ri++
				i++
			} else {
				result[i] = left[li]
				li++
				i++
			}
		}
	}
	return result
}

func sort(usage []usage_t) []usage_t {
	if len(usage) == 1 {
		return usage
	} else {
		left := sort(usage[:len(usage)/2])
		right := sort(usage[len(usage)/2:])
		result := merge(left, right)
		return result
	}
}

func (stream stream_t) next(x uint64) uint64 {
	return (stream.a*x + stream.b) % (1E9 + 7)
}

func (stream stream_t) generator() ([]usage_t, []query_t) {
	var i uint64
	x := stream.x
	usage := make([]usage_t, stream.N)
	query := make([]query_t, stream.Q)

	for i = 0; i < stream.N; i++ { // generování spotřeby
		x = stream.next(x)
		usage[i].pow = x
		if usage[i].pow > stream.N-1 {
			usage[i].pow %= stream.N
		}
		usage[i].index = i
	}
	for i = 0; i < stream.Q; i++ { // generování dotazů
		x = stream.next(x) // B
		query[i].B = x
		if query[i].B > stream.N-1 {
			query[i].B %= stream.N
		}
		x = stream.next(x) // E
		query[i].E = x
		if query[i].E > stream.N-1 {
			query[i].E %= stream.N
		}
		if query[i].E < query[i].B {
			swp := query[i].E
			query[i].E = query[i].B
			query[i].B = swp
		}
		x = stream.next(x) // K
		if x > query[i].E-query[i].B {
			query[i].K = x % (query[i].E - query[i].B + 1)
		} else {
			query[i].K = x
		}
	}
	return usage, query
}

func solve(usage []usage_t, query []query_t) uint64 {
	var XOR, index uint64

	for _, ques := range query {
		index = 0
		for _, pow := range usage {
			if pow.index < ques.B || pow.index > ques.E {
				continue
			} else {
				if index == ques.K {
					XOR ^= pow.pow
					break
				} else {
					index++
				}
			}
		}
	}
	return XOR
}

func compute(stream stream_t, out chan uint64, i uint8) {
	usage, query := stream.generator()
	usage = sort(usage)
	fmt.Fprintf(os.Stderr, "sorted: %d\n", i)
	out <- solve(usage, query)
	fmt.Fprintf(os.Stderr, "Solved: %d\n", i)
}

func main() {
	var T uint8
	var N, Q, a, b, x uint64
	var i uint8
	var task []chan uint64

	fmt.Scanf("%d", &T)
	for i = 0; i < T; i++ {
		task = append(task, make(chan uint64, 1))
	}
	for i = 0; i < T; i++ {
		var stream stream_t

		fmt.Scanf("%d%d%d%d%d", &N, &Q, &a, &b, &x)
		stream.N = N
		stream.Q = Q
		stream.a = a
		stream.b = b
		stream.x = x
		//fmt.Fprintf(os.Stderr, "Compute: %d\n", i)
		go compute(stream, task[i], i)
	}
	for {
		for i = 0; i < T; i++ {
			select {
			case res := <-task[i]:
				fmt.Printf("%d:%d\n", i, res)
			default:
				continue
			}
		}

	}
}
