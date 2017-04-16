package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	str := readString()
	a := []byte(str)
	ans := 0
	for _, v := range []byte("VK") {
		for i := 0; i < len(a); i++ {
			prev := a[i]
			a[i] = v
			res := 0
			for j := 0; j < len(a)-1; j++ {
				if a[j] == 'V' && a[j+1] == 'K' {
					res++
					j++
				}
			}
			ans = max(ans, res)
			a[i] = prev
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
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
