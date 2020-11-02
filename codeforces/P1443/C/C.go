package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func solve() {
	t := readInt()
	for i := 0; i < t; i++ {
		solveSingle()
	}
}

func solveSingle() {
	n := readInt()
	a := readInt64Array(n)
	b := readInt64Array(n)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]]
	})
	sum := make([]int64, n+2)
	sum[n] = 0
	for i := n - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + b[id[i]]
	}
	ans := sum[0]
	for i := 0; i < n; i++ {
		ans = min64(ans, max64(sum[i+1], a[id[i]]))
	}
	fmt.Println(ans)
}

func solveSingle2() {
	n := readInt()
	a := readInt64Array(n)
	b := readInt64Array(n)
	lo := int64(0)
	hi := int64(math.MaxInt64)
	ans := hi
	for lo <= hi {
		mi := (lo + hi) / 2
		if ok(a, b, mi) {
			ans = min64(ans, mi)
			hi = mi - 1
		} else {
			lo = mi + 1
		}
	}
	fmt.Println(ans)
}

func ok(a []int64, b []int64, t int64) bool {
	cost := int64(0)
	for i := 0; i < len(a); i++ {
		if a[i] > t {
			cost += b[i]
			if cost > t {
				return false
			}
		}
	}
	return cost <= t
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

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
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
