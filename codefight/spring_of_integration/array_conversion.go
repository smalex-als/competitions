package main

func arrayConversion(a []int) int {
	b := 0
	for n := len(a); n > 1; n /= 2 {
		for i := 0; i < n; i += 2 {
			if b%2 == 0 {
				a[i/2] = a[i] + a[i+1]
			} else {
				a[i/2] = a[i] * a[i+1]
			}
		}
		b++
	}
	return a[0]
}
