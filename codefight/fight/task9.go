package main

import (
	"fmt"
	"sort"
)

func main9() {

	lastBackupTime := 461620205
	changes := [][]int{{461620203, 1},
		{461620204, 2},
		{461620205, 6},
		{461620206, 5},
		{461620207, 3},
		{461620207, 5},
		{461620208, 1}}
	res := incrementalBackups(lastBackupTime, changes)
	fmt.Printf("res = %+v\n", res)
}

func incrementalBackups(lastBackupTime int, changes [][]int) []int {
	update := make(map[int]bool, 0)
	for i := 0; i < len(changes); i++ {
		if changes[i][0] > lastBackupTime {
			update[changes[i][1]] = true
		}
	}
	res := make([]int, 0)
	for k, _ := range update {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}
