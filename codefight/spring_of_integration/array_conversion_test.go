package main

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestArrayConversion(t *testing.T) {
	var tests = []struct {
		a        []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 186},
	}

	for _, c := range tests {
		actual := arrayConversion(c.a)
		expected := c.expected
		if !reflect.DeepEqual(expected, actual) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, expected, actual)
			t.FailNow()
		}
	}
}
