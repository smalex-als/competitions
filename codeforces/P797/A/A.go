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
	a := make([]int, k)
	if k == 1 {
		a[0] = n
		k = 0
	}
	for i := 2; i <= n && k > 1; i++ {
		for n%i == 0 {
			k--
			n /= i
			a[k] = i
			if k == 1 {
				if n != 1 {
					a[0] = n
					k--
				}
				break
			}
		}
	}
	if k == 0 {
		for i := 0; i < len(a); i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(a[i])
		}
	} else {
		fmt.Println(-1)
	}
}

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
