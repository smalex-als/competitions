package main

import "fmt"

func main7() {
	res := stringsCrossover([]string{"abc", "aaa", "aba", "bab"}, "bbb")
	fmt.Println("res", res)
}

func stringsCrossover(a []string, result string) int {
	res := 0
	for i := 0; i < len(a); i++ {
		if len(a[i]) == len(result) {
			for j := i + 1; j < len(a); j++ {
				if len(a[j]) == len(result) {
					ok := true
					for k := 0; k < len(a[i]); k++ {
						if a[i][k] != result[k] && a[j][k] != result[k] {
							ok = false
							break
						}
					}
					if ok {
						res++
					}
				}
			}
		}
	}
	return res
}
