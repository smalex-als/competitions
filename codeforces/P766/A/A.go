package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	a := readString()
	b := readString()
	if len(a) > len(b) {
		a, b = b, a
	}
	if len(a) != len(b) {
		fmt.Println(len(b))
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			fmt.Println(len(b))
			return
		}
	}
	fmt.Println(-1)
}

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

func min(a int, b int) int {
	if a < b {
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
