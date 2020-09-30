package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		a := readIntArray(n)
		b := readIntArray(n)
		c := readIntArray(n)
		res := make([]int, n)
		res[0] = a[0]
		for j := 1; j < n; j++ {
			prev := res[j-1]
			cur := 0
			if j == n-1 {
				if prev != a[j] && res[0] != a[j] {
					cur = a[j]
				} else if prev != b[j] && res[0] != b[j] {
					cur = b[j]
				} else if prev != c[j] && res[0] != c[j] {
					cur = c[j]
				}
			} else {
				if prev != a[j] {
					cur = a[j]
				} else if prev != b[j] {
					cur = b[j]
				} else if prev != c[j] {
					cur = c[j]
				}
			}
			res[j] = cur
		}
		for j := 0; j < n; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(res[j])
		}
		fmt.Println()
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a int64, b int64) int64 {
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
