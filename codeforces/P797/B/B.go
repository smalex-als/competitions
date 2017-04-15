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
	a := make([]int, n)
	for i := 0; i < len(a); i++ {
		a[i] = readInt()
	}
	sort.Ints(a)
	sum := int64(0)
	minPos := int64(-1)
	minNeg := int64(1)
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] > 0 {
			sum += int64(a[i])
			if a[i]%2 != 0 {
				minPos = int64(a[i])
			}
		} else {
			if a[i]%2 != 0 {
				minNeg = int64(a[i])
				break
			}
		}
	}
	if sum%2 == 0 {
		if minNeg != 1 && minPos != -1 {
			if sum-minPos > sum+minNeg {
				fmt.Println(sum - minPos)
			} else {
				fmt.Println(sum + minNeg)
			}
		} else if minNeg != 1 {
			fmt.Println(sum + minNeg)
		} else {
			fmt.Println(sum - minPos)
		}
	} else {
		fmt.Println(sum)
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
