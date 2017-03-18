package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var y []int
var x []int

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
	visited := make([]bool, n)
	ans := true
	for i := 0; i < n; i++ {
		if !visited[i] {
			cur := make([]int, 0)
			res := int64(dfs(i, matrix, visited, &cur))
			size := int64(len(cur))
			if size*(size-1) != res {
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

func dfs(u int, matrix [][]int, visited []bool, cur *[]int) int {
	*cur = append(*cur, u)
	visited[u] = true
	res := 0
	for i := 0; i < len(matrix[u]); i++ {
		v := matrix[u][i]
		if !visited[v] {
			res += dfs(v, matrix, visited, cur)
		}
		res++
	}
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
