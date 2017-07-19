package main

func main() {

}

func sudoku(grid [][]int) bool {
	for j := 0; j < 9; j++ {
		v := make([]int, 9)
		for i := 0; i < 9; i++ {
			if v[grid[j][i]] != 0 {
				return false
			}
			v[grid[j][i]]++
		}
	}
	for j := 0; j < 9; j++ {
		v := make([]int, 9)
		for i := 0; i < 9; i++ {
			if v[grid[i][j]] != 0 {
				return false
			}
			v[grid[i][j]]++
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			v := make([]int, 9)
			for k := 0; k < 9; k++ {
				value := grid[i+k/3][j+k%3]
				if v[value] != 0 {
					return false
				}
				v[value]++
			}
		}
	}
	return true
}
