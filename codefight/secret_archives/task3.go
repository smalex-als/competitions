package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main3() {
	s := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR"
	ex := "RP4pr/NP4pn/BP4pb/QP4pq/K2P2pk/BP4pb/NP4pn/RP4pr"

	res := chessNotation(s)
	fmt.Println(res == ex)
}

func chessNotation(notation string) string {
	field := parseBoard(notation)
	var buffer bytes.Buffer
	for j := 0; j < 8; j++ {
		if j > 0 {
			buffer.WriteString("/")
		}
		cnt := 0
		for i := 7; i >= 0; i-- {
			if field[i][j] != " " {
				if cnt > 0 {
					buffer.WriteString(strconv.Itoa(cnt))
					cnt = 0
				}
				buffer.WriteString(field[i][j])
			} else {
				cnt++
			}
		}
		if cnt > 0 {
			buffer.WriteString(strconv.Itoa(cnt))
			cnt = 0
		}
	}
	return buffer.String()
}

func parseBoard(notation string) [][]string {
	field := make([][]string, 8)
	for i := 0; i < 8; i++ {
		field[i] = make([]string, 8)
	}
	row := 0
	for _, v := range strings.Split(notation, "/") {
		column := 0
		for _, ch := range v {
			if ch >= '0' && ch <= '8' {
				cnt := int(ch - '0')
				for i := 0; i < cnt; i++ {
					field[row][column] = " "
					column++
				}
			} else {
				field[row][column] = string(ch)
				column++
			}
		}
		row++
	}
	return field
}
