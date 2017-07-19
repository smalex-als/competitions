package main

import "fmt"

func main() {

	roads := [][]int{
		{1, 0}, {1, 2}, {8, 5}, {9, 7},
		{5, 6}, {5, 4}, {4, 6}, {6, 7},
	}
	res := citiesConquering(10, roads)
	fmt.Printf("res = %+v\n", res)
}

func citiesConquering(n int, roads [][]int) []int {
	g := make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
	}
	for _, road := range roads {
		u := road[0]
		v := road[1]
		g[u][v] = true
		g[v][u] = true
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	days := 1
	updated := true
	for updated {
		updated = false
		for i := 0; i < n; i++ {
			if res[i] == -1 {
				cnt := 0
				for j := 0; j < n; j++ {
					if g[i][j] {
						cnt++
					}
				}
				if cnt <= 1 {
					res[i] = days
					updated = true
				}
			}
		}
		for i := 0; i < n; i++ {
			if res[i] == days {
				for j := 0; j < n; j++ {
					if g[i][j] {
						g[i][j] = false
						g[j][i] = false
					}
				}
			}
		}
		days++
	}
	return res
}
