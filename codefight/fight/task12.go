package main

import (
	"fmt"
	"sort"
	"strings"
)

func main12() {
	// redirects := [][]string{
	// 	{"godaddy.net", "godaddy.com"},
	// 	{"godaddy.org", "godaddycares.com"},
	// 	{"godady.com", "godaddy.com"},
	// 	{"godaddy.ne", "godaddy.net"},
	// }
	redirects := [][]string{
		{"a", "z"},
		{"c", "b"},
	}
	res := domainForwarding(redirects)
	fmt.Println(res)
}

func domainForwarding(redirects [][]string) [][]string {
	visited := make(map[string]bool, 0)
	preres := make(map[string][]string, 0)
	allKeys := make([]string, 0)
	for i := 0; i < len(redirects); i++ {
		keys := make(map[string]bool)
		path := make(map[string]bool)
		domainDFS(redirects, i, i, visited, path, keys)
		if len(keys) > 0 {
			keysString := make([]string, len(keys))
			pos := 0
			for k, _ := range keys {
				keysString[pos] = k
				pos++
			}
			pathString := make([]string, len(path))
			pos = 0
			for k, _ := range path {
				pathString[pos] = k
				pos++
			}
			sort.Strings(keysString)
			sort.Strings(pathString)
			allKeys = append(allKeys, strings.Join(keysString, "|"))
			preres[strings.Join(keysString, "|")] = pathString
		}
	}
	sort.Strings(allKeys)
	res := make([][]string, len(allKeys))
	for i, v := range allKeys {
		res[i] = preres[v]
	}
	return res
}

func domainDFS(redirects [][]string, index int, root int, visited map[string]bool,
	path map[string]bool, keys map[string]bool) {
	cur := redirects[index]
	if visited[cur[0]] {
		return
	}
	visited[cur[0]] = true
	path[cur[0]] = true
	path[cur[1]] = true
	keys[cur[1]] = true
	for i := 0; i < len(redirects); i++ {
		d0 := redirects[i][0]
		d1 := redirects[i][1]
		if d0 == cur[1] || d1 == cur[1] || d0 == cur[0] || d1 == cur[0] {
			domainDFS(redirects, i, root, visited, path, keys)
		}
	}
}
