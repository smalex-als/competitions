package main

import (
	"fmt"
	"sort"
)

func main17() {
	healthPoints := []int{110, 30, 50}
	attackDamage := []int{12, 11, 20}
	res := allianceVersusMonster(healthPoints, attackDamage)
	fmt.Printf("res = %+v\n", res)
}

type Warrior struct {
	h, d int
}

type ByDamage []Warrior

func (a ByDamage) Len() int           { return len(a) }
func (a ByDamage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDamage) Less(i, j int) bool { return a[i].d > a[j].d }

func allianceVersusMonster(healthPoints []int, attackDamage []int) int {
	health := healthPoints[0]
	attack := attackDamage[0]
	warriors := make([]Warrior, 0)
	for i := 1; i < len(healthPoints); i++ {
		warriors = append(warriors, Warrior{healthPoints[i], attackDamage[i]})
	}
	sort.Sort(ByDamage(warriors))
	for _, w := range warriors {
		for w.h > attack {
			health -= w.d
			w.h -= attack
		}
	}
	if health <= 0 {
		return len(warriors)
	}
	died := 0
	for _, w := range warriors {
		health -= w.d
		if health <= 0 {
			return len(warriors) - died
		}
		died++
	}
	return 0
}
