package main

import "fmt"

func main01() {
	roadRegister := [][]bool{
		{false, true, false, false},
		{false, false, true, false},
		{true, false, false, true},
		{false, false, true, false},
	}
	res := newRoadSystem(roadRegister)
	fmt.Printf("res = %+v\n", res)
}

func newRoadSystem(roadRegister [][]bool) bool {
	for i := 0; i < len(roadRegister); i++ {
		cntIn := 0
		cntOut := 0
		for j := 0; j < len(roadRegister); j++ {
			if roadRegister[i][j] {
				cntIn++
			}
			if roadRegister[j][i] {
				cntOut++
			}
		}
		if cntIn != cntOut {
			return false
		}
	}
	return true

}
