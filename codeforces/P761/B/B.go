package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	L := readInt()
	a := make([]int, n+n+n)
	b := make([]int, n+n+n)
	for i := 0; i < n; i++ {
		a[i] = readInt()
		a[i+n] = a[i] + L
		a[i+n+n] = a[i] + L + L
	}
	for i := 0; i < n; i++ {
		b[i] = readInt()
		b[i+n] = b[i] + L
		b[i+n+n] = b[i] + L + L
	}

	diffa := make([]int, len(a))
	diffb := make([]int, len(b))

	for i := 1; i < len(a); i++ {
		diffa[i] = a[i] - a[i-1]
	}
	for i := 1; i < len(b); i++ {
		diffb[i] = b[i] - b[i-1]
	}
	for y := 0; y < len(b)-n; y++ {
		for x := 0; x < len(a)-n; x++ {
			cnt := 0
			for i := 0; i < n; i++ {
				if diffb[y+i] != diffa[x+i] {
					break
				}
				cnt++
				if cnt == n {
					fmt.Println("YES")
					return
				}
			}
		}
	}
	fmt.Println("NO")
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
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
