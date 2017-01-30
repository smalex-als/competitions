package main

import (
	"fmt"
	"sort"
)

func main2() {
	fmt.Println(boxesPacking([]int{1, 3, 2}, []int{1, 3, 2}, []int{1, 3, 2}))
	fmt.Println(boxesPacking([]int{1, 1}, []int{1, 1}, []int{1, 1}))
}

func boxesPacking(length []int, width []int, height []int) bool {
	n := len(width)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		s := make([]int, 3)
		s[0] = length[i]
		s[1] = width[i]
		s[2] = height[i]
		sort.Ints(s)
		arr[i] = s
	}
	sort.Sort(BySize(arr))
	for i := 1; i < n; i++ {
		a := arr[i-1]
		b := arr[i]
		for j := 0; j < 3; j++ {
			if a[j] >= b[j] {
				return false
			}
		}
	}
	return true
}

type BySize [][]int

func (a BySize) Len() int      { return len(a) }
func (a BySize) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySize) Less(i, j int) bool {
	for k := 0; k < 3; k++ {
		if a[i][k] >= a[j][k] {
			return false
		}
	}
	return true
}
