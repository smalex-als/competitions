package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	m := readInt()
	sizes := make([]int, n)
	set := NewDJSet(n)
	for i := 0; i < m; i++ {
		v := readInt() - 1
		u := readInt() - 1
		set.Union(v, u)
		sizes[u]++
		sizes[v]++
	}
	ans := true

	byroot := make([]int, n)
	for i := 0; i < n; i++ {
		root := set.Root(i)
		byroot[root] += sizes[i]
	}
	for i := 0; i < n; i++ {
		size := int64(-set.Upper(i))
		if size > 0 {
			if size*(size-1) != int64(byroot[i]) {
				ans = false
				break
			}
		}
	}
	if ans {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
