package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve() {
	const a = "^>v<"
	const b = "^<v>"
	from := readString()
	to := readString()
	cnt := readInt() % 4
	res1 := (strings.Index(a, to) - strings.Index(a, from) + 4) % 4
	res2 := (strings.Index(b, to) - strings.Index(b, from) + 4) % 4
	if res1 == cnt && res2 == cnt {
		fmt.Println("undefined")
	} else if res1 == cnt {
		fmt.Println("cw")
	} else if res2 == cnt {
		fmt.Println("ccw")
	} else {
		fmt.Println("undefined")
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
