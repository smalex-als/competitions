package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type p struct {
	x, y int
}

func solve() {
	x1 := readInt()
	y1 := readInt()
	x2 := readInt()
	y2 := readInt()
	a := readInt()
	b := readInt()
	v := make(map[p]bool)
	q := make([]p, 0)
	q = append(q, p{x1, y1})
	v[q[0]] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.x == x2 && cur.y == y2 {
			fmt.Println("YES")
			return
		}
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i != 0 && j != 0 {
					np := p{cur.x + a*i, cur.y + b*j}
					if !v[np] && abs(np.x) <= abs(y2) && abs(np.y) <= abs(y2) {
						q = append(q, np)
						v[np] = true
					}
				}
			}
		}
	}
	fmt.Println("NO")
}

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func min(a int, b int) int {
	if a < b {
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
