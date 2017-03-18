package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	k := readInt()
	m := n - k + 1
	ok := make([]bool, m)
	for i := 0; i < m; i++ {
		ok[i] = readString() == "YES"
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}
	names := make([]string, n)
	for i := 0; i < n; i++ {
		names[i] = string('A'+(i/26)) + string('a'+(i%26))
	}
	for i := 0; i < m; i++ {
		if !ok[i] {
			res[i+k-1] = res[i]
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf(names[res[i]])
	}
	fmt.Println()
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
