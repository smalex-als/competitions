package main

import "fmt"

func main() {
	m := [][]bool{
		{true, false, false},
		{false, true, false},
		{false, false, false},
	}
	fmt.Println(minesweeper(m))
}

func minesweeper(m [][]bool) [][]int {
	h := len(m)
	w := len(m[0])
	res := make([][]int, h)
	for i := 0; i < h; i++ {
		res[i] = make([]int, w)
		for j := 0; j < w; j++ {
			sum := 0
			for ii := -1; ii < 2; ii++ {
				for jj := -1; jj < 2; jj++ {
					ny := ii + i
					nx := jj + j
					if (ii != 0 || jj != 0) &&
						ny >= 0 && ny < h &&
						nx >= 0 && nx < w &&
						m[ny][nx] {
						sum++
					}
				}
			}
			res[i][j] = sum
		}
	}
	return res
}
