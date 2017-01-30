package main

func combs(comb1 string, comb2 string) int {
	if len(comb1) > len(comb2) {
		comb1, comb2 = comb2, comb1
	}
	a1 := []rune(comb1)
	a2 := []rune(comb2)
	best := 10000
	lenA1 := len(a1)
	lenA2 := len(a2)
	a := make([]rune, 2*lenA1+lenA2)
	for i := 0; i < lenA2+lenA1; i++ {
		for j := 0; j < len(a); j++ {
			a[j] = ' '
		}
		for j := 0; j < lenA2; j++ {
			a[j+lenA1-1] = a2[j]
		}
		ok := true
		for j := 0; j < lenA1; j++ {
			if a[j+i] == a1[j] && a1[j] == '*' {
				ok = false
				break
			} else {
				if a1[j] == '*' || a[j+i] == '*' {
					a[j+i] = '*'
				} else {
					a[j+i] = a1[j]
				}
			}
		}
		if ok {
			cnt := 0
			for j := 0; j < len(a); j++ {
				if a[j] != ' ' {
					cnt++
				}
			}
			if cnt < best {
				best = cnt
			}
		}
	}
	return best
}

func crossingSum(matrix [][]int, a int, b int) int {
	n := len(matrix)
	total := 0
	for i := 0; i < len(matrix[0]); i++ {
		total += matrix[a][i]
	}
	for i := 0; i < n; i++ {
		if i != a {
			total += matrix[i][b]
		}
	}
	return total
}

func reverseOnDiagonals(matrix [][]int) [][]int {
	n := len(matrix)
	for i, j := 0, n-1; i < n/2; i++ {
		matrix[i][i], matrix[j][j] = matrix[j][j], matrix[i][i]
		matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		j--
	}
	return matrix
}

func swapDiagonals(matrix [][]int) [][]int {
	n := len(matrix)
	for i, j := 0, n-1; i < n/2; i++ {
		matrix[i][i], matrix[j][i] = matrix[j][i], matrix[i][i]
		matrix[i][j], matrix[j][j] = matrix[j][j], matrix[i][j]
		j--
	}
	return matrix
}
