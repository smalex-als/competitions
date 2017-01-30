package main

import (
	"fmt"
	"math"
	"sort"
)

func main3() {
	pkg := []int{4, 2, 5}
	boxes := [][]int{{4, 3, 5}, {5, 2, 5}}
	res := packageBoxing(pkg, boxes)
	fmt.Println(res)
}

func packageBoxing(pkg []int, boxes [][]int) int {
	sort.Ints(pkg)
	best := -1
	bestSize := math.MaxInt32
	for k, v := range boxes {
		sort.Ints(v)
		ok := true
		for i := 0; i < 3; i++ {
			if pkg[i] > v[i] {
				ok = false
				break
			}
		}
		if ok {
			size := v[0] * v[1] * v[2]
			if bestSize > size {
				best = k
				bestSize = size
			}
		}
	}
	return best
}
