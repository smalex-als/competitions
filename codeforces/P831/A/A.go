package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	a := make([]int, n)
	freq := make([]int, 1002)
	start := make([]int, 1002)
	end := make([]int, 1002)
	up := 0
	for i := 0; i < n; i++ {
		a[i] = readInt()
		if freq[a[i]] == 0 {
			start[a[i]] = i
		}
		freq[a[i]]++
		end[a[i]] = i
		up = max(up, a[i])
	}
	index := up
	for i := 1; i <= start[index]; i++ {
		if a[i] <= a[i-1] {
			fmt.Println("NO")
			return
		}
	}
	for i := start[index]; i <= end[index]; i++ {
		if a[i] != index {
			fmt.Println("NO")
			return
		}
	}
	for i := end[index] + 1; i < len(a); i++ {
		if a[i] >= a[i-1] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
