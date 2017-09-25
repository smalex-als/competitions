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
	v := make([]int64, 3)
	for i := 0; i < 3; i++ {
		v[i] = readInt64()
	}
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = readInt64()
	}
	suffix := make([]int64, n)
	for j := n - 1; j >= 0; j-- {
		if j == n-1 {
			suffix[j] = v[2] * a[j]
		} else {
			suffix[j] = max(suffix[j+1], v[2]*a[j])
		}
	}
	bestFirst := int64(math.MinInt64)
	res := int64(math.MinInt64)
	for i := 0; i < n; i++ {
		bestFirst = max(bestFirst, v[0]*a[i])
		best := bestFirst + (v[1] * a[i]) + suffix[i]
		if best > res {
			res = best
		}
	}
	fmt.Println(res)
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
