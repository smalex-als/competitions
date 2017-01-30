package switchlights

import (
	"regexp"
	"strings"
)

func timedReading(maxLength int, text string) int {
	text = strings.ToLower(text)
	r, _ := regexp.Compile("[^a-z]")
	text = r.ReplaceAllString(text, " ")
	cnt := 0
	for _, s := range strings.Split(text, " ") {
		if len(s) != 0 && len(s) <= maxLength {
			cnt++
		}
	}
	return cnt
}
