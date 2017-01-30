package main

import (
	"fmt"
	"sort"
)

func main8() {
	// sentBatches := [][]int{
	// 	{1471040000, 736273, 827482, 2738283},
	// 	{1471040005, 736273, 2738283},
	// 	{1471040010, 827482, 2738283},
	// 	{1471040015, 2738283},
	// 	{1471040025, 827482},
	// 	{1471046400, 736273, 827482, 2738283},
	// }
	// receivedMessages := [][]int{
	// 	{1471040001, 2738283},
	// 	{1471040002, 2738283},
	// 	{1471040010, 827482},
	// 	{1471040020, 2738283},
	// }
	// startingAllowance := 1

	sentBatches := [][]int{{1471046400, 1, 2, 3},
		{1471046400, 1, 2},
		{1471046400, 1, 3},
		{1471046400, 1},
		{1471046400, 1},
		{1471046400, 2, 3, 6}}
	receivedMessages := [][]int{{1471046400, 3},
		{1471046400, 3},
		{1471046400, 1},
		{1471046400, 3},
		{1471046400, 4}}
	startingAllowance := 2
	res := rateLimit(sentBatches, receivedMessages, startingAllowance)
	fmt.Println(res)
}

type ByTime [][]int

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	msg0 := a[i]
	msg1 := a[j]
	if msg0[0] != msg1[0] {
		return msg0[0] < msg1[0]
	}
	type0 := msg0[len(msg0)-2]
	type1 := msg1[len(msg1)-2]
	if type0 != type1 {
		return type0 < type1
	}
	index0 := msg0[len(msg0)-1]
	index1 := msg1[len(msg1)-1]
	return index0 < index1
}

func rateLimit(sent [][]int, received [][]int, limit int) []int {
	const timeReset = 86400
	queue := make([][]int, 0)
	for i := 0; i < len(sent); i++ {
		sent[i] = append(sent[i], 1)
		sent[i] = append(sent[i], i)
		queue = append(queue, sent[i])
	}
	for i := 0; i < len(received); i++ {
		received[i] = append(received[i], 0)
		received[i] = append(received[i], i)
		queue = append(queue, received[i])
	}
	sort.Sort(ByTime(queue))
	limitByUser := make(map[int]int)
	res := make([]int, 0)
	prevTime := 0
	for len(queue) > 0 {
		cmd := queue[0]
		queue = queue[1:]
		time := cmd[0]
		if time/timeReset != prevTime {
			limitByUser = make(map[int]int)
		}
		prevTime = time / timeReset
		cmdType := cmd[len(cmd)-2]
		index := cmd[len(cmd)-1]
		userIds := cmd[1 : len(cmd)-2]
		fmt.Println(time)
		if cmdType == 0 {
			fmt.Println("   receive ", userIds)
			for _, v := range userIds {
				limitByUser[v]++
			}
		} else {
			fmt.Println("   send ", userIds)
			ok := true
			for _, v := range userIds {
				if limitByUser[v] <= -limit {
					fmt.Println("failed", index)
					res = append(res, index)
					ok = false
					break
				}
			}
			if ok {
				for _, v := range userIds {
					limitByUser[v]--
				}
			}
		}
	}
	return res
}
