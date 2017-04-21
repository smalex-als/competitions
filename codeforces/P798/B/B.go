package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString()
	}
	ans := 1000000
outer:
	for i := 0; i < len(a[0]); i++ {
		b := a[0][i:] + a[0][:i]
		total := i
	inner:
		for j := 1; j < n; j++ {
			for k := 0; k < len(a[0]); k++ {
				c := a[j][k:] + a[j][:k]
				if c == b {
					total += k
					continue inner
				}
			}
			continue outer
		}
		ans = min(total, ans)
	}
	if ans == 1000000 {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
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
