package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main2() {
	fmt.Println(videoPart("2:20:00", "7:00:00"))
	fmt.Println(videoPart("0:02:20", "0:10:00"))
}

func videoPart(part string, total string) []int {
	a := parseInt(part)
	b := parseInt(total)
	g := gcd(a, b)
	return []int{a / g, b / g}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func parseInt(str string) int {
	r := regexp.MustCompile("([0-9]+):([0-9]+):([0-9]+)")
	p := r.FindStringSubmatch(str)
	h, _ := strconv.ParseInt(p[1], 10, 32)
	m, _ := strconv.ParseInt(p[2], 10, 32)
	s, _ := strconv.ParseInt(p[3], 10, 32)
	return int(h)*3600 + int(m)*60 + int(s)
}
