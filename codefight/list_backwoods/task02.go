package main

func threeSplit(a []int) int {
	n := len(a)
	prefix := make([]int64, n)
	for i := 0; i < n; i++ {
		prefix[i] = int64(a[i])
		if i > 0 {
			prefix[i] += prefix[i-1]
		}
	}
	suffix := make([]int64, n)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = int64(a[i])
		if i < n-1 {
			suffix[i] += suffix[i+1]
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n-1; j++ {
			sumI := prefix[i]
			sumJ := prefix[j] - prefix[i]
			sumK := suffix[j+1]
			if sumK == sumJ && sumI == sumJ {
				res++
			}
		}
	}
	return res
}
