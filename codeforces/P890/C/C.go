package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = readInt() - 1
	}
	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = -1
	}
	id := 0
	for i := n - 1; i >= 0; i-- {
		if values[i] != -1 {
			continue
		}
		id++
		next := i
		for next >= 0 && values[next] == -1 {
			values[next] = id
			next = t[next]
		}
	}
	fmt.Println(id)
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

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
