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
		n := readInt()
		m := readInt()
		g := make([][]int, n)
		for j := 0; j < n; j++ {
			g[j] = make([]int, 0)
		}
		for j := 0; j < m; j++ {
			u := readInt() - 1
			v := readInt() - 1
			g[u] = append(g[u], v)
			g[v] = append(g[v], u)
		}
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(len(g[j]))
		}
		fmt.Println()
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int64) int64 {
	if a > b {
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
