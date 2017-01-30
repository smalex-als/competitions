package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main15() {
	rices := []float32{110, 95, 70}
	notes := []string{
		"10.0% higher than in-store",
		"5.0% lower than in-store",
		"Same as in-store",
	}
	x := float32(5.0)
	res := isAdmissibleOverpayment(rices, notes, x)
	fmt.Printf("res = %+v\n", res)
}

func isAdmissibleOverpayment(prices []float32, notes []string, x float32) bool {
	re := regexp.MustCompile("[^0-9.]+")
	var sum float64
	for i := 0; i < len(notes); i++ {
		less := strings.Contains(notes[i], " lower ")
		note := re.ReplaceAllString(notes[i], "")
		if len(note) > 0 {
			pct, _ := strconv.ParseFloat(note, 64)
			price := float64(prices[i])
			if less {
				diff := price - price*100/(-pct+100)
				sum += diff
			} else {
				diff := price - price*100/(pct+100)
				sum += diff
			}
		}
	}
	if sum <= float64(x) {
		return true
	}
	return false
}
