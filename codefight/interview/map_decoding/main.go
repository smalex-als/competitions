package main

import (
	"fmt"
)

func main() {
	// fmt.Println(mapDecoding("SMALEX"))
	fmt.Println(mapDecoding("1221112111122221211221221212212212111221222212122221222112122212121212221212122221211112212212211211"))
}

func mapDecoding(message string) int {
	if message == "0" {
		return 0
	}
	memo = map[int]int64{}
	a := make([]int, len(message))
	for k, v := range message {
		a[k] = int(v - '0')
	}
	res := dfs(a, 0)
	return int(res)
}

var memo map[int]int64

func dfs(a []int, pos int) int64 {
	if value, ok := memo[pos]; ok {
		return value
	}
	if pos == len(a) {
		return 1
	}
	num := 0
	cnt := int64(0)
	for i := 0; i < 2; i++ {
		if pos+i < len(a) {
			num += a[pos+i]
			if i == 1 && num < 10 {
				break
			}
			if num > 0 && num <= 26 {
				cnt += dfs(a, pos+i+1) % 1000000007
			}
			num *= 10
		}
	}
	memo[pos] = cnt
	return cnt
}
