package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var successuful int64

func solve() {
	successuful = 0
	n := readInt()
	a := make([]int, n)
	b := make([]int, n)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		a[i] = readInt()
		b[i] = a[i]
		cnt[a[i]]++
	}
	sort.Ints(b)
	// all uniq
	if b[0] != b[1] && b[2] != b[1] && b[2] != b[0] {
		res := 0
		for i := 0; i < 3; i++ {
			if cnt[b[i]] > res {
				res = cnt[b[i]]
			}
		}
		fmt.Println(res)
		return
	}
	// all the same
	if b[0] == b[1] && b[1] == b[2] {
		v := cnt[b[0]]
		fmt.Println(int64(v) * int64(v-1) * int64(v-2) / int64(6))
		return
	}
	t := 0
	k := 0
	if b[0] == b[1] {
		t = b[0]
		k = b[2]
	} else {
		k = b[0]
		t = b[2]
	}
	v := k
	left := 0
	right := 0
	found := false
	for i := 0; i < n; i++ {
		if t == a[i] {
			if found {
				right++
			} else {
				left++
			}
		} else if v == a[i] {
			found = true
		}
	}
	successuful += int64(right * left)
	successuful += int64(right) * int64(right-1) / int64(2)
	successuful += int64(left) * int64(left-1) / int64(2)
	fmt.Println(successuful)
}

func min(a, b int) int {
	if a < b {
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
