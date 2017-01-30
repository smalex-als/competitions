package switchlights

func areSimilar(a []int, b []int) bool {
	cnt := 0
	index := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
			if cnt == 1 {
				index = i
			} else if cnt == 2 {
				if a[index] != b[i] || a[i] != b[index] {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func switchLights(a []int) []int {
	n := len(a)
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		cnt += a[i]
		if cnt%2 == 1 {
			a[i] ^= 1
		}
	}
	return a
}
