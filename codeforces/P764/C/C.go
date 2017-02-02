package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	n    int
	g    [][]int
	c    []int
	memo map[int64]bool
)

func solve() {
	memo = make(map[int64]bool)
	n = readInt()
	g = make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
	}
	for i := 0; i < n-1; i++ {
		u := readInt() - 1
		v := readInt() - 1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	c = make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = readInt()
	}
	for i := 0; i < n; i++ {
		ok := true
		for _, v := range g[i] {
			if !dfs(v, i, c[v]) {
				ok = false
				break
			}
		}
		if ok {
			fmt.Println("YES")
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println("NO")
}

func dfs(v int, from int, color int) bool {
	if c[v] != color {
		return false
	}
	key := int64(v)*math.MaxInt32 + int64(from)
	if v, ok := memo[key]; ok {
		return v
	}
	res := true
	for _, u := range g[v] {
		if u != from {
			if !dfs(u, v, color) {
				res = false
				break
			}
		}
	}
	memo[key] = res
	return res
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
