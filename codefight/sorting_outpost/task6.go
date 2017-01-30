package main

import "fmt"

func main() {
	a := uniqueDigitProducts([]int{2, 8, 121, 42, 222, 23})
	fmt.Println(a)
}

func uniqueDigitProducts(a []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		str := []rune(fmt.Sprintf("%d", a[i]))
		sum := 1
		for j := 0; j < len(str); j++ {
			sum *= int(str[j] - '0')
		}
		m[sum] = true
	}
	return len(m)
}
