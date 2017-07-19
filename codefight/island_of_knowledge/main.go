package main

import "fmt"

func main() {
	a := [][]int{
		{1, 1, 1},
		{1, 7, 1},
		{1, 1, 1},
	}
	fmt.Println(boxBlur(a))
}

func boxBlur(image [][]int) [][]int {
	res := make([][]int, 0)
	for j := 0; j < len(image)-2; j++ {
		cur := make([]int, 0)
		for i := 0; i < len(image[0])-2; i++ {
			sum := 0
			for jj := 0; jj < 3; jj++ {
				for ii := 0; ii < 3; ii++ {
					sum += image[j+jj][i+ii]
				}
			}
			cur = append(cur, sum/9)
		}
		res = append(res, cur)
	}
	return res
}
