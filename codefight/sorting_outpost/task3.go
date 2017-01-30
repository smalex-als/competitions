package main

import (
	"fmt"
	"sort"
)

func main3() {
	A := []int{4, 2, 1, 6, 5, 7, 2, 4}
	Q := [][]int{{1, 6}, {2, 4}, {3, 6}, {0, 7}, {3, 6}, {4, 4}, {5, 6}, {5, 6}, {0, 1}, {3, 4}}
	// A := []int{2, 1, 2}
	// Q := [][]int{{0, 1}}
	res := maximumSum(A, Q)
	fmt.Println(res)
}

func maximumSum(A []int, Q [][]int) int {
	cnt := make([]int, len(A))
	for i := 0; i < len(Q); i++ {
		for j := Q[i][0]; j <= Q[i][1]; j++ {
			cnt[j]++
		}
	}
	sort.Ints(cnt)
	sort.Ints(A)
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += cnt[i] * A[i]
	}
	return sum
}
