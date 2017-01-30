package main

func efficientRoadNetwork2(n int, r [][]int) bool {
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			d[i][j] = n + 1 // "infinite dist"
		}
		d[i][i] = 0 // dist to myself
	}
	// initial weights
	for _, v := range r {
		d[v[0]][v[1]] = 1
		d[v[1]][v[0]] = 1
	}
	// Floyd Warshall
	for k := range d {
		for i := range d {
			for j := range d {
				if d[i][j] > d[i][k]+d[k][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	for _, v := range d {
		for _, v := range v {
			if v > 2 {
				return false
			}
		}
	}
	return true
}
