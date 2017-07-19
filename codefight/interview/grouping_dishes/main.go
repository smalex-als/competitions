package main

import (
	"fmt"
	"sort"
)

func main() {
	dishes := [][]string{
		{"Salad", "Tomato", "Cucumber", "Salad", "Sauce"},
		{"Pizza", "Tomato", "Sausage", "Sauce", "Dough"},
		{"Quesadilla", "Chicken", "Cheese", "Sauce"},
		{"Sandwich", "Salad", "Bread", "Tomato", "Cheese"},
	}
	a := groupingDishes(dishes)
	fmt.Printf("a = %+v\n", a)
}

func groupingDishes(dishes [][]string) [][]string {
	res := make(map[string]map[string]bool)
	for _, v := range dishes {
		name := v[0]
		for i := 1; i < len(v); i++ {
			ingredient := v[i]
			if _, ok := res[ingredient]; !ok {
				res[ingredient] = map[string]bool{}
			}
			res[ingredient][name] = true
		}
	}
	names := make([]string, 0)
	for k, v := range res {
		if len(v) > 1 {
			names = append(names, k)
		}
	}
	sort.Strings(names)
	a := make([][]string, len(names))
	for i := 0; i < len(names); i++ {
		arr := make([]string, len(res[names[i]])+1)
		j := 1
		for v, _ := range res[names[i]] {
			arr[j] = v
			j++
		}
		arr[0] = names[i]
		a[i] = arr
	}
	return a
}
