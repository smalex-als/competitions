package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var used []bool

func solve() {
	n := readInt()
	used = make([]bool, n)
	adj := make([][]int, n)
	for i := 1; i < n; i++ {
		u := readInt() - 1
		v := readInt() - 1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	fmt.Printf("%f", dfs(adj, 0))
}

func dfs(adj [][]int, v int) float64 {
	used[v] = true
	cnt := 0
	var sum float64
	for _, u := range adj[v] {
		if !used[u] {
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	for _, u := range adj[v] {
		if !used[u] {
			sum += dfs(adj, u) + float64(1)
		}
	}
	return sum / float64(cnt)
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
