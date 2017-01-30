package main

import (
	"fmt"
	"sort"
)

func main4() {
	m := [][]int{{6, 4}, {2, 2}, {4, 3}}
	res := rowsRearranging(m)
	fmt.Println(res)
}

func rowsRearranging(matrix [][]int) bool {
	sort.Sort(ByRow(matrix))
	fmt.Println(matrix)
	for j := 0; j < len(matrix[0]); j++ {
		for i := 1; i < len(matrix); i++ {
			if matrix[i-1][j] >= matrix[i][j] {
				return false
			}
		}
	}
	return true
}

type ByRow [][]int

func (a ByRow) Len() int      { return len(a) }
func (a ByRow) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRow) Less(i, j int) bool {
	for k := 0; k < len(a[k]); k++ {
		if a[i][k] >= a[j][k] {
			return false
		}
	}
	return true
}
