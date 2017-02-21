package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve() {
	a := readIntArray(readInt())
	sort.Ints(a)
	mm := map[int]bool{}
	for _, v := range a {
		mm[v] = true
	}
	prefix := make([]int, len(a))
	suffix := make([]int, len(a))
	for i := len(a) - 2; i >= 0; i-- {
		prefix[i] = prefix[i+1]
		if a[i] < a[i+1] {
			prefix[i]++
		}
	}
	for i := 1; i < len(a); i++ {
		suffix[i] = suffix[i-1]
		if a[i] > a[i-1] {
			suffix[i]++
		}
	}
	res := 0
	for i, _ := range a {
		if suffix[i] > 0 && prefix[i] > 0 {
			res++
		}
	}
	fmt.Println(res)
}

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
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
