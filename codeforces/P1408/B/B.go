package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	t := readInt()
	for i := 0; i < t; i++ {
		n := readInt()
		k := readInt()
		a := readIntArray(n)
		freq := make(map[int]int, 0)
		for j := 0; j < n; j++ {
			freq[a[j]]++
		}
		if len(freq) <= k {
			fmt.Println(1)
		} else {
			if k == 1 {
				if len(freq) == 1 {
					fmt.Println(1)
				} else {
					fmt.Println(-1)
				}
			} else {
				res := solveSingle(a, k)
				fmt.Println(res)
			}
		}
	}
}

func solveSingle(a []int, k int) int {
	freq := make(map[int]int, 0)
	pos := 0
	max := 0
	for i := 0; i < len(a); i++ {
		freq[a[i]]++
		max = a[i]
		a[i] = 0
		pos = i
		if len(freq) == k {
			break
		}
	}
	for i := pos + 1; i < len(a); i++ {
		a[i] -= max
	}
	ok := true
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			ok = false
			break
		}
	}
	if ok {
		return 1
	} else {
		return 1 + solveSingle(a, k)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int64) int64 {
	if a > b {
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
