package main

import (
	"fmt"
	"regexp"
	"strings"
)

func repetitionEncryption(letter string) int {
	pattern := regexp.MustCompile("[^a-z]+")
	words := pattern.Split(strings.ToLower(letter), -1)
	res := 0
	for i := range words {
		if i > 0 && words[i] == words[i-1] {
			res++
		}
	}
	return res
}

func main7() {
	fmt.Println(repetitionEncryption("Hi, hi Jane! I'm so. So glad to to finally be able to write - WRITE!! - to you!"))
}
