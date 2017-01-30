package main

import (
	"fmt"
	"sort"
)

func main10() {
	res := runnersMeetings([]int{1, 4, 2}, []int{27, 18, 24})
	fmt.Println(res)
}

func runnersMeetings(startPosition []int, speed []int) int {
	max := -1
	for j := 0; j < 1000; j++ {
		a := make([]int, len(speed))
		for i := 0; i < len(speed); i++ {
			a[i] = startPosition[i]*60 + speed[i]*j
		}
		sort.Ints(a)
		fmt.Println(j, a)
		cnt := 1
		for i := 1; i < len(a); i++ {
			if a[i-1] == a[i] {
				cnt++
				if cnt > 1 && cnt >= max {
					max = cnt
					if max == len(speed) {
						return max
					}
				}
			} else {
				cnt = 1
			}
		}
	}

	return max
}
