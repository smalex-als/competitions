package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func solve() {
	n := readInt()
	m := readInt()
	k := readInt()
	s := readInt()
	goods := make([][]int, n)
	for i := 0; i < n; i++ {
		value := readInt() - 1
		goods[value] = append(goods[value], i)
	}
	uv := make([][]int, n)
	for i := 0; i < m; i++ {
		u := readInt() - 1
		v := readInt() - 1
		uv[u] = append(uv[u], v)
		uv[v] = append(uv[v], u)
	}
	cost := make([][]int, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, k)
	}
	for j := 0; j < k; j++ {
		queue := make([]int, 0)
		ans := make([]int, n)
		for i := 0; i < n; i++ {
			ans[i] = math.MaxInt32

		}
		for i := 0; i < len(goods[j]); i++ {
			cityId := goods[j][i]
			queue = append(queue, cityId)
			ans[cityId] = 0
		}
		for len(queue) > 0 {
			curCity := queue[0]
			queue = queue[1:]
			for _, nextCity := range uv[curCity] {
				if ans[nextCity] > ans[curCity]+1 {
					queue = append(queue, nextCity)
					ans[nextCity] = ans[curCity] + 1
				}
			}
		}
		for i := 0; i < n; i++ {
			cost[i][j] = ans[i]
		}
	}
	for i := 0; i < len(cost); i++ {
		res := 0
		sort.Ints(cost[i])
		for j := 0; j < s; j++ {
			res += cost[i][j]
		}
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(res)
	}
	fmt.Println()
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
