package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt64()
	l := readInt64() - 1
	r := readInt64() - 1
	if n == 0 {
		fmt.Println(0)
		return
	}
	tmpn := int64(1)
	var req int64
	for tmpn < n {
		tmpn <<= 1
		req = req*2 + 1
	}
	dfs(0, req-1, n, l, r)
	fmt.Println(cnt)
}

var cnt int

func dfs(lo int64, hi int64, val int64, l int64, r int64) {
	if val == 0 {
		return
	}
	if l > hi || lo > r {
		return
	}
	mid := (lo + hi) / 2
	if val%2 == 1 && l <= mid && mid <= r {
		cnt++
	}
	dfs(lo, mid-1, val/2, l, r)
	dfs(mid+1, hi, val/2, l, r)
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
