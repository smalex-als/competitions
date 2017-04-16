package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	p := readInt()
	_ = p
	bal := make([]int, n)
	power := make([]int, n)
	for i := 0; i < n; i++ {
		power[i] = readInt()
		bal[i] = readInt()
	}
	lo := float64(0)
	hi := float64(math.MaxInt32)
	ans := float64(0)
	eternity := true
	for (hi - lo) > 0.000001 {
		mid := (lo + hi) / 2
		if isOk(n, p, bal, power, mid) {
			if mid > ans {
				ans = mid
			}
			lo = mid
		} else {
			eternity = false
			hi = mid
		}
	}
	if eternity {
		fmt.Printf("-1\n")
	} else {
		fmt.Printf("%f\n", ans)
	}
}

func isOk(n int, p int, bal []int, power []int, t float64) bool {
	totalcharge := float64(p) * t
	for i := 0; i < n; i++ {
		consume := float64(power[i])*t - float64(bal[i])
		if consume > 0 {
			totalcharge -= consume
			if totalcharge < 0 {
				return false
			}
		}
	}
	return true
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
