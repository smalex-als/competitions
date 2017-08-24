package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	s := readString()
	k := readInt()
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	if len(m) >= k {
		fmt.Println(0)
		return
	}
	cur := len(m)
	res := 0
	for _, cnt := range m {
		for i := 1; i < cnt; i++ {
			cur++
			res++
			if cur >= k {
				fmt.Println(res)
				return
			}
			if cur > 26 {
				fmt.Println("impossible")
				return
			}
		}
	}
	fmt.Println("impossible")
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
