package main

import "fmt"

func main02() {
	cities := 4
	roads := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}
	res := roadsBuilding(cities, roads)
	fmt.Printf("res = %+v\n", res)
}

func roadsBuilding(cities int, roads [][]int) [][]int {
	g := make([][]bool, cities)
	for i := 0; i < cities; i++ {
		g[i] = make([]bool, cities)
	}
	for i := 0; i < len(roads); i++ {
		u := roads[i][0]
		v := roads[i][1]
		g[u][v] = true
		g[v][u] = true
	}
	res := make([][]int, 0)
	for u := 0; u < cities; u++ {
		for v := 0; v < cities; v++ {
			if u != v && !g[u][v] {
				g[u][v] = true
				g[v][u] = true
				res = append(res, []int{u, v})
			}
		}
	}
	return res
}
