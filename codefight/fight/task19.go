package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main19() {
	requests := []string{
		"I need a new window.",
		"I really, really want to replace my window!",
		"Replace my window.",
		"I want a new window.",
		"I want a new carpet, I want a new carpet, I want a new carpet.",
		"Replace my carpet",
	}
	IDs := []int{374, 2845, 83, 1848, 1837, 1500}
	threshold := 0.5
	res := spamClusterization(requests, IDs, threshold)
	fmt.Printf("res = %+v\n", res)
}

type Req struct {
	ID      int
	words   map[string]bool
	similar []int
	visited bool
}

func (r *Req) add(v int) {
	r.similar = append(r.similar, v)
}

func spamClusterization(requests []string, IDs []int, threshold float64) [][]int {
	re := regexp.MustCompile("[A-Za-z]+")
	reqs := make([]Req, len(requests))
	for i, r := range requests {
		words := make(map[string]bool)
		for _, v := range re.FindAllString(strings.ToLower(r), -1) {
			words[v] = true
		}
		reqs[i] = Req{IDs[i], words, make([]int, 0), false}
	}
	sort.Sort(ById(reqs))
	for i := 0; i < len(reqs); i++ {
		for j := i + 1; j < len(reqs); j++ {
			if jaccardIndex(reqs[i], reqs[j]) >= threshold {
				reqs[i].add(j)
				reqs[j].add(i)
			}
		}
	}
	res := make([][]int, 0)
	for i := 0; i < len(reqs); i++ {
		if !reqs[i].visited {
			v := make([]int, 0)
			dfsSimilar(&reqs[i], reqs, &v)
			if len(v) > 1 {
				sort.Ints(v)
				res = append(res, v)
			}
		}
	}
	return res
}

type ById []Req

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].ID < a[j].ID }

func dfsSimilar(req *Req, reqs []Req, v *[]int) {
	res := *v
	req.visited = true
	res = append(res, req.ID)
	for _, v := range req.similar {
		nextReq := &reqs[v]
		if !nextReq.visited {
			dfsSimilar(nextReq, reqs, &res)
		}
	}
	*v = res
}

func jaccardIndex(req1, req2 Req) float64 {
	m := make(map[string]int)
	cmn := 0
	for k, _ := range req1.words {
		m[k]++
	}
	for k, _ := range req2.words {
		m[k]++
		if m[k] == 2 {
			cmn++
		}
	}
	return float64(cmn) / float64(len(m))
}
