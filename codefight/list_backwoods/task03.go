package main

import "fmt"

func test22() {
	field := [][]byte{
		{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
		{'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		{'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
	}
	rect := []int{1, 1, 4, 3}

	res := drawRectangle(field, rect)
	for i := 0; i < len(res); i++ {
		fmt.Println(string(res[i]))
	}
}

func drawRectangle(canvas [][]byte, r []int) [][]byte {
	x0, y0, x1, y1 := r[0], r[1], r[2], r[3]
	canvas[y0][x0] = '*'
	canvas[y0][x1] = '*'
	canvas[y1][x0] = '*'
	canvas[y1][x1] = '*'
	for i := x0 + 1; i < x1; i++ {
		canvas[y0][i] = '-'
		canvas[y1][i] = '-'
	}
	for i := y0 + 1; i < y1; i++ {
		canvas[i][x0] = '|'
		canvas[i][x1] = '|'
	}
	return canvas
}
