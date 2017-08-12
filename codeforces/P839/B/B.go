package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	k := readInt()
	cnt := make([]int, 5)
	have := make([]int, 5)
	have[4] = n
	have[2] = 2 * n
	for i := 0; i < k; i++ {
		x := readInt()
		for x >= 3 {
			if have[4] > 0 {
				have[4]--
				x -= 4
			} else if have[2] > 0 {
				have[2]--
				x -= 2
			} else {
				fmt.Println("NO")
				return
			}
		}
		if x > 0 {
			cnt[x]++
		}
	}
	for cnt[2] > 0 {
		if have[2] > 0 {
			have[2]--
			cnt[2]--
		} else if have[4] > 0 {
			cnt[2]--
			have[4]--
			have[1]++
		} else {
			cnt[2]--
			cnt[1] += 2
		}
	}
	if cnt[1] <= have[1]+have[2]+have[4]*2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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
