package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func main16() {
	shoppers := [][]string{
		{"17:40", "22:30"},
		{"15:10", "16:00"},
	}
	orders := [][]string{
		{"17:30", "18:00"},
		{"15:00", "15:45"},
	}
	leadTime := []int{15, 30}
	res := busyHolidays(shoppers, orders, leadTime)
	fmt.Printf("res = %+v\n", res)
}

func timeToInt(time string) int {
	re := regexp.MustCompile("^([0-9]+):([0-9]+)$")
	m := re.FindStringSubmatch(time)
	hh, _ := strconv.Atoi(m[1])
	mm, _ := strconv.Atoi(m[2])
	return hh*60 + mm
}

func timesToInts(orders [][]string) [][]int {
	s := make([][]int, len(orders))
	for i := 0; i < len(orders); i++ {
		start := timeToInt(orders[i][0])
		end := timeToInt(orders[i][1])
		s[i] = []int{start, end}
	}
	return s
}

type Order struct {
	start,
	end,
	leadTime int
	shoppers []int
}

func timesToOrders(orders [][]string, leadTime []int) []Order {
	s := make([]Order, 0)
	for i := 0; i < len(orders); i++ {
		start := timeToInt(orders[i][0])
		end := timeToInt(orders[i][1])
		s = append(s, Order{start, end, leadTime[i], make([]int, 0)})
	}
	return s
}

type ByLen []Order

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLen) Less(i, j int) bool { return len(a[i].shoppers) < len(a[j].shoppers) }

func busyHolidays(shoppers [][]string, orders [][]string, leadTime []int) bool {
	s := timesToInts(shoppers)
	o := timesToOrders(orders, leadTime)
	for i := 0; i < len(o); i++ {
		order := &o[i]
		arr := order.shoppers
		for j := 0; j < len(s); j++ {
			start := maxTime(order.start, s[j][0])
			end := minTime(order.end, s[j][1])
			if start < end && end-start >= order.leadTime {
				arr = append(arr, j)
			}
		}
		order.shoppers = arr
	}
	sort.Sort(ByLen(o))
	visited := make([]bool, len(shoppers))
	return dfsOrders(o, visited, 0)
}

func dfsOrders(orders []Order, visited []bool, cur int) bool {
	if cur == len(orders) {
		return true
	}
	for i := 0; i < len(orders[cur].shoppers); i++ {
		index := orders[cur].shoppers[i]
		if !visited[index] {
			visited[index] = true
			if dfsOrders(orders, visited, cur+1) {
				return true
			}
			visited[index] = false
		}
	}
	return false
}

func minTime(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxTime(a, b int) int {
	if a > b {
		return a
	}
	return b
}
