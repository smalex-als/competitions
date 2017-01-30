package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main6() {
	// messages := [][]string{{"Jkl ABA ty", "111"},
	// 	{"Jkl aba TY", "222"},
	// 	{"Jkl abA Ty", "111"},
	// 	{"jkl Aba tY", "111"},
	// 	{"JKl ABA ty", "222"},
	// 	{"Jkl aba tY", "222"},
	// 	{"Jkl abA Ty", "111"},
	// 	{"jkl Aba tY klk TY", "111"},
	// 	{"Jkl aBa tY", "222"},
	// 	{"Jkl aBA Ty", "111"},
	// 	{"Jkl aBA TY", "111"}}
	// spamSignals := []string{"ty", "jk", "ab"}
	messages := [][]string{{"Sale today!", "2837273"},
		{"Unique offer!", "3873827"},
		{"Only today and only for you!", "2837273"},
		{"Sale today!", "2837273"},
		{"Unique offer!", "3873827"}}
	spamSignals := []string{"sale", "discount", "offer"}

	res := spamDetection(messages, spamSignals)
	fmt.Println(res)
}

func spamDetection(messages [][]string, spamSignals []string) []string {
	mapSpamWords := make(map[string]bool)
	for _, v := range spamSignals {
		mapSpamWords[v] = true
	}
	mapVisitSpamWords := make(map[string]bool)
	messageHasSpamWords := make(map[int]bool)

	res := make([]string, 4)

	for i := 0; i < len(res); i++ {
		res[i] = "passed"
	}
	r := regexp.MustCompile("[^a-z]")
	r2 := regexp.MustCompile("\\s{2,}")
	error1 := 0
	mapMsg := make(map[string]int)
	mapUserId := make(map[string]string)
	mapAllMsg := make(map[string]int)
	mapCntByUser := make(map[string]int)
	error4 := 0
	for i := 0; i < len(messages); i++ {
		str := r.ReplaceAllString(strings.ToLower(messages[i][0]), " ")
		str = r2.ReplaceAllString(strings.TrimSpace(str), " ")
		words := strings.Split(str, " ")
		cnt := len(words)
		if cnt < 5 {
			error1++
		}
		key := messages[i][1] + " {TO} " + messages[i][0]
		mapMsg[key]++
		mapUserId[key] = messages[i][1]
		mapAllMsg[messages[i][0]]++
		mapCntByUser[messages[i][1]]++

		for _, v := range words {
			if mapSpamWords[v] {
				mapVisitSpamWords[v] = true
				if !messageHasSpamWords[i] {
					messageHasSpamWords[i] = true
					error4++
				}
			}
		}
	}
	perror1 := error1 * 100 / len(messages)
	if perror1 >= 90 {
		n := len(messages)
		for i := 2; i <= n && i <= error1; i++ {
			if n%i == 0 && error1%i == 0 {
				n /= i
				error1 /= i
				i--
			}
		}
		res[0] = fmt.Sprintf("failed: %d/%d", error1, n)
	}
	failedUsers := make([]int, 0)
	for i, k := range mapMsg {
		pct := 100 * k / mapCntByUser[mapUserId[i]]
		if k >= 2 && pct > 50 {
			v, _ := strconv.Atoi(mapUserId[i])
			failedUsers = append(failedUsers, v)
		}
	}
	if len(failedUsers) > 0 {
		sort.Ints(failedUsers)
		failedUsersStr := make([]string, len(failedUsers))
		for i := 0; i < len(failedUsers); i++ {
			failedUsersStr[i] = strconv.Itoa(failedUsers[i])
		}
		res[1] = fmt.Sprintf("failed: %s", strings.Join(failedUsersStr, " "))
	}
	if len(messages) >= 2 {
		for k, v := range mapAllMsg {
			pct := v * 100 / len(messages)
			if pct >= 50 {
				res[2] = fmt.Sprintf("failed: %s", k)
				break
			}
		}
	}
	if error4*100/len(messages) >= 50 {
		words := make([]string, 0)
		for k, _ := range mapVisitSpamWords {
			words = append(words, k)
		}
		sort.Strings(words)
		res[3] = "failed: " + strings.Join(words, " ")
	}

	return res
}
