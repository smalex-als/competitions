package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := treeLevelSum("(0(5(6()())(14()(9()())))(7(1()())(23()())))", 2)
	fmt.Printf("res = %+v\n", res)
}

var mytree string
var pos int
var K int

func treeLevelSum(tree string, k int) int {
	mytree = tree
	pos = 0
	K = k
	return readNode(0)
}

func readNode(level int) int {
	if mytree[pos] != '(' {
		panic("expected '('")
	}
	pos++
	if mytree[pos] == ')' {
		pos++
		return 0
	}
	start := pos
	res := 0
	for pos < len(mytree) {
		if mytree[pos] >= '0' && mytree[pos] <= '9' {
			pos++
		} else {
			if level == K {
				num, _ := strconv.Atoi(mytree[start:pos])
				res += num
			}
			break
		}
	}
	res += readNode(level + 1) // left
	res += readNode(level + 1) // right
	if mytree[pos] != ')' {
		panic("expected ')'")
	}
	pos++
	return res
}
