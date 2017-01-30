package main

import "fmt"

func main14() {
	order := []int{200, 20, 15}
	shoppers := [][]int{{300, 40, 5}, {600, 40, 10}}
	res := delivery(order, shoppers)
	fmt.Printf("res = %+v\n", res)
}

func delivery(order []int, shoppers [][]int) []bool {
	res := make([]bool, len(shoppers))
	for k, shopper := range shoppers {
		dist := (order[0] + shopper[0]) + shopper[1]*shopper[2]
		lo := order[1] * shopper[1]
		hi := lo + order[2]*shopper[1]
		if lo <= dist && dist <= hi {
			res[k] = true
		}
	}
	return res
}
