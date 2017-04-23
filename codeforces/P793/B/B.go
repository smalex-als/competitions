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
	mm := make([][]byte, n)
	startX := -1
	startY := -1
	for i := 0; i < n; i++ {
		mm[i] = []byte(readString())
		if startX == -1 {
			for j := 0; j < m; j++ {
				if 'S' == mm[i][j] {
					startY = i
					startX = j
					break
				}
			}
		}
	}
	yx := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for j := range yx {
		vis := make([][]int, n)
		for i := 0; i < n; i++ {
			vis[i] = make([]int, m)
			for j := 0; j < m; j++ {
				vis[i][j] = 1000000
			}
		}

		q := make([][4]int, 0)
		q = append(q, [4]int{startY, startX, 0, j})
		vis[startY][startX] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for i, curdir := range yx {
				cnt := cur[2]
				if i != cur[3] {
					cnt++
					if cnt > 2 {
						continue
					}
				}
				ny := cur[0] + curdir[0]
				nx := cur[1] + curdir[1]
				if ny >= 0 && ny < n && nx >= 0 && nx < m &&
					mm[ny][nx] != '*' && vis[ny][nx] > cnt {
					if mm[ny][nx] == 'T' {
						fmt.Println("YES")
						return
					}
					q = append(q, [4]int{ny, nx, cnt, i})
					vis[ny][nx] = cnt
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
