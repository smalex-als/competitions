package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println(bugsAndBugfixes("Roll d6-3 and 4d4+3 to pick a weapon, and finish the boss with 3d7!"))
}

func bugsAndBugfixes(rules string) int {
	pattern := regexp.MustCompile("([1-6])?d(\\d+)([+-][0-9])?")
	formulas := pattern.FindAllStringSubmatch(rules, -1)
	fmt.Println(formulas)

	res := 0
	for _, formula := range formulas {
		rolls := 0
		if formula[1] != "" {
			rolls, _ = strconv.Atoi(formula[1])
		} else {
			rolls = 1
		}
		dieType, _ := strconv.Atoi(formula[2])
		formulaMax := rolls * dieType

		if formula[3] != "" {
			b, _ := strconv.Atoi(formula[3][1:len(formula[3])])
			if formula[3][0] == '-' {
				formulaMax = formulaMax - b
			} else {
				formulaMax = formulaMax + b
			}
		}
		res += formulaMax
	}
	return res
}
