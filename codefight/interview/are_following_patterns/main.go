package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"cat", "dog", "dog"}
	patterns := []string{"a", "b", "b"}
	res := areFollowingPatterns(strings, patterns)
	fmt.Printf("res = %+v\n", res)
}

func areFollowingPatterns(strings []string, patterns []string) bool {
	a := make([][]string, len(strings))
	for i := 0; i < len(a); i++ {
		a[i] = []string{strings[i], patterns[i]}
	}
	sort.Sort(ByName(a))
	for i := 1; i < len(a); i++ {
		if a[i][0] == a[i-1][0] {
			if a[i][1] != a[i-1][1] {
				return false
			}
		}
	}
	sort.Sort(ByName2(a))
	for i := 1; i < len(a); i++ {
		if a[i][1] == a[i-1][1] {
			if a[i][0] != a[i-1][0] {
				return false
			}
		}
	}
	return true
}

type ByName [][]string

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i][0] < a[j][0] }

type ByName2 [][]string

func (a ByName2) Len() int           { return len(a) }
func (a ByName2) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName2) Less(i, j int) bool { return a[i][1] < a[j][1] }
