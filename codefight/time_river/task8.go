package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(holiday(3, "Wednesday", "November", 2016))
}
func holiday(x int, weekDay string, month string, yearNumber int) int {
	str := fmt.Sprintf("1 %s %d", month, yearNumber)
	dt, _ := time.Parse("2 January 2006", str)
	for i := 0; ; i++ {
		ndt := dt.AddDate(0, 0, i)
		if int(ndt.Month()) != int(dt.Month()) {
			break
		}
		if ndt.Format("Monday") == weekDay {
			x--
			if x == 0 {
				return ndt.Day()
			}
		}
	}
	return -1
}
