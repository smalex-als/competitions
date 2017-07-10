package main

import "fmt"

type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func main() {
	t := &Tree{
		Value: -1,
		Left: &Tree{
			Value: 5,
		},
		Right: &Tree{
			Value: 7,
			Right: &Tree{
				Value: 1,
			},
		},
	}
	values := largestValuesInTreeRows(t)
	fmt.Printf("values = %+v\n", values)
}

func largestValuesInTreeRows(t *Tree) []int {
	res := map[int]int{}
	level := dfs(t, 0, res)
	a := make([]int, level)
	for i := 0; i < level; i++ {
		a[i] = res[i]
	}
	return a
}

func dfs(node *Tree, level int, res map[int]int) int {
	if node == nil {
		return level
	}
	cur := node.Value.(int)
	if prev, ok := res[level]; ok {
		if prev < cur {
			res[level] = cur
		}
	} else {
		res[level] = cur
	}
	return max(dfs(node.Left, level+1, res), dfs(node.Right, level+1, res))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
