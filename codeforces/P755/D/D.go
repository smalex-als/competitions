package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	// k := readInt()
	var buffer bytes.Buffer
	for i := 1; i <= n; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}
	t := NewFenwickTree(n + 10)
	t.Add(0, 1)
	t.Add(1, 1)
	fmt.Println(t.Sum(0))
	fmt.Println(t.Sum(1))
	fmt.Println(t.SumRange(1, 1))
	fmt.Printf("t = %+v\n", t)
	fmt.Println(buffer.String())
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

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

// IO

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func readInt64() int64 {
	v, _ := strconv.ParseInt(readString(), 10, 64)
	return v
}

func readIntArray(sz int) []int {
	a := make([]int, sz)
	for i := 0; i < sz; i++ {
		a[i] = readInt()
	}
	return a
}

func readInt64Array(n int) []int64 {
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = readInt64()
	}
	return res
}
