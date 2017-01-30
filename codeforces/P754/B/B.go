package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	field := make([][]int, 4)
	for i := 0; i < 4; i++ {
		field[i] = make([]int, 4)
		for j, ch := range readString() {
			switch ch {
			case 'x':
				field[i][j] = 1
			case 'o':
				field[i][j] = 2
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if field[i][j] == 0 {
				field[i][j] = 1
				if win(field) {
					fmt.Println("YES")
					return
				}
				field[i][j] = 0
			}
		}
	}
	fmt.Println("NO")
}

func win(field [][]int) bool {
	for j := 0; j < 4; j++ {
		for i := 0; i < 2; i++ {
			if field[j][i] == 1 && field[j][i+1] == 1 && field[j][i+2] == 1 {
				return true
			}
		}
	}
	for j := 0; j < 4; j++ {
		for i := 0; i < 2; i++ {
			if field[i][j] == 1 && field[i+1][j] == 1 && field[i+2][j] == 1 {
				return true
			}
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if field[i][j] == 1 && field[i+1][j+1] == 1 && field[i+2][j+2] == 1 {
				return true
			}
		}
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if field[i][3-j] == 1 && field[i+1][3-j-1] == 1 && field[i+2][3-j-2] == 1 {
				return true
			}
		}
	}
	return false
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
