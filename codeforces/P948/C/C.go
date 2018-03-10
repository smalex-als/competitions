package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve() {
	n := readInt()
	v := readIntArray(n)
	t := readIntArray(n)
	a := make([]int64, n)
	acc := int64(0)
	for i := 0; i < n; i++ {
		a[i] = int64(v[i]) + acc
		acc += int64(t[i])
	}
	sort.Sort(Int64(a))
	acc = 0
	start := 0
	cnt := 0
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		total := int64(0)
		prev := acc
		acc += int64(t[i])
		cnt++
		for j := start; j < n; j++ {
			if acc > a[j] {
				total += a[j] - prev
				cnt--
				start = j + 1
			} else {
				break
			}
		}
		total += int64(cnt) * int64(t[i])
		res[i] = total
	}
	w := bufio.NewWriter(os.Stdout)
	for i, v := range res {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, strconv.FormatInt(v, 10))
	}
	w.Flush() // Don't forget to flush!
}

type Int64 []int64

func (a Int64) Len() int           { return len(a) }
func (a Int64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64) Less(i, j int) bool { return a[i] < a[j] }

func abs(a int64) int64 {
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
