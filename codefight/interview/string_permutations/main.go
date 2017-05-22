package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(stringPermutations("CBA"))
	fmt.Println(stringPermutations("ABA"))
}

func stringPermutations(s string) []string {
	a := make([]int, len(s))
	for k, v := range s {
		a[k] = int(v)
	}
	sort.Ints(a)
	m := make(map[string]bool)
	perm(a, 0, m)
	res := make([]string, 0)
	for k, _ := range m {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}

func perm(a []int, pos int, m map[string]bool) {
	if pos == len(a) {
		b := make([]byte, len(a))
		for j := 0; j < len(a); j++ {
			b[j] = byte(a[j])
		}
		m[string(b)] = true
		return
	}
	for i := pos; i < len(a); i++ {
		for ; i+1 < len(a) && a[i] == a[i+1]; i++ {
		}
		swap(a, pos, i)
		perm(a, pos+1, m)
		swap(a, pos, i)
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
