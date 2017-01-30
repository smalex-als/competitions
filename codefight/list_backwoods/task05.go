package main

import "fmt"

func main2() {
	matrix := [][]int{
		{1, 0, 0, 2, 0, 0, 3},
		{0, 1, 0, 2, 0, 3, 0},
		{0, 0, 1, 2, 3, 0, 0},
		{8, 8, 8, 9, 4, 4, 4},
		{0, 0, 7, 6, 5, 0, 0}}
	res := starRotation(matrix, 3, []int{1, 5}, 81)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(res[i])
	}
}

func starRotation(matrix [][]int, w int, center []int, t int) [][]int {
	y := center[0]
	x := center[1]
	yx := [][]int{
		{0, 1, y, x - w/2},
		{1, 1, y - w/2, x - w/2},
		{1, 0, y - w/2, x},
		{1, -1, y - w/2, x + w/2},
	}
	allyx := [][]int{
		{0, 1, y, x - w/2},
		{1, 1, y - w/2, x - w/2},
		{1, 0, y - w/2, x},
		{1, -1, y - w/2, x + w/2},
		{0, -1, y, x + w/2},
		{-1, -1, y + w/2, x + w/2},
		{-1, 0, y + w/2, x},
		{-1, 1, y + w/2, x - w/2},
	}
	res := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, len(matrix[i]))
		copy(res[i], matrix[i])
	}
	for j, k := range yx {
		fmt.Println(j, k)
		arr := make([]int, w)
		for curY, curX, i := k[2], k[3], 0; i < w; i++ {
			arr[i] = matrix[curY][curX]
			curY += k[0]
			curX += k[1]
		}
		nk := allyx[(j+t)%8]
		for curY, curX, i := nk[2], nk[3], 0; i < w; i++ {
			res[curY][curX] = arr[i]
			curY += nk[0]
			curX += nk[1]
		}
	}
	return res
}
