package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	a := make([]int, 6)
	total := 0
	for i := 0; i < 6; i++ {
		a[i] = readInt()
		total += a[i]
	}
	if total%2 != 0 {
		fmt.Println("NO")
		return
	}
	for j := 1; j < 5; j++ {
		for k := j + 1; k < 6; k++ {
			if a[0]+a[j]+a[k] == total/2 {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
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
