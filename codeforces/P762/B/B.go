package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Mouse struct {
	Price int64
	IsUsb bool
}

func solve() {
	a := readInt()
	b := readInt()
	c := readInt()
	m := readInt()
	items := make([]Mouse, m)
	for i := 0; i < m; i++ {
		items[i] = Mouse{readInt64(), parseType(readString())}
	}
	sort.Sort(ByPrice(items))
	total := 0
	sum := int64(0)
	for i := 0; i < m; i++ {
		item := items[i]
		if item.IsUsb {
			if a > 0 {
				a--
				total++
				sum += item.Price
			} else if c > 0 {
				c--
				total++
				sum += item.Price
			}
		} else {
			if b > 0 {
				b--
				total++
				sum += item.Price
			} else if c > 0 {
				c--
				total++
				sum += item.Price
			}
		}
	}
	fmt.Printf("%d %d\n", total, sum)
}

type ByPrice []Mouse

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }

func parseType(name string) bool {
	if name == "USB" {
		return true
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
