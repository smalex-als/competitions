package main

import "fmt"

func main5() {
	fmt.Println(chessBishopDream([]int{3, 7}, []int{1, 2}, []int{-1, 1}, 1003))
}

func chessBishopDream(size []int, pos []int, dir []int, k int) []int {
	h := size[0]
	w := size[1]
	y := pos[0]
	x := pos[1]
	m := make(map[string]int)
	for i := 0; i < k; i++ {
		ny := y + dir[0]
		nx := x + dir[1]
		if (ny == -1 || ny == h) && (nx == -1 || nx == w) {
			dir[0] = -dir[0]
			dir[1] = -dir[1]
		} else if ny == -1 {
			dir[0] = -dir[0]
			y = 0
			x = nx
		} else if ny == h {
			dir[0] = -dir[0]
			y = h - 1
			x = nx
		} else if nx == -1 {
			dir[1] = -dir[1]
			x = 0
			y = ny
		} else if nx == w {
			dir[1] = -dir[1]
			x = w - 1
			y = ny
		} else {
			y = ny
			x = nx
		}
		key := fmt.Sprintf("%d,%d,%d,%d", y, x, dir[0], dir[1])
		nk, ok := m[key]
		if ok {
			if k-i > i-nk {
				k -= i - nk
			}
		} else {
			m[key] = i
		}
	}
	return []int{y, x}
}
