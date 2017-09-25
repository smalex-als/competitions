package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX = 1e6

func solve() {
	a := readInt()
	b := readInt()
	f := readInt()
	k := readInt()
	gas := b
	res := 0
	for i := 0; i < k; i++ {
		if i%2 == 0 {
			gas -= f
			if gas < 0 {
				fmt.Println(-1)
				return
			}
			if i == k-1 {
				if gas < (a - f) {
					res++
					gas = b
				}
			} else {
				if gas < (a-f)*2 {
					res++
					gas = b
				}
			}
			gas -= (a - f)
			if gas < 0 {
				fmt.Println(-1)
				return
			}
		} else {
			gas -= a - f
			if gas < 0 {
				fmt.Println(-1)
				return
			}
			if i == k-1 {
				if gas < f {
					res++
					gas = b
				}
			} else {
				if gas < f*2 {
					res++
					gas = b
				}
			}
			gas -= f
			if gas < 0 {
				fmt.Println(-1)
				return
			}
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
