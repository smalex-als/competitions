package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	t := readInt()
	for i := 0; i < t; i++ {
		solveSingle()
	}
}

func solveSingle() {
	n := readInt()
	m := readInt()
	k := readInt()
	a := make([]int, k)
	for i := 0; i < k; i++ {
		a[i] = readInt() - 1
	}
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
	}
	for i := 0; i < m; i++ {
		u := readInt() - 1
		v := readInt() - 1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	colors := make([]bool, n)
	for _, v := range a {
		if !colors[v] {
			dfs(colors, g, v)
		}
	}
	cnt := 0
	for _, color := range colors {
		if color {
			cnt++
		}
	}
	if cnt == k {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func dfs(colors []bool, g [][]int, u int) {
	colors[u] = true
	for _, v := range g[u] {
		if !colors[v] {
			dfs(colors, g, v)
		}
	}
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
