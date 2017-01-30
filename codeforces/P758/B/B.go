package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	str1 := readString()
	a := []byte(str1)
	res := make([]int, 4)
	dfs(a, 0, res)
	for i := 0; i < len(res); i++ {
		if i > 0 {
			fmt.Printf(" %d", res[i])
		} else {
			fmt.Printf("%d", res[i])
		}
	}
	fmt.Println()
}

func dfs(a []byte, pos int, res []int) bool {
	if pos == len(a) {
		return true
	}
	if a[pos] != '!' {
		if dfs(a, pos+1, res) {
			return true
		}
	} else {
		for j, v := range "RBYG" {
			ch := byte(v)
			ok := true
			for j := 0; j < 3; j++ {
				k := pos - j - 1
				if k >= 0 && a[k] == ch {
					ok = false
					break
				}
				k = pos + j + 1
				if k < len(a) && a[k] == ch {
					ok = false
					break
				}
			}
			if ok {
				a[pos] = ch
				if dfs(a, pos+1, res) {
					res[j]++
					return true
				}
				a[pos] = '!'
			}
		}
	}
	return false
}

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

func max(a, b int) int {
	if a > b {
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
