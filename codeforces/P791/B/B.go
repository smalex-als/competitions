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
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, 0)
	}
	for i := 0; i < m; i++ {
		v := readInt() - 1
		u := readInt() - 1
		matrix[v] = append(matrix[v], u)
		matrix[u] = append(matrix[u], v)
	}
	ans := "YES"
	visit := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			var size int64
			var edges int64
			dfs(i, matrix, visit, &edges, &size)
			if size*(size-1) != edges {
				ans = "NO"
				break
			}
		}
	}
	fmt.Println(ans)
}

func dfs(u int, matrix [][]int, visited []bool, edges, size *int64) {
	visited[u] = true
	*size++
	*edges += int64(len(matrix[u]))
	for _, v := range matrix[u] {
		if !visited[v] {
			dfs(v, matrix, visited, edges, size)
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
