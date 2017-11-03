package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var res map[string]bool

func solve() {
	res = make(map[string]bool)
	n := readInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		str := make([]byte, 6)
		for j := 0; j < 6; j++ {
			str[j] = byte(readInt() + '0')
		}
		a[i] = string(str)
	}
	cnt := 0
	perm(a, 0)
	for i := 1; i < 1000; i++ {
		if _, ok := res[strconv.Itoa(i)]; !ok {
			break
		}
		cnt++
	}
	fmt.Println(cnt)
}

func dfs(a []string, pos int, prefix string) {
	if len(prefix) > 0 {
		res[prefix] = true
	}
	if pos == len(a) {
		return
	}
	for i := 0; i < 6; i++ {
		if len(prefix) == 0 && a[pos][i] == '0' {
			continue
		}
		dfs(a, pos+1, prefix+string(a[pos][i]))
	}
}

func perm(a []string, from int) {
	if from+1 == len(a) {
		dfs(a, 0, "")
		return
	}
	for i := from; i < len(a); i++ {
		a[from], a[i] = a[i], a[from]
		perm(a, from+1)
		a[from], a[i] = a[i], a[from]
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int64) int64 {
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
