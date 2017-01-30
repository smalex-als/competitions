package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	str1 := readString()
	str2 := readString()
	mm := map[rune]rune{}
	for i := 0; i < len(str1); i++ {
		v := rune(str1[i])
		rep := rune(str2[i])
		if prev, ok := mm[v]; !ok {
			mm[v] = rep
		} else if prev != rep {
			fmt.Println(-1)
			return
		}
		if prev, ok := mm[rep]; !ok {
			mm[rep] = v
		} else if prev != v {
			fmt.Println(-1)
			return
		}
	}
	visited := map[rune]int{}
	lines := []string{}
	for k, v := range mm {
		if k != v {
			if visited[k] == 0 && visited[v] == 0 {
				lines = append(lines, fmt.Sprintln(string(v), string(k)))
				visited[k]++
				visited[v]++
			}
		}
	}
	fmt.Println(len(lines))
	for i := 0; i < len(lines); i++ {
		fmt.Print(lines[i])
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
