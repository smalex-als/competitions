package main

import (
	"fmt"
	"math"
)

func main25() {
	network := [][]int{
		{0, 2, 4, 8, 0, 0},
		{0, 0, 0, 9, 0, 0},
		{0, 0, 0, 0, 0, 10},
		{0, 0, 6, 0, 0, 10},
		{10, 10, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	res := dataRouteShort(4, 5, network)
	fmt.Printf("res = %+v\n", res)
}

// tags: graph max flow
func dataRouteShort(from int, to int, capacity [][]int) int {
	N := len(capacity)
	flow := make([][]int, N)
	for i := 0; i < N; i++ {
		flow[i] = make([]int, N)
	}
	path := make([]int, N)
	value := 0
	for {
		for i := 0; i < N; i++ {
			path[i] = -1
		}
		q := make([]int, 1)
		q[0] = from
		for len(q) > 0 && path[to] == -1 {
			cur := q[0]
			q = q[1:]
			for i := 0; i < N; i++ {
				if capacity[cur][i]-flow[cur][i] > 0 && path[i] == -1 {
					path[i] = cur
					q = append(q, i)
				}
			}
		}
		if path[to] == -1 {
			break
		}
		bottle := math.MaxInt32
		for i := to; i != from; i = path[i] {
			dist := capacity[path[i]][i] - flow[path[i]][i]
			if dist < bottle {
				bottle = dist
			}
		}
		for i := to; i != from; i = path[i] {
			flow[path[i]][i] += bottle
		}
		value += bottle
	}
	return value
}
