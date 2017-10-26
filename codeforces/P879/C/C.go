package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	ops := make([]string, n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		ops[i] = readString()
		a[i] = readInt()
	}
	x1 := calc(ops, a, 1023)
	x2 := calc(ops, a, 0)
	if x1 == 1023 && x2 == 0 {
		fmt.Println("0")
	} else if x1 != 1023 && x2 == 0 {
		fmt.Println("1")
		fmt.Println("&", x1)
	} else if x1 == 1023 && x2 != 0 {
		fmt.Println("1")
		fmt.Println("|", x2)
	} else {
		a1 := 0
		a2 := 0
		for i := uint(0); i < 10; i++ {
			b := 1 << i
			if b&x1 > 0 && b&x2 > 0 {
				a1 |= b
			} else if b&x1 == 0 && b&x2 == 0 {
				a1 |= b
				a2 |= b
			} else if b&x1 == 0 && b&x2 > 0 {
				a2 |= b
			}
		}
		fmt.Println("2")
		fmt.Println("|", a1)
		fmt.Println("^", a2)
	}
}

func calc(ops []string, a []int, x int) int {
	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "|":
			x |= a[i]
		case "^":
			x ^= a[i]
		case "&":
			x &= a[i]
		}
	}
	return x
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

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
