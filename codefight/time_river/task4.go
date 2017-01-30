package main

import "fmt"
import "time"

func main4() {
	fmt.Println(curiousClock("2016-08-26 22:40", "2016-08-29 10:00"))
}

func curiousClock(someTime string, leavingTime string) string {
	dt0, _ := time.Parse("2006-01-02 15:04", someTime)
	dt1, _ := time.Parse("2006-01-02 15:04", leavingTime)

	dt := dt0.Add(dt0.Sub(dt1))
	return dt.Format("2006-01-02 15:04")
}
