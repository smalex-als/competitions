package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var best = 100000

func solve() {
	n := readInt()
	m := readInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString()
		a[i] += a[i]
	}
	res := make([][]int, n)
	for k, v := range a {
		res[k] = make([]int, 3)
		minL := 10000
		minS := 10000
		minN := 10000
		for i := 0; i < len(v); i++ {
			ch := v[i]
			if ch >= '0' && ch <= '9' {
				minN = min(minN, abs(i-m))
			} else if ch == '#' || ch == '&' || ch == '*' {
				minS = min(minS, abs(i-m))
			} else {
				minL = min(minL, abs(i-m))
			}
		}
		res[k][0] = minL
		res[k][1] = minS
		res[k][2] = minN
	}
	visited := make([]bool, n)
	dfs(res, 0, visited, 0)
	fmt.Println(best)
}

func dfs(cost [][]int, column int, visited []bool, res int) {
	if column == 3 {
		best = min(best, res)
		return
	}
	for i := 0; i < len(visited); i++ {
		if !visited[i] && cost[i][column] != 10000 {
			visited[i] = true
			dfs(cost, column+1, visited, res+cost[i][column])
			visited[i] = false
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
