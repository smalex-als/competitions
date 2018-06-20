package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solve() {
	s := readString()
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	s = regexp.MustCompile("a+").ReplaceAllString(s, "a")
	s = regexp.MustCompile("b+").ReplaceAllString(s, "b")
	s = regexp.MustCompile("c+").ReplaceAllString(s, "c")
	if len(s) == 3 && s[0] == 'a' && s[1] == 'b' && s[2] == 'c' && (cnt[0] == cnt[2] || cnt[1] == cnt[2]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
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
