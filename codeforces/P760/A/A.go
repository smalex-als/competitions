package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	daysInMonths := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	m := readInt()
	d := readInt()
	days := daysInMonths[m-1] + d - 1
	weeks := days / 7
	if days%7 != 0 {
		weeks++
	}
	fmt.Println(weeks)
}

func getStats(str string) map[byte]int {
	stat := map[byte]int{}
	for _, v := range str {
		stat[byte(v)]++
	}
	return stat
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
