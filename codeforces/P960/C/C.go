package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	x := readInt()
	_ = readInt64()
	p := int64(0)
	res := make([]int64, 0)
	for i := uint(31); i >= 1; i-- {
		for x >= ((1 << i) - 1) {
			for j := uint(0); j < i; j++ {
				res = append(res, int64(1)+p*100000007)
			}
			p++
			x -= (1 << i) - 1
		}
	}
	fmt.Println(len(res))
	for i := 0; i < len(res); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(res[i])
	}
	fmt.Println()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int64) int64 {
	if a > b {
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
