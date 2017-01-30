package main

import (
	"fmt"
	"regexp"
)

func main5() {
	fmt.Println(isSubsequence("CodeFights", "CoFi"))
}

func isSubsequence(t string, s string) bool {
	pattern := ""
	for _, ch := range s {
		pattern += fmt.Sprintf(".*%s.*", regexp.QuoteMeta(string(ch)))
	}
	fmt.Println(pattern)
	re := regexp.MustCompile(pattern)
	return re.MatchString(t)
}
