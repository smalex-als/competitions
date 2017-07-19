package main

import "fmt"

func main() {
	a := []int{14, 1, 2, 3, 8, 15, 3}
	fmt.Println(tripletSum(15, a))
}

func tripletSum(x int, a []int) bool {
	inventory := make([]int, 1001)
	for i := 0; i < len(a); i++ {
		inventory[a[i]]++
	}
	for i := 0; i < len(a); i++ {
		inventory[a[i]]--
		for j := i + 1; j < len(a); j++ {
			inventory[a[j]]--
			t := x - a[i] - a[j]
			if t >= 0 && t < len(inventory) && inventory[t] > 0 {
				return true
			}
			inventory[a[j]]++
		}
		inventory[a[i]]++
	}

	return false
}
