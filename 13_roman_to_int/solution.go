package main

func romanToInt(s string) int {
	var current, last, total int
	romanIntMap := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := len(s) - 1; i >= 0; i-- {
		current = romanIntMap[s[i]]
		if current < last {
			total = total - current
		} else {
			total = total + current
		}
		last = current
	}

	return total
}
