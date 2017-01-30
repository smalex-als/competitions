package main

import "fmt"
import "math"

func main() {
	// fmt.Println(rectangleRotation(4, 6))
	fmt.Println(rectangleRotation(2, 30))
}

func rectangleRotation(a int, b int) int {
	cnt := 0
	h := math.Hypot(float64(1), float64(1))
	maxA := float64(a) / 2
	minA := -float64(a) / 2
	maxB := float64(b) / 2
	minB := -float64(b) / 2
	for j := -a; j < a; j++ {
		for i := -b; i < b; i++ {
			x := h * float64(i)
			y := h * float64(j)
			if x >= minB && x <= maxB && y >= minA && y <= maxA {
				cnt++
			}
			x = h*float64(i) + h/2
			y = h*float64(j) + h/2
			if x >= minB && x <= maxB && y >= minA && y <= maxA {
				cnt++
			}
		}
	}
	return cnt
}
