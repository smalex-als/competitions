package main

import (
	"fmt"
	"strconv"
)

func main7() {
	userId := "1851027803204441"
	res := kikCode(userId)
	fmt.Println(res)
}

func kikCode(userId string) [][][]int {
	id, _ := strconv.ParseInt(userId, 10, 64)
	a := make([]int, 0)
	for _, ch := range fmt.Sprintf("%052b", id) {
		a = append(a, int(ch-'0'))
	}
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	parts := []int{3, 4, 8, 10, 12, 15}
	offset := 0
	res := make([][][]int, 0)
	for i := 0; i < len(parts); i++ {
		next := offset + parts[i]
		cur := a[offset:next]
		size := 360 / len(cur)
		start := -1
		ranges := make([][]int, 0)
		for i := 0; i < len(cur); i++ {
			if cur[i] == 1 {
				if start == -1 {
					start = i
				}
			}
			if start >= 0 && ((i+1 == len(cur)) || (i+1 < len(cur) && cur[i+1] == 0)) {
				ranges = append(ranges, []int{start * size, (1 + i) * size})
				start = -1
			}
		}
		if len(ranges) > 1 && ranges[0][0] == 0 && ranges[len(ranges)-1][1] == 360 {
			ranges[len(ranges)-1][1] += ranges[0][1]
			ranges = ranges[1:]
		}
		for _, v := range ranges {
			item := make([][]int, 2)
			item[0] = []int{i + 1, v[0]}
			item[1] = []int{i + 1, v[1]}
			res = append(res, item)
		}
		offset = next
	}
	return res
}
