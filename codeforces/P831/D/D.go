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
	n := readInt()
	k := readInt()
	p := readInt()
	a := readIntArray(n)
	b := readIntArray(k)
	sort.Ints(a)
	sort.Ints(b)
	mini := math.MaxInt32
	for i := 0; i <= k-n; i++ {
		x := 0
		for j := 0; j < n; j++ {
			x = max(x, abs(a[j]-b[i+j])+abs(b[i+j]-p))
		}
		mini = min(mini, x)
	}
	fmt.Println(mini)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
