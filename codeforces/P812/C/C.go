package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve() {
	n := readInt()
	s := readInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readInt()
	}

	lo := 1
	hi := n
	ans := 0
	T := 0
	for lo <= hi {
		k := (lo + hi) / 2
		t := isOk(k, a, s)
		if t > 0 {
			lo = k + 1
			if ans < k {
				ans = k
				T = t
			}
		} else {
			hi = k - 1
		}
	}
	fmt.Println(ans, T)
}

func isOk(k int, a []int, s int) int {
	b := make([]int, 0)
	maxSum := int64(s)
	maxK := int64(k)
	for i := 0; i < len(a); i++ {
		v := int64(a[i]) + maxK*int64(i+1)
		if v <= maxSum {
			b = append(b, int(v))
		}
	}
	if len(b) < k {
		return 0
	}
	sort.Ints(b)
	sum := 0
	for i := 0; i < k; i++ {
		sum += b[i]
		if sum > s {
			return 0
		}
	}
	return sum
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
