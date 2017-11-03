package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	var data [24]int
	for i := 0; i < 24; i++ {
		data[i] = readInt()
	}
	if isFullSide(data, 3) && isFullSide(data, 4) {
		row0 := []int{1, 3, 5, 7, 9, 11, 20, 22}
		row1 := []int{0, 2, 4, 6, 8, 10, 23, 21}

		if check(data, row0, row1, 2) {
			fmt.Println("YES")
			return
		}
		if check(data, row0, row1, -2) {
			fmt.Println("YES")
			return
		}
	}
	if isFullSide(data, 0) && isFullSide(data, 2) {
		row0 := []int{12, 13, 4, 5, 16, 17, 20, 21}
		row1 := []int{14, 15, 6, 7, 18, 19, 22, 23}

		if check(data, row0, row1, 2) {
			fmt.Println("YES")
			return
		}
		if check(data, row0, row1, -2) {
			fmt.Println("YES")
			return
		}
	}
	if isFullSide(data, 1) && isFullSide(data, 5) {
		row0 := []int{8, 9, 18, 16, 3, 2, 13, 15}
		row1 := []int{10, 11, 19, 17, 1, 0, 12, 14}

		if check(data, row0, row1, 2) {
			fmt.Println("YES")
			return
		}
		if check(data, row0, row1, -2) {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}

func check(data [24]int, row0, row1 []int, dir int) bool {
	for i := 0; i < len(row0); i++ {
		if data[row0[i]] != data[row1[(i+dir+len(row1))%len(row1)]] {
			return false
		}
	}
	return true
}

func isFullSide(data [24]int, side int) bool {
	color := data[side*4]
	for i := 0; i < 4; i++ {
		if data[side*4+i] != color {
			return false
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
