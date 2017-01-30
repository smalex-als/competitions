package main

import "fmt"

func main10() {
	files := [][]int{
		{461618501, 3},
		{461618502, 1},
		{461618504, 2},
		{461618506, 5},
		{461618507, 6},
	}
	backups := []int{461618501, 461618502, 461618504, 461618505, 461618506}
	res := troubleFiles(files, backups)
	fmt.Printf("res = %+v\n", res)
}

func troubleFiles(files [][]int, backups []int) []int {
	prevtime := 0
	endtime := 0
	res := make([]int, len(backups))
	cancel := make([]bool, len(files))
	for index, time := range backups {
		reqTime := 0
		for i, file := range files {
			if !cancel[i] && file[0] > prevtime && file[0] <= time {
				reqTime += file[1]
			}
		}
		if reqTime > 0 {
			endtime = time + reqTime
			for i, file := range files {
				if file[0] > time && file[0] <= endtime {
					cancel[i] = true
					res[index]++
				}
			}
		}
		prevtime = time
	}
	return res
}
