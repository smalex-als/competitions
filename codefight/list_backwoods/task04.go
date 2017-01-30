package main

func volleyballPositions(formation [][]string, k int) [][]string {
	yx := [][]int{
		{3, 2},
		{1, 2},
		{0, 1},
		{1, 0},
		{3, 0},
		{2, 1},
	}
	names := [6]string{}
	for i, v := range yx {
		names[i] = formation[v[0]][v[1]]
	}
	for i, v := range yx {
		p := (i - k) % 6
		if p < 0 {
			p += 6
		}
		formation[v[0]][v[1]] = names[p]
	}
	return formation
}
