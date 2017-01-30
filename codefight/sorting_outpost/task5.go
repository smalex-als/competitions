package main

import (
	"fmt"
	"sort"
)

func main5() {
	// a := []int{152, 23, 7, 887, 243}
	// res := digitDifferenceSort(a)
	b := []int{}
	res2 := digitDifferenceSort(b)
	fmt.Println(res2)
}

func digitDifferenceSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	b := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		str := []rune(fmt.Sprintf("%d", a[i]))
		min := '9'
		max := '0'
		for j := 0; j < len(str); j++ {
			if str[j] < min {
				min = str[j]
			}
			if str[j] > max {
				max = str[j]
			}
		}
		diff := int(max - min)
		b[i] = []int{diff*100000 + i, a[i]}
	}
	sort.Sort(BySum(b))
	for i := 0; i < len(a); i++ {
		a[i] = b[i][1]
	}
	return a
}

type BySum [][]int

func (a BySum) Len() int           { return len(a) }
func (a BySum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySum) Less(i, j int) bool { return a[i][0] < a[j][0] }
