package main

import "strings"
import "fmt"

func main2() {
	text := "<table><tr><td>jQu9ABs8l</td><td>9alQS</td><td>6j</td><td>x0C</td><td>VJwINu0wjE</td></tr><tr><td>52K</td><td>w5P</td><td>K0HTHBB</td><td>76H</td><td>2Up4kl</td></tr><tr><td>d7J9bn7lx</td><td>unJT</td><td>mdICgjl</td><td>v0</td><td>LKvS1LbYBo</td></tr><tr><td>eld9</td><td>O</td><td>Yqe184E9</td><td>b45QX0313A</td><td>4M02</td></tr><tr><td>6XKiOf96</td><td>wb7</td><td>HW5535kri</td><td>81U</td><td>V64O2502a</td></tr><tr><td>o8</td><td>col7G7g</td><td>y92s3R</td><td>q1</td><td>zl0LizILrm</td></tr></table>"
	fmt.Println(htmlTable(text, 5, 4))
}

func htmlTable(table string, row int, column int) string {
	for i := 0; i <= row; i++ {
		index := strings.Index(table, "<tr>")
		if index < 0 {
			return "No such cell"
		}
		table = table[index+4:]
	}
	if strings.HasPrefix(table, "<td>") {
		table = table[4:]
		res := make([]string, 0)
		for i := 0; ; i++ {
			index := strings.Index(table, "</td>")
			if index < 0 {
				break
			}
			res = append(res, table[:index])
			table = table[index+5:]
			if strings.HasPrefix(table, "</tr>") {
				break
			}
			index = strings.Index(table, "<td>")
			table = table[index+4:]
		}
		if column < len(res) {
			return res[column]
		}
	}
	return "No such cell"
}
