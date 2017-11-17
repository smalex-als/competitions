package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func solve() {
	n := readInt()
	a := readIntArray(n)
	cnt1 := 0
	for i := 0; i < n; i++ {
		if a[i] == 1 {
			cnt1++
		}
	}
	if cnt1 > 0 {
		fmt.Println(n - cnt1)
		return
	}
	best := 100000
	for i := 0; i < n; i++ {
		h := a[i]
		for j := i + 1; j < n; j++ {
			h = gcd(h, a[j])
			if h == 1 {
				best = min(best, j-i+n-1)
			}
		}
	}
	if best == 100000 {
		fmt.Println(-1)
	} else {
		fmt.Println(best)
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
