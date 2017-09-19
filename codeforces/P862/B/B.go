package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var g [][]int
var color []int

func solve() {
	n := readInt()
	g = make([][]int, n)
	color = make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = -1
	}
	for i := 0; i < n-1; i++ {
		u := readInt() - 1
		v := readInt() - 1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	dfs(0, 1)
	cnt := 0
	for i := 0; i < n; i++ {
		if color[i] != 1 {
			cnt++
		}
	}
	res := int64(0)
	for i := 0; i < n; i++ {
		if color[i] == 1 {
			res += int64(cnt - len(g[i]))
		}
	}
	fmt.Println(res)
}

func dfs(v int, c int) {
	color[v] = c
	for i := 0; i < len(g[v]); i++ {
		next := g[v][i]
		if color[next] == -1 {
			dfs(next, ^c)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
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
