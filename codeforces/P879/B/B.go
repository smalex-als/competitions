package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	k := readInt64()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readInt()
	}
	if a[0] < a[1] {
		a[0], a[1] = a[1], a[0]
	}
	wins := int64(0)
	for i := 1; i < n; i++ {
		if a[i-1] > a[i] {
			a[i-1], a[i] = a[i], a[i-1]
			wins++
		} else {
			wins = 1
		}
		if wins == k {
			fmt.Println(a[i])
			return
		}
	}
	fmt.Println(a[len(a)-1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int64) int64 {
	if a > b {
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
