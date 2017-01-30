package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	m := readInt()
	k := readInt64()
	x := readInt()
	y := readInt()

	v := make([][]int64, n)
	for i := 0; i < n; i++ {
		v[i] = make([]int64, m)
	}
	if n == 1 {
		size := m
		mul := k / int64(size)
		k = k % int64(size)
		if mul > 0 {
			for j := 0; j < m; j++ {
				v[0][j] = mul
			}
		}
	} else {
		size := (n + n - 2) * m
		mul := k / int64(size)
		k = k % int64(size)
		if mul > 0 {
			for i := 1; i < n-1; i++ {
				for j := 0; j < m; j++ {
					v[i][j] = 2 * mul
				}
			}
			for j := 0; j < m; j++ {
				v[0][j] = mul
			}
			for j := 0; j < m; j++ {
				v[n-1][j] = mul
			}
		}
	}
	for i := 0; i < n && k > 0; i++ {
		for j := 0; j < m && k > 0; j++ {
			v[i][j]++
			k--
		}
	}
	for i := n - 2; i >= 0 && k > 0; i-- {
		for j := 0; j < m && k > 0; j++ {
			v[i][j]++
			k--
		}
	}
	// for i := 0; i < n; i++ {
	// 	fmt.Printf("v[i] = %+v\n", v[i])
	// }
	minValue := int64(math.MaxInt64)
	maxValue := int64(math.MinInt64)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			value := v[i][j]
			if value < minValue {
				minValue = value
			}
			if value > maxValue {
				maxValue = value
			}
		}
	}
	fmt.Printf("%d %d %d", maxValue, minValue, v[x-1][y-1])
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
