package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	res := treeBottom("(2 (7 (2 () ()) (6 (5 () ()) (11 () ()))) (5 () (9 (4 () ()) ())))")
	fmt.Println(res)
}

func treeBottom(tree string) []int {
	res2 := make([][]int, 0)
	res2, _ = parseNode(tree, 0, res2)
	return res2[len(res2)-1]
}

func parseNode(tree string, level int, res2 [][]int) ([][]int, string) {
	if len(tree) == 0 {
		return res2, tree
	}
	if tree[0] == '(' {
		tree = tree[1:]
	}
	if tree[0] == ')' {
		tree = tree[1:]
		return res2, tree
	}
	r := regexp.MustCompile("(\\d+) ")

	m := r.FindStringSubmatch(tree)
	if len(m) > 0 {
		num, _ := strconv.Atoi(m[1])
		if level == len(res2) {
			res2 = append(res2, make([]int, 0))
		}
		res2[level] = append(res2[level], num)
		tree = tree[len(m[0]):]
		res2, tree = parseNode(tree, level+1, res2)
		if tree[0] == ' ' {
			tree = tree[1:]
		}
		res2, tree = parseNode(tree, level+1, res2)
	}

	if tree[0] == ')' {
		tree = tree[1:]
		return res2, tree
	}
	return res2, tree
}
