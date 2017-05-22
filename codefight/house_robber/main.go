package main

import "fmt"

func main() {
	fmt.Printf("res = %+v\n", houseRobber([]int{1, 1, 1}))
}

func houseRobber(a []int) int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i > 1 {
			res[i] = max(res[i-2]+a[i], res[i-1])
		} else if i > 0 {
			res[i] = max(a[i], res[i-1])
		} else {
			res[i] = a[i]
		}
	}
	return res[len(a)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
