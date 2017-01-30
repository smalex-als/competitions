package main

import "fmt"

func main04() {
	roadRegister := [][]bool{
		{false, true, true, false},
		{true, false, true, false},
		{true, true, false, true},
		{false, false, true, false},
	}
	res := financialCrisis(roadRegister)
	fmt.Printf("res = %+v\n", res)
}

func financialCrisis(g [][]bool) [][][]bool {
	res := make([][][]bool, 0)
	for i := 0; i < len(g); i++ {
		newg := make([][]bool, 0)
		for j := 0; j < len(g); j++ {
			if i != j {
				row := make([]bool, 0)
				for k := 0; k < len(g); k++ {
					if k != i {
						row = append(row, g[j][k])
					}
				}
				newg = append(newg, row)
			}
		}
		res = append(res, newg)
	}
	return res
}
