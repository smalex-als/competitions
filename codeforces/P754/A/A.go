package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	a := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		a[i] = readInt()
		if a[i] == 0 {
			total++
		}
	}
	if total == n {
		fmt.Println("NO")
		return
	}
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		start := i
		sum := a[i]
		for sum == 0 && i+1 < n {
			i++
			sum += a[i]
		}
		for i+1 < n && a[i+1] == 0 {
			i++
			sum += a[i]
		}
		if sum == 0 {
			fmt.Println("NO")
			return
		}
		res = append(res, []int{start + 1, i + 1})
	}
	fmt.Println("YES")
	fmt.Println(len(res))
	for _, v := range res {
		fmt.Println(v[0], v[1])
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
