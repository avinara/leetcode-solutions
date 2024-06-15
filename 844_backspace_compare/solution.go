package main

import (
	"fmt"
	"strings"
)

func backspaceCompare(s string, t string) bool {
	hash := "#"
	var resA, resB []string

	for i := 0; i < len(s); i++ {
		if string(s[i]) == hash {
			if len(resA) > 0 {
				resA = resA[:len(resA)-1]
				continue
			}
		} else {
			resA = append(resA, string(s[i]))
		}
	}

	for i := 0; i < len(t); i++ {
		if string(t[i]) == hash {
			if len(resB) > 0 {
				resB = resB[:len(resB)-1]
				continue
			}

		} else {
			resB = append(resB, string(t[i]))
		}
	}

	resultA := strings.Join(resA, "")
	resultB := strings.Join(resB, "")

	fmt.Println(resultA, resultB)

	if resultA != resultB {
		return false
	}
	return true
}
