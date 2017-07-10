package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containsCloseNums([]int{0, 1, 2, 3, 5, 2}, 3))
}

func containsCloseNums(nums []int, k int) bool {
	h := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if _, ok := h[v]; !ok {
			h[v] = make([]int, 0)
		}
		h[v] = append(h[v], i)
	}
	for _, arr := range h {
		sort.Ints(arr)
		for i := 1; i < len(arr); i++ {
			if arr[i]-arr[i-1] <= k {
				return true
			}
		}
	}
	return false
}
