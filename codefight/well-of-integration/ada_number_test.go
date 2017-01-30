package switchlights

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func TestAdaNumber(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"16#1234567890ABCDEFabcdef#", true},
		{"123_456_789", true},
		{"16#123abc#", true},
		{"10#123abc#", false},
		{"10#10#123ABC#", false},
		{"123", true},
	}

	for num, c := range tests {
		actual := adaNumber(c.str)
		expected := c.expected
		if expected != actual {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("Test%d\n%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", num, filepath.Base(file), line, expected, actual)
			t.FailNow()
		}
	}

}
