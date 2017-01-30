package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	people := readInt()
	pillow := readInt()
	k := readInt()
	if people == 1 {
		fmt.Println(pillow)
		return
	}
	pillow -= people
	left := int64(k - 1)
	right := int64(people - k)
	best := 0
	lo := 0
	hi := pillow
	for lo <= hi && pillow > 0 {
		mi := lo + (hi-lo)/2
		rest := int64(pillow - mi)
		step := int64(mi - 1)
		if getSum(right, step)+getSum(left, step) <= rest {
			best = mi
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	fmt.Println(best + 1)
}

func getSum(cnt, step int64) int64 {
	var sum int64
	if cnt > 0 && step > 0 {
		if step < cnt {
			cnt = step
		}
		sum = int64(step-cnt+1+step) * cnt / 2
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
