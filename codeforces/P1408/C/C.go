package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	t := readInt()

	for j := 0; j < t; j++ {
		n := readInt()
		l := readInt()
		a := make([]int, 0)
		a = append(a, 0)
		for i := 0; i < n; i++ {
			a = append(a, readInt())
		}
		a = append(a, l)

		leftSpeed := 1
		rightSpeed := 1

		leftIndex := 0
		rightIndex := len(a) - 2
		leftAt := float64(0)
		rightAt := float64(0)

		for leftIndex < rightIndex {
			tmpLeftAt := leftAt + float64(a[leftIndex+1]-a[leftIndex])/float64(leftSpeed)
			tmpRightAt := rightAt + float64(a[rightIndex+1]-a[rightIndex])/float64(rightSpeed)
			if tmpLeftAt < tmpRightAt {
				leftIndex++
				leftSpeed++
				leftAt = tmpLeftAt
			} else {
				rightIndex--
				rightSpeed++
				rightAt = tmpRightAt
			}
		}
		curTime := max(leftAt, rightAt)
		dist := float64(a[leftIndex+1] - a[leftIndex])
		if leftAt < rightAt {
			dist -= float64(leftSpeed) * (rightAt - leftAt)
		} else {
			dist -= float64(rightSpeed) * (leftAt - rightAt)
		}
		curTime += dist / float64(rightSpeed+leftSpeed)
		fmt.Printf("%.8f\n", curTime)
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

func max(a, b float64) float64 {
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
