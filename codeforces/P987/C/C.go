package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	s := readIntArray(n)
	c := readIntArray(n)
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		l := math.MaxInt32
		r := math.MaxInt32
		for j := i + 1; j < n; j++ {
			if s[i] < s[j] {
				r = min(r, c[j])
			}
		}
		for j := i - 1; j >= 0; j-- {
			if s[j] < s[i] {
				l = min(l, c[j])
			}
		}
		if l != math.MaxInt32 && r != math.MaxInt32 {
			ans = min(ans, c[i]+r+l)
		}
	}
	if ans == math.MaxInt32 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
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
