package main

import "fmt"

func main() {
	fmt.Println(chessTriangle(2, 3))
	fmt.Println(chessTriangle(3, 3))
	fmt.Println(chessTriangle(1, 30))
	fmt.Println(chessTriangle(2, 5))
}

func chessTriangle(h int, w int) int {
	res := 0
	res += loop(h, w, 2, 3, 8)
	res += loop(h, w, 3, 2, 8)
	res += 2 * loop(h, w, 3, 3, 8)
	res += loop(h, w, 4, 2, 8)
	res += loop(h, w, 2, 4, 8)
	res += loop(h, w, 4, 3, 8)
	res += loop(h, w, 3, 4, 8)
	return res
}

func loop(h, w, sizeH, sizeW, score int) int {
	res := 0
	for i := 0; i <= h-sizeH; i++ {
		for j := 0; j <= w-sizeW; j++ {
			res += score
		}
	}
	return res
}

func lmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
