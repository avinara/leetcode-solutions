// SOLUTION NOT CORRECT

package main

func plusOne(digits []int) []int {
	length := len(digits)
	if digits[length-1] < 9 {
		digits[length-1] = digits[length-1] + 1
		return digits
	}

	for i := length - 1; i >= 0; i-- {
		if digits[i] == 9 && (i-1) >= 0 && digits[i-1] < 9 {
			digits[i] = 0
			digits[i-1] = digits[i-1] + 1
			return digits
		}

	}

	return digits
}
