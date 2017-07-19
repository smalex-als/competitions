package main

import (
	"fmt"
)

func main() {
	// res := longestIncreasingSubsequence([]int{1, 231, 2, 4, 89, 32, 12, 234, 33, 90, 34, 100})
	res2 := longestIncreasingSubsequence([]int{
		1, 231, 2, 4, 89, 32, 12, 234, 33, 90, 34, 42, 88, 15, 16, 100,
	})
	fmt.Printf("res = %+v\n", res2)
}

func longestIncreasingSubsequence(seq []int) int {
	res := 1
	lis := make([]int, 0)
	for i := 0; i < len(seq); i++ {
		val := seq[i]
		if len(lis) == 0 || lis[len(lis)-1] < val {
			lis = append(lis, val)
			if res < len(lis) {
				res = len(lis)
			}
		} else {
			index := 0
			// you can use binary search here
			// complexity will be O(n * lg(n))
			for j := 0; j < len(lis); j++ {
				if lis[j] >= val {
					index = j
					break
				}
			}
			lis[index] = val
		}
	}
	return res
}
