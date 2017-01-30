package main

import "fmt"

func main9() {
	fmt.Println(beautifulText("Look at this example of a correct text", 5, 15))
	fmt.Println(beautifulText("abc abc abc", 4, 10))
}

func beautifulText(s string, l int, r int) bool {
	for i := l; i <= r; i++ {
		if (len(s)+1)%(i+1) == 0 {
			ok := true
			for j := i; j < len(s); j += i {
				if s[j] != ' ' {
					ok = false
					break
				}
				j++
			}
			if ok {
				return true
			}
		}
	}
	return false
}
