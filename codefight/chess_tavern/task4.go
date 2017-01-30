package main

import (
	"fmt"
	"math"
	"strings"
)

func main4() {
	fmt.Println(whoseTurn("g1;g2;g3;g4"))
	// fmt.Println(whoseTurn("c3;g1;b8;g8"))
	// fmt.Println(whoseTurn("b1;g1;b8;g8"))
}

func parsePosition(c string) (x, y int) {
	return int(c[0] - 'a'), int(c[1] - '1')
}

func whoseTurn(p string) bool {
	args := strings.Split(p, ";")
	steps := make([]int, 0)
	for i := 0; i < 4; i++ {
		field := make([][]int, 8)
		for i := 0; i < 8; i++ {
			field[i] = make([]int, 8)
			for j := 0; j < 8; j++ {
				field[i][j] = math.MaxInt16
			}
		}
		x, y := parsePosition(args[i])
		dfs(x, y, field, 0)
		if i < 2 {
			steps = append(steps, field[0][1])
			steps = append(steps, field[0][6])
		} else {
			steps = append(steps, field[7][1])
			steps = append(steps, field[7][6])
		}
	}
	res := [][]int{
		{steps[0] + steps[3], steps[1] + steps[2]},
		{steps[4] + steps[7], steps[5] + steps[6]},
	}
	for i := 0; i < 2; i++ {
		diff := res[0][i] - res[1][i]
		if diff == 0 || diff == 1 {
			return diff == 0
		}
	}
	return false
}

func dfs(x, y int, field [][]int, turn int) {
	if field[y][x] < turn {
		return
	}
	field[y][x] = turn
	for i := -2; i <= 2; i++ {
		for j := -2; j <= 2; j++ {
			if abs(i)+abs(j) == 3 {
				nx := i + x
				ny := j + y
				if ny >= 0 && ny < 8 && nx >= 0 && nx < 8 {
					dfs(nx, ny, field, turn+1)
				}
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
