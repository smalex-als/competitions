package main

var res []string
var str []rune

func main() {
	generateParentheses(4)
}

func generateParentheses(n int) []string {
	str = make([]rune, n*2)
	res = make([]string, 0)
	solve(n, 0, 0, 0)

	return res
}

func solve(n int, pos int, opened int, closed int) {
	if pos == n*2 {
		res = append(res, string(str))
		return
	}
	if opened < n {
		str[pos] = '('
		solve(n, pos+1, opened+1, closed)
	}
	if opened > closed && closed < n {
		str[pos] = ')'
		solve(n, pos+1, opened, closed+1)
	}
}
