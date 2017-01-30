package main

import "fmt"

func main06() {

	roadRegister := [][]bool{
		{false, true, true, false},
		{true, false, true, false},
		{true, true, false, true},
		{false, false, true, false},
	}
	res := greatRenaming(roadRegister)
	for _, v := range res {
		fmt.Printf("v = %+v\n", v)
	}
}

func greatRenaming(roadRegister [][]bool) [][]bool {
	n := len(roadRegister)
	res := make([][]bool, n)
	for i := 0; i < n; i++ {
		res[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = roadRegister[(i+n-1)%n][(j+n-1)%n]
		}
	}
	return res
}
