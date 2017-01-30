package main

import "fmt"

func main18() {
	friendlyTroops := []int{-2, 2, 3}
	enemyTroops := []int{1, 0, 9}
	loggingCamp := []int{0, 0}
	impassableCells := [][]int{{-1, 1}}
	res := farmingResources(friendlyTroops, enemyTroops, loggingCamp, impassableCells)
	fmt.Printf("res = %+v\n", res)
}

func farmingResources(friendlyTroops, enemyTroops, loggingCamp []int,
	impassableCells [][]int) bool {
	len1 := traceLen(friendlyTroops, loggingCamp, impassableCells)
	len2 := traceLen(enemyTroops, loggingCamp, impassableCells)
	time1 := len1 * friendlyTroops[2]
	time2 := len2 * enemyTroops[2]
	if time1 <= time2 {
		return true
	}
	return false
}

func traceLen(start, end []int, impassableCells [][]int) int {
	yx := [][]int{{0, -1}, {0, 1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}}
	const N = 101
	const N_2 = 50
	var field [N][N]int
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			field[i][j] = 100000
		}
	}
	for _, v := range impassableCells {
		field[v[0]+N_2][v[1]+N_2] = -1
	}
	field[start[0]+N_2][start[1]+N_2] = 0
	queue := make([][]int, 1)
	queue[0] = []int{start[0] + N_2, start[1] + N_2}
	destY := end[0] + N_2
	destX := end[1] + N_2
	for len(queue) > 0 {
		cur := queue[0]
		curLen := field[cur[0]][cur[1]]
		queue = queue[1:]
		if cur[0] == destY && cur[1] == destX {
			return curLen
		}
		for _, v := range yx {
			ny := cur[0] + v[0]
			nx := cur[1] + v[1]
			if ny >= 0 && ny < N && nx >= 0 && nx < N && field[ny][nx] > curLen+1 {
				field[ny][nx] = curLen + 1
				queue = append(queue, []int{ny, nx})
			}
		}
	}
	return 10000
}
