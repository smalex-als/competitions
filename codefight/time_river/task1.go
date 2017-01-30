package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main1() {
	res := validTime("13:59")
	fmt.Println(res)
}

func validTime(t string) bool {
	r := regexp.MustCompile("([0-9]{2}):([0-9]{2})")
	p := r.FindStringSubmatch(t)
	h, _ := strconv.ParseInt(p[1], 10, 32)
	m, _ := strconv.ParseInt(p[2], 10, 32)
	return h >= 0 && h < 24 && m >= 0 && m <= 59
}
