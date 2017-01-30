package switchlights

import "sort"

func electionsWinners(votes []int, k int) int {
	n := len(votes)
	sort.Ints(votes)
	res := 1
	max := votes[n-1]
	if k == 0 && n > 1 {
		if votes[n-1] == votes[n-2] {
			return 0
		}
	}
	for i := n - 2; i >= 0 && votes[i]+k > max; i-- {
		res++
	}
	return res
}
