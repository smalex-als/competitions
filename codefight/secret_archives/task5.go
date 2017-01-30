package main

import "fmt"

func main5() {
	fmt.Println(firstOperationCharacter("((2 + 2) * 2) * 3 + (2 + (2 * 2))"))
}

func firstOperationCharacter(expr string) int {
	buf := []rune(expr)
	level := 0
	maxcost := 0
	maxpos := -1
	for i := 0; i < len(buf); i++ {
		ch := buf[i]
		if ch == '(' {
			level++
		} else if ch == ')' {
			level--
		} else if ch == '*' || ch == '+' {
			cost := level*3 + 1
			if ch == '*' {
				cost++
			}
			if cost > maxcost {
				maxcost = cost
				maxpos = i
			}
		}
	}
	return maxpos
}
