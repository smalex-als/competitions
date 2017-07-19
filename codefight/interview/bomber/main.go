package main

import "fmt"

func main() {
	field := [][]string{
		{"0", "0", "E", "0"},
		{"W", "0", "W", "E"},
		{"0", "E", "0", "W"},
		{"0", "W", "0", "E"},
	}
	res := bomber(field)
	fmt.Printf("res = %+v\n", res)
}

func bomber(field [][]string) int {
	W := len(field[0])
	H := len(field)
	mm := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		mm[i] = make([][]int, H)
	}
	for i := 0; i < H; i++ {
		mm[0][i] = make([]int, W)
		mm[1][i] = make([]int, W)
		mm[2][i] = make([]int, W)
		mm[3][i] = make([]int, W)
		for j := 0; j < W; j++ {
			prev := 0
			if j > 0 {
				prev = mm[0][i][j-1]
			}
			switch field[i][j] {
			case "W":
				prev = 0
			case "E":
				prev++
			}
			mm[0][i][j] = prev
		}
		for j := W - 1; j >= 0; j-- {
			prev := 0
			if j < W-1 {
				prev = mm[1][i][j+1]
			}
			switch field[i][j] {
			case "W":
				prev = 0
			case "E":
				prev++
			}
			mm[1][i][j] = prev
		}
	}
	for j := 0; j < W; j++ {
		for i := 0; i < H; i++ {
			prev := 0
			if i > 0 {
				prev = mm[2][i-1][j]
			}
			switch field[i][j] {
			case "W":
				prev = 0
			case "E":
				prev++
			}
			mm[2][i][j] = prev
		}
		for i := H - 1; i >= 0; i-- {
			prev := 0
			if i < H-1 {
				prev = mm[3][i+1][j]
			}
			switch field[i][j] {
			case "W":
				prev = 0
			case "E":
				prev++
			}
			mm[3][i][j] = prev
		}
	}

	res := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if field[i][j] != "0" {
				continue
			}
			val := 0
			for k := 0; k < 4; k++ {
				val += mm[k][i][j]
			}
			if val > res {
				//	fmt.Printf("i = %+v j = %+v\n", i, j)
				res = val
			}
		}
	}

	return res
}
