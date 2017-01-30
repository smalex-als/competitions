package switchlights

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestAreSimilar(t *testing.T) {
	var tests = []struct {
		a, b []int
		res  bool
	}{
		{[]int{1, 2, 3}, []int{2, 1, 3}, true},
		{[]int{1, 2, 2}, []int{2, 1, 1}, false},
		{[]int{1, 2, 4}, []int{1, 10, 2}, false},
	}

	for num, c := range tests {
		actual := areSimilar(c.a, c.b)
		expected := c.res
		if !reflect.DeepEqual(expected, actual) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("Test%d\n\n%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", num, filepath.Base(file), line, expected, actual)
			t.FailNow()
		}
	}
}
func TestFunction(t *testing.T) {
	var tests = []struct {
		a, expected []int
	}{
		{[]int{1, 1, 1, 1, 1}, []int{0, 1, 0, 1, 0}},
		{[]int{1, 0, 0, 1, 0, 1, 0, 1}, []int{1, 1, 1, 0, 0, 1, 1, 0}},
	}

	for _, c := range tests {
		actual := switchLights(c.a)
		expected := c.expected
		if !reflect.DeepEqual(expected, actual) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, expected, actual)
			t.FailNow()
		}
	}
}
