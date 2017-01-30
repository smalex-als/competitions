package main

import (
	"fmt"
	"sort"
)

func main26() {
	// topic_id for i-th question
	topicIds := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{1, 4},
		{2, 3},
	}
	// answer_id for i-th question
	answerIds := [][]int{
		{6, 4},
		{1, 2},
		{5},
		{3},
	}
	// answer_id, user_id, cnt
	views := [][]int{
		{2, 1, 2},
		{6, 3, 5},
		{3, 3, 0},
		{5, 1, 1},
		{4, 2, 3},
		{1, 4, 2},
	}
	res := mostViewedWriters(topicIds, answerIds, views)
	fmt.Printf("res = %+v\n", res)
}

func mostViewedWriters(topicIds [][]int, answerIds [][]int, views [][]int) [][][]int {
	res := make([][][]int, 0)
	questionsByTopic := make(map[int][]int)
	topicKeys := make([]int, 0)
	for i := 0; i < len(topicIds); i++ {
		for j := 0; j < len(topicIds[i]); j++ {
			q := topicIds[i][j]
			if _, ok := questionsByTopic[q]; !ok {
				questionsByTopic[q] = make([]int, 0)
				topicKeys = append(topicKeys, q)
			}
			questionsByTopic[q] = append(questionsByTopic[q], i)
		}
	}
	sort.Ints(topicKeys)
	for i := 0; i < len(topicKeys); i++ {
		arr := questionsByTopic[topicKeys[i]]
		acceptAnswers := make(map[int]bool)
		for j := 0; j < len(arr); j++ {
			row := answerIds[arr[j]]
			for ii := 0; ii < len(row); ii++ {
				acceptAnswers[row[ii]] = true
			}
		}
		stats := make(map[int]int)
		for _, v := range views {
			if _, ok := acceptAnswers[v[0]]; ok {
				stats[v[1]] += v[2]
			}
		}
		row := make([][]int, 0)
		for k, v := range stats {
			row = append(row, []int{k, v})
		}
		sort.Sort(ByCnt(row))
		res = append(res, row)
	}
	return res
}

type ByCnt [][]int

func (a ByCnt) Len() int      { return len(a) }
func (a ByCnt) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCnt) Less(i, j int) bool {
	if a[i][1] == a[j][1] {
		return a[i][0] < a[j][0]
	}
	return a[i][1] > a[j][1]
}
