package main

import (
	"fmt"
	"regexp"
)

func main1() {
	fmt.Println(isSentenceCorrect("This is an example of *correct* sentence."))
}

func isSentenceCorrect(sentence string) bool {
	re := regexp.MustCompile("^[A-Z][^.?!]*[.?!]$")
	return re.MatchString(sentence)
}
