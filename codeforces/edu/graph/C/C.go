package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	t := readInt()

	for j := 0; j < t; j++ {
		solveSingle()
	}
}

func solveSingle() {
	n := readInt()
	m := readInt()
	k := readInt()
	a := make([]int, k)
	for j := 0; j < k; j++ {
		a[j] = readInt() - 1
	}
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
	if !isPathOrCycle(g, a) {
		fmt.Println("none")
		return
	}
	vis := make([]bool, n)
	if a[0] == a[k-1] {
		simple := true
		for _, cur := range a[1:] {
			if vis[cur] {
				simple = false
				break
			}
			vis[cur] = true
		}
		if simple {
			fmt.Println("simple cycle")
		} else {
			fmt.Println("cycle")
		}
	} else {
		simple := true
		for _, cur := range a {
			if vis[cur] {
				simple = false
				break
			}
			vis[cur] = true
		}
		if simple {
			fmt.Println("simple path")
		} else {
			fmt.Println("path")
		}
	}
}

func isPathOrCycle(g [][]int, a []int) bool {
	cur := a[0]
	for j := 1; j < len(a); j++ {
		next := a[j]
		found := false
		for i := 0; i < len(g[cur]); i++ {
			if g[cur][i] == next {
				found = true
				break
			}
		}
		if !found {
			return false
		}
		cur = next
	}
	return true
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

func max(a, b float64) float64 {
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
