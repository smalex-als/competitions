package main

import "fmt"

func main() {
	res := composeExpression("123", 6)
	fmt.Printf("res = %+v\n", res)
}

func composeExpression(digits string, target int) []string {
	return []string{}
}
