package main

import "fmt"
import "time"

func main6() {
	fmt.Println(regularMonths("02-2016"))
}

func regularMonths(currMonth string) string {
	dt, _ := time.Parse("01-2006", currMonth)
	dt = dt.AddDate(0, 1, 0)
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 1, 0)
	}
	return dt.Format("01-2006")
}
