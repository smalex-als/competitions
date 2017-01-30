package main

import "fmt"

func main13() {
	// res := typosquatting(7, "godaddy.com")
	// res := typosquatting(9, "smalex.tv")
	res := typosquatting(9, "omg.tv")
	// res := typosquatting(1, "a.ab")
	fmt.Printf("res = %+v\n", res)
	// res := typosquatting(1, "godaddy.com")
	// res2 := typosquatting(0, "a.ab")
	// fmt.Printf("res = %+v\n", res2)
}

func typosquatting(n int, domain string) int {
	res := make(map[string]bool)
	res[domain] = true
	queue := make([]string, 1)
	queueLevel := make([]int, 1)
	queue[0] = domain
	queueLevel[0] = 0
	pos := 0
	ans := -1
	for pos < len(queue) {
		cur := queue[pos]
		level := queueLevel[pos]
		if level > ans+1 {
			ans = level - 1
		}
		if pos-1 == n {
			break
		}
		pos++
		for _, v := range traceWords(cur) {
			if !res[v] {
				res[v] = true
				queue = append(queue, v)
				queueLevel = append(queueLevel, level+1)
			}
		}
	}

	if len(queue) == pos {
		return -1
	}
	return ans
}

func traceWords(word string) []string {
	buf := []byte(word)
	items := make([]string, 0)
	for i := 0; i < len(buf)-1; i++ {
		k := i + 1
		ok := true
		for j := i; j <= k; j++ {
			if buf[j] == '.' {
				ok = false
				break
			}
		}
		if ok {
			buf[i], buf[k] = buf[k], buf[i]
			key := string(buf)
			items = append(items, key)
			buf[i], buf[k] = buf[k], buf[i]
		}
	}
	return items
}
