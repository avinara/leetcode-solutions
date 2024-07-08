package main

import (
	"fmt"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	stringArray := strings.Split(s, "")
	openParan := "("
	closedParan := ")"
	emptyString := ""

	var bracketArray []int

	for i := 0; i < len(stringArray); i++ {
		if stringArray[i] == openParan {
			bracketArray = append(bracketArray, i)
		}
		if stringArray[i] == closedParan {
			if len(bracketArray) > 0 {
				bracketArray = bracketArray[:len(bracketArray)-1]
			} else {
				stringArray[i] = emptyString
			}
		}
	}

	for i := 0; i < len(bracketArray); i++ {
		stringArray[bracketArray[i]] = emptyString
	}

	fmt.Println(strings.Join(stringArray, emptyString))
	return strings.Join(stringArray, emptyString)
}
