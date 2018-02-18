package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	sum := make([]int64, n+10)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + readInt64()
	}
	s := readInt()
	f := readInt()
	best := int64(0)
	bestHour := 0
	for i := 1; i <= n; i++ {
		from := s - i + 1
		to := f - i
		for from < 1 {
			from += n
			to += n
		}
		cur := int64(0)
		if to <= n {
			cur = sum[to] - sum[from-1]
		} else {
			cur = sum[to-n] + sum[n] - sum[from-1]
		}
		if cur > best {
			best = cur
			bestHour = i
		}
	}
	fmt.Println(bestHour)
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
