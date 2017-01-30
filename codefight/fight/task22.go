package main

import (
	"fmt"
	"sort"
)

func main22() {
	// jobs := []int{15, 30, 15, 5, 10}
	// jobs := []int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	jobs := []int{8, 7, 15, 15, 13, 6, 18, 4, 16, 1, 2, 19, 2, 15, 18, 6, 20, 16, 10, 7, 3, 7, 9, 7, 12, 1, 16, 15, 7, 12, 20, 17, 17, 4, 20, 15, 20, 6, 15, 3, 5, 17, 5, 5, 19, 17, 4, 15, 2, 7}
	servers := 9
	res := serverFarm(jobs, servers)
	fmt.Printf("res = %+v\n", res)
}

type Item struct {
	index, value int
}

type SortByValueRev []Item

func (a SortByValueRev) Len() int      { return len(a) }
func (a SortByValueRev) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByValueRev) Less(i, j int) bool {
	if a[j].value == a[i].value {
		return a[j].index > a[i].index
	}
	return a[j].value < a[i].value
}

func serverFarm(jobs []int, servers int) [][]int {
	items := make([]Item, len(jobs))
	for i := 0; i < len(jobs); i++ {
		items[i] = Item{i, jobs[i]}
	}
	sort.Sort(SortByValueRev(items))
	slots := make([][]int, servers)
	sum := make([]int, servers)
	for _, v := range items {
		minIndex := 0
		for i := 0; i < servers; i++ {
			if sum[i] < sum[minIndex] {
				minIndex = i
			}
		}
		slots[minIndex] = append(slots[minIndex], v.index)
		sum[minIndex] += v.value
	}
	return slots
}
