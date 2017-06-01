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
	a := make([]string, n)
	cnt := 0
	for i := 0; i < n; i++ {
		a[i] = readString()
		for _, v := range a[i] {
			if v == '1' {
				cnt++
			}
		}
	}
	fmt.Println(dfs(n, m+2, a, n-1, true, cnt))
}

func dfs(n, m int, a []string, y int, left bool, cnt int) int {
	if cnt <= 0 {
		return 0
	}
	need := 0
	localFound := 0
	if left {
		steps := 0
		for i := 0; i < m; i++ {
			if a[y][i] == '1' {
				need = steps
				localFound++
			}
			steps++
		}
	} else {
		steps := 0
		for i := m - 1; i >= 0; i-- {
			if a[y][i] == '1' {
				need = steps
				localFound++
			}
			steps++
		}
	}
	if localFound == 0 {
		return 1 + dfs(n, m, a, y-1, left, cnt)
	} else if localFound == cnt {
		return need
	} else {
		return min(
			need+need+1+dfs(n, m, a, y-1, left, cnt-localFound),
			m+dfs(n, m, a, y-1, !left, cnt-localFound))
	}
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
