package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type point struct {
	x    int
	l, r int
	c    int
	open int
}

func solve() {
	n := readInt()
	x := readInt()
	q := make([]point, n*2)
	for i := 0; i < n; i++ {
		l := readInt()
		r := readInt()
		c := readInt()
		q[i*2] = point{l, l, r, c, -1}
		q[i*2+1] = point{r, l, r, c, 1}
	}
	ans := math.MaxInt32
	sort.Sort(ByLeft(q))
	mm := map[int]int{}
	for i := 0; i < len(q); i++ {
		p := q[i]
		if p.open == -1 {
			y := p.r - p.l + 1
			if x >= y {
				if c, ok := mm[x-y]; ok {
					ans = min(ans, c+p.c)
				}
			}
		} else {
			y := p.r - p.l + 1
			if v, ok := mm[y]; ok {
				if p.c < v {
					mm[y] = p.c
				}
			} else {
				mm[y] = p.c
			}
		}
	}
	if ans == math.MaxInt32 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

type ByLeft []point

func (a ByLeft) Len() int      { return len(a) }
func (a ByLeft) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByLeft) Less(i, j int) bool {
	if a[i].x == a[j].x {
		return a[i].open < a[j].open
	}
	return a[i].x < a[j].x
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
