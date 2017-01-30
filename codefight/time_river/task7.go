package main

import "fmt"
import "time"

func main7() {
	fmt.Println(missedClasses(2015, []int{2, 3},
		[]string{"11-04", "02-22", "02-23", "03-07", "03-08", "05-09"}))
}

func missedClasses(year int, daysOfTheWeek []int, holidays []string) int {
	const layout = "2006-01-02"
	from, _ := time.Parse(layout, fmt.Sprintf("%d-08-31", year))
	to, _ := time.Parse(layout, fmt.Sprintf("%d-06-01", year+1))
	res := 0
	for j := year; j <= year+1; j++ {
		for i := 0; i < len(holidays); i++ {
			str := fmt.Sprintf("%d-%s", j, holidays[i])
			dt, _ := time.Parse(layout, str)
			if dt.After(from) && dt.Before(to) {
				lookup := int(dt.Weekday())
				if lookup == 0 {
					lookup = 7
				}
				for k := 0; k < len(daysOfTheWeek); k++ {
					if daysOfTheWeek[k] == lookup {
						res++
					}
				}
			}
		}
	}
	return res
}
