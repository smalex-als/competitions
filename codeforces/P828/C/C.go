package main

import (
	"bufio"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	res := make([]byte, 0)
	for i := 0; i < n; i++ {
		s := readString()
		t := readInt()
		end := 0
		for j := 0; j < t; j++ {
			pos := readInt() - 1
			for ii := max(end-pos, 0); ii < len(s); ii++ {
				for len(res) <= pos+ii {
					res = append(res, 'a')
				}
				res[pos+ii] = s[ii]
			}
			end = max(end, pos+len(s))
		}
	}
	os.Stdout.Write(res)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
