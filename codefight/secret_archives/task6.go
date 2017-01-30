package main

import "fmt"
import "strings"

func main6() {
	fmt.Println(countElements(`1234567890`))
	fmt.Println(countElements(`[[1, 2, [3]], 4]`))
	fmt.Println(countElements(`[ "one", true, "three" ]`))
}

func countElements(str string) int {
	res := 0
	for len(str) > 0 {
		if str[0] == ',' || str[0] == '[' || str[0] == ']' || str[0] == ' ' {
			str = str[1:]
		} else if str[0] == '"' {
			str = str[1:]
			index := strings.Index(str, "\"")
			if index >= 0 {
				str = str[index+1:]
			}
			res++
		} else if str[0] >= '0' && str[0] <= '9' {
			for len(str) > 0 && str[0] >= '0' && str[0] <= '9' {
				str = str[1:]
			}
			res++
		} else if strings.HasPrefix(str, "false") {
			res++
			str = str[5:]
		} else if strings.HasPrefix(str, "true") {
			str = str[4:]
			res++
		}
	}
	return res
}
