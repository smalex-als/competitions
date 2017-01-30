package switchlights

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestElectionsWinner(t *testing.T) {
	var tests = []struct {
		a        []int
		k        int
		expected int
	}{
		{[]int{2, 3, 5, 2}, 3, 2},
		{[]int{1, 3, 3, 1, 1}, 0, 0},
	}

	for _, c := range tests {
		actual := electionsWinners(c.a, c.k)
		expected := c.expected
		if !reflect.DeepEqual(expected, actual) {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, expected, actual)
			t.FailNow()
		}
	}
}
