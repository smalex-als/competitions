package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve() {
	str := readString()
	a := make([]int, len(str))
	nums := make([][]int, 10)
	ok := make([]bool, len(str))
	for k, v := range str {
		num := int(v - '0')
		a[k] = num

		nums[num] = append(nums[num], k)
	}
	vals := make([]string, 0)
	for i := 1; i < 10000; i++ {
		if i%3 == 0 {
			val := strconv.Itoa(i)
			if strings.Contains(val, "0") {
				continue
			}
			vals = append(vals, val)
		}
	}
	sort.Sort(ByLen(vals))
	for ii := 0; ii < len(vals); ii++ {
		cur, _ := strconv.Atoi(vals[ii])
		t := cur
		values := make([]int, 10)
		for t > 0 {
			c := t % 10
			if c > 0 {
				values[c]++
			}
			t /= 10
		}
		cnt := 1000000
		for i := 1; i < len(values); i++ {
			if values[i] > 0 {
				cnt = min(cnt, len(nums[i])/values[i])
			}
		}
		if cnt > 0 {
			for i := 0; i < len(values); i++ {
				for j := 0; j < cnt*values[i]; j++ {
					index := nums[i][len(nums[i])-1]
					ok[index] = true
					nums[i] = nums[i][:len(nums[i])-1]
				}
			}
		}
	}
	var buffer bytes.Buffer
	notEmpty := false
	for i := 0; i < len(str); i++ {
		if ok[i] {
			buffer.WriteString(strconv.Itoa(a[i]))
			notEmpty = true
		} else if a[i] == 0 && notEmpty {
			buffer.WriteString("0")
		}
	}
	if buffer.Len() == 0 && len(nums[0]) > 0 {
		buffer.WriteString("0")
	}
	if buffer.Len() == 0 {
		buffer.WriteString("-1")
	}
	fmt.Println(buffer.String())
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

func max(a, b int) int {
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

type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLen) Less(i, j int) bool { return len(a[i]) > len(a[j]) }
