package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	m := readInt()
	words := map[string]int{}
	for i := 0; i < n; i++ {
		words[readString()]++
	}
	common := 0
	for i := 0; i < m; i++ {
		if _, ok := words[readString()]; ok {
			common++
		}
	}
	for {
		if common > 0 {
			common--
			n--
			m--
		} else if n > 0 {
			n--
		} else {
			fmt.Println("NO")
			return
		}
		if common > 0 {
			common--
			n--
			m--
		} else if m > 0 {
			m--
		} else {
			fmt.Println("YES")
			return
		}
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
