package main

import (
	"fmt"
	"sort"
	"strings"
)

func main4() {

	catalog := [][]string{{"Books", "Classics", "Fiction"},
		{"Electronics", "Cell Phones", "Computers", "Ultimate item"},
		{"Grocery", "Beverages", "Snacks"},
		{"Snacks", "Chocolate", "Sweets"},
		{"root", "Books", "Electronics", "Grocery"}}

	updates := [][]string{{"Snacks", "Marmalade"},
		{"Fiction", "Harry Potter"},
		{"root", "T-shirts"},
		{"T-shirts", "CodeFights"}}

	res := catalogUpdate(catalog, updates)
	for i := 0; i < len(res); i++ {
		fmt.Println(strings.Join(res[i], "|"))
	}
}

func catalogUpdate(catalog [][]string, updates [][]string) [][]string {
	root := make(map[string]map[string]bool)
	for _, row := range catalog {
		rowMap := make(map[string]bool)
		for j := 1; j < len(row); j++ {
			rowMap[row[j]] = true
		}
		root[row[0]] = rowMap
	}
	for _, row := range updates {
		rowMap, exists := root[row[0]]
		if !exists {
			rowMap = make(map[string]bool)
			root[row[0]] = rowMap
		}
		for j := 1; j < len(row); j++ {
			rowMap[row[j]] = true
		}
	}
	keys := make([]string, 0, len(root))
	for k := range root {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	res := make([][]string, len(keys))
	for i, v := range keys {
		res[i] = make([]string, len(root[v])+1)
		res[i][0] = v
		names := make([]string, 0, len(root[v]))
		for k := range root[v] {
			names = append(names, k)
		}
		sort.Strings(names)
		for j := 0; j < len(names); j++ {
			res[i][j+1] = names[j]
		}
	}
	return res
}
