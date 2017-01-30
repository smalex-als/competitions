package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main6() {
	code := `
	function add($n, m,n,n) {\t
	  return n + n0 + $m;\t
  }`
	fmt.Println(programTranslation(code, []string{"n", "m"}))
}

func programTranslation(solution string, args []string) string {
	argumentVariants := strings.Join(args, "|")
	re := regexp.MustCompile("[$]?\\b(" + argumentVariants + ")\\b")
	repl := "$$$1"
	subSolution := re.ReplaceAllString(solution, "$1")
	return re.ReplaceAllString(subSolution, repl)
}
