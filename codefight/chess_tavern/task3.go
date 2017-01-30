package main

import (
	"fmt"
	"sort"
)

func main3() {
	fmt.Println(bishopDiagonal("a1", "h8"))
	fmt.Println(bishopDiagonal("d7", "f5"))
	fmt.Println(bishopDiagonal("d8", "b5"))
}

func bishopDiagonal(bishop1 string, bishop2 string) []string {
	y0, x0 := parseCoords3(bishop1)
	y1, x1 := parseCoords3(bishop2)
	dx := x1 - x0
	dy := y1 - y0
	res := []string{bishop1, bishop2}
	defer sort.Strings(res)
	if Abs3(dx) != Abs3(dy) {
		return res
	}
	dx /= Abs3(dx)
	dy /= Abs3(dy)
	dx, dy = -dx, -dy
	x0, y0 = moveFigure(x0, y0, dx, dy)
	dx, dy = -dx, -dy
	x1, y1 = moveFigure(x1, y1, dx, dy)
	res[0] = coordsToString(x0, y0)
	res[1] = coordsToString(x1, y1)
	return res
}

func moveFigure(x0, y0, dx, dy int) (int, int) {
	for {
		ny := y0 + dy
		nx := x0 + dx
		if ny < 0 || ny > 7 || nx < 0 || nx > 7 {
			return x0, y0
		}
		y0 = ny
		x0 = nx
	}
}

func Abs3(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseCoords3(coords string) (int, int) {
	x := int(coords[0] - 'a')
	y := int(coords[1] - '1')
	return y, x
}

func coordsToString(x, y int) string {
	return fmt.Sprintf("%c%d", x+'a', y+1)
}
