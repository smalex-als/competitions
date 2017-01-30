package main

import "fmt"

func main() {
	roads := [][]int{
		{0, 1, 0},
		{4, 1, 2},
		{4, 3, 4},
		{2, 3, 1},
		{2, 0, 3},
	}
	// roads := [][]int{
	// 	{2, 3, 1},
	// 	{3, 0, 0},
	// 	{2, 0, 2},
	// }
	// roads := [][]int{
	// 	{0, 6, 0},
	// 	{0, 5, 5},
	// 	{0, 4, 8},
	// 	{4, 2, 7},
	// 	{3, 2, 3},
	// 	{1, 0, 1},
	// 	{1, 5, 6},
	// 	{6, 2, 2},
	// 	{5, 4, 4},
	// }
	res := namingRoads(roads)
	fmt.Printf("res = %+v\n", res)
}

func namingRoads(roads [][]int) bool {
	V := 0
	for i := range roads {
		if roads[i][0] > V {
			V = roads[i][0]
		}
		if roads[i][1] > V {
			V = roads[i][1]
		}
	}
	V++
	g := make([][][]int, V)
	for i := range g {
		g[i] = make([][]int, 0)
	}
	for i := range roads {
		u := roads[i][0]
		v := roads[i][1]
		w := roads[i][2]
		g[u] = append(g[u], []int{v, w})
		g[v] = append(g[v], []int{u, w})
	}
	for i := range g {
		for j := 0; j < len(g[i]); j++ {
			for k := j + 1; k < len(g[i]); k++ {
				if abs(g[i][j][1]-g[i][k][1]) <= 1 {
					return false
				}
			}
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
