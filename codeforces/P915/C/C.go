package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve() {
	a := []byte(strconv.FormatInt(readInt64(), 10))
	b := []byte(strconv.FormatInt(readInt64(), 10))
	anum := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		anum[i] = int(a[i] - '0')
	}
	bnum := make([]int, len(b))
	for i := 0; i < len(b); i++ {
		bnum[i] = int(b[i] - '0')
	}
	sort.Sort(sort.Reverse(sort.IntSlice(anum)))
	if len(a) == len(b) {
		dfs(anum, bnum, 0)
	}
	for _, v := range anum {
		fmt.Print(v)
	}
	fmt.Println()
}

func dfs(anum []int, bnum []int, pos int) bool {
	if pos == len(anum) {
		return true
	}
	prev := -1
	for i := pos; i < len(anum); i++ {
		if anum[i] < bnum[pos] {
			anum[pos], anum[i] = anum[i], anum[pos]
			resort(anum, pos+1)
			return true
		}
		if anum[i] == bnum[pos] && prev != anum[i] {
			anum[pos], anum[i] = anum[i], anum[pos]
			resort(anum, pos+1)
			if dfs(anum, bnum, pos+1) {
				return true
			}
			resort(anum, pos)
			prev = anum[i]
		}
	}
	return false
}

func resort(nums []int, pos int) {
	for i := pos; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
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
