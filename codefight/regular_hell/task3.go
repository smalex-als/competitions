package main

import (
	"fmt"
	"regexp"
)

func main3() {
	fmt.Println(swapAdjacentWords("How are you today guys"))
}

func swapAdjacentWords(s string) string {
	return regexp.MustCompile("([A-Za-z]+) ([A-Za-z]+)").ReplaceAllString(s, "$2 $1")
}
