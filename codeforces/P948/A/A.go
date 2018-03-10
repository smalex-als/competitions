package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve() {
	r := readInt()
	c := readInt()
	res := make([]string, r)
	for i := 0; i < r; i++ {
		res[i] = readString()
	}
	for i := 0; i < r; i++ {
		row := res[i]
		for j := 0; j < len(row); j++ {
			if row[j] == 'W' {
				if j > 0 && row[j-1] == 'S' {
					fmt.Println("No")
					return
				}
				if j+1 < c && row[j+1] == 'S' {
					fmt.Println("No")
					return
				}
				if i > 0 && res[i-1][j] == 'S' {
					fmt.Println("No")
					return
				}
				if i+1 < r && res[i+1][j] == 'S' {
					fmt.Println("No")
					return
				}
			}
		}
		res[i] = row
	}
	fmt.Println("Yes")
	for i := 0; i < r; i++ {
		fmt.Println(strings.Replace(res[i], ".", "D", -1))
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
