package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	t := readInt()

	for j := 0; j < t; j++ {
		n := readInt()
		a := readInt64Array(n)
		ans := -1
		best := a[0]
		for i := 0; i < n; i++ {
			best = max(best, a[i])
		}
		for i := 0; i < n; i++ {
			if a[i] == best {
				if solveCase(a, i) {
					ans = i + 1
					break
				}
			}
		}
		fmt.Println(ans)
	}
}

func solveCase(a []int64, x int) bool {
	l := x - 1
	r := x + 1
	cur := a[x]
	for {
		ok := false
		for l >= 0 && a[l] < cur {
			cur++
			l--
			ok = true
		}
		for r < len(a) && a[r] < cur {
			cur++
			r++
			ok = true
		}
		if !ok {
			break
		}
	}
	return l == -1 && r == len(a)
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

func max(a, b int64) int64 {
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
