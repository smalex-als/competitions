package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	a := readString()
	b := []byte(a)
	N := len(a)
	suffix := make([]byte, N)
	suffix[N-1] = b[N-1]
	for i := N - 2; i >= 0; i-- {
		if b[i] < suffix[i+1] {
			suffix[i] = b[i]
		} else {
			suffix[i] = suffix[i+1]
		}
	}
	res := make([]byte, 0)
	buf := make([]byte, 0)
	for i := 0; i < N; i++ {
		if b[i] <= suffix[i] {
			for len(buf) > 0 && (buf[len(buf)-1] <= b[i]) {
				res = append(res, buf[len(buf)-1])
				buf = buf[:len(buf)-1]
			}
			res = append(res, b[i])
		} else {
			for len(buf) > 0 && (buf[len(buf)-1] < b[i] && buf[len(buf)-1] <= suffix[i]) {
				res = append(res, buf[len(buf)-1])
				buf = buf[:len(buf)-1]
			}
			buf = append(buf, b[i])
		}
	}
	for i := 0; i < len(buf); i++ {
		res = append(res, buf[len(buf)-i-1])
	}
	fmt.Println(string(res))
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

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLen) Less(i, j int) bool { return len(a[i]) > len(a[j]) }
