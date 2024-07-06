package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^A-Za-z0-9]`)
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
