package main

import "fmt"
import "math"

func main6() {
	fmt.Println(countBlackCells(3, 4))
	// fmt.Println(countBlackCells(2, 5))
	// fmt.Println(countBlackCells(1, 2))
	// fmt.Println(countBlackCells(1, 3))
	// fmt.Println(countBlackCells(1, 293))
	// fmt.Println(countBlackCells(33, 44))
	// fmt.Println(countBlackCells(16, 8))
	// fmt.Println(countBlackCells(4, 2))
}

func countBlackCells(dx, dy int) int {
	if dx < dy {
		dx, dy = dy, dx
	}

	if dx == dy {
		return dx + 2*(dx-1)
	}
	m := make(map[string]bool)
	for x := 1; x < dx; x++ {
		y := float64(dy) * float64(x) / float64(dx)
		if math.Floor(y) == y {
			m[fmt.Sprintf("%d, %d", int(y), x)] = true
			m[fmt.Sprintf("%d, %d", int(y)+1, x)] = true
		} else {
			m[fmt.Sprintf("%d, %d", int(y), x)] = true
			m[fmt.Sprintf("%d, %d", int(y), x+1)] = true
		}
	}
	return len(m)
}
