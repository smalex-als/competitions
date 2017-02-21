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
	k := readInt()
	x := readInt()
	a := readIntArray(n)
	prev := make([]int, n)
	prev2 := make([]int, n)
	cnt := 0
	for i := 0; i < k; i++ {
		sort.Ints(a)
		if i%2 == 1 && testEq(a, prev) {
			cnt++
		} else if i%2 == 0 && testEq(a, prev2) {
			cnt++
		} else {
			cnt = 0
		}
		if i%2 == 1 {
			copy(prev, a)
		} else {
			copy(prev2, a)
		}
		for j := 0; j < len(a); j += 2 {
			a[j] = a[j] ^ x
		}
		if cnt > 2 && (i+1)%2 == k%2 {
			break
		}
	}
	sort.Ints(a)
	fmt.Printf("%d %d\n", a[len(a)-1], a[0])
}

func testEq(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
