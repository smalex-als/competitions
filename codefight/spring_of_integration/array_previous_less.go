package main

func arrayPreviousLess(a []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = -1
		for j := i - 1; j >= 0; j-- {
			if a[j] < a[i] {
				res[i] = a[j]
				break
			}
		}
	}
	return res
}
