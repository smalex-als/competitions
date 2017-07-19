package main

import (
	"fmt"
	"sort"
)

func main() {
	res := sumSubsets([]int{1, 2, 2, 3, 4, 5}, 5)
	fmt.Printf("res = %+v\n", res)
}

var res [][]int
var visited map[string]bool

func sumSubsets(arr []int, num int) [][]int {
	res = make([][]int, 0)
	visited = make(map[string]bool, 0)
	sort.Ints(arr)
	selected := make([]bool, len(arr))
	dfs(arr, num, 0, 0, selected)
	return res
}

func dfs(arr []int, num int, pos int, sum int, selected []bool) {
	if num < sum {
		return
	}
	if num == sum {
		a := make([]int, 0)
		for i := 0; i < len(selected); i++ {
			if selected[i] {
				a = append(a, arr[i])
			}
		}
		strKey := fmt.Sprintln(a)
		if _, exists := visited[strKey]; !exists {
			visited[strKey] = true
			res = append(res, a)
		}
		return
	}
	if pos == len(arr) {
		return
	}
	selected[pos] = true
	dfs(arr, num, pos+1, sum+arr[pos], selected)
	selected[pos] = false
	dfs(arr, num, pos+1, sum, selected)
}
