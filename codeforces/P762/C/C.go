package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var vis [][]int

var ans = ""

func solve() {
	a := []byte(readString())
	b := []byte(readString())
	n := max(len(a), len(b)) + 10
	vis = make([][]int, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			vis[i][j] = -1
		}
	}
	search(a, b, 0, 0, "")
	if len(ans) > 0 {
		fmt.Println(ans)
	} else {
		fmt.Println("-")
	}
}

func search(a, b []byte, posA, posB int, cur string) {
	if vis[posA][posB] >= len(cur) {
		return
	}
	vis[posA][posB] = len(cur)
	if posA == len(a) || posB == len(b) {
		if len(ans) < len(cur) {
			ans = cur
		}
		return
	}
	if a[posA] == b[posB] {
		search(a, b, posA+1, posB+1, cur+string(a[posA]))
	}
	search(a, b, posA+1, posB, cur)
	search(a, b, posA, posB+1, cur)
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
