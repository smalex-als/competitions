package main

import "fmt"

func main03() {
	roads := [][]int{
		{3, 0},
		{0, 4},
		{5, 0},
		{2, 1},
		{1, 4},
		{2, 3},
		{5, 2},
	}
	res := efficientRoadNetwork(6, roads)
	fmt.Printf("res = %+v\n", res)
}

func efficientRoadNetwork(n int, roads [][]int) bool {
	g := make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
	}
	for i := 0; i < len(roads); i++ {
		u := roads[i][0]
		v := roads[i][1]
		g[u][v] = true
		g[v][u] = true
	}
	for i := 0; i < n; i++ {
		visited := make([]int, n)
		for j := 0; j < n; j++ {
			visited[j] = 99999
		}
		q := make([]int, 1)
		q[0] = i
		visited[i] = 0
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			l := visited[u]
			if l > 2 {
				return false
			}
			for v := 0; v < n; v++ {
				if g[u][v] && visited[v] == 99999 {
					q = append(q, v)
					visited[v] = l + 1
				}
			}
		}
		for j := 0; j < n; j++ {
			if visited[j] > 2 {
				return false
			}
		}
	}
	return true
}
