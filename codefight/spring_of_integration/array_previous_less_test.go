package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestArrayPreviousLessBig(t *testing.T) {
	a := make([]int, 1000*1000*1)
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(1000)
	}
	arrayPreviousLess(a)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Mb\n", m.Alloc/1024/1024)
}

func TestArrayPreviousLess(t *testing.T) {
	var tests = []struct {
		a        []int
		expected []int
	}{
		{[]int{3, 5, 2, 4, 5}, []int{-1, 3, -1, 2, 4}},
	}

	for _, c := range tests {
		actual := arrayPreviousLess(c.a)
		expected := c.expected
		if !reflect.DeepEqual(expected, actual) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, expected, actual)
			t.FailNow()
		}
	}
}
