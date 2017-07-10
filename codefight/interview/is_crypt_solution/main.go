package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	// crypt := []string{"SEND", "MORE", "MONEY"}
	// solution := [][]string{{"O", "0"}, {"M", "1"}, {"Y", "2"}, {"E", "5"}, {"N", "6"}, {"D", "7"}, {"R", "8"}, {"S", "9"}}
	crypt := []string{"A", "A", "A"}
	solution := [][]string{{"A", "0"}}
	fmt.Println(isCryptSolution(crypt, solution))
}

func isCryptSolution(crypt []string, solution [][]string) bool {
	m := make(map[rune]rune)
	for _, sol := range solution {
		if _, ok := m[rune(sol[0][0])]; ok {
			return false
		} else {
			m[rune(sol[0][0])] = rune(sol[1][0])
		}
	}
	nums := make([]int64, 3)
	for i, w := range crypt {
		val := decrypt(w, m)
		if val == "" {
			return false
		}
		if len(val) > 1 && strings.HasPrefix(val, "0") {
			return false
		}
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return false
		}
		nums[i] = num
	}
	if nums[0]+nums[1] != nums[2] {
		return false
	}
	return true
}

func decrypt(w string, m map[rune]rune) string {
	var buffer bytes.Buffer
	for _, v := range []rune(w) {
		if val, ok := m[v]; ok {
			buffer.WriteRune(val)
		} else {
			return ""
		}
	}
	return buffer.String()
}
