package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	grid := [][]string{
		{".", ".", ".", "1", "4", ".", ".", "2", "."},
		{".", ".", "6", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "1", ".", ".", ".", ".", ".", "."},
		{".", "6", "7", ".", ".", ".", ".", ".", "9"},
		{".", ".", ".", ".", ".", ".", "8", "1", "."},
		{".", "3", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", "7", ".", ".", "."},
		{".", ".", ".", "5", ".", ".", ".", "7", "."},
	}
	res := sudoku2(grid)
	fmt.Printf("res = %+v\n", res)
}

func sudoku2(grid [][]string) bool {
	for i := 0; i < len(grid); i += 3 {
		for j := 0; j < len(grid); j += 3 {
			a := make([]int, 10)
			for ii := 0; ii < 3; ii++ {
				for jj := 0; jj < 3; jj++ {
					sym := grid[i+ii][j+jj][0]
					if sym != '.' {
						num := sym - '0'
						if a[num] > 0 {
							return false
						}
						a[num]++
					}
				}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		a := make([]int, 10)
		for j := 0; j < len(grid); j++ {
			sym := grid[i][j][0]
			if sym != '.' {
				num := sym - '0'
				if a[num] > 0 {
					return false
				}
				a[num]++
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		a := make([]int, 10)
		for j := 0; j < len(grid); j++ {
			sym := grid[j][i][0]
			if sym != '.' {
				num := sym - '0'
				if a[num] > 0 {
					return false
				}
				a[num]++
			}
		}
	}
	return true
}
