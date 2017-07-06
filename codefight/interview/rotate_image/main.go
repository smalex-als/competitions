package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := rotateImage(a)
	for i := 0; i < len(res); i++ {
		fmt.Printf("a[i] = %+v\n", a[i])
	}
}

func rotateImage(a [][]int) [][]int {
	n := len(a)
	for j := 0; j < n; j++ {
		for i := 0; i < n/2; i++ {
			a[j][i], a[j][n-i-1] = a[j][n-i-1], a[j][i]
		}
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			a[j][i], a[n-i-1][n-j-1] = a[n-i-1][n-j-1], a[j][i]
		}
	}
	return a
}
