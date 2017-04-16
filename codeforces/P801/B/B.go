package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	x := []byte(readString())
	y := []byte(readString())
	z := make([]byte, len(x))
	for i := 0; i < len(x); i++ {
		if x[i] < y[i] {
			fmt.Println(-1)
			return
		}
		if x[i] == y[i] {
			z[i] = x[i]
		} else {
			z[i] = y[i]
		}
	}
	fmt.Println(string(z))
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
