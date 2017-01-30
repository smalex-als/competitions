package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	readInt()
	path := readString()
	prevdir := 0
	res := 1
	for _, v := range path {
		value := 0
		switch v {
		case 'R':
			value = 1
		case 'L':
			value = 2
		case 'U':
			value = 4
		case 'D':
			value = 8
		}
		prevdir |= value
		if ((prevdir&1) != 0 && (prevdir&2) != 0) ||
			((prevdir&4) != 0 && (prevdir&8) != 0) {
			res++
			prevdir = 0
		}
		prevdir |= value
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
