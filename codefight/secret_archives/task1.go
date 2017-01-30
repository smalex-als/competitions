package main

import "fmt"
import "strconv"

import "regexp"

func main1() {

	lrcLyrics := []string{
		"[00:12.00] Happy birthday dear coder,",
		"[00:17.20] Happy birthday to you!"}
	res := lrc2subRip(lrcLyrics, "00:00:20")
	// lrcLyrics := []string{
	// 	"[95:19.55] i hear babies cryin,",
	// 	"[95:23.31] i watch them grow",
	// 	"[95:26.05] theyll learn much more",
	// 	"[95:29.18] than ill ever know",
	// 	"[95:33.10] and i think to myself,",
	// 	"[95:38.44] what a wonderful world"}
	// res := lrc2subRip(lrcLyrics, "02:23:44")

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func lrc2subRip(rows []string, songLength string) []string {
	res := make([]string, 0)
	times := make([]string, len(rows)+1)
	reg := regexp.MustCompile("\\[(\\d+):(\\d+)\\.(\\d+)\\]")
	for i := 0; i < len(rows); i++ {
		p := reg.FindStringSubmatch(rows[i])
		m, _ := strconv.Atoi(p[1])
		s, _ := strconv.Atoi(p[2])
		ms, _ := strconv.Atoi(p[3])
		h := m / 60
		m = m % 60
		times[i] = fmt.Sprintf("%02d:%02d:%02d,%02d0", h, m, s, ms)
	}
	times[len(rows)] = songLength + ",000"
	for i := 0; i < len(rows); i++ {
		res = append(res, strconv.Itoa(i+1))
		res = append(res, fmt.Sprintf("%s --> %s", times[i], times[i+1]))
		res = append(res, rows[i][11:])
		if i+1 < len(rows) {
			res = append(res, "")
		}

	}
	return res
}
