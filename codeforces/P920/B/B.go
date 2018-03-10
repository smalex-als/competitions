package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Student struct {
	Pos int
	l   int
	r   int
}

func solve() {
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		students := make([]Student, n)
		for j := 0; j < n; j++ {
			students[j].Pos = j
			students[j].l = readInt()
			students[j].r = readInt()
		}
		sort.Sort(ByTime(students))
		curtime := 1
		for j := 0; j < n; j++ {
			st := students[j]
			if curtime < st.l {
				curtime = st.l
				students[j].r = curtime + 1
			} else if curtime >= st.r {
				students[j].r = 0
			} else {
				students[j].r = curtime
				curtime++
			}
		}
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(students[j].r)
		}
		fmt.Println()
	}
}

type ByTime []Student

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	if a[i].l == a[j].l {
		return a[i].Pos < a[j].Pos
	}
	return a[i].l < a[j].l
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
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
