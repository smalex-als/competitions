package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt64()
	k := readString()
	values := make([]int, 0)
	for len(k) > 0 {
		var maxValue int
		var maxIndex int
		for i := len(k) - 1; i >= 0; i-- {
			value, err := strconv.ParseInt(k[i:], 10, 64)
			if err == nil {
				if value < n {
					maxValue = int(value)
					maxIndex = i
				} else {
					break
				}
			} else {
				break
			}
		}
		if maxIndex+1 < len(k) {
			if maxValue != 0 {
				for maxIndex < len(k) && k[maxIndex] == '0' {
					maxIndex++
				}
			} else if maxValue == 0 {
				maxIndex = len(k) - 1
			}
		}
		k = k[:maxIndex]
		values = append(values, maxValue)
	}
	var res int64
	for i := len(values) - 1; i >= 0; i-- {
		res *= n
		res += int64(values[i])
	}
	fmt.Println(res)
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
