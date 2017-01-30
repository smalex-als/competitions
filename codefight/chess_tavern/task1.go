package main

import "fmt"

func main1() {
	fmt.Println(bishopAndPawn("a1", "c3"))
}

func bishopAndPawn(bishop string, pawn string) bool {
	y0, x0 := parse(bishop)
	y1, x1 := parse(pawn)
	y, x := y0-y1, x0-x1
	if y < 0 {
		y *= -1
	}
	if x < 0 {
		x *= -1
	}
	return x == y
}

func parse(coords string) (int, int) {
	return int(coords[0] - 'a'), int(coords[1] - '1')
}
