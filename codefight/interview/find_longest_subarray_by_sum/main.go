package main

import "fmt"

func main() {
	fmt.Println(findLongestSubarrayBySum(12, []int{1, 2, 3, 7, 5}))
}

func findLongestSubarrayBySum(s int, arr []int) []int {
	sum := int64(s)
	localSum := int64(0)
	start := 0
	bestLen := 0
	ret := []int{-1}
	for i := 0; i < len(arr); i++ {
		localSum += int64(arr[i])
		for localSum > sum {
			localSum -= int64(arr[start])
			start++
		}
		if localSum == sum && bestLen < (i-start+1) {
			bestLen = i - start + 1
			ret = []int{start + 1, i + 1}
		}
	}
	return ret
}
