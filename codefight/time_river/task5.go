package main

import "fmt"
import "strings"
import "strconv"

func main5() {

	//	fmt.Println(newYearCelebrations("23:35", []int{60, 90, 140}))
	fmt.Println(newYearCelebrations("00:00", []int{60, 120, 180, 250}))
}

func newYearCelebrations(takeOffTime string, minutes []int) int {
	const ny = 60 * 24
	time := strings.Split(takeOffTime, ":")
	num := parseValue(time[0])*60 + parseValue(time[1])
	if num == 0 {
		num = ny
	}
	cnt := 0
	sum := 0
	for i := 0; i < len(minutes); i++ {
		m := minutes[i] - sum
		sum += m
		nnum := num + m
		fmt.Println(num, nnum)
		if num <= ny && ny <= nnum {
			cnt++
		}
		num = nnum
		num -= 60
	}
	if num <= ny {
		cnt++
	}
	return cnt
}

func parseValue(s string) int {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int(v)
}
