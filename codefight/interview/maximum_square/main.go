package main

import "fmt"

func main() {
	m := [][]string{
		{"1", "1", "1", "1", "1", "1", "1", "1"},
		{"1", "1", "1", "1", "1", "1", "1", "0"},
		{"1", "1", "1", "1", "1", "1", "1", "0"},
		{"1", "1", "1", "1", "1", "0", "0", "0"},
		{"0", "1", "1", "1", "1", "0", "0", "0"},
	}
	res := maximalSquare(m)
	fmt.Printf("res = %+v\n", res)
}

func maximalSquare(matrix [][]string) int {
	H := len(matrix)
	if H == 0 {
		return 0
	}
	W := len(matrix[0])
	r := make([][]int, H)
	for y := 0; y < H; y++ {
		r[y] = make([]int, W)
		for x := W - 1; x >= 0; x-- {
			if matrix[y][x] == "1" {
				r[y][x] = 1
				if x+1 < W {
					r[y][x] += r[y][x+1]
				}
			}
		}
		fmt.Printf("r[y] = %+v\n", r[y])
	}
	d := make([][]int, H)
	for y := 0; y < H; y++ {
		d[y] = make([]int, W)
	}
	for x := 0; x < W; x++ {
		for y := H - 1; y >= 0; y-- {
			if matrix[y][x] == "1" {
				d[y][x] = 1
				if y+1 < H {
					d[y][x] += d[y+1][x]
				}
			}
		}
	}

	res := 0
	for x := 0; x < W; x++ {
		for y := 0; y < H; y++ {
			cnt := min(d[y][x], r[y][x])
			for i := y + 1; i < y+cnt; i++ {
				cnt = min(cnt, r[i][x])
			}
			for i := x + 1; i < x+cnt; i++ {
				cnt = min(cnt, d[y][i])
			}
			if cnt > res {
				res = cnt
			}
		}
	}
	return res * res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
