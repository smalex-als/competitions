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
		n := readInt()
		a := readIntArray(n)
		solveSingle(a)
	}
}

func solveSingle(a []int) {
	n := len(a)
	connected := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		connected[i] = make(map[int]bool)
	}
	for i := 0; i < n; i++ {
		if len(connected[i]) > 0 {
			continue
		}
		for j := 0; j < n; j++ {
			if a[i] != a[j] && !connected[i][j] {
				connected[i][j] = true
				connected[j][i] = true
				break
			}
		}
	}
	u := n - 1
	if len(connected[u]) == 0 {
		for j := 0; j < n-1; j++ {
			if a[u] != a[j] && !connected[u][j] {
				connected[u][j] = true
				connected[j][u] = true
				break
			}
		}
	}
	vis := make(map[int]bool, n)
	dfs(connected, 0, vis)
	if len(vis) == n {
		fmt.Println("YES")
		for v := 0; v < n; v++ {
			for u := range connected[v] {
				if v < u {
					fmt.Println(v+1, u+1)
				}
			}
		}
	} else {
		fmt.Println("NO")
	}
}

func dfs(connected map[int]map[int]bool, from int, vis map[int]bool) {
	vis[from] = true
	for next := range connected[from] {
		if !vis[next] {
			dfs(connected, next, vis)
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

func max(a, b int64) int64 {
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
