package main

func lengthOfLongestSubstring(s string) int {
	longest := 0

	if len(s) <= 1 {
		return len(s)
	}

	for i := 0; i < len(s); i++ {
		charsPresent := make(map[string]bool)
		for j := i; j < len(s); j++ {
			currentChar := string(s[j])
			if !charsPresent[currentChar] {
				charsPresent[currentChar] = true
				if len(charsPresent) > longest {
					longest = len(charsPresent)
				}
			} else {
				break
			}
		}
	}
	return longest
}
