package main

import (
	"fmt"
	"strings"
)

func main8() {
	fmt.Println(cyclicString("cabca"))
}

func cyclicString(s string) int {
	for i := 1; i < len(s); i++ {
		t := strings.Repeat(s[:i], 15)
		if strings.Contains(t, s) {
			return i
		}
	}
	return len(s)
}
