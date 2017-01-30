package switchlights

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func TestTimedReading(t *testing.T) {
	var tests = []struct {
		l        int
		str      string
		expected int
	}{
		{4, "The Fox asked the stork, 'How is the soup?'", 7},
		{5, "Although zebra species may have overlapping ranges, they do not interbreed.", 6},
	}

	for num, c := range tests {
		actual := timedReading(c.l, c.str)
		expected := c.expected
		if expected != actual {
			_, file, line, _ := runtime.Caller(0)
			fmt.Printf("Test%d\n%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", num, filepath.Base(file), line, expected, actual)
			t.FailNow()
		}
	}
}
