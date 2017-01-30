package main

import "fmt"
import "time"

func main3() {
	res := dayOfWeek("02-29-2016")
	fmt.Println(res)
}

func dayOfWeek(birthdayDate string) int {
	dt, _ := time.Parse("01-02-2006", birthdayDate)
	for i := 1; i < 1000; i++ {
		tmpdt := dt.AddDate(i, 0, 0)
		if tmpdt.Weekday() == dt.Weekday() && tmpdt.Month() == dt.Month() {
			return tmpdt.Year() - dt.Year()
		}
	}
	return 0
}
