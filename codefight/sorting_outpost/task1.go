package main

import (
	"fmt"
	"sort"
)

func main1() {
	res := sortByLength([]string{"abc", "", "aaa", "a", "zz"})
	fmt.Println(res)
}

func sortByLength(inputArray []string) []string {
	sort.Sort(ByLength(inputArray))
	return inputArray
}

type ByLength []string

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
