package main

import (
	"bytes"
	"fmt"
)

func main20() {
	res := losslessDataCompression("abacabadabacaba", 7)
	// res := losslessDataCompression("ababa", 7)
	fmt.Printf("res = %+v\n", res)
}

func losslessDataCompression(str string, width int) string {
	var res bytes.Buffer
	for i := 0; i < len(str); i++ {
		best := 0
		start := i - width
		if start < 0 {
			start = 0
		}
		for j := start; j < i; j++ {
			size := HasPrefix(str[i:], str[j:i])
			if size > best {
				best = size
				start = j
			}
		}
		if best > 0 {
			res.WriteString(fmt.Sprintf("(%d,%d)", start, best))
			i += best - 1
		} else {
			res.WriteString(fmt.Sprintf("%c", str[i]))
		}
	}
	return res.String()
}

func HasPrefix(str, buffer string) int {
	length := len(str)
	if length > len(buffer) {
		length = len(buffer)
	}
	n := 0
	for i := 0; i < length; i++ {
		if str[i] == buffer[i] {
			n++
		} else {
			return n
		}
	}
	return n
}
