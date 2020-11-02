package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	t := readInt()
	for i := 0; i < t; i++ {
		solveSingle()
	}
}

func solveSingle() {
	a := readInt()
	b := readInt()
	s := readString()
	values := make([][]int, 0)
	start := 0
	for start < len(s) {
		cnt := CountValues(s[start], s, start)
		if cnt > 0 {
			values = append(values, []int{int(s[start] - '0'), cnt})
		} else {
			break
		}
		start += cnt
	}
	cost := 0
	for i := 0; i < len(values)-2; i++ {
		if values[i][0] == 1 {
			cost1 := 2 * a
			cost2 := a + b*values[i+1][1]
			if cost1 > cost2 {
				cost += b * values[i+1][1]
				values[i+2][1] += values[i][1]
				values[i][0] = 0
			}
		}
	}
	for _, value := range values {
		if value[0] == 1 {
			cost += a
		}
	}
	fmt.Println(cost)
}

func CountValues(ex uint8, str string, start int) int {
	cnt := 0
	for i := start; i < len(str) && str[i] == ex; i++ {
		cnt++
	}
	return cnt
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
