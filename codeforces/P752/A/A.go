package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	readInt()
	m := readInt()
	k := readInt() - 1

	column := (k) / (m * 2)
	row := (k % (m * 2)) / 2
	row1 := (k % (m * 2)) % 2
	fmt.Printf("%d ", column+1)
	fmt.Printf("%d ", row+1)
	if row1 == 0 {
		fmt.Println("L")
	} else {
		fmt.Println("R")
	}
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
