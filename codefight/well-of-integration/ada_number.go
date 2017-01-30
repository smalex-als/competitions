package switchlights

import (
	"regexp"
	"strconv"
	"strings"
)

func adaNumber(line string) bool {
	line = strings.ToUpper(strings.Replace(line, "_", "", -1))
	r := regexp.MustCompile("^([0-9]+)#([0-9A-F]+)#$")
	r2 := regexp.MustCompile("^[0-9]+$")
	if m := r.FindStringSubmatch(line); len(m) > 0 {
		base, _ := strconv.ParseInt(m[1], 10, 32)
		if base < 2 || base > 16 {
			return false
		}
		for i := 0; i < len(m[2]); i += 14 {
			end := i + 14
			if end > len(m[2]) {
				end = len(m[2])
			}
			substr := m[2][i:end]
			_, err := strconv.ParseInt(substr, int(base), 64)
			if err != nil {
				return false
			}
		}
		return true
	} else if r2.MatchString(line) {
		return true
	}
	return false
}
