package main

import "regexp"

func replaceAllDigitsRegExp(input string) string {
	return regexp.MustCompile("[0-9]").ReplaceAllString(input, "#")
}
