package main

import "fmt"

func main4() {
	table := []string{
		"+----+--+-----+----+",
		"|abcd|56|!@#$%|qwer|",
		"|1234|78|^&=()|tyui|",
		"+----+--+-----+----+",
		"|zxcv|90|77777|stop|",
		"+----+--+-----+----+",
		"|asdf|~~|ghjkl|100$|",
		"+----+--+-----+----+",
	}

	res := cellsJoining(table, [][]int{{2, 0}, {1, 1}})
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func cellsJoining(table []string, coords [][]int) []string {
	columns := make([]int, 0)
	for k, v := range table[0] {
		if v == '+' {
			columns = append(columns, k)
		}
	}
	maxh, maxw := len(table)-1, len(table[0])-1
	rows := make([]int, 0)
	for i := 0; i < len(table); i++ {
		if table[i][0] == '+' {
			rows = append(rows, i)
		}
	}
	m := make([][]rune, len(table))
	for i := 0; i < len(table); i++ {
		m[i] = []rune(table[i])
	}
	y0, x0 := coords[1][0], coords[0][1]
	y1, x1 := coords[0][0], coords[1][1]
	for i := y0 + 1; i <= y1; i++ {
		sx := columns[x0]
		ex := columns[x1+1]
		for j := sx; j <= ex; j++ {
			if j == 0 || j == maxw {
				m[rows[i]][j] = '|'
			} else {
				if j != sx && j != ex {
					m[rows[i]][j] = ' '
				}
			}
		}
	}
	for i := x0 + 1; i <= x1; i++ {
		sy := rows[y0]
		ey := rows[y1+1]
		for j := sy; j <= ey; j++ {
			if j == 0 || j == maxh {
				m[j][columns[i]] = '-'
			} else {
				if j != sy && j != ey {
					m[j][columns[i]] = ' '
				}
			}
		}
	}
	for i := 0; i < len(table); i++ {
		table[i] = string(m[i])
	}
	return table
}
