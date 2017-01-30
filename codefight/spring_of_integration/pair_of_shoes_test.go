package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestPairOfShoes(t *testing.T) {
	var tests = []struct {
		a        [][]int
		expected bool
	}{
		{[][]int{{0, 21}, {1, 23}, {1, 21}, {0, 23}}, true},
	}

	for _, c := range tests {
		actual := pairOfShoes(c.a)
		expected := c.expected
		if !reflect.DeepEqual(expected, actual) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, expected, actual)
			t.FailNow()
		}
	}
}
