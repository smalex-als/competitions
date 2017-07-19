package main

import (
	"fmt"
)

func main() {
	fmt.Println(decodeString("abc2[b3[a]]"))
	fmt.Println(decodeString("z1[y]zzz2[abc]"))
}

func decodeString(s string) string {
	num := 0
	stackNum := make([]int, 0)
	stackString := make([]string, 0)
	cur := ""
	for i := 0; i < len(s); i++ {
		ch := byte(s[i])
		if ch >= '0' && ch <= '9' {
			num *= 10
			num += int(ch - '0')
		} else if ch == '[' {
			stackNum = append(stackNum, num)
			stackString = append(stackString, cur)
			cur = ""
			num = 0
		} else if ch == ']' {
			num = stackNum[len(stackNum)-1]
			rep := cur
			stackNum = stackNum[:len(stackNum)-1]
			cur = stackString[len(stackString)-1]
			stackString = stackString[:len(stackString)-1]
			for i := 0; i < num; i++ {
				cur += rep
			}
			num = 0
		} else {
			cur += string(ch)
		}
	}
	return cur
}
