package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const MAX = 1e6

func solve() {
	n := readInt()
	x := readInt()
	if n == 2 && x == 0 {
		fmt.Println("NO")
		return
	}
	v := make(map[int]bool)
	res := make([]int, n)
	res[0] = x
	v[x] = true
	start := 0
	for i := 0; i < n-1; i++ {
		for start < MAX {
			if !v[start] {
				need := start ^ res[i]
				if need < MAX && !v[need] {
					res[i] = start
					res[i+1] = need
					v[need] = true
					v[start] = true
					start++
					break
				}
			}
			start++
		}
		if start == MAX {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
	var buffer bytes.Buffer
	for i := 0; i < len(res); i++ {
		if i != 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(strconv.Itoa(res[i]))
	}
	fmt.Println(buffer.String())
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
