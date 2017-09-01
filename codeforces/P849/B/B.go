package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var a []int64

func isOnLine(start, end, p int) bool {
	return (a[p]-a[start])*int64(end-start) ==
		(int64(p)-int64(start))*(a[end]-a[start])
}

func isParallel(start0, end0, start1, end1 int) bool {
	return int64(start1-end1)*(a[start0]-a[end0]) ==
		int64(start0-end0)*(a[start1]-a[end1])
}

func checkAll(start0, end0 int) bool {
	start1 := -1
	end1 := -1
	for i := 0; i < len(a); i++ {
		if i != start0 && i != end0 {
			if !isOnLine(start0, end0, i) {
				if start1 == -1 {
					start1 = i
				} else if end1 == -1 {
					end1 = i
					if !isParallel(start0, end0, start1, end1) {
						return false
					}
				} else if !isOnLine(start1, end1, i) {
					return false
				}
			}
		}
	}
	return true
}

func solve() {
	n := readInt()
	a = make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = readInt64()
	}

	fail := true
	for i := 1; i+1 < n; i++ {
		if !isOnLine(0, n-1, i) {
			fail = false
			break
		}
	}
	if fail {
		fmt.Println("No")
		return
	}
	for start := 0; start < 4; start++ {
		for end := start + 1; end < min(start+4, n); end++ {
			if checkAll(start, end) {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
