package main

import (
	"fmt"
	"math"
	"sort"
)

func main5() {
	vendorsDelivery := []int{5, 4, 2, 3}
	vendorsProducts := [][]int{
		{1, 1, 1},
		{3, -1, 3},
		{-1, 2, 2},
		{5, -1, -1},
	}
	res := minimalBasketPrice(7, vendorsDelivery, vendorsProducts)
	fmt.Println(res)
}

var bestIndexes map[int]bool
var bestTime int

func minimalBasketPrice(maxPrice int, vendorsDelivery []int, vendorsProducts [][]int) []int {
	bestTime = math.MaxInt32
	res := make([]int, len(vendorsProducts[0]))
	dfs(maxPrice, vendorsDelivery, vendorsProducts, 0, 0, res, 0, 0)

	indexes := make([]int, 0)
	for k, _ := range bestIndexes {
		indexes = append(indexes, k)
	}
	sort.Ints(indexes)
	return indexes
}

func dfs(maxPrice int, vendors []int, products [][]int, x int, value int,
	res []int, cost, maxDays int) {
	if cost > maxPrice {
		return
	}
	if x == len(products[0]) {
		if bestTime > maxDays {
			values := make(map[int]bool)
			for i := 0; i < len(products[0]); i++ {
				values[res[i]] = true
			}
			bestTime = maxDays
			bestIndexes = values
		}
		return
	}
	for i := 0; i < len(vendors); i++ {
		res[x] = i
		if products[i][x] != -1 && bestTime > max(maxDays, vendors[i]) {
			dfs(maxPrice, vendors, products, x+1, 0, res,
				cost+products[i][x], max(maxDays, vendors[i]))
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
