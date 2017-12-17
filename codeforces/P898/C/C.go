package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type person struct {
	name   string
	phones []string
}

func solve() {
	n := readInt()
	book := make(map[string]*person)
	for i := 0; i < n; i++ {
		name := readString()
		k := readInt()
		p, ok := book[name]
		if !ok {
			p = &person{name: name}
			book[name] = p
		}
		for j := 0; j < k; j++ {
			p.phones = append(p.phones, readString())
		}
	}
	for _, p := range book {
		sort.Sort(ByLength(p.phones))
		a := make([]string, 0)
		for i := 0; i < len(p.phones); i++ {
			ok := true
			for j := 0; j < len(a); j++ {
				if strings.HasSuffix(a[j], p.phones[i]) {
					ok = false
					break
				}
			}
			if ok {
				a = append(a, p.phones[i])
			}
		}
		p.phones = a
	}
	fmt.Println(len(book))
	names := make([]string, 0)
	for name, _ := range book {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		p := book[name]
		fmt.Println(name, len(p.phones), strings.Join(p.phones, " "))
	}
}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return strings.Compare(s[i], s[j]) < 0
	}
	return len(s[i]) > len(s[j])
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

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
