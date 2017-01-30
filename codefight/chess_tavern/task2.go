package main

import "fmt"

func main2() {
	fmt.Println(chessKnightMoves("c2"))
}

func chessKnightMoves(cell string) int {
	y, x := parseCoords(cell)
	res := 0
	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			if Abs(i)+Abs(j) == 3 {
				if y+i >= 0 && y+i <= 7 &&
					x+j >= 0 && x+j <= 7 {
					res++
				}
			}
		}
	}
	return res
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseCoords(coords string) (int, int) {
	return int(coords[0] - 'a'), int(coords[1] - '1')
}
