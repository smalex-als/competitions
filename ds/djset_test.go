package ds

import (
	"fmt"
	"testing"
)

func TestNewDJSet(t *testing.T) {
	dj := NewDJSet(10)
	for i := 0; i < 10; i++ {
		DJSetAssertCount(t, dj, i, 1)
	}
	DJSetAssertComponents(t, dj, 10)
	for i := 1; i < 4; i++ {
		DJSetAssertEquiv(t, dj, 0, i, false)
		dj.Union(0, i)
		DJSetAssertEquiv(t, dj, 0, i, true)
		DJSetAssertComponents(t, dj, 10-i)
	}
	for i := 5; i < 10; i++ {
		DJSetAssertEquiv(t, dj, 4, i, false)
		dj.Union(4, i)
		DJSetAssertEquiv(t, dj, 4, i, true)
	}
	DJSetAssertComponents(t, dj, 2)
	dj.Union(4, 0)
	DJSetAssertEquiv(t, dj, 4, 0, true)
	fmt.Println(dj.upper)
}

func DJSetAssertEquiv(t *testing.T, dj *DJSet, x int, y int, expected bool) {
	val := dj.Equiv(x, y)
	if val != expected {
		t.Fatalf("Expected equiv of djset %v, but was %v", expected, val)
	}
}

func DJSetAssertCount(t *testing.T, dj *DJSet, root int, expected int) {
	count := dj.Count(root)
	if count != expected {
		t.Fatalf("Expected count of djset %d, but was %d", expected, count)
	}
}

func DJSetAssertComponents(t *testing.T, dj *DJSet, expected int) {
	count := dj.Components()
	if count != expected {
		t.Fatalf("Expected number of components %d, but was %d", expected, count)
	}
}
