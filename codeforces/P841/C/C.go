package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type A struct {
	value int
	index int
}

func solve() {
	n := readInt()
	a := make([]A, n)
	for i := 0; i < n; i++ {
		a[i] = A{readInt(), i}
	}
	sort.Sort(ByValue(a))
	djset := NewDJSet(n)
	for i := 0; i < n; i++ {
		cur := a[i]
		with := a[cur.index]
		if with.index == i {
			djset.Union(with.index, cur.index)
		} else {
			djset.Union(i, cur.index)
		}
	}
	byroot := make([][]string, n)
	for i := 0; i < n; i++ {
		byroot[i] = make([]string, 0)
	}
	for i := 0; i < n; i++ {
		root := djset.Root(i)
		byroot[root] = append(byroot[root], strconv.Itoa(i+1))
	}
	fmt.Println(djset.Count())
	for i := 0; i < n; i++ {
		if len(byroot[i]) > 0 {
			fmt.Printf("%d %s\n", len(byroot[i]), strings.Join(byroot[i], " "))
		}
	}
}

type ByValue []A

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

type DJSet struct {
	upper []int
}

func NewDJSet(n int) *DJSet {
	s := DJSet{upper: make([]int, n)}
	for i := 0; i < n; i++ {
		s.upper[i] = -1
	}
	return &s
}

func (s *DJSet) Root(x int) int {
	if s.upper[x] < 0 {
		return x
	} else {
		s.upper[x] = s.Root(s.upper[x])
	}
	return s.upper[x]
}

func (s *DJSet) Union(x, y int) bool {
	x = s.Root(x)
	y = s.Root(y)
	if x != y {
		if s.upper[y] < s.upper[x] {
			x, y = y, x
		}
		s.upper[x] += s.upper[y]
		s.upper[y] = x
	}
	return x == y
}

func (s *DJSet) Equiv(x, y int) bool {
	return s.Root(x) == s.Root(y)
}

func (s *DJSet) Upper(y int) int {
	return s.upper[y]
}

func (s *DJSet) Count() int {
	ct := 0
	for _, u := range s.upper {
		if u < 0 {
			ct++
		}
	}
	return ct
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
