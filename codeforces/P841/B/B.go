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

	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			f[i][j] = readInt() == 1
		}
	}
	res := int64(0)
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < m; j++ {
			if f[i][j] {
				cnt++
			}
		}
		res += getFor(cnt)
		res += getFor(m - cnt)
	}
	for i := 0; i < m; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if f[j][i] {
				cnt++
			}
		}
		res += getFor(cnt) - int64(cnt)
		res += getFor(n-cnt) - int64(n-cnt)
	}
	fmt.Println(res)
}

func getFor(n int) int64 {
	res := int64(0)
	for k := 1; k <= n; k++ {
		res += rec(n, k)
	}
	return res
}

var memo [60][60]int64

func rec(n, k int) int64 {
	if n == k {
		return 1
	}
	if k == 1 {
		return int64(n)
	}
	if memo[n][k] != 0 {
		return memo[n][k]
	}
	res := rec(n-1, k-1) + rec(n-1, k)
	memo[n][k] = res
	return res
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
