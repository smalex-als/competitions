package main

import "fmt"

func main() {
	fmt.Printf("res = %+v\n", arrayMaxConsecutiveSum2([]int{1, -2, 3, -4, 5, -3, 2, 2, 2}))
}

func arrayMaxConsecutiveSum2(a []int) int {
	best := -10000
	sum := 0
	for _, v := range a {
		if best < v {
			best = v
		}
		sum += v
		if sum < 0 {
			sum = 0
		} else if sum > best {
			best = sum
		}
	}
	return best
}
