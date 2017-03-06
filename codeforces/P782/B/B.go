package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var y []int
var x []int

func solve() {
	n := readInt()
	x = make([]int, n)
	y = make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = readInt()
	}
	for i := 0; i < n; i++ {
		y[i] = readInt()
	}
	left := float64(0)
	right := float64(math.MaxInt32)
	middle := float64(0)
	for i := 0; i < 100; i++ {
		middle1 := left + (right-left)/3
		middle2 := left + 2*(right-left)/3
		middle = (left + right) / 2
		if f(middle1) < f(middle2) {
			right = middle2
		} else {
			left = middle1
		}
	}
	fmt.Printf("%.10f\n", f(middle))
}

func f(t float64) float64 {
	var max float64
	for i := 0; i < len(x); i++ {
		s := (math.Abs(t-float64(x[i])) / float64(y[i]))
		if s > max {
			max = s
		}
	}
	return max
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
