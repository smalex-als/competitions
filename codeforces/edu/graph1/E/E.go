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
		solveSingle()
	}
}

func solveSingle() {
	n := readInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = readInt()
	}
	cntDeg := 0
	sum := int64(0)
	cntConnected := 0
	for _, deg := range a {
		if deg >= n {
			fmt.Println("NO")
			return
		}
		if deg%2 != 0 {
			cntDeg++
		}
		sum += int64(deg)
		if deg > 0 {
			cntConnected++
		}
	}
	for _, deg := range a {
		if deg > 0 && cntConnected < deg {
			fmt.Println("NO")
			return
		}
	}
	if sum != 2*int64(n)-2 {
		fmt.Println("NO")
		return
	}
	if cntDeg%2 != 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
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
