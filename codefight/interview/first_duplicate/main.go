package main

import "fmt"

func main() {
	a := []int{8, 1, 4, 8, 10, 1, 5, 7, 8, 7}
	//	a := []int{8, 4, 6, 2, 6, 4, 7, 9, 5, 8}
	// a := []int{2, 3, 3, 1, 5, 2}
	// a := []int{2, 3, 3}
	fmt.Println(firstDuplicate(a))
}

func firstDuplicate(a []int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == i+1 {
			a[i] = -a[i]
		} else if a[i] == -(i + 1) {
			return i + 1
		} else {
			if a[abs(a[i])-1] < 0 {
				return abs(a[i])
			}
			a[abs(a[i])-1] = -a[abs(a[i])-1]
		}
	}
	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
