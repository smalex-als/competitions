package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	v := make([]int, 101)
	for i := 0; i < n; i++ {
		v[readInt()]++
	}
	pos := 0
	cnt1 := 0
	cnt2 := 0
	res := make([]int, 0)
	total1 := 0
	total2 := 0
	for i := 0; i < len(v); i++ {
		if v[i] != 0 {
			res = append(res, i)
			if pos%2 == 0 {
				cnt1++
				total1 += v[i]
			} else {
				cnt2++
				total2 += v[i]
			}
			pos++
		}
	}
	if cnt1 == cnt2 && total1 == total2 && len(res) == 2 {
		fmt.Println("YES")
		for i := 0; i < len(res); i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(res[i])
		}
		fmt.Println()
	} else {
		fmt.Println("NO")
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

func min(a int, b int) int {
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
