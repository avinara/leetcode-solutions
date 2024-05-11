package main

import (
	"math"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	temp := x
	lenX := len(strconv.Itoa(x))
	newNumber := 0
	for i := lenX; temp > 0; i-- {
		newNumber = newNumber + ((temp % 10) * int(math.Pow(10, float64(i-1))))
		temp = temp / 10
	}

	if newNumber == x {
		return true
	}

	return false
}
