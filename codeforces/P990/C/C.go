package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Parentheses struct {
	sum    int
	lowest int
	cnt    int
}

func NewParentheses(str string) *Parentheses {
	sum, lowest := getSum(str)
	par := &Parentheses{}
	par.sum = sum
	par.lowest = lowest
	par.cnt = 1
	return par
}

func solve() {
	n := readInt()
	p := make(map[string]*Parentheses)
	for i := 0; i < n; i++ {
		str := readString()
		if prev, exists := p[str]; exists {
			prev.cnt++
		} else {
			p[str] = NewParentheses(str)
		}
	}
	ans := int64(0)
	for _, p1 := range p {
		if p1.lowest >= 0 {
			for _, p2 := range p {
				if p1.sum+p2.lowest >= 0 && p1.sum+p2.sum == 0 {
					ans += int64(p1.cnt) * int64(p2.cnt)
				}
			}
		}
	}
	fmt.Println(ans)
}

func getSum(str string) (int, int) {
	sum := 0
	lowest := 0
	for _, ch := range str {
		if ch == '(' {
			sum++
		} else {
			sum--
		}
		lowest = min(lowest, sum)
	}
	return sum, lowest
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
