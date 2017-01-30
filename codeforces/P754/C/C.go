package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type message struct {
	from  string
	text  string
	users map[string]bool
}

var scanner *bufio.Scanner

func solve() {
	t := readInt()
	for i := 0; i < t; i++ {
		solveTest()
	}
}

func solveTest() {
	readInt()
	id := 0
	people := make(map[string]int)
	for _, v := range strings.Split(readString(), " ") {
		people[v] = id
		id++
	}
	cnt := readInt()
	messages := make([]message, cnt)
	re := regexp.MustCompile("[ ,!\\?\\.]+")
	for j := 0; j < cnt; j++ {
		line := readString()
		p := strings.SplitN(line, ":", 2)
		users := make(map[string]bool)
		for _, v := range re.Split(p[1], -1) {
			if _, ok := people[v]; ok {
				users[v] = true
			}
		}
		messages[j] = message{p[0], p[1], users}
	}
	for i := 0; i < len(people); i++ {
		for j := 0; j < len(messages); j++ {
			if messages[j].from == "?" {
				if j-1 >= 0 && messages[j-1].from != "?" {
					u := messages[j-1].from
					if _, ok := messages[j].users[u]; !ok {
						messages[j].users[u] = true
					}
				}
				if j+1 < len(messages) && messages[j+1].from != "?" {
					u := messages[j+1].from
					if _, ok := messages[j].users[u]; !ok {
						messages[j].users[u] = true
					}
				}

				if len(messages[j].users) >= len(people) {
					fmt.Println("Impossible")
					return
				}
				if len(messages[j].users)+1 == len(people) {
					for p := range people {
						if _, ok := messages[j].users[p]; !ok {
							messages[j].from = p
							break
						}
					}
				}
			} else {
				if j-1 >= 0 && messages[j-1].from == messages[j].from {
					fmt.Println("Impossible")
					return
				}
				if j+1 < len(messages) && messages[j+1].from == messages[j].from {
					fmt.Println("Impossible")
					return
				}
				if _, ok := messages[j].users[messages[j].from]; ok {
					fmt.Println("Impossible")
					return
				}
			}
		}
	}
	for i := 0; i < len(messages); i++ {
		if messages[i].from == "?" {
			for p := range people {
				if _, ok := messages[i].users[p]; !ok {
					messages[i].from = p
					if i+1 < len(messages) {
						messages[i+1].users[p] = true
					}
					break
				}
			}
		}
	}
	for _, v := range messages {
		fmt.Printf("%s:%s\n", v.from, v.text)
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	//	scanner.Split(bufio.ScanWords)
	solve()
}

//
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
