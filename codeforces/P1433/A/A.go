package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve() {
	a := make([]string, 0)
	for i := 1; i <= 10000; i++ {
		num := strconv.Itoa(i)
		ok := true
		for j := 1; j < len(num); j++ {
			if num[j] != num[j-1] {
				ok = false
				break
			}
		}
		if ok {
			a = append(a, num)
		}
	}
	sort.Strings(a)
	t := readInt()

	for i := 0; i < t; i++ {
		x := strconv.Itoa(readInt())
		ans := 0
		for j := 0; j < len(a); j++ {
			ans += len(a[j])
			if a[j] == x {
				fmt.Println(ans)
				break
			}
		}
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

func min(a int64, b int64) int64 {
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
