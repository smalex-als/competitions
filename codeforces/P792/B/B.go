package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	k := readInt()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i + 1
	}
	res := make([]int, 0)
	cur := 1
	for i := 0; i < k; i++ {
		a := readInt()
		index := ((cur + a) % len(p)) - 1
		if index == -1 {
			index = len(p) - 1
		}
		v := p[index]
		res = append(res, v)
		p = p[:index+copy(p[index:], p[index+1:])]
		cur = index + 1
	}
	for i := 0; i < len(res); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(res[i])
	}
	fmt.Println()
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
