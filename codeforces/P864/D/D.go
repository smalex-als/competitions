package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	a := make([]int, n)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		a[i] = readInt()
		cnt[a[i]]++
	}
	b := make([]int, 0)
	for i := 1; i <= n; i++ {
		if cnt[i] == 0 {
			b = append(b, i)
		}
	}
	changes := 0
	visit := map[int]int{}
	for i := 0; i < n; i++ {
		if cnt[a[i]] > 1 {
			if visit[a[i]] > 0 || b[0] < a[i] {
				cnt[a[i]]--
				changes++
				a[i] = b[0]
				b = b[1:]
			} else {
				visit[a[i]]++
			}
		}
	}
	fmt.Println(changes)
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		if i > 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString(strconv.Itoa(a[i]))
	}
	fmt.Println(buffer.String())
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
