package main

func pairOfShoes(shoes [][]int) bool {
	a := [101][2]int{}
	for i := 0; i < len(shoes); i++ {
		a[shoes[i][1]][shoes[i][0]]++
	}
	for i := 0; i < 101; i++ {
		if a[i][0] != a[i][1] {
			return false
		}
	}
	return true
}
