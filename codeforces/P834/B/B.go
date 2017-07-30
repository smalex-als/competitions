package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	readInt()
	k := readInt()
	q := readString()
	openAt := make([]bool, 26)
	closedAt := make([]int, 26)
	for i := 0; i < len(closedAt); i++ {
		closedAt[i] = -1
	}
	for i := 0; i < len(q); i++ {
		closedAt[q[i]-'A'] = i
	}

	for i := 0; i < len(q); i++ {
		v := q[i] - 'A'
		if !openAt[v] {
			k--
			if k < 0 {
				fmt.Println("YES")
				return
			}
			openAt[v] = true
		}
		if closedAt[v] == i {
			k++
		}
	}
	fmt.Println("NO")
}

func min(a, b int) int {
	if a < b {
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
