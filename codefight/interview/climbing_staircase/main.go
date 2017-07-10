package main

import "fmt"

func main() {
	fmt.Println(climbingStaircase(4, 2))
}

func climbingStaircase(n int, k int) [][]int {
	res := make([][]int, 0)
	a := make([]int, n)
	dfs(a, 0, 0, k, &res)
	return res
}

func dfs(a []int, pos int, sum int, k int, res *[][]int) {
	if sum == len(a) {
		newa := make([]int, pos)
		copy(newa, a[:pos])
		*res = append(*res, newa)
		return
	}
	for i := 1; i <= k; i++ {
		next := sum + i
		a[pos] = i
		if next <= len(a) {
			dfs(a, pos+1, next, k, res)
		} else {
			break
		}
	}
}
