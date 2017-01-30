package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main4() {
	fmt.Println(nthNumber("8one 003number 201numbers li-000233le number444", 4))
	fmt.Println(nthNumber("some023020 num ber 033 02103 32 meh peh beh 4328", 4))
	fmt.Println(nthNumber("007 is an awesome agent", 1))
	fmt.Println(nthNumber("Everyone hates error 404", 1))
	fmt.Println(nthNumber("LaS003920tP3rEt4t04Yte0023s3t", 4))
}

func nthNumber(s string, n int) string {
	re := regexp.MustCompile("(?:0*([1-9][0-9]*)[^0-9]*){" + strconv.Itoa(n) + "}")
	return re.FindStringSubmatch(s)[1]
}
