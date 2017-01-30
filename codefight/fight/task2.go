package main

import (
	"fmt"
	"regexp"
)

func main2() {

	//	code1 := []string{"def is_even_sum(a, b):",
	//		"    return (a + b) % 2 == 0"}
	//	code2 := []string{" def is_even_sum(summand_1, summand_2):",
	//		"    return (summand_1 + summand_2) % 2 == 0"}
	//	code1 := []string{"def is_even_sum(a, b):",
	//		"    return (a + b)"}
	//	code2 := []string{"def is_even_sum(b, a):",
	//		"    return (a + b)"}
	//
	// code1 := []string{"function is_even_sum(a, b) {",
	// 	"  return (a + b) % 2 === 0;",
	// 	"}"}
	// code2 := []string{"function is_even_sum(b, a) {",
	// 	"  return (b + a) % 2 === 1;",
	// 	"}"}

	code1 := []string{"def return_first(f, s):",
		"  t = f",
		"  return f"}
	code2 := []string{"def return_first(f, s):",
		"  f = f",
		"  return f"}
	res := plagiarismCheck(code1, code2)
	fmt.Println(res)
}

func plagiarismCheck(code1 []string, code2 []string) bool {
	c1 := normalizecode(code1)
	c2 := normalizecode(code2)
	if len(c1) != len(c2) {
		return false
	}
	r := regexp.MustCompile("([a-zA-Z_][a-z_A-Z0-9]*)")
	nextId := 0
	fixed := make(map[string]bool)
	for i := 0; i < len(c1); i++ {
		if r.MatchString(c1[i]) && r.MatchString(c2[i]) {
			if c1[i] == c2[i] {
				fixed[c1[i]] = true
			} else {
				_, ok := fixed[c2[i]]
				if ok {
					look1 := c1[i]
					for j := 0; j < len(c1); j++ {
						if c1[j] == look1 {
							c1[j] = c2[i]
						}
					}
				} else {
					nextId++
					varName := fmt.Sprintf("var%d", nextId)
					look1 := c1[i]
					look2 := c2[i]
					for j := 0; j < len(c1); j++ {
						if c1[j] == look1 {
							c1[j] = varName
						}
						if c2[j] == look2 {
							c2[j] = varName
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}

func normalizecode(code1 []string) []string {
	r := regexp.MustCompile("([a-zA-Z_][a-z_A-Z0-9]*)")
	res := make([]string, 0)
	for i := 0; i < len(code1); i++ {
		row := code1[i]
		matched := r.FindAllStringIndex(row, -1)
		for i := 0; i < len(matched); i++ {
			if i > 0 {
				res = append(res, row[matched[i-1][1]:matched[i][0]])
			} else {
				res = append(res, row[0:matched[i][0]])
			}
			res = append(res, row[matched[i][0]:matched[i][1]])
		}
		if len(matched) > 0 {
			res = append(res, row[matched[len(matched)-1][1]:])
		} else {
			res = append(res, row)
		}
	}
	return res
}
