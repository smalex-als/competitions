package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main1() {
	code := []string{"ans = 0",
		"for i in range(n):",
		"    for j in range(n):",
		"    //DB 3//for j in range(1, n):",
		"    //DB 2//for j in range(n + 1):",
		"        ans += 1",
		"return ans"}
	res := taskMaker(code, 3)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func taskMaker(code []string, challengeId int) []string {
	res := make([]string, 0)
	r := regexp.MustCompile("//DB (\\d+)//")
	for i := 0; i < len(code); i++ {
		m := r.FindStringSubmatch(code[i])
		if len(m) > 0 {
			if value, _ := strconv.Atoi(m[1]); value == challengeId {
				s := code[i]
				res[len(res)-1] = strings.Replace(s, m[0], "", 1)
			}
		} else {
			res = append(res, code[i])
		}
	}
	return res
}
