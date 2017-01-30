package main

import (
	"fmt"
	"strings"
)

func mainChristmasTree() {
	res := christmasTree(8, 8)
	for k, v := range res {
		fmt.Printf("%02d %q\n", k, v)
	}
}

func christmasTree(levelNum int, levelHeight int) []string {
	res := make([]string, 100)
	res[0] = "*"
	res[1] = "*"
	res[2] = "***"
	p := 3
	for i := 0; i < levelNum; i++ {
		for j := 0; j < levelHeight; j++ {
			res[p] = strings.Repeat("*", i*2+5+j*2)
			p++
		}
	}
	for i := 0; i < levelNum; i++ {
		if levelHeight%2 == 0 {
			res[p] = strings.Repeat("*", levelHeight+1)
		} else {
			res[p] = strings.Repeat("*", levelHeight)
		}
		p++
	}
	return trimTree(res)
}

func trimTree(res []string) []string {
	var h, w int
	for i := 0; i < len(res); i++ {
		if len(res[i]) == 0 {
			break
		}
		if w < len(res[i]) {
			w = len(res[i])
		}
		h++
	}
	r := make([]string, h)
	for i := 0; i < h; i++ {
		prefix := strings.Repeat(" ", (w-len(res[i]))/2)
		r[i] = prefix + res[i] + prefix
	}
	return r
}
