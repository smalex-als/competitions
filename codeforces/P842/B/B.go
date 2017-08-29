package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type circle struct {
	x, y, r int
}

func solve() {
	r := readInt()
	d := readInt()
	n := readInt()
	circles := make([]circle, n)
	res := 0
	up := float64(r)
	down := float64(r - d)
	for i := 0; i < n; i++ {
		circles[i] = circle{readInt(), readInt(), readInt()}
		c := circles[i]
		l := math.Sqrt(float64(c.x*c.x + c.y*c.y))
		r0 := l - float64(c.r)
		r1 := l + float64(c.r)
		if up >= r1 && down <= r0 {
			res++
		}
	}

	fmt.Println(res)
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
