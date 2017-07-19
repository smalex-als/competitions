package main

import "fmt"

func main() {
	arr := []int{3, 0, -2, 6, -3, 2}
	queries := [][]int{{0, 2}, {2, 5}, {0, 5}}
	res := sumInRange(arr, queries)
	fmt.Printf("res = %+v\n", res)
}

func sumInRange(nums []int, queries [][]int) int {
	tree := NewFenwickTree(len(nums))
	for k, v := range nums {
		tree.Add(k, v)
	}
	res := 0
	for i := 0; i < len(queries); i++ {
		res += tree.SumRange(queries[i][0], queries[i][1])
		res = res % 1000000007
	}
	return res
}

type FenwickTree struct {
	t []int
	n int
}

func NewFenwickTree(n int) *FenwickTree {
	s := FenwickTree{t: make([]int, n), n: n}
	return &s
}

func (s *FenwickTree) Add(i, value int) {
	for ; i < s.n; i += (i + 1) & -(i + 1) {
		s.t[i] += value
	}
}

func (s *FenwickTree) Sum(i int) int {
	res := 0
	for ; i >= 0; i -= (i + 1) & -(i + 1) {
		res += s.t[i]
	}
	return res
}

func (s *FenwickTree) SumRange(a, b int) int {
	return s.Sum(b) - s.Sum(a-1)
}
