package main

import "fmt"

type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

func main() {
	t := &Tree{
		Value: 1,
		Left: &Tree{
			Value: 2,
			Right: &Tree{
				Value: 3,
			},
		},
		Right: &Tree{
			Value: 4,
			Left: &Tree{
				Value: 5,
			},
		},
	}
	res := traverseTree(t)
	fmt.Printf("res = %+v\n", res)
}

func traverseTree(t *Tree) []int {
	res := make([]int, 0)
	queue := make([]*Tree, 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if v, ok := (node.Value).(int); ok {
			res = append(res, v)
		}
	}
	return res
}
