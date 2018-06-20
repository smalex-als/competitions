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
	obj := map[string]string{
		"purple": "Power",
		"green":  "Time",
		"blue":   "Space",
		"orange": "Soul",
		"red":    "Reality",
		"yellow": "Mind",
	}
	values := make(map[string]int, 0)
	for i := 0; i < n; i++ {
		values[obj[readString()]]++
	}
	res := make([]string, 0)
	for _, v := range obj {
		if _, ok := values[v]; ok {
		} else {
			res = append(res, v)
		}
	}
	fmt.Println(len(res))
	sort.Strings(res)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
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
