package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	readInt()
	readInt()
	n := []rune(readString())
	m := []rune(readString())
	best := 10000
	bestIndex := 0
	for i := 0; i < len(m); i++ {
		cnt := 0
		if i+len(n) <= len(m) {
			for j := 0; j < len(n); j++ {
				if m[i+j] != n[j] {
					cnt++
				}
			}
			if cnt < best {
				best = cnt
				bestIndex = i
			}
		}
	}
	fmt.Println(best)
	cnt := 0
	for j := 0; j < len(n); j++ {
		if m[bestIndex+j] != n[j] {
			if cnt > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", j+1)
			cnt++
		}
	}
	fmt.Println()
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
