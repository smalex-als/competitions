package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve() {
	n := readInt()
	K := readInt()
	a := readIntArray(n)
	if n == 1 {
		fmt.Println(1)
		return
	}
	sort.Ints(a)
	cur := 0
	offset := 1
	ans := 0
	for offset < len(a) {
		if cur == offset || a[cur] == a[offset] {
			offset++
			continue
		}
		if a[cur] < a[offset] && a[cur]+K >= a[offset] {
			ans++
			cur++
		} else {
			cur++
		}
	}
	fmt.Println(n - ans)

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
