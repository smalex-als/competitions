package main

import "fmt"

func main() {
	t := &Tree{
		Value: 1,
		Left: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 3,
			},
			Right: &Tree{
				Value: 4,
			},
		},
		Right: &Tree{
			Value: 2,
			Left: &Tree{
				Value: 4,
			},
			Right: &Tree{
				Value: 3,
			},
		},
	}
	fmt.Println(isTreeSymmetric(t))
}

// Definition for binary tree:
type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}

type Point struct {
	n     *Tree
	level int
}

// ok I know the better solution exists
func isTreeSymmetric(t *Tree) bool {
	empty := &Tree{Value: -999999}
	if t == nil {
		return true
	}
	q := make([]*Point, 0)
	q = append(q, &Point{t, 0})
	a := make([]int, 0)
	level := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if level != cur.level {
			cont := false
			for i := 0; i < len(a); i++ {
				if a[i] != -999999 {
					cont = true
					break
				}
			}
			if !cont {
				break
			}
			for i := 0; i < len(a)/2; i++ {
				if a[i] != a[len(a)-i-1] {
					return false
				}
			}
			a = make([]int, 0)
			level = cur.level
		}
		a = append(a, cur.n.Value.(int))
		if cur.n.Left != nil {
			q = append(q, &Point{cur.n.Left, cur.level + 1})
		} else {
			q = append(q, &Point{empty, cur.level + 1})
		}
		if cur.n.Right != nil {
			q = append(q, &Point{cur.n.Right, cur.level + 1})
		} else {
			q = append(q, &Point{empty, cur.level + 1})
		}
	}
	return true
}
